package mailing_list

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/lib/middlewares"
)

func ApplyRoutes(r *gin.RouterGroup) {
	list := r.Group("/mailing-list")
	{
		list.POST("/", middlewares.Authorized, create)
		list.GET("/", middlewares.Authorized, read)
	}
}
