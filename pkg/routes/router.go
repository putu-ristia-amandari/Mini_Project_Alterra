package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Home Page Informasi Kedatangan Kapal")
	})

	// Routes Login
	// e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	// e.POST("/login", controllers.LoginCheck)

	//Routes Validation
	// e.GET("/validation", controllers.LoginUserValidation)

	return e
}
