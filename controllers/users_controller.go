package controllers

import (
	"net/http"
	"twitter-app/models"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var u models.User
	users := u.All()
	return c.JSON(http.StatusOK, users)
}

// func Create(c echo.Context) error {
// 	var id int
// 	if models.Users == nil {
// 		id = 1
// 	} else {
// 		id = models.Users[len(models.Users)-1].Id + 1
// 	}
// 	u := models.User{Id: id}
// 	if err := c.Bind(&u); err != nil {
// 		return err
// 	}
// 	models.Users = append(models.Users, u)
// 	return c.JSON(http.StatusCreated, u)
// }

// func Show(c echo.Context) error {
// 	var user models.User
// 	user.Id = 1
// 	d := database.GetDB()
// 	u := d.First(&user)
// 	return c.JSON(http.StatusOK, u)
// }

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
