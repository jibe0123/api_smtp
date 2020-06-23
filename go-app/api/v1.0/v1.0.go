package apiv1

import (
	"github.com/gin-gonic/gin"
	auth "github.com/jibe0123/api_smtp/api/v1.0/Auth"
)


func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		auth.ApplyRoutes(v1)
	}
}