package main

import (
	"fmt"
	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/ChengjinWu/imagedraw"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	ZhengmianV3()
	//Fangmian()
}

/**
出生年月日 方正黑体简体
字符大小：姓名＋号码（11点）其他（9点）
字符间距（AV）：号码（50）
字符行距：住址（12点）
**/
func ZhengmianV3() {
	textBrush, err := imagedraw.NewTextBrush("../data/font/hei.ttf", 50, image.Black, 300)
	if err != nil {
		fmt.Println(err)
	}
	backgroundImg, err := imaging.Open("../data/idcardMerge/source/shenfenzheng_zhengmian.png")
	if err != nil {
		fmt.Println(err)
	}

	backgroundImgBounds := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(backgroundImgBounds, backgroundImg.Bounds(), backgroundImg, image.ZP, draw.Src)

	//姓名
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(358, 172), `周宏楷
`)

	//性别
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(358, 312), `男
`)
	//民族
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(735, 312), `汉
	`)

	//住址
	textBrush.TextWidth = 900
	textBrush.FontSize = 55
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(358, 593), `湖北省咸宁市咸安区永安

大道175号平房
`)

	textBrushBirthday, err := imagedraw.NewTextBrush("../data/font/fzhei.ttf", 50, image.Black, 300)
	if err != nil {
		fmt.Println(err)
	}
	//出生年
	textBrushBirthday.DrawFontOnRGBA(backgroundImgBounds, image.Pt(358, 452), `2000
`)

	//出生月
	textBrushBirthday.DrawFontOnRGBA(backgroundImgBounds, image.Pt(680, 452), `8
`)

	//出生日
	textBrushBirthday.DrawFontOnRGBA(backgroundImgBounds, image.Pt(860, 452), `17
`)

	//身份证号
	textBrushIdcard, errIdcard := imagedraw.NewTextBrush("../data/font/OCR-B 10 BT.TTF", 70, image.Black, 980)
	if errIdcard != nil {
		fmt.Println(errIdcard)
	}
	textBrushIdcard.DrawFontOnRGBA(backgroundImgBounds, image.Pt(658, 916), `421202200008170514
`)

	imgRGBA, err := GetHeadImageRGBAV3("M:/goProgram/foreground.png")
	if err != nil {
		fmt.Println(err)
	}

	x0, y0 := 1155, 80
	draw.DrawMask(backgroundImgBounds, image.Rect(x0, y0, x0+636, y0+775), imgRGBA, image.ZP, imgRGBA, image.ZP, draw.Over)

	//backgroundImgBounds=fzImageV3(backgroundImgBounds)
	imaging.Save(backgroundImgBounds, "../data/idcardMerge/result/zhengmian_big.jpg")

	//固定图片大小输出
	resultImgDecode := resize.Resize(660, 422, backgroundImgBounds, resize.Lanczos3)
	imaging.Save(resultImgDecode, "../data/idcardMerge/result/zhengmian_small.jpg")
}

func FangmianV3() {
	textBrush, err := imagedraw.NewTextBrush("M:/goProgram/src/home/idcardMerge/华文仿宋 加粗.TTF", 30, image.Black, 999)
	if err != nil {
		fmt.Println(err)
	}
	backgroundImg, err := imaging.Open("M:/goProgram/shenfenzheng_fanmian.png")
	if err != nil {
		fmt.Println(err)
	}

	backgroundImgBounds := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(backgroundImgBounds, backgroundImg.Bounds(), backgroundImg, image.ZP, draw.Src)

	//发证机关
	textBrush.DrawFontOnRGBA(backgroundImgBounds, image.Pt(370, 364), `成都市公安局天府新区分局
	`)

	textBrushY, err := imagedraw.NewTextBrush("M:/goProgram/src/home/idcardMerge/OCR-B 10 BT.TTF", 30, image.Black, 999)
	if err != nil {
		fmt.Println(err)
	}
	//有效日期
	textBrushY.DrawFontOnRGBA(backgroundImgBounds, image.Pt(370, 424), `2017.08.12-2037.08.12
	`)

	imaging.Save(backgroundImgBounds, "./../fangmian.jpg")
}

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
