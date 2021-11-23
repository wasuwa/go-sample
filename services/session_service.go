package services

import (
	"twitter-app/database"
	"twitter-app/models"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SearchUser(ru *models.ReceiveUser) (*models.User, error) {
	db := database.DB()
	u := new(models.User)
	db = db.Where("email = ?", ru.Email).Find(u)
	if db.Error != nil {
		return nil, db.Error
	} else if db.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func Login(u *models.User, c echo.Context) error {
	s, err := session.Get("session", c)
	if err != nil {
		return err
	}
	s.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	s.Values["user_id"] = u.ID
	if err = s.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return nil
}

func IsLoggedin(s *sessions.Session, id int) bool {
	if s.Values["user_id"] == uint(id) {
		return true
	}
	return false
}

func ClearSession(c echo.Context, s *sessions.Session) error {
	s.Options = &sessions.Options{MaxAge: -1, Path: "/"}
	if err := s.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return nil
}
