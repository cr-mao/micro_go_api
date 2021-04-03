package initalize

import "go.uber.org/zap"

func InitLogger(){
	/*
	 1. S()可以获取一个全局的sugar,可以设置一个全局的logger
	 2. 日志是分级别的 debug，info, warn ,error, fetal ,         debug 开发环境才有效
	*/
	logger, _ :=zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}