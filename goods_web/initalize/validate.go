package initalize

import (
	"bff/goods_web/global"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var uni *ut.UniversalTranslator

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//采用tag  json的标签名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zh := zh2.New()
		en := en2.New()
		//第一个参数en备用， zh应该支持
		uni = ut.New(en, zh, en)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			return en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			return zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			return en_translations.RegisterDefaultTranslations(v, global.Trans)
		}
	}
	return
}
