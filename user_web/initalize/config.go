package initalize

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"bff/user_web/global"
)

func GetEnv(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfig() {
	v := viper.New()
	// 必须是大写， ENV
	debug := GetEnv("ENV")
	//fmt.Println(debug)
	configSuffix := "debug"
	if debug != "DEV" {
		configSuffix = "pro"
	}
	// debug 模式 路径有点问题   ，可以临时替换成绝对路径  debug 值 也是有问题
	v.SetConfigFile(fmt.Sprintf("config-%s.yaml", configSuffix))
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	err := v.Unmarshal(&global.ServerConfig)
	zap.S().Infof("初始化配置信息 %#v", global.ServerConfig)
	fmt.Println(global.ServerConfig)
	if err != nil {
		zap.S().Errorw("初始化配置失败", "msg", err.Error())
	}

	// 暂时有问题
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//	//_=v.ReadInConfig()
	//	//_=v.Unmarshal(&config)
	//	//fmt.Println(config)
	//})
	//v.WatchConfig()
}
