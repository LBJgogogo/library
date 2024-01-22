package service

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jinzhu/copier"
	db "go_backend/dao"
	"go_backend/entity"
	"go_backend/model"
	"go_backend/test"
	"go_backend/utils/Seal"
	"image/png"
)

func Register(from model.UserFrom) bool {
	// 校验用户名或账户地址是否重复
	judgment := db.UserReview(from.Username, from.ChainAccount)
	if judgment {
		log.Info("注册失败：用户名或账户地址重复")
		return false
	}
	//上链
	test.AddSealAccount(from.ChainAccount, from.Name, from.CardId)
	// 拷贝对象
	userEntity := &entity.UserEntity{}
	err := copier.Copy(&userEntity, from)
	if err != nil {
		log.Info("拷贝失败", err)
		return false
	}
	// md5加密密码
	hash := md5.Sum([]byte(from.Password))
	md5Str := hex.EncodeToString(hash[:])
	userEntity.Password = md5Str
	// 生成印章（Base64）
	buffer, err := Seal.GenerateSeal(from.Name)
	if err != nil {
		log.Info("图像编码失败", err)
		return false
	}
	// 将图像数据保存到内存中的缓冲区
	var buf bytes.Buffer
	err = png.Encode(&buf, buffer)
	if err != nil {
		log.Info("图像编码失败", err)
		return false
	}
	// 将缓冲区中的数据进行Base64编码
	sealImg := base64.StdEncoding.EncodeToString(buf.Bytes())
	userEntity.SealImg = sealImg
	//入库
	db.UsersInset(*userEntity)
	return true
}

func IsLogin(from model.LoginFrom) bool {
	hash := md5.Sum([]byte(from.Password))
	md5Str := hex.EncodeToString(hash[:])
	if db.UsersQuery(from.Username, md5Str) {
		return true
	}
	return false
}

func GetUserMessage(username string) (model.GetUserFrom, error) {
	return db.GetUserMessage(username)
}

func GetSealMessage(username string) (string, error) {
	return db.GetSealMessage(username)
}
