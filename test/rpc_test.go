package test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/JFJun/stafi-substrate-go/client"
	"github.com/JFJun/stafi-substrate-go/expand"
	"github.com/JFJun/stafi-substrate-go/models"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"testing"
)

func Test_RpcClient(t *testing.T) {
	c, err := client.New("wss://mainnet-rpc.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	var block models.SignedBlock
	err = c.C.Client.Call(&block, "chain_getBlock", "0x9832ff45a5d135a37518459cbab3331331e6ae30a0aa6298be4e03ea6c42f71b")
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
	c, err := client.New("ws://fis.rylink.io:31844")
	//c, err := client.New("wss://rpc.polkadot.io")
	c.SetPrefix(ss58.StafiPrefix)
	if err != nil {
		t.Fatal(err)
	}
	//c.SetPrefix(ss58.StafiPrefix)
	ai, err := c.GetAccountInfo("34R7rV86nwK478QbBWjvCrwdR6UVpBJuvbDwy7MDN328HaWs")
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
	c, err := client.New("wss://mainnet.chainx.org/ws")
	//c, err := client.New("")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.ChainXPrefix)
	expand.SetSerDeOptions(false)
	resp, err := c.GetBlockByNumber(1753086)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}

func Test_GetGenesisHash(t *testing.T) {
	c, err := client.New("wss://testnet-1.chainx.org/ws")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(c.GetGenesisHash())
}

func Test_CalcFee(t *testing.T) {
	c, err := client.New("wss://mainnet-rpc.stafi.io")
	//c, err := client.New("wss://rpc.polkadot.io")
	c.SetPrefix(ss58.StafiPrefix)
	if err != nil {
		t.Fatal(err)
	}
	md, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}

	_, value, err := md.MV.GetConstants("TransactionPayment", "TransactionByteFee")
	if err != nil {
		t.Fatal(err)
	}
	var tbf types.U128
	decoder := scale.NewDecoder(bytes.NewReader(value))
	err = decoder.Decode(&tbf)
	if err != nil {
		t.Fatal(err)
	}
	transactionByteFee := tbf.Uint64()
	fmt.Println("pre_Bytes:", transactionByteFee)

	_, value2, err := md.MV.GetConstants("System", "ExtrinsicBaseWeight")
	var w types.U32
	decoder2 := scale.NewDecoder(bytes.NewReader(value2))
	err = decoder2.Decode(&w)
	if err != nil {
		t.Fatal(err)
	}
	extrinsicBaseWeight := int64(w)
	fmt.Println("extrinsicWeight:", extrinsicBaseWeight)

	storage, err := types.CreateStorageKey(c.Meta, "TransactionPayment", "NextFeeMultiplier", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	var weight types.U128
	data, _ := hex.DecodeString("db78f4f0026651e6d85e32d9601a92571e107b7c907d85ea606d3cc12a7285bf")
	prehash := types.NewHash(data)
	ok, err := c.C.RPC.State.GetStorage(storage, &weight, prehash)
	if err != nil {
		t.Fatal(err)
	}
	//var r interface{}
	//err = c.C.Client.Call(&r,"state_getStorageAt",storage.Hex(),"0x822ad83fdad1b40ed35159b1e32c8b5d32d3e034b08bf2e9ebde18f3141004b9")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println("R: ",r)
	//rH:=r.(string)
	//rH = strings.TrimPrefix(rH,"0x")
	//dd,_:=hex.DecodeString(rH)
	//fmt.Println(dd)
	if !ok {
		t.Fatal(111)
	}
	fmt.Println("WeightMultiplier: ", weight)
	_, value3, err := md.MV.GetConstants("TransactionPayment", "WeightToFee")
	if err != nil {
		t.Fatal(err)
	}
	decoder3 := scale.NewDecoder(bytes.NewReader(value3))
	vec := new(expand.Vec)
	var wtfc expand.WeightToFeeCoefficient
	err = vec.ProcessVec(*decoder3, wtfc)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(vec.Value)
	fmt.Println("WTFC:", string(d))
	vv := vec.Value[0]
	hj := vv.(*expand.WeightToFeeCoefficient)
	cf := expand.NewCalcFee(hj, extrinsicBaseWeight, int64(transactionByteFee), int64(weight.Int64()))
	fee := cf.CalcPartialFee(190949000, 143)
	fmt.Println(fee)
}

func Test_GetChainName(t *testing.T) {
	c, err := client.New("wss://cc1.darwinia.network")
	if err != nil {
		t.Fatal(err)
	}
	d, err := c.C.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(d.SpecName)
}

func Test_GetLatestBlock(t *testing.T) {

}

// 1700dcea9317bceb28b52bdae9229a3794de4ca85e36d990a78f779c6fd7f27eb54102890700003c001c0000000300000034f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c6534f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c65
// 1700dcea9317bceb28b52bdae9229a3794de4ca85e36d990a78f779c6fd7f27eb54102890700003c001c0000000300000034f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c6534f61bfda344b3fad3c3e38832a91448b3c613b199eb23e5110a635d71c13c65

func Test_GetChainEvent(t *testing.T) {
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println(c.ChainName)
	//mod,_:=json.Marshal(c.Meta.AsMetadataV12.Modules)
	//fmt.Println(string(mod))
	for _, mod := range c.Meta.AsMetadataV12.Modules {
		if mod.HasEvents {
			for _, event := range mod.Events {
				//typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
				//if IsExist(typeName) {
				//	continue
				//}

				//fmt.Println(event.Name)
				for _, arg := range event.Args {
					fmt.Println(arg)
				}
			}
		}
	}
}
