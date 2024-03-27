package controllers

import "github.com/gin-gonic/gin"

type UserCtrl interface {
	GetAllUsersIsBrithday(ctx *gin.Context)
}
