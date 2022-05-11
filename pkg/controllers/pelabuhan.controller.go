package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllPelabuhanController(c echo.Context) error {
	listpelabuhan, err := repository.GetAllPelabuhan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, listpelabuhan)
}

func GetPelabuhanController(c echo.Context) error {
	pelabuhan, err := repository.GetPelabuhanById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, pelabuhan)

}

func DeletePelabuhanController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeletePelabuhanById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete data pelabuhan dengan id '" + id + "'",
	})

}

func CreatePelabuhanController(c echo.Context) error {
	id := c.Param("id")
	NamaPelabuhan := models.Pelabuhan{}
	c.Bind(&NamaPelabuhan)
	err := repository.CreateNewPelabuhan(NamaPelabuhan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes insert data pelabuhan dengan id  '" + id + "'",
	})

}

func UpdatePelabuhanController(c echo.Context) error {
	NamaPelabuhan := models.Pelabuhan{}
	c.Bind(&NamaPelabuhan)
	err := repository.UpdatePelabuhanById(c.Param("id"), NamaPelabuhan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaPelabuhan)

}
