package main

import (
	"github.com/gin-gonic/gin"
	"go_backend/controller"
	db "go_backend/dao"
	"go_backend/test"
)

func main() {
	//创建服务器引擎
	engine := gin.Default()
	//初始化数据库连接
	err := db.Init()
	db.CheckErr(err)
	//初始化区块链WeBASE平台链接
	test.StartBlockchain()
	//设置信任代理地址
	engine.SetTrustedProxies([]string{"192.168.43.117"})
	// 加载静态文件
	//engine.Static("/static", "./static")
	//// 加载模板
	//engine.LoadHTMLGlob("./views/*")
	registerRouter(engine) //注册路由
	engine.Run(":8080")
}

// 路由设置
func registerRouter(router *gin.Engine) {
	// 控制器
	new(controller.TakeOutController).Login(router) //加载TakeOutController..
}
