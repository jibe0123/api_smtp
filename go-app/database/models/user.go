package models

import (
	common "github.com/jibe0123/api_smtp/lib"
	"github.com/jinzhu/gorm"
)

// User data model
type User struct {
	gorm.Model
	Username     string
	Email  string
	PasswordHash string
	ClientID uint
}

func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"display_name": u.Email,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
	u.Email = m["display_name"].(string)
}