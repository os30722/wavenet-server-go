package postService

import (
	"encoding/json"
	"net/http"

	cerr "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/utils"
	"github.com/hepa/wavenet/vo"
)

func (po postService) GetPosts(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	pageParams, err := utils.GetPagePramas(req)
	if err != nil {
		return cerr.HttpError(err, "", 400)
	}

	ctx := req.Context()
	repo := po.postRepo

	posts, err := repo.GetPosts(ctx, *pageParams)
	if err != nil {
		return cerr.HttpError(err, "", 500)
	}

	page := *&vo.PageItem{
		TotalCounts: len(posts),
		Items:       posts,
	}

	json.NewEncoder(res).Encode(page)

	return nil
}
