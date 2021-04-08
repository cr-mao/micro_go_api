package order

import (
	"bff/order_web/api"
	"bff/order_web/forms"
	"bff/order_web/global"
	"bff/order_web/models"
	"bff/order_web/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	userId, _ := c.Get("userId")
	claims, _ := c.Get("claims")
	request := proto.OrderFilterRequest{}

	//如果是管理员用户则返回所有的订单
	model := claims.(*models.CustomClaims)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}
	pages := c.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := c.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)
	rsp, err := global.OrderSrvClient.OrderList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单列表失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	/**
	{
			"total":100,
			"data":[{}]
	}
	*/

	reMap := gin.H{
		"total": rsp.Total,
	}

	orderList := make([]interface{}, 0)

	for _, item := range rsp.Data {
		tmpMap := map[string]interface{}{}
		tmpMap["id"] = item.Id
		tmpMap["status"] = item.Status
		tmpMap["pay_type"] = item.PayType
		tmpMap["user"] = item.UserId
		tmpMap["post"] = item.Post
		tmpMap["total"] = item.Total
		tmpMap["address"] = item.Address
		tmpMap["name"] = item.Name
		tmpMap["mobile"] = item.Mobile
		tmpMap["order_sn"] = item.OrderSn
		tmpMap["add_time"] = item.AddTime
		orderList = append(orderList, tmpMap)
	}
	reMap["data"] = orderList
	c.JSON(http.StatusOK, reMap)

}

func New(c *gin.Context) {
	orderForm := forms.CreateOrderForm{}
	if err := c.ShouldBindJSON(&orderForm); err != nil {
		fmt.Println(err)
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	userId, _ := c.Get("userId")
	rsp, err := global.OrderSrvClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Address: orderForm.Address,
		Mobile:  orderForm.Mobile,
		Name:    orderForm.Name,
		Post:    orderForm.Post,
	})
	if err != nil {
		zap.S().Errorw("创建订单失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	//TODO 支付宝url
	c.JSON(http.StatusOK, gin.H{
		"id": rsp.Id,
	})
	//CreateOrderForm

}

func Detail(c *gin.Context) {

}
