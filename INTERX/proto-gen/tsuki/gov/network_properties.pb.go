// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tsuki/gov/network_properties.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NetworkProperty int32

const (
	NetworkProperty_MIN_TX_FEE                     NetworkProperty = 0
	NetworkProperty_MAX_TX_FEE                     NetworkProperty = 1
	NetworkProperty_VOTE_QUORUM                    NetworkProperty = 2
	NetworkProperty_PROPOSAL_END_TIME              NetworkProperty = 3
	NetworkProperty_PROPOSAL_ENACTMENT_TIME        NetworkProperty = 4
	NetworkProperty_MIN_PROPOSAL_END_BLOCKS        NetworkProperty = 5
	NetworkProperty_MIN_PROPOSAL_ENACTMENT_BLOCKS  NetworkProperty = 6
	NetworkProperty_ENABLE_FOREIGN_FEE_PAYMENTS    NetworkProperty = 7
	NetworkProperty_MISCHANCE_RANK_DECREASE_AMOUNT NetworkProperty = 8
	NetworkProperty_MAX_MISCHANCE                  NetworkProperty = 9
	NetworkProperty_MISCHANCE_CONFIDENCE           NetworkProperty = 10
	NetworkProperty_INACTIVE_RANK_DECREASE_PERCENT NetworkProperty = 11
	NetworkProperty_POOR_NETWORK_MAX_BANK_SEND     NetworkProperty = 12
	NetworkProperty_MIN_VALIDATORS                 NetworkProperty = 13
	NetworkProperty_JAIL_MAX_TIME                  NetworkProperty = 14
	NetworkProperty_ENABLE_TOKEN_WHITELIST         NetworkProperty = 15
	NetworkProperty_ENABLE_TOKEN_BLACKLIST         NetworkProperty = 16
)

var NetworkProperty_name = map[int32]string{
	0:  "MIN_TX_FEE",
	1:  "MAX_TX_FEE",
	2:  "VOTE_QUORUM",
	3:  "PROPOSAL_END_TIME",
	4:  "PROPOSAL_ENACTMENT_TIME",
	5:  "MIN_PROPOSAL_END_BLOCKS",
	6:  "MIN_PROPOSAL_ENACTMENT_BLOCKS",
	7:  "ENABLE_FOREIGN_FEE_PAYMENTS",
	8:  "MISCHANCE_RANK_DECREASE_AMOUNT",
	9:  "MAX_MISCHANCE",
	10: "MISCHANCE_CONFIDENCE",
	11: "INACTIVE_RANK_DECREASE_PERCENT",
	12: "POOR_NETWORK_MAX_BANK_SEND",
	13: "MIN_VALIDATORS",
	14: "JAIL_MAX_TIME",
	15: "ENABLE_TOKEN_WHITELIST",
	16: "ENABLE_TOKEN_BLACKLIST",
}

var NetworkProperty_value = map[string]int32{
	"MIN_TX_FEE":                     0,
	"MAX_TX_FEE":                     1,
	"VOTE_QUORUM":                    2,
	"PROPOSAL_END_TIME":              3,
	"PROPOSAL_ENACTMENT_TIME":        4,
	"MIN_PROPOSAL_END_BLOCKS":        5,
	"MIN_PROPOSAL_ENACTMENT_BLOCKS":  6,
	"ENABLE_FOREIGN_FEE_PAYMENTS":    7,
	"MISCHANCE_RANK_DECREASE_AMOUNT": 8,
	"MAX_MISCHANCE":                  9,
	"MISCHANCE_CONFIDENCE":           10,
	"INACTIVE_RANK_DECREASE_PERCENT": 11,
	"POOR_NETWORK_MAX_BANK_SEND":     12,
	"MIN_VALIDATORS":                 13,
	"JAIL_MAX_TIME":                  14,
	"ENABLE_TOKEN_WHITELIST":         15,
	"ENABLE_TOKEN_BLACKLIST":         16,
}

func (x NetworkProperty) String() string {
	return proto.EnumName(NetworkProperty_name, int32(x))
}

func (NetworkProperty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{0}
}

type MsgSetNetworkProperties struct {
	NetworkProperties    *NetworkProperties `protobuf:"bytes,1,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	Proposer             []byte             `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MsgSetNetworkProperties) Reset()         { *m = MsgSetNetworkProperties{} }
