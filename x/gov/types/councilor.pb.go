// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/gov/councilor.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type CouncilorStatus int32

const (
	// Waiting status
	CouncilorWaiting CouncilorStatus = 0
	// Active status
	CouncilorActive CouncilorStatus = 1
	// Inactive status
	CouncilorInactive CouncilorStatus = 2
	// Paused status
	CouncilorPaused CouncilorStatus = 3
	// Jailed status
	CouncilorJailed CouncilorStatus = 4
)

var CouncilorStatus_name = map[int32]string{
	0: "WAITING",
	1: "ACTIVE",
	2: "INACTIVE",
	3: "PAUSED",
	4: "JAILED",
}

var CouncilorStatus_value = map[string]int32{
	"WAITING":  0,
	"ACTIVE":   1,
	"INACTIVE": 2,
	"PAUSED":   3,
	"JAILED":   4,
}

func (x CouncilorStatus) String() string {
	return proto.EnumName(CouncilorStatus_name, int32(x))
}

func (CouncilorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cac0b13545050495, []int{0}
}

type MsgClaimCouncilor struct {
	Address     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
	Moniker     string                                        `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Username    string                                        `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Description string                                        `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Social      string                                        `protobuf:"bytes,5,opt,name=social,proto3" json:"social,omitempty"`
	Contact     string                                        `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"`
	Avatar      string                                        `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (m *MsgClaimCouncilor) Reset()         { *m = MsgClaimCouncilor{} }
func (m *MsgClaimCouncilor) String() string { return proto.CompactTextString(m) }
func (*MsgClaimCouncilor) ProtoMessage()    {}
func (*MsgClaimCouncilor) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac0b13545050495, []int{0}
}
func (m *MsgClaimCouncilor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimCouncilor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimCouncilor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimCouncilor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimCouncilor.Merge(m, src)
}
func (m *MsgClaimCouncilor) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimCouncilor) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimCouncilor.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimCouncilor proto.InternalMessageInfo

func (m *MsgClaimCouncilor) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgClaimCouncilor) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *MsgClaimCouncilor) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MsgClaimCouncilor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgClaimCouncilor) GetSocial() string {
	if m != nil {
		return m.Social
	}
	return ""
}

func (m *MsgClaimCouncilor) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

func (m *MsgClaimCouncilor) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type Councilor struct {
	Address           github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
	Status            CouncilorStatus                               `protobuf:"varint,2,opt,name=status,proto3,enum=tsuki.gov.CouncilorStatus" json:"status,omitempty"`
	Rank              int64                                         `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
	AbstentionCounter int64                                         `protobuf:"varint,4,opt,name=abstention_counter,json=abstentionCounter,proto3" json:"abstention_counter,omitempty"`
}

func (m *Councilor) Reset()         { *m = Councilor{} }
func (m *Councilor) String() string { return proto.CompactTextString(m) }
func (*Councilor) ProtoMessage()    {}
func (*Councilor) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac0b13545050495, []int{1}
}
func (m *Councilor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Councilor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Councilor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Councilor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Councilor.Merge(m, src)
}
func (m *Councilor) XXX_Size() int {
	return m.Size()
}
func (m *Councilor) XXX_DiscardUnknown() {
	xxx_messageInfo_Councilor.DiscardUnknown(m)
}

var xxx_messageInfo_Councilor proto.InternalMessageInfo

func (m *Councilor) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Councilor) GetStatus() CouncilorStatus {
	if m != nil {
		return m.Status
	}
	return CouncilorWaiting
}

func (m *Councilor) GetRank() int64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *Councilor) GetAbstentionCounter() int64 {
	if m != nil {
		return m.AbstentionCounter
	}
	return 0
}

func init() {
	proto.RegisterEnum("tsuki.gov.CouncilorStatus", CouncilorStatus_name, CouncilorStatus_value)
	proto.RegisterType((*MsgClaimCouncilor)(nil), "tsuki.gov.MsgClaimCouncilor")
	proto.RegisterType((*Councilor)(nil), "tsuki.gov.Councilor")
}

func init() { proto.RegisterFile("tsuki/gov/councilor.proto", fileDescriptor_cac0b13545050495) }

