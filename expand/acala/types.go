package acala

import (
	"fmt"
	"github.com/JFJun/stafi-substrate-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type AcalaEventRecords struct {
	base.BaseEventRecords
	Tokens_Transferred                              []EventTokensTransferred
	AcalaTreasury_Proposed                          []types.EventTreasuryProposed
	AcalaTreasury_Spending                          []types.EventTreasurySpending
	AcalaTreasury_Awarded                           []types.EventTreasuryAwarded
	AcalaTreasury_Rejected                          []types.EventTreasuryRejected
	AcalaTreasury_Burnt                             []types.EventTreasuryBurnt
	AcalaTreasury_Rollover                          []types.EventTreasuryRollover
	AcalaTreasury_Deposit                           []types.EventTreasuryDeposit
	AcalaTreasury_NewTip                            []types.EventTreasuryNewTip
	AcalaTreasury_TipClosing                        []types.EventTreasuryTipClosing
	AcalaTreasury_TipClosed                         []types.EventTreasuryTipClosed
	AcalaTreasury_TipRetracted                      []types.EventTreasuryTipRetracted
	AcalaTreasury_BountyProposed                    []base.EventTreasuryBountyProposed
	AcalaTreasury_BountyRejected                    []base.EventTreasuryBountyRejected
	AcalaTreasury_BountyBecameActive                []base.EventTreasuryBountyBecameActive
	AcalaTreasury_BountyAwarded                     []base.EventTreasuryBountyAwarded
	AcalaTreasury_BountyClaimed                     []base.EventTreasuryBountyClaimed
	AcalaTreasury_BountyCanceled                    []base.EventTreasuryBountyCanceled
	AcalaTreasury_BountyExtended                    []base.EventTreasuryBountyExtended
	Proxy_Announced                                 []EventProxyAnnounced
	GraduallyUpdate_GraduallyUpdateAdded            []EventGraduallyUpdateGraduallyUpdateAdded
	GraduallyUpdate_GraduallyUpdateCancelled        []EventGraduallyUpdateGraduallyUpdateCancelled
	GraduallyUpdate_Updated                         []EventGraduallyUpdateUpdated
	GeneralCouncil_Proposed                         []EventGeneralCouncilProposed
	GeneralCouncil_Voted                            []EventGeneralCouncilVoted
	GeneralCouncil_Approved                         []EventGeneralCouncilApproved
	GeneralCouncil_Disapproved                      []EventGeneralCouncilDisapproved
	GeneralCouncil_Executed                         []EventGeneralCouncilExecuted
	GeneralCouncil_MemberExecuted                   []EventGeneralCouncilMemberExecuted
	GeneralCouncil_Closed                           []EventGeneralCouncilClosed
	GeneralCouncilMembership_MemberAdded            []EventGeneralCouncilMembershipMemberAdded
	GeneralCouncilMembership_MemberRemoved          []EventGeneralCouncilMembershipMemberRemoved
	GeneralCouncilMembership_MembersSwapped         []EventGeneralCouncilMembershipMembersSwapped
	GeneralCouncilMembership_MembersReset           []EventGeneralCouncilMembershipMembersReset
	GeneralCouncilMembership_KeyChanged             []EventGeneralCouncilMembershipKeyChanged
	GeneralCouncilMembership_Dummy                  []EventGeneralCouncilMembershipDummy
	HonzonCouncil_Proposed                          []EventHonzonCouncilProposed
	HonzonCouncil_Voted                             []EventHonzonCouncilVoted
	HonzonCouncil_Approved                          []EventHonzonCouncilApproved
	HonzonCouncil_Disapproved                       []EventHonzonCouncilDisapproved
	HonzonCouncil_Executed                          []EventHonzonCouncilExecuted
	HonzonCouncil_MemberExecuted                    []EventHonzonCouncilMemberExecuted
	HonzonCouncil_Closed                            []EventHonzonCouncilClosed
	HonzonCouncilMembership_MemberAdded             []EventHonzonCouncilMembershipMemberAdded
	HonzonCouncilMembership_MemberRemoved           []EventHonzonCouncilMembershipMemberRemoved
	HonzonCouncilMembership_MembersSwapped          []EventHonzonCouncilMembershipMembersSwapped
	HonzonCouncilMembership_MembersReset            []EventHonzonCouncilMembershipMembersReset
	HonzonCouncilMembership_KeyChanged              []EventHonzonCouncilMembershipKeyChanged
	HonzonCouncilMembership_Dummy                   []EventHonzonCouncilMembershipDummy
	HomaCouncil_Proposed                            []EventHomaCouncilProposed
	HomaCouncil_Voted                               []EventHomaCouncilVoted
	HomaCouncil_Approved                            []EventHomaCouncilApproved
	HomaCouncil_Disapproved                         []EventHomaCouncilDisapproved
	HomaCouncil_Executed                            []EventHomaCouncilExecuted
	HomaCouncil_MemberExecuted                      []EventHomaCouncilMemberExecuted
	HomaCouncil_Closed                              []EventHomaCouncilClosed
	HomaCouncilMembership_MemberAdded               []EventHomaCouncilMembershipMemberAdded
	HomaCouncilMembership_MemberRemoved             []EventHomaCouncilMembershipMemberRemoved
	HomaCouncilMembership_MembersSwapped            []EventHomaCouncilMembershipMembersSwapped
	HomaCouncilMembership_MembersReset              []EventHomaCouncilMembershipMembersReset
	HomaCouncilMembership_KeyChanged                []EventHomaCouncilMembershipKeyChanged
	HomaCouncilMembership_Dummy                     []EventHomaCouncilMembershipDummy
	TechnicalCommitteeMembership_MemberAdded        []EventTechnicalCommitteeMembershipMemberAdded
	TechnicalCommitteeMembership_MemberRemoved      []EventTechnicalCommitteeMembershipMemberRemoved
	TechnicalCommitteeMembership_MembersSwapped     []EventTechnicalCommitteeMembershipMembersSwapped
	TechnicalCommitteeMembership_MembersReset       []EventTechnicalCommitteeMembershipMembersReset
	TechnicalCommitteeMembership_KeyChanged         []EventTechnicalCommitteeMembershipKeyChanged
	TechnicalCommitteeMembership_Dummy              []EventTechnicalCommitteeMembershipDummy
	Authority_Dispatched                            []EventAuthorityDispatched
	Authority_Scheduled                             []EventAuthorityScheduled
	Authority_FastTracked                           []EventAuthorityFastTracked
	Authority_Delayed                               []EventAuthorityDelayed
	Authority_Cancelled                             []EventAuthorityCancelled
	ElectionsPhragmen_NewTerm                       []EventElectionsPhragmenNewTerm
	ElectionsPhragmen_EmptyTerm                     []EventElectionsPhragmenEmptyTerm
	ElectionsPhragmen_MemberKicked                  []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_MemberRenounced               []EventElectionsPhragmenMemberRenounced
	ElectionsPhragmen_VoterReported                 []EventElectionsPhragmenVoterReported
	AcalaOracle_NewFeedData                         []EventAcalaOracleNewFeedData
	BandOracle_NewFeedData                          []EventBandOracleNewFeedData
	OperatorMembershipAcala_MemberAdded             []EventOperatorMembershipAcalaMemberAdded
	OperatorMembershipAcala_MemberRemoved           []EventOperatorMembershipAcalaMemberRemoved
	OperatorMembershipAcala_MembersSwapped          []EventOperatorMembershipAcalaMembersSwapped
	OperatorMembershipAcala_MembersReset            []EventOperatorMembershipAcalaMembersReset
	OperatorMembershipAcala_KeyChanged              []EventOperatorMembershipAcalaKeyChanged
	OperatorMembershipAcala_Dummy                   []EventOperatorMembershipAcalaDummy
	OperatorMembershipBand_MemberAdded              []EventOperatorMembershipBandMemberAdded
	OperatorMembershipBand_MemberRemoved            []EventOperatorMembershipBandMemberRemoved
	OperatorMembershipBand_MembersSwapped           []EventOperatorMembershipBandMembersSwapped
	OperatorMembershipBand_MembersReset             []EventOperatorMembershipBandMembersReset
	OperatorMembershipBand_KeyChanged               []EventOperatorMembershipBandKeyChanged
	OperatorMembershipBand_Dummy                    []EventOperatorMembershipBandDummy
	Auction_Bid                                     []EventAuctionBid
	Prices_LockPrice                                []EventPricesLockPrice
	Prices_UnlockPrice                              []EventPricesUnlockPrice
	Dex_AddLiquidity                                []EventDexAddLiquidity
	Dex_RemoveLiquidity                             []EventDexRemoveLiquidity
	Dex_Swap                                        []EventDexSwap
	AuctionManager_NewCollateralAuction             []EventAuctionManagerNewCollateralAuction
	AuctionManager_NewDebitAuction                  []EventAuctionManagerNewDebitAuction
	AuctionManager_NewSurplusAuction                []EventAuctionManagerNewSurplusAuction
	AuctionManager_CancelAuction                    []EventAuctionManagerCancelAuction
	AuctionManager_CollateralAuctionDealt           []EventAuctionManagerCollateralAuctionDealt
	AuctionManager_SurplusAuctionDealt              []EventAuctionManagerSurplusAuctionDealt
	AuctionManager_DebitAuctionDealt                []EventAuctionManagerDebitAuctionDealt
	AuctionManager_DEXTakeCollateralAuction         []EventAuctionManagerDEXTakeCollateralAuction
	Loans_PositionUpdated                           []EventLoansPositionUpdated
	Loans_ConfiscateCollateralAndDebit              []EventLoansConfiscateCollateralAndDebit
	Loans_TransferLoan                              []EventLoansTransferLoan
	Honzon_Authorization                            []EventHonzonAuthorization
	Honzon_UnAuthorization                          []EventHonzonUnAuthorization
	Honzon_UnAuthorizationAll                       []EventHonzonUnAuthorizationAll
	CdpTreasury_CollateralAuctionMaximumSizeUpdated []EventCdpTreasuryCollateralAuctionMaximumSizeUpdated
	CdpEngine_LiquidateUnsafeCDP                    []EventCdpEngineLiquidateUnsafeCDP
	CdpEngine_SettleCDPInDebit                      []EventCdpEngineSettleCDPInDebit
	CdpEngine_StabilityFeeUpdated                   []EventCdpEngineStabilityFeeUpdated
	CdpEngine_LiquidationRatioUpdated               []EventCdpEngineLiquidationRatioUpdated
	CdpEngine_LiquidationPenaltyUpdated             []EventCdpEngineLiquidationPenaltyUpdated
	CdpEngine_RequiredCollateralRatioUpdated        []EventCdpEngineRequiredCollateralRatioUpdated
	CdpEngine_MaximumTotalDebitValueUpdated         []EventCdpEngineMaximumTotalDebitValueUpdated
	CdpEngine_GlobalStabilityFeeUpdated             []EventCdpEngineGlobalStabilityFeeUpdated
	EmergencyShutdown_Shutdown                      []EventEmergencyShutdownShutdown
	EmergencyShutdown_OpenRefund                    []EventEmergencyShutdownOpenRefund
	EmergencyShutdown_Refund                        []EventEmergencyShutdownRefund
	StakingPool_MintLiquid                          []EventStakingPoolMintLiquid
	StakingPool_RedeemByUnbond                      []EventStakingPoolRedeemByUnbond
	StakingPool_RedeemByFreeUnbonded                []EventStakingPoolRedeemByFreeUnbonded
	StakingPool_RedeemByClaimUnbonding              []EventStakingPoolRedeemByClaimUnbonding
	Incentives_DepositDEXShare                      []EventIncentivesDepositDEXShare
	Incentives_WithdrawDEXShare                     []EventIncentivesWithdrawDEXShare
	AirDrop_Airdrop                                 []EventAirDropAirdrop
	AirDrop_UpdateAirdrop                           []EventAirDropUpdateAirdrop
	NFT_CreatedClass                                []EventNFTCreatedClass
	NFT_MintedToken                                 []EventNFTMintedToken
	NFT_TransferedToken                             []EventNFTTransferedToken
	NFT_BurnedToken                                 []EventNFTBurnedToken
	NFT_DestroyedClass                              []EventNFTDestroyedClass
	RenVmBridge_Minted                              []EventRenVmBridgeMinted
	RenVmBridge_Burnt                               []EventRenVmBridgeBurnt
	EVM_Log                                         []EventEVMLog
	EVM_Created                                     []EventEVMCreated
	EVM_CreatedFailed                               []EventEVMCreatedFailed
	EVM_Executed                                    []EventEVMExecuted
	EVM_ExecutedFailed                              []EventEVMExecutedFailed
	EVM_BalanceDeposit                              []EventEVMBalanceDeposit
	EVM_BalanceWithdraw                             []EventEVMBalanceWithdraw
}

