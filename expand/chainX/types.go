package chainX

import (
	"fmt"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type ChainXEventRecords struct {
	types.EventRecords
	//write by flynn   扩展对chainX的支持
	XSystem_Blacklisted                         []EventXSystemBlacklisted
	XSystem_Unblacklisted                       []EventXSystemUnblacklisted
	XAssetsRegistrar_Registered                 []EventXAssetsRegistrarRegistered
	XAssetsRegistrar_Recovered                  []EventXAssetsRegistrarRecovered
	XAssetsRegistrar_Deregistered               []EventXAssetsRegistrarDeregistered
	XAssets_Moved                               []EventXAssetsMoved
	XAssets_Issued                              []EventXAssetsIssued
	XAssets_Destroyed                           []EventXAssetsDestroyed
	XAssets_SetBalance                          []EventXAssetsSetBalance
	XAssets_BalanceSet                          []EventXAssetsBalanceSet
	XAssets_SetRestrictions                     []EventXAssetsSetRestrictions
	XStaking_Minted                             []EventXStakingMinted
	XStaking_Slashed                            []EventXStakingSlashed
	XStaking_Bonded                             []EventXStakingBonded
	XStaking_Rebonded                           []EventXStakingReBonded
	XStaking_Unbonded                           []EventXStakingUnBonded
	XStaking_Claimed                            []EventXStakingClaimed
	XStaking_Withdrawn                          []EventXStakingWithdrawn
	XStaking_ForceChilled                       []EventXStakingForceChilled
	XStaking_ForceAllWithdrawn                  []EventXStakingForceAllWithdrawn
	XMiningAsset_Claimed                        []EventXMiningAssetClaimed
	XMiningAsset_Minted                         []EventXMiningAssetMinted
	XGatewayRecords_Deposited                   []EventXGatewayRecordsDeposited
	XGatewayRecords_WithdrawalCreated           []EventXGatewayRecordsWithdrawalCreated
	XGatewayRecords_WithdrawalProcessed         []EventXGatewayRecordsWithdrawalProcessed
	XGatewayRecords_WithdrawalRecovered         []EventXGatewayRecordsWithdrawalRecovered
	XGatewayRecords_WithdrawalCanceled          []EventXGatewayRecordsWithdrawalCanceled
	XGatewayRecords_WithdrawalFinished          []EventXGatewayRecordsWithdrawalFinished
	XGatewayCommon_SetTrusteeProps              []EventXGatewayCommonSetTrusteeProps
	XGatewayCommon_ReferralBinded               []EventXGatewayCommonReferralBinded
	XGatewayCommon_TrusteeSetChanged            []EventXGatewayCommonTrusteeSetChanged
	XGatewayBitcoin_HeaderInserted              []EventXGatewayBitcoinHeaderInserted
	XGatewayBitcoin_TxProcessed                 []EventXGatewayBitcoinTxProcessed
	XGatewayBitcoin_Deposited                   []EventXGatewayBitcoinDeposited
	XGatewayBitcoin_Withdrawn                   []EventXGatewayBitcoinWithdrawn
	XGatewayBitcoin_UnclaimedDeposit            []EventXGatewayBitcoinUnclaimedDeposit
	XGatewayBitcoin_PendingDepositRemoved       []EventXGatewayBitcoinPendingDepositRemoved
	XGatewayBitcoin_WithdrawalProposalCreated   []EventXGatewayBitcoinWithdrawalProposalCreated
	XGatewayBitcoin_WithdrawalProposalVoted     []EventXGatewayBitcoinWithdrawalProposalVoted
	XGatewayBitcoin_WithdrawalProposalDropped   []EventXGatewayBitcoinWithdrawalProposalDropped
	XGatewayBitcoin_WithdrawalProposalCompleted []EventXGatewayBitcoinWithdrawalProposalCompleted
	XGatewayBitcoin_WithdrawalFatalErr          []EventXGatewayBitcoinWithdrawalFatalErr
	XSpot_NewOrder                              []EventXSpotNewOrder
	XSpot_MakerOrderUpdated                     []EventXSpotMakerOrderUpdated
	XSpot_TakerOrderUpdated                     []EventXSpotTakerOrderUpdated
	XSpot_OrderExecuted                         []EventXSpotOrderExecuted
	XSpot_CanceledOrderUpdated                  []EventXSpotCanceledOrderUpdated
	XSpot_TradingPairAdded                      []EventXSpotTradingPairAdded
	XSpot_TradingPairUpdated                    []EventXSpotTradingPairUpdated
	XSpot_PriceFluctuationUpdated               []EventXSpotPriceFluctuationUpdated
	Currencies_Transferred                      []EventCurrenciesTransferred
	Currencies_BalanceUpdated                   []EventCurrenciesBalanceUpdated
	Currencies_Deposited                        []EventCurrenciesDeposited
	Currencies_Withdrawn                        []EventCurrenciesWithdrawn
	XTransactionFee_FeePaid                     []EventXTransactionFeeFeePaid
	Democracy_Blacklisted                       []EventDemocracyBlacklisted
	Elections_ElectionError                     []EventElectionsElectionError
}

type EventXMiningAssetMinted struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventXStakingForceAllWithdrawn struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventXAssetsBalanceSet struct {
	Phase     types.Phase
	AssetId   AssetId
	AccountId types.AccountID
	AssetType AssetType
	Balance   types.U128
	Topics    []types.Hash
}
type EventElectionsElectionError struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventDemocracyBlacklisted struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

func (d *ChainXEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return d.Balances_Transfer
}
func (d *ChainXEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return d.System_ExtrinsicSuccess
}
func (d *ChainXEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return d.System_ExtrinsicFailed
}

type EventXSystemBlacklisted struct {
	Phase     types.Phase
	AccountID types.AccountID
	Topics    []types.Hash
}

type EventXSystemUnblacklisted struct {
	Phase     types.Phase
	AccountID types.AccountID
	Topics    []types.Hash
}

type EventXAssetsRegistrarRegistered struct {
	Phase   types.Phase
	AssetId AssetId
	Bool    types.Bool
	Topics  []types.Hash
}

type EventXAssetsRegistrarRecovered struct {
	Phase   types.Phase
	AssetId AssetId
	Bool    types.Bool
	Topics  []types.Hash
}

type EventXAssetsRegistrarDeregistered struct {
	Phase   types.Phase
	AssetId AssetId
	Topics  []types.Hash
}

type EventXAssetsMoved struct {
	Phase      types.Phase
	AssetId    AssetId
	AccountId1 types.AccountID
	AssetType1 AssetType
	AccountId2 types.AccountID
	AssetType2 AssetType
	Balance    types.U128
	Topics     []types.Hash
}

type EventXAssetsIssued struct {
	Phase     types.Phase
	AssetId   AssetId
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}

type EventXAssetsDestroyed struct {
	Phase     types.Phase
	AssetId   AssetId
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventXAssetsSetBalance struct {
	Phase     types.Phase
	AssetId   AssetId
	AccountId types.AccountID
	AssetType AssetType
	Balance   types.U128
	Topics    []types.Hash
}

type EventXAssetsSetRestrictions struct {
	Phase             types.Phase
	AssetId           AssetId
	AssetRestrictions AssetRestrictions
	Topics            []types.Hash
}

type EventXStakingMinted struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}

type EventXStakingSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}

