package models

// import (
// 	"database/sql"
// 	"fmt"
// 	"mini_project/pkg/database"
// 	"mini_project/pkg/helpers"
// )

// type User struct {
// 	Id       int `json:"id"`
// 	Username int `json:"username"`
// }

// func LoginCheck(username, password string) (bool, error) {
// 	var (
// 		obj User
// 		pwd string
// 	)

// 	con := database.DBConnect()

// 	sqlStatement := "SELECT * FROM users WHERE username = ?"

// 	err := con.QueryRow(sqlStatement, username).Scan(
// 		&obj.Id, &obj.Username, &pwd,
// 	)

// 	if err == sql.ErrNoRows {
// 		fmt.Println("Username Not Found")
// 		return false, err
// 	}

// 	if err != nil {
// 		fmt.Println("Query Error")
// 		return false, err
// 	}

// 	match, err := helpers.CheckPasswordHash(password, pwd)
// 	if !match {
// 		fmt.Println("Hash and Password Doesn't Match")
// 		return false, err
// 	}

// 	return true, nil
// }
