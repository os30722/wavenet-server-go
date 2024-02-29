package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hepa/wavenet/vo"
)

var errStr = []byte(`Internal Server Error`)

type ErrHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn ErrHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	if err := fn(res, req); err != nil {
		switch err.Code {
		case 500:
			fmt.Println(err)
			res.WriteHeader(500)
			res.Write(errStr)
			break

		case 400:
			fmt.Println(err)
			res.WriteHeader(400)
			res.Write([]byte(`{"err":"Bad Request"}`))
			break

		case 401:
			fmt.Println(err)
			res.WriteHeader(401)
			res.Write([]byte(`{"msg":"Unauthorized"}`))
			break

		case 409:
			res.WriteHeader(409)
			json.NewEncoder(res).Encode(vo.Message{Msg: err.Msg})
			break
		}

	}
}

type AppError struct {
	Error error
	Msg   string
	Code  int
}

func HttpError(err error, msg string, code int) *AppError {
	return &AppError{
		Error: err,
		Msg:   msg,
		Code:  code,
	}
}
