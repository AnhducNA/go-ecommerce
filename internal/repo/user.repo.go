package repo

import (
	"github.com/AnhducNA/go-ecommerce/global"
	"github.com/AnhducNA/go-ecommerce/internal/model"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (userRepo *UserRepo) GetUserByEmail(email string) (error, string) {
// 	return nil, "pong"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func (*userRepository) GetUserByEmail(email string) bool {
	row := global.Mdb.Table(TableNameGoCrmUser).Where("email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
