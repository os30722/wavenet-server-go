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

func (pd postDao) GetPosts(ctx context.Context, params vo.PageParams) ([]vo.Post, error) {
	var db = pd.db
	var vars = make([]interface{}, 0, 2)
	vars = append(vars, params.PageSize)
	var cursor string = ""
	if params.Cursor != 0 {
		cursor = " and posts.Post_id < $2 "
		vars = append(vars, params.Cursor)
	}

	var query = `select post_id, title, name as user, url from posts join users on users.user_id=posts.user_id where posts.status = 'Published'
			` + cursor + `order by posts.post_id desc limit $1`

	rows, err := db.Query(ctx, query, vars...)
	if err != nil {
		return nil, err
	}

	posts := make([]vo.Post, 0, params.PageSize)
	for rows.Next() {
		var post vo.Post
		err = rows.Scan(&post.PostId, &post.Title, &post.User, &post.Url)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	rows.Close()
	return posts, nil
}
