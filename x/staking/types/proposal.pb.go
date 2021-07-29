// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/staking/proposal.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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
	Proposer  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty" yaml:"proposer"`
	Hash      string                                        `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Reference string                                        `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
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

func (m *ProposalUnjailValidator) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *ProposalUnjailValidator) GetHash() string {
	if m != nil {
		return m.Hash
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
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xce, 0x2c, 0x4a,
	0xd4, 0x2f, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e,
	0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x49, 0xea, 0x41, 0x25, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x12, 0xfa, 0x20, 0x16, 0x44, 0x8d, 0x94, 0x64, 0x72, 0x7e, 0x71,
	0x6e, 0x7e, 0x71, 0x3c, 0x44, 0x02, 0xc2, 0x81, 0x48, 0x29, 0xed, 0x63, 0xe4, 0x12, 0x0f, 0x80,
	0x9a, 0x18, 0x9a, 0x97, 0x95, 0x98, 0x99, 0x13, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0x5f,
	0x24, 0x94, 0xc0, 0xc5, 0x01, 0xb1, 0x2c, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0xc7, 0xc9,
	0xe5, 0xd3, 0x3d, 0x79, 0xfe, 0xca, 0xc4, 0xdc, 0x1c, 0x2b, 0x25, 0x98, 0x8c, 0xd2, 0xaf, 0x7b,
	0xf2, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0x50, 0xd3, 0xa1, 0x94,
	0x6e, 0x71, 0x4a, 0xb6, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0xb1, 0x9e, 0x63, 0x72, 0xb2, 0x63, 0x4a,
	0x4a, 0x51, 0x6a, 0x71, 0x71, 0x10, 0xdc, 0x54, 0x21, 0x21, 0x2e, 0x96, 0x8c, 0xc4, 0xe2, 0x0c,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x86, 0x8b, 0xb3, 0x28, 0x35, 0x2d,
	0xb5, 0x28, 0x35, 0x2f, 0x39, 0x55, 0x82, 0x19, 0x2c, 0x81, 0x10, 0xb0, 0xe2, 0x7f, 0xb1, 0x40,
	0x9e, 0xf1, 0xd2, 0x16, 0x5d, 0x76, 0xe7, 0xfc, 0xbc, 0x92, 0xd4, 0xbc, 0x12, 0x27, 0xe7, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x44, 0x72, 0x95, 0x77, 0x66, 0x51,
	0xa2, 0x73, 0x7e, 0x51, 0xaa, 0x7e, 0x71, 0x6a, 0x76, 0x62, 0xa6, 0x7e, 0x05, 0x3c, 0x34, 0xc1,
	0x8e, 0x4b, 0x62, 0x03, 0x07, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x27, 0x9b, 0x0a,
	0x6a, 0x01, 0x00, 0x00,
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
	if !bytes.Equal(this.Proposer, that1.Proposer) {
		return false
	}
	if this.Hash != that1.Hash {
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
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Hash)))
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
	l = len(m.Hash)
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
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
			m.Hash = string(dAtA[iNdEx:postIndex])
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
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
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
