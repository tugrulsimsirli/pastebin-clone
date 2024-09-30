package data_models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string    `gorm:"type:varchar(20);unique"`
	Password     string    `gorm:"type:text"`
	CreatedDate  time.Time `gorm:"type:date;not null;default:now()"`
	ModifiedDate time.Time `gorm:"type:date;default:null"`
	IsDeleted    bool      `gorm:"type:boolean;default:false"`
}
