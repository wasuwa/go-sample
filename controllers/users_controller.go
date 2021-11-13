package controllers

import (
	"net/http"
	"strconv"
	"twitter-app/models"

	"github.com/labstack/echo/v4"
)

func IndexUser(c echo.Context) error {
	var u models.User
	uu, err := u.All()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, uu)
}

func CreateUser(c echo.Context) error {
	u := new(models.User)
	r := new(models.ReceiveUser)
	err := c.Bind(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = c.Validate(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	r.BindUser(u)
	err = u.Create()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, u)
}

func ShowUser(c echo.Context) error {
	var u models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = u.Find(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func UpdateUser(c echo.Context) error {
	u := new(models.User)
	r := new(models.ReceiveUser)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = c.Bind(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = c.Validate(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	r.BindUser(u)
	err = u.Update(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}

func DestroyUser(c echo.Context) error {
	var u models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = u.Destroy(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}
