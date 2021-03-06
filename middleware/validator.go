package middleware

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &CustomValidator{Validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.Validator.RegisterValidation("emailType", isEmailTypeValid)
	return cv.Validator.Struct(i)
}

func isEmailTypeValid(fl validator.FieldLevel) bool {
	str := `^[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z$`
	r := regexp.MustCompile(str)
	return r.MatchString(fl.Field().String())
}
