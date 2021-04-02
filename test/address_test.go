package test

import (
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"testing"
)

func Test_Address(t *testing.T) {
	pub := "6e4f8120ba0ce9fa1d34bbb604860b6dee2873c0303a7aa6453902c6e9268462"
	address, err := ss58.EncodeByPubHex(pub, ss58.KsmPrefix)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(address)
}
