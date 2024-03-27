package services

import (
	"context"

	models "github.com/Billy278/assignment_project/modules/models/gmail"
)

type GmailSrv interface {
	Created(ctx context.Context, gmailIn models.Gmail) (err error)
}