/*
https://github.com/AcalaNetwork/acala.js/blob/aa6d9ac3c3b7f2b6b694ee7b28a08964e0490095/packages/type-definitions/src/runtime.ts
*/
type AuctionId types.U32
type Amount types.I128

/*
https://github.com/AcalaNetwork/acala.js/blob/f6e977e7c874181785a0389c993b71b134b65b1d/packages/type-definitions/src/nft.ts
*/
type ClassId types.U64
type TokenId types.U64

/*
https://github.com/AcalaNetwork/acala.js/blob/5865994dd416c02a08351f15b7b6c8edd21c1981/packages/type-definitions/src/primitives.ts
*/
type AirDropCurrencyId struct {
	Value string
}

/*
https://github.com/AcalaNetwork/acala.js/blob/a7d8072fe02bd9bde593888d54b3941e297b584b/packages/type-definitions/src/index.ts
*/
type PalletsOrigin struct {
	Value string
}

func (d *PalletsOrigin) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	if b > 58 {
		return fmt.Errorf("unknown enum index %d", b)
	}
	//  太多了，不想处理了，反正也没什么用
	return nil
}

/*
https://github.com/AcalaNetwork/acala.js/blob/f6e977e7c874181785a0389c993b71b134b65b1d/packages/type-definitions/src/cdpEngine.ts
*/
type LiquidationStrategy struct {
	Value string
}

