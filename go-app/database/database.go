package database

import (
	"fmt"
	"github.com/jibe0123/api_smtp/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {

	user := os.Getenv("MYSQL_USER")
	pwd := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("DB_HOST")

	dsn := user + ":" + pwd + "@tcp("+ host +")/" + database + "?parseTime=true&charset=utf8"

	var err error

	for i:=1; i <= 3; i++ {

		db, err := gorm.Open("mysql", dsn)

		if err == nil {
			models.Migrate(db)
			fmt.Println("Connected to database")
			return db,nil
		}
		fmt.Println(err)
		time.Sleep(10 * time.Second)
	}

	return nil, err
}