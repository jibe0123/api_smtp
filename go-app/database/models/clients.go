package models

import (
	common "github.com/jibe0123/api_smtp/lib"
	"github.com/jinzhu/gorm"
)

type Clients struct {
	gorm.Model
	Users        []User `gorm:"foreignkey:ClientID"`
	CompanyName string
}


// Serialize serializes user data
func (c *Clients) Serialize() common.JSON {
	return common.JSON{
		"id":           c.ID,
		"companyName": c.CompanyName,
		"users":       c.Users,
	}
}

func (c *Clients) Read(m common.JSON) {
	c.ID = uint(m["id"].(float64))
	c.CompanyName = m["campany_name"].(string)
}