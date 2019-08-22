package main

import (
	"home/zhongzhuan/controll"
	"net/http"
)

func main() {
	/*fmt.Println(common.Config.RediszzHostname)
	fmt.Println(common.Config.RediszzPort)
	fmt.Println(common.Config.RediszzPass)
	redishandle.Test()
	return*/
	/*controll.Test()
	return*/
	http.HandleFunc("/t/", controll.WebdataTran)
	http.ListenAndServe(":16789", nil)
}
