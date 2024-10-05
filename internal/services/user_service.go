package services

import (
	"pastebin-clone/internal/http/models"
	"pastebin-clone/internal/mapper"
	"pastebin-clone/internal/repositories"

	"github.com/google/uuid"
)

type UserServiceInterface interface {
	GetUserDetail(userId uuid.UUID) (*models.UserDetailResponseModel, error)
}

type UserService struct {
	UserRepo repositories.UserRepositoryInterface
}

func NewUserService(userRepo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (u *UserService) GetUserDetail(userId uuid.UUID) (*models.UserDetailResponseModel, error) {
	user, err := u.UserRepo.GetUserDetail(userId)
	if err != nil {
		return nil, err
	}

	response := &models.UserDetailResponseModel{}
	err = mapper.Map(user, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
