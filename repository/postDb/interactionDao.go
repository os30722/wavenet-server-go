package postDb

import (
	"context"
)

func (pd postDao) LikePost(ctx context.Context, postId int, userId int) error {
	var db = pd.db
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "insert into likes(post_id,user_id) values($1,$2)", postId, userId)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "update posts set likes=likes+1 where post_id=$1", postId)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
