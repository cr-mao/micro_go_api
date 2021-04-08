package initalize

import (
	"bff/order_web/middlewares"
	router2 "bff/order_web/router"
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

	ApiGroup := router.Group("/o/v1")
	router2.InitOrderRouter(ApiGroup)
	router2.InitShopCartRouter(ApiGroup)
	return router
}
