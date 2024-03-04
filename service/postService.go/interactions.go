package postService

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	cerr "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/utils"
	"github.com/hepa/wavenet/vo"
)

func (po postService) LikeUnlikePost(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	repo := po.postRepo
	vars := mux.Vars(req)

	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return cerr.HttpError(errors.New("require post id"), 500)
	}

	ctx := req.Context()
	uid, err := utils.GetUid(ctx)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	action := req.URL.Query().Get("action")

	if action == "add" {
		err = repo.LikePost(ctx, postId, uid)
	} else if action == "remove" {
		err = repo.UnlikePost(ctx, postId, uid)
	} else {
		return cerr.HttpError(err, 500)
	}

	if err != nil {
		return cerr.HttpError(err, 500)
	}

	return nil
}

func (po postService) Comment(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	repo := po.postRepo

	ctx := req.Context()
	uid, err := utils.GetUid(ctx)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	var commment vo.Comment
	err = json.NewDecoder(req.Body).Decode(&commment)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	action := req.URL.Query().Get("action")

	if action == "add" {
		err = repo.AddComment(ctx, commment, uid)
	} else if action == "remove" {
		err = repo.RemoveCommment(ctx, commment, uid)
	} else {
		return cerr.HttpError(err, 500)
	}

	if err != nil {
		return cerr.HttpError(err, 500)
	}

	return nil
}
