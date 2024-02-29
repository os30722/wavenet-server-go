package postDb

import (
	"context"

	"github.com/hepa/wavenet/vo"
)

type PostRepo interface {
	GetPosts(ctx context.Context, pageParams vo.PageParams) ([]vo.Post, error)
}
