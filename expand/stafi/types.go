package stafi

import (
	"fmt"
	"github.com/JFJun/stafi-substrate-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type StafiEventRecords struct {
	base.BaseEventRecords
	BridgeCommon_ChainWhitelisted        []EventBridgeCommonChainWhitelisted
	BridgeCommon_FungibleTransfer        []EventBridgeCommonFungibleTransfer
	BridgeCommon_ChainFeesSet            []EventBridgeCommonChainFeesSet
	RBalances_Transfer                   []EventRBalancesTransfer
	RBalances_Minted                     []EventRBalancesMinted
	RBalances_Burned                     []EventRBalancesBurned
	RTokenRate_RateSet                   []EventRTokenRateRateSet
	RFis_NewPool                         []EventRFisNewPool
	RFis_CommissionUpdated               []EventRFisCommissionUpdated
	RFis_MaxValidatorCommissionUpdated   []EventRFisMaxValidatorCommissionUpdated
	RFis_PoolBalanceLimitUpdated         []EventRFisPoolBalanceLimitUpdated
	RFis_LiquidityBond                   []EventRFisLiquidityBond
	RFis_LiquidityUnBond                 []EventRFisLiquidityUnBond
	RFis_LiquidityWithdrawUnBond         []EventRFisLiquidityWithdrawUnBond
	RFis_ValidatorOnboard                []EventRFisValidatorOnboard
	RFis_ValidatorOffboard               []EventRFisValidatorOffboard
	RFis_TotalBondedBeforePayout         []EventRFisTotalBondedBeforePayout
	RFis_TotalBondedAfterPayout          []EventRFisTotalBondedAfterPayout
	RFis_NominateSwitchToggle            []EventRFisNominateSwitchToggle
	RFis_MinNominationNumSet             []EventRFisMinNominationNumSet
	RFis_MaxNominationNumSet             []EventRFisMaxNominationNumSet
	RFis_NominationUpdated               []EventRFisNominationUpdated
	BridgeCommon_RelayerThresholdChanged []EventBridgeCommonRelayerThresholdChanged
	BridgeCommon_RelayerAdded            []EventBridgeCommonRelayerAdded
	BridgeCommon_RelayerRemoved          []EventBridgeCommonRelayerRemoved
	BridgeCommon_ChainRemoved            []EventBridgeCommonChainRemoved
	BridgeCommon_VoteFor                 []EventBridgeCommonVoteFor
	BridgeCommon_VoteAgainst             []EventBridgeCommonVoteAgainst
	BridgeCommon_ProposalPassed          []EventBridgeCommonProposalPassed
	BridgeCommon_ProposalCancelled       []EventBridgeCommonProposalCancelled
	BridgeCommon_ProposalExecuted        []EventBridgeCommonProposalExecuted
	RFis_ValidatorPaidout                []EventRFisValidatorPaidout
}
type EventRFisValidatorPaidout struct {
	Phase      types.Phase
	EraIndex   EraIndex
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Bool       types.Bool
	Topics     []types.Hash
}
type EventBridgeCommonProposalExecuted struct {
	Phase        types.Phase
	ChainId      ChainId
	DepositNonce DepositNonce
	Topics       []types.Hash
}
type EventBridgeCommonProposalCancelled struct {
	Phase        types.Phase
	ChainId      ChainId
	DepositNonce DepositNonce
	Topics       []types.Hash
}
type EventBridgeCommonProposalPassed struct {
	Phase        types.Phase
	ChainId      ChainId
	DepositNonce DepositNonce
	Topics       []types.Hash
}
type EventBridgeCommonVoteAgainst struct {
	Phase        types.Phase
	ChainId      ChainId
	DepositNonce DepositNonce
	AccountId    types.AccountID
	Topics       []types.Hash
}
type EventBridgeCommonVoteFor struct {
	Phase        types.Phase
	ChainId      ChainId
	DepositNonce DepositNonce
	AccountId    types.AccountID
	Topics       []types.Hash
}
type EventBridgeCommonChainRemoved struct {
	Phase   types.Phase
	ChainId ChainId
	Topics  []types.Hash
}
type EventBridgeCommonRelayerRemoved struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventBridgeCommonRelayerAdded struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventBridgeCommonRelayerThresholdChanged struct {
	Phase  types.Phase
	U32    types.U32
	Topics []types.Hash
}
type EventRFisNominationUpdated struct {
	Phase      types.Phase
	EraIndex   EraIndex
	AccountIds []types.AccountID
	AccountId1 types.AccountID
	Topics     []types.Hash
}
type EventRFisMaxNominationNumSet struct {
	Phase  types.Phase
	U8     types.U8
	Topics []types.Hash
}
type EventRFisMinNominationNumSet struct {
	Phase  types.Phase
	U8     types.U8
	Topics []types.Hash
}
type EventRFisNominateSwitchToggle struct {
	Phase  types.Phase
	Bool   types.Bool
	Topics []types.Hash
}
type EventRFisTotalBondedAfterPayout struct {
	Phase    types.Phase
	EraIndex EraIndex
	Balance  types.U128
	Topics   []types.Hash
}
type EventRFisTotalBondedBeforePayout struct {
	Phase    types.Phase
	EraIndex EraIndex
	Balance  types.U128
	Topics   []types.Hash
}
type EventRFisValidatorOffboard struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID

	Topics []types.Hash
}
type EventRFisValidatorOnboard struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Topics     []types.Hash
}
type EventRFisLiquidityWithdrawUnBond struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}
type EventRFisLiquidityUnBond struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	U1281      types.U128
	U1282      types.U128
	Balance1   types.U128
	Topics     []types.Hash
}
type EventRFisLiquidityBond struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Balance1   types.U128
	U128       types.U128
	Topics     []types.Hash
}
type EventRFisPoolBalanceLimitUpdated struct {
	Phase    types.Phase
	Balance1 types.U128
	Balance2 types.U128
	Topics   []types.Hash
}
type EventRFisMaxValidatorCommissionUpdated struct {
	Phase    types.Phase
	Perbill1 base.Perbill
	Perbill2 base.Perbill
	Topics   []types.Hash
}
type EventRFisCommissionUpdated struct {
	Phase    types.Phase
	Perbill1 base.Perbill
	Perbill2 base.Perbill
	Topics   []types.Hash
}
type EventRFisNewPool struct {
	Phase     types.Phase
	VecU8     types.Bytes
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventRTokenRateRateSet struct {
	Phase    types.Phase
	RateType RateType
	Topics   []types.Hash
}
type EventRBalancesBurned struct {
	Phase     types.Phase
	AccountId types.AccountID
	RSymbol   RSymbol
	U128      types.U128
	Topics    []types.Hash
}
type EventRBalancesMinted struct {
	Phase     types.Phase
	AccountId types.AccountID
	RSymbol   RSymbol
	U128      types.U128
	Topics    []types.Hash
}
type EventRBalancesTransfer struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	RSymbol    RSymbol
	U128       types.U128
	Topics     []types.Hash
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
type RateType types.U64
type EraIndex types.U32

/*
https://github.com/stafiprotocol/stafi-bootstrap/blob/master/types.json
*/
type RSymbol struct {
	Value string
}

func (d *RSymbol) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	if b == 0 {
		d.Value = "RFIS"
	}
	return fmt.Errorf("unKnow enum index: %d", b)
}
