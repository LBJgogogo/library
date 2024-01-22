package Token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	expireTime  = 30 * time.Minute
	tokenSecret = "hnkj"
)

// GenerateToken generates a new token with username and loginTime claims
func GenerateToken(username, loginTime string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"loginName": username,
		"loginTime": loginTime,
		"exp":       time.Now().Add(expireTime).Unix(),
	})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil
}

// VerifyToken verifies the token and returns whether it is valid or not
func VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err == nil && token.Valid {
		log.Print("通过拦截器")
		return true
	}

	return false
}

// GetLoginName extracts the loginName claim from the token
func GetLoginName(tokenString string) (string, error) {
	if VerifyToken(tokenString) {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		})

		if err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if loginName, ok := claims["loginName"].(string); ok {
					return loginName, nil
				}
			}
		}
	}
	return "", fmt.Errorf("failed to get loginName from token")
}
