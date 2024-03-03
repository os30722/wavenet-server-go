package userDb

import (
	"context"

	"github.com/hepa/wavenet/vo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userDao struct {
	db *pgxpool.Pool
}

func GetUserDao(db *pgxpool.Pool) *userDao {
	return &userDao{
		db: db,
	}
}

type UserRepo interface {
	GetUserCred(ctx context.Context, email string) (*vo.UserCred, error)
	SignUp(ctx context.Context, form vo.UserForm) error
	FindDuplicate(ctx context.Context, username string, email string) ([]string, error)
}
