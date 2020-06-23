package models

import (
	common "github.com/jibe0123/api_smtp/lib"
)

type User struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Email    string
	Password string
}

func (p User) Serialize() common.JSON {
	return common.JSON{
		"id":       p.ID,
		"Username": p.Username,
		"Email":    p.Email,
	}
}
