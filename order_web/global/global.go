package global

import (
	"bff/order_web/config"
	"bff/order_web/proto"
	ut "github.com/go-playground/universal-translator"
)

var Trans ut.Translator
var ServerConfig = &config.ServerConfig{}

//全局usersrv 连接对象
var GoodsSrvClient proto.GoodsClient
var OrderSrvClient proto.OrderClient
var InventorySrvClient proto.InventoryClient

