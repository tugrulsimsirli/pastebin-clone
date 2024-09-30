package data_models

import (
	"github.com/google/uuid"
)

type Snippet struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	Content   string
	ViewCount uint
}
