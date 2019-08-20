package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	login_url           string = "https://www.duduniu.cn/front/index"
	post_login_info_url string = "https://www.duduniu.cn/front/login"

	username string = "13330295142"
	password string = "4525674692ac06e619cdb3f1b4b65b08"
)

var cookies []*http.Cookie

func main() {
	login()
}

func login() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{
		Jar: jar,
	}
	req, _ := http.NewRequest("GET", login_url, nil)
	res_index, _ := client.Do(req)
	var temp_cookies = res_index.Cookies()
	defer res_index.Body.Close()

	c := &http.Client{
		Jar: jar,
	}
	//post数据
	//post数据
	postValues := url.Values{}
	postValues.Add("loginIdType", "userId")
	postValues.Add("province", "广东省")
	postValues.Add("city", "深圳市")
	postValues.Add("domainId", "1")
	postValues.Add("loginType", "userId")
	postValues.Add("userId", "6817")
	postValues.Add("password", "dxwk159357")
	postValues.Add("charPwd", "dxwk159357")
	postValues.Add("mianze", "on")
	postURL, _ := url.Parse(post_login_info_url)
	Jar, _ := cookiejar.New(nil)
	Jar.SetCookies(postURL, temp_cookies)
	c.Jar = Jar
	res, _ := c.PostForm(post_login_info_url,
		postValues)
	cookies = res.Cookies()
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(data))
	for k, v := range cookies {
		fmt.Printf("%v=%v\n", k, v)
	}
}

/**
字符串编码转换功能
如：gbk转utf8
ConvertToString(html, "gbk", "utf-8")
*/
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
