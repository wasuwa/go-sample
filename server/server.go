package server

import (
	"twitter-app/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() (e *echo.Echo) {
	e = echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", controllers.UserIndex)
	// e.GET("/users/:id", controllers.Show)
	// e.GET("/users/new", controllers.New)
	e.POST("/users", controllers.UserCreate)
	// e.GET("/users/:id/edit", controllers.Edit)
	// e.PUT("/users/:id", controllers.Update)
	// e.DELETE("/users/:id", controllers.Destroy)

	return
}
