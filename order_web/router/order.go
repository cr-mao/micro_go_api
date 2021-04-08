package router

import (
	"bff/order_web/api/order"
	"bff/order_web/middlewares"
	"github.com/gin-gonic/gin"
)

func InitOrderRouter(group *gin.RouterGroup) {
	OrderGroup := group.Group("orders").Use(middlewares.JWTAuth())
	{
		OrderGroup.GET("", order.List)
		OrderGroup.POST("", order.New)         //新建订单
		OrderGroup.POST("/:id/", order.Detail) //订单详情
	}
}
