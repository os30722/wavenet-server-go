package postDb

import (
	"context"
	"errors"
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

func (pd postDao) UnlikePost(ctx context.Context, postId int, userId int) error {
	var db = pd.db
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	commandTag, err := tx.Exec(ctx, "delete from likes where post_id=$1 and user_id=$2", postId, userId)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("row not found")
	}

	_, err = tx.Exec(ctx, "update posts set likes=likes-1 where post_id=$1", postId)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (pd postDao) AddComment(ctx context.Context, postId int, userId int, content string) error {
	var db = pd.db
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "insert into comments(post_id,user_id,content) values($1,$2$3)", postId, userId, content)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "update posts set coments=comments+1 where post_id=$1", postId)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (pd postDao) RemoveCommment(ctx context.Context, userId int, commentId int) error {
	var db = pd.db
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	commandTag, err := tx.Exec(ctx, "delete from comments where comment_id=$1 and user_id=$2 returning post_id", commentId, userId)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("row not found")
	}

	_, err = tx.Exec(ctx, "update posts set comments=comments-1 where post_id=$1", postId)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
