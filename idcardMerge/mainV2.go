package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"log"
	"net/url"
	"os/exec"
	"path/filepath"
	"strings"
	//"github.com/ChengjinWu/imagedraw" #存在bug 需要修复
	"home/idcardMerge/imagedraw"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

/*
*1、传出图片地址身份证正反面
*2、调用 腾讯OCR 接口进行识别，返回姓名 、头像等信息
*3、调用 百度人像分割 获取处理后的图片
*4、调用图片合成程序 合成新的图片
*5、返回信息
 */

type UserInfoType struct {
	RealName  string //姓名
	Sex       string //性别
	Nation    string //民族
	Birthday  string //生日
	Address   string //住址
	Idcard    string //身份证号
	HeadImg   string //头像 图片地址
	Authority string //发证机关
	ValidDate string //有效期
}

var (
	UserInfo = &UserInfoType{}
	AppPath  string
)

func main() {
	AppPath = GetAppPath()
	params := os.Args
	if len(params) == 1 {
		fmt.Println("No data")
		return
	}
	data := params[1]
	//data := `%7B%22RealName%22%3A%22%E7%86%8A%E9%AB%98%E6%B5%B7%22%2C%22Sex%22%3A%22%E7%94%B7%22%2C%22Nation%22%3A%22%E6%B1%89%22%2C%22Birthday%22%3A%221987%5C%2F7%5C%2F8%22%2C%22Address%22%3A%22%E6%88%90%E9%83%BD%E5%B8%82%E5%A4%A9%E5%BA%9C%E6%96%B0%E5%8C%BA%E5%8D%8E%E9%98%B3%E5%A4%A9%E5%BA%9C%E5%A4%A7%E9%81%93%E5%8D%97%E6%AE%B52389%E5%8F%B71%E6%A0%8B1%E5%8D%95%E5%85%838%E6%A5%BC802%E5%8F%B7%22%2C%22Idcard%22%3A%22513922198707082852%22%2C%22HeadImg%22%3A%22513922198707082852%22%2C%22Authority%22%3A%22%E6%88%90%E9%83%BD%E5%B8%82%E5%85%AC%E5%AE%89%E5%B1%80%E5%A4%A9%E5%BA%9C%E6%96%B0%E5%8C%BA%E5%88%86%E5%B1%80%22%2C%22ValidDate%22%3A%221999.12.11-2039.12.11%22%7D`
	data_json, err := url.QueryUnescape(data)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(data_json), UserInfo)
	if err != nil {
		log.Fatal(err)
	}
	ZhengmianV3()
	return
	//FangmianV3()
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator)) + 1
	return path[:index]
}

/**
出生年月日 方正黑体简体
字符大小：姓名＋号码（11点）其他（9点）
字符间距（AV）：号码（50）
字符行距：住址（12点）
**/
func ZhengmianV3() {
	textBrush, err := imagedraw.NewTextBrush(AppPath+"../data/font/黑体.ttf", 50, image.Black, 300)
	if err != nil {
		log.Fatalln(err)
	}
	backgroundImg, err := imaging.Open(AppPath + "../data/idcardMerge/source/shenfenzheng_zhengmian.png")
	if err != nil {
		log.Fatalln(err)
	}

	backgroundImgBounds := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(backgroundImgBounds, backgroundImg.Bounds(), backgroundImg, image.ZP, draw.Src)

	//姓名
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(358, 172), UserInfo.RealName+`
`)

	//性别
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(358, 312), UserInfo.Sex+`
`)
	//民族
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(735, 312), UserInfo.Nation+`
	`)

	//住址
	textBrush.TextWidth = 700
	textBrush.FontSize = 55
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(358, 593), UserInfo.Address+`
`)

	textBrushBirthday, err := imagedraw.NewTextBrush(AppPath+"../data/font/方正黑体简体_0.ttf", 50, image.Black, 300)
	if err != nil {
		log.Fatalln(err)
	}

	Birthday := strings.Split(UserInfo.Birthday, "/")

	//出生年
	textBrushBirthday.DrawFontOnRGBA(backgroundImgBounds, image.Pt(358, 452), Birthday[0]+`
`)

	//出生月
	textBrushBirthday.DrawFontOnRGBA(backgroundImgBounds, image.Pt(680, 452), Birthday[1]+`
`)

	//出生日
	textBrushBirthday.DrawFontOnRGBA(backgroundImgBounds, image.Pt(860, 452), Birthday[2]+`
`)

	//身份证号
	textBrushIdcard, err := imagedraw.NewTextBrush(AppPath+"../data/font/OCR-B 10 BT.TTF", 70, image.Black, 980)
	if err != nil {
		log.Fatalln(err)
	}
	textBrushIdcard.DrawFontOnRGBA(backgroundImgBounds, image.Pt(658, 916), UserInfo.Idcard+`
`)

	imgRGBA, err := GetHeadImageRGBAV3(UserInfo.HeadImg)
	if err != nil {
		log.Fatalln(err)
	}

	x0, y0 := 1155, 80
	draw.DrawMask(backgroundImgBounds, image.Rect(x0, y0, x0+636, y0+775), imgRGBA, image.ZP, imgRGBA, image.ZP, draw.Over)

	//backgroundImgBounds=fzImageV3(backgroundImgBounds)
	imaging.Save(backgroundImgBounds, AppPath+"../data/idcardMerge/result/zhengmian_big.jpg")

	//固定图片大小输出
	resultImgDecode := resize.Resize(660, 422, backgroundImgBounds, resize.Lanczos3)
	imaging.Save(resultImgDecode, AppPath+"../data/idcardMerge/result/zhengmian_small.jpg")
}

