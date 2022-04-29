package main

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/routes"
	// "github.com/putu-ristia-amandari/Mini_Project_Alterra/pkg/database"
	// "github.com/putu-ristia-amandari/Mini_Project_Alterra/pkg/routes"
)

func main() {
	db, err := database.DBConnect()

	if err != nil {
		fmt.Println("Failed to connect database", err.Error())
		return
	}
	defer db.Close()
	fmt.Println("Succes to connect database")

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
