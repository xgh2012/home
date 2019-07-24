package route

import (
	"github.com/gin-gonic/gin"
	"home/wechat-transfer/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/pay/unifiedorder", controller.UnifiedOrder)
	return r
}
