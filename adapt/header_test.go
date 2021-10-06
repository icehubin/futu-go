package adapt

import (
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	headP := &Header{}
	headP.nProtoID = 1001
	headP.nProtoFmtType = 0
	headP.nProtoVer = 0
	headP.nSerialNo = 23423532
	headP.nBodyLen = 12345

	fmt.Println(headP)
	packData := headP.Pack()
	fmt.Println(packData)

	headP2 := &Header{}
	fmt.Println(headP2)
	headP2.UnPack(packData)
	fmt.Println(headP2)

}
