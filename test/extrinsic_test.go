package test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/JFJun/stafi-substrate-go/client"
	"github.com/JFJun/stafi-substrate-go/expand"
	"github.com/JFJun/stafi-substrate-go/utils"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"testing"
)

func Test_ExtrinsicDecode(t *testing.T) {
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	extrinsic := "450284ffe2d53f597d6e05231b2791346473e8ae6ffa6936f0db3ff74691ca19918dbd230170fd81ee207df240c09be80bbd07834dfbe056f242a82af0129cfc436c0f2521f8a91aefd5242d3d1405ceb6e7c58c27bbb104a41f3679b4b025c0de1ee3488f150010000600ffb001b764098de6caf2de4fc7bdf2c30dc6306e405d52b620112dff0eb6a823670b87689703ab4f"
	data, _ := hex.DecodeString(extrinsic)
	decoder := scale.NewDecoder(bytes.NewReader(data))
	meta, err := c.C.RPC.State.GetMetadataLatest()
	if err != nil {
		t.Fatal(err)
	}
	ed, err := expand.NewExtrinsicDecoder(meta)
	if err != nil {
		t.Fatal(err)
	}
	err = ed.ProcessExtrinsicDecoder(*decoder)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ed.Value)
	d, _ := json.Marshal(ed.Value)
	fmt.Println(string(d))
}
func Test_Compact(t *testing.T) {
	d := "001a"
	data, _ := hex.DecodeString(d)
	decoder := scale.NewDecoder(bytes.NewReader(data))
	var u types.UCompact
	err := decoder.Decode(&u)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(utils.UCompactToBigInt(u).Int64())
	b, err := decoder.ReadOneByte()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)

}

// 5502841cf326c5aaa5af9f0e2791e66310fe8f044faadaf12567eaa0976959d1f7731f01946d2b5a9f138a76c96fc6e34358b4236e097cf9c46e44ac05874f195a78ea133f211b89f2b531fd0982cd4e97e63d0c44a2cafa1d89d9c972badcc2b5122889
//
//f502
//ce160200 001a 0004 0500 e464f6458fb3d4b5e040fa9acda4cab9972a38281342e33462546580343daa6a0b00fe4b811704

//0x5102841cf326c5aaa5af9f0e2791e66310fe8f044faadaf12567eaa0976959d1f7731f015cb2e117826ad322bb7c4e84ac9fb6b5eebf6077fdb0f2d5dfc6ac753f25f44aee92da074ef4d648e0f323a1da4569adf2dee42b2181a1367b9fe4d3cb24a481b50276170200
//001a
// 0004 0500 a4622c0902547c3abd943d50ad4942f787c9c9cbfbb534a51ae6deba7bae0b1c07004894294f

//0x3d028456bab61e017ef136908f318444ce2c73603fcd6cd985c91dc5ae019c0a6ad706013c21dc38c5df6f17da9922a0115d202a40f77572e790827dfc0aae702206ee26f543fdf2e79ff743c3493fbdb792518cf71c56f67120953b66aec1961402f787f5000000
// 0500 56daf75bd0f09c11d798263bc79baeb77c4b4af1dbd372bbe532b1f8702b2a7e0bc0d01e3b0402
