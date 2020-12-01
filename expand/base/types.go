package base

import (
	"encoding/hex"
	"fmt"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

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

	Currencies_Transferred    []EventCurrenciesTransferred
	Currencies_BalanceUpdated []EventCurrenciesBalanceUpdated
	Currencies_Deposited      []EventCurrenciesDeposited
	Currencies_Withdrawn      []EventCurrenciesWithdrawn

	Vesting_VestingScheduleAdded    []EventVestingVestingScheduleAdded
	Vesting_Claimed                 []EventVestingClaimed
	Vesting_VestingSchedulesUpdated []EventVestingVestingSchedulesUpdated

	Multisig_NewMultisig       []types.EventMultisigNewMultisig
	Multisig_MultisigApproval  []types.EventMultisigApproval
	Multisig_MultisigExecuted  []types.EventMultisigExecuted
	Multisig_MultisigCancelled []types.EventMultisigCancelled

	Balances_ReserveRepatriated []EventBalancesReserveRepatriated
	Proxy_Announced             []EventProxyAnnounced
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

type EventCurrenciesTransferred struct {
	Phase      types.Phase
	CurrencyId CurrencyId
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventCurrenciesBalanceUpdated struct {
	Phase      types.Phase
	CurrencyId CurrencyId
	AccountId  types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventCurrenciesDeposited struct {
	Phase      types.Phase
	CurrencyId CurrencyId
	AccountId  types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventCurrenciesWithdrawn struct {
	Phase      types.Phase
	CurrencyId CurrencyId
	AccountId  types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}
type EventVestingVestingScheduleAdded struct {
	Phase           types.Phase
	From            types.AccountID
	To              types.AccountID
	VestingSchedule VestingSchedule
	Topics          []types.Hash
}

type EventVestingClaimed struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type EventVestingVestingSchedulesUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventBalancesReserveRepatriated struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	Balance types.U128
	Status  types.BalanceStatus
	Topics  []types.Hash
}
type EventProxyAnnounced struct {
	Phase  types.Phase
	Who    types.AccountID
	ID     types.AccountID
	Hash   types.Hash
	Topics []types.Hash
}

type CurrencyId types.U32

/*
https://github.com/polkadot-js/api/blob/95c4f03bc3709c58b159623ec5c3c9794e077d08/packages/types/src/interfaces/balances/definitions.ts
*/
type VestingSchedule struct {
	Offset        types.U128
	PerBlock      types.U128
	StartingBlock types.BlockNumber
}

func (d *VestingSchedule) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.Offset)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.PerBlock)
	if err != nil {
		return err
	}

	return decoder.Decode(&d.StartingBlock)
}

type OracleKey types.U32

type OracleValue types.U128

type VecU8L20 struct {
	Value string
}

func (d *VecU8L20) Decode(decoder scale.Decoder) error {
	data := make([]byte, 20)
	err := decoder.Read(data)
	if err != nil {
		return fmt.Errorf("U8L20 read bytes error: %v", err)
	}
	d.Value = hex.EncodeToString(data)
	return nil
}

type VecU8L256 struct {
	Value string
}

func (d *VecU8L256) Decode(decoder scale.Decoder) error {
	data := make([]byte, 256)
	err := decoder.Read(data)
	if err != nil {
		return fmt.Errorf("U8L256 read bytes error: %v", err)
	}
	d.Value = hex.EncodeToString(data)
	return nil
}

type VecU8L32 struct {
	Value string
}

func (d *VecU8L32) Decode(decoder scale.Decoder) error {
	data := make([]byte, 32)
	err := decoder.Read(data)
	if err != nil {
		return fmt.Errorf("U8L32 read bytes error: %v", err)
	}
	d.Value = hex.EncodeToString(data)
	return nil
}

/*
https://github.com/polkadot-js/api/blob/2906b02e413050be04980f472abb69a6991ad5e5/packages/types/src/interfaces/runtime/definitions.ts
*/

type FixedU128 types.U128

/*
https://github.com/open-web3-stack/open-web3.js/blob/2409235b00c03d0adf22258bb972a76d7aa57b4c/packages/orml-type-definitions/src/authority.ts
*/
type ScheduleTaskIndex types.U32

/*
https://github.com/polkadot-js/api/blob/2906b02e413050be04980f472abb69a6991ad5e5/packages/types/src/primitive/StorageKey.ts
*/
type StorageKey types.Bytes

/*
https://github.com/open-web3-stack/open-web3.js/blob/d6f3095f79999e9fc1d39f09a52215684398090c/packages/orml-type-definitions/src/graduallyUpdates.ts
*/
type StorageValue types.Bytes

/*
https://github.com/polkadot-js/api/blob/master/packages/types/src/interfaces/evm/types.ts
*/
type Log struct {
	Address types.H160
	Topics  []types.H256
	Data    types.Bytes
}

func (d *Log) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.Address)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Topics)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.Data)
}

/*
https://github.com/polkadot-js/api/blob/master/packages/types/src/interfaces/collective/types.ts
*/
type MemberCount types.U32
