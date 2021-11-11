// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tsuki/gov/actor.proto

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

type ActorStatus int32

const (
	// Undefined status
	ActorStatus_UNDEFINED ActorStatus = 0
	// Unclaimed status
	ActorStatus_UNCLAIMED ActorStatus = 1
	// Active status
	ActorStatus_ACTIVE ActorStatus = 2
	// Paused status
	ActorStatus_PAUSED ActorStatus = 3
	// Inactive status
	ActorStatus_INACTIVE ActorStatus = 4
	// Jailed status
	ActorStatus_JAILED ActorStatus = 5
	// Removed status
	ActorStatus_REMOVED ActorStatus = 6
)

var ActorStatus_name = map[int32]string{
	0: "UNDEFINED",
	1: "UNCLAIMED",
	2: "ACTIVE",
	3: "PAUSED",
	4: "INACTIVE",
	5: "JAILED",
	6: "REMOVED",
}

var ActorStatus_value = map[string]int32{
	"UNDEFINED": 0,
	"UNCLAIMED": 1,
	"ACTIVE":    2,
	"PAUSED":    3,
	"INACTIVE":  4,
	"JAILED":    5,
	"REMOVED":   6,
}

func (x ActorStatus) String() string {
	return proto.EnumName(ActorStatus_name, int32(x))
}

func (ActorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{0}
}

type Permissions struct {
	Blacklist            []uint32 `protobuf:"varint,1,rep,packed,name=blacklist,proto3" json:"blacklist,omitempty"`
	Whitelist            []uint32 `protobuf:"varint,2,rep,packed,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{0}
}

func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
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

type NetworkActor struct {
	Address              []byte       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Roles                []uint64     `protobuf:"varint,2,rep,packed,name=roles,proto3" json:"roles,omitempty"`
	Status               ActorStatus  `protobuf:"varint,3,opt,name=status,proto3,enum=tsuki.gov.ActorStatus" json:"status,omitempty"`
	Votes                []VoteOption `protobuf:"varint,4,rep,packed,name=votes,proto3,enum=tsuki.gov.VoteOption" json:"votes,omitempty"`
	Permissions          *Permissions `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Skin                 uint64       `protobuf:"varint,6,opt,name=skin,proto3" json:"skin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NetworkActor) Reset()         { *m = NetworkActor{} }
func (m *NetworkActor) String() string { return proto.CompactTextString(m) }
func (*NetworkActor) ProtoMessage()    {}
func (*NetworkActor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{1}
}

func (m *NetworkActor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkActor.Unmarshal(m, b)
}
func (m *NetworkActor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkActor.Marshal(b, m, deterministic)
}
func (m *NetworkActor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkActor.Merge(m, src)
}
func (m *NetworkActor) XXX_Size() int {
	return xxx_messageInfo_NetworkActor.Size(m)
}
func (m *NetworkActor) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkActor.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkActor proto.InternalMessageInfo

func (m *NetworkActor) GetAddress() []byte {
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

func (m *NetworkActor) GetStatus() ActorStatus {
	if m != nil {
		return m.Status
	}
	return ActorStatus_UNDEFINED
}

func (m *NetworkActor) GetVotes() []VoteOption {
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

type MsgWhitelistPermissions struct {
	Proposer             []byte   `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgWhitelistPermissions) Reset()         { *m = MsgWhitelistPermissions{} }
func (m *MsgWhitelistPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgWhitelistPermissions) ProtoMessage()    {}
func (*MsgWhitelistPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{2}
}

func (m *MsgWhitelistPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWhitelistPermissions.Unmarshal(m, b)
}
func (m *MsgWhitelistPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWhitelistPermissions.Marshal(b, m, deterministic)
}
func (m *MsgWhitelistPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWhitelistPermissions.Merge(m, src)
}
func (m *MsgWhitelistPermissions) XXX_Size() int {
	return xxx_messageInfo_MsgWhitelistPermissions.Size(m)
}
func (m *MsgWhitelistPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWhitelistPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWhitelistPermissions proto.InternalMessageInfo

func (m *MsgWhitelistPermissions) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgWhitelistPermissions) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgWhitelistPermissions) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

