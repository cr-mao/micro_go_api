package router

import (
	"github.com/gin-gonic/gin"

	"bff/goods_web/api/goods"
	"bff/goods_web/middlewares"
)

func InitGoodsRouter(group *gin.RouterGroup) {
	GoodsGroup := group.Group("goods").Use(middlewares.Trace())
	{
		GoodsGroup.GET("list", goods.List)
	}
}
