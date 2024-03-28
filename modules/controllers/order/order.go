package controllers

import "github.com/gin-gonic/gin"

type CtrlOrders interface {
	List(ctx *gin.Context)
	Order(ctx *gin.Context)
}
