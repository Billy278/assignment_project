package models

import "time"

type User struct {
	Id         uint64
	Name       string
	DoB        *time.Time
	Gmail      string
	Username   string
	Password   string
	Created_at *time.Time
	Updated_at *time.Time
}
