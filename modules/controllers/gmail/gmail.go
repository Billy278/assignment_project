package controllers

import "github.com/gin-gonic/gin"

type GmailCtr interface {
	Created(ctx *gin.Context)
}
