package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// route
	// e.GET("/users", Index)
	// e.GET("/users/:id", Show)
	// e.GET("/users/new", New)
	// e.POST("/users", Create)
	// e.GET("/users/:id/edit", Edit)
	// e.GET("/users/:id/edit", Put)
	// e.DELETE("/users/:id", Delete)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "hello world" + id)
}
