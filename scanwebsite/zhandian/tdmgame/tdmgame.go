/**
*站点：穿越火线官网-攻略中心
**/
package tdmgame

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"home/scanwebsite/service"
	"home/scanwebsite/writedb"
)

var (
	mainUrl = "https://www.3dmgame.com/news/"
)

//获取主页列表
func GetList()  {
	doc:=service.GetContent(mainUrl)

	//获取 大标题列表
	doc.Find(".Content_L .ul_ li").Each(func(i int, selection *goquery.Selection) {
		a:=selection.Find("a")
		href,err:=a.Attr("href")
		if(err!=false){
			documents:=service.GetContent(href)
			ParasParentDocuments(documents)
		}
	})
}

//解析父级内容
func ParasParentDocuments(documents *goquery.Document)  {
	//获取 新闻列表
	documents.Find(".Revision_list .selectpost").Each(func(i int, selection *goquery.Selection) {
		a:=selection.Children().Find("a")
		href,istrue:=a.Attr("href")
		if(istrue != false){
			GetContents(href)
		}
	})
}

func GetContents(url string)  {
	if writedb.QueryIsExist(url) == true{
		return
	}
	doc:=service.GetContent(url)
	mainContent:=doc.Find(".news_warp_center")
	html,err:=mainContent.Html()
	if(err!=nil){
		return
	}

	title:=doc.Find(".news_warp_top h1").Text()
	createdt:=doc.Find(".news_warp_top .time span").Text()
	author:=doc.Find(".news_warp_top .intem li .name").Text()
	if(err!=nil){
		return
	}

	imagesSlice := service.FindImageAll(mainContent)
	imagesSliceJson,err:= json.Marshal(imagesSlice)
	info := writedb.ActInfo{
		Title:title,
		Contents:html,
		Href:url,
		Laiyuan:mainUrl,
		Images:string(imagesSliceJson),
		Author:author,
		Created:createdt,
		From:"3dmgame",
	}
	writedb.InsertToDB(info)
}