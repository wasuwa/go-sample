package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"twitter-app/config"
)

var db *gorm.DB

func Init() {
	var err error
	c := config.Config()
	db, err = gorm.Open(postgres.Open(c.GetString("db.url")))
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return db
}

func Close() {
	d, _ := db.DB()
	d.Close()
}

func SetupTestDB() (*gorm.DB, func()) {
	config.Init("../config/environments/", "test")
	Init()
	db = db.Begin()
	return db, func() {
		db.Rollback()
		Close()
	}
}
