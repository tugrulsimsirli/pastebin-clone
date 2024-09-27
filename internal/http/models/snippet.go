package models

import (
	"github.com/google/uuid"
)

type Snippet struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	ViewCount uint      `json:"view_count"`
}
