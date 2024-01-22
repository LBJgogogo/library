package db

import (
	"go_backend/entity"
	"go_backend/model"
	"log"
)

func SealInset(entity entity.SealEntity) {
	//签章信息入库
	stmt, errs := DB.Prepare("insert into tb_seal(code,filename,type,username,img_base64) " +
		"values(?,?,?,?,?)")
	CheckErr(errs)
	_, errs = stmt.Exec(entity.Code, entity.Filename, entity.Type, entity.Username, entity.ImgBase64)
	CheckErr(errs)
	log.Print("签章信息入库成功")
}
func GetSealRecord(username, code, time string, id int) model.SealRecord {
	//查询签章记录
	froms := &model.SealRecord{}
	froms.Id = id
	froms.Code = code
	froms.Datetime = time
	errs := DB.QueryRow("select filename,type,img_base64 from tb_seal where code=? and username=?", code, username).
		Scan(&froms.Filename, &froms.Type, &froms.ImgBase64)
	CheckErr(errs)
	return *froms
}

func GetVerifyMessage(username, code string) model.SealBase64 {
	//查询签章记录
	base64 := &model.SealBase64{}
	errs := DB.QueryRow("select img_base64 from tb_seal where code=? and username=?", code, username).
		Scan(&base64.ImgBase64)
	CheckErr(errs)
	return *base64
}
func GetFileName(uname string) string {
	//查询文件名
	var fileName string
	errs := DB.QueryRow("select filename from tb_seal where username=?",
		uname).Scan(&fileName)
	CheckErr(errs)
	return fileName
}
