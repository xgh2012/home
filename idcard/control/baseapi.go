package control

import (
	"encoding/json"
	"github.com/iniconf"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/**
*其他程序传入绝对地址
**/
type ImgInPathType struct {
	FrontInImgPath string //身份证正面照
	BackInImgPath  string //身份证背面照
}

/*
*主进程  处理用户数据
 */
type UserInfoType struct {
	RealName            string //姓名
	Sex                 string //性别
	Nation              string //民族
	Birthday            string //生日
	Address             string //住址
	Idcard              string //身份证号
	NativeIdcardImg     string //原始 图片地址 腾讯返回
	HeadImg             string //头像 图片地址
	Authority           string //发证机关
	ValidDate           string //有效期
	ForegroundImg       string //百度返回头像 png
	IdcardImgSmallFront string //合成图片(小)png
	IdcardImgBigFront   string //合成图片(大) png
	IdcardImgSmallBack  string //合成图片(小)png
	IdcardImgBigBack    string //合成图片(大) png
}

//配置项
type Tglbinf struct {
	Baidu_appId     string
	Baidu_apiKey    string
	Baidu_secretKey string

	Tencent_secretId  string
	Tencent_secretKey string
}

var (
	UserInfo     UserInfoType
	ImgInPath    = &ImgInPathType{}
	AppPath      string
	Ginf         Tglbinf
	Success_Code = "P0000"
	Success_Mes  = "成功"
)
var complete = make(chan int)

func init() {
	AppPath = GetAppPath()
	/*params:=os.Args
	if len(params) == 1 {
		fmt.Println("No data")
		return
	}
	data := params[1]*/
	data := `%7B%22FrontInImgPath%22%3A%22M%3A%5C%5Ctest%5C%5Czheng.jpg%22%2C%22BackInImgPath%22%3A%22M%3A%5C%5Ctest%5C%5Cfan.jpg%22%7D`
	data_json, err := url.QueryUnescape(data)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(data_json), ImgInPath)
	if err != nil {
		log.Fatal(err)
	}

	go LoadConfig()
	<-complete
}

//加载配置文件
func LoadConfig() {
	config, err := iniconf.NewFileConf(AppPath + "cfg.ini")
	//config, err := iniconf.NewFileConf("M:/goProgram/src/home/idcard/cfg.ini")
	//config, err := iniconf.NewFileConf("/Users/xgh/go/src/home/idcard/cfg.ini")

	if err != nil {
		log.Fatal(err)
	}
	Ginf.Baidu_appId = config.String("baidu.appId")
	Ginf.Baidu_apiKey = config.String("baidu.apiKey")
	Ginf.Baidu_secretKey = config.String("baidu.secretKey")

	Ginf.Tencent_secretId = config.String("tencent.secretId")
	Ginf.Tencent_secretKey = config.String("tencent.secretKey")
	complete <- 1 // 执行完毕了，发个消息
}

//获取当前文件夹绝对路径
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator)) + 1
	return path[:index]
}

func GetRealPath(path string) string {
	return AppPath + "../" + path
}
