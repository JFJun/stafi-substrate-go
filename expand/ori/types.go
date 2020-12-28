package ori

import (
	"github.com/JFJun/stafi-substrate-go/expand/base"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
)

type OrionEventRecords struct {
	base.BaseEventRecords
	EVM_Log             []EventEVMLog
	EVM_Created         []EventEVMCreated
	EVM_CreatedFailed   []EventEVMCreatedFailed
	EVM_Executed        []EventEVMExecuted
	EVM_ExecutedFailed  []EventEVMExecutedFailed
	EVM_BalanceDeposit  []EventEVMBalanceDeposit
	EVM_BalanceWithdraw []EventEVMBalanceWithdraw
	Issue_AddStage      []EventIssueAddStage
	Issue_ChangeOperate []EventIssueChangeOperate
	Gov_OpenProposal    []EventGovOpenProposal
	Gov_CloseProposal   []EventGovCloseProposal
	Gov_Voted           []EventGovVoted
}
type EventGovVoted struct {
	Phase     types.Phase
	U32       types.U32
	AccountId types.AccountID
	Bool      types.Bool
	Balance   types.U128
	Topics    []types.Hash
}
type EventGovCloseProposal struct {
	Phase  types.Phase
	U32    types.U32
	Bool   types.Bool
	Topics []types.Hash
}
type EventGovOpenProposal struct {
	Phase  types.Phase
	U32    types.U32
	Topics []types.Hash
}
type EventIssueChangeOperate struct {
	Phase  types.Phase
	Data1  types.Bytes
	Data2  types.Bytes
	Topics []types.Hash
}
type EventIssueAddStage struct {
	Phase   types.Phase
	Data    types.Bytes
	Balance types.U128
	Topics  []types.Hash
}
type EventEVMBalanceWithdraw struct {
	Phase     types.Phase
	AccountId types.AccountID
	H160      types.H160
	U256      types.U256
	Topics    []types.Hash
}

type EventEVMBalanceDeposit struct {
	Phase     types.Phase
	AccountId types.AccountID
	H160      types.H160
	U256      types.U256
	Topics    []types.Hash
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
