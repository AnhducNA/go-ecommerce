package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AnhducNA/go-ecommerce/internal/repo"
	"github.com/AnhducNA/go-ecommerce/internal/utils/crypto"
	"github.com/AnhducNA/go-ecommerce/internal/utils/random"
	"github.com/AnhducNA/go-ecommerce/internal/utils/sendto"
	"github.com/AnhducNA/go-ecommerce/pkg/response"
)

// INTERFACE VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// hash Email:
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail:: %s\n", hashEmail)
	// check email exist
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeUserHasExists
	}
	// add otp
	otp := random.GenerateSixDigitOTP()
	fmt.Printf("OTP:: %d\n", otp)
	// save OTP in redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrorInvalidOTP
	}
	// send email OTP
	err = sendto.SendTemplateEmailOTP(
		[]string{email},
		"daiphatmanagement@gmail.com",
		"otp-auth.html",
		map[string]interface{}{
			"otp": strconv.Itoa(otp),
		},
	)

	if err != nil {
		return response.ErrorSendEmailOTP
	}
	return response.Success
}
