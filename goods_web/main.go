package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"

	"bff/goods_web/global"
	"bff/goods_web/initalize"
	"bff/goods_web/utils/register/consul"
)

func main() {

	// 初始化全局logger ， zap.S()可以获取全局的logger
	initalize.InitLogger()
	//初始化配置
	initalize.InitConfig()
	//初始化验证翻译
	if err := initalize.InitTrans("zh"); err != nil {
		panic(err.Error())
	}
	//注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//翻译问题
		// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go#L105
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
	//初始化连接 user-srv  grpc connect
	initalize.InitSrvConn()
	//初始化路由
	Router := initalize.InitRouters()
	port := global.ServerConfig.Port

	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	//  对外服务的ip
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败", err.Error())
	}
	zap.S().Infof("启动服务器端口: %d", port)

	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err = registerClient.DeRegister(serviceId); err != nil {
		zap.S().Panic("注销失败", err.Error())
	}
}