func (m *MsgSetNetworkProperties) String() string { return proto.CompactTextString(m) }
func (*MsgSetNetworkProperties) ProtoMessage()    {}
func (*MsgSetNetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{0}
}

func (m *MsgSetNetworkProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSetNetworkProperties.Unmarshal(m, b)
}
func (m *MsgSetNetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSetNetworkProperties.Marshal(b, m, deterministic)
}
func (m *MsgSetNetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetNetworkProperties.Merge(m, src)
}
func (m *MsgSetNetworkProperties) XXX_Size() int {
	return xxx_messageInfo_MsgSetNetworkProperties.Size(m)
}
func (m *MsgSetNetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetNetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetNetworkProperties proto.InternalMessageInfo

func (m *MsgSetNetworkProperties) GetNetworkProperties() *NetworkProperties {
	if m != nil {
		return m.NetworkProperties
	}
	return nil
}

func (m *MsgSetNetworkProperties) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type NetworkProperties struct {
	MinTxFee                    uint64   `protobuf:"varint,1,opt,name=min_tx_fee,json=minTxFee,proto3" json:"min_tx_fee,omitempty"`
	MaxTxFee                    uint64   `protobuf:"varint,2,opt,name=max_tx_fee,json=maxTxFee,proto3" json:"max_tx_fee,omitempty"`
	VoteQuorum                  uint64   `protobuf:"varint,3,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	ProposalEndTime             uint64   `protobuf:"varint,4,opt,name=proposal_end_time,json=proposalEndTime,proto3" json:"proposal_end_time,omitempty"`
	ProposalEnactmentTime       uint64   `protobuf:"varint,5,opt,name=proposal_enactment_time,json=proposalEnactmentTime,proto3" json:"proposal_enactment_time,omitempty"`
	MinProposalEndBlocks        uint64   `protobuf:"varint,6,opt,name=min_proposal_end_blocks,json=minProposalEndBlocks,proto3" json:"min_proposal_end_blocks,omitempty"`
	MinProposalEnactmentBlocks  uint64   `protobuf:"varint,7,opt,name=min_proposal_enactment_blocks,json=minProposalEnactmentBlocks,proto3" json:"min_proposal_enactment_blocks,omitempty"`
	EnableForeignFeePayments    bool     `protobuf:"varint,8,opt,name=enable_foreign_fee_payments,json=enableForeignFeePayments,proto3" json:"enable_foreign_fee_payments,omitempty"`
	MischanceRankDecreaseAmount uint64   `protobuf:"varint,9,opt,name=mischance_rank_decrease_amount,json=mischanceRankDecreaseAmount,proto3" json:"mischance_rank_decrease_amount,omitempty"`
	MaxMischance                uint64   `protobuf:"varint,10,opt,name=max_mischance,json=maxMischance,proto3" json:"max_mischance,omitempty"`
	MischanceConfidence         uint64   `protobuf:"varint,11,opt,name=mischance_confidence,json=mischanceConfidence,proto3" json:"mischance_confidence,omitempty"`
	InactiveRankDecreasePercent uint64   `protobuf:"varint,12,opt,name=inactive_rank_decrease_percent,json=inactiveRankDecreasePercent,proto3" json:"inactive_rank_decrease_percent,omitempty"`
	MinValidators               uint64   `protobuf:"varint,13,opt,name=min_validators,json=minValidators,proto3" json:"min_validators,omitempty"`
	PoorNetworkMaxBankSend      uint64   `protobuf:"varint,14,opt,name=poor_network_max_bank_send,json=poorNetworkMaxBankSend,proto3" json:"poor_network_max_bank_send,omitempty"`
	JailMaxTime                 uint64   `protobuf:"varint,15,opt,name=jail_max_time,json=jailMaxTime,proto3" json:"jail_max_time,omitempty"`
	EnableTokenWhitelist        bool     `protobuf:"varint,16,opt,name=enable_token_whitelist,json=enableTokenWhitelist,proto3" json:"enable_token_whitelist,omitempty"`
	EnableTokenBlacklist        bool     `protobuf:"varint,17,opt,name=enable_token_blacklist,json=enableTokenBlacklist,proto3" json:"enable_token_blacklist,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *NetworkProperties) Reset()         { *m = NetworkProperties{} }
func (m *NetworkProperties) String() string { return proto.CompactTextString(m) }
func (*NetworkProperties) ProtoMessage()    {}
func (*NetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{1}
}

func (m *NetworkProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkProperties.Unmarshal(m, b)
}
func (m *NetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkProperties.Marshal(b, m, deterministic)
}
func (m *NetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProperties.Merge(m, src)
}
func (m *NetworkProperties) XXX_Size() int {
	return xxx_messageInfo_NetworkProperties.Size(m)
}
func (m *NetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProperties proto.InternalMessageInfo

func (m *NetworkProperties) GetMinTxFee() uint64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

func (m *NetworkProperties) GetMaxTxFee() uint64 {
	if m != nil {
		return m.MaxTxFee
	}
	return 0
}

func (m *NetworkProperties) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *NetworkProperties) GetProposalEndTime() uint64 {
	if m != nil {
		return m.ProposalEndTime
	}
	return 0
}

func (m *NetworkProperties) GetProposalEnactmentTime() uint64 {
	if m != nil {
		return m.ProposalEnactmentTime
	}
	return 0
}

func (m *NetworkProperties) GetMinProposalEndBlocks() uint64 {
	if m != nil {
		return m.MinProposalEndBlocks
	}
	return 0
}

func (m *NetworkProperties) GetMinProposalEnactmentBlocks() uint64 {
	if m != nil {
		return m.MinProposalEnactmentBlocks
	}
	return 0
}

func (m *NetworkProperties) GetEnableForeignFeePayments() bool {
	if m != nil {
		return m.EnableForeignFeePayments
	}
	return false
}

func (m *NetworkProperties) GetMischanceRankDecreaseAmount() uint64 {
	if m != nil {
		return m.MischanceRankDecreaseAmount
	}
	return 0
}

func (m *NetworkProperties) GetMaxMischance() uint64 {
	if m != nil {
		return m.MaxMischance
	}
	return 0
}

func (m *NetworkProperties) GetMischanceConfidence() uint64 {
	if m != nil {
		return m.MischanceConfidence
	}
	return 0
}

func (m *NetworkProperties) GetInactiveRankDecreasePercent() uint64 {
	if m != nil {
		return m.InactiveRankDecreasePercent
	}
	return 0
}

func (m *NetworkProperties) GetMinValidators() uint64 {
	if m != nil {
		return m.MinValidators
	}
	return 0
}

func (m *NetworkProperties) GetPoorNetworkMaxBankSend() uint64 {
	if m != nil {
		return m.PoorNetworkMaxBankSend
	}
	return 0
}

func (m *NetworkProperties) GetJailMaxTime() uint64 {
	if m != nil {
		return m.JailMaxTime
	}
	return 0
}

func (m *NetworkProperties) GetEnableTokenWhitelist() bool {
	if m != nil {
		return m.EnableTokenWhitelist
	}
	return false
}

func (m *NetworkProperties) GetEnableTokenBlacklist() bool {
	if m != nil {
		return m.EnableTokenBlacklist
	}
	return false
}

func init() {
	proto.RegisterEnum("tsuki.gov.NetworkProperty", NetworkProperty_name, NetworkProperty_value)
	proto.RegisterType((*MsgSetNetworkProperties)(nil), "tsuki.gov.MsgSetNetworkProperties")
	proto.RegisterType((*NetworkProperties)(nil), "tsuki.gov.NetworkProperties")
}

func init() {
	proto.RegisterFile("tsuki/gov/network_properties.proto", fileDescriptor_98011a6048e5dde3)
}

var fileDescriptor_98011a6048e5dde3 = []byte{
	// 1026 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xcb, 0x6e, 0xdb, 0x46,
	0x17, 0x8e, 0x12, 0xc7, 0x51, 0x46, 0x92, 0x45, 0xd3, 0x8e, 0xcd, 0x9f, 0xce, 0x6f, 0xb3, 0x2e,
	0x02, 0x18, 0x01, 0x2c, 0x21, 0x4d, 0xd3, 0x45, 0x80, 0x2e, 0x28, 0x6a, 0xd4, 0xd0, 0x12, 0x2f,
	0xa1, 0xe8, 0x4b, 0xbb, 0x19, 0x50, 0xd4, 0x58, 0x66, 0x25, 0x72, 0x54, 0x92, 0x76, 0xe4, 0x37,
	0x28, 0xf8, 0x0e, 0x5c, 0xf5, 0x15, 0xfa, 0x3c, 0x5d, 0x74, 0xd1, 0x87, 0xe8, 0xaa, 0x98, 0x21,
	0x29, 0x4b, 0x56, 0xec, 0x95, 0x84, 0xf3, 0x5d, 0xce, 0x99, 0x39, 0xf3, 0x81, 0xe0, 0x9b, 0xb1,
	0x17, 0x3a, 0xcd, 0x11, 0xb9, 0x69, 0x06, 0x38, 0xfe, 0x42, 0xc2, 0x31, 0x9a, 0x86, 0x64, 0x8a,
	0xc3, 0xd8, 0xc3, 0x51, 0x63, 0x1a, 0x92, 0x98, 0xf0, 0x65, 0x4a, 0x69, 0x8c, 0xc8, 0x8d, 0xb8,
	0x3d, 0x22, 0x23, 0xc2, 0x8a, 0x4d, 0xfa, 0x2f, 0xc3, 0x0f, 0xff, 0x2c, 0x81, 0x5d, 0x2d, 0x1a,
	0xf5, 0x71, 0xac, 0x67, 0x16, 0xe6, 0xdc, 0x81, 0x3f, 0x01, 0xfc, 0xaa, 0xaf, 0x50, 0x92, 0x4a,
	0x47, 0x95, 0xef, 0xf6, 0x1a, 0x85, 0x71, 0x63, 0x45, 0x68, 0x6d, 0x06, 0x2b, 0x5e, 0x1a, 0x28,
	0x53, 0x0f, 0x12, 0xe1, 0x50, 0x78, 0x2a, 0x95, 0x8e, 0xaa, 0xad, 0x77, 0xff, 0xfe, 0x75, 0x70,
	0x3c, 0xf2, 0xe2, 0xab, 0xeb, 0x41, 0xc3, 0x25, 0x7e, 0xd3, 0x25, 0x91, 0x4f, 0xa2, 0xfc, 0xe7,
	0x38, 0x1a, 0x8e, 0x9b, 0xf1, 0xed, 0x14, 0x47, 0x0d, 0xd9, 0x75, 0xe5, 0xe1, 0x30, 0xc4, 0x51,
	0x64, 0xcd, 0x2d, 0x0e, 0xff, 0x5e, 0x07, 0x9b, 0xab, 0x03, 0xbf, 0x06, 0xc0, 0xf7, 0x02, 0x14,
	0xcf, 0xd0, 0x25, 0xc6, 0x6c, 0xd0, 0x35, 0xab, 0xec, 0x7b, 0x81, 0x3d, 0xeb, 0x60, 0xcc, 0x50,
	0x67, 0x56, 0xa0, 0x4f, 0x73, 0xd4, 0x99, 0x65, 0xe8, 0x01, 0xa8, 0xdc, 0x90, 0x18, 0xa3, 0xdf,
	0xae, 0x49, 0x78, 0xed, 0x0b, 0xcf, 0x18, 0x0c, 0x68, 0xe9, 0x33, 0xab, 0xf0, 0x6f, 0xc1, 0x66,
	0xd6, 0xde, 0x99, 0x20, 0x1c, 0x0c, 0x51, 0xec, 0xf9, 0x58, 0x58, 0x63, 0xb4, 0x7a, 0x01, 0xc0,
	0x60, 0x68, 0x7b, 0x3e, 0xe6, 0x7f, 0x00, 0xbb, 0x0b, 0x5c, 0xc7, 0x8d, 0x7d, 0x1c, 0xc4, 0x99,
	0xe2, 0x39, 0x53, 0xbc, 0xba, 0x53, 0xe4, 0x28, 0xd3, 0x7d, 0x00, 0xbb, 0xf4, 0x00, 0x4b, 0x7d,
	0x06, 0x13, 0xe2, 0x8e, 0x23, 0x61, 0x9d, 0xe9, 0xb6, 0x7d, 0x2f, 0x30, 0xef, 0x9a, 0xb5, 0x18,
	0xc6, 0xcb, 0xe0, 0xff, 0xf7, 0x64, 0x45, 0xcb, 0x5c, 0xfc, 0x82, 0x89, 0xc5, 0x25, 0x71, 0x4e,
	0xc9, 0x2d, 0x7e, 0x04, 0x7b, 0x38, 0x70, 0x06, 0x13, 0x8c, 0x2e, 0x49, 0x88, 0xbd, 0x51, 0x40,
	0x2f, 0x09, 0x4d, 0x9d, 0x5b, 0xca, 0x89, 0x84, 0xb2, 0x54, 0x3a, 0x2a, 0x5b, 0x42, 0x46, 0xe9,
	0x64, 0x8c, 0x0e, 0xc6, 0x66, 0x8e, 0xf3, 0x0a, 0xd8, 0xf7, 0xbd, 0xc8, 0xbd, 0x72, 0x02, 0x17,
	0xa3, 0xd0, 0x09, 0xc6, 0x68, 0x88, 0xdd, 0x10, 0x3b, 0x11, 0x46, 0x8e, 0x4f, 0xae, 0x83, 0x58,
	0x78, 0xc9, 0x46, 0xd8, 0x9b, 0xb3, 0x2c, 0x27, 0x18, 0xb7, 0x73, 0x8e, 0xcc, 0x28, 0xfc, 0xb7,
	0xa0, 0x46, 0x17, 0x34, 0xa7, 0x08, 0x80, 0x69, 0xaa, 0xbe, 0x33, 0xd3, 0x8a, 0x1a, 0xff, 0x0e,
	0x6c, 0xdf, 0x75, 0x72, 0x49, 0x70, 0xe9, 0x0d, 0x31, 0xe5, 0x56, 0x18, 0x77, 0x6b, 0x8e, 0x29,
	0x73, 0x88, 0x0e, 0xe7, 0xd1, 0xe3, 0x7a, 0x37, 0xf7, 0x67, 0x9b, 0xe2, 0xd0, 0xc5, 0x41, 0x2c,
	0x54, 0xb3, 0xe1, 0x0a, 0xd6, 0xe2, 0x6c, 0x66, 0x46, 0xe1, 0xdf, 0x80, 0x0d, 0x7a, 0xc7, 0x37,
	0xce, 0xc4, 0x1b, 0x3a, 0x31, 0x09, 0x23, 0xa1, 0xc6, 0x44, 0x35, 0xdf, 0x0b, 0xce, 0xe6, 0x45,
	0xfe, 0x23, 0x10, 0xa7, 0x84, 0x84, 0xa8, 0x08, 0x0e, 0x3d, 0xd0, 0x80, 0xf6, 0x8c, 0x70, 0x30,
	0x14, 0x36, 0x98, 0x64, 0x87, 0x32, 0xf2, 0xd7, 0xab, 0x39, 0xb3, 0x96, 0x13, 0x8c, 0xfb, 0x38,
	0x18, 0xf2, 0x87, 0xa0, 0xf6, 0xab, 0xe3, 0x4d, 0x98, 0x86, 0xbd, 0x95, 0x3a, 0xa3, 0x57, 0x68,
	0x51, 0x73, 0x66, 0xec, 0x85, 0x7c, 0x0f, 0x76, 0xf2, 0x3d, 0xc5, 0x64, 0x8c, 0x03, 0xf4, 0xe5,
	0xca, 0x8b, 0xf1, 0xc4, 0x8b, 0x62, 0x81, 0x63, 0x2b, 0xda, 0xce, 0x50, 0x9b, 0x82, 0xe7, 0x05,
	0xb6, 0xa2, 0x1a, 0x4c, 0x1c, 0x77, 0xcc, 0x54, 0x9b, 0x2b, 0xaa, 0x56, 0x81, 0xbd, 0xfd, 0x67,
	0x1d, 0xd4, 0x97, 0x43, 0x76, 0x4b, 0x43, 0xa4, 0xa9, 0x3a, 0xb2, 0x2f, 0x50, 0x07, 0x42, 0xee,
	0x89, 0x58, 0x4d, 0x52, 0xa9, 0xac, 0x2d, 0x44, 0x4c, 0x93, 0x2f, 0x0a, 0xb4, 0x94, 0xa3, 0x0b,
	0x11, 0x3b, 0x33, 0x6c, 0x88, 0x3e, 0x9f, 0x1a, 0xd6, 0xa9, 0xc6, 0x3d, 0x15, 0x37, 0x92, 0x54,
	0x02, 0x67, 0x4b, 0x11, 0x33, 0x2d, 0xc3, 0x34, 0xfa, 0x72, 0x0f, 0x41, 0xbd, 0x8d, 0x6c, 0x55,
	0x83, 0xdc, 0x33, 0x71, 0x2b, 0x49, 0xa5, 0xba, 0xb9, 0x1a, 0xb1, 0x05, 0xae, 0xac, 0xd8, 0x1a,
	0xd4, 0xed, 0x4c, 0xb1, 0x26, 0xfe, 0x2f, 0x49, 0xa5, 0x57, 0xe6, 0x43, 0x11, 0xa3, 0x07, 0x58,
	0xea, 0xd3, 0xea, 0x19, 0x4a, 0xb7, 0xcf, 0x3d, 0x17, 0x85, 0x24, 0x95, 0xb6, 0xb5, 0x07, 0x22,
	0x76, 0x4f, 0x56, 0xb4, 0xcc, 0xc5, 0xeb, 0xe2, 0x7e, 0x92, 0x4a, 0xa2, 0xf6, 0x68, 0xc4, 0xa0,
	0x2e, 0xb7, 0x7a, 0x10, 0x75, 0x0c, 0x0b, 0xaa, 0x3f, 0xe9, 0xf4, 0x92, 0x90, 0x29, 0xff, 0x4c,
	0x6d, 0xfa, 0xdc, 0x0b, 0xf1, 0x75, 0x92, 0x4a, 0x02, 0x7c, 0x24, 0x62, 0x9a, 0xda, 0x57, 0x3e,
	0xc9, 0xba, 0x02, 0x91, 0x25, 0xeb, 0x5d, 0xd4, 0x86, 0x8a, 0x05, 0xe5, 0x3e, 0x44, 0xb2, 0x66,
	0x9c, 0xea, 0x36, 0x57, 0x16, 0x0f, 0x92, 0x54, 0xda, 0xd3, 0x1e, 0x8f, 0x18, 0x5d, 0xd0, 0xdc,
	0x88, 0x7b, 0x29, 0x72, 0x49, 0x2a, 0x55, 0xb5, 0x7b, 0x11, 0xbb, 0xeb, 0xa4, 0x18, 0x7a, 0x47,
	0x6d, 0x43, 0xca, 0x05, 0xe2, 0x6e, 0x92, 0x4a, 0x5b, 0xda, 0xd7, 0x23, 0xa6, 0xd2, 0x1b, 0x51,
	0xcf, 0xee, 0xcf, 0x66, 0x42, 0x4b, 0x81, 0xba, 0xcd, 0x55, 0xb2, 0xe1, 0xd4, 0x47, 0x22, 0xf6,
	0x11, 0x88, 0xa6, 0x61, 0x58, 0x48, 0x87, 0xf6, 0xb9, 0x61, 0x75, 0x11, 0x9d, 0xb4, 0x45, 0xcd,
	0xfa, 0x50, 0x6f, 0x73, 0x55, 0x51, 0x4c, 0x52, 0x69, 0xc7, 0xfc, 0x7a, 0x76, 0xde, 0x80, 0x0d,
	0xba, 0x9f, 0x33, 0xb9, 0xa7, 0xb6, 0x65, 0xdb, 0xb0, 0xfa, 0x5c, 0x4d, 0xdc, 0x4c, 0x52, 0xa9,
	0xa6, 0x2d, 0xc5, 0xf3, 0x10, 0xd4, 0x4e, 0x64, 0xb5, 0xc7, 0xac, 0xd9, 0x5b, 0xd9, 0x10, 0xeb,
	0x49, 0x2a, 0x55, 0x4e, 0x96, 0x23, 0x96, 0xef, 0xc9, 0x36, 0xba, 0x50, 0x47, 0xe7, 0x9f, 0x54,
	0x1b, 0xf6, 0xd4, 0xbe, 0xcd, 0xd5, 0xb3, 0x07, 0x02, 0x1f, 0x88, 0xd8, 0x92, 0xaa, 0xd5, 0x93,
	0x95, 0x2e, 0x53, 0x71, 0x2b, 0xaa, 0x79, 0xc4, 0xc4, 0xb5, 0xdf, 0xff, 0xd8, 0x7f, 0xd2, 0xfa,
	0xf0, 0xcb, 0xfb, 0x85, 0x0f, 0x61, 0xd7, 0x0b, 0x1d, 0x85, 0x84, 0xb8, 0x19, 0xe1, 0xb1, 0xe3,
	0x35, 0x55, 0xdd, 0x86, 0xd6, 0x45, 0x93, 0x7d, 0xae, 0x8f, 0x47, 0x38, 0x68, 0x16, 0x1f, 0xfd,
	0xc1, 0x3a, 0xab, 0xbd, 0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x06, 0xbb, 0xd9, 0x2d, 0x07, 0x08,
	0x00, 0x00,
}
