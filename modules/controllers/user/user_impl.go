package controllers

import (
	"net/http"

	services "github.com/Billy278/assignment_project/modules/services/user"
	"github.com/Billy278/assignment_project/pkg/responses"
	"github.com/gin-gonic/gin"
)

type UserCtrlImpl struct {
	UserSrv services.UserSrv
}

func NewUserCtrlImpl(usersrv services.UserSrv) UserCtrl {
	return &UserCtrlImpl{
		UserSrv: usersrv,
	}
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
