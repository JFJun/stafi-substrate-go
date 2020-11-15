package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/stafi-substrate-go/client"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"testing"
)

func Test_GetStorage(t *testing.T){
	blockHash:="0x06b90d6b25eb33fb6b475027cd18d934551786cfe2a85e59f65fd519bd6f4905"
	c,err:=client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	meta,err:=c.C.RPC.State.GetMetadataLatest()
	if err != nil {
		t.Fatal(err)
	}
	storage,err:=types.CreateStorageKey(meta,"System","Events",nil,nil)
	if err != nil {
		t.Fatal(err)
	}


	key:=storage.Hex()
	var r interface{}
	err = c.C.Client.Call(&r,"state_getStorage",key,blockHash)
	if err != nil {
		t.Fatal(err)
	}

	e := types.EventRecordsRaw(types.MustHexDecodeString(r.(string)))
	events := types.EventRecords{}
	err = e.DecodeEventRecords(meta, &events)
	if err != nil {
		t.Fatal(err)
	}
	dd,_:=json.Marshal(events)
	fmt.Println(string(dd))
}
