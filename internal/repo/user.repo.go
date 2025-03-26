package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (userRepo *UserRepo) GetUserByEmail(email string) (error, string) {
	return nil, "pong"
}
