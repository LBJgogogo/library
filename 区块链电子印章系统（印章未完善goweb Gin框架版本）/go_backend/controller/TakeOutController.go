package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TakeOutController struct {
}

func (toc TakeOutController) Login(engine *gin.Engine) {
	engine.Use(toc.CorsMiddleware())
	//POST
	engine.POST("/user/register", toc.UserRegister)
	engine.POST("/user/login", toc.UserLogin)
	engine.POST("/seal/signature", toc.SealSignature)
	engine.POST("/seal/verify", toc.SealVerify)
	//GET
	engine.GET("/user/info", toc.UserMessage)
	engine.GET("/user/seal", toc.SealMessage)
	engine.GET("/seal/record", toc.SealRecordMessage)
}

func (toc TakeOutController) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
