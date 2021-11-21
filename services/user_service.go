package services

import (
	"twitter-app/database"
	"twitter-app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AllUser() ([]models.User, error) {
	db := database.DB()
	var users []models.User
	db = db.Find(&users)
	if db.Error != nil {
		return nil, db.Error
	} else if db.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return users, nil
}

func FindUser(id int) (*models.User, error) {
	db := database.DB()
	u := new(models.User)
	db = db.Where("id = ?", id).Find(u)
	if db.Error != nil {
		return nil, db.Error
	} else if db.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func CreateUser(ru *models.ReceiveUser) (*models.User, error) {
	var err error
	ru.Password, err = hashPassword(ru.Password)
	if err != nil {
		return nil, err
	}
	db := database.DB()
	u := new(models.User)
	bindUser(u, ru)
	db = db.Create(u)
	if db.Error != nil {
		return nil, db.Error
	}
	return u, nil
}

func UpdateUser(ru *models.ReceiveUser, id int) error {
	var err error
	ru.Password, err = hashPassword(ru.Password)
	if err != nil {
		return err
	}
	db := database.DB()
	u := new(models.User)
	bindUser(u, ru)
	db = db.Where("id = ?", id).Updates(u)
	 if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DestroyUser(id int) error {
	u := new(models.User)
	db := database.DB()
	db = db.Delete(u, id)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
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