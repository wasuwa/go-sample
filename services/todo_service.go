package services

import (
	"errors"
	"twitter-app/database"
	"twitter-app/models"
)

func AllUser() ([]models.User, error) {
	var users []models.User
	d := database.DB()
	d = d.Find(&users)
	return users, d.Error
}

func FindUser(id int) (*models.User, error) {
	u := new(models.User)
	db := database.DB()
	db = db.Where("id = ?", id).Find(u)
	if db.RowsAffected == 0 {
		return nil, notFoundError()
	}
	return u, db.Error
}

func notFoundError() error {
	return errors.New("record not found")
}
