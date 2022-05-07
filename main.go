package main

import (
	"mini_project/pkg/config"
	"mini_project/pkg/routes"
)

func init() {
	config.InitDBConnect()
	config.InitialMigration()
}

func main() {

	e := routes.Route()

	e.Logger.Fatal(e.Start(":1234"))
}
