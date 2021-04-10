// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.6
// source: collectionpb/collection.proto

package collectionpb

import (
	context "context"
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

//Check Token
//-------------------
//streaming data
type DataStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataReq int32 `protobuf:"varint,1,opt,name=dataReq,proto3" json:"dataReq,omitempty"`
}

func (x *DataStreamRequest) Reset() {
	*x = DataStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionpb_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStreamRequest) ProtoMessage() {}

func (x *DataStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collectionpb_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStreamRequest.ProtoReflect.Descriptor instead.
func (*DataStreamRequest) Descriptor() ([]byte, []int) {
	return file_collectionpb_collection_proto_rawDescGZIP(), []int{0}
}

func (x *DataStreamRequest) GetDataReq() int32 {
	if x != nil {
		return x.DataReq
	}
	return 0
}

type DataStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataRes string `protobuf:"bytes,1,opt,name=dataRes,proto3" json:"dataRes,omitempty"`
}

func (x *DataStreamResponse) Reset() {
	*x = DataStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionpb_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStreamResponse) ProtoMessage() {}

func (x *DataStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collectionpb_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStreamResponse.ProtoReflect.Descriptor instead.
func (*DataStreamResponse) Descriptor() ([]byte, []int) {
	return file_collectionpb_collection_proto_rawDescGZIP(), []int{1}
}

func (x *DataStreamResponse) GetDataRes() string {
	if x != nil {
		return x.DataRes
	}
	return ""
}

//streaming data
//Check Token and Role
type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckToken   string `protobuf:"bytes,1,opt,name=checkToken,proto3" json:"checkToken,omitempty"`
	CheckThingId string `protobuf:"bytes,2,opt,name=checkThingId,proto3" json:"checkThingId,omitempty"`
	CheckRole    string `protobuf:"bytes,3,opt,name=checkRole,proto3" json:"checkRole,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionpb_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collectionpb_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_collectionpb_collection_proto_rawDescGZIP(), []int{2}
}

func (x *CheckRequest) GetCheckToken() string {
	if x != nil {
		return x.CheckToken
	}
	return ""
}

func (x *CheckRequest) GetCheckThingId() string {
	if x != nil {
		return x.CheckThingId
	}
	return ""
}

func (x *CheckRequest) GetCheckRole() string {
	if x != nil {
		return x.CheckRole
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckRes string `protobuf:"bytes,1,opt,name=checkRes,proto3" json:"checkRes,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collectionpb_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collectionpb_collection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_collectionpb_collection_proto_rawDescGZIP(), []int{3}
}

func (x *CheckResponse) GetCheckRes() string {
	if x != nil {
		return x.CheckRes
	}
	return ""
}

var File_collectionpb_collection_proto protoreflect.FileDescriptor

var file_collectionpb_collection_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x0c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x2b, 0x0a, 0x0d,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x32, 0xae, 0x01, 0x0a, 0x11, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6e, 0x64,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_collectionpb_collection_proto_rawDescOnce sync.Once
	file_collectionpb_collection_proto_rawDescData = file_collectionpb_collection_proto_rawDesc
)

func file_collectionpb_collection_proto_rawDescGZIP() []byte {
	file_collectionpb_collection_proto_rawDescOnce.Do(func() {
		file_collectionpb_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_collectionpb_collection_proto_rawDescData)
	})
	return file_collectionpb_collection_proto_rawDescData
}

