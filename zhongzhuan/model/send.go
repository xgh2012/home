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
		count  = 0
	)

	count, err = conn.Read(result)
	if err != nil {

	}

	return count, result[0:count]
}
