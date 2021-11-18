package models

import "time"

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
