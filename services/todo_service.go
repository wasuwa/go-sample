package services

import (
	"fmt"
	"twitter-app/database"
	"twitter-app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AllUser() ([]models.User, error) {
	db := database.DB()
	var users []models.User
	db = db.Find(&users)
	if db.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return users, db.Error
}

func FindUser(id int) (*models.User, error) {
	db := database.DB()
	u := new(models.User)
	db = db.Where("id = ?", id).Find(u)
	if db.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return u, db.Error
}

func CreateUser(ru *models.ReceiveUser) (*models.User, error) {
	if ru.Password != "" {
		var err error
		ru.Password, err = hashPassword(ru.Password)
		if err != nil {
			return nil, err
		}
	}
	db := database.DB()
	u := new(models.User)
	bindUser(u, ru)
	fmt.Println(u)
	db = db.Create(u)
	fmt.Println(u)
	return u, db.Error
}

func UpdateUser(ru *models.ReceiveUser, id int) error {
	if ru.Password != "" {
		var err error
		ru.Password, err = hashPassword(ru.Password)
		if err != nil {
			return err
		}
	}
	db := database.DB()
	u := new(models.User)
	bindUser(u, ru)
	db = db.Where("id = ?", id).Updates(u)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
}

func DestroyUser(id int) error {
	u := new(models.User)
	db := database.DB()
	db = db.Debug().Delete(u, id)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
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
