package product

import (
	controllers "github.com/Billy278/assignment_project/modules/controllers/product"
	"github.com/gin-gonic/gin"
)

func NewProductRoute(v1 *gin.RouterGroup, productCtrl controllers.CtrlProduct) {
	g := v1.Group("/product")
	g.POST("", productCtrl.Created)
	g.GET(":id", productCtrl.FindByid)
	g.GET("", productCtrl.List)
	g.DELETE(":id", productCtrl.Deleted)

}
