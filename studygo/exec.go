package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := "/usr/local/jdk/bin/java -jar /opt/app/public/eid_decrypt.jar f2d12b9e2f5e8de336a3103a13e175720ccbbbfe63f1fa90850dc38275752d55 BNazn4Q8Zupu3/iQqyyIM202pg9f8W13yDrduAOCg6/lRHH+k/SxtyJ9HddyUXXV9qjYDO0oR7wPEuC3paWakKEqqLoUUsxUYi+60fe8I0MMA/TkhaiAF4sOEPQNT4Ev403HdqCr3WjjT2tuDopvA0o= vC0FHDsVQ0j77e5USToNJ3Msa4dmfAOExrfp/cXm5U7duVpXyY0OYIacaA1f37FRWGTslQlNPdeQiRKdQpnjXQ=="
	res, _ := exec_shell(cmd)
	fmt.Println(res)
	fmt.Println(123)
}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func exec_shell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	fmt.Println(out.String())
	return out.String(), err
}
