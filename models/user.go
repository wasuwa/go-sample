package models

import (
	"errors"
	"time"
	"twitter-app/database"
)

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name" validate:"required,max=15"`
	Email    string `json:"email" validate:"required,max=256,emailType"`
	Password string `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
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
