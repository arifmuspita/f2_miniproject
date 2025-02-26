package main

import (
	"f2_miniproject/config"
	"f2_miniproject/migration"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	migration.Migration(db)

	e := echo.New()

	e.Logger.Fatal(e.Start(":1234"))
}
