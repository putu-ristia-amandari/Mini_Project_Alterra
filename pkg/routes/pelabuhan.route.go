package routes

import (
	"mini_project/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RoutePelabuhan() *echo.Echo {
	e := echo.New()
	e.GET("/pelabuhan", controllers.GetAllPelabuhanController)
	e.GET("/pelabuhan/:id", controllers.GetPelabuhanController)
	e.DELETE("/pelabuhan/:id", controllers.DeletePelabuhanController)
	e.POST("/pelabuhan", controllers.CreatePelabuhanController)
	e.PUT("/pelabuhan/:id", controllers.UpdatePelabuhanController)
	return e
}
