/*5、站点：电玩巴士-首页要闻
http://www.tgbus.com/list/yaowen/
*/
package tgbus

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"scanwebsite/service"
	"scanwebsite/writedb"
)

var (
	mainUrl    = "http://www.tgbus.com/list/yaowen/"
	mainDomain = "http://www.tgbus.com/list/yaowen/"
)

//获取主页列表
func TgbusGetList() {
	doc := service.GetContent(mainUrl)
	doc.Find(".information-item").Each(func(i int, selection *goquery.Selection) {
		href, err := selection.Find("a").Attr("href")
		if err != false {
			TgbusGetContents(href)
		}
	})
}

func TgbusGetContents(url string) {
	if writedb.QueryIsExist(url) == true {
		return
	}

	doc := service.GetContent(url)
	mainContent := doc.Find(".left-container")
	if mainContent.Nodes == nil {
		return
	}

	title := mainContent.Find("h1").Text()
	auths := mainContent.Find("h5")
	author := auths.Children().Eq(3).Text()
	createdt := auths.Children().Eq(4).Text()

	html, err := mainContent.Find(".article-main-contentraw").Html()
	if err != nil {
		return
	}

	//获取文章内容中的图片 前三张
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
		From:     "电玩巴士",
	}
	writedb.InsertToDB(info)
}
