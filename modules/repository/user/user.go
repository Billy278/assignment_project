package repository

import (
	"context"

	models "github.com/Billy278/assignment_project/modules/models/user"
)

type UserRepo interface {
	GetAllUserIsBirthday(ctx context.Context) (resUser []models.User, err error)
}
