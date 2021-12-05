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

func CreateTweet(rt *models.ReceiveTweet, id uint) (*models.Tweet, error) {
	db := database.DB()
	t := bindTweet(rt, id)
	db = db.Debug().Create(t)
	if db.Error != nil {
		return nil, db.Error
	}
	return t, nil
}
func DestroyTweet(id int) error {
	db := database.DB()
	t := new(models.Tweet)
	db = db.Debug().Delete(t, id)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func bindTweet(rt *models.ReceiveTweet, id uint) *models.Tweet {
	t := new(models.Tweet)
	t.UserID = id
	t.Content = rt.Content
	return t
}
