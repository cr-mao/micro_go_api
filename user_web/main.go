package main

import (
	"go.uber.org/zap"

	"bff/user_web/initalize"
	"fmt"
)

func main() {
	port := 8021
	// 初始化全局logger ， zap.S()可以获取全局的logger
	initalize.InitLogger()
	Router := initalize.Routers()
	zap.S().Infof("启动服务器端口: %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

}
