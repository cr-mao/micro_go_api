
# micro service  backend for front 


框架：gin

数据库：mysql

缓存 :redis

服务注册: consul 

日志记录： zap

配置：viper 

用户认证：jwt 

图形验证码: base64Captcha


#### 便捷
```shell script
protoc --proto_path=user_web/proto --go_out=plugins=grpc:user_web/proto user.proto
```
