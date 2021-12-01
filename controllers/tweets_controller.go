package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"twitter-app/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func IndexTweet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("user_id"))
	println(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	tt, err := services.AllTweet(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, tt, " ")
}

func ShowTweet(c echo.Context) error {
	return nil
}

func CreateTweet(c echo.Context) error {
	return nil
}

func DestroyTweet(c echo.Context) error {
	return nil
}