func (d *LiquidationStrategy) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	if b == 0 {
		d.Value = "Auction"
		return nil
	}
	if b == 1 {
		d.Value = "Exchange"
		return nil
	}
	return fmt.Errorf("unknow enum index %d", b)
}

func (d *AirDropCurrencyId) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	if b == 0 {
		d.Value = "KAR"
		return nil
	}
	if b == 1 {
		d.Value = "ACA"
		return nil
	}
	return fmt.Errorf("unknow enum type %d", d)
}

/*
https://github.com/polkadot-js/api/blob/95c4f03bc3709c58b159623ec5c3c9794e077d08/packages/types/src/interfaces/staking/definitions.ts
*/
type EraIndex types.U32

/*
https://github.com/AcalaNetwork/acala.js/blob/f6e977e7c874181785a0389c993b71b134b65b1d/packages/type-definitions/src/support.ts
*/
type Rate base.FixedU128

type EventEVMBalanceWithdraw struct {
	Phase  types.Phase
	Who    types.AccountID
	H160   types.H160
	U256   types.U256
	Topics []types.Hash
}
type EventEVMBalanceDeposit struct {
	Phase  types.Phase
	Who    types.AccountID
	H160   types.H160
	U256   types.U256
	Topics []types.Hash
}
type EventEVMExecutedFailed struct {
	Phase  types.Phase
	H160   types.H160
	Topics []types.Hash
}
type EventEVMExecuted struct {
	Phase  types.Phase
	H160   types.H160
	Topics []types.Hash
}
type EventEVMCreatedFailed struct {
	Phase  types.Phase
	H160   types.H160
	Topics []types.Hash
}
type EventEVMCreated struct {
	Phase  types.Phase
	H160   types.H160
	Topics []types.Hash
}
type EventEVMLog struct {
	Phase  types.Phase
	Log    base.Log
	Topics []types.Hash
}
type EventRenVmBridgeBurnt struct {
	Phase       types.Phase
	Who         types.AccountID
	DestAddress base.VecU8L20
	Balance     types.U128
	Topics      []types.Hash
}
type EventRenVmBridgeMinted struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}
type EventNFTDestroyedClass struct {
	Phase   types.Phase
	From    types.AccountID
	ClassId ClassId
	To      types.AccountID
	Topics  []types.Hash
}
type EventNFTBurnedToken struct {
	Phase   types.Phase
	Who     types.AccountID
	ClassId ClassId
	TokenId TokenId
	Topics  []types.Hash
}
type EventNFTTransferedToken struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	ClassId ClassId
	TokenId TokenId
	Topics  []types.Hash
}
type EventNFTMintedToken struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	ClassId ClassId
	U32     types.U32
	Topics  []types.Hash
}
type EventNFTCreatedClass struct {
	Phase   types.Phase
	Who     types.AccountID
	ClassId ClassId
	Topics  []types.Hash
}
type EventAirDropUpdateAirdrop struct {
	Phase             types.Phase
	Who               types.AccountID
	AirDropCurrencyId AirDropCurrencyId
	Balance           types.U128
	Topics            []types.Hash
}
type EventAirDropAirdrop struct {
	Phase             types.Phase
	Who               types.AccountID
	AirDropCurrencyId AirDropCurrencyId
	Balance           types.U128
	Topics            []types.Hash
}
type EventIncentivesWithdrawDEXShare struct {
	Phase      types.Phase
	Who        types.AccountID
	CurrencyId base.CurrencyId
	Balance    types.U128
	Topics     []types.Hash
}
type EventIncentivesDepositDEXShare struct {
	Phase      types.Phase
	Who        types.AccountID
	CurrencyId base.CurrencyId
	Balance    types.U128
	Topics     []types.Hash
}
type EventStakingPoolRedeemByClaimUnbonding struct {
	Phase    types.Phase
	Who      types.AccountID
	EraIndex EraIndex
	Balance1 types.U128
	Balance2 types.U128
	Balance3 types.U128
	Topics   []types.Hash
}
type EventStakingPoolRedeemByFreeUnbonded struct {
	Phase    types.Phase
	Who      types.AccountID
	Balance1 types.U128
	Balance2 types.U128
	Balance3 types.U128
	Topics   []types.Hash
}
type EventStakingPoolRedeemByUnbond struct {
	Phase    types.Phase
	Who      types.AccountID
	Balance1 types.U128
	Balance2 types.U128
	Topics   []types.Hash
}
type EventStakingPoolMintLiquid struct {
	Phase    types.Phase
	Who      types.AccountID
	Balance1 types.U128
	Balance2 types.U128
	Topics   []types.Hash
}
type EventEmergencyShutdownRefund struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Value   []struct {
		CurrencyId base.CurrencyId
		Balance    types.U128
	}
	Topics []types.Hash
}
type EventEmergencyShutdownOpenRefund struct {
	Phase       types.Phase
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}
type EventEmergencyShutdownShutdown struct {
	Phase       types.Phase
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}
type EventCdpEngineGlobalStabilityFeeUpdated struct {
	Phase  types.Phase
	Rate   Rate
	Topics []types.Hash
}
type EventCdpEngineMaximumTotalDebitValueUpdated struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Balance    types.U128
	Topics     []types.Hash
}
type EventCdpEngineRequiredCollateralRatioUpdated struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Rate       Rate // Option<Ratio>
	Topics     []types.Hash
}
type EventCdpEngineLiquidationPenaltyUpdated struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Rate       Rate // Option<Ratio>
	Topics     []types.Hash
}
type EventCdpEngineLiquidationRatioUpdated struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Rate       Rate // Option<Ratio>
	Topics     []types.Hash
}
type EventCdpEngineStabilityFeeUpdated struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Rate       Rate // Option<Ratio>
	Topics     []types.Hash
}
type EventCdpEngineSettleCDPInDebit struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Who        types.AccountID
	Topics     []types.Hash
}
type EventCdpEngineLiquidateUnsafeCDP struct {
	Phase               types.Phase
	CurrencyId          base.CurrencyId
	Who                 types.AccountID
	Balance1            types.U128
	Balance2            types.U128
	LiquidationStrategy LiquidationStrategy
	Topics              []types.Hash
}
type EventCdpTreasuryCollateralAuctionMaximumSizeUpdated struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Balance    types.U128
	Topics     []types.Hash
}
type EventHonzonUnAuthorizationAll struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}
type EventHonzonUnAuthorization struct {
	Phase      types.Phase
	From       types.AccountID
	To         types.AccountID
	CurrencyId base.CurrencyId
	Topics     []types.Hash
}
type EventHonzonAuthorization struct {
	Phase      types.Phase
	From       types.AccountID
	To         types.AccountID
	CurrencyId base.CurrencyId
	Topics     []types.Hash
}
type EventLoansTransferLoan struct {
	Phase      types.Phase
	From       types.AccountID
	To         types.AccountID
	CurrencyId base.CurrencyId
	Topics     []types.Hash
}
type EventLoansConfiscateCollateralAndDebit struct {
	Phase      types.Phase
	Who        types.AccountID
	CurrencyId base.CurrencyId
	Balance1   types.U128
	Balance2   types.U128
	Topics     []types.Hash
}
type EventLoansPositionUpdated struct {
	Phase      types.Phase
	Who        types.AccountID
	CurrencyId base.CurrencyId
	Amount1    Amount
	Amount2    Amount
	Topics     []types.Hash
}
type EventAuctionManagerDEXTakeCollateralAuction struct {
	Phase      types.Phase
	AuctionId  AuctionId
	CurrencyId base.CurrencyId
	Balance1   types.U128
	Balance2   types.U128
	Topics     []types.Hash
}
type EventAuctionManagerDebitAuctionDealt struct {
	Phase     types.Phase
	AuctionId AuctionId
	Balance1  types.U128
	Who       types.AccountID
	Balance2  types.U128
	Topics    []types.Hash
}
type EventAuctionManagerSurplusAuctionDealt struct {
	Phase     types.Phase
	AuctionId AuctionId
	Balance1  types.U128
	Who       types.AccountID
	Balance2  types.U128
	Topics    []types.Hash
}
type EventAuctionManagerCollateralAuctionDealt struct {
	Phase      types.Phase
	AuctionId  AuctionId
	CurrencyId base.CurrencyId
	Balance1   types.U128
	Who        types.AccountID
	Balance2   types.U128
	Topics     []types.Hash
}
type EventAuctionManagerCancelAuction struct {
	Phase     types.Phase
	AuctionId AuctionId
	Topics    []types.Hash
}
type EventAuctionManagerNewSurplusAuction struct {
	Phase     types.Phase
	AuctionId AuctionId
	Balance   types.U128
	Topics    []types.Hash
}
type EventAuctionManagerNewDebitAuction struct {
	Phase     types.Phase
	AuctionId AuctionId
	Balance1  types.U128
	Balance2  types.U128
	Topics    []types.Hash
}
type EventAuctionManagerNewCollateralAuction struct {
	Phase      types.Phase
	AuctionId  AuctionId
	CurrencyId base.CurrencyId
	Balance1   types.U128
	Balance2   types.U128
	Topics     []types.Hash
}
type EventDexSwap struct {
	Phase       types.Phase
	Who         types.AccountID
	CurrencyIds []base.CurrencyId
	Balance1    types.U128
	Balance2    types.U128
	Topics      []types.Hash
}
type EventDexRemoveLiquidity struct {
	Phase       types.Phase
	Who         types.AccountID
	CurrencyId1 base.CurrencyId
	Balance1    types.U128
	CurrencyId2 base.CurrencyId
	Balance2    types.U128
	Balance3    types.U128
	Topics      []types.Hash
}
type EventDexAddLiquidity struct {
	Phase       types.Phase
	Who         types.AccountID
	CurrencyId1 base.CurrencyId
	Balance1    types.U128
	CurrencyId2 base.CurrencyId
	Balance2    types.U128
	Balance3    types.U128
	Topics      []types.Hash
}
type EventPricesUnlockPrice struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Topics     []types.Hash
}
type EventPricesLockPrice struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	Price      types.U128
	Topics     []types.Hash
}
type EventAuctionBid struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type EventOperatorMembershipBandDummy struct {
	Phase       types.Phase
	PhantomData types.Null
	Topics      []types.Hash
}
type EventOperatorMembershipBandKeyChanged struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventOperatorMembershipBandMembersReset struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventOperatorMembershipBandMembersSwapped struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventOperatorMembershipBandMemberRemoved struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventOperatorMembershipBandMemberAdded struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventOperatorMembershipAcalaDummy struct {
	Phase       types.Phase
	PhantomData types.Null
	Topics      []types.Hash
}
type EventOperatorMembershipAcalaKeyChanged struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventOperatorMembershipAcalaMembersReset struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventOperatorMembershipAcalaMembersSwapped struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventOperatorMembershipAcalaMemberRemoved struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventOperatorMembershipAcalaMemberAdded struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventBandOracleNewFeedData struct {
	Phase types.Phase
	Who   types.AccountID
	Value []struct {
		OracleKey   base.OracleKey
		OracleValue base.OracleValue
	}
	Topics []types.Hash
}
type EventAcalaOracleNewFeedData struct {
	Phase types.Phase
	Who   types.AccountID
	Value []struct {
		OracleKey   base.OracleKey
		OracleValue base.OracleValue
	}
	Topics []types.Hash
}
type EventElectionsPhragmenVoterReported struct {
	Phase  types.Phase
	From   types.AccountID
	To     types.AccountID
	Bool   bool
	Topics []types.Hash
}
type EventElectionsPhragmenMemberRenounced struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}
type EventElectionsPhragmenMemberKicked struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}
type EventElectionsPhragmenEmptyTerm struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventElectionsPhragmenNewTerm struct {
	Phase types.Phase
	Value []struct {
		Who     types.AccountID
		Balance types.U128
	}
	Topics []types.Hash
}
type EventAuthorityCancelled struct {
	Phase             types.Phase
	PalletsOrigin     PalletsOrigin
	ScheduleTaskIndex base.ScheduleTaskIndex
	Topics            []types.Hash
}
type EventAuthorityDelayed struct {
	Phase             types.Phase
	PalletsOrigin     PalletsOrigin
	ScheduleTaskIndex base.ScheduleTaskIndex
	BlockNumber       types.BlockNumber
	Topics            []types.Hash
}
type EventAuthorityFastTracked struct {
	Phase             types.Phase
	PalletsOrigin     PalletsOrigin
	ScheduleTaskIndex base.ScheduleTaskIndex
	BlockNumber       types.BlockNumber
	Topics            []types.Hash
}
type EventAuthorityScheduled struct {
	Phase             types.Phase
	PalletsOrigin     PalletsOrigin
	ScheduleTaskIndex base.ScheduleTaskIndex
	Topics            []types.Hash
}
type EventAuthorityDispatched struct {
	Phase          types.Phase
	DispatchResult types.DispatchResult
	Topics         []types.Hash
}
type EventTechnicalCommitteeMembershipDummy struct {
	Phase       types.Phase
	PhantomData types.Null
	Topics      []types.Hash
}
type EventTechnicalCommitteeMembershipKeyChanged struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventTechnicalCommitteeMembershipMembersReset struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventTechnicalCommitteeMembershipMembersSwapped struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventTechnicalCommitteeMembershipMemberRemoved struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventTechnicalCommitteeMembershipMemberAdded struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHomaCouncilMembershipDummy struct {
	Phase       types.Phase
	PhantomData types.Null
	Topics      []types.Hash
}
type EventHomaCouncilMembershipKeyChanged struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHomaCouncilMembershipMembersReset struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHomaCouncilMembershipMembersSwapped struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHomaCouncilMembershipMemberRemoved struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHomaCouncilMembershipMemberAdded struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHomaCouncilClosed struct {
	Phase        types.Phase
	Hash         types.Hash
	MemberCount1 base.MemberCount
	MemberCount2 base.MemberCount
	Topics       []types.Hash
}
type EventHomaCouncilMemberExecuted struct {
	Phase          types.Phase
	Hash           types.Hash
	DispatchResult types.DispatchResult
	Topics         []types.Hash
}
type EventHomaCouncilExecuted struct {
	Phase          types.Phase
	Hash           types.Hash
	DispatchResult types.DispatchResult
	Topics         []types.Hash
}
type EventHomaCouncilDisapproved struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventHomaCouncilApproved struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventHomaCouncilVoted struct {
	Phase        types.Phase
	Who          types.AccountID
	Hash         types.Hash
	Bool         bool
	MemberCount1 base.MemberCount
	MemberCount2 base.MemberCount
	Topics       []types.Hash
}
type EventHomaCouncilProposed struct {
	Phase         types.Phase
	Who           types.AccountID
	ProposalIndex types.U32
	Hash          types.Hash
	MemberCount   base.MemberCount
	Topics        []types.Hash
}
type EventHonzonCouncilMembershipDummy struct {
	Phase       types.Phase
	PhantomData types.Null
	Topics      []types.Hash
}
type EventHonzonCouncilMembershipKeyChanged struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHonzonCouncilMembershipMembersReset struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHonzonCouncilMembershipMembersSwapped struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHonzonCouncilMembershipMemberRemoved struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventHonzonCouncilMembershipMemberAdded struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventHonzonCouncilClosed struct {
	Phase        types.Phase
	Hash         types.Hash
	MemberCount1 base.MemberCount
	MemberCount2 base.MemberCount
	Topics       []types.Hash
}
type EventHonzonCouncilMemberExecuted struct {
	Phase          types.Phase
	Hash           types.Hash
	DispatchResult types.DispatchResult
	Topics         []types.Hash
}
type EventHonzonCouncilExecuted struct {
	Phase          types.Phase
	Hash           types.Hash
	DispatchResult types.DispatchResult
	Topics         []types.Hash
}
type EventHonzonCouncilDisapproved struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventHonzonCouncilApproved struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventHonzonCouncilVoted struct {
	Phase        types.Phase
	Who          types.AccountID
	Hash         types.Hash
	Bool         bool
	MemberCount1 base.MemberCount
	MemberCount2 base.MemberCount
	Topics       []types.Hash
}
type EventHonzonCouncilProposed struct {
	Phase         types.Phase
	Who           types.AccountID
	ProposalIndex types.U32
	Hash          types.U32
	MemberCount   base.MemberCount
	Topics        []types.Hash
}
type EventGeneralCouncilMembershipDummy struct {
	Phase       types.Phase
	PhantomData types.Null
	Topics      []types.Hash
}
type EventGeneralCouncilMembershipKeyChanged struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventGeneralCouncilMembershipMembersReset struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventGeneralCouncilMembershipMembersSwapped struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventGeneralCouncilMembershipMemberRemoved struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventGeneralCouncilMembershipMemberAdded struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventGeneralCouncilClosed struct {
	Phase        types.Phase
	Hash         types.Hash
	MemberCount1 base.MemberCount
	MemberCount2 base.MemberCount
	Topics       []types.Hash
}
type EventGeneralCouncilMemberExecuted struct {
	Phase          types.Phase
	Hash           types.Hash
	DispatchResult types.DispatchResult
	Topics         []types.Hash
}
type EventGeneralCouncilExecuted struct {
	Phase          types.Phase
	Hash           types.Hash
	DispatchResult types.DispatchResult
	Topics         []types.Hash
}
type EventGeneralCouncilDisapproved struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type EventGeneralCouncilApproved struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventGeneralCouncilVoted struct {
	Phase        types.Phase
	From         types.AccountID
	Hash         types.Hash
	Bool         bool
	MemberCount1 base.MemberCount
	MemberCount2 base.MemberCount
	Topics       []types.Hash
}
type EventGeneralCouncilProposed struct {
	Phase         types.Phase
	From          types.AccountID
	ProposalIndex types.U32
	Hash          types.Hash
	MemberCount   base.MemberCount
	Topics        []types.Hash
}
type EventGraduallyUpdateUpdated struct {
	Phase        types.Phase
	BlockNumber  types.BlockNumber
	StorageKey   base.StorageKey
	StorageValue base.StorageValue
	Topics       []types.Hash
}
type EventGraduallyUpdateGraduallyUpdateCancelled struct {
	Phase      types.Phase
	StorageKey base.StorageKey
	Topics     []types.Hash
}
type EventGraduallyUpdateGraduallyUpdateAdded struct {
	Phase         types.Phase
	StorageKey    base.StorageKey
	StorageValue1 base.StorageValue
	StorageValue2 base.StorageValue
	Topics        []types.Hash
}

type EventTokensTransferred struct {
	Phase      types.Phase
	CurrencyId base.CurrencyId
	From       types.AccountID
	To         types.AccountID
	Balance    types.U128
	Topics     []types.Hash
}

type EventProxyAnnounced struct {
	Phase  types.Phase
	From   types.AccountID
	To     types.AccountID
	Hash   types.Hash
	Topics []types.Hash
}
