package api

import "C"
import (
	"bff/user_web/global/response"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bff/user_web/proto"
)

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

func GetUserList(c *gin.Context) {
	ip := "127.0.0.1"
	port := "50052"
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接服务失败", "msg", err.Error())
	}
	defer conn.Close()
	client := proto.NewUserClient(conn)
	rsp, err := client.GetUserList(context.Background(), &proto.PageInfo{Pn: 1, PSize: 5})
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
