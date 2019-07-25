package main

import (
	"encoding/json"
	"fmt"
	"home/idcard/control"
	_ "home/idcard/control"
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
	LastReuslt := result{}

	//腾讯获取证件信息
	tencentRes, tencentMes := control.TencentEntrance()
	if tencentRes != control.Success_Code {
		LastReuslt = result{Result: tencentRes, Message: tencentMes}
		LastReusltString, _ := json.Marshal(LastReuslt)
		fmt.Printf("%s", LastReusltString)
	}
}
