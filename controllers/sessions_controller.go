package controllers

import (
	"net/http"
	"twitter-app/models"
	"twitter-app/services"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	ru := new(models.ReceiveUser)
	if err := c.Bind(ru); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.SearchUser(ru); err == nil {
		return nil
	} else {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
}

func Logout(c echo.Context) error {
	return nil
}
