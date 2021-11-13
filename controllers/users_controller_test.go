package controllers_test

import (
	"net/http"
	"net/http/httptest"
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


func SetTransaction() *gorm.DB {
	db := database.DB()
	db = db.Begin()
	database.SetDB(db)
	return db
}

var (
	userJSON = `{"name":"mokou","email":"katou@jyun.iti","password":"futontyan"}`
)

func TestIndexUser(t *testing.T) {
	assert := assert.New(t)
	config.Init("../config/environments/", "test")
	database.Init()
	db := SetTransaction()
	defer database.Close()
	defer db.Rollback()

	u := &models.User{
		Name: "mokou",
		Email: "katou@jyun.iti",
		Password: "futontyan",
	}
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
	config.Init("../config/environments/", "test")
	database.Init()
	db := SetTransaction()
	defer database.Close()
	defer db.Rollback()

	e := server.Router()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(controllers.CreateUser(c)) {
		assert.Equal(http.StatusCreated, rec.Code)
		assert.Contains(rec.Body.String(), "mokou")
	}
}

// func TestShowUser(t * testing.T) {
// 	assert := assert.New(t)
// 	config.Init("../config/environments/", "test")
// 	database.Init()
// 	defer database.Close()
// 	e := server.Router()

// 	req := httptest.NewRequest(http.MethodGet, "/users", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues("1")

// 	if assert.NoError(controllers.IndexUser(c)) {
// 		assert.Equal(http.StatusOK, rec.Code)
// 	}
// }

// func TestUpdateUser(t *testing.T) {
// 	assert := assert.New(t)
// 	config.Init("../config/environments/", "test")
// 	database.Init()
// 	defer database.Close()
// 	e := server.Router()

// 	/// 実行しない！
// 	// トランザクションを使う
// 	// DB 構造
// 	//  89 | alice | alice1@gmail.com | $2a$12$VpQ3EXcu.8WUPPzkbDC7SOQbCPNFxhYmBzCd9awGV1qfW6ymKRVj2 | 2021-11-12 21:28:55.573465 | 2021-11-12 21:28:55.573465
//  	// 90 | alice | alice2@gmail.com | $2a$12$rnirxEnR4PzPjUgwcdFSxOEq/isB5x3Al4QCHNwiwY.6pbcjoDjB6 | 2021-11-12 21:28:59.564403 | 2021-11-12 21:28:59.564403
//   //  91 | alice | alice3@gmail.com | $2a$12$RwCDLxgEQkkX03MGUWbD8OMdkwZJvpJxJusTrWNMCS04m1S6mdKbG | 2021-11-12 21:29:03.370536 | 2021-11-12 21:29:03.370536
// 	req := httptest.NewRequest(http.MethodPatch, "/users", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues("90")

// 	if assert.NoError(controllers.UpdateUser(c)) {
// 		assert.Equal(http.StatusNoContent, rec.Code)
// 	}
// }
