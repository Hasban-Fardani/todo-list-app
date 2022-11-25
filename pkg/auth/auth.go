package auth

import (
	"errors"

	"github.com/hasban-fardani/todo-list-app/pkg/database"
	"github.com/hasban-fardani/todo-list-app/pkg/models"
)

var db = database.Mysql

func LoginWithData(userData models.User) (models.User, bool, error) {
	var err error
	var result models.User

	err = nil

	if userData.Username != "" {
		err = db.QueryRow(
			"select * from user where username = ?", userData.Username,
		).Scan(
			&result.Id, &result.NamaLengkap, &result.Username,
			&result.LastEditedAt, &result.CreateAt, &result.Password,
			&result.Email)

	} else if userData.Email != "" {
		err = db.QueryRow(
			"select * from user where email = ?", userData.Email,
		).Scan(
			&result.Id, &result.NamaLengkap, &result.Username,
			&result.LastEditedAt, &result.CreateAt, &result.Password,
			&result.Email,
		)
	}

	if err != nil {
		return result, false, err
	}
	if userData.Password != result.Password {
		return result, false, errors.New("wrong password")
	}
	return result, true, nil
}

func LoginWithToken(tokenStr string) (bool, error) {
	ok, err := ValidateToken(tokenStr)
	if err != nil || !ok {
		return false, err
	}
	return true, nil
}

func Signup(user models.User) (bool, error) {
	_, err := db.Query(
		"insert into user (namaLengkap, username, email, password) values (?, ?, ?, ?)",
		user.NamaLengkap, user.Username, user.Email, user.Password,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
