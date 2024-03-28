package controllers

import (
	"net/http"
	"time"

	models "github.com/Billy278/assignment_project/modules/models/user"
	services "github.com/Billy278/assignment_project/modules/services/user"
	"github.com/Billy278/assignment_project/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCtrlImpl struct {
	UserSrv  services.UserSrv
	Validate *validator.Validate
}

func NewUserCtrlImpl(usersrv services.UserSrv, v *validator.Validate) UserCtrl {
	return &UserCtrlImpl{
		UserSrv:  usersrv,
		Validate: v,
	}
}
func (ctrl *UserCtrlImpl) Created(ctx *gin.Context) {
	//validasi req
	reqUser := models.User{}
	err := ctx.ShouldBindJSON(&reqUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}
	//validasi req with validator
	err = ctrl.Validate.Struct(reqUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	// validasi requsest dobstring
	layout := "2006-01-02 15:04:05.999999"
	tempDoB, err := time.Parse(layout, reqUser.DoBString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	reqUser.DoB = &tempDoB

	//cek apakah username sudah ada yg menggunakan
	err = ctrl.UserSrv.SrvFindUser(ctx, reqUser.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "Username sudah digunakan",
		})
		return
	}

	err = ctrl.UserSrv.CreatedUser(ctx, reqUser)
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
	})
}
func (ctrl *UserCtrlImpl) GetAllUsersIsBrithday(ctx *gin.Context) {
	resUser, err := ctrl.UserSrv.GetAllUserIsbrithday(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Data:    resUser,
	})
	return
}
