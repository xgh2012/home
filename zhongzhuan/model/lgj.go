package model

import (
	"bytes"
	"encoding/binary"
	"home/zhongzhuan/common"
	"log"
	"math/rand"
	"strconv"
)

//包头总长 71
type LgjHeaderStr struct {
	FSign     []byte //用于区分错包和乱包 Array [0..7] of Char
	iCMD      []byte //包头类型 short 2
	iSubCmd   byte   //包头类型子类型（用于扩展） byte 1
	iVer      []byte //协议版本标识 0：包体用INI方式分割。1：包体用json方式分割。 Short 2
	iCheckSum byte   //校验和（效验以下所有的值）。 byte 1
	ReCMD     byte   //成功标识(0 成功，1失败)   byte 1
	SBarID    []byte //网吧编号   Array [0..19] of Char
	IEnc      []byte //动态加密号   Integer 4
	iMax      []byte //最大包数，从1开始计算   Integer 4
	iCur      []byte //当前包序号，从 1开始小于等于iMax   Integer 4
	FuniqueID []byte //通信唯一识别符   TGuid 16
	Socketfd  []byte //这个只是中转时填入，中转连接的两方都不原样返回。初始创建包时填0即可    Integer 4
	iDataLen  []byte //包体长度    Integer 4
	sendData  string //包体数据
}

var (
	LHeader     LgjHeaderStr
	tmpBodyData [][]byte
)

//包头处理
/**
*可变参数
* icmd iVer SBarID
 */
func LgjBodys(params common.Params) []byte {
	var LgjAllBody []byte
	//FSign 处理 0 + 8 = 8
	LgjAllBody = byteMerge(LgjAllBody, []byte("XAPPCODE"))

	//iCMD 处理 8+2=10
	cmd, _ := strconv.ParseInt(params.Icmd, 0, 32)
	LgjAllBody = byteMerge(LgjAllBody, IntToByte(cmd)[0:2])

	//iSubCmd 处理 10+1=11
	LgjAllBody = byteMerge(LgjAllBody, []byte{byte(0)})

	//iVer 处理 11+2=13
	iver, _ := strconv.ParseInt(params.IVer, 0, 32)
	LgjAllBody = byteMerge(LgjAllBody, IntToByte(iver)[0:2])

	//iCheckSum校验和（效验以下所有的值） 13+1=14 暂时用0占位置 在body 确认后再重新赋值
	LgjAllBody = append(LgjAllBody, byte(0))

	//ReCMD 处理 14+1 = 15
	LgjAllBody = append(LgjAllBody, byte(0))

	//SBarID 处理 15+20=35
	var barid = make([]byte, 20)
	tmpbarid := []byte(params.Barid)
	copy(barid, tmpbarid)
	LgjAllBody = byteMerge(LgjAllBody, barid)

	//IEnc 处理 35+4=39
	LgjAllBody = byteMerge(LgjAllBody, IntToByte(int64(0))[0:4])

	//iMax 最大包数 处理 39+4 = 43
	LgjAllBody = byteMerge(LgjAllBody, IntToByte(int64(1))[0:4])

	//iCur 当前包序号 处理 43+4 = 47
	LgjAllBody = byteMerge(LgjAllBody, IntToByte(int64(1))[0:4])

	//FuniqueID 47+16=63
	var funiqueid []byte
	for idx := 0; idx < 16; idx++ {
		funiqueid = append(funiqueid, byte(rand.Int()))
	}
	LgjAllBody = byteMerge(LgjAllBody, funiqueid)

	//Socketfd 处理 63+4=67
	LgjAllBody = byteMerge(LgjAllBody, IntToByte(int64(0))[0:4])

	//包体处理
	body := []byte(params.Data)

	//iDataLen 处理 包体长度 67+4=71
	LgjAllBody = byteMerge(LgjAllBody, IntToByte(int64(len(body)))[0:4])

	LgjAllBody = byteMerge(LgjAllBody, body)

	//iCheckSum校验和 处理
	LgjAllBody[13] = CheckSum(LgjAllBody[14:])

	return LgjAllBody
}

//byte 数组合并
func byteMerge(srcData []byte, distData []byte) (result []byte) {
	tmpBodyData = [][]byte{srcData, distData}
	result = bytes.Join(tmpBodyData, []byte{})
	return result
}

//对结果进行处理
func LgjDescResult(count int, result []byte) (bussResult string) {
	//计算校验值
	if CheckSum(result[14:]) != result[13] {
		log.Fatal("校验值不通过")
	}
	bussResult = string(result[71:])
	return bussResult
}

//IntToByte 实现 使用小端
func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.LittleEndian, num)
	if err != nil {

	}
	return buffer.Bytes()
}

//计算校验和
func CheckSum(datas []byte) uint8 {
	var sum uint8
	lenth := len(datas)
	//不用再除 256
	for index := 0; index < lenth; index++ {
		sum += datas[index]
	}
	return sum
}
