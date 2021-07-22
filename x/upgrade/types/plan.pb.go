// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/upgrade/plan.proto

package types

import (
	fmt "fmt"
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

type Plan struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Resources            []Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources"`
	Height               int64      `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	MinUpgradeTime       int64      `protobuf:"varint,4,opt,name=min_upgrade_time,json=minUpgradeTime,proto3" json:"min_upgrade_time,omitempty"`
	RollbackChecksum     string     `protobuf:"bytes,5,opt,name=rollback_checksum,json=rollbackChecksum,proto3" json:"rollback_checksum,omitempty"`
	MaxEnrolmentDuration int64      `protobuf:"varint,6,opt,name=max_enrolment_duration,json=maxEnrolmentDuration,proto3" json:"max_enrolment_duration,omitempty"`
	InstateUpgrade       bool       `protobuf:"varint,7,opt,name=instate_upgrade,json=instateUpgrade,proto3" json:"instate_upgrade,omitempty"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c9a10c2f596dea, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plan) GetResources() []Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Plan) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Plan) GetMinUpgradeTime() int64 {
	if m != nil {
		return m.MinUpgradeTime
	}
	return 0
}

func (m *Plan) GetRollbackChecksum() string {
	if m != nil {
		return m.RollbackChecksum
	}
	return ""
}

func (m *Plan) GetMaxEnrolmentDuration() int64 {
	if m != nil {
		return m.MaxEnrolmentDuration
	}
	return 0
}

func (m *Plan) GetInstateUpgrade() bool {
	if m != nil {
		return m.InstateUpgrade
	}
	return false
}

func init() {
	proto.RegisterType((*Plan)(nil), "tsuki.upgrade.Plan")
}

func init() { proto.RegisterFile("tsuki/upgrade/plan.proto", fileDescriptor_96c9a10c2f596dea) }

