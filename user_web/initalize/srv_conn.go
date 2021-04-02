package initalize

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"bff/user_web/global"
	"bff/user_web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.ServerName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("consul 获得 user-srv ip:port 失败")
		return
	}
	client := proto.NewUserClient(conn)
	global.UserSrvClient = client
}

//初始化 grpcClient连接
func InitSrvConnBak() {
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//从consul 中 获得 user-srv ip和port
	data, err := consulClient.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.ServerName))
	//读配置文件的作废
	//ip := global.ServerConfig.UserSrvInfo.Host
	//port := global.ServerConfig.UserSrvInfo.Port
	var userSrvHost string
	var userSrvPort int
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("consul 获得 user-srv ip:port 失败")
		return
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Fatal("consul 获得 user-srv ip:port 失败")
		return
	}
	// 1.后续用户服务下线了 ，2改端口了，改ip了 负责均衡策略
	// 2.事先创建连接不用tcp3次握手了
	// 3. 一个连接多个grountine 共用 ，性能问题， todo  连接池 https://github.com/processout/grpc-go-pool/blob/master/pool.go
	//defer conn.Close()
	client := proto.NewUserClient(conn)
	global.UserSrvClient = client
}
