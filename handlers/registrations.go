package handlers

import (
	"log"
	"net/http"
	"strconv"

	"rest-api/events"

	"github.com/gin-gonic/gin"
)

func (eventHandler *EventHandler) CreateEventRegistration(context *gin.Context) {
	ctx := context.Request.Context()
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		log.Printf("Unable to parse event ID: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse data"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := events.GetEventById(ctx, eventId, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to fetch event: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	if event.ID != eventId {
		log.Printf("Unknown event ID: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unknown event id"})
		return
	}
	err = events.CreateRegistration(ctx, eventId, userId, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to insert data: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not insert data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})

}

func (eventHandler *EventHandler) CancelEventRegistration(context *gin.Context) {
	ctx := context.Request.Context()
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		log.Printf("Unable to parse event ID: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := events.GetEventById(ctx, eventId, eventHandler.dbPool)
	userId := context.GetInt64("userId")
	if err != nil {
		log.Printf("Unable to fetch event: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch event"})
		return
	}
	if event.ID != eventId {
		log.Printf("Unknown event ID: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unknown event ID"})
		return
	}
	err = events.CancelRegistration(ctx, eventId, userId, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to delete event registration: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cancel unsuccesfull"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Cancel successful"})
}
