package control

import (
	"encoding/base64"
	"encoding/json"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	_ "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"io/ioutil"
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

//var tencentComplete = make(chan string)

//入口
func TencentEntrance() (result string, message string) {
	//return
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	// 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露密钥对危及你的财产安全。
	credential := common.NewCredential(
		Ginf.Tencent_secretId, Ginf.Tencent_secretKey)

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

	//文件基本检测
	ckres, ckmessage := checkFile()
	if ckres != Success_Code {
		return ckres, ckmessage
	}

	//正面结果
	frontres, frontmessage := getFront(client, request)
	if frontres != Success_Code {
		return frontres, frontmessage
	}

	//背面结果
	backres, backmessage := getBack(client, request)
	if backres != Success_Code {
		return backres, backmessage
	}
	/*go getFront(client, request)
	<- tencentComplete

	go getBack(client, request)
	<- tencentComplete*/
	return Success_Code, Success_Mes
}

//检测正反面文件是否存在
func checkFile() (result string, message string) {
	if ImgInPath.FrontInImgPath == "" {
		return "GT0001", "正面照地址不全"
	}
	_, err := os.Stat(ImgInPath.BackInImgPath)
	if err != nil {
		return "GT0002", "正面照读取失败"
	}

	if ImgInPath.BackInImgPath == "" {
		return "GT0003", "背面照地址不全"
	}
	_, err = os.Stat(ImgInPath.BackInImgPath)
	if err != nil {
		return "GT0004", "背面照读取失败"
	}
	return Success_Code, "检测成功"
}

//获取正面信息
func getFront(client *ocr.Client, request *ocr.IDCardOCRRequest) (result string, message string) {
	filedir := ImgInPath.FrontInImgPath
	fileContents, err := ioutil.ReadFile(filedir)
	if err != nil {
		return "GT0005", "正面照获取失败"
	}

	imgBase64 := base64.StdEncoding.EncodeToString(fileContents)
	request.ImageBase64 = common.StringPtr(imgBase64)
	request.CardSide = common.StringPtr("FRONT")
	request.Config = common.StringPtr(`{"CropIdCard":true,"CropPortrait":true}`)

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.IDCardOCR(request)
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		return "GT0006", "初始化TX接口失败"
	}

	//结果赋值----------开始
	UserInfo.RealName = *response.Response.Name
	UserInfo.Sex = *response.Response.Sex
	UserInfo.Nation = *response.Response.Nation
	UserInfo.Birthday = *response.Response.Birth
	UserInfo.Address = *response.Response.Address
	UserInfo.Idcard = *response.Response.IdNum

	UserInfo.NativeIdcardImg = "data/idcard/result/tencent/" + UserInfo.Idcard + "_native.png"
	UserInfo.HeadImg = "data/idcard/result/tencent/" + UserInfo.Idcard + "_head.png"
	//结果赋值----------结束

	responseRss := &AdvancedInfo{}
	strs := *response.Response.AdvancedInfo
	json.Unmarshal([]byte(strs), responseRss)

	if responseRss.IdCard != "" {
		idcardDecode, err := base64.StdEncoding.DecodeString(responseRss.IdCard)
		if err != nil {
			return "GT0007", "身份证信息解析失败"
		}
		err = ioutil.WriteFile(GetRealPath(UserInfo.NativeIdcardImg), idcardDecode, 0666)
		if err != nil {
			return "GT0008", "保存身份证信息失败"
		}
	}

	if responseRss.Portrait != "" {
		PortraitDecode, err := base64.StdEncoding.DecodeString(responseRss.Portrait)
		if err != nil {
			return "GT0009", "头像信息解析失败"
		}
		filesidcard := AppPath + "../" + UserInfo.HeadImg
		err = ioutil.WriteFile(filesidcard, PortraitDecode, 0666)
		if err != nil {
			return "GT0010", "保存身份证信息失败"
		}
	}
	//tencentComplete <- "frontDone"
	return Success_Code, "成功"
}

//获取反面信息
func getBack(client *ocr.Client, request *ocr.IDCardOCRRequest) (result string, message string) {
	filedir := ImgInPath.BackInImgPath
	fileContents, err := ioutil.ReadFile(filedir)
	if err != nil {
		return "GT0011", "背面照获取失败"
	}
	imgBase64 := base64.StdEncoding.EncodeToString(fileContents)

	request.ImageBase64 = common.StringPtr(imgBase64)
	request.CardSide = common.StringPtr("BACK")
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.IDCardOCR(request)
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		return "GT0012", "初始化TX接口失败"
	}

	UserInfo.Authority = *response.Response.Authority
	UserInfo.ValidDate = *response.Response.ValidDate

	return Success_Code, Success_Mes
	//tencentComplete <- "backDone"
}
