package server

import (
	"twitter-app/controllers"
	mw "twitter-app/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() (e *echo.Echo) {
	e = echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		ContentTypeNosniff: "application/json",
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Validator = mw.NewValidator()

	// users
	e.GET("/users", controllers.IndexUser)
	e.GET("/users/:id", controllers.ShowUser)
	e.POST("/users", controllers.CreateUser)
	e.PATCH("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DestroyUser)

	// sessions
	e.POST("/login", controllers.Login)
	e.DELETE("/logout", controllers.Logout)

	return
}
