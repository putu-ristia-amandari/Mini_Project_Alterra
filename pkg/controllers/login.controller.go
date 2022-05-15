package controllers

import (
	"mini_project/pkg/helpers"
	"mini_project/pkg/middleware"
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

	// return c.String(http.StatusOK, "Berhasil Login")

	user := models.User{}
	token, err := middleware.CreateJwtToken(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"User Token": token,
	})
}

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

// func UserLogin(c echo.Context) error {
// 	user := models.User{}
// 	c.Bind(&user)

// 	err := database.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Debug().Error
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "Tidak Berhasil Login",
// 			"error":   err.Error(),
// 		})
// 	}
// 	token, err := middleware.CreateJwtToken(user.Username, user.Password)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message":    "berhasil login",
// 		"user token": token,
// 	})
// }
