package database

import (
	"github.com/url-shortener/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

const DB_URI = "postgres://postgres@localhost/url"

func Init() {
	var err error
	Db, err = gorm.Open(postgres.Open(DB_URI))
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&entities.Shortcut{})
}
