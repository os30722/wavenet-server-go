package postDb

import (
	"context"

	"github.com/hepa/wavenet/vo"
	"github.com/jackc/pgx/v5"
)

func (pd postDao) UploadPost(ctx context.Context, upload vo.PostUpload) error {
	var db = pd.db
	_, err := db.Exec(ctx, "insert into posts(title, description, status, url, user_id) values ($1, $2, $3, $4, $5)", upload.Title, upload.Description,
		"Published", upload.FileName, upload.UserId)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	return nil
}

func (pd postDao) GetPosts(ctx context.Context, userId int, params *vo.PageParams) ([]vo.Post, error) {

	var db = pd.db
	var vars = make([]interface{}, 0, 3)

	vars = append(vars, userId, params.PageSize)

	var cursor string = ""
	if params.Cursor != 0 {
		cursor = " and posts.Post_id < $3 "
		vars = append(vars, params.Cursor)
	}

	var query = `select posts.post_id, title, username as user, url, likes, likes.post_id is not null as user_liked, comments
		from posts join users on users.user_id=posts.user_id 
		left join likes on likes.post_id = posts.post_id and users.user_id = $1 
		where posts.status = 'Published'
			` + cursor + `order by posts.post_id desc limit $2`

	rows, err := db.Query(ctx, query, vars...)
	if err != nil {
		return nil, err
	}

	posts := make([]vo.Post, 0, params.PageSize)
	for rows.Next() {
		var post vo.Post
		err = rows.Scan(&post.PostId, &post.Title, &post.UserName, &post.Url, &post.Likes, &post.UserLiked, &post.Comments)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	rows.Close()
	return posts, nil
}
