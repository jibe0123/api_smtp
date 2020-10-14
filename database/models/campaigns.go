package models

import (
	common "github.com/jibe0123/api_smtp/lib"
	"time"
)

type Campaigns struct {
	ID               uint64 `gorm:"primary_key"`
	Name             string `gorm:"size:255"`
	DateStart        time.Time
	OrganismId       uint64
	Subject          string `gorm:"size:255"`
	Content          string `gorm:"size:1023"`
	RecipientsListId uint64
}

func (c *Campaigns) Serialize() common.JSON {
	return common.JSON{
		"id":        c.ID,
		"Name":      c.Name,
		"DateStart": c.DateStart,
		"Subject":   c.Subject,
		"Content":   c.Content,
	}
}
