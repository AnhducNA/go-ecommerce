package service

import "github.com/AnhducNA/go-ecommerce/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) GetUserByEmail(email string) (error, string) {
	err, user := userService.userRepo.GetUserByEmail(email)
	if err != nil {
		return err, ""
	}
	return nil, user
}
