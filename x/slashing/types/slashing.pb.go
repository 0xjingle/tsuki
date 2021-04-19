// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slashing.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// ValidatorSigningInfo defines a validator's signing info for monitoring their
// liveness activity.
type ValidatorSigningInfo struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// height at which validator was first a candidate OR was activated
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty" yaml:"start_height"`
	// timestamp validator cannot be activated until
	InactiveUntil time.Time `protobuf:"bytes,3,opt,name=inactive_until,json=inactiveUntil,proto3,stdtime" json:"inactive_until" yaml:"inactive_until"`
	// whether or not a validator has been tombstoned (killed out of validator
	// set)
	Tombstoned bool `protobuf:"varint,4,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	// missed blocks sequentially
	Mischance int64 `protobuf:"varint,5,opt,name=mischance,proto3" json:"mischance,omitempty"`
	// last signed block height by the validator
	LastPresentBlock int64 `protobuf:"varint,6,opt,name=last_present_block,json=lastPresentBlock,proto3" json:"last_present_block,omitempty"`
	// missed blocks counter (to avoid scanning the array every time)
	MissedBlocksCounter int64 `protobuf:"varint,7,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty"`
	// count produced blocks so far by a validator
	ProducedBlocksCounter int64 `protobuf:"varint,8,opt,name=produced_blocks_counter,json=producedBlocksCounter,proto3" json:"produced_blocks_counter,omitempty"`
}

func (m *ValidatorSigningInfo) Reset()      { *m = ValidatorSigningInfo{} }
func (*ValidatorSigningInfo) ProtoMessage() {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f622956ca78100, []int{0}
}
func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetInactiveUntil() time.Time {
	if m != nil {
		return m.InactiveUntil
	}
	return time.Time{}
}

func (m *ValidatorSigningInfo) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *ValidatorSigningInfo) GetMischance() int64 {
	if m != nil {
		return m.Mischance
	}
	return 0
}

func (m *ValidatorSigningInfo) GetLastPresentBlock() int64 {
	if m != nil {
		return m.LastPresentBlock
	}
	return 0
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

func (m *ValidatorSigningInfo) GetProducedBlocksCounter() int64 {
	if m != nil {
		return m.ProducedBlocksCounter
	}
	return 0
}

// Params represents the parameters used for by the slashing module.
type Params struct {
	MaxMischance             int64         `protobuf:"varint,2,opt,name=max_mischance,json=maxMischance,proto3" json:"max_mischance,omitempty" yaml:"max_mischance"`
	DowntimeInactiveDuration time.Duration `protobuf:"bytes,3,opt,name=downtime_inactive_duration,json=downtimeInactiveDuration,proto3,stdduration" json:"downtime_inactive_duration" yaml:"downtime_inactive_duration"`
	// the number of blocks validator miss to start counting mischance
	// default 10, modifiable by governance
	MischanceConfidence int64 `protobuf:"varint,4,opt,name=mischance_confidence,json=mischanceConfidence,proto3" json:"mischance_confidence,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f622956ca78100, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxMischance() int64 {
	if m != nil {
		return m.MaxMischance
	}
	return 0
}

func (m *Params) GetDowntimeInactiveDuration() time.Duration {
	if m != nil {
		return m.DowntimeInactiveDuration
	}
	return 0
}

func (m *Params) GetMischanceConfidence() int64 {
	if m != nil {
		return m.MischanceConfidence
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidatorSigningInfo)(nil), "tsuki.slashing.ValidatorSigningInfo")
	proto.RegisterType((*Params)(nil), "tsuki.slashing.Params")
}

func init() { proto.RegisterFile("slashing.proto", fileDescriptor_31f622956ca78100) }

