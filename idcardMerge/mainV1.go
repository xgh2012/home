package main

import (
	"home/idcardMerge/model"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

var (
	frontImgDecode image.Image
	frontBounds    image.Rectangle
	frontRgba      *image.RGBA
)

func main() {

	//获取背景 正面
	frontImgDecode, frontBounds, frontRgba = model.GetFront()
	draw.Draw(frontRgba, frontBounds, frontImgDecode, image.ZP, draw.Src)

	//头像处理
	idcardHead()

	//灰化
	//frontRgba=hdImage(frontRgba)

	//输出到文件
	imgw, _ := os.Create("M:/goProgram/watermarked.jpg")
	jpeg.Encode(imgw, frontRgba, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()
}

//头像处理
func idcardHead() {
	//头像处理---------------开始
	//调整头像位置
	headOffset := image.Pt(413, 72)
	//获取 正面 头像
	headImgDecode, headBounds, _ := model.GetHeadImg()

	draw.Draw(frontRgba, headBounds.Add(headOffset), headImgDecode, image.ZP, draw.Over)
	//头像处理---------------结束
}

//图片灰化处理
func hdImage(m image.Image) *image.RGBA {
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
func fzImage(m image.Image) *image.RGBA {
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
