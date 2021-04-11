package goods

import (
	"context"
	"fmt"
	"net/http"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"bff/goods_web/api"
	"bff/goods_web/global"
	"bff/goods_web/proto"
)

func List(c *gin.Context) {
	request := &proto.GoodsFilterRequest{}
	request.TopCategroy = 1
	isHot := c.DefaultQuery("ih", "0")
	if isHot == "1" {
		request.IsHot = true
	}

	//这块代码 到e.Exit() 进行限速了
	e, b := sentinel.Entry("goods-list", sentinel.WithTrafficType(base.Inbound))
	if b != nil {
		// 限流了
		fmt.Println("not pass")
		c.JSON(http.StatusTooManyRequests, gin.H{
			"msg": "请稍后重试",
		})
		return
	}
	// TODO 再研究一下 ( 客户端请求 进入拦截器(  otgrpc.OpenTracingClientInterceptor ) ,上下文信息中有jaeger ，和parent span)
	ct := context.WithValue(context.Background(), "ginContext", c)
	rsp, err := global.GoodsSrvClient.GoodsList(ct, request)
	e.Exit()
	if err != nil {
		zap.S().Errorw("查询商品失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	result := map[string]interface{}{
		"total": rsp.Total,
		"goods": rsp.Goods,
	}
	c.JSON(http.StatusOK, result)

}
