// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/collectives/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// proposal to transfer accumulated donations to a specific account
type ProposalCollectiveSendDonation struct {
	Name    string                                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string                                    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Amounts []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,rep,name=amounts,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amounts"`
}

func (m *ProposalCollectiveSendDonation) Reset()         { *m = ProposalCollectiveSendDonation{} }
func (m *ProposalCollectiveSendDonation) String() string { return proto.CompactTextString(m) }
func (*ProposalCollectiveSendDonation) ProtoMessage()    {}
func (*ProposalCollectiveSendDonation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6173a479da28660d, []int{0}
}
func (m *ProposalCollectiveSendDonation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalCollectiveSendDonation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalCollectiveSendDonation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalCollectiveSendDonation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalCollectiveSendDonation.Merge(m, src)
}
func (m *ProposalCollectiveSendDonation) XXX_Size() int {
	return m.Size()
}
func (m *ProposalCollectiveSendDonation) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalCollectiveSendDonation.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalCollectiveSendDonation proto.InternalMessageInfo

func (m *ProposalCollectiveSendDonation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProposalCollectiveSendDonation) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// proposal to update staking collective, e.g. change description, owners, contributors, spending-pool list, claim period, etc.
type ProposalCollectiveUpdate struct {
	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description      string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Status           CollectiveStatus       `protobuf:"varint,3,opt,name=status,proto3,enum=tsuki.collectives.CollectiveStatus" json:"status,omitempty"`
	DepositWhitelist DepositWhitelist       `protobuf:"bytes,4,opt,name=deposit_whitelist,json=depositWhitelist,proto3" json:"deposit_whitelist"`
	OwnersWhitelist  OwnersWhitelist        `protobuf:"bytes,5,opt,name=owners_whitelist,json=ownersWhitelist,proto3" json:"owners_whitelist"`
	SpendingPools    []WeightedSpendingPool `protobuf:"bytes,6,rep,name=spending_pools,json=spendingPools,proto3" json:"spending_pools"`
	ClaimStart       uint64                 `protobuf:"varint,7,opt,name=claim_start,json=claimStart,proto3" json:"claim_start,omitempty"`
	ClaimPeriod      uint64                 `protobuf:"varint,8,opt,name=claim_period,json=claimPeriod,proto3" json:"claim_period,omitempty"`
	ClaimEnd         uint64                 `protobuf:"varint,9,opt,name=claim_end,json=claimEnd,proto3" json:"claim_end,omitempty"`
	VoteQuorum       uint64                 `protobuf:"varint,10,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	VotePeriod       uint64                 `protobuf:"varint,11,opt,name=vote_period,json=votePeriod,proto3" json:"vote_period,omitempty"`
	VoteEnactment    uint64                 `protobuf:"varint,12,opt,name=vote_enactment,json=voteEnactment,proto3" json:"vote_enactment,omitempty"`
}

func (m *ProposalCollectiveUpdate) Reset()         { *m = ProposalCollectiveUpdate{} }
func (m *ProposalCollectiveUpdate) String() string { return proto.CompactTextString(m) }
func (*ProposalCollectiveUpdate) ProtoMessage()    {}
func (*ProposalCollectiveUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6173a479da28660d, []int{1}
}
func (m *ProposalCollectiveUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalCollectiveUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalCollectiveUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalCollectiveUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalCollectiveUpdate.Merge(m, src)
}
func (m *ProposalCollectiveUpdate) XXX_Size() int {
	return m.Size()
}
func (m *ProposalCollectiveUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalCollectiveUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalCollectiveUpdate proto.InternalMessageInfo

func (m *ProposalCollectiveUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProposalCollectiveUpdate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProposalCollectiveUpdate) GetStatus() CollectiveStatus {
	if m != nil {
		return m.Status
	}
	return CollectiveActive
}

func (m *ProposalCollectiveUpdate) GetDepositWhitelist() DepositWhitelist {
	if m != nil {
		return m.DepositWhitelist
	}
	return DepositWhitelist{}
}

func (m *ProposalCollectiveUpdate) GetOwnersWhitelist() OwnersWhitelist {
	if m != nil {
		return m.OwnersWhitelist
	}
	return OwnersWhitelist{}
}

func (m *ProposalCollectiveUpdate) GetSpendingPools() []WeightedSpendingPool {
	if m != nil {
		return m.SpendingPools
	}
	return nil
}

func (m *ProposalCollectiveUpdate) GetClaimStart() uint64 {
	if m != nil {
		return m.ClaimStart
	}
	return 0
}

func (m *ProposalCollectiveUpdate) GetClaimPeriod() uint64 {
	if m != nil {
		return m.ClaimPeriod
	}
	return 0
}

func (m *ProposalCollectiveUpdate) GetClaimEnd() uint64 {
	if m != nil {
		return m.ClaimEnd
	}
	return 0
}

func (m *ProposalCollectiveUpdate) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *ProposalCollectiveUpdate) GetVotePeriod() uint64 {
	if m != nil {
		return m.VotePeriod
	}
	return 0
}

func (m *ProposalCollectiveUpdate) GetVoteEnactment() uint64 {
	if m != nil {
		return m.VoteEnactment
	}
	return 0
}

// proposal to remove/delete Staking Collective from registry
type ProposalCollectiveRemove struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ProposalCollectiveRemove) Reset()         { *m = ProposalCollectiveRemove{} }
func (m *ProposalCollectiveRemove) String() string { return proto.CompactTextString(m) }
func (*ProposalCollectiveRemove) ProtoMessage()    {}
func (*ProposalCollectiveRemove) Descriptor() ([]byte, []int) {
	return fileDescriptor_6173a479da28660d, []int{2}
}
func (m *ProposalCollectiveRemove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalCollectiveRemove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalCollectiveRemove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalCollectiveRemove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalCollectiveRemove.Merge(m, src)
}
func (m *ProposalCollectiveRemove) XXX_Size() int {
	return m.Size()
}
func (m *ProposalCollectiveRemove) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalCollectiveRemove.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalCollectiveRemove proto.InternalMessageInfo

func (m *ProposalCollectiveRemove) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ProposalCollectiveSendDonation)(nil), "tsuki.collectives.ProposalCollectiveSendDonation")
	proto.RegisterType((*ProposalCollectiveUpdate)(nil), "tsuki.collectives.ProposalCollectiveUpdate")
	proto.RegisterType((*ProposalCollectiveRemove)(nil), "tsuki.collectives.ProposalCollectiveRemove")
}

func init() { proto.RegisterFile("tsuki/collectives/proposal.proto", fileDescriptor_6173a479da28660d) }

var fileDescriptor_6173a479da28660d = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe3, 0x5f, 0xf3, 0x6b, 0x9a, 0x4d, 0xff, 0xb1, 0xe2, 0xb0, 0x14, 0xc9, 0x71, 0x2b,
	0x01, 0xbe, 0xc4, 0x96, 0xca, 0xad, 0x17, 0xa4, 0xa4, 0x15, 0x42, 0x1c, 0x08, 0x8e, 0xaa, 0x4a,
	0x5c, 0xa2, 0xad, 0x77, 0x95, 0xac, 0x62, 0xef, 0x18, 0xef, 0x26, 0x85, 0xb7, 0xe0, 0x11, 0x38,
	0x73, 0xe6, 0x21, 0x2a, 0x4e, 0x3d, 0x22, 0x0e, 0x15, 0x4a, 0x0e, 0xf0, 0x18, 0xc8, 0x6b, 0xa7,
	0x75, 0xe3, 0x9c, 0xbc, 0xfb, 0x9d, 0xcf, 0x7e, 0x67, 0x3c, 0x9a, 0x41, 0xed, 0x89, 0x48, 0xa9,
	0x1f, 0x42, 0x14, 0xf1, 0x50, 0x8b, 0x19, 0x57, 0x7e, 0x92, 0x42, 0x02, 0x8a, 0x46, 0x5e, 0x92,
	0x82, 0x06, 0xbc, 0x9f, 0x01, 0x5e, 0x09, 0x38, 0x78, 0x3c, 0x82, 0x11, 0x98, 0xa0, 0x9f, 0x9d,
	0x72, 0xee, 0xe0, 0x49, 0x08, 0x2a, 0x06, 0x35, 0xcc, 0x03, 0xf9, 0xa5, 0x08, 0x1d, 0x56, 0x72,
	0xdc, 0x9f, 0x73, 0xe4, 0xe8, 0x9b, 0x85, 0xec, 0x7e, 0x91, 0xb8, 0x77, 0x17, 0x1c, 0x70, 0xc9,
	0x4e, 0x41, 0x52, 0x2d, 0x40, 0x62, 0x8c, 0xea, 0x92, 0xc6, 0x9c, 0x58, 0x8e, 0xe5, 0x36, 0x03,
	0x73, 0xc6, 0x04, 0x35, 0x28, 0x63, 0x29, 0x57, 0x8a, 0xfc, 0x67, 0xe4, 0xe5, 0x15, 0xbf, 0x41,
	0x0d, 0x1a, 0xc3, 0x54, 0x6a, 0x45, 0x36, 0x9c, 0x0d, 0xb7, 0xd9, 0xf5, 0xaf, 0x6f, 0xdb, 0xb5,
	0x5f, 0xb7, 0xed, 0x17, 0x23, 0xa1, 0xc7, 0xd3, 0x4b, 0x2f, 0x84, 0xb8, 0xa8, 0xb2, 0xf8, 0x74,
	0x14, 0x9b, 0xf8, 0xfa, 0x73, 0xc2, 0x95, 0xd7, 0x03, 0x21, 0x83, 0xe5, 0xfb, 0x93, 0xbd, 0xbf,
	0x5f, 0xdb, 0xd6, 0x8f, 0xef, 0x9d, 0x46, 0x0f, 0xa4, 0xe6, 0x52, 0x1f, 0xfd, 0xa9, 0x23, 0x52,
	0x2d, 0xf6, 0x3c, 0x61, 0x54, 0xf3, 0xb5, 0x65, 0x3a, 0xa8, 0xc5, 0xb8, 0x0a, 0x53, 0x91, 0x64,
	0x7f, 0x52, 0x94, 0x5a, 0x96, 0xf0, 0x09, 0xda, 0x54, 0x9a, 0xea, 0x69, 0x56, 0xad, 0xe5, 0xee,
	0x1e, 0x1f, 0x79, 0xab, 0x6d, 0xf7, 0x4a, 0x6d, 0x31, 0x64, 0x50, 0xbc, 0xc0, 0xe7, 0xe8, 0x11,
	0xe3, 0x09, 0x28, 0xa1, 0x87, 0x57, 0x63, 0xa1, 0x79, 0x24, 0x94, 0x26, 0x75, 0xc7, 0x72, 0x5b,
	0xeb, 0x6c, 0x4e, 0x73, 0xf4, 0x62, 0x49, 0x76, 0xeb, 0x59, 0x63, 0x82, 0x7d, 0xb6, 0xa2, 0xe3,
	0x00, 0xed, 0xc3, 0x95, 0xe4, 0xa9, 0x2a, 0xb9, 0xfe, 0x6f, 0x5c, 0x0f, 0xab, 0xae, 0xef, 0x0c,
	0xb9, 0x6a, 0xba, 0x07, 0x0f, 0x65, 0x3c, 0x40, 0xbb, 0x2a, 0xe1, 0x92, 0x09, 0x39, 0x1a, 0x26,
	0x00, 0x91, 0x22, 0x9b, 0xce, 0x86, 0xdb, 0x3a, 0x7e, 0x5e, 0x75, 0xbc, 0xe0, 0x62, 0x34, 0xd6,
	0x9c, 0x0d, 0x0a, 0xbe, 0x0f, 0x10, 0x15, 0xb6, 0x3b, 0xaa, 0xa4, 0x29, 0xdc, 0x46, 0xad, 0x30,
	0xa2, 0x22, 0x1e, 0x2a, 0x4d, 0x53, 0x4d, 0x1a, 0x8e, 0xe5, 0xd6, 0x03, 0x64, 0xa4, 0x41, 0xa6,
	0xe0, 0x43, 0xb4, 0x9d, 0x03, 0x09, 0x4f, 0x05, 0x30, 0xb2, 0x65, 0x88, 0xfc, 0x51, 0xdf, 0x48,
	0xf8, 0x29, 0x6a, 0xe6, 0x08, 0x97, 0x8c, 0x34, 0x4d, 0x7c, 0xcb, 0x08, 0x67, 0x92, 0x65, 0x09,
	0x66, 0xa0, 0xf9, 0xf0, 0xe3, 0x14, 0xd2, 0x69, 0x4c, 0x50, 0x9e, 0x20, 0x93, 0xde, 0x1b, 0xe5,
	0x0e, 0x28, 0xfc, 0x5b, 0xf7, 0x40, 0x61, 0xff, 0x0c, 0xed, 0x1a, 0x80, 0x4b, 0x1a, 0xea, 0x98,
	0x4b, 0x4d, 0xb6, 0x0d, 0xb3, 0x93, 0xa9, 0x67, 0x4b, 0xb1, 0x3a, 0x69, 0xaf, 0xd6, 0x0d, 0x5a,
	0xc0, 0x63, 0x98, 0xad, 0x1d, 0xb4, 0x8a, 0x41, 0xf7, 0xf5, 0xf5, 0xdc, 0xb6, 0x6e, 0xe6, 0xb6,
	0xf5, 0x7b, 0x6e, 0x5b, 0x5f, 0x16, 0x76, 0xed, 0x66, 0x61, 0xd7, 0x7e, 0x2e, 0xec, 0xda, 0x87,
	0x4e, 0x69, 0x0f, 0xde, 0x8a, 0x94, 0xf6, 0x20, 0xe5, 0xbe, 0xe2, 0x13, 0x2a, 0xfc, 0x4f, 0x0f,
	0x76, 0xd5, 0xac, 0xc4, 0xe5, 0xa6, 0xd9, 0xd3, 0x97, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xeb,
	0x04, 0x0f, 0xe5, 0x30, 0x04, 0x00, 0x00,
}

func (this *ProposalCollectiveSendDonation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProposalCollectiveSendDonation)
	if !ok {
		that2, ok := that.(ProposalCollectiveSendDonation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if len(this.Amounts) != len(that1.Amounts) {
		return false
	}
	for i := range this.Amounts {
		if !this.Amounts[i].Equal(that1.Amounts[i]) {
			return false
		}
	}
	return true
}
func (this *ProposalCollectiveUpdate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProposalCollectiveUpdate)
	if !ok {
		that2, ok := that.(ProposalCollectiveUpdate)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !this.DepositWhitelist.Equal(&that1.DepositWhitelist) {
		return false
	}
	if !this.OwnersWhitelist.Equal(&that1.OwnersWhitelist) {
		return false
	}
	if len(this.SpendingPools) != len(that1.SpendingPools) {
		return false
	}
	for i := range this.SpendingPools {
		if !this.SpendingPools[i].Equal(&that1.SpendingPools[i]) {
			return false
		}
	}
	if this.ClaimStart != that1.ClaimStart {
		return false
	}
	if this.ClaimPeriod != that1.ClaimPeriod {
		return false
	}
	if this.ClaimEnd != that1.ClaimEnd {
		return false
	}
	if this.VoteQuorum != that1.VoteQuorum {
		return false
	}
	if this.VotePeriod != that1.VotePeriod {
		return false
	}
	if this.VoteEnactment != that1.VoteEnactment {
		return false
	}
	return true
}
func (this *ProposalCollectiveRemove) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProposalCollectiveRemove)
	if !ok {
		that2, ok := that.(ProposalCollectiveRemove)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (m *ProposalCollectiveSendDonation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalCollectiveSendDonation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalCollectiveSendDonation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amounts) > 0 {
		for iNdEx := len(m.Amounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Amounts[iNdEx].Size()
				i -= size
				if _, err := m.Amounts[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProposalCollectiveUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalCollectiveUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalCollectiveUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoteEnactment != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VoteEnactment))
		i--
		dAtA[i] = 0x60
	}
	if m.VotePeriod != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VotePeriod))
		i--
		dAtA[i] = 0x58
	}
	if m.VoteQuorum != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VoteQuorum))
		i--
		dAtA[i] = 0x50
	}
	if m.ClaimEnd != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ClaimEnd))
		i--
		dAtA[i] = 0x48
	}
	if m.ClaimPeriod != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ClaimPeriod))
		i--
		dAtA[i] = 0x40
	}
	if m.ClaimStart != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ClaimStart))
		i--
		dAtA[i] = 0x38
	}
	if len(m.SpendingPools) > 0 {
		for iNdEx := len(m.SpendingPools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendingPools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size, err := m.OwnersWhitelist.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.DepositWhitelist.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Status != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProposalCollectiveRemove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalCollectiveRemove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalCollectiveRemove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposalCollectiveSendDonation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Amounts) > 0 {
		for _, e := range m.Amounts {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *ProposalCollectiveUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovProposal(uint64(m.Status))
	}
	l = m.DepositWhitelist.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.OwnersWhitelist.Size()
	n += 1 + l + sovProposal(uint64(l))
	if len(m.SpendingPools) > 0 {
		for _, e := range m.SpendingPools {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if m.ClaimStart != 0 {
		n += 1 + sovProposal(uint64(m.ClaimStart))
	}
	if m.ClaimPeriod != 0 {
		n += 1 + sovProposal(uint64(m.ClaimPeriod))
	}
	if m.ClaimEnd != 0 {
		n += 1 + sovProposal(uint64(m.ClaimEnd))
	}
	if m.VoteQuorum != 0 {
		n += 1 + sovProposal(uint64(m.VoteQuorum))
	}
	if m.VotePeriod != 0 {
		n += 1 + sovProposal(uint64(m.VotePeriod))
	}
	if m.VoteEnactment != 0 {
		n += 1 + sovProposal(uint64(m.VoteEnactment))
	}
	return n
}

func (m *ProposalCollectiveRemove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalCollectiveSendDonation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProposalCollectiveSendDonation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalCollectiveSendDonation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.Amounts = append(m.Amounts, v)
			if err := m.Amounts[len(m.Amounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProposalCollectiveUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProposalCollectiveUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalCollectiveUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CollectiveStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositWhitelist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DepositWhitelist.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnersWhitelist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnersWhitelist.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendingPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendingPools = append(m.SpendingPools, WeightedSpendingPool{})
			if err := m.SpendingPools[len(m.SpendingPools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimStart", wireType)
			}
			m.ClaimStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimPeriod", wireType)
			}
			m.ClaimPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimEnd", wireType)
			}
			m.ClaimEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimEnd |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteQuorum", wireType)
			}
			m.VoteQuorum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteQuorum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotePeriod", wireType)
			}
			m.VotePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteEnactment", wireType)
			}
			m.VoteEnactment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteEnactment |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProposalCollectiveRemove) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProposalCollectiveRemove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalCollectiveRemove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
