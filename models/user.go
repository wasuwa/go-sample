package models

import (
	"twitter-app/database"
)

type User struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
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
	u.Id = id
	d := database.GetDB()
	d = d.Take(u)
	return d.Error
}

func (u *User) Update(id int) error {
	u.Id = id
	d := database.GetDB()
	d = d.Updates(u)
	return d.Error
}
