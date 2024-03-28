package models

import (
	"time"
)

type Order struct {
	Id          uint64
	UserId      uint64 `json:"user_id" validate:"required"`
	ProductId   uint64 `json:"product_id" validate:"required"`
	Qty         uint64 `json:"qty" validate:"required"`
	Promo       float64
	Promo_code  string `json:"promo_code"`
	TotalPrize  float64
	TotalPaid   float64 `json:"total_paid" validate:"required,number"`
	TotalReturn float64
	Created_At  *time.Time
}
