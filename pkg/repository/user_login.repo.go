package repository

import (
	"database/sql"
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/helpers"
	"mini_project/pkg/models"
)

func LoginCheck(username, password string) (bool, error) {
	// user := models.User{}
	// err := database.DB.First("username = ? AND password = ?", user.Username, user.Password).First(&user).Debug().Error
	// err := database.DB.First(&models.User{}, "username = ?", username).Debug().Error
	// if err != nil {
	// 	fmt.Println(err)
	// 	return false, err
	// }
	var pwd string
	user := models.User{}
	err := database.DB.First(&models.User{}, "username = ?", username).Scan(
		&user).Debug().Error

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password Doesn't Match")
		return false, err
	}

	return true, nil
}

func CreateNewUser(user models.User) error {
	err := database.DB.Save(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// func LoginCheck(username, password string) (bool, error) {
// 	err := database.DB.First(&models.User{}, "username = ?", username).Debug().Error
// 	if err ==  {
// 		fmt.Println("Username Not Found")
// 		return false, err
// 	}
// 	if err != nil {
// 		fmt.Println("Query Error")
// 		return false, err
// 	}

// 	var pwd string
// 	match, err := helpers.CheckPasswordHash(password, pwd)
// 	if !match {
// 		fmt.Println("Hash and Password Doesn't Match")
// 		return false, err
// 	}

// 	return true, nil
// }

// func LoginUserValidation(username, role string) {
// 	validate := validator.New()

// 	user := &models.User{
// 		Id: id,
// 		Username: username,
// 		Role: role,
// 		Created_At: created_at,
// 		Updated_At: updated_at,
// 	}

// 	err := validate.Struct(user)
// 	if err != nil {
// 		return (err)
// 	}
// }
