package main

import (
	"pastebin-clone/configs"
	"pastebin-clone/internal/db"
	"pastebin-clone/internal/handlers"
	"pastebin-clone/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadConfig()

	db.ConnectDB()
	db.MigrateDB()

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "API is up and running!"})
	})

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)
	e.POST("/refresh-token", handlers.RefreshToken)

	e.GET("/protected", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome to the protected route!"})
	}, middlewares.JWTMiddleware)

	e.Start(":8080")
}
