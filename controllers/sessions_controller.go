package controllers

import (
	"net/http"
	"twitter-app/models"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	ru := new(models.ReceiveUser)
	if err := c.Bind(ru); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Logout(c echo.Context) error {
	return nil
}
