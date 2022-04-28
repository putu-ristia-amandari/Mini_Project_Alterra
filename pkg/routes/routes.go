package routes

import (
	"github.com/MINI_PROJECT_ALTERRA/pkg/controllers"

	"github.com/labstack/echo"
)

func New() e *echo.Echo {
	e := echo.New

	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	return e
}