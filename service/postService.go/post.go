package postService

import (
	"encoding/json"
	"net/http"
	"strings"

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
	reader, err := utils.GetPartReader(req)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	ctx := req.Context()
	uid, err := utils.GetUid(ctx)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	repo := po.postRepo

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
