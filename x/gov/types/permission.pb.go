// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PermValue int32

const (
	// PERMISSION_ZERO is a no-op permission.
	PermZero PermValue = 0
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set Permissions to other actors.
	PermSetPermissions PermValue = 1
	// PERMISSION_CLAIM_VALIDATOR defines the permission that allows to Claim a validator Seat.
	PermClaimValidator PermValue = 2
	// PERMISSION_CLAIM_COUNCILOR defines the permission that allows to Claim a Councilor Seat.
	PermClaimCouncilor PermValue = 3
	// PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL defines the permission needed to create proposals for setting permissions.
	PermCreateSetPermissionsProposal PermValue = 4
	// PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set permissions.
	PermVoteSetPermissionProposal PermValue = 5
	//  PERMISSION_UPSERT_TOKEN_ALIAS
	PermUpsertTokenAlias PermValue = 6
	// PERMISSION_CHANGE_TX_FEE
	PermChangeTxFee PermValue = 7
	// PERMISSION_UPSERT_TOKEN_RATE
	PermUpsertTokenRate PermValue = 8
	// PERMISSION_UPSERT_ROLE makes possible to add, modify and assign roles.
	PermUpsertRole PermValue = 9
	// PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermCreateUpsertDataRegistryProposal PermValue = 10
	// PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermVoteUpsertDataRegistryProposal PermValue = 11
	// PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission needed to create proposals for setting network property.
	PermCreateSetNetworkPropertyProposal PermValue = 12
	// PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set network property.
	PermVoteSetNetworkPropertyProposal PermValue = 13
	// PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to create proposals for upsert token Alias.
	PermCreateUpsertTokenAliasProposal PermValue = 14
	// PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to vote proposals for upsert token.
	PermVoteUpsertTokenAliasProposal PermValue = 15
	// PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES defines the permission needed to create proposals for setting poor network messages
	PermCreateSetPoorNetworkMessagesProposal PermValue = 16
	// PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL defines the permission needed to vote proposals to set poor network messages
	PermVoteSetPoorNetworkMessagesProposal PermValue = 17
	// PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to create proposals for upsert token rate.
	PermCreateUpsertTokenRateProposal PermValue = 18
	// PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to vote proposals for upsert token rate.
	PermVoteUpsertTokenRateProposal PermValue = 19
	// PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to create a proposal to unjail a validator.
	PermCreateUnjailValidatorProposal PermValue = 20
	// PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to vote a proposal to unjail a validator.
	PermVoteUnjailValidatorProposal PermValue = 21
	// PERMISSION_CREATE_CREATE_ROLE_PROPOSAL defines the permission needed to create a proposal to create a role.
	PermCreateRoleProposal PermValue = 22
	// PERMISSION_VOTE_CREATE_ROLE_PROPOSAL defines the permission needed to vote a proposal to create a role.
	PermVoteCreateRoleProposal PermValue = 23
	// PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to create a proposal to blacklist/whitelisted tokens
	PermCreateTokensWhiteBlackChangeProposal PermValue = 24
	// PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to vote on blacklist/whitelisted tokens proposal
	PermVoteTokensWhiteBlackChangeProposal PermValue = 25
	// PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to create a proposal to reset whole validator rank
	PermCreateResetWholeValidatorRankProposal PermValue = 26
	// PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to vote on reset whole validator rank proposal
	PermVoteResetWholeValidatorRankProposal PermValue = 27
)

var PermValue_name = map[int32]string{
	0:  "PERMISSION_ZERO",
	1:  "PERMISSION_SET_PERMISSIONS",
	2:  "PERMISSION_CLAIM_VALIDATOR",
	3:  "PERMISSION_CLAIM_COUNCILOR",
	4:  "PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL",
	5:  "PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL",
	6:  "PERMISSION_UPSERT_TOKEN_ALIAS",
	7:  "PERMISSION_CHANGE_TX_FEE",
	8:  "PERMISSION_UPSERT_TOKEN_RATE",
	9:  "PERMISSION_UPSERT_ROLE",
	10: "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL",
	11: "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL",
	12: "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL",
	13: "PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL",
	14: "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	15: "PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	16: "PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES",
	17: "PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL",
	18: "PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL",
	19: "PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL",
	20: "PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL",
	21: "PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL",
	22: "PERMISSION_CREATE_CREATE_ROLE_PROPOSAL",
	23: "PERMISSION_VOTE_CREATE_ROLE_PROPOSAL",
	24: "PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
	25: "PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
	26: "PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
	27: "PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
}

