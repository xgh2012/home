package urlparams

import (
	"home/appconf/golog"
	"net/url"
)

type ParamsInfo struct {
	Prov      []string //省
	Area      []string //市
	Dist      []string //区
	Sys       []string //系统类型 IOS/AND
	Client    []string //客户端 lgj = 龙管家、baoan=易上网、eauth=E实名
	Evn       []string //是否测试环境
	Isgetconf []string //是否是单纯获取配置内容，供后台使用
}

func Urlparams(vars url.Values) ParamsInfo {
	var params ParamsInfo
	params.Prov, _ = vars["prov"]
	params.Area, _ = vars["area"]
	params.Dist, _ = vars["dist"]
	params.Sys, _ = vars["sys"]
	params.Client, _ = vars["client"]
	params.Isgetconf, _ = vars["isgetconf"]
	params.Evn, _ = vars["evn"]

	golog.Info.Println("params:", vars, "\r")
	return params
}
