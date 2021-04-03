
# micro service  backend for front 


golang版本要求 1.15+

```
├── user_web  用户bff层
    ├── api   具体业务处理
    ├── config 配置
    ├── forms  表单
    ├── global 全局变量
    ├── initalize 初始化相关
    ├──  middlewares  请求中间件
    ├──  models  数据模型 非数据库的...
    ├──  proto  存放proto文件，pb文件
    ├──  router 路由
    ├──  utils  工具包
    ├──  validator 验证相关
    ├──  main.go 入口
    ├──  config-debug.yaml 配置文件 
├── scripts  存放脚本

```


**技术选型**

- 框架：gin

- 配置中心: nacos

- 缓存 :redis

- 服务发现: [grpc-consul-resolver](https://github.com/mbobakov/grpc-consul-resolver)

- 日志记录： zap

- 配置：viper 

- 用户认证：jwt 



quickstart 

```shell 
#启动consul
#下载mac版二进制consul,默认服务端口 127.0.0.1:8500

$ consul agent -dev

#启动nacos 
$ docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:latest


# 生成proto代码
$ protoc --proto_path=user_web/proto --go_out=plugins=grpc:user_web/proto user.proto
$  protoc --proto_path=goods_web/proto --go_out=plugins=grpc:goods_web/proto goods.proto
```



