package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllKedatanganKapalController(c echo.Context) error {
	data, err := repository.GetAllKedatanganKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                           "Kamu Berhasil Melihat Semua Data Kedatangan Kapal",
		"Daftar Kedatangan Kapal Perikanan": data,
	})
}

func GetKedatanganKapalController(c echo.Context) error {
	kapal, err := repository.GetKedatanganKapalById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}

func DeleteKedatanganKapalController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteKedatanganKapalById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete data kedatangan kapal dengan id '" + id + "'",
	})

}

func CreateKedatanganKapalController(c echo.Context) error {
	id := c.Param("id")
	kapal := models.KedatanganKapal{}
	c.Bind(&kapal)
	err := repository.CreateNewKedatanganKapal(kapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes insert data kedatangan kapal dengan id '" + id + "'",
	})

}

func UpdateKedatanganKapalController(c echo.Context) error {
	kapal := models.KedatanganKapal{}
	c.Bind(&kapal)
	err := repository.UpdateKedatanganKapalById(c.Param("id"), kapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}
