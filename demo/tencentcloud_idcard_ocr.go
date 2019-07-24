package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	_ "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/***
*身份证 识别 ocr
*参数见腾讯中心-身份证识别 ： https://cloud.tencent.com/document/product/866/33524
 */

/*IdCard，身份证照片，请求 CropIdCard 时返回；
Portrait，人像照片，请求 CropPortrait 时返回；*/
type AdvancedInfo struct {
	IdCard   string
	Portrait string
}

var (
	SecretId  = "SecretId"
	SecretKey = "SecretKey"
)

func main() {
	//xx()
	//return
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	// 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露密钥对危及你的财产安全。
	credential := common.NewCredential(
		SecretId, SecretKey)

	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用POST方法。
	// 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
	cpf.HttpProfile.ReqMethod = "POST"
	// SDK有默认的超时时间，非必要请不要进行调整。
	// 如有需要请在代码中查阅以获取最新的默认值。
	cpf.HttpProfile.ReqTimeout = 30
	// SDK会自动指定域名。通常是不需要特地指定域名的
	cpf.HttpProfile.Endpoint = "ocr.ap-guangzhou.tencentcloudapi.com"

	// 实例化要请求产品的client对象
	// 第二个参数是地域信息
	client, _ := ocr.NewClient(credential, "ap-guangzhou", cpf)
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定NewEnglishCompositionCorrectRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := ocr.NewIDCardOCRRequest()

	getFront(client, request)
	getBack(client, request)

}

func getFront(client *ocr.Client, request *ocr.IDCardOCRRequest) {
	filedir := "M:/goProgram/zheng.jpg"
	fileContents, err := ioutil.ReadFile(filedir)
	if err != nil {
		log.Fatal(err)
	}
	imgBase64 := base64.StdEncoding.EncodeToString(fileContents)

	request.ImageBase64 = common.StringPtr(imgBase64)
	request.CardSide = common.StringPtr("FRONT")
	request.Config = common.StringPtr(`{"CropIdCard":true,"CropPortrait":true}`)
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.IDCardOCR(request)
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	if err == nil {
		//日志开始
		var f *os.File
		filedir1 := "M:/goProgram/123.log"
		f, err = os.Create(filedir1) //创建文件
		if err != nil {
			fmt.Println("file create fail")
			return
		}
		//将文件写进去
		n, err1 := io.WriteString(f, response.ToJsonString())
		if err1 != nil {
			fmt.Println("write error", err1)
			return
		}
		fmt.Println("写入的字节数是：", n)
		//日志开始--结束
	}
	responseRss := &AdvancedInfo{}
	strs := *response.Response.AdvancedInfo
	json.Unmarshal([]byte(strs), responseRss)

	if responseRss.IdCard != "" {
		idcardDecode, err := base64.StdEncoding.DecodeString(responseRss.IdCard)
		if err != nil {
			fmt.Println("idcardDecode error", err)
			return
		}
		filesidcard := "M:/goProgram/sfzzhanpian.png"
		ioutil.WriteFile(filesidcard, idcardDecode, 0666)
	}

	if responseRss.Portrait != "" {
		PortraitDecode, err := base64.StdEncoding.DecodeString(responseRss.Portrait)
		if err != nil {
			fmt.Println("PortraitDecode error", err)
			return
		}
		filesidcard := "M:/goProgram/PortraitDecode.png"
		ioutil.WriteFile(filesidcard, PortraitDecode, 0666)
	}
}

func getBack(client *ocr.Client, request *ocr.IDCardOCRRequest) {
	filedir := "M:/goProgram/fan.jpg"
	fileContents, err := ioutil.ReadFile(filedir)
	if err != nil {
		log.Fatal(err)
	}
	imgBase64 := base64.StdEncoding.EncodeToString(fileContents)

	request.ImageBase64 = common.StringPtr(imgBase64)
	request.CardSide = common.StringPtr("BACK")
	request.Config = common.StringPtr(`{"CropIdCard":true,"CropPortrait":true}`)
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.IDCardOCR(request)
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	if err == nil {
		//日志开始
		var f *os.File
		filedir1 := "M:/goProgram/1234.log"
		f, err = os.Create(filedir1) //创建文件
		if err != nil {
			fmt.Println("file create fail")
			return
		}
		//将文件写进去
		n, err1 := io.WriteString(f, response.ToJsonString())
		if err1 != nil {
			fmt.Println("write error", err1)
			return
		}
		fmt.Println("写入的字节数是：", n)
		//日志开始--结束
	}
	responseRss := &AdvancedInfo{}
	strs := *response.Response.AdvancedInfo
	json.Unmarshal([]byte(strs), responseRss)

	if responseRss.IdCard != "" {
		idcardDecode, err := base64.StdEncoding.DecodeString(responseRss.IdCard)
		if err != nil {
			fmt.Println("idcardDecode error", err)
			return
		}
		filesidcard := "M:/goProgram/sfzzhanpian_fan.png"
		ioutil.WriteFile(filesidcard, idcardDecode, 0666)
	}

	if responseRss.Portrait != "" {
		PortraitDecode, err := base64.StdEncoding.DecodeString(responseRss.Portrait)
		if err != nil {
			fmt.Println("PortraitDecode error", err)
			return
		}
		filesidcard := "M:/goProgram/PortraitDecode_fan.png"
		ioutil.WriteFile(filesidcard, PortraitDecode, 0666)
	}
}
