package authService

import (
	"encoding/json"
	"net/http"
	"strings"

	cerr "github.com/hepa/wavenet/middleware"
	"github.com/hepa/wavenet/vo"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func (au authService) Login(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	var repo = au.userRepo
	var userInput vo.UserCred

	err := json.NewDecoder(req.Body).Decode(&userInput)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	if userInput.Email == "" && userInput.Pass == "" {
		return cerr.HttpError(err, 500)
	}

	userCred, err := repo.GetUserCred(req.Context(), strings.ToLower(userInput.Email))
	if err != nil {
		if err == pgx.ErrNoRows {
			return cerr.HttpError(err, 401)
		}
		return cerr.HttpError(err, 500)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userCred.Pass), []byte(userInput.Pass))
	if err != nil {
		return cerr.HttpError(err, 401)
	}

	json.NewEncoder(res).Encode(vo.Message{Msg: "Successful"})
	return nil
}

func (au authService) SignUp(res http.ResponseWriter, req *http.Request) *cerr.AppError {
	var form vo.UserForm
	var repo = au.userRepo
	var ctx = req.Context()

	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	// @TODO = To validate user input

	duplicates, err := repo.FindDuplicate(ctx, form.Username, form.Email)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	if len(duplicates) > 0 {
		return cerr.HttpErrorWithMsg(err, strings.Join(duplicates, "|"), 409)
	}

	_, err = repo.SignUp(ctx, form)
	if err != nil {
		return cerr.HttpError(err, 500)
	}

	json.NewEncoder(res).Encode(vo.Message{Msg: "Successful"})
	return nil

}
