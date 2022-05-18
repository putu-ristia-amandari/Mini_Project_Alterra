package repository

import (
	"database/sql"
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/helpers"
	"mini_project/pkg/models"
)

func LoginCheck(username, password string) (bool, error) {
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

	match, err := helpers.CheckPasswordHash(password, user.Password)
	if !match {
		fmt.Println("Hash and Password Doesn't Match")
		return false, err
	}

	return true, nil
}

func CreateNewUser(user models.User) error {
	err := database.DB.Create(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
