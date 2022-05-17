package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllAlatTangkapController(c echo.Context) error {
	data, err := repository.GetAllAlatTangkap()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":             "Kamu Berhasil Melihat Semua Data Alat Tangkap",
		"Daftar Alat Tangkap": data,
	})
}

func GetAlatTangkapController(c echo.Context) error {
	AlatTangkap, err := repository.GetAlatTangkapById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, AlatTangkap)

}

func DeleteAlatTangkapController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteAlatTangkapById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menghapus Data Pelabuhan dengan id '" + id + "'",
	})

}

func CreateAlatTangkapController(c echo.Context) error {
	NamaAlatTangkap := models.AlatTangkap{}
	c.Bind(&NamaAlatTangkap)
	err := repository.CreateNewAlatTangkap(NamaAlatTangkap)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Alat Tangkap",
	})
}

func UpdateAlatTangkapController(c echo.Context) error {
	NamaAlatTangkap1 := models.AlatTangkap{}
	c.Bind(&NamaAlatTangkap1)
	err := repository.UpdateAlatTangkapById(c.Param("id"), NamaAlatTangkap1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Memperbaharui Data",
	})

}
