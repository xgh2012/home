/*6、久游网英雄联盟-新闻资讯
http://lol.uuu9.com/List_5138.shtml
*/
package u9

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"home/scanwebsite/service"
	"home/scanwebsite/writedb"
	"strings"
)

var (
	mainUrl = "http://lol.uuu9.com/List_5138.shtml"
	mainDomain = "http://lol.uuu9.com/List_5138.shtml"
)

//获取主页列表
func U9GetList()  {
	doc:=service.GetContent(mainUrl)
	doc.Find(".textlist li").Each(func(i int, selection *goquery.Selection) {
		href,err:=selection.Find("a").Attr("href")
		if err!=false{
			U9GetContents(href)
		}
	})
}

func U9GetContents(url string)  {
	if writedb.QueryIsExist(url) == true{
		return
	}

	doc:=service.GetContent(url)
	mainContent := doc.Find(".robing_con")
	if mainContent.Nodes == nil{
		return
	}

	title:=mainContent.Find("h1").Text()
	title = service.ConvertToString(title, "gbk", "utf-8")

	auths:=service.ConvertToString(mainContent.Find("h4").Text(),"gbk", "utf-8")
	auths_arr := strings.Split(auths," 聽聽 ")
	if len(auths_arr) <3{
		return
	}
	createdt := string([]rune(auths)[:17])
	author:=auths_arr[2]

	text:=mainContent.Find("#textdetail")
	if len(text.Nodes) == 0 {
		text =mainContent.Find("textdetail")
	}
	text.Find("h4").Remove()
	text.Find(".Introduction").Remove()
	html,err:=text.Html()
	if err!=nil || html == ""{
		return
	}

	html = service.ConvertToString(html,"gbk", "utf-8")

	imagesSlice := service.FindImageAll(mainContent)
	imagesSliceJson,err:= json.Marshal(imagesSlice)

	info := writedb.ActInfo{
		Title:title,
		Contents:html,
		Href:url,
		Laiyuan:mainDomain,
		Author:author,
		Images:string(imagesSliceJson),
		Created:createdt,
		From:"U9网",
	}
	writedb.InsertToDB(info)
}