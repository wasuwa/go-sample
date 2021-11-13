package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"testing"
	"twitter-app/config"
	"twitter-app/controllers"
	"twitter-app/database"
	"twitter-app/models"
	"twitter-app/server"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setup() (*gorm.DB, func()) {
	config.Init("../config/environments/", "test")
	database.Init()
	db := database.DB()
	db = db.Begin()
	database.SetDB(db)
	return db, func() {
		db.Rollback()
		database.Close()
	}
}

var (
	u = &models.User{
		ID:       0,
		Name:     "mokou",
		Email:    "katou@jyun.iti",
		Password: "futontyan",
	}
	uJSON = `{"name":"god","email":"takada@ken.shi","password":"zetsuen"}`
)

func TestIndexUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := setup()
	defer teardown()

	db.Create(u)

	e := server.Router()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(controllers.IndexUser(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.Contains(rec.Body.String(), "mokou")
	}
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	_, teardown := setup()
	defer teardown()

	e := server.Router()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(uJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(controllers.CreateUser(c)) {
		assert.Equal(http.StatusCreated, rec.Code)
		assert.Contains(rec.Body.String(), "god")
	}
}

func TestShowUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := setup()
	defer teardown()

	db.Create(u)
	db.Find(u)
	id := strconv.Itoa(int(u.ID))

	e := server.Router()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(controllers.ShowUser(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.Contains(rec.Body.String(), id)
	}
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := setup()
	defer teardown()

	db.Create(u)
	db.Find(u)
	id := strconv.Itoa(int(u.ID))

	e := server.Router()
	req := httptest.NewRequest(http.MethodPatch, "/users/:id", strings.NewReader(uJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(controllers.UpdateUser(c)) {
		assert.Equal(http.StatusNoContent, rec.Code)
	}
}

func TestDestroyUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := setup()
	defer teardown()

	db.Create(u)
	db.Find(u)
	id := strconv.Itoa(int(u.ID))

	e := server.Router()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(controllers.DestroyUser(c)) {
		assert.Equal(http.StatusNoContent, rec.Code)
	}
}
