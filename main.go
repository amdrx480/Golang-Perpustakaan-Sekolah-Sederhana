package main

import (
	"perpustakaan/database"
	"perpustakaan/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load(".env")
	database.InitDatabase()

	database.Migrate()

	e := echo.New() // buat instance aplikasi echo baru

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
