package model

import (
	"fmt"
	"home/zhongzhuan/common"
	"io"
	"log"
	"net"
)

//向中转发送数据
func DoSend(sendData []byte) (int, []byte) {
	conn, err := net.Dial("tcp", common.GlobalParams.Zhongzhuan)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(conn, string(sendData))
	if err != nil {
		fmt.Println("err = ", err.Error())
	}

	var (
		result = make([]byte, 1024)
		read   = true
		count  = 0
	)

	for read {
		count, err = conn.Read(result)
		read = (err != nil)
	}

	//fmt.Println(string(result[0:count]))
	return count, result[0:count]
}
