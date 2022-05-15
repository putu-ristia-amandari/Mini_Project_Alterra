package middleware

import (
	"fmt"
	"mini_project/pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

var SecretKey = []byte("asdf.123")

func CreateJwtToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":         user.Id,
		"username":   user.Username,
		"expired_at": time.Now().Add(time.Hour * 1).Unix(),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString(SecretKey)
	if err != nil {
		fmt.Println(err)
	}
	return token, nil

}

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: SecretKey,
})
