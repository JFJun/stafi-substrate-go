package rio

//
//import (
//	"fmt"
//	"github.com/JFJun/stafi-substrate-go/expand/base"
//	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
//	"github.com/stafiprotocol/go-substrate-rpc-client/types"
//)
//
////todo 目前还无法对接这个币，许多的types官方包里面都没有说明
//
//type RioEventRecords struct {
//	base.BaseEventRecords
//	Oracle_NewFeedData		[]EventOracleNewFeedData
//	RioAssets_Transferred		[]EventRioAssetsTransferred
//	RioAssets_Created		[]EventRioAssetsCreated
//	RioAssets_UpdateAssetRestriction		[]EventRioAssetsUpdateAssetRestriction
//	RioAssets_Revoke		[]EventRioAssetsRevoke
//	RioAssetsExt_Holder		[]EventRioAssetsExtHolder
//	RioPaymentFee_AccountChanged		[]EventRioPaymentFeeAccountChanged
//	RioPaymentFee_FeeDeposit		[]EventRioPaymentFeeFeeDeposit
//	RioGateway_AuthChanged		[]EventRioGatewayAuthChanged
//	RioGateway_SupportedAssetAdded		[]EventRioGatewaySupportedAssetAdded
//	RioGateway_SupportedAssetRemoved		[]EventRioGatewaySupportedAssetRemoved
//	RioGateway_WithdrawaFeeSetted		[]EventRioGatewayWithdrawaFeeSetted
//	RioGateway_NewDepositAddrInfoOfAssetId		[]EventRioGatewayNewDepositAddrInfoOfAssetId
//	RioGateway_NewDepositIndex		[]EventRioGatewayNewDepositIndex
//	RioGateway_MaxDepositCountSetted		[]EventRioGatewayMaxDepositCountSetted
//	RioGateway_NewDepositRecord		[]EventRioGatewayNewDepositRecord
//	RioGateway_NewPendingWithdrawRecord		[]EventRioGatewayNewPendingWithdrawRecord
//	RioGateway_WithdrawRebroadcasted		[]EventRioGatewayWithdrawRebroadcasted
//	RioGateway_WithdrawStatusChanged		[]EventRioGatewayWithdrawStatusChanged
//	RioGateway_UnsafeSetWithdrawState		[]EventRioGatewayUnsafeSetWithdrawState
//	RioGateway_UnsafeRemoveWithdrawRecord		[]EventRioGatewayUnsafeRemoveWithdrawRecord
//	RioPrices_LockPrice		[]EventRioPricesLockPrice
//	RioPrices_UnlockPrice		[]EventRioPricesUnlockPrice
//	RioRoot_ModifyManager		[]EventRioRootModifyManager
//	RioRoot_LockedRFuelIssued		[]EventRioRootLockedRFuelIssued
//}
//
//
//
//
//
//type EventOracleNewFeedData struct {
//	Phase    types.Phase
//	Who 		types.AccountID
//	Value 		[]struct{
//		OracleKey	OracleKey
//		OracleValue	OracleValue
//	}
//	Topics []types.Hash
//}
//type EventRioAssetsTransferred struct {
//	Phase    types.Phase
//	CurrencyId	base.CurrencyId
//	From 		types.AccountID
//	To 			types.AccountID
//	Balance		types.U128
//	Topics []types.Hash
//}
//
//type EventRioAssetsCreated struct {
//	Phase    types.Phase
//	CurrencyId	base.CurrencyId
//	Topics []types.Hash
//}
//
//type EventRioAssetsUpdateAssetRestriction struct {
//	Phase    types.Phase
//	CurrencyId	base.CurrencyId
//	Restrictions	Restrictions
//	Topics []types.Hash
//}
//type EventRioAssetsRevoke struct {
//	Phase    types.Phase
//	CurrencyId	base.CurrencyId
//	Topics []types.Hash
//}
//
//type EventRioAssetsExtHolder struct {
//	Phase    types.Phase
//	Who 		types.AccountID
//	Topics []types.Hash
//}
//type EventRioPaymentFeeAccountChanged struct {
//	Phase    types.Phase
//	Who 	types.AccountID
//	Topics []types.Hash
//}
//
//type EventRioPaymentFeeFeeDeposit struct {
//	Phase    types.Phase
//	Balance	types.U128
//	Topics []types.Hash
//}
//
//type EventRioGatewayAuthChanged struct {
//	Phase    types.Phase
//	Who 	types.AccountID
//	Auths	Auths
//	Topics []types.Hash
//}
//
//type EventRioGatewaySupportedAssetAdded struct {
//	Phase    types.Phase
//	Who 		types.AccountID
//	CurrencyId	base.CurrencyId
//	Balance 	types.U128
//	Topics []types.Hash
//}
//type EventRioGatewaySupportedAssetRemoved struct {
//	Phase    types.Phase
//	Who 	types.AccountID
//	CurrencyId	base.CurrencyId
//	Topics []types.Hash
//}
//
//type EventRioGatewayWithdrawaFeeSetted struct {
//	Phase    types.Phase
//	Who 	types.AccountID
//	CurrencyId	base.CurrencyId
//	Balance	types.U128
//	Topics []types.Hash
//}
//
//type EventRioGatewayNewDepositAddrInfoOfAssetId struct {
//	Phase    types.Phase
//	CurrencyId	base.CurrencyId
//	DepositAddrInfo	DepositAddrInfo
//	Topics []types.Hash
//}
//type EventRioGatewayNewDepositIndex struct {
//	Phase    types.Phase
//	Who 	types.AccountID
//	U64 	types.U64
//	Topics []types.Hash
//}
//type EventRioGatewayMaxDepositCountSetted struct {
//	Phase    types.Phase
//	U64		types.U64
//	Topics []types.Hash
//}
//
//type EventRioGatewayNewDepositRecord struct {
//	Phase    types.Phase
//	CurrencyId	base.CurrencyId
//	Who 		types.AccountID
//	Balance		types.U128
//	TxHash 		TxHash
//	Topics []types.Hash
//}
//type EventRioGatewayNewPendingWithdrawRecord struct {
//	Phase    types.Phase
//	U64 	types.U64
//	CurrencyId	base.CurrencyId
//	Who	 		types.AccountID
//	Balance 	types.Bytes128
//	Topics []types.Hash
//}
//
//type EventRioGatewayWithdrawRebroadcasted struct {
//	Phase    types.Phase
//	U64 	types.U64
//	Who 	types.AccountID
//	WithdrawState	WithdrawState
//	Topics []types.Hash
//}
//
//type EventRioGatewayWithdrawStatusChanged struct {
//	Phase    types.Phase
//	U64 		types.U64
//	Who 		types.AccountID
//	WithdrawState1 	WithdrawState
//	WithdrawState2	WithdrawState
//
//	Topics []types.Hash
//}
//type EventRioGatewayUnsafeSetWithdrawState struct {
//	Phase    types.Phase
//	U64 	types.U64
//	WithdrawState	WithdrawState
//	Topics []types.Hash
//}
//
//type EventRioGatewayUnsafeRemoveWithdrawRecord struct {
//	Phase    types.Phase
//	U64		types.U64
//	Topics []types.Hash
//}
//
//type EventRioPricesLockPrice struct {
//	Phase    types.Phase
//	CurrencyId	base.CurrencyId
//	Price	Price
//	Topics []types.Hash
//}
//
//type EventRioPricesUnlockPrice struct {
//	Phase    types.Phase
//	CurrencyId	base.CurrencyId
//	Topics []types.Hash
//}
//
//type EventRioRootModifyManager struct {
//	Phase    types.Phase
//	Who 	types.AccountID
//	Bool 	bool
//	Topics []types.Hash
//}
//type EventRioRootLockedRFuelIssued struct {
//	Phase    types.Phase
//	Who 		types.AccountID
//	Balance		types.U128
//	Topics []types.Hash
//}
////---------------------------------
//
//
///*
//https://github.com/RioDefi/riochain
//*/
//type Auths struct {
//	Value string
//}
//
//func (d *Auths)Decode(decoder scale.Decoder)error{
//	b,err := decoder.ReadOneByte()
//	if err != nil {
//		return err
//	}
//	if b==0 {
//		d.Value = "All"
//		return nil
//	}
//	if b==1 {
//		d.Value = "Deposit"
//		return nil
//	}
//	if b==2 {
//		d.Value = "Withdraw"
//		return nil
//	}
//	if b==3 {
//		d.Value = "Refund"
//		return nil
//	}
//	if b==4 {
//		d.Value = "Mark"
//		return nil
//	}
//	return fmt.Errorf("unknown type index %d",b)
//}
//
//type TxHash types.H256
//type Price 	types.U128
