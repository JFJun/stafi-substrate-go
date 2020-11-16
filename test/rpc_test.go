package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/JFJun/stafi-substrate-go/client"
	"github.com/JFJun/stafi-substrate-go/expand"
	"github.com/JFJun/stafi-substrate-go/models"
	"testing"
)

func Test_RpcClient(t *testing.T) {
	c, err := client.New("wss://mainnet-rpc.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	var block models.SignedBlock
	err = c.C.Client.Call(&block, "chain_getBlock", "0xcbf137e5da22eed249580356c3c5cd074ad0e2d7621906ae4a30687a14f2da4c")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
	h, err := c.C.RPC.Chain.GetFinalizedHead()
	fmt.Println(h.Hex())

}

func Test_Zstring(t *testing.T) {
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	meta, err := c.C.RPC.State.GetMetadataLatest()
	if err != nil {
		t.Fatal(err)
	}
	//callIdx,err:=meta.FindCallIndex("Utility.batch")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(callIdx)
	d, _ := json.Marshal(meta.AsMetadataV12.Modules)
	fmt.Println(string(d))
	me, err := expand.NewMetadataExpand(meta)
	if err != nil {
		t.Fatal(err)
	}
	callIdx, err := me.MV.GetCallIndex("Utility", "batch")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(callIdx)
	m, cc, err := me.MV.FindNameByCallIndex(callIdx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(m, cc)
}

func Test_GetAccountInfo(t *testing.T) {
	c, err := client.New("wss://mainnet-rpc.stafi.io")
	//c, err := client.New("wss://rpc.polkadot.io")
	c.SetPrefix(ss58.StafiPrefix)
	if err != nil {
		t.Fatal(err)
	}
	//c.SetPrefix(ss58.StafiPrefix)
	ai, err := c.GetAccountInfo("34mqJ3JebpbRfudQFH2uDKXNsKoNUX4AWuoTgJhronNcENCh")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(ai)
	fmt.Println(string(d))
	//fmt.Println(uint32(ai.Nonce))
	//address:="32ZWhveKAYJp1CKbP7TZQUTqtcdDdGXcEYnsfDwdZ6Y3qMB3"
	////address = "12KkURmLnQcQQRvVNm5cj5uaBtUL5LVQBXUwmQ6oBHBMLwkG"
	//pub, err := ss58.DecodeToPub(address)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//storage, err := types.CreateStorageKey(c.Meta, "System", "Account", pub, nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//key:=storage.Hex()
	//var res interface{}
	//err = c.C.Client.Call(&res,"state_getStorageAt", key)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(res.(string))
	//decoder:=scale.NewDecoder(bytes.NewReader(types.MustHexDecodeString(res.(string))))
	//var ai types.AccountInfo
	//err = decoder.Decode(&ai)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//d,_:=json.Marshal(ai)
	//fmt.Println(string(d))
	//fis: 0x 30000000 00 fdff24a8131000000000000000000000 00000000000000000000000000000000 00000000000000000000000000000000 00000000000000000000000000000000
	//dot: 0x 12000000 02000000 f571d256491b000000000000000000000 0000000000000000000000000000000 00e057eb481b00000000000000000000 00e057eb481b00000000000000000000
}

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New("wss://mainnet-rpc.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	//types.SetSerDeOptions(types.SerDeOptions{NoPalletIndices: true})
	c.SetPrefix(ss58.StafiPrefix)
	resp, err := c.GetBlockByNumber(995445)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}

func Test_GetGenesisHash(t *testing.T) {
	c, err := client.New("wss://mainnet-rpc.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(c.GetGenesisHash())
}

// 1700dcea9317bceb28b52bdae9229a3794de4ca85e36d990a78f779c6fd7f27eb54102890700003c001c0000000300000034f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c6534f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c65
// 1700dcea9317bceb28b52bdae9229a3794de4ca85e36d990a78f779c6fd7f27eb54102890700003c001c0000000300000034f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c6534f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c65
