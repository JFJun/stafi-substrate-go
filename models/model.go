package models

type Bytes []byte
type SignedBlock struct {
	Block         Block `json:"block"`
	Justification Bytes `json:"justification"`
}

type Block struct {
	Extrinsics []string `json:"extrinsics"`
	Header     Header   `json:"header"`
}

type Header struct {
	ParentHash     string      `json:"parentHash"`
	Number         string      `json:"number"`
	StateRoot      string      `json:"stateRoot"`
	ExtrinsicsRoot string      `json:"extrinsicsRoot"`
	Digest         interface{} `json:"digest"`
}

type BlockResponse struct {
	Height     int64                `json:"height"`
	ParentHash string               `json:"parent_hash"`
	BlockHash  string               `json:"block_hash"`
	Timestamp  int64                `json:"timestamp"`
	Extrinsic  []*ExtrinsicResponse `json:"extrinsic"`
}

type ExtrinsicResponse struct {
	Type            string `json:"type"`   //Transfer or another
	Status          string `json:"status"` //success or fail
	Txid            string `json:"txid"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
	Signature       string `json:"signature"`
	Nonce           int64  `json:"nonce"`
	Era             string `json:"era"`
	ExtrinsicIndex  int    `json:"extrinsic_index"`
	EventIndex      int    `json:"event_index"`
	ExtrinsicLength int    `json:"extrinsic_length"`
}

type EventResult struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	ExtrinsicIdx int    `json:"extrinsic_idx"`
	EventIdx     int    `json:"event_idx"`
	Status       string `json:"status"`
	Weight       int64  `json:"weight"` //权重
}

type ExtrinsicDecodeResponse struct {
	AccountId          string                 `json:"account_id"`
	CallCode           string                 `json:"call_code"`
	CallModule         string                 `json:"call_module"`
	Era                string                 `json:"era"`
	Nonce              int64                  `json:"nonce"`
	VersionInfo        string                 `json:"version_info"`
	Signature          string                 `json:"signature"`
	Params             []ExtrinsicDecodeParam `json:"params"`
	CallModuleFunction string                 `json:"call_module_function"`
	Length             int                    `json:"length"`
}

type ExtrinsicDecodeParam struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"value_raw"`
}

type UtilityParamsValue struct {
	CallModule   string                  `json:"call_module"`
	CallFunction string                  `json:"call_function"`
	CallIndex    string                  `json:"call_index"`
	CallArgs     []UtilityParamsValueArg `json:"call_args"`
}

type UtilityParamsValueArg struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"value_raw"`
}
