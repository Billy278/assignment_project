package services

import (
	"context"
	"fmt"
	"time"

	models "github.com/Billy278/assignment_project/modules/models/user"
	repository "github.com/Billy278/assignment_project/modules/repository/user"
	"github.com/Billy278/assignment_project/pkg/crypto"
)

type UserSrvImpl struct {
	RepoUser repository.UserRepo
}

func NewUserSrvImpl(repouser repository.UserRepo) UserSrv {
	return &UserSrvImpl{
		RepoUser: repouser,
	}
}
func (srv *UserSrvImpl) CreatedUser(ctx context.Context, userIn models.User) (err error) {
	fmt.Println("services CreatedUser")
	tNow := time.Now()
	userIn.Created_at = &tNow
	userIn.Updated_at = &tNow
	// hach password
	hashPass, err := crypto.GenereteHash(userIn.Password)
	if err != nil {
		return
	}
	userIn.Password = hashPass
	err = srv.RepoUser.CreatedUser(ctx, userIn)
	if err != nil {
		return
	}
	return
}
func (srv *UserSrvImpl) GetAllUserIsbrithday(ctx context.Context) (resUser []models.User, err error) {
	fmt.Println("services GetAllUserIsbrithday")
	resUser, err = srv.RepoUser.GetAllUserIsBirthday(ctx)
	if err != nil {
		return
	}

	return
}

func (srv *UserSrvImpl) SrvFindUser(ctx context.Context, username string) (err error) {
	fmt.Println("SrvFindUser")
	err = srv.RepoUser.RepoFindUser(ctx, username)
	return
}
