package models

import (
	common "github.com/jibe0123/api_smtp/lib"
	"github.com/jinzhu/gorm"
)

type Recipient struct {
	gorm.Model
	Email         string
	MailingListID uint64
}

func (r *Recipient) Serialize() common.JSON {
	return common.JSON{
		"id":    r.ID,
		"Email": r.Email,
	}
}
