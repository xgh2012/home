package main

import (
	_ "fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/skip2/go-qrcode"
	"image/png"
	"log"
	"net/http"
	"os"

	//"github.com/tuotoo/qrcode"
	_ "image/png"
	_ "io/ioutil"
	_ "log"
	_ "net/http"
)

func Qrcode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	bt, _ := makacode(r.FormValue("txt"))
	w.Write(bt)
	//return
	//r.ParseForm()
	//
	//png,_:=qrcode.Encode("https://example.org", qrcode.Medium, 256)
	//fmt.Println(png)
	//w.Write(png)
	//
	////// Create the barcode
	//qrCode, _ := qr.Encode("http://app.longguanjia.so/bar/index.php?m=ajax&a=makeqrcode&barid=test0000000102&size=4&margin=2&isdown=1", qr.M, qr.Auto)
	//
	//// Scale the barcode to 200x200 pixels
	//qrCode, _ = barcode.Scale(qrCode, 200, 200)
	//
	//// create the output file
	//file, _ := os.Create("qrcode.png")
	//defer file.Close()
	//
	//// encode the barcode as png
	//png.Encode(file, qrCode)
}

func main() {
	//decodeQrcode()
	//return
	http.HandleFunc("/qrcode", Qrcode)
	log.Fatal(http.ListenAndServe("127.0.0.1:80", nil))
}

/**
*解码二维码
**/
func decodeQrcode() {
	//req,_ := http.Get("https://cli.clewm.net/file/2014/12/10/10febbdfabe543c7dd27d74fb4f411f3.png")
	//body,_:=ioutil.ReadAll(req.Body)
	//filename :="test.png"
	//err:=ioutil.WriteFile(filename, body,0644)
	//if err!=nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fi, err := os.Open(filename)
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//defer fi.Close()
	//qrmatrix, err := qrcode.Decode(fi)
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(qrmatrix.Content)
}

/**
*生成二维码
**/
func makacode(text string) ([]byte, error) {
	// Create the barcode
	qrCode, _ := qr.Encode(text, qr.M, qr.Auto)
	//
	//// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 256, 256)
	//
	//// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()
	//
	//// encode the barcode as png
	png.Encode(file, qrCode)

	return qrcode.Encode(text, qrcode.Medium, 256)
}
