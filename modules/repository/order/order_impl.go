package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/Billy278/assignment_project/modules/models/orders"
)

type OrderRepoImpl struct {
}

func NewOrderRepoImpl() OrderRepo {
	return &OrderRepoImpl{}
}

func (repo *OrderRepoImpl) RepoCreate(ctx context.Context, tx *sql.Tx, orderIn models.Order) (resOrder models.Order, err error) {
	fmt.Println("RepoOrderCreate")
	sqlCreate := "INSERT INTO orders(user_id,product_id,qty,promo,promo_code,total_price,total_paid,total_return, created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id"
	row, err := tx.QueryContext(ctx, sqlCreate, orderIn.UserId, orderIn.ProductId, orderIn.Qty, orderIn.Promo, orderIn.Promo_code, orderIn.TotalPrize, orderIn.TotalPaid, orderIn.TotalReturn, orderIn.Created_At)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resOrder.Id)
		if err != nil {
			return
		}
	}
	resOrder.UserId = orderIn.UserId
	resOrder.ProductId = orderIn.ProductId
	resOrder.Promo = orderIn.Promo
	resOrder.Qty = orderIn.Qty
	resOrder.TotalPrize = orderIn.TotalPrize
	resOrder.TotalPaid = orderIn.TotalPaid
	resOrder.TotalReturn = orderIn.TotalReturn
	resOrder.Created_At = orderIn.Created_At
	return
}
func (repo *OrderRepoImpl) RepoList(ctx context.Context, db *sql.DB) (resOrder []models.Order, err error) {
	fmt.Println("RepoOrderList")
	sqlList := `SELECT o.id,o.user_id,o.product_id,o.promo,o.total_price,o.total_paid,o.total_return, o.created_at FROM orders AS o`
	row, err := db.QueryContext(ctx, sqlList)
	if err != nil {
		return
	}
	defer row.Close()
	order := models.Order{}

	for row.Next() {
		err = row.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Promo, &order.TotalPrize, &order.TotalPaid, &order.TotalReturn, &order.Created_At)
		if err != nil {
			return
		}
		resOrder = append(resOrder, order)
	}

	return
}
