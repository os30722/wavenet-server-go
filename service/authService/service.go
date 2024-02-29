package authService

import "github.com/hepa/wavenet/repository/userDb"

type authService struct {
	userRepo userDb.UserRepo
}

func NewAuthService(repo userDb.UserRepo) *authService {
	return &authService{
		userRepo: repo,
	}
}
