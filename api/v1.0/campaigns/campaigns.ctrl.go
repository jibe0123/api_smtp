package campaigns

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/api_smtp/database/models"
	"github.com/jinzhu/gorm"
	"time"
)

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Name      string    `json:"name" binding:"required"`
		DateStart time.Time `json:"date_start" binding:"required"`
		Subject   string    `json:"subject" binding:"required"`
		Content   string    `json:"content" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// user := c.MustGet("user").(models.User)
	campaigns := models.Campaigns{Name: requestBody.Name, DateStart: requestBody.DateStart, Subject: requestBody.Subject, Content: requestBody.Content}
	db.NewRecord(campaigns)
	db.Create(&campaigns)
	c.JSON(200, campaigns.Serialize())
}
