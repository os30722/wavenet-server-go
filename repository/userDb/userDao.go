package userDb

import (
	"context"

	"github.com/hepa/wavenet/vo"
	"golang.org/x/crypto/bcrypt"
)

func (ur userDao) GetUserCred(ctx context.Context, email string) (*vo.UserCred, error) {
	var db = ur.db
	var userCred vo.UserCred

	err := db.QueryRow(ctx, "select user_id, email, password from users where email=$1 limit 1", email).Scan(&userCred.Id, &userCred.Email, &userCred.Pass)
	if err != nil {
		return nil, err
	}

	return &userCred, nil
}

func (ur userDao) SignUp(ctx context.Context, form vo.UserForm) error {
	var db = ur.db
	hash, err := bcrypt.GenerateFromPassword([]byte(form.Pass), 10)
	if err != nil {
		return err
	}
	_, err = db.Exec(ctx, "insert into  users(name,dob,gender,username,email,password) values($1,$2,$3,$4,$5,$6)",
		form.Name, form.Dob, string(form.Gender[0]), form.Username, form.Email, hash)
	if err != nil {
		return err
	}

	return nil
}

func (ur userDao) FindDuplicate(ctx context.Context, username string, email string) ([]string, error) {
	var db = ur.db
	var res struct {
		email    string
		username string
	}
	duplicates := make([]string, 0, 2)
	rows, err := db.Query(ctx, "select email, username from users where email=$1 or username=$2", email, username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&res.email, &res.username)
		if err != nil {
			return nil, err
		}

		if username == res.username {
			duplicates = append(duplicates, "username")
		}
		if email == res.email {
			duplicates = append(duplicates, "email")
		}
	}

	return duplicates, nil
}
