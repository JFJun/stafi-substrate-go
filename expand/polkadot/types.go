package polkadot

import (
	"github.com/JFJun/stafi-substrate-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type PolkadotEventRecords struct {
	base.BaseEventRecords
	Claims_Claimed                    []EventClaimsClaimed
	ElectionsPhragmen_VoterReported   []EventElectionsPhragmenVoterReported
	ElectionsPhragmen_MemberRenounced []EventElectionsPhragmenMemberRenounced
	ElectionsPhragmen_MemberKicked    []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_ElectionError   []EventElectionsPhragmenElectionError
	ElectionsPhragmen_EmptyTerm       []EventElectionsPhragmenEmptyTerm
	ElectionsPhragmen_NewTerm         []EventElectionsPhragmenNewTerm
	Democracy_Blacklisted             []EventDemocracyBlacklisted
}

func (p PolkadotEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return p.Balances_Transfer
}

func (p PolkadotEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return p.System_ExtrinsicSuccess
}

func (p PolkadotEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return p.System_ExtrinsicFailed
}

type EventDemocracyBlacklisted struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type EventElectionsPhragmenNewTerm struct {
	Phase   types.Phase
	NewTerm []PhragmenNewTermItem
	Topics  []types.Hash
}

type EventElectionsPhragmenEmptyTerm struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventElectionsPhragmenElectionError struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventElectionsPhragmenMemberKicked struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenMemberRenounced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenVoterReported struct {
	Phase  types.Phase
	Who1   types.AccountID
	Who2   types.AccountID
	Bool   types.Bool
	Topics []types.Hash
}
type EventClaimsClaimed struct {
	Phase           types.Phase
	AccountId       types.AccountID
	EthereumAddress base.VecU8L20
	Balance         types.U128
	Topics          []types.Hash
}

type PhragmenNewTermItem struct {
	AccountId types.AccountID
	Balance   types.U128
}

func (d *PhragmenNewTermItem) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.AccountId)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Balance)
	if err != nil {
		return err
	}

	return nil
}
