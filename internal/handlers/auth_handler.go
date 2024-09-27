package handlers

import (
	"net/http"
	"pastebin-clone/configs"
	"pastebin-clone/internal/db"
	"pastebin-clone/internal/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error saving user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}

func Login(c echo.Context) error {
	var input models.User
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	var user models.User

	if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	userID := strconv.FormatUint(uint64(user.ID), 10)

	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configs.AppConfig.JWTSecretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
