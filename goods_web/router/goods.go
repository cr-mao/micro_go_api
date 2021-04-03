package router

import (
	"bff/goods_web/middlewares"
	"github.com/gin-gonic/gin"

	"bff/goods_web/api"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("user")
	{
		userGroup.GET("list", middlewares.JWTAuth(), api.GetUserList)
		userGroup.POST("login", api.PassWordLogin)
	}
}
