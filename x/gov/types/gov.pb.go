// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gov.proto

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

type MsgWhitelistPermissions struct {
	Address     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
	Permissions []uint32                                      `protobuf:"varint,2,rep,packed,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *MsgWhitelistPermissions) Reset()         { *m = MsgWhitelistPermissions{} }
func (m *MsgWhitelistPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgWhitelistPermissions) ProtoMessage()    {}
func (*MsgWhitelistPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb02393240bc858d, []int{0}
}
func (m *MsgWhitelistPermissions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWhitelistPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWhitelistPermissions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWhitelistPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWhitelistPermissions.Merge(m, src)
}
func (m *MsgWhitelistPermissions) XXX_Size() int {
	return m.Size()
}
func (m *MsgWhitelistPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWhitelistPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWhitelistPermissions proto.InternalMessageInfo

func (m *MsgWhitelistPermissions) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgWhitelistPermissions) GetPermissions() []uint32 {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type MsgBlacklistPermissions struct {
	Address     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
	Permissions []uint32                                      `protobuf:"varint,2,rep,packed,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *MsgBlacklistPermissions) Reset()         { *m = MsgBlacklistPermissions{} }
func (m *MsgBlacklistPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgBlacklistPermissions) ProtoMessage()    {}
func (*MsgBlacklistPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb02393240bc858d, []int{1}
}
func (m *MsgBlacklistPermissions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBlacklistPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBlacklistPermissions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBlacklistPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBlacklistPermissions.Merge(m, src)
}
func (m *MsgBlacklistPermissions) XXX_Size() int {
	return m.Size()
}
func (m *MsgBlacklistPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBlacklistPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBlacklistPermissions proto.InternalMessageInfo

func (m *MsgBlacklistPermissions) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgBlacklistPermissions) GetPermissions() []uint32 {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type NetworkActor struct {
	Address     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
	Roles       []uint64                                      `protobuf:"varint,2,rep,packed,name=roles,proto3" json:"roles,omitempty"`
	Status      uint32                                        `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Votes       []uint32                                      `protobuf:"varint,4,rep,packed,name=votes,proto3" json:"votes,omitempty"`
	Permissions *Permissions                                  `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Skin        uint64                                        `protobuf:"varint,6,opt,name=skin,proto3" json:"skin,omitempty"`
}

func (m *NetworkActor) Reset()         { *m = NetworkActor{} }
func (m *NetworkActor) String() string { return proto.CompactTextString(m) }
func (*NetworkActor) ProtoMessage()    {}
func (*NetworkActor) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb02393240bc858d, []int{2}
}
func (m *NetworkActor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkActor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkActor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkActor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkActor.Merge(m, src)
}
func (m *NetworkActor) XXX_Size() int {
	return m.Size()
}
func (m *NetworkActor) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkActor.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkActor proto.InternalMessageInfo

func (m *NetworkActor) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *NetworkActor) GetRoles() []uint64 {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *NetworkActor) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *NetworkActor) GetVotes() []uint32 {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *NetworkActor) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *NetworkActor) GetSkin() uint64 {
	if m != nil {
		return m.Skin
	}
	return 0
}

