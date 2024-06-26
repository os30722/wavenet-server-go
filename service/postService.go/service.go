package postService

import (
	"github.com/gorilla/mux"
	md "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/repository/postDb"
)

type postService struct {
	postRepo postDb.PostRepo
}

func newService(repo postDb.PostRepo) *postService {
	return &postService{
		postRepo: repo,
	}
}

func RegisterService(router *mux.Router, repo postDb.PostRepo) {
	postService := newService(repo)
	router.Handle("/getPosts", md.ErrHandler(postService.GetPosts)).Methods("GET")
	router.Handle("/upload", md.ErrHandler(postService.UploadPost)).Methods("POST")
	router.Handle("/upload", md.ErrHandler(postService.UploadPost)).Methods("POST")

	router.Handle("/like/{id}", md.ErrHandler(postService.LikeUnlikePost)).Methods("POST")
	router.Handle("/likes/{id}", md.ErrHandler(postService.GetLikes)).Methods("GET")

	router.Handle("/comment", md.ErrHandler(postService.Comment)).Methods("POST")
	router.Handle("/comments/{id}", md.ErrHandler(postService.GetComments)).Methods("GET")
}
