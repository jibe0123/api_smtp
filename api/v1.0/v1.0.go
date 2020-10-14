package apiv1

import (
	"github.com/gin-gonic/gin"
	auth "github.com/jibe0123/api_smtp/api/v1.0/auth"
	"github.com/jibe0123/api_smtp/api/v1.0/campaigns"
	"github.com/jibe0123/api_smtp/api/v1.0/clients"
	mailinglist "github.com/jibe0123/api_smtp/api/v1.0/mailing-list"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		auth.ApplyRoutes(v1)
		mailinglist.ApplyRoutes(v1)
		clients.ApplyRoutes(v1)
		campaigns.ApplyRoutes(v1)
	}
}
