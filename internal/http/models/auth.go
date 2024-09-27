package models

import "github.com/google/uuid"

type RegisterRequestModel struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"password"`
}

type RegisterResponseModel struct {
	ID uuid.UUID `json:"id" example:"b8bba550-3b82-4fa8-9617-8d3c0ab69989"`
}

type LoginRequestModel struct {
	Username string `json:"username" example:"userName"`
	Password string `json:"password" example:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" example:"some-refresh-token"`
}

// SuccessResponse represents a successful response with a message
type SuccessResponse struct {
	Message string `json:"message" example:"User registered successfully"`
}

// TokenResponse represents the response with a JWT token
type TokenResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}

// AccessTokenResponse represents the response with an access token
type AccessTokenResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}

// ErrorResponse represents a generic error response
type ErrorResponse struct {
	Message string `json:"message" example:"Invalid credentials"`
}
