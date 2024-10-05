package models

import "github.com/google/uuid"

type RegisterRequestModel struct {
	Email    string `json:"email" example:"johndoe@johndoe.com"`
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"password"`
}

type RegisterResponseModel struct {
	ID uuid.UUID `json:"id" example:"b8bba550-3b82-4fa8-9617-8d3c0ab69989"`
}

type LoginRequestModel struct {
	Email    string `json:"email" example:"johndoe@johndoe.com"`
	Password string `json:"password" example:"password"`
}

type LoginResponseModel struct {
	UserID       uuid.UUID `json:"user_id" example:"b8bba550-3b82-4fa8-9617-8d3c0ab69989"`
	AccessToken  string    `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ..."`
	RefreshToken string    `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ExpireDate   string    `json:"expire_date" example:"1970-01-01 00:00:00"`
}

type RefreshTokenRequestModel struct {
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ..."`
}

type RefreshTokenResponseModel struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ..."`
	RefreshToken string `json:"refresh_token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ..."` // Opsiyonel
	ExpireDate   string `json:"expire_date" example:"1970-01-01 00:00:00"`
}

// ErrorResponse represents a generic error response
type ErrorResponse struct {
	Message string `json:"message" example:"Invalid credentials"`
}
