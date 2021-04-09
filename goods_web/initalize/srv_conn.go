package initalize

import (
	"fmt"
	"github.com/opentracing/opentracing-go"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"bff/goods_web/global"
	"bff/goods_web/proto"
	"bff/goods_web/utils/otgrpc"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.ServerName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatal("consul 获得 goods-srv ip:port 失败")
		return
	}
	global.GoodsSrvClient = proto.NewGoodsClient(conn)
}
