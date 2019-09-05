package main

import (
	"fmt"
	"home/app/common"
)

func main() {
	fmt.Println("Padding Test........")
	res1 := common.AesEncrypt("1234567890123456")
	fmt.Println(res1)
	fmt.Println("AES ECB加密解密测试........")
	res2 := common.AesDecrypt(res1)
	fmt.Println(res2)
}
