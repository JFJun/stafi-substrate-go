package test

import (
	"bytes"
	"fmt"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"testing"
)

func Test_decode(t *testing.T) {
	fmt.Println(len("60020000000000000100000080b8dc3c030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"))
	d := "80b8dc3c030000000000000000000000"
	decoder := scale.NewDecoder(bytes.NewReader(types.MustHexDecodeString(d)))
	var target types.U128
	err := decoder.Decode(&target)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(target)

}

// 60020000
//00000000
//01000000
//80b8dc3c030000000000000000000000
//00000000000000000000000000000000
//00000000000000000000000000000000
//00000000000000000000000000000000
