package services

import (
	"twitter-app/database"
	"twitter-app/models"

	"gorm.io/gorm"
)

func AllTweet(id int) (*models.ResponseTweet, error) {
	db := database.DB()
	tt := new([]models.Tweet)
	db = db.Where("user_id = ?", id).Find(tt)
	if db.Error != nil {
		return nil, db.Error
	} else if db.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	rt := new(models.ResponseTweet)
	rt.Tweets = tt
	return rt, nil
}
