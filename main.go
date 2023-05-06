package main

import (
	"perpustakaan/database"
	"perpustakaan/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase()

	database.Migrate()

	e := echo.New() // buat instance aplikasi echo baru

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
