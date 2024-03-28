package controllers

import (
	"net/http"
	"strconv"

	models "github.com/Billy278/assignment_project/modules/models/products"
	services "github.com/Billy278/assignment_project/modules/services/product"
	"github.com/Billy278/assignment_project/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlProductimpl struct {
	ProductSrv services.SrvProduct
	Validate   *validator.Validate
}

func NewCtrlProductimpl(productsrv services.SrvProduct, v *validator.Validate) CtrlProduct {
	return &CtrlProductimpl{
		ProductSrv: productsrv,
		Validate:   v,
	}
}

func (ctrl *CtrlProductimpl) List(ctx *gin.Context) {
	resProduct, err := ctrl.ProductSrv.SrvList(ctx)
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
		Data:    resProduct,
	})
}
func (ctrl *CtrlProductimpl) FindByid(ctx *gin.Context) {
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	resProduct, err := ctrl.ProductSrv.SrvFindByid(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusNotFound,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, responses.Response{
		Code:    http.StatusAccepted,
		Success: true,
		Message: responses.Success,
		Data:    resProduct,
	})
}
func (ctrl *CtrlProductimpl) Created(ctx *gin.Context) {
	reqProduct := models.Product{}
	err := ctx.ShouldBindJSON(&reqProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	//validasi req
	err = ctrl.Validate.Struct(reqProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}

	resProduct, err := ctrl.ProductSrv.SrvCreate(ctx, reqProduct)
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
		Data:    resProduct,
	})
}

func (ctrl *CtrlProductimpl) Deleted(ctx *gin.Context) {
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	err = ctrl.ProductSrv.SrvDelete(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusNotFound,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: responses.Success,
	})
}
func (ctrl *CtrlProductimpl) ConvertId(ctx *gin.Context) (id uint64, err error) {
	reqId := ctx.Param("id")
	id, err = strconv.ParseUint(reqId, 10, 64)
	if err != nil {
		return
	}
	return
}
