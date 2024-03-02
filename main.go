package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hepa/wavenet/database"
	"github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/repository/postDb"
	"github.com/hepa/wavenet/repository/userDb"
	"github.com/hepa/wavenet/service/postService.go"
	"github.com/hepa/wavenet/service/userService"
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

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	db, err := database.GetPostgres()
	if err != nil || db == nil {
		log.Fatal(err)
	}

	userDao := userDb.GetUserDao(db)
	postDao := postDb.GetPostDao(db)

	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(middleware.Authenticator)

	userRouter := router.PathPrefix("/user").Subrouter()
	userService.RegisterService(userRouter, userDao)

	postRouter := authRouter.PathPrefix("/posts").Subrouter()
	postService.RegisterService(postRouter, postDao)

	router.HandleFunc("/", homeRoute)

	return router
}
