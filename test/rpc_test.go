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
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.PolkadotPrefix)
	ai, err := c.GetAccountInfo("1exaAg2VJRQbyUBAeXcktChCAqjVP9TUxF3zo23R2T6EGdE")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(uint32(ai.Nonce))
}

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New("wss://mainnet-rpc.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	//types.SetSerDeOptions(types.SerDeOptions{NoPalletIndices: true})
	c.SetPrefix(ss58.StafiPrefix)
	resp, err := c.GetBlockByNumber(987114)
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
