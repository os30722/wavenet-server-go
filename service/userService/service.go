package userService

import (
	"github.com/gorilla/mux"
	md "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/repository/userDb"
)

type userService struct {
	userRepo userDb.UserRepo
}

func newService(repo userDb.UserRepo) *userService {
	return &userService{
		userRepo: repo,
	}
}

func RegisterService(router *mux.Router, repo userDb.UserRepo) {
	userService := newService(repo)
	router.Handle("/login", md.ErrHandler(userService.Login)).Methods("Post")
	router.Handle("/signup", md.ErrHandler(userService.SignUp)).Methods("Post")
}
