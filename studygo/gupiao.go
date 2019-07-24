package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	t := time.Now().Unix()
	//int64
	//fmt.Println(reflect.TypeOf(t))
	//int to string
	t1 := strconv.FormatInt(t, 10)
	//t2 := strconv.Itoa(20)
	//t2 := strconv.Itoa(t1) //这个方法直接报错

	goget(t1, "sz002385")

}

func goget(t1 string, daima string) {
	//字符串拼接URL
	url := "https://hq.sinajs.cn/etag.php?_="
	url += t1
	url += "=&list="
	url += daima

	req, err := http.Get(url)
	if err != nil {

	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {

	}
	var i int
	for index, e := range body {
		i = index + 1
		if e == 44 {
			break
		}
	}
	bodystring := string(body[i:])
	str := strings.Trim(bodystring, ";")
	result := strings.Split(str, ",")
	//fmt.Println(result) //string
	//fmt.Println(reflect.TypeOf(result)) //string
	//
	//c,err:= strconv.ParseFloat(result[1],64)
	//if err!=nil{
	//
	//}
	//
	//jinkai,_:= strconv.ParseFloat(result[0],64) //今  开
	//zuoshou,_:= strconv.ParseFloat(result[1],64) //昨  收
	//zuidi,_:= strconv.ParseFloat(result[4],64) //最  低
	//zuigao,_:= strconv.ParseFloat(result[3],64) //最  高
	//dangqian, _ := strconv.ParseFloat(result[2], 64) //当前价格

	//fmt.Println(reflect.TypeOf(c))
	//fmt.Println(jinkai,zuidi,zuigao,zuoshou,dangqian)

	//fmt.Println(time.Now().YearDay())
	myprice := 5.710
	dangqianprice, _ := strconv.ParseFloat(result[2], 64)
	chayi := myprice - dangqianprice

	chayistr := strconv.FormatFloat(chayi, 'f', -1, 64)

	text := time.Now().Format("2006-01-02 15:04:05") + "\n代码：" + daima + "\n价格：" + result[2] + "\n今开：" + result[0] + ",昨收：" + result[1] + "\n最高：" + result[3] + ",最低：" + result[4] + "\n还差：" + chayistr + "\n"
	//ioutil.WriteFile("C:/Users/xgh/Desktop/123.txt", []byte(text), 0644)

	//var s Serverslice
	//str1 := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//json.Marshaler([]byte(str1), &s)
	//fmt.Println(s)
	//fmt.Printf(reflect.TypeOf(xx))
	fmt.Println(text)
	/*r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var t int
	go func() {

		for 1<2  {
			fmt.Println()
			t =r.Intn(1000000)*time.Microsecond
			fmt.Println(reflect.TypeOf(t))
			time.Sleep(t)
		}
	}()*/

	if chayi <= 0.1 {
		min, max := 1, 2
		for min < max {
			fmt.Println("请按回车键关闭...")
			var s string
			fmt.Scanf("%s", &s)
			min++
		}
	}

	/**
		0.78
	+2.03%
	39.12
	涨停：42.16跌停：34.50
	2019-02-18 14:43:17
	今  开：	38.38	成交量：	3.18万手	振  幅：	3.03%
	最  高：	39.51	成交额：	1.24亿元	换手率：	0.46%
	最  低：	38.35	总市值：	472.96亿	市净率：	2.83
	昨  收：	38.33	流通市值：	272.17亿	市盈率TTM：	35.38
	*/

	/**
	38.380 38.330 39.120 39.510 38.350 39.110 39.130 3177372 123499535.000 100 39.110 51900
	39.100 700 39.090 200 39.050 1800 39.040 27000 39.130 1800 39.140 3600 39.150 8800 39.160 9550 39.170 2019-02-18 14:43:02 00
	*/
	//var hwnd winapi.HWND
	//r:= winapi.MessageBox(nil, "Hello, World!", "Hi!", MB_OK )
	//fmt.Println(r)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	fmt.Println(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
