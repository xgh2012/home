package controll

import (
	"encoding/json"
	"home/zhongzhuan/common"
	"home/zhongzhuan/model"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var pool = common.RedisPool

const (
	Time_Layout_1 = "2006-01-02 15:04:05"
	Time_Layout_2 = "20060102150405"
)

//http接收数据
func WebdataTran(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"result":"1"}`)) //告诉对方收到了

	r.ParseForm() //解析参数
	//数据赋值
	common.GlobalParams.Barid = r.FormValue("barid")
	common.GlobalParams.Zhongzhuan = r.FormValue("zhongzhuan")
	common.GlobalParams.IVer = r.FormValue("iVer")
	common.GlobalParams.Itype = r.FormValue("itype")
	common.GlobalParams.Icmd = r.FormValue("icmd")
	common.GlobalParams.NoRedis = r.FormValue("NoRedis")
	common.GlobalParams.Data = r.FormValue("data")
	common.GlobalParams.Token = r.FormValue("token")

	go getRes()

	/*if len(r.Header) > 0 {
		for k,v := range r.Header {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	fmt.Println("打印Form参数列表：")
	r.ParseForm()
	if len(r.Form) > 0 {
		for k,v := range r.Form {
			fmt.Printf("%s  %s=%s\n", reflect.TypeOf(k),k, v[0])

		}
	}*/
}

//获取最终业务结果
func getRes() {
	ParseData()
	sendData := model.LgjBodys(common.GlobalParams)
	resultCount, resultRcev := model.DoSend(sendData)
	bussRes := model.LgjDescResult(resultCount, resultRcev)
	client := pool.Get()
	defer client.Close()
	common.GlobalParams.Token = common.GlobalParams.Token[14:] + time.Now().Format(Time_Layout_2)
	client.Do("SET", common.GlobalParams.Token, bussRes)
	client.Do("EXPIRE", common.GlobalParams.Token, 500)
}

//数据解析
func ParseData() {
	var err error

	//解析url Urldecode
	common.GlobalParams.Zhongzhuan, err = url.QueryUnescape(common.GlobalParams.Zhongzhuan)
	if err != nil {
	}

	//解析数据 Urldecode
	common.GlobalParams.Data, err = url.QueryUnescape(common.GlobalParams.Data)
	if err != nil {

	}

	//整数	协议版本标识：0：包体用INI方式分割，1：包体用json方式分割。
	if common.GlobalParams.IVer == "0" && len(common.GlobalParams.Data) > 0 {
		var IVerf interface{}
		err = json.Unmarshal([]byte(common.GlobalParams.Data), &IVerf)
		if err != nil {

		}
		var tmpData1 string
		IVerm := IVerf.(map[string]interface{})
		for IVk, IVv := range IVerm {
			switch IVv.(type) {
			case string:
				tmpData1 = tmpData1 + IVk + "=" + IVv.(string) + "\r\n"
				break
			case int:
				tmpData1 = tmpData1 + IVk + "=" + strconv.Itoa(IVv.(int)) + "\r\n"
				break
			case float64:
				tmpData1 = tmpData1 + IVk + "=" + strconv.FormatFloat(IVv.(float64), 'f', 2, 64) + "\r\n"
				break
			case bool:
				tmpData1 = tmpData1 + IVk + "=" + strconv.FormatBool(IVv.(bool)) + "\r\n"
				break
			default:
				break
			}
		}
		//重新赋值
		common.GlobalParams.Data = tmpData1
	}
}
