package server

import (
	routeGmail "github.com/Billy278/assignment_project/modules/route/gmail"
	routeOrder "github.com/Billy278/assignment_project/modules/route/order"
	routeProduct "github.com/Billy278/assignment_project/modules/route/product"
	routePromo "github.com/Billy278/assignment_project/modules/route/promo"
	routeUser "github.com/Billy278/assignment_project/modules/route/user"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	g := gin.Default()
	g.Use(gin.Recovery())
	handler := InitServer()
	v := g.Group("api/")
	routeGmail.NewGmailRoute(v, handler.GmailCtrl)
	routeUser.NewUserRoute(v, handler.UserCtrl)
	routePromo.NewPromoRoute(v, handler.PromoCtrl)
	routeProduct.NewProductRoute(v, handler.ProductCtl)
	routeOrder.NewOrderRoute(v, handler.OrderCtl)
	g.Run(":9090")
}
