package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	models "github.com/Billy278/assignment_project/modules/models/promo"
)

type PromoRepoImpl struct {
	DB *sql.DB
}

func NewPromoRepoImpl(db *sql.DB) PromoRepo {
	return &PromoRepoImpl{
		DB: db,
	}
}
func (repo *PromoRepoImpl) Created(ctx context.Context, promoIn models.Promo) (err error) {
	fmt.Println("Repo Promo")
	sql := "INSERT INTO promo (kode_promo,token,created_at)VALUES ($1,$2,$3) ON CONFLICT (kode_promo) DO UPDATE set token = $4"
	_, err = repo.DB.ExecContext(ctx, sql, promoIn.Kode_Promo, promoIn.Token, promoIn.Created_at, promoIn.Token)
	if err != nil {
		return err
	}
	return
}

func (repo *PromoRepoImpl) GetToken(ctx context.Context, Promo string) (resToken string, err error) {
	fmt.Println("repo Gettoken")
	sql := "SELECT token FROM promo WHERE kode_promo=$1"
	rows, err := repo.DB.QueryContext(ctx, sql, Promo)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&resToken)
		if err != nil {
			return
		}
	} else {
		err = errors.New("NOT FOUND")
	}
	return
}
