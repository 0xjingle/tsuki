// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/staking/staking.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorStatus int32

const (
	// Undefined status
	Undefined ValidatorStatus = 0
	// Active status
	Active ValidatorStatus = 1
	// Inactive status
	Inactive ValidatorStatus = 2
	// Paused status
	Paused ValidatorStatus = 3
	// Jailed status
	Jailed ValidatorStatus = 4
)

var ValidatorStatus_name = map[int32]string{
	0: "UNDEFINED",
	1: "ACTIVE",
	2: "INACTIVE",
	3: "PAUSED",
	4: "JAILED",
}

var ValidatorStatus_value = map[string]int32{
	"UNDEFINED": 0,
	"ACTIVE":    1,
	"INACTIVE":  2,
	"PAUSED":    3,
	"JAILED":    4,
}

func (x ValidatorStatus) String() string {
	return proto.EnumName(ValidatorStatus_name, int32(x))
}

func (ValidatorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84a0e4e3724532dc, []int{0}
}

type Validator struct {
	ValKey github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=val_key,json=valKey,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"val_key,omitempty" yaml:"val_key"`
	PubKey *types.Any                                    `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty" yaml:"pub_key"`
	Status ValidatorStatus                               `protobuf:"varint,3,opt,name=status,proto3,enum=tsuki.staking.ValidatorStatus" json:"status,omitempty"`
	// percentage (e.g. 0.1 == 10%) commission that the validator charges from all block reward
	Commission github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=commission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission"`
	// To judge validator performance a streak and rank properties should be created (as part of each validator status data).
	// The streak would imply consecutive number of times that given validator managed to successfully propose a block (since the last time he failed) that was accepted into the blockchain state. The streak property should be zeroed every time validator misses to propose a block and the mischance property is incremented. You can treat streak in similar way to kill-streaks in video games - which imply your short term performance.
	// The rank property is a long term statistics implying the "longest" streak that validator ever achieved, it can be expressed as rank = MAX(rank, streak). Under certain circumstances we should however decrease the rank of the validator. If the mischance property is incremented, the rank should be decremented by X (default 10), that is rank = MAX(rank - X, 0). Every time node status changes to inactive the rank should be divided by 2, that is rank = FLOOR(rank / 2)
	// The streak and rank will enable governance to judge real life performance of validators on the mainnet or testnet, and potentially propose eviction of the weakest and least reliable operators.
	Rank   int64 `protobuf:"varint,5,opt,name=rank,proto3" json:"rank,omitempty"`
	Streak int64 `protobuf:"varint,6,opt,name=streak,proto3" json:"streak,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a0e4e3724532dc, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetValKey() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValKey
	}
	return nil
}

func (m *Validator) GetPubKey() *types.Any {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *Validator) GetStatus() ValidatorStatus {
	if m != nil {
		return m.Status
	}
	return Undefined
}

func (m *Validator) GetRank() int64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *Validator) GetStreak() int64 {
	if m != nil {
		return m.Streak
	}
	return 0
}

type ValidatorJailInfo struct {
	Time time.Time `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time"`
}

func (m *ValidatorJailInfo) Reset()         { *m = ValidatorJailInfo{} }
func (m *ValidatorJailInfo) String() string { return proto.CompactTextString(m) }
func (*ValidatorJailInfo) ProtoMessage()    {}
func (*ValidatorJailInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a0e4e3724532dc, []int{1}
}
func (m *ValidatorJailInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorJailInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorJailInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorJailInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorJailInfo.Merge(m, src)
}
func (m *ValidatorJailInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorJailInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorJailInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorJailInfo proto.InternalMessageInfo

func (m *ValidatorJailInfo) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("tsuki.staking.ValidatorStatus", ValidatorStatus_name, ValidatorStatus_value)
	proto.RegisterType((*Validator)(nil), "tsuki.staking.Validator")
	proto.RegisterType((*ValidatorJailInfo)(nil), "tsuki.staking.ValidatorJailInfo")
}

func init() { proto.RegisterFile("tsuki/staking/staking.proto", fileDescriptor_84a0e4e3724532dc) }

