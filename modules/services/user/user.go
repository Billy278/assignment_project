package services

import (
	"context"

	models "github.com/Billy278/assignment_project/modules/models/user"
)

type UserSrv interface {
	GetAllUserIsbrithday(ctx context.Context) (resUser []models.User, err error)
}
