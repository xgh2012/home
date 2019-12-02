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
	newsmainDomain = "http://www.duowan.com/news/"
	newsDomain     = "http://www.duowan.com"
)

func NewsGetList() {
	doc := service.GetContent(newsmainDomain)
	doc.Find(".day-item").Each(func(i int, selection *goquery.Selection) {
		href, err := selection.Find(".item-fl a").Attr("href")
		if err != false {
			url := ""
			if strings.Index(href, "http") != -1 {
				url = href
			} else {
				url = newsDomain + href
			}
			NewsGetContents(url)
		} else {
			//fmt.Println(href)
		}
	})
}

func NewsGetContents(url string) {
	if writedb.QueryIsExist(url) == true {
		return
	}

	doc := service.GetContent(url)
	mainContent := doc.Find(".art-cont")
	if mainContent.Nodes == nil {
		return
	}

	title := mainContent.Find("h1").Text()
	auths := mainContent.Find(".meta-info")
	createdt := auths.Children().First().Text()
	author := auths.Children().Eq(1).Text()

	html, err := mainContent.Find("#text").Html()
	if err != nil {
		return
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
		From:     "多玩游戏网",
	}
	writedb.InsertToDB(info)
}
