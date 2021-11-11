package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"twitter-app/controllers"
	"twitter-app/database"
	"twitter-app/factories"
	"twitter-app/server"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIndexUser(t *testing.T) {
	factories.Seed()
	defer database.Close()
	e := server.Router()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.IndexUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// mock DB を用意する必要がある
		println(rec.Body.String())
	}

	factories.Rollback()
}
