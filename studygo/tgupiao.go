package main

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

type Signer struct {
	FontSize   float64
	Dpi        float64
	font       *truetype.Font
	startPoint image.Point
	signPoint  image.Point
}

func main() {
	sign := NewSigner(72, "M:/goProgram/src/github.com/golang/freetype/testdata/luxisr.ttf", 240, 350, 360)
	img, _, err := sign.DrawStringImage("sdsdfd")
	if err != nil {
		log.Fatalln(err)
		return
	}

	outFile, err := os.Create("out_test_1.png")
	if err != nil {
		log.Println("open 10png ", err)
		os.Exit(1)
	}

	defer outFile.Close()

	var b bytes.Buffer

	//b := bufio.NewWriter(outFile)
	err = png.Encode(&b, img)

	if err != nil {
		log.Println(" write file ", err)
		os.Exit(1)
	}

	outFile.Write(b.Bytes())

	fmt.Printf("ok \n")

	fmt.Println("Wrote out.png OK.")
}

//字符初始化
func initFont(fontfile string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	return font, nil
}

func NewSigner(dpi float64, fontfile string, size float64, x, y int) *Signer {
	font, err := initFont(fontfile)
	if err != nil {
		return nil
	}
	return &Signer{
		FontSize:   size,
		font:       font,
		Dpi:        dpi,
		startPoint: image.ZP,
		signPoint:  image.Point{X: x, Y: y},
	}
}

//code码的写入
func (this *Signer) DrawStringCode(scode string) (*image.RGBA, fixed.Point26_6, error) {
	var err error
	var my_fix fixed.Point26_6
	fgn := image.NewUniform(color.Gray16{0x9cd3})
	fgt := image.NewUniform(color.Gray16{0x4a48})
	fg, bg := fgn, image.Transparent

	rgba := image.NewRGBA(image.Rect(0, 0, this.signPoint.X, this.signPoint.Y))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	//设置文字对像
	c := freetype.NewContext()
	c.SetDPI(this.Dpi)           //分辨率
	c.SetFont(this.font)         //字符
	c.SetFontSize(this.FontSize) //大小
	c.SetClip(rgba.Bounds())     //背景
	c.SetDst(rgba)               //
	c.SetSrc(fg)
	//c.SetHinting(font.HintingNone)

	// Draw the text.
	//计算出x的位置
	pt := freetype.Pt(0, int(c.PointToFixed(this.FontSize)>>6))

	c.SetSrc(fgt)
	my_fix, err = c.DrawString("yqm：", pt)
	if err != nil {
		return rgba, my_fix, err
	}

	pt.X += c.PointToFixed(90)
	c.SetSrc(image.Black)
	my_fix, err = c.DrawString(scode, pt)
	//fmt.Println(my_fix)
	if err != nil {
		return rgba, my_fix, err
	}
	return rgba, my_fix, nil
}

//画一个带有text的图片
func (this *Signer) DrawStringImage(scode string) (*image.RGBA, fixed.Point26_6, error) {
	var err error
	var fixd_my fixed.Point26_6
	fgn := image.NewUniform(color.Gray16{0x9cd3})
	fgt := image.NewUniform(color.Gray16{0x4a48})
	fg, bg := fgn, image.Transparent

	rgba := image.NewRGBA(image.Rect(0, 0, this.signPoint.X, this.signPoint.Y))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	//设置文字对像
	c := freetype.NewContext()
	c.SetDPI(this.Dpi)       //分辨率
	c.SetFont(this.font)     //字符
	c.SetFontSize(36)        //大小
	c.SetClip(rgba.Bounds()) //背景
	c.SetDst(rgba)           //
	c.SetSrc(fg)
	//c.SetHinting(font.HintingNone)

	// Draw the text.
	//计算出x的位置
	c.SetSrc(fgt)
	//c.SetSrc(image.Black)
	pt := freetype.Pt(0, int(c.PointToFixed(this.FontSize)>>6))

	fixd_my, err = c.DrawString(fmt.Sprintf("I'am %s", scode), pt)
	if err != nil {
		return rgba, fixd_my, err
	}
	//fmt.Println(fixd_my)
	//err = png.Encode(b, rgba)
	return rgba, fixd_my, nil
}
