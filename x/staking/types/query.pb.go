// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ValidatorByAddressRequest struct {
	ValAddr github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=val_addr,json=valAddr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"val_addr,omitempty" yaml:"val_addr"`
}

func (m *ValidatorByAddressRequest) Reset()         { *m = ValidatorByAddressRequest{} }
func (m *ValidatorByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*ValidatorByAddressRequest) ProtoMessage()    {}
func (*ValidatorByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}
func (m *ValidatorByAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorByAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorByAddressRequest.Merge(m, src)
}
func (m *ValidatorByAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorByAddressRequest proto.InternalMessageInfo

func (m *ValidatorByAddressRequest) GetValAddr() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValAddr
	}
	return nil
}

type ValidatorByMonikerRequest struct {
	Moniker string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
}

func (m *ValidatorByMonikerRequest) Reset()         { *m = ValidatorByMonikerRequest{} }
func (m *ValidatorByMonikerRequest) String() string { return proto.CompactTextString(m) }
func (*ValidatorByMonikerRequest) ProtoMessage()    {}
func (*ValidatorByMonikerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}
func (m *ValidatorByMonikerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorByMonikerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorByMonikerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorByMonikerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorByMonikerRequest.Merge(m, src)
}
func (m *ValidatorByMonikerRequest) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorByMonikerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorByMonikerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorByMonikerRequest proto.InternalMessageInfo

func (m *ValidatorByMonikerRequest) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

type ValidatorResponse struct {
	Validator Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator"`
}

