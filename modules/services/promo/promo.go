package services

import (
	"context"
)

type PromoSrv interface {
	CreatedPromo(ctx context.Context) (err error)
	GetToken(ctx context.Context, promo string) (resToken string, err error)
}
