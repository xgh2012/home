/*9、站点：游侠网-最新资讯
https://www.ali213.net/news/
*/
package youxia

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"home/scanwebsite/service"
	"home/scanwebsite/writedb"
	"strings"
)

var (
	mainUrl    = "https://www.ali213.net/news/game/"
	mainDomain = "https://www.ali213.net/news/game/"
)

//获取主页列表
func YoyxiaGetList() {
	doc := service.GetContent(mainUrl)
	leftList := doc.Find(".n_lone")
	leftList.Children().Each(func(i int, selection *goquery.Selection) {
		href, err := selection.Find(".lone_f_l a").Attr("href")
		if err != false {
			YoyxiaGetContents(href)
		}
	})
}

func YoyxiaGetContents(url string) {
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

	title = doc.Find(".ns_t4 h1").Text()
	auths := doc.Find(".newstag").Children().First().Text()
	auths_arr := strings.Split(auths, " ")
	if len(auths_arr) < 2 {
		return
	}
	createdt = string([]rune(auths)[:16]) + ":00"
	author = auths_arr[1]
	auths_arr1 := strings.Split(author, "    ")
	if len(auths_arr1) > 2 {
		author = auths_arr1[2]
	} else {
		author = ""
	}
	mainContent = doc.Find("#Content")
	html, err = mainContent.Html()

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
		From:     "游侠网",
	}
	writedb.InsertToDB(info)
}
