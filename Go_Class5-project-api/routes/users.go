package routes

import (
	"net/http"

	"api.com/models"
	"github.com/gin-gonic/gin"
)

func getAllUsers(context *gin.Context) {
	// context.JSON(http.StatusOK, gin.H{"message": "vai toma"})
	events, error := models.GetAllUsers()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not stablish connection"})
		return
	}
	context.JSON(http.StatusOK, events)
}


func signup(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, user)
}

func login(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Hashing comparison error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}