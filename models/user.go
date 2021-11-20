package models

type User struct {
	Base
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
}

type ReceiveUser struct {
	Name     string `json:"name"     validate:"required,max=15"`
	Email    string `json:"email"    validate:"required,max=256,emailType"`
	Password string `json:"password" validate:"required,min=6"`
}