type EventXStakingBonded struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventXStakingReBonded struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	AccountId3 types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}
type EventXStakingUnBonded struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventXStakingClaimed struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventXStakingWithdrawn struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventXStakingForceChilled struct {
	Phase        types.Phase
	SessionIndex types.U32
	AccountIds   []types.AccountID
	Topics       []types.Hash
}

type EventXMiningAssetClaimed struct {
	Phase     types.Phase
	AccountId types.AccountID
	AssetId   AssetId
	Balance   types.U128
	Topics    []types.Hash
}

type EventXGatewayRecordsDeposited struct {
	Phase     types.Phase
	AccountId types.AccountID
	AssetId   AssetId
	Balance   types.U128
	Topics    []types.Hash
}

type EventXGatewayRecordsWithdrawalCreated struct {
	Phase              types.Phase
	WithdrawalRecordId WithdrawalRecordId
	WithdrawalRecord   WithdrawalRecord
	Topics             []types.Hash
}

type EventXGatewayRecordsWithdrawalProcessed struct {
	Phase              types.Phase
	WithdrawalRecordId WithdrawalRecordId
	Topics             []types.Hash
}

type EventXGatewayRecordsWithdrawalRecovered struct {
	Phase              types.Phase
	WithdrawalRecordId WithdrawalRecordId
	Topics             []types.Hash
}

type EventXGatewayRecordsWithdrawalCanceled struct {
	Phase              types.Phase
	WithdrawalRecordId WithdrawalRecordId
	WithdrawalState    WithdrawalState
	Topics             []types.Hash
}

type EventXGatewayRecordsWithdrawalFinished struct {
	Phase              types.Phase
	WithdrawalRecordId WithdrawalRecordId
	WithdrawalState    WithdrawalState
	Topics             []types.Hash
}

