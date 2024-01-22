package db

import (
	"fmt"
	"go_backend/entity"
	"go_backend/model"
)

func UserReview(uname, chainAccount string) bool {
	//注册信息查重
	rows, errs := DB.Query("select *from tb_user where username=? or chain_account=?", uname, chainAccount)
	CheckErr(errs)
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}

func UsersInset(entity entity.UserEntity) {
	//注册信息入库
	stmt, err := DB.Prepare("insert into tb_user(username,password,chain_account,name,card_id,phone,email,seal_img) " +
		"values(?,?,?,?,?,?,?,?)")
	CheckErr(err)
	re, err := stmt.Exec(entity.Username, entity.Password, entity.ChainAccount, entity.Name, entity.CardId, entity.Phone, entity.Email, entity.SealImg)
	CheckErr(err)
	fmt.Println(re.LastInsertId())
}
func UsersQuery(uname, pwd string) bool {
	//登陆验证
	rows, errs := DB.Query("select *from tb_user where username=? and password=?", uname, pwd)
	CheckErr(errs)
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}
func GetUserMessage(uname string) (model.GetUserFrom, error) {
	//查询个人信息
	fmt.Println("uname:", uname)
	from := &model.GetUserFrom{}
	errs := DB.QueryRow("select username,chain_account,name,card_id,phone,email from tb_user where username=?",
		uname).Scan(&from.Username, &from.ChainAccount, &from.Name, &from.CardId, &from.Phone, &from.Email)
	return *from, errs
}
func GetSealMessage(uname string) (string, error) {
	//查询印章信息
	var sealstr string
	errs := DB.QueryRow("select seal_img from tb_user where username=?",
		uname).Scan(&sealstr)
	return sealstr, errs
}
func GetChainAccount(uname string) string {
	//查询账户地址
	var chain_account string
	errs := DB.QueryRow("select chain_account from tb_user where username=?",
		uname).Scan(&chain_account)
	CheckErr(errs)
	return chain_account
}
func GetCardId(uname string) string {
	//查询身份证
	var cardId string
	errs := DB.QueryRow("select card_id from tb_user where username=?",
		uname).Scan(&cardId)
	CheckErr(errs)
	return cardId
}
