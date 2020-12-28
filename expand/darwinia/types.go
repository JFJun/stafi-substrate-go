package darwinia

import (
	"github.com/JFJun/stafi-substrate-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type DarwiniaEventRecords struct {
	base.BaseEventRecords
	Kton_Endowed                                    []EventKtonEndowed
	Kton_DustLost                                   []EventKtonDustLost
	Kton_Transfer                                   []EventKtonTransfer
	Kton_BalanceSet                                 []EventKtonBalanceSet
	Kton_Deposit                                    []EventKtonDeposit
	Kton_Reserved                                   []EventKtonReserved
	Kton_Unreserved                                 []EventKtonUnreserved
	Kton_ReserveRepatriated                         []EventKtonReserveRepatriated
	Staking_Slash                                   []EventStakingSlash
	Staking_BondRing                                []EventStakingBondRing
	Staking_BondKton                                []EventStakingBondKton
	Staking_UnbondRing                              []EventStakingUnbondRing
	Staking_UnbondKton                              []EventStakingUnbondKton
	Staking_ClaimDepositsWithPunish                 []EventStakingClaimDepositsWithPunish
	ElectionsPhragmen_NewTerm                       []EventElectionsPhragmenNewTerm
	ElectionsPhragmen_EmptyTerm                     []EventElectionsPhragmenEmptyTerm
	ElectionsPhragmen_ElectionError                 []EventElectionsPhragmenElectionError
	ElectionsPhragmen_MemberKicked                  []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_MemberRenounced               []EventElectionsPhragmenMemberRenounced
	ElectionsPhragmen_VoterReported                 []EventElectionsPhragmenVoterReported
	Claims_Claimed                                  []EventClaimsClaimed
	EthereumBacking_RedeemRing                      []EventEthereumBackingRedeemRing
	EthereumBacking_RedeemKton                      []EventEthereumBackingRedeemKton
	EthereumBacking_RedeemDeposit                   []EventEthereumBackingRedeemDeposit
	EthereumRelay_Affirmed                          []EventEthereumRelayAffirmed
	EthereumRelay_DisputedAndAffirmed               []EventEthereumRelayDisputedAndAffirmed
	EthereumRelay_Extended                          []EventEthereumRelayExtended
	EthereumRelay_NewRound                          []EventEthereumRelayNewRound
	EthereumRelay_GameOver                          []EventEthereumRelayGameOver
	EthereumRelay_RemoveConfirmedParcel             []EventEthereumRelayRemoveConfirmedParcel
	EthereumRelay_VerifyReceipt                     []EventEthereumRelayVerifyReceipt
	EthereumRelay_Pended                            []EventEthereumRelayPended
	EthereumRelay_GuardVoted                        []EventEthereumRelayGuardVoted
	EthereumRelay_PendingRelayHeaderParcelConfirmed []EventEthereumRelayPendingRelayHeaderParcelConfirmed
	EthereumRelay_PendingRelayHeaderParcelRejected  []EventEthereumRelayPendingRelayHeaderParcelRejected
	Treasury_Awarded                                []EventTreasuryAwarded
	Treasury_Rejected                               []EventTreasuryRejected
	Treasury_Burnt                                  []EventTreasuryBurnt
	Treasury_Rollover                               []EventTreasuryRollover
	Treasury_DepositRing                            []EventTreasuryDepositRing
	Treasury_DepositKton                            []EventTreasuryDepositKton
	Treasury_BountyProposed                         []EventTreasuryBountyProposed
	Treasury_BountyRejected                         []EventTreasuryBountyRejected
	Treasury_BountyBecameActive                     []EventTreasuryBountyBecameActive
	Treasury_BountyAwarded                          []EventTreasuryBountyAwarded
	Treasury_BountyClaimed                          []EventTreasuryBountyClaimed
	Treasury_BountyCanceled                         []EventTreasuryBountyCanceled
	Treasury_BountyExtended                         []EventTreasuryBountyExtended
	Proxy_Announced                                 []EventProxyAnnounced
	Multisig_NewMultisig                            []types.EventMultisigNewMultisig
	Multisig_MultisigApproval                       []types.EventMultisigApproval
	Multisig_MultisigExecuted                       []types.EventMultisigExecuted
	Multisig_MultisigCancelled                      []types.EventMultisigCancelled
	CrabIssuing_DummyEvent                          []EventCrabIssuingDummyEvent
}

