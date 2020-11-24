package test

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_decode(t *testing.T) {
	//d := "fdff24a8131000000000000000000000"
	//decoder := scale.NewDecoder(bytes.NewReader(types.MustHexDecodeString(d)))
	//var target types.U128
	//err := decoder.Decode(&target)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(target)
	s := "74657374"
	d, _ := hex.DecodeString(s)
	fmt.Println(string(d))
}
