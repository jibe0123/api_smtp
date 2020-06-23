package auth

import (
	"github.com/gin-gonic/gin"
	common "github.com/jibe0123/api_smtp/lib"
)

type JSON = common.JSON

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


