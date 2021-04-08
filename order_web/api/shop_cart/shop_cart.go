package shop_cart

import (
	"bff/order_web/api"
	"bff/order_web/global"
	"bff/order_web/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func List(c *gin.Context) {
	//获取购物车商品

	userId, _ := c.Get("userId")
	rsp, err := global.OrderSrvClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("[list]查询购物车列表失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	ids := make([]int32, 0)
	for _, item := range rsp.Data {
		ids = append(ids, item.GoodsId)
	}
	if len(ids) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}

	fmt.Println(ids)

	// 请求商品服务获取商品信息
	goodsRsp, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: ids,
	})

	if err != nil {
		zap.S().Errorw("批量查询商品失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	reMap := gin.H{
		"total": rsp.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		for _, good := range goodsRsp.Goods {
			if good.Id == item.GoodsId {
				tmpMap := map[string]interface{}{}
				tmpMap["id"] = item.Id
				tmpMap["goods_id"] = item.GoodsId
				tmpMap["good_name"] = good.Name
				tmpMap["good_image"] = good.GoodsFrontImage
				tmpMap["good_price"] = good.ShopPrice
				tmpMap["nums"] = item.Nums
				tmpMap["checked"] = item.Checked
				goodsList = append(goodsList, tmpMap)
			}
		}
	}
	reMap["data"] = goodsList
	c.JSON(http.StatusOK, reMap)

}

func New(c *gin.Context) {
  //检测库存够不够
	//  调增加购物车srv接口
}

func Delete(c *gin.Context) {

}

func Update(c *gin.Context) {

}