var fileDescriptor_84a0e4e3724532dc = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4b, 0x6b, 0xd4, 0x50,
	0x14, 0xce, 0x6d, 0x63, 0xda, 0xb9, 0xad, 0x75, 0x0c, 0xa5, 0x8c, 0x51, 0x93, 0x30, 0x0b, 0x19,
	0x0b, 0x4d, 0xa0, 0x22, 0x48, 0x77, 0x99, 0x87, 0x90, 0x56, 0x87, 0x92, 0x3e, 0x40, 0x11, 0xca,
	0x9d, 0xe4, 0x36, 0x5e, 0xf2, 0xb8, 0x21, 0x37, 0x29, 0xe6, 0x1f, 0x48, 0x57, 0xdd, 0xba, 0x18,
	0x10, 0xfc, 0x0b, 0xfe, 0x88, 0xe2, 0xaa, 0x4b, 0x71, 0x31, 0xca, 0xcc, 0xc6, 0xb5, 0x4b, 0x41,
	0x90, 0x3c, 0x66, 0x1c, 0xeb, 0xc6, 0xd5, 0xb9, 0xe7, 0x7c, 0xe7, 0x9c, 0x7c, 0xdf, 0x77, 0x02,
	0x25, 0x8f, 0xc4, 0x48, 0x67, 0x09, 0xf2, 0x48, 0xe8, 0x4e, 0xa3, 0x16, 0xc5, 0x34, 0xa1, 0xe2,
	0x6a, 0x8e, 0x69, 0x55, 0x4d, 0x5a, 0x77, 0xa9, 0x4b, 0x0b, 0x40, 0xcf, 0x5f, 0x65, 0x8f, 0x74,
	0xc7, 0xa5, 0xd4, 0xf5, 0xb1, 0x5e, 0x64, 0x83, 0xf4, 0x54, 0x47, 0x61, 0x36, 0x85, 0x6c, 0xca,
	0x02, 0xca, 0x4e, 0xca, 0x99, 0x32, 0xa9, 0x20, 0xe5, 0xfa, 0x54, 0x42, 0x02, 0xcc, 0x12, 0x14,
	0x44, 0x55, 0xc3, 0xdd, 0xbf, 0x68, 0x45, 0x31, 0x8d, 0x28, 0x43, 0x7e, 0x09, 0x36, 0x7f, 0x2d,
	0xc0, 0xda, 0x31, 0xf2, 0x89, 0x83, 0x12, 0x1a, 0x8b, 0xaf, 0xe0, 0xd2, 0x19, 0xf2, 0x4f, 0x3c,
	0x9c, 0x35, 0x80, 0x0a, 0x5a, 0xab, 0xed, 0xce, 0x8f, 0x91, 0xb2, 0x96, 0xa1, 0xc0, 0xdf, 0x69,
	0x56, 0x40, 0xf3, 0xe7, 0x48, 0xd9, 0x72, 0x49, 0xf2, 0x3a, 0x1d, 0x68, 0x36, 0x0d, 0x2a, 0x2e,
	0x55, 0xd8, 0x62, 0x8e, 0xa7, 0x27, 0x59, 0x84, 0x99, 0x76, 0x8c, 0x7c, 0xc3, 0x71, 0x62, 0xcc,
	0x98, 0x25, 0x9c, 0x21, 0x7f, 0x0f, 0x67, 0xe2, 0x0b, 0xb8, 0x14, 0xa5, 0x83, 0x62, 0xfb, 0x82,
	0x0a, 0x5a, 0x2b, 0xdb, 0xeb, 0x5a, 0xc9, 0x5d, 0x9b, 0x72, 0xd7, 0x8c, 0x30, 0x6b, 0x6f, 0xfe,
	0xf9, 0x66, 0xd5, 0xde, 0xfc, 0xf4, 0x71, 0x6b, 0xbd, 0x12, 0x6d, 0xc7, 0x59, 0x94, 0x50, 0x6d,
	0x3f, 0x1d, 0xec, 0xe1, 0xcc, 0x12, 0xa2, 0x22, 0x8a, 0x8f, 0xa1, 0xc0, 0x12, 0x94, 0xa4, 0xac,
	0xb1, 0xa8, 0x82, 0xd6, 0xda, 0xf6, 0x7d, 0x6d, 0xde, 0x6f, 0x6d, 0xa6, 0xf0, 0xa0, 0x68, 0xb2,
	0xaa, 0x66, 0xb1, 0x0f, 0xa1, 0x4d, 0x83, 0x80, 0x30, 0x46, 0x68, 0xd8, 0xe0, 0x55, 0xd0, 0xaa,
	0xb5, 0xb5, 0xcb, 0x91, 0xc2, 0x7d, 0x19, 0x29, 0x0f, 0xfe, 0x43, 0x64, 0x17, 0xdb, 0xd6, 0xdc,
	0x06, 0x51, 0x84, 0x7c, 0x8c, 0x42, 0xaf, 0x71, 0x43, 0x05, 0xad, 0x45, 0xab, 0x78, 0x8b, 0x1b,
	0x39, 0xb5, 0x18, 0x23, 0xaf, 0x21, 0x14, 0xd5, 0x2a, 0xdb, 0xe1, 0xbf, 0xbf, 0x57, 0x40, 0xf3,
	0x39, 0xbc, 0x3d, 0x23, 0xb7, 0x8b, 0x88, 0x6f, 0x86, 0xa7, 0x54, 0x7c, 0x02, 0xf9, 0xfc, 0x88,
	0xc5, 0x0d, 0x56, 0xb6, 0xa5, 0x7f, 0x5c, 0x3a, 0x9c, 0x5e, 0xb8, 0xbd, 0x9c, 0x93, 0xbd, 0xf8,
	0xaa, 0x00, 0xab, 0x98, 0xd8, 0x7c, 0x07, 0xe0, 0xad, 0x6b, 0x62, 0xc5, 0x7b, 0xb0, 0x76, 0xd4,
	0xef, 0xf6, 0x9e, 0x9a, 0xfd, 0x5e, 0xb7, 0xce, 0x49, 0x37, 0xcf, 0x87, 0x6a, 0xed, 0x28, 0x74,
	0xf0, 0x29, 0x09, 0xb1, 0x93, 0xd3, 0x33, 0x3a, 0x87, 0xe6, 0x71, 0xaf, 0x0e, 0x24, 0x78, 0x3e,
	0x54, 0x05, 0xc3, 0x4e, 0xc8, 0x19, 0x16, 0x25, 0xb8, 0x6c, 0xf6, 0x2b, 0x64, 0x41, 0x5a, 0x3d,
	0x1f, 0xaa, 0xcb, 0x66, 0x88, 0x4a, 0x6c, 0x03, 0x0a, 0xfb, 0xc6, 0xd1, 0x41, 0xaf, 0x5b, 0x5f,
	0x2c, 0x67, 0xf6, 0x51, 0xca, 0xca, 0x5d, 0xbb, 0x86, 0xf9, 0xac, 0xd7, 0xad, 0xf3, 0x65, 0x3d,
	0x57, 0x84, 0x1d, 0x89, 0x7f, 0xfb, 0x41, 0xe6, 0xda, 0x9d, 0xcb, 0xb1, 0x0c, 0xae, 0xc6, 0x32,
	0xf8, 0x36, 0x96, 0xc1, 0xc5, 0x44, 0xe6, 0xae, 0x26, 0x32, 0xf7, 0x79, 0x22, 0x73, 0x2f, 0x1f,
	0xce, 0x59, 0xbd, 0x47, 0x62, 0xd4, 0xa1, 0x31, 0xd6, 0x19, 0xf6, 0x10, 0xd1, 0xdf, 0xcc, 0x7e,
	0xdc, 0xc2, 0xf1, 0x81, 0x50, 0x98, 0xf0, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x68,
	0x28, 0x2e, 0x6c, 0x03, 0x00, 0x00,
}

