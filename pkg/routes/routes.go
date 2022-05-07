package routes

import (
	"mini_project/pkg/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Home Page Informasi Kedatangan Kapal")
	})

	e.GET("/kedatangan", controllers.GetAllKedatanganKapalController)
	e.GET("/kedatangan/:id", controllers.GetKedatanganKapalController)
	e.DELETE("/kedatangan/:id", controllers.DeleteKedatanganKapalController)
	e.POST("/kedatangan", controllers.CreateKedatanganKapalController)
	e.PUT("/kedatangan/:id", controllers.UpdateKedatanganKapalController)

	return e
}
