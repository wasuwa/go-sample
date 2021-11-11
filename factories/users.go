package factories

import (
	"strconv"
	"twitter-app/config"
	"twitter-app/database"
	"twitter-app/models"

	"gorm.io/gorm"
)

var db *gorm.DB

func Seed() {
	config.Init("../config/environments/", "test")
	database.Init()
	db = database.DB()

	db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < 10; i++ {
			u := &models.User{
				Name:     "mokou",
				Email:    "mokou" + strconv.Itoa(i) + "@gmail.com",
				Password: "password",
			}
			tx.Create(u)
		}
		return nil
	})
}

func Rollback() {
	db.Rollback()
}
