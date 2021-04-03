package api

import (
	"bff/user_web/global"
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

//获得图片验证码
func GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}

// 短信验证码
func GenerateSmsCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	length := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(length)])
	}
	return sb.String()
}

func GetSmsCode(ctx *gin.Context) {
	// 号码
	smsCode := GenerateSmsCode(4)
	tel := "18711111111"
	//存redis 中
	redisCli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	redisCli.Set(context.Background(), tel, smsCode, 60*time.Second)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送验证成功",
	})
}
