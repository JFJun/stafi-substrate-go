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

	from := "5DkswVFmWPUwPkmqMUEvavvso2HMdiyY71ixA2e52Ynwzvtg"
	to := "5H4N5JZHuqkprDKSR9SJeTMivbQQ94WrxeFELxh45ACoZFQC"

	amount := uint64(1234567890000)
	c, err := client.New("wss://node-6714447553211260928.rz.onfinality.io/ws")
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
	nonce := uint64(acc.Nonce)
	//types.SetSerDeOptions(types.SerDeOptions{NoPalletIndices: true})
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
	from := "32qwhN8jf2nqa8nh6rjLhbfJ8kRjb1TgYFeLQyvRT8yPcNFr"
	to := "3441UQVQ9W172grCEXCKo66faJcxPMbnDTTvFxUwD2x3LtiY"
	nonce := uint64(15)
	amount := uint64(123456)
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.StafiPrefix)
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
			nonce = 13
			amount = 1000000000000
		} else {
			nonce = 14
			amount = 2000000000000
		}
		transaction := tx.CreateTransaction(from, to, amount, nonce)
		transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(),
			c.GetGenesisHash())
		transaction.SetSpecVersionAndCallId(uint32(v.SpecVersion), uint32(v.TransactionVersion), callIdx)
		tt, err := transaction.SignTransaction("0000", crypto.Sr25519Type)
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
