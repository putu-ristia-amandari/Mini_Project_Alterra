package controllers

// import (
// 	"net/http"

// 	validator "github.com/go-playground/validator/v10"
// 	"github.com/labstack/echo/v4"
// )

// type LoginUser struct {
// 	Username string `validate:"required"`
// 	Role     string `validate:"role=admin"`
// }

// func LoginUserValidationController(c echo.Context) error {
// 	v := validator.New()

// 	role := "user"

// 	err := v.Var(role, "required, role")
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"message": "for admin only"
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{
// 		"message": "success"
// 	})

// }
