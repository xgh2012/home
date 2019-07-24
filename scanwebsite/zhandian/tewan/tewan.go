/**
*站点：穿越火线官网-攻略中心
**/
package tewan

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"home/scanwebsite/service"
	"home/scanwebsite/writedb"
)

var (
	mainUrl = "http://lol.te5.com/xinwenzhongxin/list_73_2.html"
)

//获取主页列表
func TwGetList()  {
	doc:=service.GetContent(mainUrl)
	leftList := doc.Find(".listLol .arc_left")
	leftList.Children().Each(func(i int, selection *goquery.Selection) {
		href,err:=selection.Attr("href")
		if err!=false{
			TwGetContents(href)
		}
	})
}

func TwGetContents(url string)  {
	if writedb.QueryIsExist(url) == true{
		return
	}
	doc:=service.GetContent(url)
	selection := doc.Find(".mainbg")
	title := selection.Find("h1").Text()

	html,err:=selection.Find("#text").Html()
	if err!=nil {
		return
	}
	if html == ""{
		html,err=selection.Find(".artical-bd").Html()
		if err!=nil {
			return
		}
	}
	if html == ""{
		return
	}

	arcinfo := selection.Find(".arcinfo")
	createdt := arcinfo.Find(".time").Text()
	author := arcinfo.Find(".author").Text()


	imagesSlice := service.FindImageAll(selection.Find("#text"))
	imagesSliceJson,err:= json.Marshal(imagesSlice)

	info := writedb.ActInfo{
		Title:title,
		Contents:html,
		Href:url,
		Laiyuan:mainUrl,
		Author:author,
		Images:string(imagesSliceJson),
		Created:createdt,
		From:"特玩网",
	}
	writedb.InsertToDB(info)
}