var fileDescriptor_31f622956ca78100 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xc7, 0x73, 0x6d, 0x49, 0xd3, 0x6b, 0x53, 0xa1, 0x6b, 0xa2, 0x9a, 0x08, 0xd9, 0xa9, 0xa7,
	0x08, 0x81, 0x2d, 0x8a, 0xc4, 0x10, 0x89, 0x25, 0x61, 0x68, 0x85, 0x90, 0x2a, 0xf3, 0x32, 0xb0,
	0x58, 0x17, 0xfb, 0xe2, 0x9c, 0x62, 0xdf, 0x45, 0x77, 0x67, 0x48, 0x67, 0x24, 0xe6, 0x8e, 0x1d,
	0x3b, 0xf2, 0x01, 0xf8, 0x10, 0x1d, 0x3b, 0x32, 0x05, 0x94, 0x2c, 0xcc, 0xf9, 0x04, 0xc8, 0xbe,
	0xd8, 0x6d, 0xa8, 0xd8, 0xf2, 0xfc, 0x7f, 0xff, 0xe7, 0xc9, 0xf3, 0x72, 0x86, 0xfb, 0x32, 0xc6,
	0x72, 0x44, 0x59, 0xe4, 0x4c, 0x04, 0x57, 0x1c, 0xd5, 0xc7, 0x54, 0x60, 0xa7, 0x10, 0x5b, 0x8d,
	0x88, 0x47, 0x3c, 0x27, 0x6e, 0xf6, 0x4b, 0x9b, 0x5a, 0x66, 0xc4, 0x79, 0x14, 0x13, 0x37, 0x8f,
	0x06, 0xe9, 0xd0, 0x0d, 0x53, 0x81, 0x15, 0xe5, 0x6c, 0xc5, 0xad, 0x7f, 0xb9, 0xa2, 0x09, 0x91,
	0x0a, 0x27, 0x13, 0x6d, 0xb0, 0x7f, 0x6c, 0xc2, 0xc6, 0x47, 0x1c, 0xd3, 0x10, 0x2b, 0x2e, 0xde,
	0xd1, 0x88, 0x51, 0x16, 0x9d, 0xb2, 0x21, 0x47, 0x06, 0xdc, 0xc6, 0x61, 0x28, 0x88, 0x94, 0x06,
	0x68, 0x83, 0xce, 0x8e, 0x57, 0x84, 0xa8, 0x0b, 0xf7, 0xa4, 0xc2, 0x42, 0xf9, 0x23, 0x42, 0xa3,
	0x91, 0x32, 0x36, 0xda, 0xa0, 0xb3, 0xd9, 0x3b, 0x5c, 0xce, 0xac, 0x83, 0x73, 0x9c, 0xc4, 0x5d,
	0xfb, 0x2e, 0xb5, 0xbd, 0xdd, 0x3c, 0x3c, 0xc9, 0x23, 0x14, 0xc2, 0x7d, 0xca, 0x70, 0xa0, 0xe8,
	0x67, 0xe2, 0xa7, 0x4c, 0xd1, 0xd8, 0xd8, 0x6c, 0x83, 0xce, 0xee, 0x71, 0xcb, 0xd1, 0x8d, 0x3a,
	0x45, 0xa3, 0xce, 0xfb, 0xa2, 0xd1, 0xde, 0xd1, 0xf5, 0xcc, 0xaa, 0x2c, 0x67, 0x56, 0x53, 0x57,
	0x5f, 0xcf, 0xb7, 0x2f, 0x7e, 0x59, 0xc0, 0xab, 0x17, 0xe2, 0x87, 0x4c, 0x43, 0x26, 0x84, 0x8a,
	0x27, 0x03, 0xa9, 0x38, 0x23, 0xa1, 0xb1, 0xd5, 0x06, 0x9d, 0x9a, 0x77, 0x47, 0x41, 0x8f, 0xe1,
	0x4e, 0x42, 0x65, 0x30, 0xc2, 0x2c, 0x20, 0xc6, 0x83, 0xac, 0x7d, 0xef, 0x56, 0x40, 0x4f, 0x21,
	0x8a, 0xb1, 0x54, 0xfe, 0x44, 0x10, 0x49, 0x98, 0xf2, 0x07, 0x31, 0x0f, 0xc6, 0x46, 0x35, 0xb7,
	0x3d, 0xcc, 0xc8, 0x99, 0x06, 0xbd, 0x4c, 0x47, 0xc7, 0xb0, 0x99, 0x50, 0x29, 0x49, 0xa8, 0x7d,
	0xd2, 0x0f, 0x78, 0xca, 0x14, 0x11, 0xc6, 0x76, 0x9e, 0x70, 0xa0, 0x61, 0xee, 0x95, 0x7d, 0x8d,
	0xd0, 0x4b, 0x78, 0x38, 0x11, 0x3c, 0x4c, 0x83, 0xfb, 0x59, 0xb5, 0x3c, 0xab, 0x59, 0xe0, 0xb5,
	0xbc, 0x6e, 0xed, 0xf2, 0xca, 0xaa, 0xfc, 0xb9, 0xb2, 0x80, 0xfd, 0x75, 0x03, 0x56, 0xcf, 0xb0,
	0xc0, 0x89, 0x44, 0xaf, 0x60, 0x3d, 0xc1, 0x53, 0xff, 0x76, 0x20, 0x7d, 0x0f, 0x63, 0x39, 0xb3,
	0x1a, 0x7a, 0x63, 0x6b, 0xd8, 0xf6, 0xf6, 0x12, 0x3c, 0x7d, 0x5b, 0x4e, 0xfb, 0x0d, 0xc0, 0x56,
	0xc8, 0xbf, 0xb0, 0xec, 0x61, 0xf8, 0xe5, 0x6e, 0x8b, 0x67, 0xb4, 0x3a, 0xcf, 0xa3, 0x7b, 0xe7,
	0x79, 0xbd, 0x32, 0xf4, 0x9e, 0xad, 0xae, 0x73, 0xa4, 0xff, 0xeb, 0xff, 0xa5, 0xec, 0xcb, 0xec,
	0x52, 0x46, 0x61, 0x38, 0x5d, 0xf1, 0xa2, 0x10, 0x7a, 0x0e, 0x1b, 0x65, 0x93, 0x7e, 0xc0, 0xd9,
	0x90, 0x86, 0x24, 0x1b, 0x67, 0xab, 0xdc, 0xa3, 0x66, 0xfd, 0x12, 0xf5, 0x4e, 0xbe, 0xcf, 0x4d,
	0x70, 0x3d, 0x37, 0xc1, 0xcd, 0xdc, 0x04, 0xbf, 0xe7, 0x26, 0xb8, 0x58, 0x98, 0x95, 0x9b, 0x85,
	0x59, 0xf9, 0xb9, 0x30, 0x2b, 0x9f, 0x9e, 0x44, 0x54, 0x8d, 0xd2, 0x81, 0x13, 0xf0, 0xc4, 0x7d,
	0x43, 0x05, 0xee, 0x73, 0x41, 0x5c, 0x49, 0xc6, 0x98, 0xba, 0x53, 0xb7, 0xf8, 0xae, 0x5c, 0x75,
	0x3e, 0x21, 0x72, 0x50, 0xcd, 0x07, 0x7b, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x60, 0x10, 0x9e,
	0x62, 0x85, 0x03, 0x00, 0x00,
}

