package repository

import (
	"context"

	models "github.com/Billy278/assignment_project/modules/models/promo"
)

type PromoRepo interface {
	Created(ctx context.Context, promoIn models.Promo) (err error)
	GetToken(ctx context.Context, Promo string) (resToken string, err error)
}
