package postService

import "github.com/hepa/wavenet/repository/postDb"

type postService struct {
	postRepo postDb.PostRepo
}

func NewAuthService(repo postDb.PostRepo) *postService {
	return &postService{
		postRepo: repo,
	}
}
