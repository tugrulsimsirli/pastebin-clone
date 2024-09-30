package repositories

import (
	"pastebin-clone/internal/db"
	data_models "pastebin-clone/internal/db/data-models"
	"pastebin-clone/internal/mapper"
	"pastebin-clone/internal/repositories/dto"
)

type UserRepositoryInterface interface {
	GetUserByUsername(username string) (*dto.UserDto, error)
}

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByUsername(username string) (*dto.UserDto, error) {
	var user data_models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	var response *dto.UserDto

	err := mapper.Map(user, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
