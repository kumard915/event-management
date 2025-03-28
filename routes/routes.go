package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/events", CreateEvent)
	server.PUT("/events/:id", UpdateEvent)    // Add update route
	server.DELETE("/events/:id", DeleteEvent) // Add delete route
}
