package routes

import (
	"mini_project/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RouteAlat() *echo.Echo {
	e := echo.New()
	e.GET("/alat", controllers.GetAllAlatTangkapController)
	e.GET("/alat/:id", controllers.GetAlatTangkapController)
	e.DELETE("/alat/:id", controllers.DeleteAlatTangkapController)
	e.POST("/alat", controllers.CreateAlatTangkapController)
	e.PUT("/alat/:id", controllers.UpdateAlatTangkapController)
	return e
}
