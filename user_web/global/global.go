package global

import (
	"bff/user_web/config"
	ut "github.com/go-playground/universal-translator"
)

var Trans ut.Translator
var ServerConfig = &config.ServerConfig{}
