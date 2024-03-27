package models

import "time"

type Gmail struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name" validate:"required"`
	Promo      string `json:"promo" validate:"required"`
	Message    string `json:"message" `
	Receiver   string `json:"receiver" validate:"required"`
	Created_at *time.Time
}
