package controllers

import (
	"errors"
	"net/http"
	"twitter-app/models"
	"twitter-app/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {
	ru := new(models.ReceiveUser)
	if err := c.Bind(ru); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	u, err := services.SearchUser(ru)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if u.Password != ru.Password {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("password is incorrect"))
	}
	return nil
}

func Logout(c echo.Context) error {
	return nil
}
