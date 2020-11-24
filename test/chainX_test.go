package test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/crypto"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/JFJun/stafi-substrate-go/client"
	"github.com/JFJun/stafi-substrate-go/expand"
	"github.com/JFJun/stafi-substrate-go/models"
	"github.com/JFJun/stafi-substrate-go/tx"
	"testing"
)

func Test_Chain_GetBlockByNumber(t *testing.T) {
	c, err := client.New("wss://rpc.kulupu.corepaper.org/ws")
	if err != nil {
		t.Fatal(err)
	}
	dd, _ := json.Marshal(c.Meta.AsMetadataV12.Modules)
	fmt.Println(string(dd))
	fmt.Println(c.ChainName)
	for _, mod := range c.Meta.AsMetadataV12.Modules {
		if mod.HasEvents {
			for _, event := range mod.Events {
				fmt.Printf("%s_%s\n", mod.Name, event.Name)
				fmt.Println(event.Args)
				fmt.Println("------------------------------------------------")
			}
		}
	}
	c.SetPrefix(ss58.SubstratePrefix)
	block, err := c.GetBlockByNumber(60165)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
}

func Test_ChainXTX(t *testing.T) {
	from := "5Fq9MpKxdjzCWEHHtqZ6rdYkKUtW4qwmJV4VHwKBan2hxRyL"
	to := "5E2dFRZoSbXE4at8QjHPxfx8eWA9mvLFbH64ZE3wTAsEwVFu"
	nonce := uint64(0)
	amount := uint64(123456)
	c, err := client.New("wss://testnet-1.chainx.org/ws")
	if err != nil {
		t.Fatal(err)
	}
	v, err := c.C.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v.TransactionVersion)
	fmt.Println(v.SpecVersion)
	//meta,err:=c.C.RPC.State.GetMetadataLatest()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//types.SerDeOptionsFromMetadata(meta)

	//types.SetSerDeOptions(types.SerDeOptions{NoPalletIndices: true})
	transaction := tx.CreateTransaction(from, to, amount, nonce)
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(),
		c.GetGenesisHash())
	ed, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}
	callIdx, err := ed.MV.GetCallIndex("Balances", "transfer")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(callIdx)
	transaction.SetSpecVersionAndCallId(uint32(v.SpecVersion), uint32(v.TransactionVersion), callIdx)
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

func Test_GetChainXBlock(t *testing.T) {
	c, err := client.New("wss://testnet-1.chainx.org/ws")
	if err != nil {
		t.Fatal(err)
	}
	var block models.SignedBlock
	err = c.C.Client.Call(&block, "chain_getBlock", "0x8a98b97126880b930b938d69e05fed10b170be06aa15f68318f2e9efd24e490d")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
	h, err := c.C.RPC.Chain.GetFinalizedHead()
	fmt.Println(h.Hex())
	//0x3d02
	//84
	//ff 64fe111943aa50763968bc260dc5a0eacfc0348e7f5eca8e98902749ecbb645c
	//01 763970418e96e1c64512f3e720afa73d35fef3f802b69f43001adf447a559f02
	//   38bbe79a2d0da20019650a503933f1a66d12dbb092524c4af71f70bd42a4a181
	//   c502cc000603
	//  ff56e2771f82b7845270c4f57990435dbacb44982a33ba1a202218a567ae39580d0300943577
}

func Test_CreateAddress(t *testing.T) {
	pub := "a69958eee5de0cb8fb250eba9c4b4ab1675468e68e49a5ebcac22fa9340fe938"
	pubBytes, _ := hex.DecodeString(pub)
	pubBytes = append([]byte{0xff}, pubBytes...)
	address, err := ss58.Encode(pubBytes, ss58.SubstratePrefix)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(address)
}

func Test_GetChainXAccountInfo(t *testing.T) {
	c, err := client.New("wss://testnet-1.chainx.org/ws")
	if err != nil {
		t.Fatal(err)
	}
	acc, err := c.GetAccountInfo("5Fq9MpKxdjzCWEHHtqZ6rdYkKUtW4qwmJV4VHwKBan2hxRyL")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(acc.Nonce)
}
