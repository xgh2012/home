package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {
	//fmt.Println("Hello, world Println",122)
	//fmt.Println(os.Args[2:])

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	a := strings.Join(os.Args[1:], " ")
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)

	cmd := exec.Command("echo", "hello")
	buf, _ := cmd.Output() // 错误处理略

	print(string(buf))
}
