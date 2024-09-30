package main

import (
	"pastebin-clone/configs"
	"pastebin-clone/internal/db"
	"pastebin-clone/internal/http/handlers"
	"pastebin-clone/internal/repositories"
	"pastebin-clone/internal/services"

	"github.com/labstack/echo/v4"

	_ "pastebin-clone/internal/http/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	configs.LoadConfig()

	db.ConnectDB()
	db.MigrateDB()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "API is up and running!"})
	})

	authRepo := repositories.NewAuthRepository()
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)
	e.POST("/refresh-token", authHandler.RefreshToken)

	e.Start(":8080")
}
