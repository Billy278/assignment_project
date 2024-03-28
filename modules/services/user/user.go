package services

import (
	"context"

	models "github.com/Billy278/assignment_project/modules/models/user"
)

type UserSrv interface {
	GetAllUserIsbrithday(ctx context.Context) (resUser []models.User, err error)
	CreatedUser(ctx context.Context, userIn models.User) (err error)
	SrvFindUser(ctx context.Context, username string) (err error)
}
