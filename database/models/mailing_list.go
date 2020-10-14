package models

import (
	common "github.com/jibe0123/api_smtp/lib"
	"github.com/jinzhu/gorm"
)

type MailingList struct {
	gorm.Model
	ClientID uint
	Name     string
}

func (m *MailingList) Serialize() common.JSON {
	return common.JSON{
		"id":   m.ID,
		"Name": m.Name,
	}
}
