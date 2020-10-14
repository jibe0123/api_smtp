package models

import (
	common "github.com/jibe0123/api_smtp/lib"
	"github.com/jinzhu/gorm"
)

type Clients struct {
	gorm.Model
	Users        []User `gorm:"foreignkey:ClientID"`
	MailingLists []MailingList
	CompanyName  string
}

func (c *Clients) Serialize() common.JSON {
	return common.JSON{
		"id":          c.ID,
		"companyName": c.CompanyName,
		"mailingList": c.MailingLists,
	}
}
