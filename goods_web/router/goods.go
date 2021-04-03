package router

import (
	"bff/goods_web/api/goods"
	"github.com/gin-gonic/gin"
)

func InitGoodsRouter(group *gin.RouterGroup) {
	GoodsGroup := group.Group("goods")
	{
		GoodsGroup.GET("list", goods.List)
	}
}