type EventXGatewayCommonSetTrusteeProps struct {
	Phase                        types.Phase
	AccountId                    types.AccountID
	Chain                        Chain
	GenericTrusteeIntentionProps GenericTrusteeIntentionProps
	Topics                       []types.Hash
}

type EventXGatewayCommonReferralBinded struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	Chain      Chain
	AccountId2 types.AccountID
	Topics     []types.Hash
}

type EventXGatewayCommonTrusteeSetChanged struct {
	Phase                     types.Phase
	Chain                     Chain
	U32                       types.U32
	GenericTrusteeSessionInfo GenericTrusteeSessionInfo
	Topics                    []types.Hash
}

type EventXGatewayBitcoinHeaderInserted struct {
	Phase  types.Phase
	H256   types.H256
	Topics []types.Hash
}

type EventXGatewayBitcoinTxProcessed struct {
	Phase      types.Phase
	H2561      types.H256
	H2562      types.H256
	BtcTxState BtcTxState
	Topics     []types.Hash
}

type EventXGatewayBitcoinDeposited struct {
	Phase     types.Phase
	H256      types.H256
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}

type EventXGatewayBitcoinWithdrawn struct {
	Phase   types.Phase
	H256    types.H256
	U32s    []types.U32
	Balance types.U128
	Topics  []types.Hash
}

type EventXGatewayBitcoinUnclaimedDeposit struct {
	Phase      types.Phase
	H256       types.H256
	BtcAddress BtcAddress
	Topics     []types.Hash
}
type EventXGatewayBitcoinPendingDepositRemoved struct {
	Phase      types.Phase
	AccountId  types.AccountID
	Balance    types.U128
	H256       types.H256
	BtcAddress BtcAddress
	Topics     []types.Hash
}

type EventXGatewayBitcoinWithdrawalProposalCreated struct {
	Phase     types.Phase
	AccountId types.AccountID
	U32s      []types.U32
	Topics    []types.Hash
}

type EventXGatewayBitcoinWithdrawalProposalVoted struct {
	Phase     types.Phase
	AccountId types.AccountID
	Bool      types.Bool
	Topics    []types.Hash
}

type EventXGatewayBitcoinWithdrawalProposalDropped struct {
	Phase  types.Phase
	U32    types.U32
	U32s   []types.U32
	Topics []types.Hash
}

type EventXGatewayBitcoinWithdrawalProposalCompleted struct {
	Phase  types.Phase
	H256   types.H256
	Topics []types.Hash
}

type EventXGatewayBitcoinWithdrawalFatalErr struct {
	Phase  types.Phase
	H2561  types.H256
	H2562  types.H256
	Topics []types.Hash
}

type EventXSpotNewOrder struct {
	Phase         types.Phase
	TradingPairId TradingPairId
	AccountId     types.AccountID
	Balance       types.U128
	Price         Price
	BlockNumber   types.BlockNumber
	Topics        []types.Hash
}

type EventXSpotMakerOrderUpdated struct {
	Phase         types.Phase
	TradingPairId TradingPairId
	AccountId     types.AccountID
	Balance       types.U128
	Price         Price
	BlockNumber   types.BlockNumber
	Topics        []types.Hash
}

type EventXSpotTakerOrderUpdated struct {
	Phase         types.Phase
	TradingPairId TradingPairId
	AccountId     types.AccountID
	Balance       types.U128
	Price         Price
	BlockNumber   types.BlockNumber
	Topics        []types.Hash
}
type EventXSpotOrderExecuted struct {
	Phase       types.Phase
	AccountId   types.AccountID
	Balance     types.U128
	BlockNumber types.BlockNumber
	Price       Price
	Topics      []types.Hash
}

type EventXSpotCanceledOrderUpdated struct {
	Phase         types.Phase
	TradingPairId TradingPairId
	AccountId     types.AccountID
	Balance       types.U128
	Price         Price
	BlockNumber   types.BlockNumber
	Topics        []types.Hash
}

type EventXSpotTradingPairAdded struct {
	Phase              types.Phase
	TradingPairProfile TradingPairProfile
	Topics             []types.Hash
}

type EventXSpotTradingPairUpdated struct {
	Phase              types.Phase
	TradingPairProfile TradingPairProfile
	Topics             []types.Hash
}

