package control

import (
	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"home/idcard/imagedraw"
	"image"
	"image/draw"
	"image/png"
	"os"
	"strings"
)

/**
出生年月日 方正黑体简体
字符大小：姓名＋号码（11点）其他（9点）
字符间距（AV）：号码（50）
字符行距：住址（12点）
**/
func GetZhengMian() (result string, message string) {
	textBrush, err := imagedraw.NewTextBrush(GetRealPath("data/font/黑体.ttf"), 50, image.Black, 300)
	if err != nil {
		return "GZ0001", "打开字体失败"
	}
	backgroundImg, err := imaging.Open(GetRealPath("data/idcard/source/shenfenzheng_zhengmian.png"))
	if err != nil {
		return "GZ0002", "打开正面背景失败"
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

	textBrushBirthday, err := imagedraw.NewTextBrush(GetRealPath("data/font/方正黑体简体_0.ttf"), 50, image.Black, 300)
	if err != nil {
		return "GZ0003", "打开字体失败"
	}

	Birthday := strings.Split(UserInfo.Birthday, "/")
	if len(Birthday) < 3 {
		return "GZ0004", "生日信息获取错误"
	}

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
	textBrushIdcard, err := imagedraw.NewTextBrush(GetRealPath("data/font/OCR-B 10 BT.TTF"), 70, image.Black, 980)
	if err != nil {
		return "GZ0005", "打开字体失败"
	}
	textBrushIdcard.DrawFontOnRGBA(backgroundImgBounds, image.Pt(658, 916), UserInfo.Idcard+`
`)

	imgRGBA, err := GetHeadImageRGBA(GetRealPath(UserInfo.ForegroundImg))
	if err != nil {
		return "GZ0006", "打开头像失败"
	}

	x0, y0 := 1155, 80
	draw.DrawMask(backgroundImgBounds, image.Rect(x0, y0, x0+636, y0+775), imgRGBA, image.ZP, imgRGBA, image.ZP, draw.Over)

	//backgroundImgBounds=fzImageV3(backgroundImgBounds)
	UserInfo.IdcardImgBigFront = "data/idcard/result/merge/zhengmian" + UserInfo.Idcard + "_big.jpg"
	UserInfo.IdcardImgSmallFront = "data/idcard/result/merge/zhengmian" + UserInfo.Idcard + "_small.jpg"

	err = imaging.Save(backgroundImgBounds, GetRealPath(UserInfo.IdcardImgBigFront))
	if err != nil {
		return "GZ0007", "保存身份证失败大图"
	}
	//固定图片大小输出
	resultImgDecode := resize.Resize(660, 422, backgroundImgBounds, resize.Lanczos3)
	err = imaging.Save(resultImgDecode, GetRealPath(UserInfo.IdcardImgSmallFront))
	if err != nil {
		return "GZ0008", "保存身份证失败小图"
	}
	return Success_Code, Success_Mes
}

//合成身份证反面
func GetFanMian() (result string, message string) {
	textBrush, err := imagedraw.NewTextBrush(GetRealPath("data/font/黑体.ttf"), 60, image.Black, 1280)
	if err != nil {
		return "GF0001", "打开字体失败"
	}
	backgroundImg, err := imaging.Open(GetRealPath("data/idcard/source/shenfenzheng_fanmian.png"))
	if err != nil {
		return "GF0002", "打开反面背景失败"
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

	UserInfo.IdcardImgBigBack = "data/idcard/result/merge/fanmian" + UserInfo.Idcard + "_big.jpg"
	UserInfo.IdcardImgSmallBack = "data/idcard/result/merge/fanmian" + UserInfo.Idcard + "_small.jpg"

	err = imaging.Save(backgroundImgBounds, GetRealPath(UserInfo.IdcardImgBigBack))
	if err != nil {
		return "GF0003", "保存身份证失败大图"
	}
	resultImgDecode := resize.Resize(660, 422, backgroundImgBounds, resize.Lanczos3)
	err = imaging.Save(resultImgDecode, GetRealPath(UserInfo.IdcardImgSmallBack))
	if err != nil {
		return "GF0004", "保存身份证失败小图"
	}
	return Success_Code, Success_Mes
}

func GetHeadImageRGBA(iamgePath string) (*image.RGBA, error) {
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
