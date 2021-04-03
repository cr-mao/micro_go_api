package initalize

import (
	"bff/goods_web/middlewares"
	router2 "bff/goods_web/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"success": true,
		})
	})
	router.Use(middlewares.Cors())
	ApiGroup := router.Group("/g/v1")
	router2.InitGoodsRouter(ApiGroup)
	return router
}
