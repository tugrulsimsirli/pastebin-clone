package main

import (
	"pastebin-clone/configs"
	"pastebin-clone/internal/db"
	"pastebin-clone/internal/http/handlers"
	"pastebin-clone/internal/http/middlewares"

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

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)
	e.POST("/refresh-token", handlers.RefreshToken)

	e.POST("/snippet", handlers.CreateSnippet)

	e.GET("/protected", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome to the protected route!"})
	}, middlewares.JWTMiddleware)

	e.Start(":8080")
}