var file_collectionpb_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_collectionpb_collection_proto_goTypes = []interface{}{
	(*DataStreamRequest)(nil),  // 0: collection.DataStreamRequest
	(*DataStreamResponse)(nil), // 1: collection.DataStreamResponse
	(*CheckRequest)(nil),       // 2: collection.CheckRequest
	(*CheckResponse)(nil),      // 3: collection.CheckResponse
}
var file_collectionpb_collection_proto_depIdxs = []int32{
	2, // 0: collection.CollectionService.CheckTokenAndRole:input_type -> collection.CheckRequest
	0, // 1: collection.CollectionService.SendData:input_type -> collection.DataStreamRequest
	3, // 2: collection.CollectionService.CheckTokenAndRole:output_type -> collection.CheckResponse
	1, // 3: collection.CollectionService.SendData:output_type -> collection.DataStreamResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_collectionpb_collection_proto_init() }
func file_collectionpb_collection_proto_init() {
	if File_collectionpb_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collectionpb_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStreamRequest); i {
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
		file_collectionpb_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStreamResponse); i {
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
		file_collectionpb_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_collectionpb_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
			RawDescriptor: file_collectionpb_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collectionpb_collection_proto_goTypes,
		DependencyIndexes: file_collectionpb_collection_proto_depIdxs,
		MessageInfos:      file_collectionpb_collection_proto_msgTypes,
	}.Build()
	File_collectionpb_collection_proto = out.File
	file_collectionpb_collection_proto_rawDesc = nil
	file_collectionpb_collection_proto_goTypes = nil
	file_collectionpb_collection_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CollectionServiceClient is the client API for CollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectionServiceClient interface {
	//khai bao API
	//rpc CheckToken(TokenRequest) returns (TokenResponse){} //check token
	CheckTokenAndRole(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	SendData(ctx context.Context, opts ...grpc.CallOption) (CollectionService_SendDataClient, error)
}

type collectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionServiceClient(cc grpc.ClientConnInterface) CollectionServiceClient {
	return &collectionServiceClient{cc}
}

func (c *collectionServiceClient) CheckTokenAndRole(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/collection.CollectionService/CheckTokenAndRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) SendData(ctx context.Context, opts ...grpc.CallOption) (CollectionService_SendDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CollectionService_serviceDesc.Streams[0], "/collection.CollectionService/SendData", opts...)
	if err != nil {
		return nil, err
	}
	x := &collectionServiceSendDataClient{stream}
	return x, nil
}

type CollectionService_SendDataClient interface {
	Send(*DataStreamRequest) error
	CloseAndRecv() (*DataStreamResponse, error)
	grpc.ClientStream
}

type collectionServiceSendDataClient struct {
	grpc.ClientStream
}

func (x *collectionServiceSendDataClient) Send(m *DataStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *collectionServiceSendDataClient) CloseAndRecv() (*DataStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DataStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CollectionServiceServer is the server API for CollectionService service.
type CollectionServiceServer interface {
	//khai bao API
	//rpc CheckToken(TokenRequest) returns (TokenResponse){} //check token
	CheckTokenAndRole(context.Context, *CheckRequest) (*CheckResponse, error)
	SendData(CollectionService_SendDataServer) error
}

// UnimplementedCollectionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCollectionServiceServer struct {
}

func (*UnimplementedCollectionServiceServer) CheckTokenAndRole(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTokenAndRole not implemented")
}
func (*UnimplementedCollectionServiceServer) SendData(CollectionService_SendDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SendData not implemented")
}

func RegisterCollectionServiceServer(s *grpc.Server, srv CollectionServiceServer) {
	s.RegisterService(&_CollectionService_serviceDesc, srv)
}

func _CollectionService_CheckTokenAndRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).CheckTokenAndRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.CollectionService/CheckTokenAndRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).CheckTokenAndRole(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_SendData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CollectionServiceServer).SendData(&collectionServiceSendDataServer{stream})
}

type CollectionService_SendDataServer interface {
	SendAndClose(*DataStreamResponse) error
	Recv() (*DataStreamRequest, error)
	grpc.ServerStream
}

type collectionServiceSendDataServer struct {
	grpc.ServerStream
}

func (x *collectionServiceSendDataServer) SendAndClose(m *DataStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *collectionServiceSendDataServer) Recv() (*DataStreamRequest, error) {
	m := new(DataStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CollectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "collection.CollectionService",
	HandlerType: (*CollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckTokenAndRole",
			Handler:    _CollectionService_CheckTokenAndRole_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendData",
			Handler:       _CollectionService_SendData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "collectionpb/collection.proto",
}
