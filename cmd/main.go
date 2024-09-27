package main

import (
	"pastebin-clone/configs"
	"pastebin-clone/internal/db"
	"pastebin-clone/internal/handlers"
	"pastebin-clone/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {
	// Config dosyasını yükle
	configs.LoadConfig()

	// Veritabanı bağlantısı ve migrasyon
	db.ConnectDB()
	db.MigrateDB()

	// Echo instance başlat
	e := echo.New()

	// Sağlık kontrolü endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "API is up and running!"})
	})

	// Kayıt ve giriş rotaları
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	// JWT ile korunan bir rota
	e.GET("/protected", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome to the protected route!"})
	}, middlewares.JWTMiddleware)

	// Sunucuyu başlat
	e.Start(":8080")
}
