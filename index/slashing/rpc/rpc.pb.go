// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: index/slashing/rpc/rpc.proto

package rpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tipsetkey string              `protobuf:"bytes,1,opt,name=tipsetkey,proto3" json:"tipsetkey,omitempty"`
	Miners    map[string]*Slashes `protobuf:"bytes,2,rep,name=miners,proto3" json:"miners,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_slashing_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_index_slashing_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_index_slashing_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *Index) GetTipsetkey() string {
	if x != nil {
		return x.Tipsetkey
	}
	return ""
}

func (x *Index) GetMiners() map[string]*Slashes {
	if x != nil {
		return x.Miners
	}
	return nil
}

type Slashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epochs []uint64 `protobuf:"varint,1,rep,packed,name=epochs,proto3" json:"epochs,omitempty"`
}

func (x *Slashes) Reset() {
	*x = Slashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_slashing_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slashes) ProtoMessage() {}

func (x *Slashes) ProtoReflect() protoreflect.Message {
	mi := &file_index_slashing_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slashes.ProtoReflect.Descriptor instead.
func (*Slashes) Descriptor() ([]byte, []int) {
	return file_index_slashing_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *Slashes) GetEpochs() []uint64 {
	if x != nil {
		return x.Epochs
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_slashing_rpc_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_index_slashing_rpc_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_index_slashing_rpc_rpc_proto_rawDescGZIP(), []int{2}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index *Index `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_slashing_rpc_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_index_slashing_rpc_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_index_slashing_rpc_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetIndex() *Index {
	if x != nil {
		return x.Index
	}
	return nil
}

var File_index_slashing_rpc_rpc_proto protoreflect.FileDescriptor

var file_index_slashing_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x72,
	0x70, 0x63, 0x22, 0xbc, 0x01, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x70, 0x73, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x70, 0x73, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x06, 0x6d, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x1a, 0x56, 0x0a, 0x0b, 0x4d, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x21, 0x0a, 0x07, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x32, 0x56, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x48, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_index_slashing_rpc_rpc_proto_rawDescOnce sync.Once
	file_index_slashing_rpc_rpc_proto_rawDescData = file_index_slashing_rpc_rpc_proto_rawDesc
)

func file_index_slashing_rpc_rpc_proto_rawDescGZIP() []byte {
	file_index_slashing_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_index_slashing_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_index_slashing_rpc_rpc_proto_rawDescData)
	})
	return file_index_slashing_rpc_rpc_proto_rawDescData
}

var file_index_slashing_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_index_slashing_rpc_rpc_proto_goTypes = []interface{}{
	(*Index)(nil),       // 0: index.slashing.rpc.Index
	(*Slashes)(nil),     // 1: index.slashing.rpc.Slashes
	(*GetRequest)(nil),  // 2: index.slashing.rpc.GetRequest
	(*GetResponse)(nil), // 3: index.slashing.rpc.GetResponse
	nil,                 // 4: index.slashing.rpc.Index.MinersEntry
}
var file_index_slashing_rpc_rpc_proto_depIdxs = []int32{
	4, // 0: index.slashing.rpc.Index.miners:type_name -> index.slashing.rpc.Index.MinersEntry
	0, // 1: index.slashing.rpc.GetResponse.index:type_name -> index.slashing.rpc.Index
	1, // 2: index.slashing.rpc.Index.MinersEntry.value:type_name -> index.slashing.rpc.Slashes
	2, // 3: index.slashing.rpc.RPCService.Get:input_type -> index.slashing.rpc.GetRequest
	3, // 4: index.slashing.rpc.RPCService.Get:output_type -> index.slashing.rpc.GetResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_index_slashing_rpc_rpc_proto_init() }
func file_index_slashing_rpc_rpc_proto_init() {
	if File_index_slashing_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_index_slashing_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_index_slashing_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slashes); i {
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
		file_index_slashing_rpc_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_index_slashing_rpc_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
			RawDescriptor: file_index_slashing_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_index_slashing_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_index_slashing_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_index_slashing_rpc_rpc_proto_msgTypes,
	}.Build()
	File_index_slashing_rpc_rpc_proto = out.File
	file_index_slashing_rpc_rpc_proto_rawDesc = nil
	file_index_slashing_rpc_rpc_proto_goTypes = nil
	file_index_slashing_rpc_rpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type rPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCServiceClient(cc grpc.ClientConnInterface) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/index.slashing.rpc.RPCService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServiceServer is the server API for RPCService service.
type RPCServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

// UnimplementedRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRPCServiceServer struct {
}

func (*UnimplementedRPCServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterRPCServiceServer(s *grpc.Server, srv RPCServiceServer) {
	s.RegisterService(&_RPCService_serviceDesc, srv)
}

func _RPCService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.slashing.rpc.RPCService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "index.slashing.rpc.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RPCService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index/slashing/rpc/rpc.proto",
}
