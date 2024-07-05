package routes

import (
	"api.com/middlewares"
	"github.com/gin-gonic/gin"
)

// in the main function server is a pointer, therefore the type of
// paramter server here is a pointer to the gin Engine.
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getAllEvents)

	server.POST("/events", middlewares.Authenticate, createEvent)
	authenticatedGroup := server.Group("/")
	authenticatedGroup.Use(middlewares.Authenticate)
	authenticatedGroup.GET("/events/:id", getEventbyID)
	authenticatedGroup.PUT("/events/:id", updateEvent)
	authenticatedGroup.DELETE("/events/:id", deleteEvent)
	authenticatedGroup.POST("/events/:id/registerusers", addEventUser)
	authenticatedGroup.DELETE("/events/:id/registerusers", deleteEventUser)

	server.GET("/allusers", getAllUsers)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