type EventXSpotPriceFluctuationUpdated struct {
	Phase              types.Phase
	TradingPairProfile TradingPairProfile
	PriceFluctuation   PriceFluctuation
	Topics             []types.Hash
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

type EventXTransactionFeeFeePaid struct {
	Phase      types.Phase
	AccountID1 types.AccountID
	Balance1   types.U128
	AccountID2 types.AccountID
	Balance2   types.U128
	Topics     []types.Hash
}

//  ------------------chainX struct -------------------//
// https://github.com/chainx-org/chainx.js-v2/blob/945a8c2f5bf7a6f058275736aad06b3a301d112f/packages/api/src/chainx/chainx-types.ts
type AssetId types.U32
type AssetType struct {
	EnumId int
	Value  string
}

func (at *AssetType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}
	switch b {
	case 0:
		at.Value = "Usable"
	case 1:
		at.Value = "Locked"
	case 2:
		at.Value = "Reserved"
	case 3:
		at.Value = "ReservedWithdrawal"
	case 4:
		at.Value = "ReservedDexSpot"
	default:
		return fmt.Errorf("unkown enum id %d", b)
	}
	return nil
}

type AssetRestrictions struct {
	Bits types.U32
}

func (a *AssetRestrictions) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&a.Bits)
	if err != nil {
		return err
	}
	return nil
}

type WithdrawalRecordId types.U32
type AddrStr types.Text
type Memo types.Text
type WithdrawalRecord struct {
	AssetId   AssetId
	Applicant types.AccountID
	Balance   types.U128
	Addr      AddrStr
	Ext       Memo
	Height    types.BlockNumber
}

func (w *WithdrawalRecord) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&w.AssetId)
	if err != nil {
		return err
	}
	err = decoder.Decode(&w.Applicant)
	if err != nil {
		return err
	}
	err = decoder.Decode(&w.Balance)
	if err != nil {
		return err
	}
	err = decoder.Decode(&w.Addr)
	if err != nil {
		return err
	}
	err = decoder.Decode(&w.Ext)
	if err != nil {
		return err
	}
	return decoder.Decode(&w.Height)
}

type WithdrawalState struct {
	EnumIdx int
	Value   string
}

func (at *WithdrawalState) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}
	switch b {
	case 0:
		at.Value = "Applying"
	case 1:
		at.Value = "Processing"
	case 2:
		at.Value = "NormalFinish"
	case 3:
		at.Value = "RootFinish"
	case 4:
		at.Value = "NormalCancel"
	case 5:
		at.Value = "RootCancel"
	default:
		return fmt.Errorf("unkown enum id %d", b)
	}
	return nil
}

type Chain struct {
	EnumIdx int
	Value   string
}

func (at *Chain) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}
	switch b {
	case 0:
		at.Value = "ChainX"
	case 1:
		at.Value = "Bitcoin"
	case 2:
		at.Value = "Ethereum"
	case 3:
		at.Value = "Polkadot"
	default:
		return fmt.Errorf("unkown enum id %d", b)
	}
	return nil
}

type GenericTrusteeIntentionProps struct {
	About      types.Text
	HotEntity  types.Bytes
	ColdEntity types.Bytes
}

func (d *GenericTrusteeIntentionProps) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.About)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.HotEntity)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.ColdEntity)
}

type GenericTrusteeSessionInfo struct {
	TrusteeList []types.AccountID
	Threshold   types.U16
	HotAddress  types.Bytes
	ColdAddress types.Bytes
}

func (d *GenericTrusteeSessionInfo) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.TrusteeList)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Threshold)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.HotAddress)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.ColdAddress)
}

type BtcTxResult struct {
	Value string
}

func (at *BtcTxResult) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}
	switch b {
	case 0:
		at.Value = "Success"
	case 1:
		at.Value = "Failed"
	default:
		return fmt.Errorf("unkown enum id %d", b)
	}
	return nil
}

type BtcTxType struct {
	Value string
}

func (at *BtcTxType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}
	switch b {
	case 0:
		at.Value = "Withdrawal"
	case 1:
		at.Value = "Deposit"
	case 2:
		at.Value = "HotAndCold"
	case 3:
		at.Value = "TrusteeTransition"
	case 4:
		at.Value = "Irrelevance"
	default:
		return fmt.Errorf("unkown enum id %d", b)
	}
	return nil
}

type BtcTxState struct {
	Result BtcTxResult
	TxType BtcTxType
}

func (d *BtcTxState) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.Result)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.TxType)
}

type BtcAddress types.Text

type TradingPairId types.U32

type Price types.U128

type CurrencyPair struct {
	Base  AssetId
	Quote AssetId
}

func (d *CurrencyPair) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.Base)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.Quote)
}

type TradingPairProfile struct {
	Id           TradingPairId
	CurrencyPair CurrencyPair
	PipDecimals  types.U32
	TickDecimals types.U32
	Tradable     types.Bool
}

func (d *TradingPairProfile) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.Id)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.CurrencyPair)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.PipDecimals)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.TickDecimals)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.Tradable)
}

type CurrencyId types.U32

type PriceFluctuation types.U32
