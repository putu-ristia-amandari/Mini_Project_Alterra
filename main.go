package main

import (
	"github.com/putu-ristia-amandari/Mini_Project_Alterra/pkg/routes"
)

func main() {
	db.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
