package routes

import (
	"example.com/restapi/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Auth Routes
	authGroup := server.Group("/")
	{
		authGroup.POST("signup", signup)
		authGroup.POST("login", login)
	}

	// Event Routes (Protected Routes with Authentication Middleware)
	eventGroup := server.Group("/events")
	{
		eventGroup.GET("", GetEvents)
		eventGroup.GET("/:id", GetEvent)
		eventGroup.POST("", middleware.Authenticate, CreateEvent)
		eventGroup.PUT("/:id", UpdateEvent)
		eventGroup.DELETE("/:id", DeleteEvent)

	}
}
