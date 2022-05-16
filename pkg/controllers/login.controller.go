package controllers

import (
	"mini_project/pkg/helpers"
	"mini_project/pkg/middleware"
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	NewUser := models.User{}
	c.Bind(&NewUser)

	hash, _ := helpers.HashPassword(NewUser.Password)

	NewUser.Password = hash

	err := repository.CreateNewUser(NewUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Membuat Data User",
	})

}

func LoginController(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := repository.LoginCheck(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if !res {
		return echo.ErrUnauthorized
	}

	user := models.User{}
	token, err := middleware.CreateJwtToken(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"User Token": token,
	})
}
