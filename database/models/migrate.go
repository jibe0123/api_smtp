package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Clients{}, &MailingList{}, &Campaigns{})
	fmt.Println("Auto Migration has been processed")
}
