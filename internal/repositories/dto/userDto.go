package dto

import "github.com/google/uuid"

type UserDto struct {
	ID       uuid.UUID `json:"id" example:"b8bba550-3b82-4fa8-9617-8d3c0ab69989"`
	Username string    `json:"username" example:"johndoe"`
	Password string    `json:"password" example:"password"`
}
