package common

import (
	"log"
	"os"
)

func LogInit(path string, name string) {
	realpath := path + name
	file, err := os.OpenFile(realpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	LogInfo = log.New(file, "", log.Ldate|log.Ltime)
	WriteSuccessLog("start program")
}

/**
*记录成功日志
**/
func WriteSuccessLog(str string) {
	LogInfo.Println(Now(), str, "\r")
}
