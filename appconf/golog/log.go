package golog

import (
	"log"
	"os"
)

var (
	Info *log.Logger
)

func init() {
	file, err := os.OpenFile("./log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	Info = log.New(file, "", log.Ltime)
}
