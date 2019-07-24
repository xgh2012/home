package service

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/nfnt/resize"
	_ "github.com/nfnt/resize"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)
var (
	dir1 = time.Now().Format("2006-01-02")
)

func GetContent(url string) (*goquery.Document)  {
	fmt.Println(url)
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil{
		log.Fatal(err)
	}
	return doc
}

//从html 中查找前三个图片
func FindImageAll(mainContent *goquery.Selection) (imagesSlice []string) {
	var ImagesSlice []string

	imgNum := 0
	mainContent.Find("img").Each(func(i int, selection *goquery.Selection) {
		imgSrc,err:= selection.Attr("src")
		if err==false || imgNum>2 {
			return
		}
		filename :=ImageDownload(imgSrc)
		if len(filename)<2 {
			return
		}
		imgNum++
		filepath := ImgServerHost+string([]rune(filename)[len(RootPath):])
		//fmt.Println(filepath)
		ImagesSlice=append(ImagesSlice,filepath)
	})
	return ImagesSlice
}

//下载图片文件
func ImageDownload(imagPath string) (filename string) {
	//创建文件夹
	dir := ImgPath+"/"+dir1
	aa, err := os.Stat(dir)
	if err!=nil || aa==nil{
		os.Mkdir(dir,0755)
	}
	if strings.Index(imagPath,"http://") ==-1 && strings.Index(imagPath,"https://") ==-1 {
		if strings.Index(imagPath,"//") ==0{
			imagPath = "http:"+imagPath
		}else {
			return ""
		}
	}
	resp, err := http.Get(imagPath)
	if err!=nil{
		log.Fatalln(err)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return ""
	}

	defer resp.Body.Close()

	path := strings.Split(imagPath, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}

	realName := ImageRealPath(name)
	out, err := os.Create(realName)

	if err!=nil{
		return
	}

	io.Copy(out, bytes.NewReader(body))
	checkOver := ImageAttrCheck(realName)
	if checkOver == false{
		return ""
	}

	newFileName := ResizeImage(realName)
	//删除原图
	os.Remove(realName)
	return newFileName
}

func ImageRealPath(name string) (realPath string)  {
	realPath = ImgPath+"/"+dir1+"/"+name
	return realPath
}

//文件属性检测
func ImageAttrCheck(fileName string) (isTure bool)  {
	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()
	image, _, err := image.Decode(reader)
	if err != nil {
		return false
	}

	x:= image.Bounds().Size().X
	y:= image.Bounds().Size().Y

	if x<100 || y<100 {
		return false
	}

	fileType := strings.Split(fileName, ".")
	fileTypeName := fileType[len(fileType)-1]
	allowArr := []string{"jpg", "jpeg", "png", "gif"}
	exists,index:=in_array(fileTypeName,allowArr)

	if exists == false || index==-1 {
		return false
	}
	return  true
}

//更新图片，按一定比例 或尺寸裁剪
func ResizeImage(name string) string  {
	file, err := os.Open(name)
	if err != nil {
		return ""
	}
	fileType := strings.Split(name, ".")
	fileTypeName := fileType[len(fileType)-1]
	newFileName := fileType[0]+"_new."+fileTypeName

	// decode jpeg into image.Image
	var  img image.Image
	switch fileTypeName {
		case "jpeg":
		case "jpg":
			img, err = jpeg.Decode(file)
			break
		case "png":
			img, err = png.Decode(file)
			break
		case "gif":
			img, err = gif.Decode(file)
			break
		default:
			return ""
	}

	if err != nil {
		return ""
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	if img==nil {
		return ""
	}
	m := resize.Resize(320, 180, img, resize.Lanczos3)

	out, err := os.Create(newFileName)
	if err != nil {
		return ""
	}
	defer out.Close()

	// write new image to file
	switch fileTypeName {
		case "jpeg":
		case "jpg":
			jpeg.Encode(out, m, nil)
			break
		case "png":
			png.Encode(out, m)
			break
		case "gif":
			gif.Encode(out, m,nil)
			break
		default:
			return ""
	}
	return newFileName
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

//写入文件
func WriteStringToFile(filepath, content string) {
	//打开文件，没有则创建，有则append内容
	w1, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	w1.Write([]byte(content))

	w1.Close()
}

/**
	字符串编码转换功能
	如：gbk转utf8
	ConvertToString(html, "gbk", "utf-8")
*/
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}