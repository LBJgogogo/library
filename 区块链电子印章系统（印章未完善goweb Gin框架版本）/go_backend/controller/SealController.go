package controller

import (
	"github.com/gin-gonic/gin"
	db "go_backend/dao"
	"go_backend/model"
	"go_backend/service"
	"go_backend/utils/Token"
	"log"
	"net/http"
)

func (toc TakeOutController) SealSignature(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	username, err := Token.GetLoginName(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token has expired"})
		return
	}

	var SealFrom model.SealFrom
	if err = c.ShouldBindJSON(&SealFrom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	log.Println(SealFrom)
	chainAccount := db.GetChainAccount(username)
	if service.SealSignature(SealFrom, username, chainAccount) {
		// 返回成功的响应
		c.JSON(http.StatusOK, gin.H{"message": "register ok"})
	} else {
		// 返回失败的响应
		c.JSON(http.StatusForbidden, gin.H{"message": "register fail"})
	}
}
func (toc TakeOutController) SealRecordMessage(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	username, err := Token.GetLoginName(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token has expired"})
		return
	}
	chainAccount := db.GetChainAccount(username)
	data, err := service.GetRecords(username, chainAccount)
	if err == nil {
		// 填充结构体并返回成功的响应
		response := model.MessageResponse{
			Message: "Request ok ",
			Data:    data,
		}
		c.JSON(http.StatusOK, response)
	} else {
		// 返回失败的响应
		c.JSON(http.StatusForbidden, gin.H{"message": "Request fail"})
	}
}
func (toc TakeOutController) SealVerify(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	username, err := Token.GetLoginName(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token has expired"})
		return
	}
	var SealBase64 model.SealBase64
	if err = c.ShouldBindJSON(&SealBase64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	chainAccount := db.GetChainAccount(username)
	data, bl := service.SealVerify(username, chainAccount, SealBase64)
	if bl {
		// 填充结构体并返回成功的响应
		response := model.MessageResponse{
			Message: "verify ok ",
			Data:    data,
		}
		log.Println("验签成功")
		c.JSON(http.StatusOK, response)
	} else {
		log.Println("验签失败")
		c.JSON(http.StatusNoContent, gin.H{"message": "verify fail"})
	}
}
