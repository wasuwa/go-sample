package models

import (
	"errors"
	"twitter-app/database"

	"github.com/go-playground/validator/v10"
)

type (
	User struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (u *User) All() ([]User, error) {
	var users []User
	d := database.GetDB()
	d = d.Find(&users)
	return users, d.Error
}

func (u *User) Create() error {
	d := database.GetDB()
	d = d.Create(u)
	return d.Error
}

func (u *User) Find(id int) error {
	d := database.GetDB()
	d = d.Where("id = ?", id).Take(u)
	return d.Error
}

func (u *User) Update(id int) error {
	d := database.GetDB()
	d = d.Debug().Where("id = ?", id).Updates(u)
	if d.RowsAffected == 0 {
		err := errors.New("record not found")
		return err
	}
	return d.Error
}

func (u *User) Destroy(id int) error {
	d := database.GetDB()
	d = d.Delete(u, id)
	if d.RowsAffected == 0 {
		err := errors.New("record not found")
		return err
	}
	return d.Error
}
