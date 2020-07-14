package db

import (
	"log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    _ "github.com/mattn/go-sqlite3"
)

var instance *gorm.DB

// Get gets opened db instance
func Get() *gorm.DB {
	if instance != nil {
		return instance
	}

	db, err := gorm.Open("postgres", "user=itotatsuhiko dbname=go_server sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	instance = db
	return instance
}
