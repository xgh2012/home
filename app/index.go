package main

import (
	"fmt"
	"home/app/common"
)

func main() {
	fmt.Println("Padding Test........")
	common.TestPadding()
	fmt.Println("AES ECB加密解密测试........")
	common.TestEncryptDecrypt()
}
