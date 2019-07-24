//微信下单 被微信风控 IP被封 禁止下单
//解决方法 从多个出口ip进行发送订单 PHP封装发送参数传递给go程序，go程序请求微信并返回给php

package main

import (
	"home/wechat-transfer/route"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := route.Router()
	r.Run(":9999")
}
