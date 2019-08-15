package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bytes1, err := ioutil.ReadFile("M:/goProgram/src/home/test/20190713192612.pu")
	if err != nil {
		log.Fatal(err)
	}

	bytes2, err := ioutil.ReadFile("M:/goProgram/src/home/test/20190713202615.pu")
	if err != nil {
		log.Fatal(err)
	}

	length := len(bytes1)
	bytes3 := make([]byte, length)
	bytes4 := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes3[i] = bytes1[i] ^ bytes2[i]
		bytes4[i] = (bytes1[i]) + (bytes2[i])
	}
	//fmt.Println("bytes4 TypeOf:",reflect.TypeOf(bytes4[4102]))

	fmt.Println("bytes4 read:", bytes1)
	fmt.Println("bytes4 read:", bytes2)
	fmt.Println("bytes5 read:", bytes4)

	//ioutil.WriteFile("M:/goProgram/src/home/test/test.pu",bytes4,0766)
}
