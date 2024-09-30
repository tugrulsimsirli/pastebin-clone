package repositories

import (
	"pastebin-clone/internal/db"
	"pastebin-clone/internal/repositories/dto"

	"github.com/google/uuid"
)

type AuthRepositoryInterface interface {
	CreateUser(username string, password string) (uuid.UUID, error)
	GetUserByUsername(username string) (*dto.User, error)
}

type AuthRepository struct{}

func NewAuthRepository() AuthRepositoryInterface {
	return &AuthRepository{}
}

func (r *AuthRepository) CreateUser(username string, hashedPassword string) (uuid.UUID, error) {
	newUser := dto.User{
		ID:       uuid.New(),
		Username: username,
		Password: hashedPassword,
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		return uuid.Nil, err
	}

	return newUser.ID, nil
}

func (r *AuthRepository) GetUserByUsername(username string) (*dto.User, error) {
	var user dto.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
