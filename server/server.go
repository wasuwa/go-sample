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

	e.GET("/users", controllers.IndexUser)
	e.GET("/users/:id", controllers.ShowUser)
	// e.GET("/users/new", controllers.New)
	e.POST("/users", controllers.CreateUser)
	// e.GET("/users/:id/edit", controllers.Edit)
	e.PUT("/users/:id", controllers.UpdateUser)
	// e.DELETE("/users/:id", controllers.Destroy)

	return
}
