package control

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	_ "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var access_token = "24.2e0c37451c777a35ba7e85b943ace144.2592000.1566657442.282335-16864563"

/**
 * 人像分割接口
 *
 * @param string $image - 图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
 * @param array $options - 可选参数对象，key: value都为string类型
 * @description options列表:
 *   type 可以通过设置type参数，自主设置返回哪些结果图，避免造成带宽的浪费<br>1）可选值说明：<br>labelmap - 二值图像，需二次处理方能查看分割效果<br>scoremap - 人像前景灰度图<br>foreground - 人像前景抠图，透明背景<br>2）type 参数值可以是可选值的组合，用逗号分隔；如果无此参数默认输出全部3类结果图
 * @return array
 */
type baidu struct {
	Foreground string `json:"foreground"`
}

func BaiduEntrance() (result string, message string) {

	//TODO 获取AccessToken，并缓存
	// getAccessToken()
	return bodySeg()
}

//获取Access Token
func getAccessToken() {
	url := "https://aip.baidubce.com/oauth/2.0/token"

	postString := "grant_type=client_credentials&client_id=" + Ginf.Baidu_apiKey + "&client_secret=" + Ginf.Baidu_secretKey

	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(postString))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

//人像分割接口
func bodySeg() (result string, message string) {
	urls := "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_seg?access_token=" + access_token
	content, err := ioutil.ReadFile(UserInfo.HeadImg)
	if err != nil {
		return "GB0001", "打开用户头像失败"
	}
	contentBase64 := base64.StdEncoding.EncodeToString(content)
	contentBase64 = url.QueryEscape(contentBase64)

	postString := "image=" + contentBase64
	resp, err := http.Post(urls,
		"application/x-www-form-urlencoded",
		strings.NewReader(postString))

	if err != nil {
		return "GB0002", "请求用户头像失败"
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "GB0003", "解析用户头像失败"
	}

	var ress baidu
	err = json.Unmarshal(body, &ress)
	if err != nil {
		return "GB0004", "解析用户头像失败"
	}

	imgDecode, err := base64.StdEncoding.DecodeString(ress.Foreground)
	if err != nil {
		return "GB0005", "解析用户头像失败"
	}
	UserInfo.ForegroundImg, UserInfo.ForegroundImgUlr = GetIdcardImageSavePath("result/baidu/" + UserInfo.Idcard + "_foreground.png")
	err = ioutil.WriteFile(UserInfo.ForegroundImg, imgDecode, 0666)
	if err != nil {
		return "GB0006", "存储用户头像失败"
	}
	return Success_Code, "成功"
}
