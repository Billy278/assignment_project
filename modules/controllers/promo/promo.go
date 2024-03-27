package controllers

import "github.com/gin-gonic/gin"

type PromoCtrl interface {
	Created(ctx *gin.Context)
	GetToken(ctx *gin.Context)
}