var fileDescriptor_cac0b13545050495 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0xeb, 0xb5, 0xa4, 0x9d, 0x41, 0xa3, 0x35, 0x03, 0x85, 0x1c, 0xd2, 0x50, 0x84, 0x34,
	0x21, 0x35, 0x11, 0x70, 0xdb, 0x05, 0x65, 0x59, 0x85, 0x32, 0x60, 0x9a, 0xba, 0xc1, 0x24, 0x24,
	0x84, 0x5c, 0xc7, 0x0a, 0x56, 0x9b, 0xb8, 0xb2, 0x9d, 0x8a, 0xfd, 0x03, 0xd4, 0x13, 0x67, 0xa4,
	0x4a, 0x48, 0xfc, 0x13, 0x4e, 0x1c, 0x77, 0xe4, 0x34, 0x4d, 0xed, 0x85, 0x33, 0x47, 0x4e, 0x28,
	0x4e, 0xda, 0x4d, 0xfd, 0x01, 0x9c, 0xe2, 0xd7, 0xef, 0xf3, 0x7d, 0x8e, 0xde, 0xcf, 0x86, 0xe6,
	0x90, 0x09, 0xec, 0xc5, 0x7c, 0xe2, 0x11, 0x9e, 0xa5, 0x84, 0x8d, 0xb8, 0x70, 0xc7, 0x82, 0x2b,
	0x8e, 0x1a, 0xb9, 0xe3, 0xc6, 0x7c, 0x62, 0x6d, 0xc7, 0x3c, 0xe6, 0x7a, 0xd3, 0xcb, 0x57, 0x85,
	0xdf, 0xf9, 0xba, 0x01, 0x5b, 0xaf, 0x65, 0x1c, 0x8c, 0x30, 0x4b, 0x82, 0x65, 0x2d, 0x7a, 0x0f,
	0xeb, 0x38, 0x8a, 0x04, 0x95, 0xd2, 0x04, 0x0e, 0xd8, 0xb9, 0xb5, 0x17, 0xfc, 0xb9, 0x68, 0x6f,
	0x9d, 0xe1, 0x64, 0xb4, 0xdb, 0x29, 0x8d, 0xce, 0xdf, 0x8b, 0x76, 0x37, 0x66, 0xea, 0x63, 0x36,
	0x70, 0x09, 0x4f, 0x3c, 0xc2, 0x65, 0xc2, 0x65, 0xf9, 0xe9, 0xca, 0x68, 0xe8, 0xa9, 0xb3, 0x31,
	0x95, 0xae, 0x4f, 0x88, 0x5f, 0x54, 0xf4, 0x97, 0x3d, 0x91, 0x09, 0xeb, 0x09, 0x4f, 0xd9, 0x90,
	0x0a, 0x73, 0xc3, 0x01, 0x3b, 0x9b, 0xfd, 0xa5, 0x44, 0x16, 0x6c, 0x64, 0x92, 0x8a, 0x14, 0x27,
	0xd4, 0xac, 0x6a, 0x6b, 0xa5, 0x91, 0x03, 0x6f, 0x46, 0x54, 0x12, 0xc1, 0xc6, 0x8a, 0xf1, 0xd4,
	0xac, 0x69, 0xfb, 0xfa, 0x16, 0xba, 0x07, 0x0d, 0xc9, 0x09, 0xc3, 0x23, 0xf3, 0x86, 0x36, 0x4b,
	0x95, 0x9f, 0x47, 0x78, 0xaa, 0x30, 0x51, 0xa6, 0x51, 0x9c, 0x57, 0xca, 0xbc, 0x02, 0x4f, 0xb0,
	0xc2, 0xc2, 0xac, 0x17, 0x15, 0x85, 0xda, 0xad, 0xfd, 0xfe, 0xd6, 0x06, 0x9d, 0x4b, 0x00, 0x37,
	0xff, 0x5b, 0x28, 0x4f, 0xa0, 0x21, 0x15, 0x56, 0x99, 0xd4, 0x99, 0x6c, 0x3d, 0xbd, 0xef, 0x2e,
	0x47, 0xe7, 0xae, 0xfe, 0xe1, 0x58, 0x03, 0xfd, 0x12, 0x44, 0x08, 0xd6, 0x04, 0x4e, 0x87, 0x3a,
	0xa9, 0x6a, 0x5f, 0xaf, 0x51, 0x17, 0x22, 0x3c, 0x90, 0x8a, 0xa6, 0x79, 0x22, 0x1f, 0xf2, 0xeb,
	0xa0, 0xa8, 0xd0, 0x61, 0x55, 0xfb, 0xad, 0x2b, 0x27, 0x28, 0x8c, 0xc7, 0x3f, 0x00, 0xbc, 0xbd,
	0xd6, 0x1e, 0x3d, 0x80, 0xf5, 0x53, 0x3f, 0x3c, 0x09, 0x0f, 0x5f, 0x34, 0x2b, 0xd6, 0xf6, 0x74,
	0xe6, 0x34, 0x57, 0xc4, 0x29, 0x66, 0x8a, 0xa5, 0x31, 0x6a, 0x43, 0xc3, 0x0f, 0x4e, 0xc2, 0xb7,
	0xbd, 0x26, 0xb0, 0xee, 0x4c, 0x67, 0xce, 0x55, 0x0f, 0x9f, 0x28, 0x36, 0xa1, 0xe8, 0x21, 0x6c,
	0x84, 0x87, 0x25, 0xb2, 0x61, 0xdd, 0x9d, 0xce, 0x9c, 0xd6, 0x0a, 0x09, 0x53, 0x5c, 0x40, 0x6d,
	0x68, 0x1c, 0xf9, 0x6f, 0x8e, 0x7b, 0xfb, 0xcd, 0xea, 0x5a, 0x97, 0x23, 0x9c, 0x49, 0x1a, 0xe5,
	0xc0, 0x81, 0x1f, 0xbe, 0xea, 0xed, 0x37, 0x6b, 0x6b, 0xc0, 0x01, 0x66, 0x23, 0x1a, 0x59, 0xb5,
	0xcf, 0xdf, 0xed, 0xca, 0xde, 0xf3, 0x9f, 0x73, 0x1b, 0x9c, 0xcf, 0x6d, 0x70, 0x39, 0xb7, 0xc1,
	0x97, 0x85, 0x5d, 0x39, 0x5f, 0xd8, 0x95, 0x5f, 0x0b, 0xbb, 0xf2, 0xee, 0xd1, 0xb5, 0x61, 0xbc,
	0x64, 0x02, 0x07, 0x5c, 0x50, 0x4f, 0xd2, 0x21, 0x66, 0xde, 0x27, 0xfd, 0x5e, 0xf4, 0x3c, 0x06,
	0x86, 0x7e, 0x0c, 0xcf, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x57, 0x99, 0xc7, 0x35, 0x48, 0x03,
	0x00, 0x00,
}

