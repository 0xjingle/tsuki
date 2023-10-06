// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/staking/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type ProposalUnjailValidator struct {
	Proposer  string `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	ValAddr   string `protobuf:"bytes,2,opt,name=valAddr,proto3" json:"valAddr,omitempty"`
	Reference string `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
}

func (m *ProposalUnjailValidator) Reset()         { *m = ProposalUnjailValidator{} }
func (m *ProposalUnjailValidator) String() string { return proto.CompactTextString(m) }
func (*ProposalUnjailValidator) ProtoMessage()    {}
func (*ProposalUnjailValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_557d21e45da18a9e, []int{0}
}
func (m *ProposalUnjailValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalUnjailValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalUnjailValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalUnjailValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalUnjailValidator.Merge(m, src)
}
func (m *ProposalUnjailValidator) XXX_Size() int {
	return m.Size()
}
func (m *ProposalUnjailValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalUnjailValidator.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalUnjailValidator proto.InternalMessageInfo

func (m *ProposalUnjailValidator) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *ProposalUnjailValidator) GetValAddr() string {
	if m != nil {
		return m.ValAddr
	}
	return ""
}

func (m *ProposalUnjailValidator) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func init() {
	proto.RegisterType((*ProposalUnjailValidator)(nil), "tsuki.staking.ProposalUnjailValidator")
}

func init() { proto.RegisterFile("tsuki/staking/proposal.proto", fileDescriptor_557d21e45da18a9e) }

var fileDescriptor_557d21e45da18a9e = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xce, 0x2c, 0x4a,
	0xd4, 0x2f, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e,
	0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x49, 0xea, 0x41, 0x25, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x12, 0xfa, 0x20, 0x16, 0x44, 0x8d, 0x94, 0x64, 0x72, 0x7e, 0x71,
	0x6e, 0x7e, 0x71, 0x3c, 0x44, 0x02, 0xc2, 0x81, 0x48, 0x29, 0xd5, 0x71, 0x89, 0x07, 0x40, 0x0d,
	0x0c, 0xcd, 0xcb, 0x4a, 0xcc, 0xcc, 0x09, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x12,
	0x92, 0xe2, 0xe2, 0x80, 0xd8, 0x95, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7,
	0x0b, 0x49, 0x70, 0xb1, 0x97, 0x25, 0xe6, 0x38, 0xa6, 0xa4, 0x14, 0x49, 0x30, 0x81, 0xa5, 0x60,
	0x5c, 0x21, 0x19, 0x2e, 0xce, 0xa2, 0xd4, 0xb4, 0xd4, 0xa2, 0xd4, 0xbc, 0xe4, 0x54, 0x09, 0x66,
	0xb0, 0x1c, 0x42, 0xc0, 0x8a, 0xff, 0xc5, 0x02, 0x79, 0xc6, 0x53, 0x5b, 0x74, 0xd9, 0x9d, 0xf3,
	0xf3, 0x4a, 0x52, 0xf3, 0x4a, 0x9c, 0x9c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0x4a, 0x33, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0x3b, 0xb3,
	0x28, 0xd1, 0x39, 0xbf, 0x28, 0x55, 0xbf, 0x38, 0x35, 0x3b, 0x31, 0x53, 0xbf, 0x02, 0x1e, 0x18,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xbf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x01, 0x34, 0x43, 0x88, 0x29, 0x01, 0x00, 0x00,
}

func (this *ProposalUnjailValidator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProposalUnjailValidator)
	if !ok {
		that2, ok := that.(ProposalUnjailValidator)
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
	if this.Proposer != that1.Proposer {
		return false
	}
	if this.ValAddr != that1.ValAddr {
		return false
	}
	if this.Reference != that1.Reference {
		return false
	}
	return true
}
func (m *ProposalUnjailValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalUnjailValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalUnjailValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reference) > 0 {
		i -= len(m.Reference)
		copy(dAtA[i:], m.Reference)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Reference)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValAddr) > 0 {
		i -= len(m.ValAddr)
		copy(dAtA[i:], m.ValAddr)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ValAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Proposer)))
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
func (m *ProposalUnjailValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.ValAddr)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Reference)
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
func (m *ProposalUnjailValidator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProposalUnjailValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalUnjailValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
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
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddr", wireType)
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
			m.ValAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reference", wireType)
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
			m.Reference = string(dAtA[iNdEx:postIndex])
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
