package data_models

import (
	"time"

	"github.com/google/uuid"
)

type Snippet struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid"`
	Title        string    `gorm:"type:varchar(255)" json:"title"`
	Content      string    `gorm:"type:varchar;not null"`
	ViewCount    uint      `gorm:"type:integer;default:0"`
	IsPublic     bool      `gorm:"type:boolean;default:false"`
	CreatedDate  time.Time `gorm:"type:date;not null;default:now()"`
	ModifiedDate time.Time `gorm:"type:date;default:null"`
	IsDeleted    bool      `gorm:"type:boolean;default:false"`
}
