package main

import (
	"fmt"
	"home/zz/lgj"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	C_HeadLgjLen = 71
	Info         *log.Logger
)

type Tlgjsendata struct {
	Stoken   string
	Sbarid   string
	Icmd     int
	IVer     int
	INoRedis int
	Ixtype   int
	Imark    int
	Sdata    string
}

type Tlgjlistnode struct {
	obj  Tlgjsendata
	next *Tlgjlistnode
}

type Tlgjlist struct {
	head    *Tlgjlistnode
	tail    *Tlgjlistnode
	mymutex sync.Mutex
}

func main() {
	WriteSuccessLog("start")
	//lgj.LgjDescResult(lgj.DoSend(lgj.SendData()))
	/*var sendObj Tlgjsendata
	sendObj.Icmd = 2006
	sendObj.INoRedis = 0
	sendObj.IVer = 0
	sendObj.Sdata = "%7B%22CardNo%22%3A%22513922198707082852%22%2C%22ccxs%22%3A%22123%22%2C%22money%22%3A%225.65%22%2C%22notify_host%22%3A%22api2.topfreeweb.net%22%7D"
	sendObj.Sbarid = "44030610001028"
	sendObj.Stoken = "1234567890123456"

	//Tlgjlist.Adddata(sendObj)
	return*/

	//创建连接池
	var tcpsrv = "zhongzhuan.topfreeweb.net:50001"
	go CteateAssociation(tcpsrv)

	var i = 0
	for i < 500 {
		time.Sleep(5 * time.Second)
		i++
	}
}

//创建协程
func CteateAssociation(tcpsrv string) (t string) {
	var icount = 10 //默认每个中转10个长连接，不够再加
	var ino, i int
	var bok bool
	tmpchannel := make(chan int, icount*3)
	for i = 1; i <= icount; i++ {
		go ReadAssociationLgj(tcpsrv, tmpchannel, i)
	}

	for {
		ino, bok = <-tmpchannel
		WriteSuccessLog(strconv.Itoa(ino))
		time.Sleep(time.Duration(5) * time.Second)
		if ino > 0 && bok {
			if (ino > 0) && (ino < 11) {
				WriteSuccessLog("ExGet:龙管家中转异常,5秒后重连")
				go ReadAssociationLgj(tcpsrv, tmpchannel, ino)
			}
		}
	}
}

//读协程 龙管家的
func ReadAssociationLgj(tcpsrv string, ch chan int, ino int) {
	var (
		result     []byte
		bussResult string
		b          int
		iok        bool
	)

	conn, err := net.DialTimeout("tcp", tcpsrv, time.Second*6)
	if err != nil {
		WriteSuccessLog("conn lgjzhongzhuan faild,err:" + err.Error())
		ch <- ino
		return
	}
	defer conn.Close()
	var tmpsendlist *Tlgjlist
	tmpsendlist = new(Tlgjlist)

	go longconnlgj_write(conn.(*net.TCPConn), tmpsendlist)

	for {
		b, err = io.ReadFull(conn, result)
		if (err != nil) || (b < C_HeadLgjLen) {
			WriteSuccessLog("read1 lgjzhongzhuan faild,err:")
			break
		}
		bussResult, iok = lgj.LgjDescResult(result)
		if iok == false {
			WriteSuccessLog("read2 lgjzhongzhuan faild,err:" + err.Error())
		}
	}
	WriteSuccessLog(bussResult)
	ch <- ino
}

//写协程
func longconnlgj_write(clt *net.TCPConn, datalist *Tlgjlist) {
	var (
		iok    bool
		err    error
		icount = 0
	)

	fmt.Println(datalist)

	for {
		sendObj := lgj.SendData()
		time.Sleep(time.Duration(20) * time.Second)

		if iok || (icount > 2000) {
			_, err = clt.Write(sendObj)
			if err != nil {
				clt.Close()
				WriteSuccessLog("send lgjzhongzhuan head faild:" + err.Error())
				break
			}
			icount = 0
		} else {
			icount = icount + 1
		}
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
}

func WriteSuccessLog(str string) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	Info = log.New(file, "", log.Ldate|log.Ltime)
	fmt.Println(str)
	Info.Println(Now(), str, "\r")
}
func Now() (t string) {
	t = time.Now().Format("2006-01-02 15:04:05")
	return
}

// 添加需要发送的数据 生产者
func (p *Tlgjlist) Adddata(data Tlgjsendata) {
	p.mymutex.Lock()
	defer p.mymutex.Unlock()
	var node *Tlgjlistnode
	node = new(Tlgjlistnode)
	node.obj = data
	node.next = nil
	if p.head == nil {
		p.head = node
		p.tail = node
	} else {
		p.tail.next = node
		p.tail = node
	}
	fmt.Println(p.tail.next)
}
