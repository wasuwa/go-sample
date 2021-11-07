package controllers

import (
	"net/http"
	"strconv"
	"twitter-app/models"

	"github.com/labstack/echo/v4"
)

func IndexUser(c echo.Context) error {
	var u models.User
	users, err := u.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	u := new(models.User)
	err := c.Bind(u)
	if err != nil {
		return err
	}
	err = u.Create()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func ShowUser(c echo.Context) error {
	var u models.User
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = u.Find(i)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func UpdateUser(c echo.Context) error {
	u := new(models.User)
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = c.Bind(u)
	if err != nil {
		return err
	}
	err = u.Update(i)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, nil)
}

func DestroyUser(c echo.Context) error {
	var u models.User
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = u.Destroy(i)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, nil)
}
