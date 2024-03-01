package authService

import (
	"github.com/gorilla/mux"
	md "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/repository/userDb"
)

type authService struct {
	userRepo userDb.UserRepo
}

func newService(repo userDb.UserRepo) *authService {
	return &authService{
		userRepo: repo,
	}
}

func RegisterService(router *mux.Router, repo userDb.UserRepo) {
	authService := newService(repo)
	router.Handle("/login", md.ErrHandler(authService.Login)).Methods("Post")
	router.Handle("/signup", md.ErrHandler(authService.SignUp)).Methods("Post")
}
