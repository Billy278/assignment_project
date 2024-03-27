package services

import (
	"context"
	"fmt"

	models "github.com/Billy278/assignment_project/modules/models/user"
	repository "github.com/Billy278/assignment_project/modules/repository/user"
)

type UserSrvImpl struct {
	RepoUser repository.UserRepo
}

func NewUserSrvImpl(repouser repository.UserRepo) UserSrv {
	return &UserSrvImpl{
		RepoUser: repouser,
	}
}
func (srv *UserSrvImpl) GetAllUserIsbrithday(ctx context.Context) (resUser []models.User, err error) {
	fmt.Println("services GetAllUserIsbrithday")
	resUser, err = srv.RepoUser.GetAllUserIsBirthday(ctx)
	if err != nil {
		return
	}

	return
}
