package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"math/big"
	"strings"
)

func RemoveHex0x(hexStr string)string{
	if	strings.HasPrefix(hexStr,"0x"){
		return hexStr[2:]
	}
	return hexStr
}

func BytesToHex(b []byte) string {
	c := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(c, b)
	return string(c)
}

func U256(v string) *big.Int {
	v = strings.TrimPrefix(v, "0x")
	bn := new(big.Int)
	n, _ := bn.SetString(v, 16)
	return n
}

func UCompactToBigInt(u types.UCompact)*big.Int{
	b:=big.Int(u)
	return &b
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsNumberString(str string)bool{
	for _,a:=range str{
		if a >57 || a<48 {
			return false
		}
	}
	return true
}
func IntToHex(i interface{}) string {
	return fmt.Sprintf("%x", i)
}

func CheckStructData(object interface{}){
	d,_:=json.Marshal(object)
	fmt.Println(string(d))
}

func AddressToPublicKey(address string) string {
	if address == "" {
		return ""
	}
	pub, err := ss58.DecodeToPub(address)

	if err != nil {
		return ""
	}
	if len(pub) != 32 {
		return ""
	}
	pubHex := hex.EncodeToString(pub)
	return pubHex
}

func Remove0X(hexData string) string {
	if strings.HasPrefix(hexData, "0x") {
		return hexData[2:]
	}
	return hexData
}

func ZeroBytes(data []byte){
	for i,_:=range data{
		data[i] = 0
	}
}