type Permissions struct {
	Blacklist []uint32 `protobuf:"varint,1,rep,packed,name=blacklist,proto3" json:"blacklist,omitempty"`
	Whitelist []uint32 `protobuf:"varint,2,rep,packed,name=whitelist,proto3" json:"whitelist,omitempty"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb02393240bc858d, []int{3}
}
func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return m.Size()
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetBlacklist() []uint32 {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *Permissions) GetWhitelist() []uint32 {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgWhitelistPermissions)(nil), "tsuki.gov.MsgWhitelistPermissions")
	proto.RegisterType((*MsgBlacklistPermissions)(nil), "tsuki.gov.MsgBlacklistPermissions")
	proto.RegisterType((*NetworkActor)(nil), "tsuki.gov.NetworkActor")
	proto.RegisterType((*Permissions)(nil), "tsuki.gov.Permissions")
}

func init() { proto.RegisterFile("gov.proto", fileDescriptor_eb02393240bc858d) }

var fileDescriptor_eb02393240bc858d = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xf6, 0xd6, 0xb2, 0x5b, 0xaf, 0xed, 0x1e, 0x84, 0xdb, 0x8a, 0x52, 0x64, 0x21, 0x28, 0xe8,
	0x62, 0x09, 0xda, 0x43, 0xc1, 0x97, 0x62, 0xfb, 0x54, 0x4a, 0x42, 0xd0, 0x25, 0x10, 0xc8, 0x41,
	0x96, 0x17, 0x79, 0x91, 0xe4, 0x31, 0x3b, 0x6b, 0x39, 0x7e, 0x8b, 0x3c, 0x82, 0x1f, 0x27, 0x47,
	0x1f, 0x73, 0x32, 0xc1, 0xbe, 0xe4, 0x9c, 0x63, 0x0e, 0x21, 0x58, 0x3f, 0x58, 0x7e, 0x80, 0x40,
	0x4e, 0x3b, 0xf3, 0xcd, 0xee, 0xce, 0x37, 0xdf, 0x7c, 0xb4, 0x11, 0x40, 0x62, 0xcf, 0x05, 0x48,
	0x50, 0x3f, 0x85, 0x5c, 0x78, 0x76, 0x00, 0xc9, 0xf7, 0x4e, 0x00, 0x01, 0xa4, 0xa0, 0x73, 0x88,
	0xb2, 0xba, 0xb9, 0x26, 0xf4, 0xdb, 0x19, 0x06, 0x97, 0x53, 0x2e, 0x59, 0xc4, 0x51, 0x5e, 0x30,
	0x11, 0x73, 0x44, 0x0e, 0x33, 0x54, 0xaf, 0xe9, 0x47, 0x6f, 0x32, 0x11, 0x0c, 0x51, 0x23, 0x06,
	0xb1, 0x5a, 0xc3, 0xd1, 0xd3, 0xb6, 0xfb, 0x79, 0xe5, 0xc5, 0x51, 0xdf, 0xcc, 0x0b, 0xe6, 0xf3,
	0xb6, 0xdb, 0x0b, 0xb8, 0x9c, 0x2e, 0xc6, 0xb6, 0x0f, 0xb1, 0xe3, 0x03, 0xc6, 0x80, 0xf9, 0xd1,
	0xc3, 0x49, 0xe8, 0xc8, 0xd5, 0x9c, 0xa1, 0x3d, 0xf0, 0xfd, 0x41, 0xf6, 0xc2, 0x2d, 0xfe, 0x54,
	0x0d, 0xda, 0x9c, 0x1f, 0xbb, 0x69, 0x1f, 0x8c, 0xaa, 0xd5, 0x76, 0xcb, 0x50, 0x5f, 0x79, 0x5c,
	0x77, 0x49, 0x41, 0x71, 0x18, 0x79, 0x7e, 0xf8, 0x4e, 0x29, 0xbe, 0x10, 0xda, 0x3a, 0x67, 0x72,
	0x09, 0x22, 0x1c, 0xf8, 0x12, 0xc4, 0x5b, 0xf3, 0xea, 0xd0, 0x9a, 0x80, 0x88, 0x65, 0x8c, 0x14,
	0x37, 0x4b, 0xd4, 0xaf, 0xb4, 0x8e, 0xd2, 0x93, 0x0b, 0xd4, 0xaa, 0x06, 0xb1, 0xda, 0x6e, 0x9e,
	0x1d, 0x6e, 0x27, 0x20, 0x19, 0x6a, 0x4a, 0xca, 0x3f, 0x4b, 0xd4, 0x3f, 0xa7, 0xb3, 0xd5, 0x0c,
	0x62, 0x35, 0x7f, 0x7d, 0xb1, 0x0b, 0xbf, 0xd8, 0x25, 0x99, 0x4f, 0x46, 0x56, 0x55, 0xaa, 0x60,
	0xc8, 0x67, 0x5a, 0xdd, 0x20, 0x96, 0xe2, 0xa6, 0xb1, 0xf9, 0x8f, 0x36, 0xcb, 0x6b, 0xf9, 0x41,
	0x1b, 0xe3, 0x62, 0x5d, 0x1a, 0x49, 0xbb, 0x1e, 0x81, 0x43, 0x75, 0x59, 0xf8, 0x2d, 0xd7, 0xf4,
	0x08, 0x0c, 0xff, 0xde, 0xed, 0x74, 0xb2, 0xd9, 0xe9, 0xe4, 0x61, 0xa7, 0x93, 0xdb, 0xbd, 0x5e,
	0xd9, 0xec, 0xf5, 0xca, 0xfd, 0x5e, 0xaf, 0x5c, 0xfd, 0x2c, 0xa9, 0xf5, 0x9f, 0x0b, 0x6f, 0x04,
	0x82, 0x39, 0xc8, 0x42, 0x8f, 0x3b, 0x37, 0x4e, 0x00, 0x49, 0x26, 0xd8, 0xb8, 0x9e, 0x3a, 0xfb,
	0xf7, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x80, 0x8e, 0xf1, 0x06, 0x03, 0x00, 0x00,
}

func (this *MsgWhitelistPermissions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgWhitelistPermissions)
	if !ok {
		that2, ok := that.(MsgWhitelistPermissions)
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
	if len(this.Permissions) != len(that1.Permissions) {
		return false
	}
	for i := range this.Permissions {
		if this.Permissions[i] != that1.Permissions[i] {
			return false
		}
	}
	return true
}
func (this *MsgBlacklistPermissions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgBlacklistPermissions)
	if !ok {
		that2, ok := that.(MsgBlacklistPermissions)
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
	if len(this.Permissions) != len(that1.Permissions) {
		return false
	}
	for i := range this.Permissions {
		if this.Permissions[i] != that1.Permissions[i] {
			return false
		}
	}
	return true
}
func (m *MsgWhitelistPermissions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWhitelistPermissions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWhitelistPermissions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		dAtA2 := make([]byte, len(m.Permissions)*10)
		var j1 int
		for _, num := range m.Permissions {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGov(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBlacklistPermissions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBlacklistPermissions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBlacklistPermissions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		dAtA4 := make([]byte, len(m.Permissions)*10)
		var j3 int
		for _, num := range m.Permissions {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintGov(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NetworkActor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkActor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkActor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Skin != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.Skin))
		i--
		dAtA[i] = 0x30
	}
	if m.Permissions != nil {
		{
			size, err := m.Permissions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Votes) > 0 {
		dAtA7 := make([]byte, len(m.Votes)*10)
		var j6 int
		for _, num := range m.Votes {
			for num >= 1<<7 {
				dAtA7[j6] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j6++
			}
			dAtA7[j6] = uint8(num)
			j6++
		}
		i -= j6
		copy(dAtA[i:], dAtA7[:j6])
		i = encodeVarintGov(dAtA, i, uint64(j6))
		i--
		dAtA[i] = 0x22
	}
	if m.Status != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Roles) > 0 {
		dAtA9 := make([]byte, len(m.Roles)*10)
		var j8 int
		for _, num := range m.Roles {
			for num >= 1<<7 {
				dAtA9[j8] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j8++
			}
			dAtA9[j8] = uint8(num)
			j8++
		}
		i -= j8
		copy(dAtA[i:], dAtA9[:j8])
		i = encodeVarintGov(dAtA, i, uint64(j8))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Permissions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Permissions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Permissions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Whitelist) > 0 {
		dAtA11 := make([]byte, len(m.Whitelist)*10)
		var j10 int
		for _, num := range m.Whitelist {
			for num >= 1<<7 {
				dAtA11[j10] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j10++
			}
			dAtA11[j10] = uint8(num)
			j10++
		}
		i -= j10
		copy(dAtA[i:], dAtA11[:j10])
		i = encodeVarintGov(dAtA, i, uint64(j10))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Blacklist) > 0 {
		dAtA13 := make([]byte, len(m.Blacklist)*10)
		var j12 int
		for _, num := range m.Blacklist {
			for num >= 1<<7 {
				dAtA13[j12] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j12++
			}
			dAtA13[j12] = uint8(num)
			j12++
		}
		i -= j12
		copy(dAtA[i:], dAtA13[:j12])
		i = encodeVarintGov(dAtA, i, uint64(j12))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgWhitelistPermissions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Permissions) > 0 {
		l = 0
		for _, e := range m.Permissions {
			l += sovGov(uint64(e))
		}
		n += 1 + sovGov(uint64(l)) + l
	}
	return n
}

func (m *MsgBlacklistPermissions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Permissions) > 0 {
		l = 0
		for _, e := range m.Permissions {
			l += sovGov(uint64(e))
		}
		n += 1 + sovGov(uint64(l)) + l
	}
	return n
}

func (m *NetworkActor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Roles) > 0 {
		l = 0
		for _, e := range m.Roles {
			l += sovGov(uint64(e))
		}
		n += 1 + sovGov(uint64(l)) + l
	}
	if m.Status != 0 {
		n += 1 + sovGov(uint64(m.Status))
	}
	if len(m.Votes) > 0 {
		l = 0
		for _, e := range m.Votes {
			l += sovGov(uint64(e))
		}
		n += 1 + sovGov(uint64(l)) + l
	}
	if m.Permissions != nil {
		l = m.Permissions.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	if m.Skin != 0 {
		n += 1 + sovGov(uint64(m.Skin))
	}
	return n
}

func (m *Permissions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Blacklist) > 0 {
		l = 0
		for _, e := range m.Blacklist {
			l += sovGov(uint64(e))
		}
		n += 1 + sovGov(uint64(l)) + l
	}
	if len(m.Whitelist) > 0 {
		l = 0
		for _, e := range m.Whitelist {
			l += sovGov(uint64(e))
		}
		n += 1 + sovGov(uint64(l)) + l
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgWhitelistPermissions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: MsgWhitelistPermissions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWhitelistPermissions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Permissions = append(m.Permissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGov
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGov
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Permissions) == 0 {
					m.Permissions = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGov
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Permissions = append(m.Permissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *MsgBlacklistPermissions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: MsgBlacklistPermissions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBlacklistPermissions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Permissions = append(m.Permissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGov
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGov
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Permissions) == 0 {
					m.Permissions = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGov
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Permissions = append(m.Permissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *NetworkActor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: NetworkActor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkActor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Roles = append(m.Roles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGov
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGov
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Roles) == 0 {
					m.Roles = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGov
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Roles = append(m.Roles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Votes = append(m.Votes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGov
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGov
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Votes) == 0 {
					m.Votes = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGov
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Votes = append(m.Votes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Permissions == nil {
				m.Permissions = &Permissions{}
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skin", wireType)
			}
			m.Skin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Skin |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *Permissions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: Permissions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Permissions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Blacklist = append(m.Blacklist, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGov
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGov
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Blacklist) == 0 {
					m.Blacklist = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGov
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Blacklist = append(m.Blacklist, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Blacklist", wireType)
			}
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Whitelist = append(m.Whitelist, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGov
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGov
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Whitelist) == 0 {
					m.Whitelist = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGov
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Whitelist = append(m.Whitelist, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Whitelist", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
