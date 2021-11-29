package server

import (
	"twitter-app/controllers"
	mw "twitter-app/middleware"
	"twitter-app/utils"

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
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(utils.LoadEnv("COOKIE_KEY")))))

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

	// tweet
	e.GET("/users/tweets", controllers.IndexTweet)
	e.GET("/users/tweets/:id", controllers.ShowTweet)
	e.POST("/users/tweets", controllers.CreateTweet)
	e.DELETE("/users/tweets", controllers.DestroyTweet)

	return
}
