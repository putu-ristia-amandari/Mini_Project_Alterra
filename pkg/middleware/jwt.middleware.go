package middleware

import (
	"fmt"
	"mini_project/pkg/constans"
	"mini_project/pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

func CreateJwtToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"id":         user.Id,
		"username":   user.Username,
		"expired_at": time.Now().Add(time.Hour * 1).Unix(),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(constans.SECRET_KEY))
	if err != nil {
		fmt.Println(err)
	}
	return token, nil

}

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: ([]byte(constans.SECRET_KEY)),
})
