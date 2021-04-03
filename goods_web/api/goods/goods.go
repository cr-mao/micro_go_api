package goods

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"

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

	rsp, err := global.GoodsSrvClient.GoodsList(context.Background(), request)
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
