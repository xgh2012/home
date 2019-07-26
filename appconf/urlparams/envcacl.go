package urlparams

// 2 是否获取测试文件
func Istest(params ParamsInfo) (prefilepath string, istest bool) {

	if len(params.Evn) > 0 && params.Evn[0] == "dev" { //测试环境
		prefilepath = "./confile/devfile/"
		istest = true
	} else {
		prefilepath = "./confile/prdfile/"
		istest = false
	}
	return prefilepath, istest
}

//计算是否是只获取配置文件
func IsGetconf(params ParamsInfo) bool {
	isGetConf := false
	if len(params.Isgetconf) > 0 && params.Isgetconf[0] == "1" { //测试环境
		isGetConf = true
	} else {
		isGetConf = false
	}
	return isGetConf
}
