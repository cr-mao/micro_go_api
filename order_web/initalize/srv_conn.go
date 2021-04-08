package initalize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"bff/order_web/global"
	"bff/order_web/proto"
)

func InitSrvConn() {

	consulInfo := global.ServerConfig.ConsulInfo
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.ServerName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("consul 获得 goods-srv ip:port 失败")
		return
	}
	global.GoodsSrvClient = proto.NewGoodsClient(conn)

	conn, err = grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.OrderSrvInfo.ServerName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("consul 获得 order-srv ip:port 失败")
		return
	}
	global.OrderSrvClient = proto.NewOrderClient(conn)

	conn, err = grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.InventorySrvInfo.ServerName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("consul 获得 inventory-srv ip:port 失败")
		return
	}
	global.InventorySrvClient = proto.NewInventoryClient(conn)

}
