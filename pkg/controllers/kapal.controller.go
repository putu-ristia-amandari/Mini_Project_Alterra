package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllKapalController(c echo.Context) error {
	data, err := repository.GetAllKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "Kamu Berhasil Melihat Semua Data Kapal Perikanan",
		"Daftar Kapal Perikanan": data,
	})
}

func GetKapalController(c echo.Context) error {
	kapal, err := repository.GetKapalById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}

func DeleteKapalController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteKapalById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete data kapal dengan id '" + id + "'",
	})

}

func CreateKapalController(c echo.Context) error {
	id := c.Param("id")
	NamaKapal := models.Kapal{}
	c.Bind(&NamaKapal)
	err := repository.CreateNewKapal(NamaKapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes insert data kapal dengan id '" + id + "'",
	})

}

func UpdateKapalController(c echo.Context) error {
	NamaKapal := models.Kapal{}
	c.Bind(&NamaKapal)
	err := repository.UpdateKapalById(c.Param("id"), NamaKapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaKapal)

}
