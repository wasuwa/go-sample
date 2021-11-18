package services

import (
	"twitter-app/database"
	"twitter-app/models"
)

func AllUser() ([]models.User, error) {
	var users []models.User
	d := database.DB()
	d = d.Find(&users)
	return users, d.Error
}
