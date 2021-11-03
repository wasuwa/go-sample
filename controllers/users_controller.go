package controllers

import (
	"net/http"
	"strconv"
	"twitter-app/models"

	"github.com/labstack/echo/v4"
)

var users []models.User

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func Create(c echo.Context) error {
	var id int
	if users == nil {
		id = 1
	} else {
		id = users[len(users) - 1].Id + 1
	}
	u := models.User{Id: id}
	if err := c.Bind(&u); err != nil {
		return err
	}
	users = append(users, u)
	return c.JSON(http.StatusCreated, u)
}

func Show(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	u := users[i - 1]
	return c.JSON(http.StatusOK, u)
}

func Update(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	u := &users[i - 1]
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
