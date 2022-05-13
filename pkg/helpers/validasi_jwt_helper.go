package helpers

// import (
// 	"errors"
// 	"fmt"
// 	"strings"

// 	"mini_project/pkg/models"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo"
// )

// var jwtKey = []byte("secret")

// func ParsingJWT(c echo.Context) (User models.User, err error) {
// 	token := c.Request().Header.Get("Authorization")
// 	arrToken := strings.Split(token, " ")
// 	if len(arrToken) < 2 {
// 		err = errors.New("Header Authorization Invalid Value")
// 		return User, err
// 	}
// 	tokenJwt, err := jwt.Parse(arrToken[1], func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected Signing Method: %v", token.Header["alg"])
// 		}
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		return User, err
// 	}

// 	if !tokenJwt.Valid {
// 		err = errors.New("Token is Invalid")
// 		return User, err
// 	}
// 	payloadJWT := tokenJwt.Claims.(jwt.MapClaims)
// 	floatId := payloadJWT["id"].(float64)
// 	User.Id = int(floatId)
// 	User.Role = payloadJWT["role"].(string)
// 	User.Username = payloadJWT["username"].(string)
// 	return User, err
// }