var fileDescriptor_96c9a10c2f596dea = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x93, 0x36, 0x5f, 0x3f, 0x6a, 0x50, 0x29, 0x56, 0x55, 0xac, 0x0e, 0x21, 0x62, 0x21,
	0x08, 0x29, 0x91, 0x80, 0x89, 0xb1, 0x85, 0x89, 0x05, 0x45, 0xb0, 0xb0, 0x44, 0x6e, 0x7a, 0x95,
	0x58, 0x89, 0xed, 0xc8, 0x71, 0xa4, 0xf2, 0x16, 0x3c, 0x04, 0x0f, 0xd3, 0xb1, 0x23, 0x13, 0x42,
	0xed, 0x8b, 0xa0, 0xa6, 0x2e, 0x7f, 0x26, 0x5f, 0x9f, 0xdf, 0xf1, 0xf5, 0xd1, 0x41, 0xc7, 0x39,
	0x53, 0x34, 0xac, 0xcb, 0x54, 0xd1, 0x19, 0x84, 0x65, 0x41, 0x45, 0x50, 0x2a, 0xa9, 0x25, 0x3e,
	0xd8, 0x80, 0xc0, 0x80, 0xd1, 0x20, 0x95, 0xa9, 0x6c, 0x40, 0xb8, 0x99, 0xb6, 0x9e, 0xd1, 0xe8,
	0xcf, 0x63, 0x73, 0x6e, 0xd9, 0xe9, 0x5b, 0x0b, 0x39, 0x0f, 0x05, 0x15, 0x18, 0x23, 0x47, 0x50,
	0x0e, 0xc4, 0xf6, 0x6c, 0xbf, 0x1b, 0x35, 0x33, 0xbe, 0x41, 0x5d, 0x05, 0x95, 0xac, 0x55, 0x02,
	0x15, 0x69, 0x79, 0x6d, 0x7f, 0xff, 0x72, 0x18, 0xfc, 0xfe, 0x30, 0x88, 0x0c, 0x1e, 0x3b, 0x8b,
	0x8f, 0x13, 0x2b, 0xfa, 0xb1, 0xe3, 0x21, 0xea, 0x64, 0xc0, 0xd2, 0x4c, 0x93, 0xb6, 0x67, 0xfb,
	0xed, 0xc8, 0xdc, 0xb0, 0x8f, 0xfa, 0x9c, 0x89, 0xd8, 0x2c, 0x88, 0x35, 0xe3, 0x40, 0x9c, 0xc6,
	0xd1, 0xe3, 0x4c, 0x3c, 0x6d, 0xe5, 0x47, 0xc6, 0x01, 0x5f, 0xa0, 0x23, 0x25, 0x8b, 0x62, 0x4a,
	0x93, 0x3c, 0x4e, 0x32, 0x48, 0xf2, 0xaa, 0xe6, 0xe4, 0x5f, 0x13, 0xaf, 0xbf, 0x03, 0x13, 0xa3,
	0xe3, 0x6b, 0x34, 0xe4, 0x74, 0x1e, 0x83, 0x50, 0xb2, 0xe0, 0x20, 0x74, 0x3c, 0xab, 0x15, 0xd5,
	0x4c, 0x0a, 0xd2, 0x69, 0x96, 0x0f, 0x38, 0x9d, 0xdf, 0xed, 0xe0, 0xad, 0x61, 0xf8, 0x0c, 0x1d,
	0x32, 0x51, 0x69, 0xaa, 0x61, 0x17, 0x88, 0xfc, 0xf7, 0x6c, 0x7f, 0x2f, 0xea, 0x19, 0xd9, 0xe4,
	0x19, 0x4f, 0x16, 0x2b, 0xd7, 0x5e, 0xae, 0x5c, 0xfb, 0x73, 0xe5, 0xda, 0xaf, 0x6b, 0xd7, 0x5a,
	0xae, 0x5d, 0xeb, 0x7d, 0xed, 0x5a, 0xcf, 0xe7, 0x29, 0xd3, 0x59, 0x3d, 0x0d, 0x12, 0xc9, 0xc3,
	0x7b, 0xa6, 0xe8, 0x44, 0x2a, 0x08, 0x2b, 0xc8, 0x29, 0x0b, 0xe7, 0xdf, 0x9d, 0xeb, 0x97, 0x12,
	0xaa, 0x69, 0xa7, 0xa9, 0xfc, 0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xdb, 0x11, 0x93, 0xcd,
	0x01, 0x00, 0x00,
}

func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InstateUpgrade {
		i--
		if m.InstateUpgrade {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.MaxEnrolmentDuration != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.MaxEnrolmentDuration))
		i--
		dAtA[i] = 0x30
	}
	if len(m.RollbackChecksum) > 0 {
		i -= len(m.RollbackChecksum)
		copy(dAtA[i:], m.RollbackChecksum)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.RollbackChecksum)))
		i--
		dAtA[i] = 0x2a
	}
	if m.MinUpgradeTime != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.MinUpgradeTime))
		i--
		dAtA[i] = 0x20
	}
	if m.Height != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Resources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlan(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = e.Size()
			n += 1 + l + sovPlan(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovPlan(uint64(m.Height))
	}
	if m.MinUpgradeTime != 0 {
		n += 1 + sovPlan(uint64(m.MinUpgradeTime))
	}
	l = len(m.RollbackChecksum)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.MaxEnrolmentDuration != 0 {
		n += 1 + sovPlan(uint64(m.MaxEnrolmentDuration))
	}
	if m.InstateUpgrade {
		n += 2
	}
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, Resource{})
			if err := m.Resources[len(m.Resources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUpgradeTime", wireType)
			}
			m.MinUpgradeTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUpgradeTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollbackChecksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollbackChecksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEnrolmentDuration", wireType)
			}
			m.MaxEnrolmentDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxEnrolmentDuration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstateUpgrade", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
			m.InstateUpgrade = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlan
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlan
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
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)