var PermValue_value = map[string]int32{
	"PERMISSION_ZERO":                                       0,
	"PERMISSION_SET_PERMISSIONS":                            1,
	"PERMISSION_CLAIM_VALIDATOR":                            2,
	"PERMISSION_CLAIM_COUNCILOR":                            3,
	"PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL":            4,
	"PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL":              5,
	"PERMISSION_UPSERT_TOKEN_ALIAS":                         6,
	"PERMISSION_CHANGE_TX_FEE":                              7,
	"PERMISSION_UPSERT_TOKEN_RATE":                          8,
	"PERMISSION_UPSERT_ROLE":                                9,
	"PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL":       10,
	"PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL":         11,
	"PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL":       12,
	"PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL":         13,
	"PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL":         14,
	"PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL":           15,
	"PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES":           16,
	"PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL":    17,
	"PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL":          18,
	"PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL":            19,
	"PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL":           20,
	"PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL":             21,
	"PERMISSION_CREATE_CREATE_ROLE_PROPOSAL":                22,
	"PERMISSION_VOTE_CREATE_ROLE_PROPOSAL":                  23,
	"PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL":  24,
	"PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL":    25,
	"PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL": 26,
	"PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL":   27,
}

func (x PermValue) String() string {
	return proto.EnumName(PermValue_name, int32(x))
}

func (PermValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}

func init() {
	proto.RegisterEnum("tsuki.gov.PermValue", PermValue_name, PermValue_value)
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }

var fileDescriptor_c837ef01cbda0ad8 = []byte{
	// 889 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xc7, 0x1d, 0x28, 0x25, 0x5d, 0x4a, 0xa3, 0x2a, 0x21, 0x0d, 0x0b, 0x35, 0x1b, 0x48, 0x43,
	0x9a, 0xb6, 0xf1, 0x40, 0x81, 0x19, 0x86, 0x0b, 0x46, 0x71, 0x36, 0x89, 0xb0, 0x63, 0x79, 0x56,
	0x72, 0x4c, 0x33, 0xd3, 0xd1, 0x6c, 0xd2, 0xc5, 0x16, 0x56, 0xbc, 0x1e, 0x49, 0x29, 0xed, 0x1b,
	0x30, 0x7b, 0xc5, 0x0b, 0xec, 0x15, 0x2f, 0xc3, 0x1d, 0xb9, 0xe4, 0x92, 0x49, 0x5e, 0xa4, 0x23,
	0x27, 0xd1, 0xea, 0xcb, 0x4e, 0xae, 0xfc, 0xa5, 0xf3, 0x3b, 0xff, 0xf3, 0x3f, 0x1f, 0x63, 0xa0,
	0x8d, 0x58, 0x70, 0xec, 0x85, 0xa1, 0xc7, 0x87, 0x1b, 0xa3, 0x80, 0x47, 0x5c, 0x9f, 0x1d, 0x78,
	0x01, 0xdd, 0xe8, 0xf1, 0xd7, 0x70, 0xa1, 0xc7, 0x7b, 0x7c, 0xfc, 0x65, 0x2d, 0x7e, 0x77, 0xf1,
	0xfb, 0xfa, 0xbf, 0xf7, 0xc1, 0x9d, 0x36, 0x0b, 0x8e, 0xf7, 0xa9, 0x7f, 0xc2, 0xf4, 0x65, 0x30,
	0xd7, 0xc6, 0x64, 0xcf, 0xb4, 0x6d, 0xd3, 0x6a, 0xb9, 0x07, 0x98, 0x58, 0x5a, 0x05, 0xde, 0x15,
	0x12, 0xcd, 0xc6, 0xcf, 0x1c, 0xb0, 0x80, 0xeb, 0x3f, 0x00, 0x98, 0x7a, 0xc4, 0xc6, 0x8e, 0xab,
	0x3e, 0xda, 0xda, 0x0c, 0x5c, 0x14, 0x12, 0xe9, 0xf1, 0xd3, 0x36, 0x8b, 0xda, 0x89, 0x9a, 0x30,
	0x17, 0x57, 0x6f, 0x1a, 0xe6, 0x9e, 0xbb, 0x6f, 0x34, 0xcd, 0x2d, 0xc3, 0xb1, 0x88, 0xf6, 0x9e,
	0x8a, 0xab, 0xfb, 0xd4, 0x8b, 0xe5, 0x78, 0xaf, 0x68, 0xc4, 0x83, 0xd2, 0xb8, 0xba, 0xd5, 0x69,
	0xd5, 0xcd, 0xa6, 0x45, 0xb4, 0xf7, 0x73, 0x71, 0x75, 0x7e, 0x32, 0x3c, 0xf2, 0x7c, 0x1e, 0xe8,
	0x0e, 0x58, 0x4f, 0xc7, 0x11, 0x6c, 0x38, 0x38, 0x2f, 0xd7, 0x6d, 0x13, 0xab, 0x6d, 0xd9, 0x46,
	0x53, 0xbb, 0x05, 0x57, 0x84, 0x44, 0x68, 0xcc, 0x09, 0x18, 0x8d, 0x58, 0x56, 0x7d, 0x3b, 0xe0,
	0x23, 0x1e, 0x52, 0x5f, 0xb7, 0xc0, 0x5a, 0x8a, 0xba, 0x6f, 0x4d, 0x63, 0x7e, 0x00, 0x97, 0x85,
	0x44, 0x0f, 0xc7, 0xee, 0xf2, 0x1c, 0x31, 0x01, 0xfe, 0x04, 0x1e, 0xa6, 0x80, 0x9d, 0xb6, 0x8d,
	0x89, 0xe3, 0x3a, 0x56, 0x03, 0xb7, 0x5c, 0xa3, 0x69, 0x1a, 0xb6, 0x76, 0x1b, 0x2e, 0x09, 0x89,
	0x16, 0xe2, 0xd0, 0xce, 0x28, 0x64, 0x41, 0xe4, 0xf0, 0x01, 0x1b, 0x1a, 0xbe, 0x47, 0x43, 0xfd,
	0x1b, 0xb0, 0x94, 0xae, 0x71, 0xd7, 0x68, 0xed, 0x60, 0xd7, 0xf9, 0xd5, 0xdd, 0xc6, 0x58, 0xfb,
	0x10, 0xce, 0x0b, 0x89, 0xe6, 0xc6, 0x15, 0xf5, 0xe9, 0xb0, 0xc7, 0x9c, 0x37, 0xdb, 0x8c, 0xe9,
	0x3f, 0x82, 0xcf, 0x27, 0xe5, 0x23, 0x86, 0x83, 0xb5, 0x59, 0xf8, 0x40, 0x48, 0x34, 0x9f, 0x4b,
	0x47, 0x68, 0xc4, 0xf4, 0x0d, 0xb0, 0x58, 0x0c, 0x25, 0x56, 0x13, 0x6b, 0x77, 0xa0, 0x2e, 0x24,
	0xba, 0xa7, 0x82, 0x08, 0xf7, 0x99, 0xfe, 0x12, 0xd4, 0x8a, 0x1d, 0xb8, 0x0c, 0xdb, 0x32, 0x1c,
	0xc3, 0x25, 0x78, 0xc7, 0xb4, 0x1d, 0xf2, 0x42, 0x59, 0x06, 0xe0, 0x9a, 0x90, 0x68, 0x45, 0xb5,
	0xe1, 0x02, 0xb7, 0x45, 0x23, 0x4a, 0x58, 0xcf, 0x0b, 0xa3, 0xe0, 0x6d, 0xe2, 0xdc, 0x0b, 0xf0,
	0x2c, 0xdf, 0x8a, 0xe9, 0xf0, 0x8f, 0xe0, 0xaa, 0x90, 0xe8, 0xcb, 0xab, 0x7e, 0x4c, 0x41, 0x97,
	0x2a, 0x8f, 0xfb, 0xdc, 0xc2, 0x4e, 0xd7, 0x22, 0x8d, 0x31, 0x13, 0x13, 0x27, 0x05, 0xbf, 0x9b,
	0x57, 0x6e, 0xb3, 0xa8, 0xc5, 0xa2, 0x3f, 0x78, 0x30, 0x88, 0xb1, 0x2c, 0x88, 0xa6, 0x2a, 0x9f,
	0x0e, 0xff, 0x38, 0xab, 0xfc, 0xc6, 0xe8, 0xac, 0xe7, 0xa9, 0xa9, 0x52, 0xe8, 0x7b, 0x0a, 0x9d,
	0x76, 0x5c, 0x0d, 0x59, 0x82, 0xee, 0x80, 0x27, 0x13, 0xfc, 0x2e, 0x05, 0xcf, 0xa9, 0x8d, 0x52,
	0x6e, 0x97, 0x60, 0x5f, 0x66, 0xb0, 0xe9, 0x3d, 0xb5, 0x2c, 0x92, 0x78, 0xb2, 0x87, 0x6d, 0xdb,
	0xd8, 0xc1, 0xb6, 0xa6, 0xc1, 0xa7, 0x42, 0xa2, 0xb5, 0xec, 0xa2, 0x72, 0x1e, 0x5c, 0x1a, 0xb2,
	0xc7, 0xc2, 0x90, 0xf6, 0x98, 0xc2, 0x1f, 0x82, 0x6f, 0x4b, 0x17, 0xb6, 0x0c, 0xae, 0xc4, 0xdf,
	0x87, 0xeb, 0x42, 0xa2, 0xd5, 0xf4, 0xea, 0x4e, 0xc9, 0xd1, 0x05, 0x4f, 0xaf, 0x31, 0x3d, 0x5e,
	0x2d, 0x45, 0xd7, 0xe1, 0x23, 0x21, 0xd1, 0x72, 0xa9, 0xe7, 0xf1, 0xa6, 0x25, 0x60, 0x3b, 0x73,
	0xc3, 0x8a, 0x96, 0x67, 0xb1, 0xf3, 0xf0, 0x2b, 0x21, 0xd1, 0x17, 0x25, 0x8e, 0x67, 0xa0, 0xfb,
	0x65, 0x86, 0x77, 0x5a, 0xbf, 0x18, 0x66, 0x53, 0x1d, 0x64, 0x45, 0x5d, 0x28, 0x88, 0x1d, 0xfe,
	0x4e, 0x3d, 0x3f, 0x39, 0xd0, 0x09, 0x97, 0x80, 0xc7, 0x05, 0xb1, 0x13, 0xa9, 0x9f, 0xe4, 0xb4,
	0x4e, 0x60, 0x6e, 0x83, 0xd5, 0xa2, 0xd6, 0xcb, 0x97, 0xf8, 0xf2, 0x28, 0xe0, 0x22, 0x84, 0x42,
	0xa2, 0x45, 0x25, 0x33, 0x3e, 0x41, 0x09, 0x67, 0x17, 0xac, 0xe4, 0xb5, 0x95, 0x52, 0x1e, 0xc0,
	0xaa, 0x90, 0x08, 0x5e, 0xc9, 0x2a, 0x21, 0xfd, 0x06, 0xbe, 0x2b, 0x2a, 0x1a, 0x77, 0xc3, 0x76,
	0xbb, 0xbb, 0xa6, 0x83, 0xdd, 0xcd, 0xa6, 0x51, 0x6f, 0x5c, 0x1d, 0xe3, 0x84, 0xbc, 0x94, 0x9f,
	0xdb, 0x71, 0x63, 0xc2, 0x6e, 0xdf, 0x8b, 0xd8, 0xa6, 0x4f, 0x8f, 0x06, 0x17, 0x47, 0x7a, 0xda,
	0xdc, 0xde, 0x20, 0xcb, 0xa7, 0xd9, 0xb9, 0xbd, 0x26, 0x47, 0x1f, 0x7c, 0x5f, 0xac, 0x85, 0xe0,
	0x78, 0x3f, 0xba, 0xbb, 0xb1, 0x2f, 0xaa, 0x71, 0xc4, 0x68, 0x35, 0x54, 0x1a, 0x08, 0x9f, 0x09,
	0x89, 0x1e, 0xa7, 0xcc, 0x66, 0x21, 0x8b, 0xba, 0x7d, 0xee, 0xb3, 0xa4, 0x87, 0x84, 0x0e, 0x07,
	0x49, 0xa6, 0x57, 0xe0, 0x79, 0xbe, 0x9a, 0x9b, 0xe4, 0xf9, 0x0c, 0x3e, 0x11, 0x12, 0x7d, 0x7d,
	0x55, 0xce, 0x35, 0x59, 0xe0, 0xad, 0x3f, 0xff, 0xae, 0x56, 0x36, 0x7f, 0xfe, 0xe7, 0xac, 0x3a,
	0x73, 0x7a, 0x56, 0x9d, 0xf9, 0xff, 0xac, 0x3a, 0xf3, 0xd7, 0x79, 0xb5, 0x72, 0x7a, 0x5e, 0xad,
	0xfc, 0x77, 0x5e, 0xad, 0x1c, 0x3c, 0xea, 0x79, 0x51, 0xff, 0xe4, 0x70, 0xe3, 0x88, 0x1f, 0xd7,
	0x1a, 0x5e, 0x40, 0xeb, 0x3c, 0x60, 0xb5, 0x90, 0x0d, 0xa8, 0x57, 0x7b, 0x53, 0xeb, 0xf1, 0xd7,
	0xb5, 0xe8, 0xed, 0x88, 0x85, 0x87, 0xb7, 0xc7, 0xff, 0x8c, 0x9e, 0xbf, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0x09, 0xbd, 0xf1, 0x4d, 0x09, 0x00, 0x00,
}
