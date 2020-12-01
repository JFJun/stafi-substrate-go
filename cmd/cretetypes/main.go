package main

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/stafi-substrate-go/client"
)

func main() {
	url := "http://fis.rylink.io:31833"
	err := createTypes("fis", url)
	if err != nil {
		panic(err)
	}
}

func createBaseTypesJson(url string) error {
	c, err := client.New(url)
	if err != nil {
		return err
	}
	tj := new(TypesJson)
	tj.CoinName = "base"
	tj.SpecVersion = c.SpecVersion
	tj.ChainName = c.ChainName

	for _, mod := range c.Meta.AsMetadataV11.Modules {
		tj.Version = "v11"
		if mod.HasEvents {
			for _, event := range mod.Events {
				typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
				if IsExist(typeName) {
					var t EventType
					t.Name = typeName
					for _, arg := range event.Args {
						t.SubTypes = append(t.SubTypes, string(arg))
					}
					tj.Types = append(tj.Types, t)
				}
			}
		}
	}

	for _, mod := range c.Meta.AsMetadataV12.Modules {
		tj.Version = "v12"
		if mod.HasEvents {
			for _, event := range mod.Events {
				typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
				if IsExist(typeName) {
					var t EventType
					t.Name = typeName
					for _, arg := range event.Args {
						t.SubTypes = append(t.SubTypes, string(arg))
					}
					tj.Types = append(tj.Types, t)
				}
			}
		}
	}
	d, err := json.Marshal(tj)
	if err != nil {
		return err
	}
	fmt.Println(string(d))
	return nil
}

func createTypes(coinName, url string) error {
	c, err := client.New(url)
	if err != nil {
		return err
	}
	tj := new(TypesJson)
	tj.CoinName = coinName
	tj.SpecVersion = c.SpecVersion
	tj.ChainName = c.ChainName

	for _, mod := range c.Meta.AsMetadataV11.Modules {
		tj.Version = "v11"
		if mod.HasEvents {
			for _, event := range mod.Events {
				typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
				if !IsExist(typeName) {
					var t EventType
					t.Name = typeName
					for _, arg := range event.Args {
						t.SubTypes = append(t.SubTypes, string(arg))
					}
					tj.Types = append(tj.Types, t)
				}
			}
		}
	}

	for _, mod := range c.Meta.AsMetadataV12.Modules {
		tj.Version = "v12"
		if mod.HasEvents {
			for _, event := range mod.Events {
				typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
				if !IsExist(typeName) {
					var t EventType
					t.Name = typeName
					for _, arg := range event.Args {
						t.SubTypes = append(t.SubTypes, string(arg))
					}
					tj.Types = append(tj.Types, t)
				}
			}
		}
	}

	d, err := json.Marshal(tj)
	if err != nil {
		return err
	}
	fmt.Println(string(d))
	return nil
}

