package middleware

import (
	"log"
	"net/http"

	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		log.Printf("Missing authorization data")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User unauthorized"})
		return
	}
	parsedToken, err := utils.ValidateToken(token)
	if err != nil {
		log.Printf("User unathorized: %v\n", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User unauthorized"})
		return
	}

	userId := utils.ExtractData(parsedToken)
	context.Set("userId", userId)
	context.Next()
}
