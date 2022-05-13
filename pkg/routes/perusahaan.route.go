package routes

import (
	"mini_project/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RoutePeusahaan() *echo.Echo {
	e := echo.New()
	e.GET("/perusahaan", controllers.GetAllPerusahaanController)
	e.GET("/perusahaan/:id", controllers.GetPerusahaanController)
	e.DELETE("/perusahaan/:id", controllers.DeletePerusahaanController)
	e.POST("/perusahaan", controllers.CreatePerusahaanController)
	e.PUT("/perusahaan/:id", controllers.UpdatePerusahaanController)
	return e
}
