package utils

import (
	"net/http"
	"strconv"

	"github.com/hepa/wavenet/vo"
)

func GetPagePramas(req *http.Request) (*vo.PageParams, error) {
	params := req.URL.Query()

	cursor, err := strconv.Atoi(params.Get("cursor"))
	if err != nil {
		return nil, err
	}

	pageSize, err := strconv.Atoi(params.Get("pagesize"))
	if err != nil || cursor < 0 || pageSize > 200 || pageSize < 1 {
		return nil, err
	}

	return &vo.PageParams{
		Cursor:   cursor,
		PageSize: pageSize,
	}, nil
}
