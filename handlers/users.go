package handlers

import (
	"log"
	"net/http"

	"rest-api/user"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func (eventHandler *EventHandler) SignupUser(context *gin.Context) {
	ctx := context.Request.Context()
	var usr user.User
	err := context.ShouldBindJSON(&usr)

	if err != nil {
		log.Printf("Unable to parse data: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	err = usr.Create(ctx, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to signup user: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not signup user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (eventHandler *EventHandler) Login(context *gin.Context) {
	ctx := context.Request.Context()
	var usr user.User
	err := context.ShouldBindJSON(&usr)
	if err != nil {
		log.Printf("Unable to parse data: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	err = usr.ValidateCredentials(ctx, eventHandler.dbPool)
	if err != nil {
		log.Printf("Invalid credentials: %v\n", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	jwt, err := utils.GenerateJWT(usr.Email, usr.ID)

	if err != nil {
		log.Printf("Unable to authenticate user: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt})

}
