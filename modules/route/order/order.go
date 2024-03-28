package order

import (
	controllers "github.com/Billy278/assignment_project/modules/controllers/order"
	"github.com/gin-gonic/gin"
)

func NewOrderRoute(v1 *gin.RouterGroup, orderCtrl controllers.CtrlOrders) {
	g := v1.Group("/order")
	g.POST("", orderCtrl.Order)
	g.GET("", orderCtrl.List)

}
