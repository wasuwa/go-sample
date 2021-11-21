package models

type User struct {
	Base
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"unique"`
	Password      string `json:"password"`
	RememberToken string `json:"rememberToken"`
}

type ReceiveUser struct {
	Name     string `validate:"required,max=15"`
	Email    string `validate:"required,max=256,emailType"`
	Password string `validate:"required,min=6"`
}
