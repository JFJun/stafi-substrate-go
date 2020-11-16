package test

import (
	"bytes"
	"fmt"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"testing"
)

func Test_decode(t *testing.T) {
	d := "fdff24a8131000000000000000000000"
	decoder := scale.NewDecoder(bytes.NewReader(types.MustHexDecodeString(d)))
	var target types.U128
	err := decoder.Decode(&target)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(target)
}
