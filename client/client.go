package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/JFJun/stafi-substrate-go/expand"
	"github.com/JFJun/stafi-substrate-go/models"
	"github.com/JFJun/stafi-substrate-go/utils"
	"github.com/shopspring/decimal"
	gsrc "github.com/stafiprotocol/go-substrate-rpc-client"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"golang.org/x/crypto/blake2b"
	"log"
	"strconv"
)

type Client struct {
	C                  *gsrc.SubstrateAPI
	Meta               *types.Metadata
	prefix             []byte //币种的前缀
	SpecVersion        int
	TransactionVersion int
	genesisHash        string
}

func New(url string) (*Client, error) {
	c := new(Client)
	var err error
	c.C, err = gsrc.NewSubstrateAPI(url)
	if err != nil {
		return nil, err
	}
	err = c.checkRuntimeVersion()
	if err != nil {
		return nil, err
	}
	c.prefix = ss58.SubstratePrefix //默认prefix
	return c, nil
}

func (c *Client) checkRuntimeVersion() error {
	v, err := c.C.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return fmt.Errorf("init runtime version error,err=%v", err)
	}
	c.TransactionVersion = int(uint32(v.TransactionVersion))
	specVersion := int(uint32(v.SpecVersion))
	if specVersion != c.SpecVersion {
		c.Meta, err = c.C.RPC.State.GetMetadataLatest()
		if err != nil {
			return fmt.Errorf("init metadata error: %v", err)
		}
		c.SpecVersion = specVersion
	}
	return nil
}

func (c *Client) GetGenesisHash() string {
	if c.genesisHash != "" {
		return c.genesisHash
	}
	hash, err := c.C.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return ""
	}
	c.genesisHash = hash.Hex()
	return hash.Hex()
}

func (c *Client) SetPrefix(prefix []byte) {
	c.prefix = prefix
}
func (c *Client) GetBlockByNumber(height int64) (*models.BlockResponse, error) {
	hash, err := c.C.RPC.Chain.GetBlockHash(uint64(height))
	if err != nil {
		return nil, fmt.Errorf("get block hash error:%v,height:%d", err, height)
	}
	blockHash := hash.Hex()

	return c.GetBlockByHash(blockHash)
}

func (c *Client) GetBlockByHash(blockHash string) (*models.BlockResponse, error) {
	var (
		block *models.SignedBlock
		err   error
	)
	err = c.checkRuntimeVersion()
	if err != nil {
		return nil, err
	}
	err = c.C.Client.Call(&block, "chain_getBlock", blockHash)
	if err != nil {
		return nil, fmt.Errorf("get block error: %v", err)
	}
	blockResp := new(models.BlockResponse)
	number, _ := strconv.ParseInt(utils.RemoveHex0x(block.Block.Header.Number), 16, 64)
	blockResp.Height = number
	blockResp.ParentHash = block.Block.Header.ParentHash
	blockResp.BlockHash = blockHash
	if len(block.Block.Extrinsics) > 0 {
		err = c.parseExtrinsicByDecode(block.Block.Extrinsics, blockResp)
		if err != nil {
			return nil, err
		}

		err = c.parseExtrinsicByStorage(blockHash, blockResp)
		if err != nil {
			return nil, err
		}
	}
	return blockResp, nil
}

type parseBlockExtrinsicParams struct {
	from, to, sig, era, txid string
	nonce                    int64
	extrinsicIdx             int
}

