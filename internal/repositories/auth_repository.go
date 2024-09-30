package repositories

import (
	"pastebin-clone/internal/db"
	data_models "pastebin-clone/internal/db/data-models"

	"github.com/google/uuid"
)

type AuthRepositoryInterface interface {
	CreateUser(username string, password string) (uuid.UUID, error)
}

type AuthRepository struct{}

func NewAuthRepository() AuthRepositoryInterface {
	return &AuthRepository{}
}

func (r *AuthRepository) CreateUser(username string, hashedPassword string) (uuid.UUID, error) {
	newUser := data_models.User{
		ID:       uuid.New(),
		Username: username,
		Password: hashedPassword,
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		return uuid.Nil, err
	}

	return newUser.ID, nil
}