func (d DarwiniaEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return d.Balances_Transfer
}

func (d DarwiniaEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return d.System_ExtrinsicSuccess
}

func (d DarwiniaEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return d.System_ExtrinsicFailed
}

type EventKtonEndowed struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}
type EventKtonDustLost struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type EventKtonTransfer struct {
	Phase  types.Phase
	From   types.AccountID
	To     types.AccountID
	Value  types.U128
	Topics []types.Hash
}
type EventKtonBalanceSet struct {
	Phase    types.Phase
	Who      types.AccountID
	Free     types.U128
	Reserved types.U128
	Topics   []types.Hash
}
type EventKtonDeposit struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}
type EventKtonReserved struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}
type EventKtonUnreserved struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}
type EventKtonReserveRepatriated struct {
	Phase             types.Phase
	From              types.AccountID
	To                types.AccountID
	Balance           types.U128
	DestinationStatus types.BalanceStatus
	Topics            []types.Hash
}

// EventStakingSlash is emitted when one validator (and its nominators) has been slashed by the given amount
type EventStakingSlash struct {
	Phase       types.Phase
	AccountID   types.AccountID
	RingBalance types.U128
	KtonBalance types.U128
	Topics      []types.Hash
}

type EventStakingBondRing struct {
	Phase       types.Phase
	RingBalance types.U128
	TsInMs1     types.U64
	TsInMs2     types.U64
	Topics      []types.Hash
}

type EventStakingBondKton struct {
	Phase       types.Phase
	KtonBalance types.U128
	Topics      []types.Hash
}

type EventStakingUnbondRing struct {
	Phase       types.Phase
	RingBalance types.U128
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}

type EventStakingUnbondKton struct {
	Phase       types.Phase
	KtonBalance types.U128
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}

type EventStakingClaimDepositsWithPunish struct {
	Phase       types.Phase
	Who         types.AccountID
	KtonBalance types.U128
	Topics      []types.Hash
}
type EventElectionsPhragmenNewTerm struct {
	Phase  types.Phase
	ABs    []AccountBalance
	Topics []types.Hash
}

type EventElectionsPhragmenEmptyTerm struct {
	Phase types.Phase

	Topics []types.Hash
}

type EventElectionsPhragmenElectionError struct {
	Phase types.Phase

	Topics []types.Hash
}

type EventElectionsPhragmenMemberKicked struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventElectionsPhragmenMemberRenounced struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventElectionsPhragmenVoterReported struct {
	Phase  types.Phase
	Who    types.AccountID
	ID     types.AccountID
	Bool   bool
	Topics []types.Hash
}

type EventClaimsClaimed struct {
	Phase       types.Phase
	Who         types.AccountID
	AddressT    base.VecU8L20 //[u8;20]
	RingBalance types.U128
	Topics      []types.Hash
}

type EventEthereumBackingRedeemRing struct {
	Phase                    types.Phase
	Who                      types.AccountID
	Balance                  types.U128
	EthereumTransactionIndex types.H256
	Topics                   []types.Hash
}

type EventEthereumBackingRedeemKton struct {
	Phase                    types.Phase
	Who                      types.AccountID
	Balance                  types.U128
	EthereumTransactionIndex types.H256
	Topics                   []types.Hash
}

type EventEthereumBackingRedeemDeposit struct {
	Phase                    types.Phase
	Who                      types.AccountID
	DepositId                types.U256
	RingBalance              types.U128
	EthereumTransactionIndex types.H256
	Topics                   []types.Hash
}

type EventEthereumRelayAffirmed struct {
	Phase              types.Phase
	Who                types.AccountID
	RelayAffirmationId RelayAffirmationId
	Topics             []types.Hash
}

type EventEthereumRelayDisputedAndAffirmed struct {
	Phase              types.Phase
	Who                types.AccountID
	RelayAffirmationId RelayAffirmationId
	Topics             []types.Hash
}

type EventEthereumRelayExtended struct {
	Phase              types.Phase
	Who                types.AccountID
	RelayAffirmationId RelayAffirmationId
	Topics             []types.Hash
}

type EventEthereumRelayNewRound struct {
	Phase                types.Phase
	EthereumBlockNumber  types.U64
	EthereumBlockNumbers []types.U64
	Topics               []types.Hash
}

