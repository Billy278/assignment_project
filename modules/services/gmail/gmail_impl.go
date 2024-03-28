package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	models "github.com/Billy278/assignment_project/modules/models/gmail"
	repository "github.com/Billy278/assignment_project/modules/repository/gmail"
)

type GmailSrvImpl struct {
	GmailRepo repository.Gmail
}

func NewGmailSrvImpl(gmairepo repository.Gmail) GmailSrv {
	return &GmailSrvImpl{
		GmailRepo: gmairepo,
	}
}

func (srv *GmailSrvImpl) Created(ctx context.Context, gmailIn models.Gmail) (err error) {
	fmt.Println("Gmail Services")
	tNow := time.Now()
	gmailIn.Created_at = &tNow
	gmailIn.Message = fmt.Sprintf("happy birthday to our beloved customer %v, today especially for you we are providing a special birthday promo code with the promo code %v  with a 20 percent discount for all product valid 1 x 24 after the message is sent ", gmailIn.Name, gmailIn.Promo)
	err = srv.GmailRepo.Created(ctx, gmailIn)
	if err != nil {
		err = errors.New("Faild to send Message To gmail")
		return
	}
	return err
}
