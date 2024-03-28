package repository

import (
	"context"
	"database/sql"

	models "github.com/Billy278/assignment_project/modules/models/products"
)

type ProductRepo interface {
	RepoList(ctx context.Context) (resProduct []models.Product, err error)
	RepoFindByid(ctx context.Context, id uint64) (resProduct models.Product, err error)
	RepoCreate(ctx context.Context, productIn models.Product) (resProduct models.Product, err error)
	RepoDelete(ctx context.Context, id uint64) (err error)
	RepoUpdateTx(ctx context.Context, db *sql.Tx, productIn models.Product) (resProduct models.Product, err error)
	RepoFindByidTx(ctx context.Context, db *sql.Tx, id uint64) (resProduct models.Product, err error)
}
