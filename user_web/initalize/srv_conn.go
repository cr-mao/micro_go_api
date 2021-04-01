package initalize

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"bff/user_web/global"
	"bff/user_web/proto"
)

//初始化 grpcClient连接
func InitSrvConn() {
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
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
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接服务失败", "msg", err.Error())
	}
	//defer conn.Close()
	client := proto.NewUserClient(conn)
	global.UserSrvClient = client
}
