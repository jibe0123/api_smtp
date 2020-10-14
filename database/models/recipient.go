package models

import (
	common "github.com/jibe0123/api_smtp/lib"
	"github.com/jinzhu/gorm"
)

type Recipient struct {
	gorm.Model
	Name             string `gorm:"size:255"`
	Email            string `gorm:"size:255"`
	MailingList      MailingList
	RecipientsListID uint64
}

func (r *Recipient) Serialize() common.JSON {
	return common.JSON{
		"id":    r.ID,
		"Email": r.Email,
	}
}
