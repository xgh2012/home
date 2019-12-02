/*7、站点：多玩游戏网-24h轮播
http://www.duowan.com/news/*/

package duowan

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"home/scanwebsite/service"
	"home/scanwebsite/writedb"
	"strings"
)

var (
	yaowenmainDomain = "http://news.duowan.com/"
	yaowenDomain     = "http://news.duowan.com"
)

func YaowenGetList() {
	doc := service.GetContent(yaowenmainDomain)
	doc.Find(".mod-news ul li").Each(func(i int, selection *goquery.Selection) {
		href, err := selection.Find("a").Attr("href")
		if err != false {
			url := ""
			if strings.Index(href, "http") != -1 {
				url = href
			} else {
				url = yaowenDomain + href
			}
			YaowenGetContents(url)
		} else {
			//fmt.Println(href)
		}
	})
}

func YaowenGetContents(url string) {
	if writedb.QueryIsExist(url) == true {
		return
	}

	doc := service.GetContent(url)

	var (
		mainContent *goquery.Selection
		title       string
		createdt    string
		author      string
		html        string
		err         error
	)

	if strings.Index(url, "news.duowan.com") != -1 {
		mainContent = doc.Find(".artical-wrap")
		if mainContent.Nodes == nil {
			return
		}

		title = mainContent.Find("h1").Text()
		auths := mainContent.Find(".desc-wrap")
		createdt = auths.Children().First().Text()
		author = auths.Children().Eq(3).Text()
		html, err = mainContent.Find(".artical-bd").Html()

	} else if strings.Index(url, "lol.duowan.com") != -1 {
		mainContent = doc.Find(".ZQ-page--article")
		title = mainContent.Find("h1").Text()
		auths := mainContent.Find("address")
		createdt = auths.Children().First().Text()
		author = auths.Children().Eq(3).Text()

		html, err = mainContent.Find("#text").Html()
	} else {
		return
	}

	if err != nil {
		return
	}

	imagesSlice := service.FindImageAll(mainContent)
	imagesSliceJson, _ := json.Marshal(imagesSlice)

	info := writedb.ActInfo{
		Title:    title,
		Contents: html,
		Href:     url,
		Laiyuan:  mainDomain,
		Author:   author,
		Images:   string(imagesSliceJson),
		Created:  createdt,
		From:     "多玩电竞",
	}
	writedb.InsertToDB(info)
}
