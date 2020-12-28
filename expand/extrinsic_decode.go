package expand

/*
扩展：解析extrinsic
	substrate2.0的extrinsic都是这样，所以这里的变动其实很小
	这里编写都是为了与github.com/JFJun/substrate-go保持一制，所以会显得有点混乱
*/
import (
	"fmt"
	"github.com/JFJun/stafi-substrate-go/utils"
	"github.com/huandu/xstrings"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type ExtrinsicDecoder struct {
	ExtrinsicLength     int              `json:"extrinsic_length"`
	VersionInfo         string           `json:"version_info"`
	ContainsTransaction bool             `json:"contains_transaction"`
	Address             string           `json:"address"`
	Signature           string           `json:"signature"`
	SignatureVersion    int              `json:"signature_version"`
	Nonce               int              `json:"nonce"`
	Era                 string           `json:"era"`
	Tip                 string           `json:"tip"`
	CallIndex           string           `json:"call_index"`
	CallModule          string           `json:"call_module"`
	CallModuleFunction  string           `json:"call_module_function"`
	Params              []ExtrinsicParam `json:"params"`
	me                  *MetadataExpand
	Value               interface{}
}

type ExtrinsicParam struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"value_raw"`
}

func NewExtrinsicDecoder(meta *types.Metadata) (*ExtrinsicDecoder, error) {
	ed := new(ExtrinsicDecoder)
	var err error
	ed.me, err = NewMetadataExpand(meta)
	if err != nil {
		return nil, err
	}
	return ed, nil
}

func (ed *ExtrinsicDecoder) ProcessExtrinsicDecoder(decoder scale.Decoder) error {
	var length types.UCompact
	err := decoder.Decode(&length)
	if err != nil {
		return fmt.Errorf("decode extrinsic: length error: %v", err)
	}
	ed.ExtrinsicLength = int(utils.UCompactToBigInt(length).Int64())
	vi, err := decoder.ReadOneByte()
	if err != nil {
		return fmt.Errorf("decode extrinsic: read version info error: %v", err)
	}
	ed.VersionInfo = utils.BytesToHex([]byte{vi})
	ed.ContainsTransaction = utils.U256(ed.VersionInfo).Int64() >= 80
	//大多数都是84了，所以只处理84
	if ed.VersionInfo == "04" || ed.VersionInfo == "84" {
		if ed.ContainsTransaction {
			// 1. 解析from地址
			var address Address
			err = decoder.Decode(&address)
			if err != nil {
				return fmt.Errorf("decode extrinsic: decode address error: %v", err)
			}
			ed.Address = address.Value
			//2。解析签名版本
			var sv types.U8
			err = decoder.Decode(&sv)
			if err != nil {
				return fmt.Errorf("decode extrinsic: decode signature version error: %v", err)
			}
			ed.SignatureVersion = int(sv)
			// 3。 解析签名
			if ed.SignatureVersion == 2 {
				//解析 ecdsa signature
				sig := make([]byte, 65)
				err = decoder.Read(sig)
				if err != nil {
					return fmt.Errorf("decode extrinsic: decode ecdsa signature error: %v", err)
				}
				ed.Signature = utils.BytesToHex(sig)
			} else {
				// 解析 sr25519 signature
				var sig types.Signature
				err = decoder.Decode(&sig)
				if err != nil {
					return fmt.Errorf("decode extrinsic: decode sr25519 signature error: %v", err)
				}
				ed.Signature = sig.Hex()
			}
			// 4. 解析era
			var era types.ExtrinsicEra
			err = decoder.Decode(&era)
			if err != nil {
				return fmt.Errorf("decode extrinsic: decode era error: %v", err)
			}
			if era.IsMortalEra {
				eraBytes := []byte{era.AsMortalEra.First, era.AsMortalEra.Second}
				ed.Era = utils.BytesToHex(eraBytes)
			}
			//5. 解析nonce
			var nonce types.UCompact
			err = decoder.Decode(&nonce)
			if err != nil {
				return fmt.Errorf("decode extrinsic: decode nonce error: %v", err)
			}
			//new

			ed.Nonce = int(utils.UCompactToBigInt(nonce).Int64())
			// 6.解析tip
			var tip types.UCompact

			err = decoder.Decode(&tip)
			if err != nil {
				return fmt.Errorf("decode tip error: %v", err)
			}
			ed.Tip = fmt.Sprintf("%d", utils.UCompactToBigInt(tip).Int64())
		}
		//处理callIndex
		callIndex := make([]byte, 2)
		err = decoder.Read(callIndex)
		if err != nil {
			return fmt.Errorf("decode extrinsic: read call index bytes error: %v", err)
		}
		ed.CallIndex = xstrings.RightJustify(utils.IntToHex(callIndex[0]), 2, "0") +
			xstrings.RightJustify(utils.IntToHex(callIndex[1]), 2, "0")
	} else {
		return fmt.Errorf("extrinsics version %s is not support", ed.VersionInfo)
	}
	if ed.CallIndex != "" {
		_ = ed.decodeCallIndex(decoder)
		//if err != nil {
		//	return err
		//}
	}
	result := map[string]interface{}{
		"extrinsic_length": ed.ExtrinsicLength,
		"version_info":     ed.VersionInfo,
	}
	if ed.ContainsTransaction {
		result["account_id"] = ed.Address
		result["signature"] = ed.Signature
		result["nonce"] = ed.Nonce
		result["era"] = ed.Era
	}
	if ed.CallIndex != "" {
		result["call_code"] = ed.CallIndex
		result["call_module_function"] = ed.CallModuleFunction
		result["call_module"] = ed.CallModule
	}
	result["nonce"] = ed.Nonce
	result["era"] = ed.Era
	result["tip"] = ed.Tip
	result["params"] = ed.Params
	result["length"] = ed.ExtrinsicLength
	ed.Value = result
	return nil
}

func (ed *ExtrinsicDecoder) decodeCallIndex(decoder scale.Decoder) error {
	var err error
	//避免指针为空
	defer func() {
		if errs := recover(); errs != nil {
			err = fmt.Errorf("decode call catch panic ,err=%v", errs)
		}
	}()
	//	解析 call index
	// 这里只能硬编码了，因为decode函数的原因，无法动态的根据type name去解析
	// 这里我只解析自己想要的，比如说Timestamp,Balance.transfer,Utility.batch
	modName, callName, err := ed.me.MV.FindNameByCallIndex(ed.CallIndex)
	if err != nil {
		return fmt.Errorf("decode call: %v", err)
	}
	ed.CallModule = modName
	ed.CallModuleFunction = callName
	switch modName {
	case "Timestamp":
		if callName == "set" {
			//Compact<Moment>
			var u types.UCompact
			err = decoder.Decode(&u)
			if err != nil {
				return fmt.Errorf("decode call: decode Timestamp.set error: %v", err)
			}

			ed.Params = append(ed.Params,
				ExtrinsicParam{
					Name:  "now",
					Type:  "Compact<Moment>",
					Value: utils.UCompactToBigInt(u).Int64(),
				})
		}
	case "Balances":
		if callName == "transfer" || callName == "transfer_keep_alive" {
			// 0 ---> 	Address
			var addrValue string
			var address Address
			err = decoder.Decode(&address)
			if err != nil {
				return fmt.Errorf("decode call: decode Balances.transfer.Address error: %v", err)
			}
			addrValue = address.Value

			ed.Params = append(ed.Params,
				ExtrinsicParam{
					Name:     "dest",
					Type:     "Address",
					Value:    addrValue,
					ValueRaw: addrValue,
				})
			// 1 ----> Compact<Balance>
			var b types.UCompact
			err = decoder.Decode(&b)
			if err != nil {
				return fmt.Errorf("decode call: decode Balances.transfer.Compact<Balance> error: %v", err)
			}

			ed.Params = append(ed.Params,
				ExtrinsicParam{
					Name:  "value",
					Type:  "Compact<Balance>",
					Value: utils.UCompactToBigInt(b).Int64(),
				})
		}
	case "Utility":
		if callName == "batch" {
			// 0--> calls   Vec<Call>
			var tc TransferCall
			vec := new(Vec)
			err := vec.ProcessVec(decoder, tc)
			if err != nil {
				return fmt.Errorf("decode call: decode Utility.batch error: %v", err)
			}
			//utils.CheckStructData(vec.Value)
			ep := ExtrinsicParam{}
			ep.Name = "calls"
			ep.Type = "Vec<Call>"
			var result []interface{}
			for _, value := range vec.Value {
				tcv := value.(*TransferCall)
				//检查一下是否为BalanceTransfer
				data := tcv.Value.(map[string]interface{})
				callIndex := data["call_index"].(string)
				btCallIdx, err := ed.me.MV.GetCallIndex("Balances", "transfer")
				if err != nil {
					return fmt.Errorf("decode Utility.batch: get  Balances.transfer call index error: %v", err)
				}
				btkaCallIdx, err := ed.me.MV.GetCallIndex("Balances", "transfer_keep_alive")
				if err != nil {
					return fmt.Errorf("decode Utility.batch: get  Balances.transfer_keep_alive call index error: %v", err)
				}
				if callIndex == btCallIdx || callIndex == btkaCallIdx {
					mn, cn, err := ed.me.MV.FindNameByCallIndex(callIndex)
					if err != nil {
						return fmt.Errorf("decode Utility.batch: get call index error: %v", err)
					}
					if mn != "Balances" {
						return fmt.Errorf("decode Utility.batch:  call module name is not 'Balances' ,NAME=%s", mn)
					}
					data["call_function"] = cn
					data["call_module"] = mn
					result = append(result, data)
				}
			}
			ep.Value = result
			ed.Params = append(ed.Params, ep)
		}
	default:
		// unsopport
		return nil

	}
	return nil
}

//----------support for bifrost
