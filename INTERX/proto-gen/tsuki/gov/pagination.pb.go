// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tsuki/gov/pagination.proto

package gov

import (
	fmt "fmt"
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

// PageRequest is to be embedded in gRPC request messages for efficient
// pagination. Ex:
//
//  message SomeRequest {
//          Foo some_parameter = 1;
//          PageRequest pagination = 2;
//  }
type PageRequest struct {
	// key is a value returned in PageResponse.next_key to begin
	// querying the next page most efficiently. Only one of offset or key
	// should be set.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// offset is a numeric offset that can be used when key is unavailable.
	// It is less efficient than using key. Only one of offset or key should
	// be set.
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// limit is the total number of results to be returned in the result page.
	// If left empty it will default to a value to be set by each app.
	Limit uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// count_total is set to true  to indicate that the result set should include
	// a count of the total number of items available for pagination in UIs. count_total
	// is only respected when offset is used. It is ignored when key is set.
	CountTotal           bool     `protobuf:"varint,4,opt,name=count_total,json=countTotal,proto3" json:"count_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageRequest) Reset()         { *m = PageRequest{} }
func (m *PageRequest) String() string { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()    {}
func (*PageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_454d1e027b6baed6, []int{0}
}

func (m *PageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageRequest.Unmarshal(m, b)
}
func (m *PageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageRequest.Marshal(b, m, deterministic)
}
func (m *PageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageRequest.Merge(m, src)
}
func (m *PageRequest) XXX_Size() int {
	return xxx_messageInfo_PageRequest.Size(m)
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

// PageResponse is to be embedded in gRPC response messages where the corresponding
// request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
type PageResponse struct {
	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	NextKey []byte `protobuf:"bytes,1,opt,name=next_key,json=nextKey,proto3" json:"next_key,omitempty"`
	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total                uint64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageResponse) Reset()         { *m = PageResponse{} }
func (m *PageResponse) String() string { return proto.CompactTextString(m) }
func (*PageResponse) ProtoMessage()    {}
func (*PageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_454d1e027b6baed6, []int{1}
}

func (m *PageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageResponse.Unmarshal(m, b)
}
func (m *PageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageResponse.Marshal(b, m, deterministic)
}
func (m *PageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageResponse.Merge(m, src)
}
func (m *PageResponse) XXX_Size() int {
	return xxx_messageInfo_PageResponse.Size(m)
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
	proto.RegisterType((*PageRequest)(nil), "tsuki.gov.PageRequest")
	proto.RegisterType((*PageResponse)(nil), "tsuki.gov.PageResponse")
}

func init() {
	proto.RegisterFile("tsuki/gov/pagination.proto", fileDescriptor_454d1e027b6baed6)
}

var fileDescriptor_454d1e027b6baed6 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0xa9, 0x9b, 0xb3, 0xbc, 0xed, 0x20, 0x41, 0xa4, 0x3b, 0x59, 0x76, 0xea, 0xc5, 0xe6,
	0x30, 0x3c, 0x0b, 0x8a, 0x07, 0x19, 0x88, 0x84, 0x1d, 0xc4, 0xcb, 0xc8, 0xc6, 0x5b, 0x0c, 0xdd,
	0xf2, 0x6a, 0xf3, 0x3a, 0xec, 0xbf, 0x97, 0x24, 0xca, 0x6e, 0xf9, 0xbe, 0xc0, 0xfb, 0x78, 0x0f,
	0xe6, 0x8d, 0xed, 0xb4, 0x34, 0x74, 0x92, 0xad, 0x36, 0xd6, 0x69, 0xb6, 0xe4, 0xea, 0xb6, 0x23,
	0x26, 0x91, 0x87, 0xaf, 0xda, 0xd0, 0x69, 0xe1, 0x60, 0xfa, 0xae, 0x0d, 0x2a, 0xfc, 0xee, 0xd1,
	0xb3, 0xb8, 0x86, 0x51, 0x83, 0x43, 0x91, 0x95, 0x59, 0x35, 0x53, 0xe1, 0x29, 0x6e, 0x61, 0x42,
	0xfb, 0xbd, 0x47, 0x2e, 0x2e, 0xca, 0xac, 0x1a, 0xab, 0x3f, 0x12, 0x37, 0x70, 0x79, 0xb0, 0x47,
	0xcb, 0xc5, 0x28, 0xea, 0x04, 0xe2, 0x0e, 0xa6, 0x3b, 0xea, 0x1d, 0x6f, 0x98, 0x58, 0x1f, 0x8a,
	0x71, 0x99, 0x55, 0xb9, 0x82, 0xa8, 0xd6, 0xc1, 0x2c, 0x1e, 0x61, 0x96, 0x7a, 0xbe, 0x25, 0xe7,
	0x51, 0xcc, 0x21, 0x77, 0xf8, 0xc3, 0x9b, 0x73, 0xf5, 0x2a, 0xf0, 0x0a, 0x87, 0x50, 0x48, 0x53,
	0x52, 0x38, 0xc1, 0xd3, 0xc3, 0xe7, 0xd2, 0x58, 0xfe, 0xea, 0xb7, 0xf5, 0x8e, 0x8e, 0x72, 0x65,
	0x3b, 0xfd, 0x4c, 0x1d, 0x4a, 0x8f, 0x8d, 0xb6, 0xf2, 0xf5, 0x6d, 0xfd, 0xa2, 0x3e, 0x64, 0x5c,
	0xf2, 0xde, 0xa0, 0x93, 0xff, 0x27, 0xd8, 0x4e, 0xa2, 0x5b, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x7c, 0x02, 0xaa, 0x15, 0x01, 0x00, 0x00,
}
