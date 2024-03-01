package postDb

import (
	"context"

	"github.com/hepa/wavenet/vo"
	"github.com/jackc/pgx/v5"
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

func (pd postDao) UploadPost(ctx context.Context, upload vo.PostUpload) error {
	var db = pd.db
	err := db.QueryRow(ctx, "insert into posts(title, description, status, url) values ($1, $2, $3, $4)", upload.Title, upload.Description,
		"Published", upload.FileName).Scan()
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	return nil
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
