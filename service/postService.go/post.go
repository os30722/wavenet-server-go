package postService

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	cerr "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/utils"
	"github.com/hepa/wavenet/utils/ffmpeg"
	"github.com/hepa/wavenet/vo"
)

func (po postService) GetPosts(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	pageParams, err := utils.GetPagePramas(req)
	if err != nil {
		return cerr.HttpError(err, 400)
	}

	ctx := req.Context()
	uid, err := utils.GetUid(ctx)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	repo := po.postRepo

	posts, err := repo.GetPosts(ctx, uid, pageParams)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	page := &vo.PageItem{
		TotalCounts: len(posts),
		Items:       posts,
	}

	json.NewEncoder(res).Encode(page)

	return nil
}

func (po postService) UploadPost(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	repo := po.postRepo

	reader, err := utils.GetPartReader(req)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	ctx := req.Context()
	uid, err := utils.GetUid(ctx)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	var upload vo.PostUpload
	upload.UserId = uid

	upload.Title, err = reader.NextTextPart("title")
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	upload.Description, err = reader.NextTextPart("description")
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	fileName, err := reader.NextFilePart(utils.FilePart{
		Field: "recording",
	})
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	upload.FileName = strings.Split(fileName, ".")[0]
	upload.Extention = strings.Split(fileName, ".")[1]

	go ffmpeg.EncodeAudioFile(upload.FileName, upload.Extention)

	err = repo.UploadPost(ctx, upload)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	return nil
}

func (po postService) GetComments(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	repo := po.postRepo
	vars := mux.Vars(req)

	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return cerr.HttpError(errors.New("require post id"), 500)
	}

	pageParams, err := utils.GetPagePramas(req)
	if err != nil {
		return cerr.HttpError(err, 400)
	}

	ctx := req.Context()
	uid, err := utils.GetUid(ctx)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	comments, err := repo.GetComments(ctx, postId, uid, pageParams)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	page := &vo.PageItem{
		TotalCounts: len(comments),
		Items:       comments,
	}

	json.NewEncoder(res).Encode(page)

	return nil
}

func (po postService) GetLikes(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	repo := po.postRepo
	vars := mux.Vars(req)

	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return cerr.HttpError(errors.New("require post id"), 500)
	}

	pageParams, err := utils.GetPagePramas(req)
	if err != nil {
		return cerr.HttpError(err, 400)
	}

	ctx := req.Context()
	uid, err := utils.GetUid(ctx)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	likes, err := repo.GetLikes(ctx, postId, uid, pageParams)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	page := &vo.PageItem{
		TotalCounts: len(likes),
		Items:       likes,
	}

	json.NewEncoder(res).Encode(page)

	return nil
}
