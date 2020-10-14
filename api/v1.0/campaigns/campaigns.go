package campaigns

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/lib/middlewares"
)

func ApplyRoutes(r *gin.RouterGroup) {
	campaigns := r.Group("/campaigns")
	{
		campaigns.POST("/", middlewares.Authorized, create)
	}
}
