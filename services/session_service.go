package services

import (
	"twitter-app/database"
	"twitter-app/models"
)

func SearchUser(ru *models.ReceiveUser) error {
	db := database.DB()
	u  := new(models.User)
	db = db.Where("email = ?", ru.Email).Take(u)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
