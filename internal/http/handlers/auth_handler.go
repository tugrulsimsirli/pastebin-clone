package handlers

import (
	"net/http"
	"pastebin-clone/internal/http/models"
	"pastebin-clone/internal/services"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService services.AuthServiceInterface // Dependency Injection ile AuthService alınıyor
}

func NewAuthHandler(authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

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
func (h *AuthHandler) Register(c echo.Context) error {
	var req models.RegisterRequestModel
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	userID, err := h.AuthService.RegisterUser(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error saving user"})
	}

	return c.JSON(http.StatusOK, models.RegisterResponseModel{ID: userID})
}

// Login godoc
// @Summary      User login
// @Description  Logs in a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body models.LoginRequestModel true "User credentials"
// @Success      200  {object} models.LoginResponseModel
// @Failure      400  {object} models.ErrorResponse
// @Failure      401  {object} models.ErrorResponse
// @Router       /login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var req models.LoginRequestModel
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	accessToken, refreshToken, expireDate, err := h.AuthService.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.LoginResponseModel{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireDate:   expireDate,
	})
}

// RefreshToken godoc
// @Summary      Refresh JWT token
// @Description  Refreshes an access token using a refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body models.RefreshTokenRequestModel true "Refresh token"
// @Success      200  {object} models.RefreshTokenResponseModel
// @Failure      400  {object} models.ErrorResponse
// @Failure      401  {object} models.ErrorResponse
// @Router       /refresh-token [post]
func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var req models.RefreshTokenRequestModel
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
	}

	accessToken, expireDate, err := h.AuthService.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.RefreshTokenResponseModel{
		AccessToken: accessToken,
		ExpireDate:  expireDate,
	})
}
