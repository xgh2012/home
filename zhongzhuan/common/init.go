package common

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

/*	参数
	token	字符串	最短16字节，最长不超过64字节，表示这个包的唯一性（最后16字节转发收银台）；
    zhognzhuan	中转url + port ，urlencode后的值
    barid	字符串	网吧编号，最长不能超过20个字符
    icmd	整数	命令号
    iVer	整数	协议版本标识：0：包体用INI方式分割，1：包体用json方式分割。
    itype	计费类型	0龙管家 1嘟嘟牛 2丕微 3万象 5宾馆
    data	字符串	采用urlcode进行编码（json格式）
*/
type Params struct {
	Token      string //唯一token
	Zhongzhuan string //中转url + port
	Barid      string //网吧编号
	Icmd       string //整数	命令号
	IVer       string //协议版本标识
	Itype      string //计费类型
	Data       string //字符串
	NoRedis    string //是否存入redis 0不存 1存
}

var (
	ConfigPath = "M:/goProgram/src/test.conf"
	Config     GlobalConfig
	RedisPool  *redis.Pool //创建redis连接池

	LogInfo *log.Logger //日志信息

	//ProgramStartTime = Now() //程序开启时间
	GlobalParams Params //全局变量
)

func init() {
	//加载配置文件
	Config, _ = LoadConf()
	LogInit(Config.LogPath, Config.LogFile)

	WriteSuccessLog("加载配置文件:" + ConfigPath)

	//创建redis连接池
	RedisPool = RedisNewPool(Config.RediszzHostname, Config.RediszzPort, Config.RediszzPass)

	WriteSuccessLog("创建redis连接池:" + Config.RediszzHostname)

}

//获取当前时间
func Now() (t string) {
	t = time.Now().Format("2006-01-02 15:04:05")
	return t
}
