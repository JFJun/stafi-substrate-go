package stafi

import (
	"github.com/JFJun/stafi-substrate-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type StafiEventRecords struct {
	base.BaseEventRecords
	BridgeCommon_ChainWhitelisted []EventBridgeCommonChainWhitelisted
	BridgeCommon_FungibleTransfer []EventBridgeCommonFungibleTransfer
	BridgeCommon_ChainFeesSet     []EventBridgeCommonChainFeesSet
}
type EventBridgeCommonChainFeesSet struct {
	Phase   types.Phase
	ChainId ChainId
	Balance types.U128
	Topics  []types.Hash
}

type EventBridgeCommonChainWhitelisted struct {
	Phase   types.Phase
	ChainId ChainId
	Topics  []types.Hash
}
type EventBridgeCommonFungibleTransfer struct {
	Phase        types.Phase
	Who          types.AccountID
	ChainId      ChainId
	DepositNonce DepositNonce
	ResourceId   base.VecU8L32
	U256         types.U256
	Data         types.Bytes
	Topics       []types.Hash
}

type ChainId types.U8
type DepositNonce types.U64
