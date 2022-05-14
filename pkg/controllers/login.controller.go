package controllers

import (
	"mini_project/pkg/helpers"
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CreateUserController(c echo.Context) error {
	NewUser := models.User{}
	c.Bind(&NewUser)
	err := repository.CreateNewUser(NewUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Membuat Data User",
	})

}
