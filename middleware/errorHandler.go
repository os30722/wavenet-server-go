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

		case 400:
			fmt.Println(err)
			res.WriteHeader(400)
			json.NewEncoder(res).Encode(vo.Message{Msg: "Bad Request"})

		case 401:
			fmt.Println(err)
			res.WriteHeader(401)
			json.NewEncoder(res).Encode(vo.Message{Msg: "Unauthorized"})

		case 409:
			res.WriteHeader(409)
			json.NewEncoder(res).Encode(vo.Message{Msg: err.Msg})
		}

	}
}

type AppError struct {
	Error error
	Msg   string
	Code  int
}

func HttpErrorWithMsg(err error, msg string, code int) *AppError {
	return &AppError{
		Error: err,
		Msg:   msg,
		Code:  code,
	}
}

func HttpError(err error, code int) *AppError {
	return &AppError{
		Error: err,
		Code:  code,
	}
}