var existTypes = []string{
	"Balances_Endowed",
	"Balances_DustLost",
	"Balances_Transfer",
	"Balances_BalanceSet",
	"Balances_Deposit",
	"Balances_Reserved",
	"Balances_Unreserved",
	"Balances_ReservedRepatriated",
	"Grandpa_NewAuthorities",
	"Grandpa_Paused",
	"Grandpa_Resumed",
	"ImOnline_HeartbeatReceived",
	"ImOnline_AllGood",
	"ImOnline_SomeOffline",
	"Indices_IndexAssigned",
	"Indices_IndexFreed",
	"Indices_IndexFrozen",
	"Offences_Offence",
	"Session_NewSession",
	"Staking_EraPayout",
	"Staking_Reward",
	"Staking_Slash",
	"Staking_OldSlashingReportDiscarded",
	"Staking_StakingElection",
	"Staking_SolutionStored",
	"Staking_Bonded",
	"Staking_Unbonded",
	"Staking_Withdrawn",
	"System_ExtrinsicSuccess",
	"System_ExtrinsicFailed",
	"System_CodeUpdated",
	"System_NewAccount",
	"System_KilledAccount",
	"Assets_Issued",
	"Assets_Transferred",
	"Assets_Destroyed",
	"Democracy_Proposed",
	"Democracy_Tabled",
	"Democracy_ExternalTabled",
	"Democracy_Started",
	"Democracy_Passed",
	"Democracy_NotPassed",
	"Democracy_Cancelled",
	"Democracy_Executed",
	"Democracy_Delegated",
	"Democracy_Undelegated",
	"Democracy_Vetoed",
	"Democracy_PreimageNoted",
	"Democracy_PreimageUsed",
	"Democracy_PreimageInvalid",
	"Democracy_PreimageMissing",
	"Democracy_PreimageReaped",
	"Democracy_Unlocked",
	"Council_Proposed",
	"Council_Voted",
	"Council_Approved",
	"Council_Disapproved",
	"Council_Executed",
	"Council_MemberExecuted",
	"Council_Closed",
	"TechnicalCommittee_Proposed",
	"TechnicalCommittee_Voted",
	"TechnicalCommittee_Approved",
	"TechnicalCommittee_Disapproved",
	"TechnicalCommittee_Executed",
	"TechnicalCommittee_MemberExecuted",
	"TechnicalCommittee_Closed",
	"TechnicalMembership_MemberAdded",
	"TechnicalMembership_MemberRemoved",
	"TechnicalMembership_MembersSwapped",
	"TechnicalMembership_MembersReset",
	"TechnicalMembership_KeyChanged",
	"Elections_NewTerm",
	"Elections_EmptyTerm",
	"Elections_MemberKicked",
	"Elections_MemberRenounced",
	"Elections_VoterReported",
	"Identity_IdentitySet",
	"Identity_IdentityCleared",
	"Identity_IdentityKilled",
	"Identity_JudgementRequested",
	"Identity_JudgementUnrequested",
	"Identity_JudgementGiven",
	"Identity_RegistrarAdded",
	"Identity_SubIdentityAdded",
	"Identity_SubIdentityRemoved",
	"Identity_SubIdentityRevoked",
	"Society_Founded",
	"Society_Bid",
	"Society_Vouch",
	"Society_AutoUnbid",
	"Society_Unbid",
	"Society_Unvouch",
	"Society_Inducted",
	"Society_SuspendedMemberJudgement",
	"Society_CandidateSuspended",
	"Society_MemberSuspended",
	"Society_Challenged",
	"Society_Vote",
	"Society_DefenderVote",
	"Society_NewMaxMembers",
	"Society_Unfounded",
	"Society_Deposit",
	"Recovery_RecoveryCreated",
	"Recovery_RecoveryInitiated",
	"Recovery_RecoveryVouched",
	"Recovery_RecoveryClosed",
	"Recovery_AccountRecovered",
	"Recovery_RecoveryRemoved",
	"Vesting_VestingUpdated",
	"Vesting_VestingCompleted",
	"Scheduler_Scheduled",
	"Scheduler_Canceled",
	"Scheduler_Dispatched",
	"Proxy_ProxyExecuted",
	"Proxy_AnonymousCreated",
	"Sudo_Sudid",
	"Sudo_KeyChanged",
	"Sudo_SudoAsDone",
	"Treasury_Proposed",
	"Treasury_Spending",
	"Treasury_Awarded",
	"Treasury_Rejected",
	"Treasury_Burnt",
	"Treasury_Rollover",
	"Treasury_Deposit",
	"Treasury_NewTip",
	"Treasury_TipClosing",
	"Treasury_TipClosed",
	"Treasury_TipRetracted",
	"Contracts_Instantiated",
	"Contracts_Evicted",
	"Contracts_Restored",
	"Contracts_CodeStored",
	"Contracts_ScheduleUpdated",
	"Contracts_ContractExecution",
	"Utility_BatchInterrupted",
	"Utility_BatchCompleted",
	"Multisig_New",
	"Multisig_Approval",
	"Multisig_Executed",
	"Multisig_Cancelled",
	"Treasury_BountyProposed",
	"Treasury_BountyRejected",
	"Treasury_BountyBecameActive",
	"Treasury_BountyAwarded",
	"Treasury_BountyClaimed",
	"Treasury_BountyCanceled",
	"Treasury_BountyExtended",
	"TechnicalMembership_Dummy",
	"Currencies_Transferred",
	"Currencies_BalanceUpdated",
	"Currencies_Deposited",
	"Currencies_Withdrawn",
	"Vesting_VestingScheduleAdded",
	"Vesting_Claimed",
	"Vesting_VestingSchedulesUpdated",
	"Multisig_NewMultisig",
	"Multisig_MultisigApproval",
	"Multisig_MultisigExecuted",
	"Multisig_MultisigCancelled",
	"Balances_ReserveRepatriated",
	"Proxy_Announced",
}

// 269-108 = 161 = 141+4+15
func IsExist(typeName string) bool {
	for _, v := range existTypes {
		if typeName == v {
			return true
		}
	}
	return false
}

type TypesJson struct {
	CoinName    string      `json:"coin_name"`
	ChainName   string      `json:"chain_name"`
	Version     string      `json:"version"`
	SpecVersion int         `json:"spec_version"`
	Types       []EventType `json:"types"`
}

type EventType struct {
	Name     string   `json:"name"`
	SubTypes []string `json:"sub_types"`
}
