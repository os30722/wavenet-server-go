package postDb

import (
	"context"
	"errors"

	"github.com/hepa/wavenet/vo"
	"github.com/jackc/pgx/v5/pgconn"
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

func (pd postDao) AddComment(ctx context.Context, comment vo.Comment, userId int) error {
	var db = pd.db
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var commandTag pgconn.CommandTag

	if comment.ParentId == 0 {
		_, err = tx.Exec(ctx, "insert into comments(post_id,user_id,content) values($1,$2,$3)", comment.PostId, userId, comment.Content)
	} else {

		commandTag, err = tx.Exec(ctx, `insert into comments(post_id, user_id,content,parent_id) select 
			$1,$2,$3,$4
		 where exists (select 1 from comments where post_id=$1 and comment_id=$4)`,
			comment.PostId, userId, comment.Content, comment.ParentId)
		if err != nil {
			return err
		}
		if commandTag.RowsAffected() == 0 {
			return errors.New("wrong insert")
		}

		_, err = tx.Exec(ctx, "update comments set child_comments=child_comments+1 where comment_id=$1", comment.ParentId)
	}

	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "update posts set comments=comments+1 where post_id=$1", comment.PostId)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (pd postDao) RemoveCommment(ctx context.Context, comment vo.Comment, userId int) error {
	var db = pd.db
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var commandTag pgconn.CommandTag
	rowsAffected := 0
	commandTag, err = tx.Exec(ctx, "delete from comments where parent_id = $1", comment.ComentId)
	if err != nil {
		return err
	}
	rowsAffected += int(commandTag.RowsAffected())

	var parentId, postId int
	err = tx.QueryRow(ctx, "delete from comments where comment_id=$1 and user_id=$2 returning coalesce(parent_id, 0), coalesce(post_id,0)",
		comment.ComentId, userId).Scan(&parentId, &postId)

	if err != nil {
		return err
	}

	if postId == 0 {
		return errors.New("row not found")
	}
	rowsAffected += 1

	if parentId != 0 {
		_, err = tx.Exec(ctx, "update comments set child_comments=child_comments-1 where comment_id=$1", comment.ParentId)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(ctx, "update posts set comments=comments-$2 where post_id=$1", postId, rowsAffected)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
