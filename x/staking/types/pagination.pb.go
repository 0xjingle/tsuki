// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pagination.proto

package types

import (
	fmt "fmt"
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

type PageRequest struct {
	Key        []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Offset     uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit      uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	CountTotal bool   `protobuf:"varint,4,opt,name=count_total,json=countTotal,proto3" json:"count_total,omitempty"`
}

func (m *PageRequest) Reset()         { *m = PageRequest{} }
func (m *PageRequest) String() string { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()    {}
func (*PageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_567bfb3a87c868dd, []int{0}
}
func (m *PageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageRequest.Merge(m, src)
}
func (m *PageRequest) XXX_Size() int {
	return m.Size()
}
func (m *PageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PageRequest proto.InternalMessageInfo

func (m *PageRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PageRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *PageRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PageRequest) GetCountTotal() bool {
	if m != nil {
		return m.CountTotal
	}
	return false
}

type PageResponse struct {
	NextKey []byte `protobuf:"bytes,1,opt,name=next_key,json=nextKey,proto3" json:"next_key,omitempty"`
	Total   uint64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *PageResponse) Reset()         { *m = PageResponse{} }
func (m *PageResponse) String() string { return proto.CompactTextString(m) }
func (*PageResponse) ProtoMessage()    {}
func (*PageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_567bfb3a87c868dd, []int{1}
}
func (m *PageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageResponse.Merge(m, src)
}
func (m *PageResponse) XXX_Size() int {
	return m.Size()
}
func (m *PageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PageResponse proto.InternalMessageInfo

func (m *PageResponse) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

func (m *PageResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*PageRequest)(nil), "tsuki.staking.PageRequest")
	proto.RegisterType((*PageResponse)(nil), "tsuki.staking.PageResponse")
}

func init() { proto.RegisterFile("pagination.proto", fileDescriptor_567bfb3a87c868dd) }

var fileDescriptor_567bfb3a87c868dd = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x63, 0x5a, 0x4a, 0xe5, 0x66, 0xa8, 0x2c, 0x84, 0xc2, 0x62, 0xa2, 0x4e, 0x61, 0x89,
	0x07, 0x3e, 0x00, 0x89, 0x8e, 0x5d, 0x50, 0xc4, 0xc4, 0x52, 0xb9, 0xd5, 0x35, 0x58, 0x69, 0x7d,
	0x21, 0xbe, 0x48, 0xcd, 0x5f, 0xf0, 0x59, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x23, 0x28, 0x71, 0x24,
	0xb6, 0x7b, 0x6f, 0xb8, 0xa7, 0x3b, 0xbe, 0x2c, 0x75, 0x6e, 0xac, 0x26, 0x83, 0x36, 0x2d, 0x2b,
	0x24, 0x14, 0x61, 0x61, 0x2a, 0x9d, 0x3a, 0xd2, 0x85, 0xb1, 0xf9, 0xca, 0xf2, 0xc5, 0xab, 0xce,
	0x21, 0x83, 0xcf, 0x1a, 0x1c, 0x89, 0x25, 0x9f, 0x14, 0xd0, 0x44, 0x2c, 0x66, 0x49, 0x98, 0xf5,
	0xa3, 0xb8, 0xe3, 0x33, 0x3c, 0x1c, 0x1c, 0x50, 0x74, 0x15, 0xb3, 0x64, 0x9a, 0x8d, 0x24, 0x6e,
	0xf9, 0xf5, 0xd1, 0x9c, 0x0c, 0x45, 0x93, 0x41, 0x7b, 0x10, 0x0f, 0x7c, 0xb1, 0xc7, 0xda, 0xd2,
	0x96, 0x90, 0xf4, 0x31, 0x9a, 0xc6, 0x2c, 0x99, 0x67, 0x7c, 0x50, 0x6f, 0xbd, 0x59, 0x3d, 0xf3,
	0xd0, 0xf7, 0x5c, 0x89, 0xd6, 0x81, 0xb8, 0xe7, 0x73, 0x0b, 0x67, 0xda, 0xfe, 0x57, 0x6f, 0x7a,
	0xde, 0x40, 0xd3, 0x17, 0xfc, 0x16, 0x1f, 0xf6, 0xf0, 0xb2, 0xfe, 0x6e, 0x25, 0xbb, 0xb4, 0x92,
	0xfd, 0xb6, 0x92, 0x7d, 0x75, 0x32, 0xb8, 0x74, 0x32, 0xf8, 0xe9, 0x64, 0xf0, 0xfe, 0x98, 0x1b,
	0xfa, 0xa8, 0x77, 0xe9, 0x1e, 0x4f, 0x6a, 0x63, 0x2a, 0xbd, 0xc6, 0x0a, 0x94, 0x83, 0x42, 0x1b,
	0x75, 0x56, 0xe3, 0xbd, 0x8a, 0x9a, 0x12, 0xdc, 0x6e, 0x36, 0xbc, 0xe2, 0xe9, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x98, 0xb2, 0xa0, 0x64, 0x1e, 0x01, 0x00, 0x00,
}

func (m *PageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CountTotal {
		i--
		if m.CountTotal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Limit != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x18
	}
	if m.Offset != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintPagination(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Total != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if len(m.NextKey) > 0 {
		i -= len(m.NextKey)
		copy(dAtA[i:], m.NextKey)
		i = encodeVarintPagination(dAtA, i, uint64(len(m.NextKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPagination(dAtA []byte, offset int, v uint64) int {
	offset -= sovPagination(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovPagination(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovPagination(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovPagination(uint64(m.Limit))
	}
	if m.CountTotal {
		n += 2
	}
	return n
}

func (m *PageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NextKey)
	if l > 0 {
		n += 1 + l + sovPagination(uint64(l))
	}
	if m.Total != 0 {
		n += 1 + sovPagination(uint64(m.Total))
	}
	return n
}

func sovPagination(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPagination(x uint64) (n int) {
	return sovPagination(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: PageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
				return ErrInvalidLengthPagination
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPagination
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountTotal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
			m.CountTotal = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPagination
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPagination
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
func (m *PageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: PageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
				return ErrInvalidLengthPagination
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPagination
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextKey = append(m.NextKey[:0], dAtA[iNdEx:postIndex]...)
			if m.NextKey == nil {
				m.NextKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPagination
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPagination
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
func skipPagination(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
				return 0, ErrInvalidLengthPagination
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPagination
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPagination
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPagination        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPagination          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPagination = fmt.Errorf("proto: unexpected end of group")
)
