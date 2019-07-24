package duowan

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"home/scanwebsite/service"
	"home/scanwebsite/writedb"
	"strings"
)

var (
	mainDomain = "http://lol.duowan.com/tag/131026281878.html"
	LolDomain = "http://lol.duowan.com"
)


func LolGetList()  {
	doc:=service.GetContent(mainDomain)
	doc.Find(".m-list li").Each(func(i int, selection *goquery.Selection){
		href,err:=selection.Find(".item-cover a").Attr("href")
		if(err!=false){
			url:=""
			if(strings.Index(href,"http://")!=-1){
				url=href
			}else{
				url=LolDomain+href
			}
			LolGetContents(url)
		}else{
			//fmt.Println(href)
		}
	})
}

func LolGetContents(url string)  {
	if writedb.QueryIsExist(url) == true{
		return
	}
	doc:=service.GetContent(url)
	if  strings.Index(url,"news.duowan.com")!=-1 {
		doc.Find(".artical-wrap").Each(func(i int, s *goquery.Selection){
			/*html,err:=s.Html()
			if(err!=nil){
				log.Fatal(err)
			}
			fmt.Println(html)*/
			//writedb.InsertToDB(url,html)
		})
	}else{
		doc.Find(".ZQ-page--article").Each(func(i int, s *goquery.Selection){
			title:=s.Find("h1").Text()
			auths:=s.Find("address")
			createdt := auths.Children().First().Text()
			author := auths.Children().Eq(3).Text()

			html,err:=s.Find("#text").Html()
			if(err!=nil){
				return
			}

			imagesSlice := service.FindImageAll(s.Find("#text"))
			imagesSliceJson,err:= json.Marshal(imagesSlice)

			info := writedb.ActInfo{
				Title:title,
				Contents:html,
				Href:url,
				Laiyuan:mainDomain,
				Author:author,
				Images:string(imagesSliceJson),
				Created:createdt,
				From:"多玩游戏",
			}
			writedb.InsertToDB(info)
		})
	}
}