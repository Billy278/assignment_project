package gmail

import (
	controllers "github.com/Billy278/assignment_project/modules/controllers/gmail"
	"github.com/gin-gonic/gin"
)

func NewGmailRoute(v1 *gin.RouterGroup, GmailCtrl controllers.GmailCtr) {
	g := v1.Group("/gmail")
	g.POST("", GmailCtrl.Created)

}
