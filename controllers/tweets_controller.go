package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"twitter-app/models"
	"twitter-app/services"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func IndexTweet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("user_id"))
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

func CreateTweet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	s, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !services.IsLoggedin(s, id) {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("login required").Error())
	}
	rt := new(models.ReceiveTweet)
	if err := c.Bind(rt); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(rt); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	t, err := services.CreateTweet(rt, uint(id))
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusCreated, t, " ")
}

func DestroyTweet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	s, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !services.IsLoggedin(s, id) {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("login required").Error())
	}
	tid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err = services.DestroyTweet(tid); errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err = services.ClearSession(c, s); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}
