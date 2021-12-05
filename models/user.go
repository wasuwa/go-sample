package models

type User struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type ReceiveUser struct {
	Name     string `validate:"required,max=15"`
	Email    string `validate:"required,max=256,emailType"`
	Password string `validate:"required,min=6"`
}

type ResponseUser struct {
	Users *[]User `json:"users"`
}
