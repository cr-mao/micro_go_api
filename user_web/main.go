package main

import (
	"bff/user_web/global"
	"fmt"

	"go.uber.org/zap"

	"bff/user_web/initalize"
)

func main() {

	// 初始化全局logger ， zap.S()可以获取全局的logger
	initalize.InitLogger()
	initalize.InitConfig()
	Router := initalize.Routers()
	port := global.ServerConfig.Port
	zap.S().Infof("启动服务器端口: %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

}
