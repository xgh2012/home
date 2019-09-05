package common

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

//AES ECB模式的加密解密
type AesTool struct {
	//128 192  256位的其中一个 长度 对应分别是 16 24  32字节长度
	Key       []byte
	BlockSize int
}

func NewAesTool(key []byte, blockSize int) *AesTool {
	return &AesTool{Key: key, BlockSize: blockSize}
}

func (this *AesTool) padding(src []byte) []byte {
	//填充个数
	paddingCount := aes.BlockSize - len(src)%aes.BlockSize
	if paddingCount == 0 {
		return src
	} else {
		//填充数据
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

//unpadding
func (this *AesTool) unPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return nil
}

func (this *AesTool) Encrypt(src []byte) ([]byte, error) {
	fmt.Println(this.Key)

	//key只能是 16 24 32长度
	block, err := aes.NewCipher([]byte(this.Key))
	if err != nil {
		return nil, err
	}
	//padding
	src = this.padding(src)
	//返回加密结果
	encryptData := make([]byte, len(src))
	//存储每次加密的数据
	tmpData := make([]byte, this.BlockSize)

	//分组分块加密
	for index := 0; index < len(src); index += this.BlockSize {
		block.Encrypt(tmpData, src[index:index+this.BlockSize])
		copy(encryptData, tmpData)
	}
	return encryptData, nil
}
func (this *AesTool) Decrypt(src []byte) ([]byte, error) {
	//key只能是 16 24 32长度
	block, err := aes.NewCipher([]byte(this.Key))
	if err != nil {
		return nil, err
	}
	//返回加密结果
	decryptData := make([]byte, len(src))
	//存储每次加密的数据
	tmpData := make([]byte, this.BlockSize)

	//分组分块加密
	for index := 0; index < len(src); index += this.BlockSize {
		block.Decrypt(tmpData, src[index:index+this.BlockSize])
		copy(decryptData, tmpData)
	}
	return this.unPadding(decryptData), nil
}

//测试padding  unpadding
func TestPadding() {
	tool := NewAesTool([]byte{}, 16)
	src := []byte{1, 2, 3, 4, 5}
	src = tool.padding(src)
	fmt.Println(src)
	src = tool.unPadding(src)
	fmt.Println(src)
}

//测试AES ECB 加密解密
func TestEncryptDecrypt() {
	key := []byte("{003FDED7-1D4D-}")
	blickSize := len(key)
	tool := NewAesTool(key, blickSize)
	encryptData, _ := tool.Encrypt([]byte("32334erew32"))
	fmt.Println(encryptData)
	encryptData, _ = base64.StdEncoding.DecodeString("Wdl7vH1I8EKFSIfHVAkiBQ==")
	decryptData, _ := tool.Decrypt([]byte(encryptData))
	fmt.Println(string(decryptData))
}

//加密数据
func AesEncrypt(str string) string {
	key := []byte("{003FDED7-1D4D-}")
	blickSize := len(key)
	tool := NewAesTool(key, blickSize)
	encryptData, _ := tool.Encrypt([]byte(str))
	res := base64.StdEncoding.EncodeToString(encryptData)
	return res
}

//解密数据
func AesDecrypt(str string) string {
	key := []byte("{003FDED7-1D4D-}")
	blickSize := len(key)
	tool := NewAesTool(key, blickSize)
	encryptData, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(str)
	decryptData, _ := tool.Decrypt([]byte(encryptData))
	return string(decryptData)
}
