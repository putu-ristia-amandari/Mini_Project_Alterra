package routes

import (
	"mini_project/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RouteMuatan() *echo.Echo {
	e := echo.New()
	e.GET("/muatan", controllers.GetAllMuatanController)
	e.GET("/muatan/:id", controllers.GetMuatanController)
	e.DELETE("/muatan/:id", controllers.DeleteMuatanController)
	e.POST("/muatan", controllers.CreateMuatanController)
	e.PUT("/muatan/:id", controllers.UpdateMuatanController)
	return e
}
