package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"home/wechat-transfer/common"
	"io/ioutil"
	"net/http"
)

func UnifiedOrder(c *gin.Context) {
	body := c.Request.Body
	contentType := `application/xml; charset="utf-8"`
	resp, err := common.HttpClient().Post(common.WechatUnifiedOrderUrl, contentType, body)
	if err != nil {
		logrus.Error(err)
		c.String(http.StatusBadGateway, "wechat error:%s", err.Error())
		return
	}
	defer resp.Body.Close()
	resultResponseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		c.String(http.StatusBadGateway, "server error:%s", err.Error())
		return
	}
	c.Data(http.StatusOK, contentType, resultResponseBody)
}
