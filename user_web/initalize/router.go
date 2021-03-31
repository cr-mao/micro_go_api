package initalize

import (
	router2 "bff/user_web/router"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	ApiGroup := router.Group("/v1")
	router2.InitUserRouter(ApiGroup)
	return router
}