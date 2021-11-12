package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"twitter-app/config"
	"twitter-app/controllers"
	"twitter-app/database"
	"twitter-app/server"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIndexUser(t *testing.T) {
	assert := assert.New(t)
	config.Init("../config/environments/", "test")
	database.Init()
	defer database.Close()
	e := server.Router()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(controllers.IndexUser(c)) {
		assert.Equal(http.StatusOK, rec.Code)
	}
}

var userJSON = `{"name":"mokou","email":"mokou@saru.moko","password":"mokomoko"}`

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	config.Init("../config/environments/", "test")
	database.Init()
	defer database.Close()
	e := server.Router()

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(controllers.CreateUser(c)) {
		assert.Equal(http.StatusCreated, rec.Code)
	}
}

func TestShowUser(t * testing.T) {
	assert := assert.New(t)
	config.Init("../config/environments/", "test")
	database.Init()
	defer database.Close()
	e := server.Router()

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(controllers.IndexUser(c)) {
		assert.Equal(http.StatusOK, rec.Code)
	}
}
