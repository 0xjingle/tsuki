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
	// Undefined status
	CouncilorUndefined CouncilorStatus = 0
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
	0: "UNDEFINED",
	1: "ACTIVE",
	2: "INACTIVE",
	3: "PAUSED",
	4: "JAILED",
}

var CouncilorStatus_value = map[string]int32{
	"UNDEFINED": 0,
	"ACTIVE":    1,
	"INACTIVE":  2,
	"PAUSED":    3,
	"JAILED":    4,
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
	return CouncilorUndefined
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
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x8b, 0xd3, 0x4e,
	0x18, 0x6d, 0xda, 0xfe, 0xd2, 0x76, 0x7e, 0xb2, 0xb6, 0xa3, 0x2e, 0x31, 0x87, 0x34, 0x54, 0x16,
	0x16, 0xa1, 0x09, 0xea, 0x6d, 0x2f, 0x92, 0x4d, 0x2b, 0x74, 0xd5, 0xb2, 0x74, 0xad, 0x07, 0x41,
	0x64, 0x3a, 0x19, 0xe3, 0xd0, 0x66, 0xa6, 0xcc, 0x4c, 0x8b, 0xfb, 0x1f, 0x48, 0x4f, 0x9e, 0x85,
	0x82, 0xe0, 0xff, 0x22, 0x1e, 0xf7, 0xe8, 0x69, 0x59, 0xda, 0x8b, 0x67, 0x8f, 0x9e, 0x24, 0x93,
	0xb4, 0xbb, 0xf4, 0x0f, 0xf0, 0x94, 0x79, 0xf3, 0xde, 0xf7, 0xbe, 0xf0, 0xbe, 0xf9, 0x80, 0x35,
	0xa6, 0x02, 0xf9, 0x31, 0x9f, 0xfb, 0x98, 0xcf, 0x18, 0xa6, 0x13, 0x2e, 0xbc, 0xa9, 0xe0, 0x8a,
	0xc3, 0x6a, 0xca, 0x78, 0x31, 0x9f, 0xdb, 0x77, 0x63, 0x1e, 0x73, 0x7d, 0xe9, 0xa7, 0xa7, 0x8c,
	0x6f, 0x7d, 0x29, 0x82, 0xc6, 0x4b, 0x19, 0x87, 0x13, 0x44, 0x93, 0x70, 0x53, 0x0b, 0xdf, 0x82,
	0x0a, 0x8a, 0x22, 0x41, 0xa4, 0xb4, 0x0c, 0xd7, 0x38, 0xbc, 0x75, 0x1c, 0xfe, 0xbe, 0x6c, 0xee,
	0x9d, 0xa3, 0x64, 0x72, 0xd4, 0xca, 0x89, 0xd6, 0x9f, 0xcb, 0x66, 0x3b, 0xa6, 0xea, 0xc3, 0x6c,
	0xe4, 0x61, 0x9e, 0xf8, 0x98, 0xcb, 0x84, 0xcb, 0xfc, 0xd3, 0x96, 0xd1, 0xd8, 0x57, 0xe7, 0x53,
	0x22, 0xbd, 0x00, 0xe3, 0x20, 0xab, 0x18, 0x6c, 0x3c, 0xa1, 0x05, 0x2a, 0x09, 0x67, 0x74, 0x4c,
	0x84, 0x55, 0x74, 0x8d, 0xc3, 0xda, 0x60, 0x03, 0xa1, 0x0d, 0xaa, 0x33, 0x49, 0x04, 0x43, 0x09,
	0xb1, 0x4a, 0x9a, 0xda, 0x62, 0xe8, 0x82, 0xff, 0x23, 0x22, 0xb1, 0xa0, 0x53, 0x45, 0x39, 0xb3,
	0xca, 0x9a, 0xbe, 0x79, 0x05, 0xf7, 0x81, 0x29, 0x39, 0xa6, 0x68, 0x62, 0xfd, 0xa7, 0xc9, 0x1c,
	0xa5, 0xfd, 0x30, 0x67, 0x0a, 0x61, 0x65, 0x99, 0x59, 0xbf, 0x1c, 0xa6, 0x15, 0x68, 0x8e, 0x14,
	0x12, 0x56, 0x25, 0xab, 0xc8, 0xd0, 0x51, 0xf9, 0xd7, 0xd7, 0xa6, 0xd1, 0xba, 0x32, 0x40, 0xed,
	0x9f, 0x85, 0xf2, 0x08, 0x98, 0x52, 0x21, 0x35, 0x93, 0x3a, 0x93, 0xbd, 0xc7, 0xf7, 0xbd, 0xcd,
	0xe8, 0xbc, 0xed, 0x3f, 0x9c, 0x69, 0xc1, 0x20, 0x17, 0x42, 0x08, 0xca, 0x02, 0xb1, 0xb1, 0x4e,
	0xaa, 0x34, 0xd0, 0x67, 0xd8, 0x06, 0x10, 0x8d, 0xa4, 0x22, 0x2c, 0x4d, 0xe4, 0x5d, 0xfa, 0x1c,
	0x14, 0x11, 0x3a, 0xac, 0xd2, 0xa0, 0x71, 0xcd, 0x84, 0x19, 0xf1, 0xf0, 0xbb, 0x01, 0x6e, 0xef,
	0xd8, 0xc3, 0x03, 0x50, 0x1b, 0xf6, 0x3b, 0xdd, 0x67, 0xbd, 0x7e, 0xb7, 0x53, 0x2f, 0xd8, 0xfb,
	0x8b, 0xa5, 0x0b, 0xb7, 0x9a, 0x21, 0x8b, 0xc8, 0x7b, 0xca, 0x48, 0x04, 0x9b, 0xc0, 0x0c, 0xc2,
	0x57, 0xbd, 0xd7, 0xdd, 0xba, 0x61, 0xdf, 0x59, 0x2c, 0xdd, 0x6b, 0x9f, 0x00, 0x2b, 0x3a, 0x27,
	0xf0, 0x01, 0xa8, 0xf6, 0xfa, 0xb9, 0xa4, 0x68, 0xdf, 0x5b, 0x2c, 0xdd, 0xc6, 0x56, 0xd2, 0x63,
	0x28, 0x13, 0x35, 0x81, 0x79, 0x1a, 0x0c, 0xcf, 0xba, 0x9d, 0x7a, 0x69, 0xc7, 0xe5, 0x14, 0xcd,
	0x64, 0xd6, 0xe6, 0x24, 0xe8, 0xbd, 0xe8, 0x76, 0xea, 0xe5, 0x1d, 0xc1, 0x09, 0xa2, 0x13, 0x12,
	0xd9, 0xe5, 0x4f, 0xdf, 0x9c, 0xc2, 0xf1, 0xd3, 0x1f, 0x2b, 0xc7, 0xb8, 0x58, 0x39, 0xc6, 0xd5,
	0xca, 0x31, 0x3e, 0xaf, 0x9d, 0xc2, 0xc5, 0xda, 0x29, 0xfc, 0x5c, 0x3b, 0x85, 0x37, 0x07, 0x37,
	0x06, 0xf2, 0x9c, 0x0a, 0x14, 0x72, 0x41, 0x7c, 0x49, 0xc6, 0x88, 0xfa, 0x1f, 0xf5, 0xce, 0xe8,
	0x99, 0x8c, 0x4c, 0xbd, 0x10, 0x4f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xab, 0x19, 0x41, 0xc8,
	0x4c, 0x03, 0x00, 0x00,
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