func (this *ValidatorSigningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfo)
	if !ok {
		that2, ok := that.(ValidatorSigningInfo)
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
	if this.Address != that1.Address {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if !this.InactiveUntil.Equal(that1.InactiveUntil) {
		return false
	}
	if this.Tombstoned != that1.Tombstoned {
		return false
	}
	if this.Mischance != that1.Mischance {
		return false
	}
	if this.LastPresentBlock != that1.LastPresentBlock {
		return false
	}
	if this.MissedBlocksCounter != that1.MissedBlocksCounter {
		return false
	}
	if this.ProducedBlocksCounter != that1.ProducedBlocksCounter {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MaxMischance != that1.MaxMischance {
		return false
	}
	if this.DowntimeInactiveDuration != that1.DowntimeInactiveDuration {
		return false
	}
	if this.MischanceConfidence != that1.MischanceConfidence {
		return false
	}
	return true
}
func (m *ValidatorSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProducedBlocksCounter != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.ProducedBlocksCounter))
		i--
		dAtA[i] = 0x40
	}
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x38
	}
	if m.LastPresentBlock != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.LastPresentBlock))
		i--
		dAtA[i] = 0x30
	}
	if m.Mischance != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.Mischance))
		i--
		dAtA[i] = 0x28
	}
	if m.Tombstoned {
		i--
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.InactiveUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.InactiveUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSlashing(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if m.StartHeight != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSlashing(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MischanceConfidence != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.MischanceConfidence))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DowntimeInactiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeInactiveDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSlashing(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if m.MaxMischance != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.MaxMischance))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintSlashing(dAtA []byte, offset int, v uint64) int {
	offset -= sovSlashing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSlashing(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovSlashing(uint64(m.StartHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.InactiveUntil)
	n += 1 + l + sovSlashing(uint64(l))
	if m.Tombstoned {
		n += 2
	}
	if m.Mischance != 0 {
		n += 1 + sovSlashing(uint64(m.Mischance))
	}
	if m.LastPresentBlock != 0 {
		n += 1 + sovSlashing(uint64(m.LastPresentBlock))
	}
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovSlashing(uint64(m.MissedBlocksCounter))
	}
	if m.ProducedBlocksCounter != 0 {
		n += 1 + sovSlashing(uint64(m.ProducedBlocksCounter))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxMischance != 0 {
		n += 1 + sovSlashing(uint64(m.MaxMischance))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeInactiveDuration)
	n += 1 + l + sovSlashing(uint64(l))
	if m.MischanceConfidence != 0 {
		n += 1 + sovSlashing(uint64(m.MischanceConfidence))
	}
	return n
}

func sovSlashing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSlashing(x uint64) (n int) {
	return sovSlashing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorSigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.InactiveUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mischance", wireType)
			}
			m.Mischance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mischance |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPresentBlock", wireType)
			}
			m.LastPresentBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastPresentBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducedBlocksCounter", wireType)
			}
			m.ProducedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProducedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlashing
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMischance", wireType)
			}
			m.MaxMischance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMischance |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DowntimeInactiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DowntimeInactiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MischanceConfidence", wireType)
			}
			m.MischanceConfidence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MischanceConfidence |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlashing
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
func skipSlashing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
				return 0, ErrInvalidLengthSlashing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSlashing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSlashing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSlashing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlashing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSlashing = fmt.Errorf("proto: unexpected end of group")
)
