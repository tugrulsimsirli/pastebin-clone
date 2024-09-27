package handlers

import (
	"net/http"
	"pastebin-clone/internal/db"
	"pastebin-clone/internal/http/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateSnippet(c echo.Context) error {
	var newSnippet models.Snippet
	if err := c.Bind(&newSnippet); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	newSnippet.UserID = userID

	if err := db.DB.Create(&newSnippet).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error saving snippet"})
	}

	return c.JSON(http.StatusOK, newSnippet)
}
