package mailing_list

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/database/models"
	"github.com/jinzhu/gorm"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Name string `json:"name" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}


	user := c.MustGet("user").(models.User)
	client := db.Model(&user).Association("ClientID")

	fmt.Printf("CLIENTID : %#v\n", user)
	fmt.Printf("CLIENT : %#v\n", client)
	mailingList := models.MailingList{Name: requestBody.Name}
	db.NewRecord(mailingList)
	db.Create(&mailingList)
	c.JSON(200, mailingList.Serialize())
}