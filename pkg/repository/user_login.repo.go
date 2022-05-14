package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/helpers"
	"mini_project/pkg/models"
)

func CreateNewUser(user models.User) error {
	err := database.DB.Save(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func LoginCheck(username, password string) (bool, error) {
	err := database.DB.First(&models.User{}, "username = ?", username).Debug().Error
	// if err ==  {
	// 	fmt.Println("Username Not Found")
	// 	return false, err
	// }
	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	var pwd string
	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password Doesn't Match")
		return false, err
	}

	return true, nil
}

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
