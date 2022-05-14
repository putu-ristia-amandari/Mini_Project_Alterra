package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=[${time_rfc3339_nano}], method=[${method}], host=[${host}], status=[${status}], uri=[${uri}]\n",
	}))
}
