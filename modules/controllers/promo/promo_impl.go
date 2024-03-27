package controllers

import (
	"net/http"

	services "github.com/Billy278/assignment_project/modules/services/promo"
	"github.com/Billy278/assignment_project/pkg/responses"
	"github.com/gin-gonic/gin"
)

type PromoCtrlImpl struct {
	PromoSrv services.PromoSrv
}

func NewPromoCtrlImpl(promosrv services.PromoSrv) PromoCtrl {
	return &PromoCtrlImpl{
		PromoSrv: promosrv,
	}
}

func (ctrl *PromoCtrlImpl) GetToken(ctx *gin.Context) {
	promo, ok := ctx.GetQuery("kode_promo")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusNotFound,
			Success: false,
			Message: responses.NotFound,
		})
		return
	}

	resPromo, err := ctrl.PromoSrv.GetToken(ctx, promo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusNotFound,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Data:    resPromo,
	})
}
func (ctrl *PromoCtrlImpl) Created(ctx *gin.Context) {
	err := ctrl.PromoSrv.CreatedPromo(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "Success created and Send Kode Promo",
	})

}
