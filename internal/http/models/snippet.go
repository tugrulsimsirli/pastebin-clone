package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateSnippetRequestModel struct {
	Title   string `json:"title" example:"Sample Snippet"`
	Content string `json:"content" example:"This is a sample snippet content"`
}

type UpdateSnippetRequestModel struct {
	Title   *string `json:"title,omitempty" example:"Updated Snippet"`
	Content *string `json:"content,omitempty" example:"Updated snippet content"`
}

type SnippetResponseModel struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	ViewCount    uint      `json:"view_count"`
	IsPublic     bool      `json:"is_public"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
	IsDeleted    bool      `json:"is_deleted"`
}
