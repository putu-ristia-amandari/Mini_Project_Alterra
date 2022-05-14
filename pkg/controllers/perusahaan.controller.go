package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllPerusahaanController(c echo.Context) error {
	data, err := repository.GetAllPerusahaanKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":           "Kamu Berhasil Melihat Semua Data Perusahaan Kapal Perikanan",
		"Daftar Perusahaan": data,
	})
}

func GetPerusahaanController(c echo.Context) error {
	kapal, err := repository.GetPerusahaanById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}

func DeletePerusahaanController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeletePerusahaanById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete data perusahaan kapal dengan id '" + id + "'",
	})

}

func CreatePerusahaanController(c echo.Context) error {
	id := c.Param("id")
	NamaPerusahaan := models.PerusahaanKapal{}
	c.Bind(&NamaPerusahaan)
	err := repository.CreateNewPerusahaan(NamaPerusahaan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes insert data perusahaan kapal dengan id  '" + id + "'",
	})

}

func UpdatePerusahaanController(c echo.Context) error {
	NamaPerusahaan := models.PerusahaanKapal{}
	c.Bind(&NamaPerusahaan)
	err := repository.UpdatePerusahaanById(c.Param("id"), NamaPerusahaan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaPerusahaan)

}
