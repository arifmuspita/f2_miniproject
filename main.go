package main

import (
	"f2_miniproject/config"
	"f2_miniproject/handler"
	"f2_miniproject/repository"
	"f2_miniproject/routes"
	"f2_miniproject/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	// migration.Migration(db)

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	e := echo.New()

	routes.Init(e, userHandler)

	e.Logger.Fatal(e.Start(":1234"))
}
