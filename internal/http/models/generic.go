package models

import "github.com/google/uuid"

type IdResponseModel struct {
	ID uuid.UUID `json:"id" example:"b8bba550-3b82-4fa8-9617-8d3c0ab69989"`
}

type BooleanRequestModel struct {
	Bool bool `json:"bool" example:"true"`
}
