package router

import (
	"bff/user_web/middlewares"
	"github.com/gin-gonic/gin"

	"bff/user_web/api"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("user")
	{
		userGroup.GET("list", middlewares.JWTAuth(), api.GetUserList)
		userGroup.POST("login", api.PassWordLogin)
	}
}
