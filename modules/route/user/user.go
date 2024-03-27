package user

import (
	controllers "github.com/Billy278/assignment_project/modules/controllers/user"
	"github.com/gin-gonic/gin"
)

func NewUserRoute(v1 *gin.RouterGroup, userCtrl controllers.UserCtrl) {
	g := v1.Group("/user")
	g.GET("", userCtrl.GetAllUsersIsBrithday)

}
