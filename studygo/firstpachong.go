package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Array2Csv(data [][]string, filepath string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	w.WriteAll(data)
	w.Flush()

	file.Close()
}

func getDoc(url string) *goquery.Document {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func getData(doc *goquery.Document) [][]string {
	var list [][]string
	doc.Find("#searchResults>a").Each(func(i int, s *goquery.Selection) {
		productId, _ := s.Attr("data-product-id")
		styleId, _ := s.Attr("data-style-id")
		brandName := s.Find(".brandName").Text()
		productName := s.Find(".productName").Text()
		price6pm := s.Find(".price-6pm").Text()
		discountText := s.Find(".discount").Text() // (50% off MSRP $59.95)
		reg := regexp.MustCompile(`(\d+%) off MSRP (\$\d+.\d+)`)
		discountArr := reg.FindStringSubmatch(discountText)
		discount := discountArr[1]
		priceOrigin := discountArr[2]
		href, _ := s.Attr("href")
		href = "http://www.6pm.com/" + href
		item := []string{productId, styleId, productName, brandName, priceOrigin, price6pm, discount, href}
		fmt.Println(item)
		list = append(list, item)
	})
	return list
}

func getAllIndexPages(url string) []*goquery.Document {
	var docs []*goquery.Document

	doc := getDoc(url)
	nextPage := doc.Find("#resultWrap .pagination>a").Last().Prev().Text()
	nextPage = strings.Trim(nextPage, ". ")
	pageNum, _ := strconv.Atoi(nextPage)
	for i := 1; i < pageNum; i++ {
		tmpUrl := fmt.Sprintf("%v&p=%v", url, i)
		fmt.Println(tmpUrl)
		tmpdoc := getDoc(tmpUrl)
		docs = append(docs, tmpdoc)
	}
	return docs
}

func main() {
	// 根据来源取得所有列表页面对象
	pages := getAllIndexPages("http://www.6pm.com/null/4gIBMIIDA7OVAQ.zso?s=isNew/desc/goLiveDate/desc/recentSalesStyle/desc/")
	//  遍历列表页面取得商品数组
	var list [][]string

	for p := range pages {
		data := getData(pages[p])
		list = append(list, data...)
	}
	// 保存数组到csv
	Array2Csv(list, "list-data.csv")
}
