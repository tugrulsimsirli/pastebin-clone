package main

import (
	"pastebin-clone/configs"
	"pastebin-clone/internal/bootstrap"
	"pastebin-clone/internal/db"

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

	// Dependency injection
	bootstrap.RegisterHandlers(e)

	e.Start(":8080")
}
