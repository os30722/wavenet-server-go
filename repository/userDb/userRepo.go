package userDb

import (
	"context"

	"github.com/hepa/wavenet/vo"
)

type UserRepo interface {
	GetUserCred(ctx context.Context, email string) (*vo.UserCred, error)
	SignUp(ctx context.Context, form vo.UserForm) (int, error)
	FindDuplicate(ctx context.Context, username string, email string) ([]string, error)
}
