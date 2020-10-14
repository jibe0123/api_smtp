package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/database/models"
	common "github.com/jibe0123/api_smtp/lib"
	"github.com/jinzhu/gorm"
)

// Clients type alias
type Clients = models.Clients

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		CompanyName string `json:"company_name" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var exists Clients
	if err := db.Where("company_name = ?", requestBody.CompanyName).First(&exists).Error; err == nil {
		c.AbortWithStatus(409)
		return
	}

	clients := Clients{
		CompanyName: requestBody.CompanyName,
	}

	db.NewRecord(clients)
	db.Create(&clients)
	c.JSON(200, clients.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var Clients []Clients

	if cursor == "" {
		if err := db.Preload("Users").Limit(10).Order("id desc").Find(&Clients).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("Users").Limit(10).Order("id desc").Where(condition, cursor).Find(&Clients).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(Clients)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = Clients[i].Serialize()
	}

	c.JSON(200, serialized)
}
