package services

import (
	"context"

	models "github.com/Billy278/assignment_project/modules/models/orders"
)

type OrderSrv interface {
	SrvList(ctx context.Context) (resOrder []models.Order, err error)
	SrvCreate(ctx context.Context, orderIn models.Order, isgetPromo bool) (resOrder models.Order, err error)
	SrvGetTokenWithPromo(ctx context.Context, kode_promo string) (token string, err error)
}
