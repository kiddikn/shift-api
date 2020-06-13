// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: shift.proto

package shift

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

type GetShiftMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUser string `protobuf:"bytes,1,opt,name=target_user,json=targetUser,proto3" json:"target_user,omitempty"`
}

func (x *GetShiftMessage) Reset() {
	*x = GetShiftMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shift_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShiftMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShiftMessage) ProtoMessage() {}

func (x *GetShiftMessage) ProtoReflect() protoreflect.Message {
	mi := &file_shift_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShiftMessage.ProtoReflect.Descriptor instead.
func (*GetShiftMessage) Descriptor() ([]byte, []int) {
	return file_shift_proto_rawDescGZIP(), []int{0}
}

func (x *GetShiftMessage) GetTargetUser() string {
	if x != nil {
		return x.TargetUser
	}
	return ""
}

type ShiftResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Shift string `protobuf:"bytes,2,opt,name=shift,proto3" json:"shift,omitempty"`
}

func (x *ShiftResponse) Reset() {
	*x = ShiftResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shift_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShiftResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShiftResponse) ProtoMessage() {}

func (x *ShiftResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shift_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShiftResponse.ProtoReflect.Descriptor instead.
func (*ShiftResponse) Descriptor() ([]byte, []int) {
	return file_shift_proto_rawDescGZIP(), []int{1}
}

func (x *ShiftResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShiftResponse) GetShift() string {
	if x != nil {
		return x.Shift
	}
	return ""
}

var File_shift_proto protoreflect.FileDescriptor

var file_shift_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x68, 0x69, 0x66, 0x74, 0x22, 0x32, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x66, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x0d, 0x53, 0x68, 0x69, 0x66,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x68, 0x69, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68,
	0x69, 0x66, 0x74, 0x32, 0x43, 0x0a, 0x05, 0x53, 0x68, 0x69, 0x66, 0x74, 0x12, 0x3a, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x66, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x69, 0x66, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x66, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x14, 0x2e, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x53, 0x68, 0x69, 0x66, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x68, 0x69, 0x66,
	0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x68, 0x69, 0x66, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shift_proto_rawDescOnce sync.Once
	file_shift_proto_rawDescData = file_shift_proto_rawDesc
)

func file_shift_proto_rawDescGZIP() []byte {
	file_shift_proto_rawDescOnce.Do(func() {
		file_shift_proto_rawDescData = protoimpl.X.CompressGZIP(file_shift_proto_rawDescData)
	})
	return file_shift_proto_rawDescData
}

var file_shift_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shift_proto_goTypes = []interface{}{
	(*GetShiftMessage)(nil), // 0: shift.GetShiftMessage
	(*ShiftResponse)(nil),   // 1: shift.ShiftResponse
}
var file_shift_proto_depIdxs = []int32{
	0, // 0: shift.Shift.GetShift:input_type -> shift.GetShiftMessage
	1, // 1: shift.Shift.GetShift:output_type -> shift.ShiftResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shift_proto_init() }
func file_shift_proto_init() {
	if File_shift_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shift_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShiftMessage); i {
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
		file_shift_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShiftResponse); i {
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
			RawDescriptor: file_shift_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shift_proto_goTypes,
		DependencyIndexes: file_shift_proto_depIdxs,
		MessageInfos:      file_shift_proto_msgTypes,
	}.Build()
	File_shift_proto = out.File
	file_shift_proto_rawDesc = nil
	file_shift_proto_goTypes = nil
	file_shift_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShiftClient is the client API for Shift service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShiftClient interface {
	GetShift(ctx context.Context, in *GetShiftMessage, opts ...grpc.CallOption) (*ShiftResponse, error)
}

type shiftClient struct {
	cc grpc.ClientConnInterface
}

func NewShiftClient(cc grpc.ClientConnInterface) ShiftClient {
	return &shiftClient{cc}
}

func (c *shiftClient) GetShift(ctx context.Context, in *GetShiftMessage, opts ...grpc.CallOption) (*ShiftResponse, error) {
	out := new(ShiftResponse)
	err := c.cc.Invoke(ctx, "/shift.Shift/GetShift", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShiftServer is the server API for Shift service.
type ShiftServer interface {
	GetShift(context.Context, *GetShiftMessage) (*ShiftResponse, error)
}

// UnimplementedShiftServer can be embedded to have forward compatible implementations.
type UnimplementedShiftServer struct {
}

func (*UnimplementedShiftServer) GetShift(context.Context, *GetShiftMessage) (*ShiftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShift not implemented")
}

func RegisterShiftServer(s *grpc.Server, srv ShiftServer) {
	s.RegisterService(&_Shift_serviceDesc, srv)
}

func _Shift_GetShift_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShiftMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShiftServer).GetShift(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shift.Shift/GetShift",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShiftServer).GetShift(ctx, req.(*GetShiftMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shift_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shift.Shift",
	HandlerType: (*ShiftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShift",
			Handler:    _Shift_GetShift_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shift.proto",
}