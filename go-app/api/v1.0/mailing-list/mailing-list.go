package mailing_list


import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/lib/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	list := r.Group("/mailing-list")
	{
		list.GET("/", middlewares.Authorized, ping)
	}
}
