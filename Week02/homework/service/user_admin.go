package service

import (
	"geek.com/lesson3-4/homework/model"
	"github.com/pkg/errors"
)

func GetAllUser() ([]model.User, error) {
	var users = make([]model.User, 0)
	users, query_err := model.QueryUser()

	if query_err != nil {
		return nil, errors.WithMessage(query_err, "query user info error")
	}

	return users, nil

}
