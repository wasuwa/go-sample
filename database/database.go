package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var d *gorm.DB

func Init() {
	var err error
	dsn := "host=localhost user=suwayouta dbname=twitter sslmode=disable"
	d, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return d
}

func Close() {
	d, _ := d.DB()
	d.Close()
}
