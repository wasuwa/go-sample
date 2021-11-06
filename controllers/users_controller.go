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
	var u models.User
	i, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Update(i)
	return c.JSON(http.StatusNoContent, nil)
}

// func Update(c echo.Context) error {
// 	var user models.User
// 	i := c.Param("id")
// 	u := user.Find(i)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	return c.JSON(http.StatusCreated, u)
// }

// func Destroy(c echo.Context) error {
// 	i := c.Param("id")
// 	n, _ := strconv.Atoi(i)
// 	models.Users = append(models.Users[:n-1], models.Users[n])
// 	return c.JSON(http.StatusNoContent, nil)
// }
