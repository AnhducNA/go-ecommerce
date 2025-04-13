package repo

import (
	"fmt"
	"time"

	"github.com/AnhducNA/go-ecommerce/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}
type userAuthRepository struct {
}

func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	// panic("unimplemented")
	key := fmt.Sprintf("user:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
