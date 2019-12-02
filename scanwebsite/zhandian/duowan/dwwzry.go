package duowan

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"home/scanwebsite/service"
	"home/scanwebsite/writedb"
	"log"
	"strings"
)

var (
	wzryMainDomain = "http://wzry.duowan.com/tag/327319645493.html"
	wzryDomain     = "http://wzry.duowan.com/"
)

func WzryGetList() {
	doc := service.GetContent(wzryMainDomain)
	doc.Find(".list-pictxt li").Each(func(i int, s *goquery.Selection) {
		a := s.Find(".item-cont .title")
		href, err := a.Attr("href")
		if err != false {
			url := ""
			if strings.Index(href, "http://") != -1 {
				url = href
			} else {
				url = wzryDomain + href
			}
			WzryGetContents(url)
		} else {
		}
	})
}

func WzryGetContents(url string) {
	if writedb.QueryIsExist(url) == true {
		return
	}
	doc := service.GetContent(url)

	article := doc.Find(".col-box-bd article")
	title := article.Find("h1").Text()
	auths := article.Find("address")
	createdt := auths.Children().First().Text()
	author, _ := auths.Children().Eq(2).Html()
	if author == "作者：" {
		author = ""
	}

	mainContent := article.Find("#text")

	if mainContent.Nodes == nil {
		return
	}

	html, err := mainContent.Html()
	if err != nil {
		log.Fatal(err)
	}

	imagesSlice := service.FindImageAll(mainContent)
	imagesSliceJson, err := json.Marshal(imagesSlice)

	info := writedb.ActInfo{
		Title:    title,
		Contents: html,
		Href:     url,
		Laiyuan:  mainDomain,
		Author:   author,
		Images:   string(imagesSliceJson),
		Created:  createdt,
		From:     "多玩游戏",
	}
	writedb.InsertToDB(info)
}
