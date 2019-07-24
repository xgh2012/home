package common

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	httpClient    http.Client
	httpTransport http.Transport
	ipList        = []string{
		"10.0.0.1",
		"10.0.0.2",
	}
	ipStatus = map[string]bool{
		"10.0.0.1": true,
		"10.0.0.2": true,
	}
)

const WechatUnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder" //微信下单地址

func init() {
	httpTransport = http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	httpClient = http.Client{
		Timeout: time.Duration(5 * time.Second),
		//Transport: &httpTransport,
	}
	go checkIpStatus()
}

func HttpClient() *http.Client {
	localTCPAddr := getLocalTCPAddr()
	httpTransport.DialContext = (&net.Dialer{
		LocalAddr: &localTCPAddr,
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext
	return &httpClient
}

var autoIncrementKey int64 = 0

func getLocalTCPAddr() net.TCPAddr {
	key := autoIncrementKey % int64(len(ipList))
	ipString := ipList[key]
	return net.TCPAddr{
		IP: net.ParseIP(ipString),
	}
}

func checkIpStatus() {
	for {
		for ip := range ipStatus {
			localTCPAddr := net.TCPAddr{
				IP: net.ParseIP(ip),
			}
			httpTransport.DialContext = (&net.Dialer{
				LocalAddr: &localTCPAddr,
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext

			resp, err := httpClient.Post(WechatUnifiedOrderUrl, "", strings.NewReader(""))
			if err != nil {
				ipStatus[ip] = false
				logrus.Error(err)
				continue
			}
			ipStatus[ip] = resp.StatusCode == http.StatusOK
			resultResponseBody, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				logrus.Error(err)
				continue
			}
			logrus.Info(string(resultResponseBody))
		}
		var list []string
		for ip, status := range ipStatus {
			if status {
				list = append(list, ip)
			}
		}
		ipList = list
		time.Sleep(60 * time.Second)
	}
}
