package controllers

import (
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllTipeKapalController(c echo.Context) error {
	data, err := repository.GetAllTipeKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":           "Kamu Berhasil Melihat Semua Data Tipe Kapal Perikanan",
		"Daftar Tipe Kapal": data,
	})
}

func GetTipeKapalController(c echo.Context) error {
	TipeKapal, err := repository.GetTipeKapalById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, TipeKapal)

}
