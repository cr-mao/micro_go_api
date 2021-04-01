package api

import (
	"bff/user_web/middlewares"
	"bff/user_web/models"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bff/user_web/forms"
	"bff/user_web/global"
	"bff/user_web/global/response"
	"bff/user_web/proto"
)

//移除 struct 名字  默认struct.field_name   只要 field_name
func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误" + e.Message(),
				})
			}
			return
		}
	}
}

func HandleValidateError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}

func GetUserList(c *gin.Context) {

	//fmt.Println(c.Get("userId"))
	//拿上下文信息
	//if claims, ok := c.Get("claims"); ok {
	//	if c, ok := claims.(*models.CustomClaims); ok {
	//		fmt.Println(c.ID)
	//	}
	//}

	ip := global.ServerConfig.UserSrvInfo.Host
	port := global.ServerConfig.UserSrvInfo.Port
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接服务失败", "msg", err.Error())
	}
	defer conn.Close()
	client := proto.NewUserClient(conn)

	pn := c.DefaultQuery("pn", "1")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := client.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})

	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	result := make([]interface{}, 0)

	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: response.Jsontime(time.Unix(int64(value.BirthDay), 0)),
			Mobile:   value.Mobile,
		}
		result = append(result, user)
		//data := make(map[string]interface{})
		//data["id"] = value.Id
		//data["name"] = value.NickName
		//data["birthday"] = value.BirthDay
		//data["mobile"] = value.Mobile
		//result = append(result, data)
	}
	c.JSON(http.StatusOK, result)
}

func PassWordLogin(c *gin.Context) {
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidateError(c, err)
		return
	}
	//验证码 验证
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	ip := global.ServerConfig.UserSrvInfo.Host
	port := global.ServerConfig.UserSrvInfo.Port
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接服务失败", "msg", err.Error())
	}
	defer conn.Close()
	client := proto.NewUserClient(conn)
	//登录逻辑

	if resp, err := client.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				//fmt.Println(e.Code())
				//fmt.Println(e.Err())
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		//检测密码
		if rsp, err := client.CheckPassword(context.Background(), &proto.PasswrodCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: resp.PassWord,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"mobile": "登录失败",
			})
		} else {
			if rsp.Success {
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(resp.Id),
					NickName:    resp.NickName,
					AuthorityId: 1,
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),                                             //签名生效时间
						ExpiresAt: time.Now().Unix() + int64(global.ServerConfig.JWTInfo.Expire), //失效时间
						Issuer:    global.ServerConfig.JWTInfo.Issuer,
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"msg":        "登录成功",
					"token":      token,
					"id":         resp.Id,
					"nick_name":  resp.NickName,
					"expired_at": (time.Now().Unix() + int64(global.ServerConfig.JWTInfo.Expire)) * 1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}
		}
	}
}
