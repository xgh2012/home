package main

import (
	"encoding/json"
	"fmt"
	"home/idcard/control"
)

type result struct {
	Result  string
	Message string
}

/*
*1、传出图片地址身份证正反面
*2、调用 腾讯OCR 接口进行识别，返回姓名 、头像等信息
*3、调用 百度人像分割 获取处理后的图片
*4、调用图片合成程序 合成新的图片
*5、返回信息
 */
func main() {
	//1、传出图片地址身份证正反面
	//control.init() baseapi.go

	//2调用 腾讯OCR 接口进行识别，返回姓名 、头像等信息
	tencentRes, tencentMes := control.TencentEntrance()
	if tencentRes != control.Success_Code {
		LastReuslt := result{Result: tencentRes, Message: tencentMes}
		output(LastReuslt)
		return
	}

	//3、调用 百度人像分割 获取处理后的图片
	baiduRes, baiduMes := control.BaiduEntrance()
	if baiduRes != control.Success_Code {
		LastReuslt := result{Result: baiduRes, Message: baiduMes}
		output(LastReuslt)
		return
	}

	//4、调用图片合成程序 合成新的图片 正面
	zmRes, zmMes := control.GetZhengMian()
	if zmRes != control.Success_Code {
		LastReuslt := result{Result: zmRes, Message: zmMes}
		output(LastReuslt)
		return
	}

	//5、调用图片合成程序 合成新的图片 反面
	fmRes, fmMes := control.GetFanMian()
	if fmRes != control.Success_Code {
		LastReuslt := result{Result: fmRes, Message: fmMes}
		output(LastReuslt)
		return
	}
	output(control.UserInfo)
}

func output(LastReuslt interface{}) {
	LastReusltString, _ := json.Marshal(LastReuslt)
	fmt.Printf("%s", LastReusltString)
}
