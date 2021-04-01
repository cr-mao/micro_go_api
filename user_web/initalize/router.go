package initalize

import (
	"bff/user_web/middlewares"
	router2 "bff/user_web/router"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Cors())
	ApiGroup := router.Group("/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)
	return router
}
