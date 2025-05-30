package service

import (
	"errors"

	"github.com/janeefajjustin/task-1/models"
	"github.com/janeefajjustin/task-1/repo"
	"github.com/janeefajjustin/task-1/utils"
)

type UserService struct {
	//UserRepo *repo.UserRepo
}

type ServiceInterface interface {
	CompareUsernameandPassword(u *models.User) error
}

// func NewUserService(userRepo *repo.UserRepo) UserService {
// 	return UserService{
// 		UserRepo: userRepo,
// 	}
// }

func (s UserService) CompareUsernameandPassword(u *models.User) error {
	// retrivedPassword, err := repo.ValidateCredentials(u)
	r := repo.RepoInterface(&repo.UserRepo{})
	retrivedPassword, err := r.ValidateCredentials(u)

	if err != nil {
		return err
	}
	if utils.CheckPasswordHash(retrivedPassword, u.Password) != true {
		return errors.New("password is invalid")
	}
	return nil

}
