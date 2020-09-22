// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cosmos/tx.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PostTransactionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Data      string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Codespace string `protobuf:"bytes,4,opt,name=codespace,proto3" json:"codespace,omitempty"`
	Hash      string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *PostTransactionResult) Reset() {
	*x = PostTransactionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTransactionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTransactionResult) ProtoMessage() {}

func (x *PostTransactionResult) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTransactionResult.ProtoReflect.Descriptor instead.
func (*PostTransactionResult) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_proto_rawDescGZIP(), []int{0}
}

func (x *PostTransactionResult) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PostTransactionResult) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *PostTransactionResult) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *PostTransactionResult) GetCodespace() string {
	if x != nil {
		return x.Codespace
	}
	return ""
}

func (x *PostTransactionResult) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// PostTransactionRequest is the request type for the tx/PostTransaction RPC method.
type PostTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// transaction hash.
	Tx []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *PostTransactionRequest) Reset() {
	*x = PostTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTransactionRequest) ProtoMessage() {}

func (x *PostTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTransactionRequest.ProtoReflect.Descriptor instead.
func (*PostTransactionRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_proto_rawDescGZIP(), []int{1}
}

func (x *PostTransactionRequest) GetTx() []byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

// PostTransactionResponse is the response type for the tx/PostTransaction RPC method.
type PostTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId   string                 `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Block     uint64                 `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	BlockTime string                 `protobuf:"bytes,3,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	Timestamp uint64                 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Response  *PostTransactionResult `protobuf:"bytes,5,opt,name=response,proto3" json:"response,omitempty"`
	Error     *Error                 `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Signature string                 `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	Hash      string                 `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *PostTransactionResponse) Reset() {
	*x = PostTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTransactionResponse) ProtoMessage() {}

func (x *PostTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTransactionResponse.ProtoReflect.Descriptor instead.
func (*PostTransactionResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_proto_rawDescGZIP(), []int{2}
}

func (x *PostTransactionResponse) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *PostTransactionResponse) GetBlock() uint64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *PostTransactionResponse) GetBlockTime() string {
	if x != nil {
		return x.BlockTime
	}
	return ""
}

func (x *PostTransactionResponse) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PostTransactionResponse) GetResponse() *PostTransactionResult {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *PostTransactionResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PostTransactionResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PostTransactionResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_cosmos_tx_proto protoreflect.FileDescriptor

var file_cosmos_tx_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x78, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01,
	0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x22, 0x28, 0x0a, 0x16, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x78, 0x22, 0x9c, 0x02,
	0x0a, 0x17, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x78, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x32, 0xac, 0x01, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x9c, 0x01, 0x0a,
	0x0f, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22,
	0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x92,
	0x41, 0x29, 0x0a, 0x02, 0x74, 0x78, 0x12, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x20, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x20, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69, 0x72, 0x61, 0x43, 0x6f,
	0x72, 0x65, 0x2f, 0x73, 0x65, 0x6b, 0x61, 0x69, 0x2f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x58, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x92, 0x41, 0x44, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a,
	0x01, 0x01, 0x72, 0x38, 0x0a, 0x0c, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69, 0x72, 0x61, 0x43, 0x6f, 0x72, 0x65, 0x2f,
	0x73, 0x65, 0x6b, 0x61, 0x69, 0x2f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x58, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_tx_proto_rawDescOnce sync.Once
	file_cosmos_tx_proto_rawDescData = file_cosmos_tx_proto_rawDesc
)

func file_cosmos_tx_proto_rawDescGZIP() []byte {
	file_cosmos_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_tx_proto_rawDescData)
	})
	return file_cosmos_tx_proto_rawDescData
}

var file_cosmos_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cosmos_tx_proto_goTypes = []interface{}{
	(*PostTransactionResult)(nil),   // 0: cosmos.tx.PostTransactionResult
	(*PostTransactionRequest)(nil),  // 1: cosmos.tx.PostTransactionRequest
	(*PostTransactionResponse)(nil), // 2: cosmos.tx.PostTransactionResponse
	(*Error)(nil),                   // 3: interx.Error
}
var file_cosmos_tx_proto_depIdxs = []int32{
	0, // 0: cosmos.tx.PostTransactionResponse.response:type_name -> cosmos.tx.PostTransactionResult
	3, // 1: cosmos.tx.PostTransactionResponse.error:type_name -> interx.Error
	1, // 2: cosmos.tx.Transaction.PostTransaction:input_type -> cosmos.tx.PostTransactionRequest
	2, // 3: cosmos.tx.Transaction.PostTransaction:output_type -> cosmos.tx.PostTransactionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_tx_proto_init() }
func file_cosmos_tx_proto_init() {
	if File_cosmos_tx_proto != nil {
		return
	}
	file_interx_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTransactionResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_tx_proto_msgTypes,
	}.Build()
	File_cosmos_tx_proto = out.File
	file_cosmos_tx_proto_rawDesc = nil
	file_cosmos_tx_proto_goTypes = nil
	file_cosmos_tx_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	PostTransaction(ctx context.Context, in *PostTransactionRequest, opts ...grpc.CallOption) (*PostTransactionResponse, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) PostTransaction(ctx context.Context, in *PostTransactionRequest, opts ...grpc.CallOption) (*PostTransactionResponse, error) {
	out := new(PostTransactionResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.Transaction/PostTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	PostTransaction(context.Context, *PostTransactionRequest) (*PostTransactionResponse, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) PostTransaction(context.Context, *PostTransactionRequest) (*PostTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostTransaction not implemented")
}

func RegisterTransactionServer(s *grpc.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_PostTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).PostTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.Transaction/PostTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).PostTransaction(ctx, req.(*PostTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.tx.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostTransaction",
			Handler:    _Transaction_PostTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/tx.proto",
}
