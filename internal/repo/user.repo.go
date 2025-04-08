package repo

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
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
