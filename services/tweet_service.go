package services

import (
	"twitter-app/database"
	"twitter-app/models"

	"gorm.io/gorm"
)

func AllTweet(id int) (*[]models.Tweet, error) {
	db := database.DB()
	tt := new([]models.Tweet)
	db = db.Where("user_id = ?", id).Find(tt)
	if db.Error != nil {
		return nil, db.Error
	} else if db.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return tt, nil
}
