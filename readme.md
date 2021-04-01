



# micro service  backend for front 


框架：gin
日志记录： zap
配置文件：viper 
用户认证：jwt 




#### 便捷
```shell script
protoc --proto_path=user_web/proto --go_out=plugins=grpc:user_web/proto user.proto
```
