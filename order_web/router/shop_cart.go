package router

import (
	"github.com/gin-gonic/gin"

	"bff/order_web/api/shop_cart"
	"bff/order_web/middlewares"
)

func InitShopCartRouter(group *gin.RouterGroup) {
	shopCartGroup := group.Group("shopcarts").Use(middlewares.JWTAuth())
	{
		shopCartGroup.GET("", shop_cart.List)          //购物车列
		shopCartGroup.DELETE("/:id", shop_cart.Delete) //购物车删除条目
		shopCartGroup.POST("", shop_cart.New)          //添加条目
		shopCartGroup.PATCH("/:id", shop_cart.Update)  //修改条目
	}
}
