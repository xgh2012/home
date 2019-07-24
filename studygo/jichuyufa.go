package main

import "fmt"

const freezingF, boilingF = 32.0, 212.0

func main() {
	//1、命名方式 驼峰法
	var boilingFx float64
	boilingFx = 0

	fmt.Println(boilingFx)

	//2、声明常量
	//const freezingFm,boilingFm= 32.0,212.0 注意常量作用域,这样声明只能在当前函数中使用
	changliang()

	//3变量声明  var 变量名字 类型 = 表达式
	var Xgh string = "xxx"
	fmt.Println(Xgh)

	//4 变量赋值
	var x int
	x = 1 // 命名变量的赋值
	/**p = true                   // 通过指针间接赋值
	person.name = "bob"         // 结构体字段赋值
	count[x] = count[x] * scale // 数组、slice或map的元素赋值*/
}

func changliang() {
	//2、声明常量
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