func (this *MsgClaimCouncilor) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgClaimCouncilor)
	if !ok {
		that2, ok := that.(MsgClaimCouncilor)
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
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if this.Moniker != that1.Moniker {
		return false
	}
	if this.Username != that1.Username {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Social != that1.Social {
		return false
	}
	if this.Contact != that1.Contact {
		return false
	}
	if this.Avatar != that1.Avatar {
		return false
	}
	return true
}
func (m *MsgClaimCouncilor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimCouncilor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimCouncilor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Avatar) > 0 {
		i -= len(m.Avatar)
		copy(dAtA[i:], m.Avatar)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Avatar)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Contact) > 0 {
		i -= len(m.Contact)
		copy(dAtA[i:], m.Contact)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Contact)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Social) > 0 {
		i -= len(m.Social)
		copy(dAtA[i:], m.Social)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Social)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Councilor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Councilor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Councilor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AbstentionCounter != 0 {
		i = encodeVarintCouncilor(dAtA, i, uint64(m.AbstentionCounter))
		i--
		dAtA[i] = 0x20
	}
	if m.Rank != 0 {
		i = encodeVarintCouncilor(dAtA, i, uint64(m.Rank))
		i--
		dAtA[i] = 0x18
	}
	if m.Status != 0 {
		i = encodeVarintCouncilor(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCouncilor(dAtA []byte, offset int, v uint64) int {
	offset -= sovCouncilor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgClaimCouncilor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	l = len(m.Social)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	l = len(m.Contact)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	return n
}

func (m *Councilor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovCouncilor(uint64(m.Status))
	}
	if m.Rank != 0 {
		n += 1 + sovCouncilor(uint64(m.Rank))
	}
	if m.AbstentionCounter != 0 {
		n += 1 + sovCouncilor(uint64(m.AbstentionCounter))
	}
	return n
}

func sovCouncilor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCouncilor(x uint64) (n int) {
	return sovCouncilor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgClaimCouncilor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCouncilor
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
			return fmt.Errorf("proto: MsgClaimCouncilor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimCouncilor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Social", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Social = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contact = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCouncilor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCouncilor
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
func (m *Councilor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCouncilor
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
			return fmt.Errorf("proto: Councilor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Councilor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CouncilorStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbstentionCounter", wireType)
			}
			m.AbstentionCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AbstentionCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCouncilor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCouncilor
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
func skipCouncilor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCouncilor
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
					return 0, ErrIntOverflowCouncilor
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
					return 0, ErrIntOverflowCouncilor
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
				return 0, ErrInvalidLengthCouncilor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCouncilor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCouncilor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCouncilor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCouncilor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCouncilor = fmt.Errorf("proto: unexpected end of group")
)
