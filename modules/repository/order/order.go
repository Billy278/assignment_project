package repository

import (
	"context"
	"database/sql"

	models "github.com/Billy278/assignment_project/modules/models/orders"
)

type OrderRepo interface {
	RepoCreate(ctx context.Context, tx *sql.Tx, orderIn models.Order) (resOrder models.Order, err error)
	RepoList(ctx context.Context, db *sql.DB) (resOrder []models.Order, err error)
}