func (this *Validator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Validator)
	if !ok {
		that2, ok := that.(Validator)
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
	if !bytes.Equal(this.ValKey, that1.ValKey) {
		return false
	}
	if !this.PubKey.Equal(that1.PubKey) {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !this.Commission.Equal(that1.Commission) {
		return false
	}
	if this.Rank != that1.Rank {
		return false
	}
	if this.Streak != that1.Streak {
		return false
	}
	return true
}
func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Streak != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Streak))
		i--
		dAtA[i] = 0x30
	}
	if m.Rank != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Rank))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Commission.Size()
		i -= size
		if _, err := m.Commission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Status != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStaking(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValKey) > 0 {
		i -= len(m.ValKey)
		copy(dAtA[i:], m.ValKey)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorJailInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorJailInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorJailInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintStaking(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintStaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovStaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValKey)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovStaking(uint64(m.Status))
	}
	l = m.Commission.Size()
	n += 1 + l + sovStaking(uint64(l))
	if m.Rank != 0 {
		n += 1 + sovStaking(uint64(m.Rank))
	}
	if m.Streak != 0 {
		n += 1 + sovStaking(uint64(m.Streak))
	}
	return n
}

func (m *ValidatorJailInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func sovStaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStaking(x uint64) (n int) {
	return sovStaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValKey = append(m.ValKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ValKey == nil {
				m.ValKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ValidatorStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Streak", wireType)
			}
			m.Streak = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Streak |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func (m *ValidatorJailInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: ValidatorJailInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorJailInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func skipStaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
				return 0, ErrInvalidLengthStaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStaking = fmt.Errorf("proto: unexpected end of group")
)
