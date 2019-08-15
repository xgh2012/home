package main

import (
	"home/zhongzhuan/controll"
	"net/http"
)

func main() {
	controll.Test()
	return
	//model.GetHeader()
	http.HandleFunc("/", controll.WebdataTran)
	http.ListenAndServe("127.0.0.1:56789", nil)
}
