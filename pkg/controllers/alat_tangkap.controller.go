package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllAlatTangkapController(c echo.Context) error {
	ListAlatTangkap, err := repository.GetAllAlatTangkap()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, ListAlatTangkap)
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
		"message": "succes delete data alat tangkap dengan id '" + id + "'",
	})

}

func CreateAlatTangkapController(c echo.Context) error {
	id := c.Param("id")
	NamaAlatTangkap := models.AlatTangkap{}
	c.Bind(&NamaAlatTangkap)
	err := repository.CreateNewAlatTangkap(NamaAlatTangkap)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes insert data alat tangkap dengan id '" + id + "'",
	})

}

func UpdateAlatTangkapController(c echo.Context) error {
	NamaAlatTangkap := models.AlatTangkap{}
	c.Bind(&NamaAlatTangkap)
	err := repository.UpdateAlatTangkapById(c.Param("id"), NamaAlatTangkap)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaAlatTangkap)

}
