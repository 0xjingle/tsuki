// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/staking/staking.proto

package types

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgClaimValidator struct {
	Moniker string                                        `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	ValKey  github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,2,opt,name=val_key,json=valKey,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"val_key,omitempty" yaml:"val_key"`
	PubKey  *types.Any                                    `protobuf:"bytes,3,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty" yaml:"pub_key"`
}

func (m *MsgClaimValidator) Reset()         { *m = MsgClaimValidator{} }
func (m *MsgClaimValidator) String() string { return proto.CompactTextString(m) }
func (*MsgClaimValidator) ProtoMessage()    {}
func (*MsgClaimValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a0e4e3724532dc, []int{0}
}
func (m *MsgClaimValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimValidator.Merge(m, src)
}
func (m *MsgClaimValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimValidator proto.InternalMessageInfo

func (m *MsgClaimValidator) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *MsgClaimValidator) GetValKey() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValKey
	}
	return nil
}

func (m *MsgClaimValidator) GetPubKey() *types.Any {
	if m != nil {
		return m.PubKey
	}
	return nil
}

// MsgClaimValidatorResponse defines the Msg/ClaimValidator response type.
type MsgClaimValidatorResponse struct {
}

func (m *MsgClaimValidatorResponse) Reset()         { *m = MsgClaimValidatorResponse{} }
func (m *MsgClaimValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimValidatorResponse) ProtoMessage()    {}
func (*MsgClaimValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a0e4e3724532dc, []int{1}
}
func (m *MsgClaimValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimValidatorResponse.Merge(m, src)
}
func (m *MsgClaimValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimValidatorResponse proto.InternalMessageInfo

type Validator struct {
	ValKey github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=val_key,json=valKey,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"val_key,omitempty" yaml:"val_key"`
	PubKey *types.Any                                    `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty" yaml:"pub_key"`
	Status ValidatorStatus                               `protobuf:"varint,3,opt,name=status,proto3,enum=tsuki.staking.ValidatorStatus" json:"status,omitempty"`
	// To judge validator performance a streak and rank properties should be created (as part of each validator status data).
	// The streak would imply consecutive number of times that given validator managed to successfully propose a block (since the last time he failed) that was accepted into the blockchain state. The streak property should be zeroed every time validator misses to propose a block and the mischance property is incremented. You can treat streak in similar way to kill-streaks in video games - which imply your short term performance.
	// The rank property is a long term statistics implying the "longest" streak that validator ever achieved, it can be expressed as rank = MAX(rank, streak). Under certain circumstances we should however decrease the rank of the validator. If the mischance property is incremented, the rank should be decremented by X (default 10), that is rank = MAX(rank - X, 0). Every time node status changes to inactive the rank should be divided by 2, that is rank = FLOOR(rank / 2)
	// The streak and rank will enable governance to judge real life performance of validators on the mainnet or testnet, and potentially propose eviction of the weakest and least reliable operators.
	Rank   int64 `protobuf:"varint,4,opt,name=rank,proto3" json:"rank,omitempty"`
	Streak int64 `protobuf:"varint,5,opt,name=streak,proto3" json:"streak,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a0e4e3724532dc, []int{2}
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
	return fileDescriptor_84a0e4e3724532dc, []int{3}
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
	proto.RegisterType((*MsgClaimValidator)(nil), "tsuki.staking.MsgClaimValidator")
	proto.RegisterType((*MsgClaimValidatorResponse)(nil), "tsuki.staking.MsgClaimValidatorResponse")
	proto.RegisterType((*Validator)(nil), "tsuki.staking.Validator")
	proto.RegisterType((*ValidatorJailInfo)(nil), "tsuki.staking.ValidatorJailInfo")
}

func init() { proto.RegisterFile("tsuki/staking/staking.proto", fileDescriptor_84a0e4e3724532dc) }

var fileDescriptor_84a0e4e3724532dc = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xce, 0x25, 0x21, 0x6d, 0xae, 0xa5, 0xa4, 0xa7, 0xaa, 0x4a, 0x5d, 0xb0, 0xad, 0x2c, 0x84,
	0x4a, 0xb5, 0xa5, 0x20, 0x24, 0xd4, 0x2d, 0x5f, 0x48, 0x69, 0x69, 0x55, 0xb9, 0x1f, 0x12, 0x15,
	0x52, 0x75, 0x89, 0xaf, 0xe6, 0xe4, 0x8f, 0xb3, 0x7c, 0x4e, 0x85, 0xff, 0x01, 0xea, 0xd4, 0x95,
	0xa1, 0x12, 0x12, 0x33, 0x1b, 0x3f, 0xa2, 0x62, 0xea, 0xc8, 0x54, 0x50, 0xbb, 0x30, 0x30, 0x31,
	0x32, 0xa1, 0x9c, 0xcf, 0xfd, 0xca, 0xc0, 0x02, 0xd3, 0xf9, 0xbd, 0xe7, 0x79, 0x5f, 0xbf, 0xcf,
	0xa3, 0xe7, 0xa0, 0xe2, 0xd2, 0x08, 0x9b, 0x3c, 0xc6, 0x2e, 0x0d, 0x9c, 0xec, 0x34, 0xc2, 0x88,
	0xc5, 0x0c, 0x4d, 0x8f, 0x30, 0x43, 0xde, 0x29, 0x73, 0x0e, 0x73, 0x98, 0x00, 0xcc, 0xd1, 0x57,
	0xca, 0x51, 0x16, 0x1c, 0xc6, 0x1c, 0x8f, 0x98, 0xa2, 0xea, 0x0f, 0x0f, 0x4c, 0x1c, 0x24, 0x19,
	0x34, 0x60, 0xdc, 0x67, 0x7c, 0x3f, 0xed, 0x49, 0x0b, 0x09, 0x69, 0x77, 0xbb, 0x62, 0xea, 0x13,
	0x1e, 0x63, 0x3f, 0x94, 0x84, 0xc5, 0x5b, 0x6b, 0x85, 0x11, 0x0b, 0x19, 0xc7, 0x5e, 0x0a, 0xd6,
	0x7e, 0x02, 0x38, 0xbb, 0xce, 0x9d, 0xb6, 0x87, 0xa9, 0xbf, 0x8b, 0x3d, 0x6a, 0xe3, 0x98, 0x45,
	0xa8, 0x0a, 0x27, 0x7c, 0x16, 0x50, 0x97, 0x44, 0x55, 0xa0, 0x83, 0x7a, 0xd9, 0xca, 0x4a, 0xf4,
	0x1a, 0x4e, 0x1c, 0x62, 0x6f, 0xdf, 0x25, 0x49, 0x35, 0xaf, 0x83, 0xfa, 0x74, 0xab, 0xfd, 0xeb,
	0x5c, 0x9b, 0x49, 0xb0, 0xef, 0xad, 0xd4, 0x24, 0x50, 0xfb, 0x7d, 0xae, 0x2d, 0x3b, 0x34, 0x7e,
	0x33, 0xec, 0x1b, 0x03, 0xe6, 0xcb, 0x6d, 0xe5, 0xb1, 0xcc, 0x6d, 0xd7, 0x8c, 0x93, 0x90, 0x70,
	0x63, 0x17, 0x7b, 0x4d, 0xdb, 0x8e, 0x08, 0xe7, 0x56, 0xe9, 0x10, 0x7b, 0x6b, 0x24, 0x41, 0xaf,
	0xe0, 0x44, 0x38, 0xec, 0x8b, 0xe9, 0x05, 0x1d, 0xd4, 0xa7, 0x1a, 0x73, 0x46, 0xaa, 0xce, 0xc8,
	0xd4, 0x19, 0xcd, 0x20, 0x69, 0x2d, 0x5d, 0xff, 0x53, 0xd2, 0x6b, 0x5f, 0x3e, 0x2f, 0xcf, 0x49,
	0x5b, 0x06, 0x51, 0x12, 0xc6, 0xcc, 0xd8, 0x1c, 0xf6, 0xd7, 0x48, 0x62, 0x95, 0x42, 0x71, 0xae,
	0x14, 0x7f, 0x7c, 0xd0, 0x40, 0x6d, 0x11, 0x2e, 0x8c, 0xa9, 0xb5, 0x08, 0x0f, 0x59, 0xc0, 0x49,
	0xed, 0x53, 0x1e, 0x96, 0xaf, 0x3d, 0xb8, 0xa1, 0x14, 0xfc, 0x57, 0xa5, 0xf9, 0x7f, 0xab, 0x14,
	0x3d, 0x83, 0x25, 0x1e, 0xe3, 0x78, 0xc8, 0x85, 0x87, 0x33, 0x8d, 0x47, 0xc6, 0xcd, 0xec, 0x19,
	0x57, 0x0a, 0xb7, 0x04, 0xc9, 0x92, 0x64, 0x84, 0x60, 0x31, 0xc2, 0x81, 0x5b, 0x2d, 0xea, 0xa0,
	0x5e, 0xb0, 0xc4, 0x37, 0x9a, 0x1f, 0x8d, 0x8a, 0x08, 0x76, 0xab, 0xf7, 0xc4, 0xad, 0xac, 0xa4,
	0x99, 0xeb, 0x70, 0xf6, 0x6a, 0xd8, 0x2a, 0xa6, 0x5e, 0x2f, 0x38, 0x60, 0xe8, 0x39, 0x2c, 0x8e,
	0x02, 0x28, 0x3c, 0x9b, 0x6a, 0x28, 0x63, 0xaa, 0xb6, 0xb3, 0x74, 0xb6, 0x26, 0x4f, 0xcf, 0xb5,
	0xdc, 0xf1, 0x37, 0x0d, 0x58, 0xa2, 0x63, 0xe9, 0x3d, 0x80, 0x0f, 0xee, 0x2c, 0x87, 0x1e, 0xc2,
	0xf2, 0xce, 0x46, 0xa7, 0xfb, 0xa2, 0xb7, 0xd1, 0xed, 0x54, 0x72, 0xca, 0xfd, 0xa3, 0x13, 0xbd,
	0xbc, 0x13, 0xd8, 0xe4, 0x80, 0x06, 0xc4, 0x1e, 0xad, 0xd7, 0x6c, 0x6f, 0xf7, 0x76, 0xbb, 0x15,
	0xa0, 0xc0, 0xa3, 0x13, 0xbd, 0xd4, 0x1c, 0xc4, 0xf4, 0x90, 0x20, 0x05, 0x4e, 0xf6, 0x36, 0x24,
	0x92, 0x57, 0xa6, 0x8f, 0x4e, 0xf4, 0xc9, 0x5e, 0x80, 0x53, 0x6c, 0x1e, 0x96, 0x36, 0x9b, 0x3b,
	0x5b, 0xdd, 0x4e, 0xa5, 0x90, 0xf6, 0x6c, 0xe2, 0x21, 0x4f, 0x67, 0xad, 0x36, 0x7b, 0x2f, 0xbb,
	0x9d, 0x4a, 0x31, 0xbd, 0x1f, 0x29, 0x22, 0xb6, 0x52, 0x7c, 0xf7, 0x51, 0xcd, 0x35, 0x30, 0x2c,
	0xac, 0x73, 0x07, 0xed, 0xc1, 0x99, 0x3b, 0x2f, 0x45, 0xbb, 0x6d, 0xee, 0x58, 0xb8, 0x94, 0xc7,
	0x7f, 0x21, 0x64, 0xe9, 0x6b, 0xb5, 0x4f, 0x2f, 0x54, 0x70, 0x76, 0xa1, 0x82, 0xef, 0x17, 0x2a,
	0x38, 0xbe, 0x54, 0x73, 0x67, 0x97, 0x6a, 0xee, 0xeb, 0xa5, 0x9a, 0xdb, 0x7b, 0x72, 0x23, 0x62,
	0x6b, 0x34, 0xc2, 0x6d, 0x16, 0x11, 0x93, 0x13, 0x17, 0x53, 0xf3, 0xed, 0xd5, 0xbb, 0x16, 0x49,
	0xeb, 0x97, 0x84, 0xcf, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xbe, 0x7f, 0xae, 0x8b,
	0x04, 0x00, 0x00,
}

func (this *MsgClaimValidator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgClaimValidator)
	if !ok {
		that2, ok := that.(MsgClaimValidator)
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
	if this.Moniker != that1.Moniker {
		return false
	}
	if !bytes.Equal(this.ValKey, that1.ValKey) {
		return false
	}
	if !this.PubKey.Equal(that1.PubKey) {
		return false
	}
	return true
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
	if this.Rank != that1.Rank {
		return false
	}
	if this.Streak != that1.Streak {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ClaimValidator defines a method for claiming a new validator.
	ClaimValidator(ctx context.Context, in *MsgClaimValidator, opts ...grpc.CallOption) (*MsgClaimValidatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ClaimValidator(ctx context.Context, in *MsgClaimValidator, opts ...grpc.CallOption) (*MsgClaimValidatorResponse, error) {
	out := new(MsgClaimValidatorResponse)
	err := c.cc.Invoke(ctx, "/tsuki.staking.Msg/ClaimValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ClaimValidator defines a method for claiming a new validator.
	ClaimValidator(context.Context, *MsgClaimValidator) (*MsgClaimValidatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ClaimValidator(ctx context.Context, req *MsgClaimValidator) (*MsgClaimValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimValidator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ClaimValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tsuki.staking.Msg/ClaimValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimValidator(ctx, req.(*MsgClaimValidator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tsuki.staking.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimValidator",
			Handler:    _Msg_ClaimValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tsuki/staking/staking.proto",
}

func (m *MsgClaimValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
		dAtA[i] = 0x1a
	}
	if len(m.ValKey) > 0 {
		i -= len(m.ValKey)
		copy(dAtA[i:], m.ValKey)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
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
		dAtA[i] = 0x28
	}
	if m.Rank != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Rank))
		i--
		dAtA[i] = 0x20
	}
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
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintStaking(dAtA, i, uint64(n3))
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
func (m *MsgClaimValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.ValKey)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func (m *MsgClaimValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
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
func (m *MsgClaimValidator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
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
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
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
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) < 0 {
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
func (m *MsgClaimValidatorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) < 0 {
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
		case 5:
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
			if skippy < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) < 0 {
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
			if skippy < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) < 0 {
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
