package main

import (
	"net/http"
	"twitter-app/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// route
	e.GET("/users", Index)
	// e.GET("/users/:id", Show)
	// e.GET("/users/new", New)
	e.POST("/users", Create)
	// e.GET("/users/:id/edit", Edit)
	// e.GET("/users/:id/edit", Put)
	// e.DELETE("/users/:id", Delete)

	e.Logger.Fatal(e.Start(":1323"))
}

var users []models.User

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func Create(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	users = append(users, *u)
	return c.JSON(http.StatusCreated, u)
}
