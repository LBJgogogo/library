package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	DB  *sqlx.DB
	err error
)

// 初始化
func Init() error {
	DB, err = sqlx.Open("mysql", "root:root@tcp(localhost:3306)/test")
	CheckErr(err)
	CheckErr(DB.Ping())
	//测试连接
	log.Println("数据库连接成功")
	return nil
}

func CheckErr(errs error) {
	if errs != nil {
		panic(errs)
	}
}
