package routes

import (
	"net/http"
	"strconv"

	"api.com/models"
	"api.com/utils"
	"github.com/gin-gonic/gin"
)

func getAllEvents(context *gin.Context) {
	// context.JSON(http.StatusOK, gin.H{"message": "vai toma"})
	events, error := models.GetAllEvents()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not stablish connection"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	
	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	event.UserID = userId

	error := event.Save()

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func getEventbyID(context *gin.Context) {
	// Hexadecimal system 10
	// 64 bytes
	eventId, error := strconv.ParseInt(context.Param("id"), 10, 64)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event Id"})
		return
	}
	// context.JSON(http.StatusOK, events)

	event, error := models.GetEventByID(eventId)

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventId, error := strconv.ParseInt(context.Param("id"), 10, 64)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event Id"})
		return
	}

	_, error = models.GetEventByID(eventId)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	var updatedEventStruct models.Event
	error = context.ShouldBindJSON(&updatedEventStruct)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event"})
		return
	}

	updatedEventStruct.ID = eventId
	error = updatedEventStruct.Update()

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
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

	event.Delete()

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully!"})
}
