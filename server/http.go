package server

import (
	routeGmail "github.com/Billy278/assignment_project/modules/route/gmail"
	routePromo "github.com/Billy278/assignment_project/modules/route/promo"
	routeUser "github.com/Billy278/assignment_project/modules/route/user"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	g := gin.Default()
	g.Use(gin.Recovery())
	handler := InitServer()
	v := g.Group("api/")
	routeGmail.NewGmailRoute(v, handler.GmailSrv)
	routeUser.NewUserRoute(v, handler.UserSrv)
	routePromo.NewPromoRoute(v, handler.PromoSrv)
	g.Run(":9090")
}
