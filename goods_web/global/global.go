package global

import (
	"bff/goods_web/config"
	"bff/goods_web/proto"
	ut "github.com/go-playground/universal-translator"
)

var Trans ut.Translator
var ServerConfig = &config.ServerConfig{}

//全局usersrv 连接对象
var GoodsSrvClient proto.GoodsClient
