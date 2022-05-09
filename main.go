package main

import (
	"mini_project/pkg/database"
	"mini_project/pkg/routes"
)

func init() {
	database.InitDBConnect()
	database.InitialMigration()
}

func main() {

	e := routes.Route()

	e.Logger.Fatal(e.Start(":1234"))
}
