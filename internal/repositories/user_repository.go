package repositories

import (
	"pastebin-clone/internal/db"
	data_models "pastebin-clone/internal/db/data-models"
	"pastebin-clone/internal/mapper"
	"pastebin-clone/internal/repositories/dto"

	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	GetUserByEmail(email string) (*dto.UserDto, error)
	GetUserDetail(userID uuid.UUID) (*dto.UserDetailDto, error)
}

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByEmail(email string) (*dto.UserDto, error) {
	var user data_models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	response := &dto.UserDto{}

	err := mapper.Map(user, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *UserRepository) GetUserDetail(userID uuid.UUID) (*dto.UserDetailDto, error) {
	var user data_models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	response := &dto.UserDetailDto{}

	err := mapper.Map(user, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
