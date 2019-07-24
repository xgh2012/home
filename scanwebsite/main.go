package  main

import (
	"home/scanwebsite/zhandian/duowan"
	"home/scanwebsite/zhandian/tdmgame"
	"home/scanwebsite/zhandian/tewan"
	"home/scanwebsite/zhandian/u9"
	"home/scanwebsite/zhandian/youxia"
)

func main (){
	tdmgame.GetList()
	//return
	u9.U9GetList()
	//tgbus.TgbusGetList()
	youxia.YoyxiaGetList()
	duowan.YaowenGetList()
	duowan.NewsGetList()
	duowan.WzryGetList()
	duowan.LolGetList()
	tewan.TwGetList()

}