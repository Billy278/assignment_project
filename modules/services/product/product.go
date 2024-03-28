package services

import (
	"context"

	models "github.com/Billy278/assignment_project/modules/models/products"
)

type SrvProduct interface {
	SrvList(ctx context.Context) (resProduct []models.Product, err error)
	SrvFindByid(ctx context.Context, id uint64) (resProduct models.Product, err error)
	SrvCreate(ctx context.Context, productIn models.Product) (resProduct models.Product, err error)
	SrvDelete(ctx context.Context, id uint64) (err error)
}