type MsgRemoveWhitelistedPermissions struct {
	Proposer             []byte   `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRemoveWhitelistedPermissions) Reset()         { *m = MsgRemoveWhitelistedPermissions{} }
func (m *MsgRemoveWhitelistedPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveWhitelistedPermissions) ProtoMessage()    {}
func (*MsgRemoveWhitelistedPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{3}
}

func (m *MsgRemoveWhitelistedPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRemoveWhitelistedPermissions.Unmarshal(m, b)
}
func (m *MsgRemoveWhitelistedPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRemoveWhitelistedPermissions.Marshal(b, m, deterministic)
}
func (m *MsgRemoveWhitelistedPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveWhitelistedPermissions.Merge(m, src)
}
func (m *MsgRemoveWhitelistedPermissions) XXX_Size() int {
	return xxx_messageInfo_MsgRemoveWhitelistedPermissions.Size(m)
}
func (m *MsgRemoveWhitelistedPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveWhitelistedPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveWhitelistedPermissions proto.InternalMessageInfo

func (m *MsgRemoveWhitelistedPermissions) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgRemoveWhitelistedPermissions) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgRemoveWhitelistedPermissions) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

type MsgBlacklistPermissions struct {
	Proposer             []byte   `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBlacklistPermissions) Reset()         { *m = MsgBlacklistPermissions{} }
func (m *MsgBlacklistPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgBlacklistPermissions) ProtoMessage()    {}
func (*MsgBlacklistPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{4}
}

func (m *MsgBlacklistPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBlacklistPermissions.Unmarshal(m, b)
}
func (m *MsgBlacklistPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBlacklistPermissions.Marshal(b, m, deterministic)
}
func (m *MsgBlacklistPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBlacklistPermissions.Merge(m, src)
}
func (m *MsgBlacklistPermissions) XXX_Size() int {
	return xxx_messageInfo_MsgBlacklistPermissions.Size(m)
}
func (m *MsgBlacklistPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBlacklistPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBlacklistPermissions proto.InternalMessageInfo

func (m *MsgBlacklistPermissions) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgBlacklistPermissions) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgBlacklistPermissions) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

