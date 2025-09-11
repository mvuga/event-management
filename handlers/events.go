package handlers

import (
	"log"
	"net/http"
	"strconv"

	"rest-api/events"

	"github.com/gin-gonic/gin"
)

func (eventHandler *EventHandler) GetEvents(context *gin.Context) {
	ctx := context.Request.Context()
	events, err := events.GetEvents(ctx, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to get all events: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func (eventHandler *EventHandler) CreateEvent(context *gin.Context) {
	ctx := context.Request.Context()
	var event events.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		log.Printf("Unable to parse data: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	event.UserId = context.GetInt64("userId")
	err = event.CreateEvent(ctx, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to insert data: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not insert data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event:": event})
}

func (eventHandler *EventHandler) GetEventById(context *gin.Context) {
	ctx := context.Request.Context()
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		log.Printf("Unable to parse event ID: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	event, err := events.GetEventById(ctx, eventId, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to fetch event: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func (eventHandler *EventHandler) UpdateEvent(context *gin.Context) {
	ctx := context.Request.Context()
	var updatedEvent events.Event
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		log.Printf("Unable to parse event ID: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	_, err = events.GetEventById(ctx, eventId, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to fetch event: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		log.Printf("Unable to parse data: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse data"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent(ctx, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to update event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Update sucessfull"})

}

func (eventHandler *EventHandler) DeleteEvent(context *gin.Context) {
	ctx := context.Request.Context()
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		log.Printf("Unable parse event ID: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := events.GetEventById(ctx, eventId, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to fetch event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	if event.UserId != context.GetInt64("userId") {
		log.Printf("User unathorized")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User unauthorized"})
		return
	}
	err = events.DeleteEvent(ctx, eventId, eventHandler.dbPool)
	if err != nil {
		log.Printf("Unable to delete event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete successful"})

}
