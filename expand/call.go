package expand

import (
	"encoding/hex"
	"fmt"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

func NewCall(callIdx string,args ...interface{})(types.Call,error){
	if len(callIdx)!=4 {
		return types.Call{},fmt.Errorf("callIdx length is not equal 4,len=%d",len(callIdx))
	}
	m,err:=hex.DecodeString(callIdx[:2])
	if err != nil {
		return types.Call{}, err
	}
	n,err:=hex.DecodeString(callIdx[2:])
	if err != nil {
		return types.Call{},err
	}
	c:=types.CallIndex{SectionIndex: m[0],MethodIndex: n[0]}
	var a []byte
	for _, arg := range args {
		e, err := types.EncodeToBytes(arg)
		if err != nil {
			return types.Call{}, err
		}
		a = append(a, e...)
	}
	return types.Call{c, a}, nil
}
