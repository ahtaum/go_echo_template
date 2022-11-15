package routes

import (
	User "go_temp/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	g := e.Group("/apiV1")
	g.GET("/", User.GetUsers)
	g.POST("/add", User.StoreUser)
	g.PUT("/update", User.UpdateUser)
	g.DELETE("/delete", User.DeleteUser)

	return e
}
