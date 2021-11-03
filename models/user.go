package models

import "strconv"

type User struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
}

var Users []User

func (u *User) Find(id string) *User {
	i, _ := strconv.Atoi(id)
	return &Users[i-1]
}
