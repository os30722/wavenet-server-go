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
	LikePost(ctx context.Context, postId int, userId int) error
}
