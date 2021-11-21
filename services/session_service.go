package services

import (
	"twitter-app/database"
	"twitter-app/models"
)

func SearchUser(ru *models.ReceiveUser) (*models.User, error) {
	db := database.DB()
	u := new(models.User)
	db = db.Where("email = ?", ru.Email).Take(u)
	if db.Error != nil {
		return nil, db.Error
	}
	return u, nil
}
