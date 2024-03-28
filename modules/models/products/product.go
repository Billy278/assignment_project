package models

import "time"

type Product struct {
	Id         uint64
	Name       string  `json:"name" validate:"required"`
	Stock      uint64  `json:"stock" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
	Created_At *time.Time
	Updated_At *time.Time
}
