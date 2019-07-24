package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"net/http"
)

func main() {
	// Load HTML file.
	/*f, err := os.Open(`./examples/test.html`)
	if err != nil {
		panic(err)
	}*/

	res, err := http.Get("http://dota.uuu9.com/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	f:=res.Body
	// Parse HTML document.
	doc, err := htmlquery.Parse(f)
	if err != nil{
		panic(err)
	}

	// Option 1: using xpath's expr to matches nodes.
	expr := xpath.MustCompile("count(//div[@class='article'])")
	//fmt.Printf("%f \n", expr.Evaluate(htmlquery.CreateXPathNavigator(doc)).(float64))

	expr = xpath.MustCompile("/html/body/div[@class='all']/div[@class='main']/div[@class='cl p10'][1]/div[@class='news']/div[@class='content'][1]/div[@class='newslist_box'][1]/ul[@class='newslist']")
	expr = xpath.MustCompile("/html/body/div[@class='all']/div[@class='main']/div[@class='cl p10'][1]/div[@class='news']/div[@class='content'][1]/div[@class='newslist_box'][1]/ul[@class='newslist']")
	iter := expr.Evaluate(htmlquery.CreateXPathNavigator(doc)).(*xpath.NodeIterator)
	for iter.MoveNext() {
		//fmt.Printf("%s \n", service.ConvertToString(iter.Current().Value(),"gbk","utf-8")) // output href
		//fmt.Printf("%s \n", service.ConvertToString(iter.Current().Value(),"gbk","utf-8")) // output href
	}
	nodes := htmlquery.Find(doc,"/html/body/div[@class='all']/div[@class='main']/div[@class='cl p10'][1]/div[@class='news']/div[@class='content'][1]/div[@class='newslist_box'][1]/ul[@class='newslist']")
	for _,node:=range nodes {
		url := htmlquery.FindOne(node, "./a/@href")
		fmt.Println(url)
		fmt.Println(node.Attr)
	}

	// Option 2: using build-in functions Find() to matches nodes.
	//for _, n := range htmlquery.Find(doc, "//a/@href") {
	//	fmt.Printf("%s \n", htmlquery.SelectAttr(n, "href")) // output href
	//}
}