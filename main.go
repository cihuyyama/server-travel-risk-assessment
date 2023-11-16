package main

import (
	"travel-risk-assessment/database"
	"travel-risk-assessment/initializers"
	"travel-risk-assessment/router"
)

func main() {
	initializers.LoadEnvVariables()

	database.Initialize()

	database.Connect()

	database.Migrate()

	r := router.SetupRouter()

	r.Run()
}
