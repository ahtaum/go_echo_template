package routes

import (
	User "go_temp/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", User.GetUsers)
	e.POST("/add", User.StoreUser)
	e.PUT("/update", User.UpdateUser)
	e.DELETE("/delete", User.DeleteUser)

	return e
}
