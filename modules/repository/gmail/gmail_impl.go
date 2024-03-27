package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/Billy278/assignment_project/modules/models/gmail"
)

type GmailRepoImpl struct {
	DB *sql.DB
}

func NewGmailRepoImpl(db *sql.DB) Gmail {
	return &GmailRepoImpl{
		DB: db,
	}
}

func (repo *GmailRepoImpl) Created(ctx context.Context, gmailIn models.Gmail) (err error) {
	fmt.Println("Repo gmail Impl")
	sql := "INSERT INTO gmail (message,receiver,created_at) values($1,$2,$3)"

	_, err = repo.DB.ExecContext(ctx, sql, gmailIn.Message, gmailIn.Receiver, gmailIn.Created_at)
	if err != nil {
		return
	}
	return
}
