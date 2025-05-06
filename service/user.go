package service

import (
	"errors"

	"github.com/janeefajjustin/task-1/models"
	"github.com/janeefajjustin/task-1/repo"
	"github.com/janeefajjustin/task-1/utils"
)

func CompareUsernameandPassword(u *models.User) error {
	retrivedPassword, err := repo.ValidateCredentials(u)
	if err != nil {
		return err
	}
	if utils.CheckPasswordHash(retrivedPassword, u.Password) != true {
		return errors.New("password is invalid")
	}
	return nil

}
