package controllers

import (
	"net/http"
	"strconv"

	User "go_temp/models"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	result, err := User.FetchAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreUser(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")
	about := c.FormValue("about")

	result, err := User.AddUser(name, password, about)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	password := c.FormValue("password")
	about := c.FormValue("about")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := User.UpdateUser(conv_id, name, password, about)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := User.DeleteUser(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
