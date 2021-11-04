package models

import "twitter-app/database"

type User struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
}

func (u *User) All() []User {
	var users []User
	d := database.GetDB()
	d.Find(&users)
	return users
}

func (u *User) Create() {
	d := database.GetDB()
	d.Create(u)
}
