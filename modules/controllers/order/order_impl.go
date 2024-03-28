package controllers

import (
	"net/http"
	"strconv"

	models "github.com/Billy278/assignment_project/modules/models/orders"
	modelsToken "github.com/Billy278/assignment_project/modules/models/token"
	services "github.com/Billy278/assignment_project/modules/services/order"
	"github.com/Billy278/assignment_project/pkg/crypto"
	"github.com/Billy278/assignment_project/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlOrdersImpl struct {
	OrderSrv services.OrderSrv
	Validate *validator.Validate
}

func NewCtrlOrdersImpl(ordersrv services.OrderSrv, v *validator.Validate) CtrlOrders {
	return &CtrlOrdersImpl{
		OrderSrv: ordersrv,
		Validate: v,
	}
}
func (ctrl *CtrlOrdersImpl) List(ctx *gin.Context) {
	resOrder, err := ctrl.OrderSrv.SrvList(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, responses.Response{
		Code:    http.StatusAccepted,
		Success: true,
		Message: responses.Success,
		Data:    resOrder,
	})
}
func (ctrl *CtrlOrdersImpl) Order(ctx *gin.Context) {
	// bind req
	reqOrder := models.Order{}
	err := ctx.ShouldBindJSON(&reqOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}
	//validate reqOrder with validator
	err = ctrl.Validate.Struct(&reqOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	var isgetPromo bool
	if reqOrder.Promo_code != "" {
		// get Token with string promo
		token, err := ctrl.OrderSrv.SrvGetTokenWithPromo(ctx, reqOrder.Promo_code)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
				Code:    http.StatusBadRequest,
				Success: false,
				Message: err.Error(),
			})
			return
		}
		// validasi token
		var claim modelsToken.AccessClaimPromo
		err = crypto.ParseAndVerifyToken(token, &claim)
		if err != nil {
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
					Code:    http.StatusBadRequest,
					Success: false,
					Message: "invalid Token",
				})
				return
			}
		}
		id := strconv.Itoa(int(reqOrder.UserId))
		// cek apakah user yg menclaim token sama idnya dengan id user yg dapat promo

		if claim.UserId != id {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				Code:    http.StatusUnauthorized,
				Success: false,
				Message: responses.Unauthorized,
			})
			return

		}
		isgetPromo = true

	}

	resOrder, err := ctrl.OrderSrv.SrvCreate(ctx, reqOrder, isgetPromo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, responses.Response{
		Code:    http.StatusCreated,
		Success: true,
		Message: responses.Success,
		Data:    resOrder,
	})

}
