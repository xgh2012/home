/**
*站点：穿越火线官网-攻略中心
**/
package youwangxk

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"scanwebsite/service"
	"scanwebsite/writedb"
	"strings"
)

var (
	mainUrl = "https://www.gamersky.com/news/"
	//mainDomain = "https://cf.qq.com/cp"
)

//获取主页列表
func YwnewsGetList() {
	doc := service.GetContent(mainUrl)
	doc.Find(".tit .tt").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		YwnewsGetContents(href)
	})
}

func YwnewsGetContents(url string) {
	doc := service.GetContent(url)
	var (
		title    = ""
		createdt = ""
		author   = ""
		html     = ""
	)
	if strings.Index(url, "acg.gamersky.com") != -1 {
		leftContent := doc.Find(".Mid_L")
		title = leftContent.Find(".MidL_title h1").Text()
		auths := leftContent.Find(".detail")
		createdt = auths.Children().First().Text()
		author = auths.Children().Eq(3).Text()
		html, err := leftContent.Find(".MidL_con").Html()
		if err != nil {
			return
		}
		if html == "" {
			return
		}
		info := writedb.ActInfo{
			Title:    title,
			Contents: html,
			Href:     url,
			Laiyuan:  mainUrl,
			Author:   author,
			Created:  createdt,
		}
		fmt.Println(info)
	} else if strings.Index(url, "ol.gamersky.com") != -1 {
		leftContent := doc.Find(".Mid2L_ctt")
		title = leftContent.Find(".Mid2L_tit h1").Text()
		html, err := leftContent.Find(".Mid2L_con").Html()
		if err != nil {
			log.Fatal(err)
			return
		}
		if html == "" {
			return
		}
	}

	return
	info := writedb.ActInfo{
		Title:    title,
		Contents: html,
		Href:     url,
		Laiyuan:  mainUrl,
		Author:   author,
		Created:  createdt,
		From:     "游民星空",
	}
	writedb.InsertToDB(info)
}
