package controllers

import (
	"mini_project/pkg/database"
	"mini_project/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllKedatanganKapalController(c echo.Context) error {
	listkapal, err := database.GetAllKedatanganKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, listkapal)
}

func GetKedatanganKapalController(c echo.Context) error {
	kapal, err := database.GetKedatanganKapalById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}

func DeleteKedatanganKapalController(c echo.Context) error {
	id := c.Param("id")
	err := database.DeleteKedatanganKapalById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete data kedatangan kapal dengan id '" + id + "'",
	})

}

func CreateKedatanganKapalController(c echo.Context) error {
	kapal := models.KedatanganKapal{}
	c.Bind(&kapal)
	err := database.CreateNewKedatanganKapal(kapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}

func UpdateKedatanganKapalController(c echo.Context) error {
	kapal := models.KedatanganKapal{}
	c.Bind(&kapal)
	err := database.UpdateKedatanganKapalById(c.Param("id"), kapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}
