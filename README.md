#   stafi go sdk
    **************** 如果要使用(2021-02-23升级)polkadot，请使用https://github.com/JFJun/bifrost-go.git   ****************************
## 介绍
    1. 该包不仅可用于波卡系列的区块解析以及离线签名。
        目前该包支持以下币种：
            stafi(FIS),kusama(KSM),chainX2.0(PCX),darwinia(CRING,RING)
    2. 这个包其实是对github.com/JFJun/substrate-go包的升级，所以功能与两者相似，只不过这个包更简洁更稳定。
    3. 发送交易建议使用新的方法，功能都是一样，进行了简单的封装，相对之前看着更加简洁，可以查看相应的测试方法：
        https://github.com/JFJun/stafi-substrate-go/test/tx2_test.go
## 使用
### 1. 解析区块
        // 初始化客户端
        c, err := client.New("wss://rpc.polkadot.io")
    	if err != nil {
    		t.Fatal(err)
    	}
    	//expand.SetSerDeOptions(false)
    	// 设置地址的前缀，默认是 0x2a
    	c.SetPrefix(ss58.StafiPrefix)
    	resp, err := c.GetBlockByNumber(2517230)
    	if err != nil {
    		t.Fatal(err)
    	}
    	d, _ := json.Marshal(resp)
    	fmt.Println(string(d))
        
### 2. Balances.transfer转账
    // 1. 初始化客户端
    c,err:=client.New("wss://crab.darwinia.network")
    if err != nil {
    	t.Fatal(err)
    }
    //2. 下面这句代码因为每个链的不同而不同，我看源代码是因为一个Indices这个module的有无决定的
    //  反正如果提交交易包含 MultiSignature错误什么的，就把 这个bool值设置为相反就行了
    expand.SetSerDeOptions(false)
    //3. 设置链的前缀
    c.SetPrefix(ss58.StafiPrefix)
    //4。 创建交易
    acc,err := c.GetAccountInfo(from)
    if err != nil {
        t.Fatal(err)
    }
    nonce := uint64(acc.Nonce)
    transaction:=tx.CreateTransaction(from,to,amount,nonce)
    //5. 设置交易签名需要的参数
    callIdx,err:=ed.MV.GetCallIndex("Balances","transfer")
    if err != nil {
        t.Fatal(err)
    }
    transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(),c.GetGenesisHash())
    .SetSpecVersionAndCallId(uint32(c.SpecVersion),uint32(c.TransactionVersion),callIdx)
    .SetTip(0)          //自行选择设置，可增加给矿工的手续费
    .SetEra(0,0)        //自行选择设置，设置交易如果处于pending状态，最多能存活多少个块
    //6. 签名交易
    tt,err:=transaction.SignTransaction("私钥",crypto.Sr25519Type)
    if err != nil {
    	t.Fatal(err)
    }
    // 7. 提交交易
    var result interface{}
    err = c.C.Client.Call(&result,"author_submitExtrinsic",tt)
    if err != nil {
    	t.Fatal(err)
    }
    txid:=result.(string)
    fmt.Println(txid)

### 3. Balances.transfer_keep_alive转账
    *** 这个转账模式我没有试过，不过按理来说适合Balances.transfer一样的 ***
    将Balances.transfer中的第5部获取callIdx改为
    callIdx,err:=ed.MV.GetCallIndex("Balances","transfer_keep_alive")
        if err != nil {
            t.Fatal(err)
        }
    其他的都和Balances.transfer一样
    
### 4. Utility.batch转账
    将Balances.transfer的第四部改为：
    pa:=map[string]int64{
        "to1":amount1,
        "to2"amount2,
        ...
    }
    uitlityCallIdx,err:=ed.MV.GetCallIndex("Utility","batch")
   if err != nil {
       t.Fatal(err)
   }
    acc,err := c.GetAccountInfo(from)
    if err != nil {
        t.Fatal(err)
    }
    nonce := uint64(acc.Nonce)
    transaction:=tx.CreateUtilityBatchTransaction(from,nonce,pa,uitlityCallIdx)
    
    其他的步骤和Balances.transfer是一样的
