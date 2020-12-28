package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/crypto"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/JFJun/stafi-substrate-go/client"
	"github.com/JFJun/stafi-substrate-go/expand"
	"github.com/JFJun/stafi-substrate-go/tx"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"testing"
)

func Test_tx(t *testing.T) {
	from := "eYGvmRP1fu418SSaypzN84S58YCm9c8SST9V2NriPbEMPP9"
	to := "gA7aiTz144UvgFtPEboN6JjoQgGGUiuM9Wx3NRzETw6gCW6"

	amount := uint64(50000000000000)
	c, err := client.New("wss://testnet.liebi.com")
	if err != nil {
		t.Fatal(err)
	}
	v, err := c.C.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v.TransactionVersion)
	fmt.Println(v.SpecVersion)
	acc, err := c.GetAccountInfo(from)
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.BifrostPrefix)
	nonce := uint64(acc.Nonce)
	fmt.Println(nonce)
	fmt.Println(c.GetGenesisHash())
	types.SetSerDeOptions(types.SerDeOptions{NoPalletIndices: true})
	transaction := tx.CreateTransaction(from, to, amount, nonce)
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(),
		c.GetGenesisHash())
	ed, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}
	btCall, err := ed.MV.GetCallIndex("Balances", "transfer")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(btCall)
	transaction.SetSpecVersionAndCallId(uint32(v.SpecVersion), uint32(v.TransactionVersion), btCall)
	tt, err := transaction.SignTransaction("000", crypto.Sr25519Type)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tt)
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", tt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_CreateUtilityBatch(t *testing.T) {
	from := "5DkswVFmWPUwPkmqMUEvavvso2HMdiyY71ixA2e52Ynwzvtg"
	to := "5H4N5JZHuqkprDKSR9SJeTMivbQQ94WrxeFELxh45ACoZFQC"
	nonce := uint64(16)
	//amount := uint64(123456)
	c, err := client.New("wss://api.crust.network/")
	if err != nil {
		t.Fatal(err)
	}
	v, err := c.C.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		t.Fatal(err)
	}
	acc, err := c.GetAccountInfo(from)
	if err != nil {
		t.Fatal(err)
	}
	nonce = uint64(acc.Nonce)
	pa := make(map[string]uint64)
	pa[to] = 100
	pa["5Hmy8BVAXAdaL6uxd41WJV4rhhWCNsXzekFRfuwLDkke9nG4"] = 1000000000
	types.SetSerDeOptions(types.SerDeOptions{NoPalletIndices: true})
	transaction := tx.CreateUtilityBatchTransaction(from, nonce, pa, "1100")
	transaction.SetGenesisHashAndBlockHash("0x34f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c65",
		"0x34f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c65")
	transaction.SetSpecVersionAndCallId(uint32(v.SpecVersion), uint32(v.TransactionVersion), "1700")
	tt, err := transaction.SignTransaction("00000", crypto.Sr25519Type)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tt)
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", tt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
	d, _ := json.Marshal(result)
	fmt.Println(string(d))

}

func Test_TxWithMemo(t *testing.T) {
	from := "5Fq9MpKxdjzCWEHHtqZ6rdYkKUtW4qwmJV4VHwKBan2hxRyL"
	to := "5E2dFRZoSbXE4at8QjHPxfx8eWA9mvLFbH64ZE3wTAsEwVFu"
	nonce := uint64(1)
	amount := uint64(123456)
	c, err := client.New("wss://testnet-1.chainx.org/ws")
	if err != nil {
		t.Fatal(err)
	}
	v, err := c.C.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		t.Fatal(err)
	}
	ed, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}
	ubCall, err := ed.MV.GetCallIndex("Utility", "batch")
	if err != nil {
		t.Fatal(err)
	}
	btCall, err := ed.MV.GetCallIndex("Balances", "transfer")
	if err != nil {
		t.Fatal(err)
	}
	srCall, err := ed.MV.GetCallIndex("System", "remark")
	transaction := tx.CreateTransactionWithMemo(from, to, "test", amount, nonce, ubCall, srCall)
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(),
		c.GetGenesisHash())
	transaction.SetSpecVersionAndCallId(uint32(v.SpecVersion), uint32(v.TransactionVersion), btCall)
	tt, err := transaction.SignTransaction("", crypto.Sr25519Type)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tt)
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", tt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
	d, _ := json.Marshal(result)
	fmt.Println(string(d))
}

