package adapt

import (
	"encoding/binary"
)

/**
Python define：MESSAGE_HEAD_FMT = "<1s1sI2B2I20s8s"
Length: 1 + 1 + 4 + 2*1 + 2*4 + 20 + 8 = 44
*/

type Header struct {
	szHeaderFlag  [2]uint8  //包头起始标志，固定为“FT”
	nProtoID      uint32    //协议 ID
	nProtoFmtType uint8     //协议格式类型，0 为 Protobuf 格式，1 为 Json 格式
	nProtoVer     uint8     //协议版本，用于迭代兼容，目前填 0
	nSerialNo     uint32    //包序列号，用于对应请求包和回包，要求递增
	nBodyLen      uint32    //包体长度
	arrBodySHA1   [20]uint8 //包体原始数据(解密后)的 SHA1 哈希值
	arrReserved   [8]uint8  //保留 8 字节扩展
}

const HEADER_LENGTH = 44

func (h *Header) GetProtoID() uint32 {
	return h.nProtoID
}

func (h *Header) SetProtoID(protoID uint32) {
	h.nProtoID = protoID
}

func (h *Header) GetProtoFmtType() uint8 {
	return h.nProtoFmtType
}

func (h *Header) SetProtoFmtType(protoFmtType uint8) {
	h.nProtoFmtType = protoFmtType
}

func (h *Header) GetProtoVer() uint8 {
	return h.nProtoVer
}

func (h *Header) SetProtoVer(protoVer uint8) {
	h.nProtoVer = protoVer
}

func (h *Header) GetSeriaNo() uint32 {
	return h.nSerialNo
}

func (h *Header) SetSeriaNo(seriaNo uint32) {
	h.nSerialNo = seriaNo
}

func (h *Header) GetBodyLen() uint32 {
	return h.nBodyLen
}

func (h *Header) SetBodyLen(bodyLen uint32) {
	h.nBodyLen = bodyLen
}

func (h *Header) GetBodySHA1() [20]uint8 {
	return h.arrBodySHA1
}

func (h *Header) SetBodySHA1(bodySHA1 []uint8) {
	for i := 0; i < len(bodySHA1) && i < len(h.arrBodySHA1); i++ {
		h.arrBodySHA1[i] = bodySHA1[i]
	}
}

func (h *Header) GetReserved() [8]uint8 {
	return h.arrReserved
}

func (h *Header) SetReserved(reserved []uint8) {
	for i := 0; i < len(reserved) && i < len(h.arrReserved); i++ {
		h.arrReserved[i] = reserved[i]
	}
}

func (h *Header) Pack() []byte {
	ret := make([]byte, 2, HEADER_LENGTH)
	start := 0
	ret[start] = byte('F')
	start++
	ret[start] = byte('T')
	start++
	v := make([]byte, 4)
	binary.LittleEndian.PutUint32(v, h.nProtoID)
	ret = append(ret, v...)
	start += 4
	ret = append(ret, byte(h.nProtoFmtType))
	start++
	ret = append(ret, byte(h.nProtoVer))
	start++
	binary.LittleEndian.PutUint32(v, h.nSerialNo)
	ret = append(ret, v...)
	start += 4
	binary.LittleEndian.PutUint32(v, h.nBodyLen)
	ret = append(ret, v...)
	start += 4
	ret = append(ret, []byte(h.arrBodySHA1[:])...)
	ret = append(ret, []byte(h.arrReserved[:])...)
	return ret
}

func (h *Header) UnPack(data []byte) {
	if len(data) < HEADER_LENGTH {
		return
	}
	start := 0
	h.szHeaderFlag[0] = uint8(data[start])
	start++
	h.szHeaderFlag[1] = uint8(data[start])
	start++
	h.nProtoID = binary.LittleEndian.Uint32(data[start : start+4])
	start += 4
	h.nProtoFmtType = uint8(data[start])
	start += 1
	h.nProtoVer = uint8(data[start])
	start += 1
	h.nSerialNo = binary.LittleEndian.Uint32(data[start : start+4])
	start += 4
	h.nBodyLen = binary.LittleEndian.Uint32(data[start : start+4])
	start += 4
	var i int
	for i = 0; i < len(h.arrBodySHA1); i++ {
		h.arrBodySHA1[i] = uint8(data[start])
		start++
	}
	for i = 0; i < len(h.arrReserved); i++ {
		h.arrReserved[i] = uint8(data[start])
		start++
	}
}
