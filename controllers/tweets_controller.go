package controllers

import (
	"errors"
	"net/http"
	"twitter-app/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func IndexTweet(e echo.Context) error {
	tt, err := services.AllTweet()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.JSONPretty(http.StatusOK, tt, " ")
}

func ShowTweet(e echo.Context) error {
	return nil
}

func CreateTweet(e echo.Context) error {
	return nil
}

func DestroyTweet(e echo.Context) error {
	return nil
}
