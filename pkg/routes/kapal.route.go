package routes

import (
	"mini_project/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RouteKapal() *echo.Echo {
	e := echo.New()
	// e.GET("/kapal", controllers.GetAllKapalController, middleware.IsAuthenticated)
	e.GET("/kapal/:id", controllers.GetKapalController)
	e.DELETE("/kapal/:id", controllers.DeleteKapalController)
	e.POST("/kapal", controllers.CreateKapalController)
	e.PUT("/kapal/:id", controllers.UpdateKapalController)
	return e
}
