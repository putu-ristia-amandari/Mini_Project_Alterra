package controllers

import (
	"mini_project/pkg/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllKedatanganKapal(c echo.Context) error {
	result, err := models.GetAllKedatanganKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})

	}
	return c.JSON(http.StatusOK, result)
}
