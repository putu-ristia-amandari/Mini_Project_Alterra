package routes

import (
	"net/http"
	

	"github.com/labstack/echo/v4"
)

func New() e *echo.Echo {
	e := echo.New

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello!")
	})

	// e.GET("/kedatangans", controller.GetkedatangansController)
	// e.GET("/kedatangans/:id", controller.GetkedatanganController)
	// e.POST("/kedatangans", controller.CreatekedatanganController)
	// e.DELETE("/kedatangans/:id", controller.DeletekedatanganlController)
	// e.PUT("/kedatangans/:id", controller.UpdatekedatanganController)

	return e
}