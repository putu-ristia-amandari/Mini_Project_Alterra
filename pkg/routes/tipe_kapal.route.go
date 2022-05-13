package routes

import (
	"mini_project/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RouteTipeKapal() *echo.Echo {
	e := echo.New()
	e.GET("/tipe", controllers.GetAllTipeKapalController)
	e.GET("/tipe/:id", controllers.GetTipeKapalController)
	e.DELETE("/tipe/:id", controllers.DeleteTipeKapalController)
	e.POST("/tipe", controllers.CreateTipeKapalController)
	e.PUT("/tipe/:id", controllers.UpdateTipeKapalController)
	return e
}
