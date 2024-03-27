package promo

import (
	controllers "github.com/Billy278/assignment_project/modules/controllers/promo"
	"github.com/gin-gonic/gin"
)

func NewPromoRoute(v1 *gin.RouterGroup, promoCtrl controllers.PromoCtrl) {
	g := v1.Group("/promo")
	g.POST("", promoCtrl.Created)
	g.GET("", promoCtrl.GetToken)

}
