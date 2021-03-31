package router

import (
	"github.com/gin-gonic/gin"

	"bff/user_web/api"
)

func InitUserRouter(group *gin.RouterGroup) {

	userGroup := group.Group("/user")
	{
		userGroup.GET("/list", api.GetUserList)
	}

}
