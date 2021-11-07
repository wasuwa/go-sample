package models

import (
	"errors"
	"twitter-app/database"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

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
