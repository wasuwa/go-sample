package services

import (
	"errors"
	"twitter-app/database"
	"twitter-app/models"

	"golang.org/x/crypto/bcrypt"
)

func AllUser() ([]models.User, error) {
	d := database.DB()
	var users []models.User
	d = d.Find(&users)
	return users, d.Error
}

func FindUser(id int) (*models.User, error) {
	db := database.DB()
	u := new(models.User)
	db = db.Where("id = ?", id).Find(u)
	if db.RowsAffected == 0 {
		return nil, notFoundError()
	}
	return u, db.Error
}

func CreateUser(ru *models.ReceiveUser) (*models.User, error) {
	db := database.DB()
	u := new(models.User)
	if ru.Password != "" {
		var err error
		ru.Password, err = hashPassword(ru.Password)
		if err != nil {
			return nil, err
		}
	}
	bindUser(u, ru)
	db = db.Create(u)
	return u, db.Error
}

func hashPassword(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	return string(h), err
}

func bindUser(u *models.User, ru *models.ReceiveUser) {
	u.Name 		 = ru.Name
	u.Email 	 = ru.Email
	u.Password = ru.Password
}

func notFoundError() error {
	return errors.New("record not found")
}
