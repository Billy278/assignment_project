package models

import "time"

type Promo struct {
	Id         uint64
	Kode_Promo string `json:"kode_promo" validate:"required"`
	Token      string
	Created_at *time.Time
}
