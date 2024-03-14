package postDb

import (
	"context"
	"strconv"

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

	query := `select posts.post_id, title, username as user, url, likes, post_likes.post_id is not null as user_liked, comments
		from posts join users on users.user_id=posts.user_id 
		left join post_likes on post_likes.post_id = posts.post_id and users.user_id = $1 
		where posts.status = 'Published'
			` + cursor + `order by posts.post_id desc limit $2`

	rows, err := db.Query(ctx, query, vars...)
	defer rows.Close()
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

	return posts, nil
}

func (pd postDao) GetComments(ctx context.Context, postId int, userId int, parentId int, params *vo.PageParams) ([]vo.Comment, error) {
	var db = pd.db
	var vars = make([]interface{}, 0, 5)

	vars = append(vars, postId, userId, params.PageSize)

	var cursor string = ""
	if parentId != 0 {
		vars = append(vars, parentId)
		cursor = " and parent_id =  $" + strconv.Itoa(len(vars))
	} else {
		cursor = " and parent_id is null "
	}

	if params.Cursor != 0 {
		vars = append(vars, params.Cursor)
		cursor += " and comments.comment_id < $" + strconv.Itoa(len(vars))
	}

	query := `select comment_id, content, username, replies_count from comments 
	join users on comments.user_id = users.user_id where post_id=$1 ` + cursor + ` order by comments.user_id=$2 desc, 
	comment_id desc limit $3`

	rows, err := db.Query(ctx, query, vars...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := make([]vo.Comment, 0, params.PageSize)
	for rows.Next() {
		var comment vo.Comment
		err = rows.Scan(&comment.CommentId, &comment.Content, &comment.Username, &comment.RepliesCount)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (pd postDao) GetLikes(ctx context.Context, postId int, userId int, params *vo.PageParams) ([]vo.Like, error) {
	var db = pd.db
	var vars = make([]interface{}, 0, 4)

	vars = append(vars, postId, userId, params.PageSize)

	var cursor string = ""
	if params.Cursor != 0 {
		cursor = " and post_likes.like_id < $4 "
		vars = append(vars, params.Cursor)
	}

	query := `select username from post_likes
	join users on post_likes.user_id = users.user_id where post_id=$1 ` + cursor + ` order by post_likes.user_id=$2 desc, 
	like_id desc limit $3`

	rows, err := db.Query(ctx, query, vars...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	likes := make([]vo.Like, 0, params.PageSize)
	for rows.Next() {
		var like vo.Like
		err = rows.Scan(&like.Username)
		if err != nil {
			return likes, err
		}
		likes = append(likes, like)
	}

	return likes, nil
}
