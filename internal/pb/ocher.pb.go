// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.3
// source: ocher.proto

package pb

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

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*Status_Init_
	//	*Status_Error_
	//	*Status_Finish_
	Event isStatus_Event `protobuf_oneof:"event"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_ocher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_ocher_proto_rawDescGZIP(), []int{0}
}

func (m *Status) GetEvent() isStatus_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Status) GetInit() *Status_Init {
	if x, ok := x.GetEvent().(*Status_Init_); ok {
		return x.Init
	}
	return nil
}

func (x *Status) GetError() *Status_Error {
	if x, ok := x.GetEvent().(*Status_Error_); ok {
		return x.Error
	}
	return nil
}

func (x *Status) GetFinish() *Status_Finish {
	if x, ok := x.GetEvent().(*Status_Finish_); ok {
		return x.Finish
	}
	return nil
}

type isStatus_Event interface {
	isStatus_Event()
}

type Status_Init_ struct {
	Init *Status_Init `protobuf:"bytes,1,opt,name=init,proto3,oneof"`
}

type Status_Error_ struct {
	Error *Status_Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type Status_Finish_ struct {
	Finish *Status_Finish `protobuf:"bytes,3,opt,name=finish,proto3,oneof"`
}

func (*Status_Init_) isStatus_Event() {}

func (*Status_Error_) isStatus_Event() {}

func (*Status_Finish_) isStatus_Event() {}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Queue string `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Args  []byte `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_ocher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_ocher_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *Task) GetArgs() []byte {
	if x != nil {
		return x.Args
	}
	return nil
}

type Status_Init struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Queue string `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *Status_Init) Reset() {
	*x = Status_Init{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status_Init) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status_Init) ProtoMessage() {}

func (x *Status_Init) ProtoReflect() protoreflect.Message {
	mi := &file_ocher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status_Init.ProtoReflect.Descriptor instead.
func (*Status_Init) Descriptor() ([]byte, []int) {
	return file_ocher_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Status_Init) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Status_Init) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

type Status_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   []byte `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status_Error) Reset() {
	*x = Status_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status_Error) ProtoMessage() {}

func (x *Status_Error) ProtoReflect() protoreflect.Message {
	mi := &file_ocher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status_Error.ProtoReflect.Descriptor instead.
func (*Status_Error) Descriptor() ([]byte, []int) {
	return file_ocher_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Status_Error) GetError() []byte {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Status_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Status_Finish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Status_Finish) Reset() {
	*x = Status_Finish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status_Finish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status_Finish) ProtoMessage() {}

func (x *Status_Finish) ProtoReflect() protoreflect.Message {
	mi := &file_ocher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status_Finish.ProtoReflect.Descriptor instead.
func (*Status_Finish) Descriptor() ([]byte, []int) {
	return file_ocher_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Status_Finish) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_ocher_proto protoreflect.FileDescriptor

var file_ocher_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x63, 0x68, 0x65, 0x72, 0x22, 0xa1, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x28, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6f, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x48, 0x00, 0x52, 0x06,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x1a, 0x2c, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x1a, 0x37, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x0a,
	0x06, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x32, 0x3f, 0x0a, 0x05, 0x4f, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x0d, 0x2e, 0x6f, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x0b, 0x2e, 0x6f, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x6d,
	0x6b, 0x75, 0x7a, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6f,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocher_proto_rawDescOnce sync.Once
	file_ocher_proto_rawDescData = file_ocher_proto_rawDesc
)

func file_ocher_proto_rawDescGZIP() []byte {
	file_ocher_proto_rawDescOnce.Do(func() {
		file_ocher_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocher_proto_rawDescData)
	})
	return file_ocher_proto_rawDescData
}

var file_ocher_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ocher_proto_goTypes = []interface{}{
	(*Status)(nil),        // 0: ocher.Status
	(*Task)(nil),          // 1: ocher.Task
	(*Status_Init)(nil),   // 2: ocher.Status.Init
	(*Status_Error)(nil),  // 3: ocher.Status.Error
	(*Status_Finish)(nil), // 4: ocher.Status.Finish
}
var file_ocher_proto_depIdxs = []int32{
	2, // 0: ocher.Status.init:type_name -> ocher.Status.Init
	3, // 1: ocher.Status.error:type_name -> ocher.Status.Error
	4, // 2: ocher.Status.finish:type_name -> ocher.Status.Finish
	0, // 3: ocher.Ocher.QueueTransactional:input_type -> ocher.Status
	1, // 4: ocher.Ocher.QueueTransactional:output_type -> ocher.Task
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ocher_proto_init() }
func file_ocher_proto_init() {
	if File_ocher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_ocher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_ocher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status_Init); i {
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
		file_ocher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status_Error); i {
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
		file_ocher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status_Finish); i {
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
	file_ocher_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Status_Init_)(nil),
		(*Status_Error_)(nil),
		(*Status_Finish_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ocher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ocher_proto_goTypes,
		DependencyIndexes: file_ocher_proto_depIdxs,
		MessageInfos:      file_ocher_proto_msgTypes,
	}.Build()
	File_ocher_proto = out.File
	file_ocher_proto_rawDesc = nil
	file_ocher_proto_goTypes = nil
	file_ocher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OcherClient is the client API for Ocher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OcherClient interface {
	QueueTransactional(ctx context.Context, opts ...grpc.CallOption) (Ocher_QueueTransactionalClient, error)
}

type ocherClient struct {
	cc grpc.ClientConnInterface
}

func NewOcherClient(cc grpc.ClientConnInterface) OcherClient {
	return &ocherClient{cc}
}

func (c *ocherClient) QueueTransactional(ctx context.Context, opts ...grpc.CallOption) (Ocher_QueueTransactionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ocher_serviceDesc.Streams[0], "/ocher.Ocher/QueueTransactional", opts...)
	if err != nil {
		return nil, err
	}
	x := &ocherQueueTransactionalClient{stream}
	return x, nil
}

type Ocher_QueueTransactionalClient interface {
	Send(*Status) error
	Recv() (*Task, error)
	grpc.ClientStream
}

type ocherQueueTransactionalClient struct {
	grpc.ClientStream
}

func (x *ocherQueueTransactionalClient) Send(m *Status) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ocherQueueTransactionalClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OcherServer is the server API for Ocher service.
type OcherServer interface {
	QueueTransactional(Ocher_QueueTransactionalServer) error
}

// UnimplementedOcherServer can be embedded to have forward compatible implementations.
type UnimplementedOcherServer struct {
}

func (*UnimplementedOcherServer) QueueTransactional(Ocher_QueueTransactionalServer) error {
	return status.Errorf(codes.Unimplemented, "method QueueTransactional not implemented")
}

func RegisterOcherServer(s *grpc.Server, srv OcherServer) {
	s.RegisterService(&_Ocher_serviceDesc, srv)
}

func _Ocher_QueueTransactional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OcherServer).QueueTransactional(&ocherQueueTransactionalServer{stream})
}

type Ocher_QueueTransactionalServer interface {
	Send(*Task) error
	Recv() (*Status, error)
	grpc.ServerStream
}

type ocherQueueTransactionalServer struct {
	grpc.ServerStream
}

func (x *ocherQueueTransactionalServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ocherQueueTransactionalServer) Recv() (*Status, error) {
	m := new(Status)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Ocher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocher.Ocher",
	HandlerType: (*OcherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueueTransactional",
			Handler:       _Ocher_QueueTransactional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ocher.proto",
}