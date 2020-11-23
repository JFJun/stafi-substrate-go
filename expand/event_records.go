package expand

import (
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/JFJun/stafi-substrate-go/expand/chainX"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type IEventRecords interface {
	GetBalancesTransfer() []types.EventBalancesTransfer
	GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess
	GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed
}

type DefaultEventRecords struct {
	types.EventRecords
}

func (d *DefaultEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return d.Balances_Transfer
}
func (d *DefaultEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return d.System_ExtrinsicSuccess
}
func (d *DefaultEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return d.System_ExtrinsicFailed
}

func DecodeEventRecords(meta *types.Metadata, rawData string, prefix []byte) (IEventRecords, error) {
	e := types.EventRecordsRaw(types.MustHexDecodeString(rawData))
	var ier IEventRecords
	switch string(prefix) {
	case string(ss58.ChainXPrefix), string(ss58.SubstratePrefix):
		var events chainX.ChainXEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	default:
		var events DefaultEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	}
	return ier, nil
}
