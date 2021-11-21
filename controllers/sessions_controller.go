package controllers

import (
	"errors"
	"net/http"
	"twitter-app/models"
	"twitter-app/services"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {
	ru := new(models.ReceiveUser)
	if err := c.Bind(ru); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	u, err := services.SearchUser(ru)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(ru.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("password is incorrect").Error())
	}
	// TODO: ログイン処理をする
	if err := services.Login(u, c); err != nil {
		
	}
	// TODO: user{}ではなく、sessionの情報を返すべき
	return c.JSONPretty(http.StatusCreated, u, " ")
}

func Logout(c echo.Context) error {
	return nil
}