type EventEthereumRelayGameOver struct {
	Phase               types.Phase
	EthereumBlockNumber types.U64
	Topics              []types.Hash
}

type EventEthereumRelayRemoveConfirmedParcel struct {
	Phase               types.Phase
	EthereumBlockNumber types.U64
	Topics              []types.Hash
}

type EventEthereumRelayVerifyReceipt struct {
	Phase               types.Phase
	Who                 types.AccountID
	EthereumReceipt     EthereumReceipt
	EthereumBlockNumber types.U64
	Topics              []types.Hash
}

type EventEthereumRelayPended struct {
	Phase               types.Phase
	EthereumBlockNumber types.U64
	Topics              []types.Hash
}

type EventEthereumRelayGuardVoted struct {
	Phase               types.Phase
	EthereumBlockNumber types.U64
	Bool                bool
	Topics              []types.Hash
}

type EventEthereumRelayPendingRelayHeaderParcelConfirmed struct {
	Phase               types.Phase
	EthereumBlockNumber types.U64
	Data                types.Bytes
	Topics              []types.Hash
}

type EventEthereumRelayPendingRelayHeaderParcelRejected struct {
	Phase               types.Phase
	EthereumBlockNumber types.U64
	Topics              []types.Hash
}

type EventTreasuryAwarded struct {
	Phase         types.Phase
	ProposalIndex types.U32
	RingBalance   types.U128
	KtonBalance   types.U128
	Who           types.AccountID
	Topics        []types.Hash
}

type EventTreasuryRejected struct {
	Phase         types.Phase
	ProposalIndex types.U32
	RingBalance   types.U128
	KtonBalance   types.U128
	Topics        []types.Hash
}

type EventTreasuryBurnt struct {
	Phase       types.Phase
	RingBalance types.U128
	KtonBalance types.U128
	Topics      []types.Hash
}

type EventTreasuryRollover struct {
	Phase       types.Phase
	RingBalance types.U128
	KtonBalance types.U128
	Topics      []types.Hash
}

type EventTreasuryDepositRing struct {
	Phase       types.Phase
	RingBalance types.U128
	Topics      []types.Hash
}

type EventTreasuryDepositKton struct {
	Phase       types.Phase
	KtonBalance types.U128
	Topics      []types.Hash
}

type EventTreasuryBountyProposed struct {
	Phase       types.Phase
	BountyIndex BountyIndex
	Topics      []types.Hash
}

type EventTreasuryBountyRejected struct {
	Phase       types.Phase
	BountyIndex BountyIndex
	RingBalance types.U128
	Topics      []types.Hash
}

type EventTreasuryBountyBecameActive struct {
	Phase       types.Phase
	BountyIndex BountyIndex
	Topics      []types.Hash
}

type EventTreasuryBountyAwarded struct {
	Phase       types.Phase
	BountyIndex BountyIndex
	Who         types.AccountID
	Topics      []types.Hash
}

type EventTreasuryBountyClaimed struct {
	Phase       types.Phase
	BountyIndex BountyIndex
	RingBalance types.U128
	Who         types.AccountID
	Topics      []types.Hash
}

type EventTreasuryBountyCanceled struct {
	Phase       types.Phase
	BountyIndex BountyIndex
	Topics      []types.Hash
}

type EventTreasuryBountyExtended struct {
	Phase       types.Phase
	BountyIndex BountyIndex
	Topics      []types.Hash
}

type EventProxyAnnounced struct {
	Phase  types.Phase
	Who    types.AccountID
	ID     types.AccountID
	Hash   types.Hash
	Topics []types.Hash
}

type EventCrabIssuingDummyEvent struct {
	Phase       types.Phase
	Who         types.AccountID
	RingBalance types.U128
	MappedRing  types.U128
	Topics      []types.Hash
}

//---------------------
type AccountBalance struct {
	Who     types.AccountID
	Balance types.U128
}

type EthereumReceipt struct {
	GasUsed  types.U256
	LogBloom base.VecU8L256
	Logs     []types.Null
	Outcome  types.Null
}

func (d *EthereumReceipt) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.GasUsed)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.LogBloom)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Logs)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.Outcome)
}

type RelayAffirmationId struct {
	RelayHeaderId types.U64
	Round         types.U32
	Index         types.U32
}

func (d *RelayAffirmationId) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.RelayHeaderId)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Round)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.Index)
}

type BountyIndex types.U32
