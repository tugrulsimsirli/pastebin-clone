package dto

import (
	"time"

	"github.com/google/uuid"
)

type SnippetDto struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	ViewCount    uint      `json:"view_count"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
	IsDeleted    bool      `json:"is_deleted"`
}
