package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hepa/wavenet/database"
	md "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/repository/postDb"
	"github.com/hepa/wavenet/repository/userDb"
	"github.com/hepa/wavenet/service/authService"
	"github.com/hepa/wavenet/service/postService.go"
)

func main() {
	port := ":3000"

	server := http.Server{
		Addr:    port,
		Handler: router(),
	}

	server.ListenAndServe()
}

func homeRoute(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`{"msg": "At you service sir"}`))
}

func router() *mux.Router {
	router := mux.NewRouter()

	db, err := database.GetPostgres()
	if err != nil || db == nil {
		log.Fatal(err)
	}

	userDao := userDb.GetUserDao(db)
	postDao := postDb.GetPostDao(db)

	authService := authService.NewAuthService(userDao)
	router.Handle("/auth/login", md.ErrHandler(authService.Login)).Methods("Post")
	router.Handle("/auth/signup", md.ErrHandler(authService.SignUp)).Methods("Post")

	postService := postService.NewAuthService(postDao)
	router.Handle("/posts/getPost", md.ErrHandler(postService.GetPosts)).Methods("GET")

	router.HandleFunc("/", homeRoute)

	return router
}
