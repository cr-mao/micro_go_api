package global

import (
	"bff/user_web/config"
	"bff/user_web/proto"
	ut "github.com/go-playground/universal-translator"
)

var Trans ut.Translator
var ServerConfig = &config.ServerConfig{}

//全局usersrv 连接对象
var UserSrvClient proto.UserClient
