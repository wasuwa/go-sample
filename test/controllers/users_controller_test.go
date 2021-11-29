package controllers_test

import (
	"fmt"
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
)

var (
	base = &models.Base{ID: 0}
	user = &models.User{
		Name:     "takada",
		Email:    "god@example.com",
		Password: "kenshi",
	}
	tests = []struct {
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

func init() {
	config.ResetPath()
}

func TestIndexUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	e := server.Router()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("ユーザーが見つからずエラーが返ること", func(t *testing.T) {
		assert.Error(controllers.IndexUser(c))
	})
	t.Run("ユーザーを取得できること", func(t *testing.T) {
		db.Create(user)
		assert.NoError(controllers.IndexUser(c))
		assert.Equal(http.StatusOK, rec.Code)
	})
}

func TestShowUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	e := server.Router()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	db.Create(user)

	t.Run("パラメーターが不正でエラーが返ること", func(t *testing.T) {
		c.SetParamValues("test")
		assert.Error(controllers.ShowUser(c))
	})
	t.Run("正しく通ること", func(t *testing.T) {
		id := strconv.Itoa(int(user.ID))
		c.SetParamValues(id)
		assert.NoError(controllers.ShowUser(c))
		assert.Equal(http.StatusOK, rec.Code)
	})
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	_, teardown := database.SetupTestDB()
	defer teardown()

	e := server.Router()
	for _, tc := range tests {
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

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	db.Create(user)
	id := strconv.Itoa(int(user.ID))
	e := server.Router()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPatch, "/users/:id", strings.NewReader(tc.input))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(id)
			services.Login(user, c)
			if err := controllers.UpdateUser(c); tc.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
				assert.Equal(http.StatusNoContent, rec.Code)
			}
		})
	}
}

func TestDestroyUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	e := server.Router()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("test")
	db.Create(user)
	assert.Error(controllers.DestroyUser(c))

	id := strconv.Itoa(int(user.ID))
	c.SetParamValues(id)
	assert.NoError(controllers.DestroyUser(c))
	assert.Equal(http.StatusNoContent, rec.Code)
}
