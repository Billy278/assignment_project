package models

import "time"

type User struct {
	Id         uint64
	Name       string `json:"name" validate:"required"`
	DoB        *time.Time
	DoBString  string `json:"dob" validate:"required"`
	Gmail      string `json:"gmail" validate:"required"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Created_at *time.Time
	Updated_at *time.Time
}
