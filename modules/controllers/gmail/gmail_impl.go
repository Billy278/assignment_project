package controllers

import (
	"net/http"

	models "github.com/Billy278/assignment_project/modules/models/gmail"
	services "github.com/Billy278/assignment_project/modules/services/gmail"
	"github.com/Billy278/assignment_project/pkg/responses"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type GmailCtrImpl struct {
	Validate *validator.Validate
	GmailSrv services.GmailSrv
}

func NewGmailCtrImpl(v *validator.Validate, gmailsrv services.GmailSrv) GmailCtr {
	return &GmailCtrImpl{
		Validate: v,
		GmailSrv: gmailsrv,
	}

}
func (ctrl *GmailCtrImpl) Created(ctx *gin.Context) {
	// validasi req much json
	gmailIn := models.Gmail{}
	err := ctx.ShouldBindJSON(gmailIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	// validasi req
	err = ctrl.Validate.Struct(gmailIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = ctrl.GmailSrv.Created(ctx, gmailIn)
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
		Message: "Success Send Gmail",
	})
}
