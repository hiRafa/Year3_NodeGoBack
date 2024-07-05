package routes

import (
	"net/http"
	"strconv"

	"api.com/models"
	"github.com/gin-gonic/gin"
)

func addEventUser(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, error := strconv.ParseInt(context.Param("id"), 10, 64)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event Id"})
		return
	}

	event, error := models.GetEventByID(eventId)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	error = event.AddEventUser(userId)

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func deleteEventUser(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, error := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	event.DeleteEventUser(userId)
	
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User deleted!"})
}
