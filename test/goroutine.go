package main

import (
	"fmt"
	"time"
)

var complete = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s%d ", "x", i)
		time.Sleep(500 * time.Microsecond)
		complete <- i // 执行完毕了，发个消息
	}

}

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s%d ", "loop", i)
	}
}

func main() {
	go loop()
	x := <-complete // 直到线程跑完, 取到消息. main在此阻塞住
	fmt.Println(x)
	loop1()
}
