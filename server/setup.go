package server

import (
	"github.com/Billy278/assignment_project/db"
	controllers "github.com/Billy278/assignment_project/modules/controllers/gmail"
	ctrlOrder "github.com/Billy278/assignment_project/modules/controllers/order"
	ctrlProduct "github.com/Billy278/assignment_project/modules/controllers/product"
	ctrlPromo "github.com/Billy278/assignment_project/modules/controllers/promo"
	ctrlUser "github.com/Billy278/assignment_project/modules/controllers/user"
	repository "github.com/Billy278/assignment_project/modules/repository/gmail"
	repoOrder "github.com/Billy278/assignment_project/modules/repository/order"
	repoProduct "github.com/Billy278/assignment_project/modules/repository/product"
	repoPromo "github.com/Billy278/assignment_project/modules/repository/promo"
	repoUser "github.com/Billy278/assignment_project/modules/repository/user"
	services "github.com/Billy278/assignment_project/modules/services/gmail"
	servOrder "github.com/Billy278/assignment_project/modules/services/order"
	servProduct "github.com/Billy278/assignment_project/modules/services/product"
	servPromo "github.com/Billy278/assignment_project/modules/services/promo"
	servUser "github.com/Billy278/assignment_project/modules/services/user"

	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	GmailCtrl  controllers.GmailCtr
	PromoCtrl  ctrlPromo.PromoCtrl
	UserCtrl   ctrlUser.UserCtrl
	ProductCtl ctrlProduct.CtrlProduct
	OrderCtl   ctrlOrder.CtrlOrders
}

func InitServer() Handlers {
	v := validator.New()
	datastore := db.NewDBPostges()
	// repo
	repoGmail := repository.NewGmailRepoImpl(datastore)
	repopromo := repoPromo.NewPromoRepoImpl(datastore)
	repouser := repoUser.NewUserRepoImpl(datastore)
	repoproduct := repoProduct.NewRepoProductImpl(datastore)
	repoorder := repoOrder.NewOrderRepoImpl()

	// services
	srvGmail := services.NewGmailSrvImpl(repoGmail)
	srvpromo := servPromo.NewPromoSrvImpl(repopromo)
	srvuser := servUser.NewUserSrvImpl(repouser)
	srvproduct := servProduct.NewSrvProductImpl(repoproduct)
	srvorder := servOrder.NewOrderSrvImpl(datastore, repoorder, repoproduct)
	// controller
	ctrlGmail := controllers.NewGmailCtrImpl(v, srvGmail)
	ctrlpromo := ctrlPromo.NewPromoCtrlImpl(srvpromo)
	ctrluser := ctrlUser.NewUserCtrlImpl(srvuser, v)
	ctrlproduct := ctrlProduct.NewCtrlProductimpl(srvproduct, v)
	ctrlorder := ctrlOrder.NewCtrlOrdersImpl(srvorder, v)

	return Handlers{
		GmailCtrl:  ctrlGmail,
		PromoCtrl:  ctrlpromo,
		UserCtrl:   ctrluser,
		ProductCtl: ctrlproduct,
		OrderCtl:   ctrlorder,
	}
}
