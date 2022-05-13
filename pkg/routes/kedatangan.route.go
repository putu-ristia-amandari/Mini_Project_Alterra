package routes

import (
	"mini_project/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RouteKedatangan() *echo.Echo {
	e := echo.New()
	e.GET("/kedatangan", controllers.GetAllKedatanganKapalController)
	e.GET("/kedatangan/:id", controllers.GetKedatanganKapalController)
	e.DELETE("/kedatangan/:id", controllers.DeleteKedatanganKapalController)
	e.POST("/kedatangan", controllers.CreateKedatanganKapalController)
	e.PUT("/kedatangan/:id", controllers.UpdateKedatanganKapalController)
	return e
}
