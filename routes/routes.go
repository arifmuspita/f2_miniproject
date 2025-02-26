package routes

import (
	"f2_miniproject/handler"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo,
	userHandler handler.UserHandler,
) {

	v1 := e.Group("/v1")
	v1.POST("/register", userHandler.Register)
	v1.POST("/login", userHandler.Login)

}
