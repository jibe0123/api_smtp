package mailing_list

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/database/models"
	common "github.com/jibe0123/api_smtp/lib"
	"github.com/jinzhu/gorm"
)

// Clients type alias
type MailingList = models.MailingList

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Name string `json:"name" binding:"required"`
	}
	var requestBody RequestBody
	var client models.Clients

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	user := c.MustGet("user").(models.User)
	db.First(&client, user.ClientID)

	fmt.Printf("USER  : %#v\n", user)
	fmt.Printf("CLIENT : %#v\n", client)
	mailingList := models.MailingList{Name: requestBody.Name, ClientID: client.ID}
	db.NewRecord(mailingList)
	db.Create(&mailingList)
	c.JSON(200, mailingList.Serialize())
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var MailingList []MailingList

	if cursor == "" {
		if err := db.Limit(10).Order("id desc").Find(&MailingList).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Limit(10).Order("id desc").Where(condition, cursor).Find(&MailingList).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(MailingList)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = MailingList[i].Serialize()
	}

	c.JSON(200, serialized)
}
