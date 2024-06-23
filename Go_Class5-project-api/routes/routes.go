package routes

import "github.com/gin-gonic/gin"

//in the main function server is a pointer, therefore the type of
// paramter server here is a pointer to the gin Engine.
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getAllEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEventbyID)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
