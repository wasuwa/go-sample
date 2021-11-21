package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
	// "strconv"
	"strings"

	"testing"
	"twitter-app/config"
	"twitter-app/controllers"
	"twitter-app/database"
	"twitter-app/models"
	"twitter-app/server"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	t time.Time
	base = &models.Base {
		ID: 0,
		CreatedAt: t,
		UpdatedAt: t,
	}
	user = &models.User{
		Name:     "takada",
		Email:    "god@example.com",
		Password: "kenshi",
	}
	testcases = []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			"正しく通ること",
			`{"name":"mokou","email":"mokou@example.com","password":"orange"}`,
			false,
		},
		{
			"パスワードが6文字以上の制限でエラーが返ること",
			`{"name":"mokou","email":"mokou@example.com","password":"apple"}`,
			true,
		},
		{
			"名前は15文字以下の制限でエラーが返ること",
			fmt.Sprintf(`{"name":"%s","email":"mokou@example.com","password":"orange"}`, strings.Repeat("mokou", 4)),
			true,
		},
		{
			"メールアドレスは256文字以下の制限でエラーが返ること",
			fmt.Sprintf(`{"name":"mokou","email":"%s","password":"orange"}`, strings.Repeat("mokou", 49)+"@example.com"),
			true,
		},
		{
			"メールアドレスのフォーマットでエラーが返ること",
			`{"name":"mokou","email":"examplecom","password":"orange"}`,
			true,
		},
		{
			"名前は必須の制限でエラーが返ること",
			`{"email":"mokou@example.com","password":"orange"}`,
			true,
		},
		{
			"メールアドレスは必須の制限でエラーが返ること",
			`{"name":"mokou","password":"orange"}`,
			true,
		},
		{
			"パスワードは必須の制限でエラーが返ること",
			`{"name":"mokou","email":"mokou@example.com"}`,
			true,
		},
	}
)

func TestIndexUser(t *testing.T) {
	assert := assert.New(t)
	config.ResetPath()
	db, teardown := database.SetupTestDB()
	defer teardown()

	e := server.Router()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.Error(controllers.IndexUser(c))
	db.Create(user)
	assert.NoError(controllers.IndexUser(c))
	assert.Equal(http.StatusOK, rec.Code)
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	_, teardown := database.SetupTestDB()
	defer teardown()

	e := server.Router()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tc.input))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if err := controllers.CreateUser(c); tc.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
				assert.Equal(http.StatusCreated, rec.Code)
			}
		})
	}
}

// func TestShowUser(t *testing.T) {
// 	assert := assert.New(t)
// 	db, teardown := database.SetupTestDB()
// 	defer teardown()

// 	db.Create(user)
// 	id := strconv.Itoa(int(user.ID))
// 	e := server.Router()
// 	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues(id)

// 	assert.NoError(controllers.ShowUser(c))
// 	assert.Equal(http.StatusOK, rec.Code)
// }

// func TestUpdateUser(t *testing.T) {
// 	assert := assert.New(t)
// 	db, teardown := database.SetupTestDB()
// 	defer teardown()

// 	db.Create(user)
// 	id := strconv.Itoa(int(user.ID))
// 	e := server.Router()
// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			req := httptest.NewRequest(http.MethodPatch, "/users/:id", strings.NewReader(tc.input))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetParamNames("id")
// 			c.SetParamValues(id)
// 			if err := controllers.UpdateUser(c); tc.wantErr {
// 				assert.Error(err)
// 			} else {
// 				assert.NoError(err)
// 				assert.Equal(http.StatusNoContent, rec.Code)
// 			}
// 		})
// 	}
// }

// func TestDestroyUser(t *testing.T) {
// 	assert := assert.New(t)
// 	db, teardown := database.SetupTestDB()
// 	defer teardown()

// 	db.Create(user)
// 	id := strconv.Itoa(int(user.ID))
// 	e := server.Router()
// 	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues(id)

// 	assert.NoError(controllers.DestroyUser(c))
// 	assert.Equal(http.StatusNoContent, rec.Code)
// }
