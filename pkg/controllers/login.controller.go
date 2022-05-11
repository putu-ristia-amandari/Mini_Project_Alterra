package controllers

// import (
// 	"mini_project/pkg/helpers"
// 	"mini_project/pkg/models"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func LoginCheck(c echo.Context) error {
// 	username := c.FromValue("username")
// 	password := c.FromValue("password")

// 	res, err := models.User(username, password)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"messages": err.Error(),
// 		})
// 	}

// 	if !res {
// 		return echo.ErrUnauthorized
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.jwt.MapClaims {
// 			"id":         user.Id,
// 			"username":   user.Username,
// 			"role":       user.Role,
// 			"created_at": timeNow,
// 			"updated_at": timeNow,
// 			"expired_at": timeNow().Add(time.Hour *72).Unix(),

// 			t, err := token.SignedString([]byte("iuid"))
// 			if err != nil {
// 				return err
// 			}

// 			return c.JSON(http.StatusOK, echo.Map{
// 				"token": t,
// 			})
// 	}

// }

// func GenerateHashPassword(c echo.Context) error {
// 	password := c.Param("password")

// 	hash, _ := helpers.HashPassword(password)

// 	return c.JSON(http.StatusOK, hash)
// }