func (c *Client) parseExtrinsicByDecode(extrinsics []string, blockResp *models.BlockResponse) error {
	var (
		params    []parseBlockExtrinsicParams
		timestamp int64
		//idx int
	)
	defer func() {
		if err := recover(); err != nil {
			blockResp.Timestamp = timestamp
			blockResp.Extrinsic = []*models.ExtrinsicResponse{}
			log.Printf("parse %d block extrinsic error,Err=[%v]", blockResp.Height, err)
		}
	}()

	for i, extrinsic := range extrinsics {
		extrinsic = utils.Remove0X(extrinsic)
		data, err := hex.DecodeString(extrinsic)
		if err != nil {
			return fmt.Errorf("hex.decode extrinsic error: %v", err)
		}
		decoder := scale.NewDecoder(bytes.NewReader(data))
		ed, err := expand.NewExtrinsicDecoder(c.Meta)
		if err != nil {
			return fmt.Errorf("new extrinsic decode error: %v", err)
		}
		err = ed.ProcessExtrinsicDecoder(*decoder)
		if err != nil {
			return fmt.Errorf("decode extrinsic error: %v", err)
		}
		var resp models.ExtrinsicDecodeResponse
		d, _ := json.Marshal(ed.Value)
		if len(d) == 0 {
			return errors.New("unknown extrinsic decode response")
		}
		err = json.Unmarshal(d, &resp)
		if err != nil {
			return fmt.Errorf("json unmarshal extrinsic decode error: %v", err)
		}

		switch resp.CallModule {
		case "Timestamp":
			for _, param := range resp.Params {
				if param.Name == "now" {
					timestamp = int64(param.Value.(float64))
				}
			}
		case "Balances":
			if resp.CallModuleFunction == "transfer" || resp.CallModuleFunction == "transfer_keep_alive" {
				blockData := parseBlockExtrinsicParams{}
				blockData.from, _ = ss58.EncodeByPubHex(resp.AccountId, c.prefix)
				blockData.era = resp.Era
				blockData.sig = resp.Signature
				blockData.nonce = resp.Nonce
				blockData.extrinsicIdx = i
				blockData.txid = c.createTxHash(extrinsic)
				for _, param := range resp.Params {
					if param.Name == "dest" {

						blockData.to, _ = ss58.EncodeByPubHex(param.Value.(string), c.prefix)
					}
				}
				params = append(params, blockData)
			}

		case "Utility":
			if resp.CallModuleFunction == "batch" {
				for _, param := range resp.Params {
					if param.Name == "calls" {
						switch param.Value.(type) {
						case []interface{}:

							d, _ := json.Marshal(param.Value)
							var values []models.UtilityParamsValue
							err = json.Unmarshal(d, &values)
							if err != nil {
								continue
							}

							for _, value := range values {
								if value.CallModule == "Balances" {
									if value.CallFunction == "transfer" || value.CallFunction == "transfer_keep_alive" {
										if len(value.CallArgs) > 0 {
											for _, arg := range value.CallArgs {
												if arg.Name == "dest" {
													blockData := parseBlockExtrinsicParams{}
													blockData.from, _ = ss58.EncodeByPubHex(resp.AccountId, c.prefix)
													blockData.era = resp.Era
													blockData.sig = resp.Signature
													blockData.nonce = resp.Nonce
													blockData.extrinsicIdx = i
													blockData.txid = c.createTxHash(extrinsic)
													blockData.to, _ = ss58.EncodeByPubHex(arg.ValueRaw, c.prefix)
													params = append(params, blockData)
												}
											}
										}
									}
								}
							}
						default:
							continue
						}
					}
				}
			}
		default:
			//todo  add another call_module 币种不同可能使用的call_module不一样
			continue
		}
	}
	blockResp.Timestamp = timestamp
	//解析params
	if len(params) == 0 {
		blockResp.Extrinsic = []*models.ExtrinsicResponse{}
		return nil
	}
	blockResp.Extrinsic = make([]*models.ExtrinsicResponse, len(params))
	for idx, param := range params {
		// write by jun 2020/06/18
		// 避免不同高度出现相同txid的情况  详情高度： 552851  552911
		//txid := fmt.Sprintf("%s_%d-%d", param.txid, blockResp.Height, param.extrinsicIdx)
		e := new(models.ExtrinsicResponse)
		e.Signature = param.sig
		e.FromAddress = param.from
		e.ToAddress = param.to
		e.Nonce = param.nonce
		e.Era = param.era
		e.ExtrinsicIndex = param.extrinsicIdx
		//e.Txid = txid
		e.Txid = param.txid
		blockResp.Extrinsic[idx] = e
	}
	//utils.CheckStructData(blockResp)
	return nil
}

