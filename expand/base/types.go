package base

import "github.com/stafiprotocol/go-substrate-rpc-client/types"

type BaseEventRecords struct {
	types.EventRecords
	Treasury_BountyProposed     []EventTreasuryBountyProposed
	Treasury_BountyRejected     []EventTreasuryBountyRejected
	Treasury_BountyBecameActive []EventTreasuryBountyBecameActive
	Treasury_BountyAwarded      []EventTreasuryBountyAwarded
	Treasury_BountyClaimed      []EventTreasuryBountyClaimed
	Treasury_BountyCanceled     []EventTreasuryBountyCanceled
	Treasury_BountyExtended     []EventTreasuryBountyExtended
	TechnicalMembership_Dummy   []EventTechnicalMembershipDummy
}

func (d *BaseEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return d.Balances_Transfer
}
func (d *BaseEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return d.System_ExtrinsicSuccess
}
func (d *BaseEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return d.System_ExtrinsicFailed
}

type EventTreasuryBountyProposed struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}

type EventTreasuryBountyRejected struct {
	Phase       types.Phase
	BountyIndex types.U32
	Balance     types.U128
	Topics      []types.Hash
}

type EventTreasuryBountyBecameActive struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}

type EventTreasuryBountyAwarded struct {
	Phase       types.Phase
	BountyIndex types.U32
	Who         types.AccountID
	Topics      []types.Hash
}

type EventTreasuryBountyClaimed struct {
	Phase       types.Phase
	BountyIndex types.U32
	Balance     types.U128
	Who         types.AccountID
	Topics      []types.Hash
}

type EventTreasuryBountyCanceled struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}

type EventTreasuryBountyExtended struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}

type EventTechnicalMembershipDummy struct {
	Phase       types.Phase
	PhantomData types.Null
	Topics      []types.Hash
}
