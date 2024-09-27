package main

import (
	"net/http"
	"pastebin-clone/internal/db"

	"github.com/labstack/echo/v4"
)

func main() {
	// Veritabanına bağlan
	db.ConnectDB()
	db.MigrateDB()

	// Echo'yu başlat
	e := echo.New()

	// Sağlık kontrolü endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "API is up and running!",
		})
	})

	// Sunucuyu başlat
	e.Start(":8080")
}
