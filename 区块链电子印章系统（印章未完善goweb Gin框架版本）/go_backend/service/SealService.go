package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	db "go_backend/dao"
	"go_backend/entity"
	"go_backend/model"
	"go_backend/test"
	"strings"
)

func SealSignature(from model.SealFrom, uname, chainAccount string) bool {
	code := uuid.New().String() // 获取签名文档uuid编码
	sealEntity := &entity.SealEntity{}
	err := copier.Copy(&sealEntity, from)
	if err != nil {
		log.Info("拷贝失败", err)
		return false
	}
	hash := sha256.Sum256([]byte(from.ImgBase64))
	hashStr := hex.EncodeToString(hash[:])
	//上链
	test.SealSignature(chainAccount, hashStr, code)

	str := strings.Split(from.Filename, ".")
	sealEntity.Filename = str[0]
	sealEntity.Type = str[1]
	sealEntity.Code = code
	sealEntity.Username = uname
	//入库
	db.SealInset(*sealEntity)
	return true
}

func GetRecords(username, chainAccount string) ([]model.SealRecord, error) {
	//链上查询哈希
	hxstr := test.GetAllHash(chainAccount)
	//链上查询签章记录信息
	length := len(hxstr)
	froms := &[]model.SealRecord{}
	for i := 0; i < length; i++ {
		code, time := test.GetSignature(chainAccount, hxstr[i])
		from := db.GetSealRecord(username, code, time, i+1)
		*froms = append(*froms, from)
	}
	return *froms, nil
}
func SealVerify(username, chainAccount string, base64 model.SealBase64) (model.SealVerifyMessage, bool) {
	//链上查询哈希
	hxstr := test.GetAllHash(chainAccount)
	//链上查询签章记录信息
	length := len(hxstr)
	froms := &model.SealVerifyMessage{}
	for i := 0; i < length; i++ {
		code, time := test.GetSignature(chainAccount, hxstr[i])
		result := db.GetVerifyMessage(username, code)
		if result.ImgBase64 == base64.ImgBase64 {
			froms.Name = username
			froms.CardId = db.GetCardId(username)
			froms.ChainAccount = chainAccount
			froms.Filename = db.GetFileName(username)
			froms.Code = code
			froms.Datetime = time
			return *froms, true
		}
	}
	return *froms, false
}