func (c *Client) parseExtrinsicByStorage(blockHash string, blockResp *models.BlockResponse) error {
	var (
		storage types.StorageKey
		err     error
	)
	defer func() {
		if err1 := recover(); err1 != nil {
			err = fmt.Errorf("panic decode event: %v", err1)
		}
	}()

	storage, err = types.CreateStorageKey(c.Meta, "System", "Events", nil, nil)
	if err != nil {
		return fmt.Errorf("create storage key error: %v", err)
	}
	key := storage.Hex()
	var result interface{}
	err = c.C.Client.Call(&result, "state_getStorageAt", key, blockHash)
	if err != nil {
		return fmt.Errorf("get storage data error: %v", err)
	}
	e := types.EventRecordsRaw(types.MustHexDecodeString(result.(string)))
	var events types.EventRecords
	err = e.DecodeEventRecords(c.Meta, &events)
	if err != nil {
		return fmt.Errorf("decode event data error: %v", err)
	}
	var res []models.EventResult
	failedMap := make(map[int]bool)
	if len(events.Balances_Transfer) > 0 {
		//有失败的交易
		for _, failed := range events.System_ExtrinsicFailed {
			if failed.Phase.IsApplyExtrinsic {
				extrinsicIdx := failed.Phase.AsApplyExtrinsic
				//记录到失败的map中
				failedMap[int(extrinsicIdx)] = true
			}
		}
		for _, ebt := range events.Balances_Transfer {
			if !ebt.Phase.IsApplyExtrinsic {
				continue
			}
			extrinsicIdx := int(ebt.Phase.AsApplyExtrinsic)
			var r models.EventResult
			r.ExtrinsicIdx = extrinsicIdx
			fromHex := hex.EncodeToString(ebt.From[:])
			r.From, err = ss58.EncodeByPubHex(fromHex, c.prefix)
			if err != nil {
				r.From = ""
				continue
			}
			toHex := hex.EncodeToString(ebt.To[:])
			r.To, err = ss58.EncodeByPubHex(toHex, c.prefix)
			if err != nil {
				r.To = ""
				continue
			}
			r.Amount = decimal.NewFromInt(ebt.Value.Int64()).String()
			res = append(res, r)
		}

	}
	for _, e := range blockResp.Extrinsic {
		e.Status = "fail"
		e.Type = "transfer"
		if len(res) > 0 {
			for _, r := range res {
				if e.ExtrinsicIndex == r.ExtrinsicIdx {
					if e.ToAddress == r.To {
						if failedMap[e.ExtrinsicIndex] {
							e.Status = "fail"
						} else {
							e.Status = "success"
						}
						e.Type = "transfer"
						e.Amount = r.Amount
						e.ToAddress = r.To
						//计算手续费
						e.Fee = c.calcFee(&events, e.ExtrinsicIndex)
					} else {
						e.Status = fmt.Sprintf("to address is not equal,a1=[%s],a2=[%s]", e.ToAddress, r.To)
					}
				}
			}
		}
	}

	return nil
}
func (c *Client) calcFee(events *types.EventRecords, extrinsicIdx int) string {
	var (
		fee = decimal.Zero
	)

	for _, bd := range events.Balances_Deposit {
		if bd.Phase.IsApplyExtrinsic && int(bd.Phase.AsApplyExtrinsic) == extrinsicIdx {
			fee = fee.Add(decimal.NewFromInt(bd.Balance.Int64()))
		}
	}
	for _, td := range events.Treasury_Deposit {
		if td.Phase.IsApplyExtrinsic && int(td.Phase.AsApplyExtrinsic) == extrinsicIdx {
			fee = fee.Add(decimal.NewFromInt(td.Deposited.Int64()))
		}
	}
	return fee.String()
}

func (c *Client) createTxHash(extrinsic string) string {
	data, _ := hex.DecodeString(utils.RemoveHex0x(extrinsic))
	d := blake2b.Sum256(data)
	return "0x" + hex.EncodeToString(d[:])
}

func (c *Client) GetAccountInfo(address string) (*types.AccountInfo, error) {

	var (
		storage types.StorageKey
		err     error
		pub     []byte
	)
	defer func() {
		if err1 := recover(); err1 != nil {
			err = fmt.Errorf("panic decode event: %v", err1)
		}
	}()
	err = c.checkRuntimeVersion()
	if err != nil {
		return nil, err
	}
	pub, err = ss58.DecodeToPub(address)
	if err != nil {
		return nil, fmt.Errorf("ss58 decode address error: %v", err)
	}
	storage, err = types.CreateStorageKey(c.Meta, "System", "Account", pub, nil)
	if err != nil {
		return nil, fmt.Errorf("create System.Account storage error: %v", err)
	}
	var accountInfo types.AccountInfo
	var ok bool
	ok, err = c.C.RPC.State.GetStorageLatest(storage, &accountInfo)
	if err != nil || !ok {
		return nil, fmt.Errorf("get account info error: %v", err)
	}
	return &accountInfo, nil
}
