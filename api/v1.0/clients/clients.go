package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/lib/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	clients := r.Group("/clients")
	{
		clients.POST("/", middlewares.Authorized, create)
		clients.GET("/", middlewares.Authorized, list)
	}
}
