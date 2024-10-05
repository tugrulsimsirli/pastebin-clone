package repositories

import (
	"pastebin-clone/internal/db"
	data_models "pastebin-clone/internal/db/data-models"

	"github.com/google/uuid"
)

type AuthRepositoryInterface interface {
	CreateUser(email string, username string, password string) (uuid.UUID, error)
}

type AuthRepository struct{}

func NewAuthRepository() AuthRepositoryInterface {
	return &AuthRepository{}
}

func (r *AuthRepository) CreateUser(email string, username string, hashedPassword string) (uuid.UUID, error) {
	newUser := data_models.User{
		ID:       uuid.New(),
		Email:    email,
		Username: username,
		Password: hashedPassword,
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		return uuid.Nil, err
	}

	return newUser.ID, nil
}