type MsgRemoveBlacklistedPermissions struct {
	Proposer             []byte   `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRemoveBlacklistedPermissions) Reset()         { *m = MsgRemoveBlacklistedPermissions{} }
func (m *MsgRemoveBlacklistedPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveBlacklistedPermissions) ProtoMessage()    {}
func (*MsgRemoveBlacklistedPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{5}
}

func (m *MsgRemoveBlacklistedPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRemoveBlacklistedPermissions.Unmarshal(m, b)
}
func (m *MsgRemoveBlacklistedPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRemoveBlacklistedPermissions.Marshal(b, m, deterministic)
}
func (m *MsgRemoveBlacklistedPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveBlacklistedPermissions.Merge(m, src)
}
func (m *MsgRemoveBlacklistedPermissions) XXX_Size() int {
	return xxx_messageInfo_MsgRemoveBlacklistedPermissions.Size(m)
}
func (m *MsgRemoveBlacklistedPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveBlacklistedPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveBlacklistedPermissions proto.InternalMessageInfo

func (m *MsgRemoveBlacklistedPermissions) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgRemoveBlacklistedPermissions) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgRemoveBlacklistedPermissions) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

func init() {
	proto.RegisterEnum("tsuki.gov.ActorStatus", ActorStatus_name, ActorStatus_value)
	proto.RegisterType((*Permissions)(nil), "tsuki.gov.Permissions")
	proto.RegisterType((*NetworkActor)(nil), "tsuki.gov.NetworkActor")
	proto.RegisterType((*MsgWhitelistPermissions)(nil), "tsuki.gov.MsgWhitelistPermissions")
	proto.RegisterType((*MsgRemoveWhitelistedPermissions)(nil), "tsuki.gov.MsgRemoveWhitelistedPermissions")
	proto.RegisterType((*MsgBlacklistPermissions)(nil), "tsuki.gov.MsgBlacklistPermissions")
	proto.RegisterType((*MsgRemoveBlacklistedPermissions)(nil), "tsuki.gov.MsgRemoveBlacklistedPermissions")
}

func init() {
	proto.RegisterFile("tsuki/gov/actor.proto", fileDescriptor_c4584595efc3386e)
}

var fileDescriptor_c4584595efc3386e = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x8d, 0x6c, 0xd9, 0x71, 0xc6, 0x49, 0x30, 0x83, 0xdb, 0x08, 0x51, 0x62, 0xa1, 0x95, 0x09,
	0xd8, 0xa2, 0x09, 0xa5, 0x90, 0x9d, 0x62, 0xab, 0xa0, 0x34, 0x76, 0x82, 0xf2, 0x68, 0x29, 0x74,
	0x31, 0x91, 0xa6, 0xca, 0xa0, 0xc7, 0x08, 0xcd, 0xc4, 0x21, 0x7f, 0x50, 0xfc, 0x0f, 0x86, 0x42,
	0x97, 0xfd, 0x96, 0xfe, 0x42, 0xb6, 0xed, 0xa6, 0x14, 0x0a, 0xdd, 0x74, 0x55, 0xa4, 0xb1, 0x64,
	0x6d, 0xbb, 0xe8, 0x26, 0x2b, 0x8f, 0xcf, 0xb9, 0xf7, 0x72, 0xcf, 0x39, 0x68, 0x06, 0x74, 0x03,
	0x92, 0x22, 0xc3, 0xa7, 0x33, 0x03, 0xb9, 0x9c, 0xa6, 0xc3, 0x24, 0xa5, 0x9c, 0xc2, 0x56, 0x86,
	0x0e, 0x7d, 0x3a, 0x53, 0xbb, 0x3e, 0xf5, 0x69, 0x0e, 0x1a, 0xd9, 0x49, 0xf0, 0xea, 0x4e, 0xd9,
	0x95, 0xa4, 0x34, 0xa1, 0x0c, 0x85, 0x82, 0xd0, 0x6d, 0xd0, 0x3e, 0xc3, 0x69, 0x44, 0x18, 0x23,
	0x34, 0x66, 0xf0, 0x19, 0xd8, 0xb8, 0x0e, 0x91, 0x1b, 0x84, 0x84, 0x71, 0x45, 0xd2, 0xea, 0xfd,
	0x2d, 0x67, 0x05, 0x64, 0xec, 0xdd, 0x0d, 0xe1, 0x38, 0x67, 0x6b, 0x82, 0x2d, 0x01, 0xfd, 0x4b,
	0x0d, 0x6c, 0x4e, 0x31, 0xbf, 0xa3, 0x69, 0x60, 0x66, 0xab, 0xc1, 0xf7, 0x60, 0x1d, 0x79, 0x5e,
	0x8a, 0x19, 0x53, 0x24, 0x4d, 0xea, 0x6f, 0x1e, 0x8d, 0x7e, 0x3d, 0xf4, 0xb6, 0xef, 0x51, 0x14,
	0x1e, 0xea, 0x4b, 0x42, 0xff, 0xf3, 0xd0, 0x1b, 0xf8, 0x84, 0xdf, 0xdc, 0x5e, 0x0f, 0x5d, 0x1a,
	0x19, 0x2e, 0x65, 0x11, 0x65, 0xcb, 0x9f, 0x01, 0xf3, 0x02, 0x83, 0xdf, 0x27, 0x98, 0x0d, 0x4d,
	0xd7, 0x35, 0x45, 0x87, 0x53, 0xcc, 0x84, 0x5d, 0xd0, 0x48, 0x69, 0x88, 0x59, 0xbe, 0x89, 0xec,
	0x88, 0x3f, 0x70, 0x00, 0x9a, 0x8c, 0x23, 0x7e, 0xcb, 0x94, 0xba, 0x26, 0xf5, 0xb7, 0xf7, 0x9f,
	0x0c, 0x0b, 0x6b, 0x86, 0xf9, 0x56, 0xe7, 0x39, 0xe9, 0x2c, 0x8b, 0xe0, 0x1e, 0x68, 0xcc, 0x28,
	0xc7, 0x4c, 0x91, 0xb5, 0x7a, 0x7f, 0x7b, 0xbf, 0xbb, 0xaa, 0xbe, 0xa2, 0x1c, 0x9f, 0x26, 0x9c,
	0xd0, 0xd8, 0x11, 0x25, 0xf0, 0x25, 0x68, 0x27, 0x2b, 0xaf, 0x94, 0x86, 0x26, 0xf5, 0xdb, 0xd5,
	0xf9, 0x15, 0x23, 0x9d, 0x6a, 0x25, 0x84, 0x40, 0x66, 0x01, 0x89, 0x95, 0xa6, 0x26, 0xf5, 0x65,
	0x27, 0x3f, 0xeb, 0x3f, 0x24, 0xb0, 0x33, 0x61, 0xfe, 0x9b, 0xc2, 0xbe, 0x6a, 0x0a, 0x13, 0xd0,
	0x12, 0x31, 0xe1, 0x74, 0xe9, 0xdc, 0xf3, 0x7f, 0xf7, 0xa9, 0x1c, 0x51, 0xcd, 0xa1, 0xf6, 0x1f,
	0x72, 0xd8, 0x05, 0x60, 0x25, 0x36, 0x77, 0x7d, 0xcb, 0xa9, 0x20, 0x87, 0xf2, 0xf7, 0x4f, 0x3d,
	0x49, 0xff, 0x2d, 0x81, 0xde, 0x84, 0xf9, 0x0e, 0x8e, 0xe8, 0x0c, 0x97, 0xaa, 0xb1, 0xf7, 0xd8,
	0x75, 0x7f, 0x13, 0x39, 0x1f, 0x15, 0x1f, 0xd1, 0xa3, 0xd5, 0xab, 0xff, 0xac, 0x26, 0x5c, 0xea,
	0x7d, 0xc4, 0x09, 0xef, 0x7d, 0x95, 0x40, 0xbb, 0x72, 0xa9, 0x64, 0xf7, 0xe3, 0xe5, 0x74, 0x6c,
	0xbd, 0xb2, 0xa7, 0xd6, 0xb8, 0xb3, 0xa6, 0x6e, 0xcd, 0x17, 0xda, 0xc6, 0x65, 0xec, 0xe1, 0x0f,
	0x24, 0xc6, 0x9e, 0x60, 0x47, 0x27, 0xa6, 0x3d, 0xb1, 0xc6, 0x1d, 0xa9, 0x60, 0xdd, 0x10, 0x91,
	0x08, 0x7b, 0xf0, 0x29, 0x68, 0x9a, 0xa3, 0x0b, 0xfb, 0xca, 0xea, 0xd4, 0x54, 0x30, 0x5f, 0x68,
	0x4d, 0xd3, 0xe5, 0x64, 0x86, 0x33, 0xfc, 0xcc, 0xbc, 0x3c, 0xb7, 0xc6, 0x9d, 0xba, 0xc0, 0xcf,
	0xd0, 0x2d, 0xc3, 0x1e, 0x54, 0x41, 0xcb, 0x9e, 0x2e, 0x3b, 0x64, 0x75, 0x73, 0xbe, 0xd0, 0x5a,
	0x76, 0x8c, 0xca, 0x9e, 0x63, 0xd3, 0x3e, 0xb1, 0xc6, 0x9d, 0x86, 0xe8, 0x39, 0x46, 0x24, 0xc4,
	0x1e, 0x54, 0xc0, 0xba, 0x63, 0x4d, 0x4e, 0xaf, 0xac, 0x71, 0xa7, 0xa9, 0xb6, 0xe7, 0x0b, 0x6d,
	0x5d, 0x84, 0xe5, 0xa9, 0xf2, 0xc7, 0xcf, 0xbb, 0x6b, 0x47, 0x2f, 0xde, 0x1d, 0x54, 0x9c, 0x7a,
	0x4d, 0x52, 0x34, 0xa2, 0x29, 0x36, 0x18, 0x0e, 0x10, 0x31, 0xec, 0xe9, 0x85, 0xe5, 0xbc, 0x35,
	0xf2, 0x67, 0x63, 0xe0, 0xe3, 0xd8, 0x28, 0x9e, 0x94, 0xeb, 0x66, 0x8e, 0x1d, 0xfc, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x64, 0xae, 0xf9, 0x9b, 0x06, 0x00, 0x00,
}
