package routes

import (
	// "mini_project/pkg/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Home Page Informasi Kedatangan Kapal")
	})

	// e.GET("/kedatangan", controllers.GetAllKedatanganKapal)
	// e.GET("/kedatangan/:id", controllers.GetKedatanganKapal)
	// e.POST("/kedatangan", controllers.CreateKedatanganKapal)
	// e.DELETE("/kedatangan/:id", controllers.DeleteKedatanganKapal)
	// e.PUT("/kedatangan/:id", controllers.UpdateKedatanganKapal)

	return e
}
