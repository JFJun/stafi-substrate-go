package crust

import (
	"fmt"
	"github.com/JFJun/stafi-substrate-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type CRustEventRecords struct {
	base.BaseEventRecords
	Offences_Offence []EventOffencesOffence
	//Treasury_BountyProposed			[]EventTreasuryBountyProposed
	//Treasury_BountyRejected			[]EventTreasuryBountyRejected
	//Treasury_BountyBecameActive		[]EventTreasuryBountyBecameActive
	//Treasury_BountyAwarded			[]EventTreasuryBountyAwarded
	//Treasury_BountyClaimed			[]EventTreasuryBountyClaimed
	//Treasury_BountyCanceled			[]EventTreasuryBountyCanceled
	//Treasury_BountyExtended			[]EventTreasuryBountyExtended
	//TechnicalMembership_Dummy		[]EventTechnicalMembershipDummy
	Swork_RegisterSuccess       []EventSworkRegisterSuccess
	Swork_WorksReportSuccess    []EventSworkWorksReportSuccess
	Swork_ABUpgradeSuccess      []EventSworkABUpgradeSuccess
	Swork_ChillSuccess          []EventSworkChillSuccess
	Swork_EnclaveUpgradeSuccess []EventSworkEnclaveUpgradeSuccess
	Market_StorageOrderSuccess  []EventMarketStorageOrderSuccess
	Market_RegisterSuccess      []EventMarketRegisterSuccess
	Market_PledgeSuccess        []EventMarketPledgeSuccess
	Market_SetAliasSuccess      []EventMarketSetAliasSuccess
	Market_PaysOrderSuccess     []EventMarketPaysOrderSuccess
	Candy_CandyIssued           []EventCandyCandyIssued
	Candy_CandyTransferred      []EventCandyCandyTransferred
	Candy_CandyBurned           []EventCandyCandyBurned
	Candy_BondEthSuccess        []EventCandyBondEthSuccess
}

type EventOffencesOffence struct {
	Phase          types.Phase
	Kind           types.Bytes16
	OpaqueTimeSlot types.Bytes
	Bool           bool
	Topics         []types.Hash
}

//type EventTreasuryBountyProposed struct {
//	Phase       types.Phase
//	BountyIndex types.U32
//	Topics      []types.Hash
//}
//
//type EventTreasuryBountyRejected struct {
//	Phase       types.Phase
//	BountyIndex types.U32
//	Balance 	types.U128
//	Topics      []types.Hash
//}
//
//type EventTreasuryBountyBecameActive struct {
//	Phase       types.Phase
//	BountyIndex types.U32
//	Topics      []types.Hash
//}
//
//type EventTreasuryBountyAwarded struct {
//	Phase       types.Phase
//	BountyIndex types.U32
//	Who         types.AccountID
//	Topics      []types.Hash
//}
//
//type EventTreasuryBountyClaimed struct {
//	Phase       types.Phase
//	BountyIndex types.U32
//	Balance 	types.U128
//	Who         types.AccountID
//	Topics      []types.Hash
//}
//
//type EventTreasuryBountyCanceled struct {
//	Phase       types.Phase
//	BountyIndex types.U32
//	Topics      []types.Hash
//}
//
//type EventTreasuryBountyExtended struct {
//	Phase       types.Phase
//	BountyIndex types.U32
//	Topics      []types.Hash
//}
//
//type EventTechnicalMembershipDummy struct {
//	Phase       types.Phase
//	PhantomData types.Null
//	Topics      []types.Hash
//}

type EventSworkRegisterSuccess struct {
	Phase         types.Phase
	Who           types.AccountID
	SworkerPubKey SworkerPubKey
	Topics        []types.Hash
}

type EventSworkWorksReportSuccess struct {
	Phase         types.Phase
	Who           types.AccountID
	SworkerPubKey SworkerPubKey
	Topics        []types.Hash
}

type EventSworkABUpgradeSuccess struct {
	Phase          types.Phase
	Who            types.AccountID
	SworkerPubKey1 SworkerPubKey
	SworkerPubKey2 SworkerPubKey
	Topics         []types.Hash
}

type EventSworkChillSuccess struct {
	Phase         types.Phase
	Who           types.AccountID
	SworkerPubKey SworkerPubKey
	Topics        []types.Hash
}

type EventSworkEnclaveUpgradeSuccess struct {
	Phase         types.Phase
	Who           types.AccountID
	SworkerPubKey SworkerPubKey
	Topics        []types.Hash
}

type EventMarketStorageOrderSuccess struct {
	Phase      types.Phase
	Who        types.AccountID
	SorderInfo struct {
		Who     types.AccountID
		Balance types.U128
	}
	SorderStatus SorderStatus
	Topics       []types.Hash
}

type EventMarketRegisterSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventMarketPledgeSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventMarketSetAliasSuccess struct {
	Phase      types.Phase
	Who        types.AccountID
	FileAlias1 FileAlias
	FileAlias2 FileAlias
	Topics     []types.Hash
}

type EventMarketPaysOrderSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventCandyCandyIssued struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type EventCandyCandyTransferred struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type EventCandyCandyBurned struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type EventCandyBondEthSuccess struct {
	Phase      types.Phase
	Who        types.AccountID
	EthAddress EthAddress
	Topics     []types.Hash
}

//----------------------
/*
https://github.com/crustio/crust-api/blob/4a0d5c49dc3d07d0ef8d838d8e5bed42c923e68f/src/services/index.ts
*/

type EthAddress types.Bytes
type FileAlias types.Bytes
type SworkerPubKey types.Bytes

type OrderStatus struct {
	Value string
}

func (d *OrderStatus) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	if b == 0 {
		d.Value = "Success"
		return nil
	}
	if b == 0 {
		d.Value = "Failed"
		return nil
	}
	if b == 0 {
		d.Value = "Pending"
		return nil
	}
	return fmt.Errorf("unkown order status index: %d", b)
}

type SorderStatus struct {
	CompletedOn types.BlockNumber
	ExpiredOn   types.BlockNumber
	Status      OrderStatus
	ClaimedAt   types.BlockNumber
}

func (d *SorderStatus) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.CompletedOn)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.ExpiredOn)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Status)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.ClaimedAt)
}
