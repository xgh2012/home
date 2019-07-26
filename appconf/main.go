package main

import (
	"fmt"
	"github.com/iniconf"
	"home/appconf/filehandle"
	"home/appconf/golog"
	"home/appconf/provareadist"
	"home/appconf/urlparams"
	"log"
	"net/http"
)

var (
	realfilename = "conflist.ini"
	qzx          = "xxx"
)

func main() {
	golog.Info.Println("start program\r")
	httpSrv()
}

func httpSrv() {
	conf, err := iniconf.NewFileConf("./confile/cfgv2.ini")
	if err != nil {
		fmt.Println("error:load cfg.ini error:", err, "\r")
		return
	}

	ipport := conf.String("inf.ipport")
	http.HandleFunc("/srv/index.html", webdataTran)
	err1 := http.ListenAndServe(ipport, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err1)
	}
}

func webdataTran(w http.ResponseWriter, r *http.Request) {
	params := urlparams.Urlparams(r.URL.Query())
	prefilepath, istest := urlparams.Istest(params)
	var zfile = ""
	if len(params.Client) > 0 && params.Client[0] == "eauth" {
		if istest {
			zfile = "./confile/tmp/dev/eauthconflist.ini"
		} else {
			zfile = "./confile/tmp/prd/eauthconflist.ini"
		}
	} else {
		isGetConf := urlparams.IsGetconf(params) //计算是否是只获取配置文件

		areaExist, filepath := provareadist.Caclarea(prefilepath, params)
		fileExist, filename := provareadist.GetFilename(prefilepath, filepath, isGetConf)

		if isGetConf == true && (fileExist == false || areaExist == false) {
			str := "conffile not exist"
			w.Write([]byte(str))
		}

		zfile = filename
		isExist, _ := filehandle.PathExists(zfile)
		if isExist == false {
			zfile = prefilepath + realfilename + qzx
		}
	}
	golog.Info.Println("读取配置文件:", zfile, "\r")

	var str string
	str = filehandle.ReadFile(zfile)
	w.Write([]byte(str))
}
