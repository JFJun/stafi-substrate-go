package expand

import (
	"github.com/JFJun/stafi-substrate-go/expand/chainX"
	"github.com/JFJun/stafi-substrate-go/expand/darwinia"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"strings"
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

func DecodeEventRecords(meta *types.Metadata, rawData string, chainName string) (IEventRecords, error) {
	e := types.EventRecordsRaw(types.MustHexDecodeString(rawData))
	var ier IEventRecords
	switch strings.ToLower(chainName) {
	case "chainx":
		var events chainX.ChainXEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "crab", "darwinia":
		var events darwinia.DarwiniaEventRecords
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