func (m *ValidatorResponse) Reset()         { *m = ValidatorResponse{} }
func (m *ValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*ValidatorResponse) ProtoMessage()    {}
func (*ValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}
func (m *ValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorResponse.Merge(m, src)
}
func (m *ValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorResponse proto.InternalMessageInfo

func (m *ValidatorResponse) GetValidator() Validator {
	if m != nil {
		return m.Validator
	}
	return Validator{}
}

func init() {
	proto.RegisterType((*ValidatorByAddressRequest)(nil), "tsuki.staking.ValidatorByAddressRequest")
	proto.RegisterType((*ValidatorByMonikerRequest)(nil), "tsuki.staking.ValidatorByMonikerRequest")
	proto.RegisterType((*ValidatorResponse)(nil), "tsuki.staking.ValidatorResponse")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2c, 0x4d, 0x2d,
	0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0xce, 0x2c, 0x4a, 0xd4, 0x2b, 0x2e,
	0x49, 0xcc, 0xce, 0xcc, 0x4b, 0x97, 0xe2, 0x85, 0x32, 0x20, 0x92, 0x52, 0x22, 0xe9, 0xf9, 0xe9,
	0xf9, 0x60, 0xa6, 0x3e, 0x88, 0x05, 0x11, 0x55, 0xaa, 0xe1, 0x92, 0x0c, 0x4b, 0xcc, 0xc9, 0x4c,
	0x49, 0x2c, 0xc9, 0x2f, 0x72, 0xaa, 0x74, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x0e, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x8a, 0xe7, 0xe2, 0x28, 0x4b, 0xcc, 0x89, 0x4f, 0x4c, 0x49, 0x29,
	0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x71, 0x72, 0xf9, 0x74, 0x4f, 0x9e, 0xbf, 0x32, 0x31, 0x37,
	0xc7, 0x4a, 0x09, 0x26, 0xa3, 0xf4, 0xeb, 0x9e, 0xbc, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x31, 0x94, 0xd2, 0x2d, 0x4e, 0xc9,
	0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x0b, 0x4b, 0xcc, 0x81, 0x19, 0xcf, 0x5e, 0x06, 0x61,
	0x2b, 0x99, 0xa2, 0xd8, 0xee, 0x9b, 0x9f, 0x97, 0x99, 0x9d, 0x5a, 0x04, 0xb3, 0x5d, 0x82, 0x8b,
	0x3d, 0x17, 0x22, 0x02, 0xb6, 0x9c, 0x33, 0x08, 0xc6, 0x55, 0x0a, 0xe0, 0x12, 0x84, 0x6b, 0x0b,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2, 0xe6, 0xe2, 0x2c, 0x83, 0x09, 0x82, 0x35,
	0x70, 0x1b, 0x89, 0xeb, 0x21, 0x07, 0x88, 0x1e, 0xc2, 0x2a, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82,
	0x10, 0xea, 0x8d, 0x4e, 0x33, 0x72, 0xb1, 0x06, 0x82, 0x42, 0x52, 0x28, 0x81, 0x4b, 0x08, 0x33,
	0x40, 0x84, 0xd4, 0x71, 0x99, 0x84, 0x16, 0x64, 0x52, 0xf2, 0x38, 0x14, 0xc2, 0x9c, 0xa9, 0xc4,
	0x80, 0x66, 0x03, 0xd4, 0xd3, 0x78, 0x6c, 0x40, 0x0d, 0x16, 0x22, 0x6c, 0x70, 0x72, 0x3e, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8,
	0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x4d, 0xa4, 0x88, 0xf2, 0xce, 0x2c, 0x4a,
	0x74, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x4e, 0xcd, 0x4e, 0xcc, 0xd4, 0xaf, 0xd0, 0x87, 0x1a, 0x09,
	0x89, 0xaf, 0x24, 0x36, 0x70, 0x02, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x80, 0x0f,
	0xd2, 0x62, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Validators queries a validator by address.
	ValidatorByAddress(ctx context.Context, in *ValidatorByAddressRequest, opts ...grpc.CallOption) (*ValidatorResponse, error)
	// Validators queries a validator by moniker.
	ValidatorByMoniker(ctx context.Context, in *ValidatorByMonikerRequest, opts ...grpc.CallOption) (*ValidatorResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ValidatorByAddress(ctx context.Context, in *ValidatorByAddressRequest, opts ...grpc.CallOption) (*ValidatorResponse, error) {
	out := new(ValidatorResponse)
	err := c.cc.Invoke(ctx, "/tsuki.staking.Query/ValidatorByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorByMoniker(ctx context.Context, in *ValidatorByMonikerRequest, opts ...grpc.CallOption) (*ValidatorResponse, error) {
	out := new(ValidatorResponse)
	err := c.cc.Invoke(ctx, "/tsuki.staking.Query/ValidatorByMoniker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Validators queries a validator by address.
	ValidatorByAddress(context.Context, *ValidatorByAddressRequest) (*ValidatorResponse, error)
	// Validators queries a validator by moniker.
	ValidatorByMoniker(context.Context, *ValidatorByMonikerRequest) (*ValidatorResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ValidatorByAddress(ctx context.Context, req *ValidatorByAddressRequest) (*ValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorByAddress not implemented")
}
func (*UnimplementedQueryServer) ValidatorByMoniker(ctx context.Context, req *ValidatorByMonikerRequest) (*ValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorByMoniker not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ValidatorByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tsuki.staking.Query/ValidatorByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorByAddress(ctx, req.(*ValidatorByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorByMoniker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorByMonikerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorByMoniker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tsuki.staking.Query/ValidatorByMoniker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorByMoniker(ctx, req.(*ValidatorByMonikerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tsuki.staking.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidatorByAddress",
			Handler:    _Query_ValidatorByAddress_Handler,
		},
		{
			MethodName: "ValidatorByMoniker",
			Handler:    _Query_ValidatorByMoniker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query.proto",
}

func (m *ValidatorByAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorByAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorByAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValAddr) > 0 {
		i -= len(m.ValAddr)
		copy(dAtA[i:], m.ValAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorByMonikerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorByMonikerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorByMonikerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorByAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ValidatorByMonikerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Validator.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorByAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ValidatorByAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorByAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValAddr = append(m.ValAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ValAddr == nil {
				m.ValAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ValidatorByMonikerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ValidatorByMonikerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorByMonikerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
