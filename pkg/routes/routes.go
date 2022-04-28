package routes

import (
	"github.com/Mini_Project_Alterra/pkg/controllers"

	"github.com/labstack/echo"
)

func New() e *echo.Echo {
	e := echo.New

	e.GET("/kapals", controller.GetKapalsController)
	e.GET("/kapals/:id", controller.GetKapalController)
	e.POST("/kapals", controller.CreateKapalController)
	e.DELETE("/kapals/:id", controller.DeleteKapalController)
	e.PUT("/kapals/:id", controller.UpdateKapalController)

	return e
}