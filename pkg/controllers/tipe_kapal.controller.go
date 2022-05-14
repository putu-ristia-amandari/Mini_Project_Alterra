package controllers

import (
	"mini_project/pkg/models"
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

func DeleteTipeKapalController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteTipeKapalById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete data tipe kapal dengan id '" + id + "'",
	})

}

func CreateTipeKapalController(c echo.Context) error {
	id := c.Param("id")
	TipeKapal := models.TipeKapal{}
	c.Bind(&TipeKapal)
	err := repository.CreateNewTipeKapal(TipeKapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes insert data tipe kapal dengan id '" + id + "'",
	})

}

func UpdateTipeKapalController(c echo.Context) error {
	TipeKapal := models.TipeKapal{}
	c.Bind(&TipeKapal)
	err := repository.UpdateTipeKapalById(c.Param("id"), TipeKapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, TipeKapal)

}
