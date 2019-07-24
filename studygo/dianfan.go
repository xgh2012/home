package main

import (
	"fmt"
	"time"
)

func main() {
	weekday := time.Now().Weekday()
	if weekday != time.Sunday && weekday != time.Saturday {
		fmt.Println("点饭时间到了...")
		min, max := 1, 2
		for min < max {
			fmt.Println("请按回车键关闭...")
			var s string
			fmt.Scanf("%s", &s)
			min++
		}
	}
}
