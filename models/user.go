package models

import (
	"errors"
	"time"
	"twitter-app/database"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey,autoincrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ReceiveUser struct {
	Name     string `json:"name"     validate:"required,max=15"`
	Email    string `json:"email"    validate:"required,max=256,emailType"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *User) All() ([]User, error) {
	var users []User
	d := database.DB()
	d = d.Find(&users)
	return users, d.Error
}

func (u *User) Create() error {
	var err error
	d := database.DB()
	u.Password, err = u.hashPassword()
	if err != nil {
		return err
	}
	d = d.Create(u)
	return d.Error
}

func (u *User) Find(id int) error {
	d := database.DB()
	d = d.Where("id = ?", id).Take(u)
	return d.Error
}

func (u *User) Update(id int) error {
	var err error
	d := database.DB()
	u.Password, err = u.hashPassword()
	if err != nil {
		return err
	}
	d = d.Where("id = ?", id).Updates(u)
	if d.RowsAffected == 0 {
		err = errors.New("record not found")
		return err
	}
	return d.Error
}

func (u *User) Destroy(id int) error {
	d := database.DB()
	d = d.Delete(u, id)
	if d.RowsAffected == 0 {
		err := errors.New("record not found")
		return err
	}
	return d.Error
}

func (r *ReceiveUser) BindUser(u *User) {
	u.Name 		 = r.Name
	u.Email 	 = r.Email
	u.Password = r.Password
}

func (u *User) hashPassword() (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	return string(h), err
}
