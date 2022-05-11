package repository

// import (
// 	"fmt"
// 	"mini_project/pkg/database"
// 	"mini_project/pkg/helpers"
// 	"mini_project/pkg/models"
// )

// func LoginCheck(username, password string) (models.User, bool error) {
// 	var (
// 		User models.User
// 		pwd string
// 	)
// 	err := database.DB.Where(&User, "username = ?", username).First(&user).Error
// 	if err != nil {
// 		fmt.Println("Query Error")
// 	}
// 	return User, err

// 	match, err := helpers.CheckPasswordHash(password, pwd)
// 	if !match {
// 		fmt.Println("Hash and Password Doesn't Match")
// 		return false, err
// 	}

// 	return true, nil
// }

// func LoginUserValidation(username, role string) {
// 	validate := validator.New()

// 	user := User{
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
