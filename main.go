package main

import (
	"mini_project/pkg/database"
	"mini_project/pkg/routes"
)

func main() {
	DB, err := database.DBConnect()

	if err != nil {
		panic(err)
	}

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
