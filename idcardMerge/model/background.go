package model

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

/**
*获取身份证正面背景
**/
func GetFront() (frontImgDecode image.Image,frontBounds image.Rectangle,frontRgba *image.RGBA) {
	zhengmian, _ := os.Open("M:/goProgram/image.png")
	frontImgDecode, _ = png.Decode(zhengmian)
	defer zhengmian.Close()

	frontBounds = frontImgDecode.Bounds()
	frontRgba = image.NewRGBA(frontBounds)
	return  frontImgDecode,frontBounds,frontRgba
}

func GetHeadImg()(headImgDecode image.Image,headBounds image.Rectangle,headRgba *image.RGBA)  {
	touxiang, _ := os.Open("M:/goProgram/watermark.jpg")
	headImgDecode, _ = jpeg.Decode( touxiang)
	defer  touxiang.Close()

	//固定图片大小
	headImgDecode=resize.Resize(206, 241, headImgDecode, resize.Lanczos3)

	headBounds = headImgDecode.Bounds()
	headRgba = image.NewRGBA(headBounds)
	return  headImgDecode,headBounds,headRgba
}