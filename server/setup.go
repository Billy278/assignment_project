package server

import (
	"github.com/Billy278/assignment_project/db"
	controllers "github.com/Billy278/assignment_project/modules/controllers/gmail"
	ctrlPromo "github.com/Billy278/assignment_project/modules/controllers/promo"
	ctrlUser "github.com/Billy278/assignment_project/modules/controllers/user"
	repository "github.com/Billy278/assignment_project/modules/repository/gmail"
	repoPromo "github.com/Billy278/assignment_project/modules/repository/promo"
	repoUser "github.com/Billy278/assignment_project/modules/repository/user"
	services "github.com/Billy278/assignment_project/modules/services/gmail"
	servPromo "github.com/Billy278/assignment_project/modules/services/promo"
	servUser "github.com/Billy278/assignment_project/modules/services/user"

	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	GmailSrv controllers.GmailCtr
	PromoSrv ctrlPromo.PromoCtrl
	UserSrv  ctrlUser.UserCtrl
}

func InitServer() Handlers {
	v := validator.New()
	datastore := db.NewDBPostges()
	// repo
	repoGmail := repository.NewGmailRepoImpl(datastore)
	repopromo := repoPromo.NewPromoRepoImpl(datastore)
	repouser := repoUser.NewUserRepoImpl(datastore)
	// services
	srvGmail := services.NewGmailSrvImpl(repoGmail)
	srvpromo := servPromo.NewPromoSrvImpl(repopromo)
	srvuser := servUser.NewUserSrvImpl(repouser)
	// controller
	ctrlGmail := controllers.NewGmailCtrImpl(v, srvGmail)
	ctrlpromo := ctrlPromo.NewPromoCtrlImpl(srvpromo)
	ctrluser := ctrlUser.NewUserCtrlImpl(srvuser)

	return Handlers{
		GmailSrv: ctrlGmail,
		PromoSrv: ctrlpromo,
		UserSrv:  ctrluser,
	}
}