/*func FangmianV3() {
	textBrush, err := imagedraw.NewTextBrush("../data/font/黑体.ttf", 60, image.Black, 1280)
	if err != nil {
		log.Fatalln(err)
	}
	backgroundImg, err := imaging.Open("../data/idcardMerge/source/shenfenzheng_fanmian.png")
	if err != nil {
		log.Fatalln(err)
	}

	backgroundImgBounds := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(backgroundImgBounds, backgroundImg.Bounds(), backgroundImg, image.ZP, draw.Src)

	//有效日期
	textBrush.TextDpi = 75
	//textBrush.FontSize = 65
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(780, 937), UserInfo.ValidDate+`
	`)

	//发证机关
	textBrush.TextDpi = 80
	//textBrush.FontSize = 60
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(780, 795), UserInfo.Authority+`
	`)

	imaging.Save(backgroundImgBounds, "../data/idcardMerge/result/fangmian_big.jpg")
	resultImgDecode := resize.Resize(660, 422, backgroundImgBounds, resize.Lanczos3)
	imaging.Save(resultImgDecode, "../data/idcardMerge/result/fangmian_small.jpg")
}*/

func GetHeadImageRGBAV3(iamgePath string) (*image.RGBA, error) {
	img, err := imaging.Open(iamgePath)
	if err != nil {
		return nil, err
	}
	touxiang, _ := os.Open(iamgePath)
	headImgDecode, _ := png.Decode(touxiang)
	defer touxiang.Close()

	//固定图片大小
	headImgDecode = resize.Resize(636, 775, headImgDecode, resize.Lanczos3)

	headBounds := headImgDecode.Bounds()
	headRgba := image.NewRGBA(headBounds)
	err = graphics.Scale(headRgba, img)
	return headRgba, nil
}

//图片灰化处理
func hdImageV3(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			_, g, _, a := colorRgb.RGBA()
			g_uint8 := uint8(g >> 8)
			a_uint8 := uint8(a >> 8)
			newRgba.SetRGBA(i, j, color.RGBA{g_uint8, g_uint8, g_uint8, a_uint8})
		}
	}
	return newRgba
}

//图片色彩反转
func fzImageV3(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			r_uint8 := uint8(r >> 8)
			g_uint8 := uint8(g >> 8)
			b_uint8 := uint8(b >> 8)
			a_uint8 := uint8(a >> 8)
			r_uint8 = 255 - r_uint8
			g_uint8 = 255 - g_uint8
			b_uint8 = 255 - b_uint8
			newRgba.SetRGBA(i, j, color.RGBA{r_uint8, g_uint8, b_uint8, a_uint8})
		}
	}
	return newRgba
}
