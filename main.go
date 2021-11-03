package main

import (
	"twitter-app/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// route
	e.GET("/users", controllers.Index)
	e.GET("/users/:id", controllers.Show)
	// e.GET("/users/new", controllers.New)
	e.POST("/users", controllers.Create)
	// e.GET("/users/:id/edit", controllers.Edit)
	e.PUT("/users/:id", controllers.Update)
	// e.DELETE("/users/:id", controllers.Destroy)

	e.Logger.Fatal(e.Start(":1323"))
}
