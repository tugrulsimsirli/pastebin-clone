package handlers

import (
	"net/http"
	"pastebin-clone/configs"
	"pastebin-clone/internal/db"
	"pastebin-clone/internal/http/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary      User registration
// @Description  Registers a new user and returns success message
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body models.RegisterRequestModel true "User registration data"
// @Success      200  {object} models.RegisterResponseModel
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /register [post]
func Register(c echo.Context) error {
	var newUser models.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	newUser.ID = uuid.New()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashedPassword)

	if err := db.DB.Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error saving user"})
	}

	return c.JSON(http.StatusOK, models.RegisterResponseModel{ID: newUser.ID})
}

// Login godoc
// @Summary      User login
// @Description  Logs in a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body models.User true "User credentials"
// @Success      200  {object} models.TokenResponse
// @Failure      400  {object} models.ErrorResponse
// @Failure      401  {object} models.ErrorResponse
// @Router       /login [post]
func Login(c echo.Context) error {
	var input models.User
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	var storedUser models.User

	if err := db.DB.Where("username = ?", input.Username).First(&storedUser).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid credentials"})
	}

	claims := &jwt.RegisteredClaims{
		Subject:   storedUser.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configs.AppConfig.JWTSecretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Could not generate token"})
	}

	return c.JSON(http.StatusOK, models.TokenResponse{Token: tokenString})
}

// RefreshToken godoc
// @Summary      Refresh JWT token
// @Description  Refreshes an access token using a refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body models.RefreshTokenRequest true "Refresh token"
// @Success      200  {object} models.AccessTokenResponse
// @Failure      400  {object} models.ErrorResponse
// @Failure      401  {object} models.ErrorResponse
// @Router       /refresh-token [post]
func RefreshToken(c echo.Context) error {
	var body models.RefreshTokenRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	token, err := jwt.Parse(body.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid refresh token"})
		}
		return []byte(configs.AppConfig.JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid or expired refresh token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid token claims"})
	}

	userID := claims["sub"].(string)

	accessTokenClaims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(configs.AppConfig.JWTSecretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Could not generate access token"})
	}

	return c.JSON(http.StatusOK, models.AccessTokenResponse{AccessToken: accessTokenString})
}
