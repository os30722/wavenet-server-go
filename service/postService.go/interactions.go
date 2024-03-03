package postService

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	cerr "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/utils"
)

func (po postService) LikeUnlikePost(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	repo := po.postRepo
	vars := mux.Vars(req)

	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return cerr.HttpError(errors.New("require post id"), 500)
	}

	action := req.URL.Query().Get("action")
	if action != "add" && action != "remove" {
		return cerr.HttpError(err, 500)
	}

	ctx := req.Context()
	uid, err := utils.GetUid(ctx)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	if action == "add" {
		err = repo.LikePost(ctx, postId, uid)
	} else {
		err = repo.UnlikePost(ctx, postId, uid)
	}

	if err != nil {
		return cerr.HttpError(err, 500)
	}

	return nil
}
