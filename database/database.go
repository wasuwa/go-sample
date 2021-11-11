package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"twitter-app/config"
)

var d *gorm.DB

func Init() {
	var err error
	c := config.Config()
	d, err = gorm.Open(postgres.Open(c.GetString("db.url")))
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return d
}

func Close() {
	d, _ := d.DB()
	d.Close()
}
