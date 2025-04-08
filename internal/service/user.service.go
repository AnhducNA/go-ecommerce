package service

import (
	"github.com/AnhducNA/go-ecommerce/internal/repo"
	"github.com/AnhducNA/go-ecommerce/pkg/response"
)

// INTERFACE VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserServie(userRepo repo.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeUserHasExists
	}
	return response.Success
}
