package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_backend/model"
	"go_backend/service"
	"go_backend/utils/Token"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (toc TakeOutController) UserRegister(c *gin.Context) {
	var registerFrom model.UserFrom

	if err := c.ShouldBindJSON(&registerFrom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	log.Println(registerFrom)
	if service.Register(registerFrom) {
		// 返回成功的响应
		c.JSON(http.StatusOK, gin.H{"message": "register ok"})
	} else {
		// 返回失败的响应
		c.JSON(http.StatusForbidden, gin.H{"message": "register fail"})
	}
}

func (toc TakeOutController) UserLogin(c *gin.Context) {
	var loginFrom model.LoginFrom
	if err := c.ShouldBindJSON(&loginFrom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	log.Println(loginFrom)
	if service.IsLogin(loginFrom) {
		// 生成JWT Token
		tokenString, err := Token.GenerateToken(loginFrom.Username, strconv.FormatInt(time.Now().Unix(), 10))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		// 返回成功的响应
		c.JSON(http.StatusOK, gin.H{"message": "login ok", "token": tokenString})
	} else {
		// 返回失败的响应
		c.JSON(http.StatusForbidden, gin.H{"message": "login fail"})
	}
}
func (toc TakeOutController) UserMessage(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	username, err := Token.GetLoginName(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token has expired"})
		return
	}
	data, err := service.GetUserMessage(username)
	if err == nil {
		// 填充结构体并返回成功的响应
		response := model.MessageResponse{
			Message: "Request ok ",
			Data:    data,
		}
		fmt.Println(data)
		c.JSON(http.StatusOK, response)
	} else {
		// 返回失败的响应
		c.JSON(http.StatusForbidden, gin.H{"message": "Request fail"})
	}
}
func (toc TakeOutController) SealMessage(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	username, err := Token.GetLoginName(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token has expired"})
		return
	}
	data, err := service.GetSealMessage(username)
	if err == nil {
		response := model.MessageResponse{
			Message: "Request ok ",
			Data:    data,
		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Request fail"})
	}
}
