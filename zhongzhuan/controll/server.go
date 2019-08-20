package controll

import (
	"encoding/json"
	"home/zhongzhuan/common"
	"home/zhongzhuan/model"
	"net/http"
	"net/url"
	"strconv"
)

//http接收数据
func WebdataTran(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"result":"1"}`)) //告诉对方收到了

	r.ParseForm() //解析参数
	//数据赋值
	common.GlobalParams.Barid = r.PostFormValue("barid")
	common.GlobalParams.Zhongzhuan = r.PostFormValue("zhongzhuan")
	common.GlobalParams.IVer = r.PostFormValue("iVer")
	common.GlobalParams.Itype = r.PostFormValue("itype")
	common.GlobalParams.Icmd = r.PostFormValue("icmd")
	common.GlobalParams.NoRedis = r.PostFormValue("NoRedis")
	common.GlobalParams.Data = r.PostFormValue("data")
	common.GlobalParams.Token = r.PostFormValue("token")

	go Test()
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

func Test() {
	//数据赋值 测试数据
	common.GlobalParams.Barid = "44030610001028"
	common.GlobalParams.Zhongzhuan = "zhongzhuan.topfreeweb.net%3A50001"
	common.GlobalParams.IVer = "0"
	common.GlobalParams.Itype = "0"
	common.GlobalParams.Icmd = "2006"
	common.GlobalParams.NoRedis = "1"
	common.GlobalParams.Data = "%7B%22CardNo%22%3A%22513922198707082852%22%2C%22ccxs%22%3A%22123%22%2C%22money%22%3A%225.65%22%2C%22notify_host%22%3A%22api2.topfreeweb.net%22%7D"
	common.GlobalParams.Token = "LZ2zhongzhuan:8f0653dff832b4bb2ee0b079e8124b5d"

	ParseData()

	model.LgjHeader()

	sendData := model.LgjBobys()

	resultCount, resultRcev := model.DoSend(sendData)

	model.LgjDescResult(resultCount, resultRcev)

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
