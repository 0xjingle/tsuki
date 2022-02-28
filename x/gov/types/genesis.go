package types

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
)

// DefaultGenesis returns the default CustomGo genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NextRoleId: 3,
		Roles: []Role{
			{
				Id:          uint32(RoleSudo),
				Sid:         "sudo",
				Description: "Sudo role",
			},
			{
				Id:          uint32(RoleValidator),
				Sid:         "validator",
				Description: "Validator role",
			},
		},
		RolePermissions: map[uint64]*Permissions{
			RoleSudo: NewPermissions([]PermValue{
				PermSetPermissions,
				PermClaimValidator,
				PermClaimCouncilor,
				PermUpsertTokenAlias,
				// PermChangeTxFee, // do not give this permission to sudo account - test does not pass
				PermUpsertTokenRate,
				PermUpsertRole,
				PermCreateSetNetworkPropertyProposal,
				PermVoteSetNetworkPropertyProposal,
				PermCreateUpsertDataRegistryProposal,
				PermVoteUpsertDataRegistryProposal,
				PermCreateUpsertTokenAliasProposal,
				PermVoteUpsertTokenAliasProposal,
				PermCreateUpsertTokenRateProposal,
				PermVoteUpsertTokenRateProposal,
				PermCreateUnjailValidatorProposal,
				PermVoteUnjailValidatorProposal,
				PermCreateRoleProposal,
				PermVoteCreateRoleProposal,
				PermCreateSetProposalDurationProposal,
				PermVoteSetProposalDurationProposal,
				PermCreateTokensWhiteBlackChangeProposal,
				PermVoteTokensWhiteBlackChangeProposal,
				PermCreateSetPoorNetworkMessagesProposal,
				PermVoteSetPoorNetworkMessagesProposal,
				PermWhitelistAccountPermissionProposal,
				PermVoteWhitelistAccountPermissionProposal,
				PermCreateResetWholeValidatorRankProposal,
				PermVoteResetWholeValidatorRankProposal,
				PermCreateSoftwareUpgradeProposal,
				PermVoteSoftwareUpgradeProposal,
				PermSetClaimValidatorPermission,
				PermBlacklistAccountPermissionProposal,
				PermVoteBlacklistAccountPermissionProposal,
				PermRemoveWhitelistedAccountPermissionProposal,
				PermVoteRemoveWhitelistedAccountPermissionProposal,
				PermRemoveBlacklistedAccountPermissionProposal,
				PermVoteRemoveBlacklistedAccountPermissionProposal,
				PermWhitelistRolePermissionProposal,
				PermVoteWhitelistRolePermissionProposal,
				PermBlacklistRolePermissionProposal,
				PermVoteBlacklistRolePermissionProposal,
				PermRemoveWhitelistedRolePermissionProposal,
				PermVoteRemoveWhitelistedRolePermissionProposal,
				PermRemoveBlacklistedRolePermissionProposal,
				PermVoteRemoveBlacklistedRolePermissionProposal,
				PermAssignRoleToAccountProposal,
				PermVoteAssignRoleToAccountProposal,
				PermUnassignRoleFromAccountProposal,
				PermVoteUnassignRoleFromAccountProposal,
				PermRemoveRoleProposal,
				PermVoteRemoveRoleProposal,
			}, nil),
			uint64(RoleValidator): NewPermissions([]PermValue{PermClaimValidator}, nil),
		},
		StartingProposalId: 1,
		NetworkProperties: &NetworkProperties{
			MinTxFee:                    100,
			MaxTxFee:                    1000000,
			VoteQuorum:                  33,
			MinimumProposalEndTime:      300, // 300 seconds / 5 mins
			ProposalEnactmentTime:       300, // 300 seconds / 5 mins
			MinProposalEndBlocks:        2,
			MinProposalEnactmentBlocks:  1,
			EnableForeignFeePayments:    true,
			MischanceRankDecreaseAmount: 10,
			MischanceConfidence:         10,
			MaxMischance:                110,
			InactiveRankDecreasePercent: 50,      // 50%
			PoorNetworkMaxBankSend:      1000000, // 1M ukex
			MinValidators:               1,
			UnjailMaxTime:               600, // 600  seconds / 10 mins
			EnableTokenWhitelist:        false,
			EnableTokenBlacklist:        true,
			MinIdentityApprovalTip:      200,
			UniqueIdentityKeys:          "moniker,username",
			UbiHardcap:                  6000_000,
		},
		ExecutionFees: []*ExecutionFee{
			{
				TransactionType:   tsukitypes.MsgTypeClaimValidator,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypeClaimCouncilor,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "claim-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "vote-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "submit-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "veto-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypeUpsertTokenAlias,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypeActivate,
				ExecutionFee:      100,
				FailureFee:        1000,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypePause,
				ExecutionFee:      100,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypeUnpause,
				ExecutionFee:      100,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
		},
		PoorNetworkMessages: &AllowedMessages{
			Messages: []string{
				tsukitypes.MsgTypeSubmitProposal,
				tsukitypes.MsgTypeSetNetworkProperties,
				tsukitypes.MsgTypeVoteProposal,
				tsukitypes.MsgTypeClaimCouncilor,
				tsukitypes.MsgTypeWhitelistPermissions,
				tsukitypes.MsgTypeBlacklistPermissions,
				tsukitypes.MsgTypeCreateRole,
				tsukitypes.MsgTypeAssignRole,
				tsukitypes.MsgTypeRemoveRole,
				tsukitypes.MsgTypeWhitelistRolePermission,
				tsukitypes.MsgTypeBlacklistRolePermission,
				tsukitypes.MsgTypeRemoveWhitelistRolePermission,
				tsukitypes.MsgTypeRemoveBlacklistRolePermission,
				tsukitypes.MsgTypeClaimValidator,
				tsukitypes.MsgTypeActivate,
				tsukitypes.MsgTypePause,
				tsukitypes.MsgTypeUnpause,
				tsukitypes.MsgTypeRegisterIdentityRecords,
				tsukitypes.MsgTypeEditIdentityRecord,
				tsukitypes.MsgTypeRequestIdentityRecordsVerify,
				tsukitypes.MsgTypeHandleIdentityRecordsVerifyRequest,
				tsukitypes.MsgTypeCancelIdentityRecordsVerifyRequest,
			},
		},
		LastIdentityRecordId:        0,
		LastIdRecordVerifyRequestId: 0,
	}
}
