package routes

import (
	"rest-api/handlers"
	"rest-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(server *gin.Engine, eventHandler *handlers.EventHandler) {

	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
	server.GET("/events", eventHandler.GetEvents)
	server.GET("/events/:eventId", eventHandler.GetEventById)
	server.POST("/signup", eventHandler.SignupUser)
	server.POST("/login", eventHandler.Login)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", eventHandler.CreateEvent)
	authenticated.PUT("/events/:eventId", eventHandler.UpdateEvent)
	authenticated.DELETE("/events/:eventId", eventHandler.DeleteEvent)
	authenticated.POST("/events/:eventId/register", eventHandler.CreateEventRegistration)
	authenticated.DELETE("/events/:eventId/cancel", eventHandler.CancelEventRegistration)

}
