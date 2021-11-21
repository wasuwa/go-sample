package services

import (
	"twitter-app/database"
	"twitter-app/models"

	"gorm.io/gorm"
)

func SearchUser(ru *models.ReceiveUser) (*models.User, error) {
	db := database.DB()
	u := new(models.User)
	db = db.Where("email = ?", ru.Email).Find(u)
	if db.Error != nil {
		return nil, db.Error
	} else if db.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}
