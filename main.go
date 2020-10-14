package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/api"
	"github.com/jibe0123/api_smtp/database"
	"github.com/jibe0123/api_smtp/lib/middlewares"
)


func main() {
	port := "8080"
	db, _ := database.Initialize()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(database.Inject(db))
	r.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(r) // apply api router
	r.Run(":" + port)

}