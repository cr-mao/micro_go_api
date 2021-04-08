package initalize

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"bff/order_web/config"
	"bff/order_web/global"
	"bff/order_web/utils"
)

func GetEnv(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfigbak() {
	v := viper.New()
	// 必须是大写， ENV
	debug := GetEnv("ENV")
	//fmt.Println(debug)
	configSuffix := "debug"
	//生成环境 使用随机端口 ，开发环境使用 固定端口
	if debug != "DEV" {
		configSuffix = "pro"
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
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

func InitConfig() {
	v := viper.New()
	// 必须是大写， ENV
	debug := GetEnv("ENV")
	//fmt.Println(debug)
	//configSuffix := "debug"
	//生成环境 使用随机端口 ，开发环境使用 固定端口
	// debug 模式 路径有点问题   ，可以临时替换成绝对路径  debug 值 也是有问题
	//v.SetConfigFile(fmt.Sprintf("config-%s.yaml", configSuffix))
	//读取nacos配置a
	v.SetConfigFile("nacos.yaml")
	if err := v.ReadInConfig(); err != nil {
		zap.S().Fatalf("读取配置失败", "msg", err.Error())
	}
	nacosConfig := config.NacosConfig{}
	err := v.Unmarshal(&nacosConfig)
	if err != nil {
		zap.S().Fatalf("初始化配置失败1", "msg", err.Error())
	}
	zap.S().Infof("nacosConfig===> %#v", global.ServerConfig)
	content, err := GetConfigFromNacos(nacosConfig.Host, nacosConfig.Port, nacosConfig.NameSpace, nacosConfig.DataId, nacosConfig.Group)
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("初始化配置失败2", "msg", err.Error())
	}
	zap.S().Infof("serverConfig===> %#v", global.ServerConfig)
	if debug != "DEV" {
		//configSuffix = "pro"
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}
}

// 从nacos 读取配置
func GetConfigFromNacos(ip string, port int, namespace, dataId, group string) (string, error) {
	sc := []constant.ServerConfig{
		{
			IpAddr: ip,
			Port:   uint64(port),
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         namespace, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return "", err
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	if err != nil {
		return content, err
	}
	err = client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		return "", err
	}
	return content, nil
}
