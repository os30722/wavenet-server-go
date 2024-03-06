package postDb

import (
	"context"

	"github.com/hepa/wavenet/vo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postDao struct {
	db *pgxpool.Pool
}

func GetPostDao(db *pgxpool.Pool) *postDao {
	return &postDao{
		db: db,
	}
}

type PostRepo interface {
	GetPosts(ctx context.Context, userId int, pageParams *vo.PageParams) ([]vo.Post, error)
	UploadPost(ctx context.Context, upload vo.PostUpload) error

	// Interactions
	LikePost(ctx context.Context, postId int, userId int) error
	UnlikePost(ctx context.Context, postId int, userId int) error
	GetLikes(ctx context.Context, postId int, userId int, params *vo.PageParams) ([]vo.Like, error)

	AddComment(ctx context.Context, comment vo.CommentForm, userId int) error
	RemoveCommment(ctx context.Context, comment vo.CommentForm, userId int) error
	GetComments(ctx context.Context, postId int, userId int, params *vo.PageParams) ([]vo.Comment, error)
}