func Test_FakeDeposit(t *testing.T) {
	from := "5RHGBggpMDvQ9HtjjFpK7oETuc51DAj4QcATk6HH2dswNZ98"
	to := "5U2RLJHbQ1VQJ1bLACxYMtxSQVxRPqfp34WAUxfb5zpUEwaQ"
	nonce := uint64(15)
	amount := uint64(123456)
	c, err := client.New("")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.ChainXPrefix)
	v, err := c.C.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		t.Fatal(err)
	}
	ed, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}
	callIdx, err := ed.MV.GetCallIndex("Balances", "transfer")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		if i == 0 {
			nonce = 5
			amount = 1000000
		} else {
			nonce = 6
			amount = 7000000
		}
		transaction := tx.CreateTransaction(from, to, amount, nonce)
		transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(),
			c.GetGenesisHash())
		transaction.SetSpecVersionAndCallId(uint32(v.SpecVersion), uint32(v.TransactionVersion), callIdx)
		tt, err := transaction.SignTransaction("", crypto.Sr25519Type)
		if err != nil {
			t.Fatal(err)
		}
		var result interface{}
		err = c.C.Client.Call(&result, "author_submitExtrinsic", tt)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(result)
	}
}

// 0x390284ff 8ce4f854296af0a2fa35faaf6f6577fb46d63d833d1a24a219d604506d151328 01 56e907f93a77685966e939d84c93366ff9895087c2d7f4cbe39ebaf750fe114004e7d01e836d4893d2c7cb8728915fb3d9908195ab70c8e1c53bcce25f52d48a 00  4102 00 0500ffc4b1c12fd91e7c199b4a3da3a3adee7bfd97f35dee81d58a670de7b294a7fa7402890700
// 0x3102 84  8ce4f854296af0a2fa35faaf6f6577fb46d63d833d1a24a219d604506d151328 01 260d90add44e26c8c290314711f30022fc33c7e2c626e8fabbfac7ff16ea04585e7e69f420d2a3dd7adf0147cb9fbaf41838708f5e78bbdf3a3eb649f834de8a 00 4102 00  0500c4b1c12fd91e7c199b4a3da3a3adee7bfd97f35dee81d58a670de7b294a7fa7402890700

// 0x350284ff4adffe0994aac9e292470b27eac94e505532ac1a22ae17012ddad445e6b78019018638639fb20aab85c72d5078439e6534a7b49039c4499ecb7d416d08792e7a4dbc4bc60db33a4681303a728ab3370a9e0d960ef4f60feef7c1eb391fc3538084003c001700ffdcea9317bceb28b52bdae9229a3794de4ca85e36d990a78f779c6fd7f27eb54102890700
// 0x350284ff4adffe0994aac9e292470b27eac94e505532ac1a22ae17012ddad445e6b78019019c7d33500b6cf5bf2da5291c948d1a333766155b570e745a31647f589bafa50b8f62048ec7d4196f1242482944e5f64d0f829d0a2ffbfa7e9f308eacc44a538d003c001700ffdcea9317bceb28b52bdae9229a3794de4ca85e36d990a78f779c6fd7f27eb54102890700

// 0x 2d02 84 8ce4f854296af0a2fa35faaf6f6577fb46d63d833d1a24a219d604506d151328 01
//			  da7245068281d7bd5e1e3db63b3b1d8c43664f5393687f663d7d80c9515f590f
//			  af5dd9be85eaebb83de8cbb53d368a77e7bd3b0a139602ddafaf0e1736e4038d
//			  000c000500c4b1c12fd91e7c199b4a3da3a3adee7bfd97f35dee81d58a670de7b294a7fa7402890700
// 0x 2d02 84 8ce4f854296af0a2fa35faaf6f6577fb46d63d833d1a24a219d604506d151328 01
//			  288e08acd89507664ca4cd2684157921a7a3df810e9bbda30669edb63d7bcc79
//			  505d4e87b5d525cb0b8f1c497b0fb9aa36776ee9fc1315188119f309a821cb8f
//			  000c000500c4b1c12fd91e7c199b4a3da3a3adee7bfd97f35dee81d58a670de7b294a7fa7402890700
