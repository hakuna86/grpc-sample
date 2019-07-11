// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatsample/chatsample.proto

package chatsample

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	common "github.com/hakuna86/grpc-sample/protos/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Process int32

const (
	Process_MAKE_BLOCK Process = 0
)

var Process_name = map[int32]string{
	0: "MAKE_BLOCK",
}

var Process_value = map[string]int32{
	"MAKE_BLOCK": 0,
}

func (x Process) String() string {
	return proto.EnumName(Process_name, int32(x))
}

func (Process) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c37368e4cf609c74, []int{0}
}

type Code struct {
	Code                 Process  `protobuf:"varint,1,opt,name=code,proto3,enum=chatsample.Process" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Code) Reset()         { *m = Code{} }
func (m *Code) String() string { return proto.CompactTextString(m) }
func (*Code) ProtoMessage()    {}
func (*Code) Descriptor() ([]byte, []int) {
	return fileDescriptor_c37368e4cf609c74, []int{0}
}

func (m *Code) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Code.Unmarshal(m, b)
}
func (m *Code) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Code.Marshal(b, m, deterministic)
}
func (m *Code) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code.Merge(m, src)
}
func (m *Code) XXX_Size() int {
	return xxx_messageInfo_Code.Size(m)
}
func (m *Code) XXX_DiscardUnknown() {
	xxx_messageInfo_Code.DiscardUnknown(m)
}

var xxx_messageInfo_Code proto.InternalMessageInfo

func (m *Code) GetCode() Process {
	if m != nil {
		return m.Code
	}
	return Process_MAKE_BLOCK
}

func init() {
	proto.RegisterEnum("chatsample.Process", Process_name, Process_value)
	proto.RegisterType((*Code)(nil), "chatsample.Code")
}

func init() { proto.RegisterFile("chatsample/chatsample.proto", fileDescriptor_c37368e4cf609c74) }

var fileDescriptor_c37368e4cf609c74 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x5b, 0x28, 0x0a, 0xef, 0x30, 0x46, 0x0a, 0xa2, 0xdd, 0x45, 0x7a, 0x71, 0x08, 0x26,
	0x73, 0x93, 0xe1, 0xd5, 0x95, 0x9d, 0xe6, 0x70, 0xcc, 0x9b, 0x17, 0x49, 0xb3, 0x67, 0x32, 0x5c,
	0xf7, 0x42, 0x93, 0x0a, 0xfa, 0xd7, 0x4b, 0x7f, 0xc8, 0x7a, 0xf0, 0x94, 0xe4, 0x93, 0xc7, 0xfb,
	0xfe, 0x80, 0x91, 0x32, 0xd2, 0x3b, 0x59, 0xd8, 0x03, 0x8a, 0xd3, 0x95, 0xdb, 0x92, 0x3c, 0x31,
	0x38, 0x91, 0x64, 0xa4, 0x89, 0xf4, 0x01, 0x45, 0xf3, 0x93, 0x57, 0x1f, 0x02, 0x0b, 0xeb, 0xbf,
	0xdb, 0xc1, 0x24, 0x56, 0x54, 0x14, 0x74, 0x14, 0xed, 0xd1, 0xc2, 0x54, 0x40, 0x94, 0xd1, 0x0e,
	0xd9, 0x0d, 0x44, 0x8a, 0x76, 0x78, 0x19, 0x5e, 0x87, 0xe3, 0xc1, 0x34, 0xe6, 0x3d, 0x99, 0x4d,
	0x49, 0x0a, 0x9d, 0xdb, 0x36, 0x03, 0xb7, 0x57, 0x70, 0xde, 0x01, 0x36, 0x00, 0x58, 0x3f, 0xad,
	0x96, 0xef, 0x8b, 0xe7, 0x97, 0x6c, 0x35, 0x0c, 0xa6, 0x3f, 0x00, 0x5b, 0xab, 0x5e, 0xb1, 0xfc,
	0xda, 0x2b, 0x64, 0x73, 0x88, 0x32, 0x23, 0x3d, 0x8b, 0x79, 0x27, 0x58, 0xbf, 0xd6, 0xe8, 0x9c,
	0xd4, 0x98, 0xfc, 0x07, 0xd3, 0x60, 0x1c, 0x4e, 0x42, 0xf6, 0x00, 0xd1, 0xa6, 0x72, 0x86, 0x5d,
	0xf0, 0x36, 0x0c, 0xff, 0x0b, 0xc3, 0x97, 0x75, 0x98, 0x64, 0xd8, 0xf7, 0x56, 0x7b, 0x4f, 0x83,
	0x49, 0xb8, 0x98, 0xbd, 0xdd, 0xeb, 0xbd, 0x37, 0x55, 0x5e, 0x2f, 0x16, 0x46, 0x7e, 0x56, 0x47,
	0xf9, 0x38, 0x17, 0xba, 0xb4, 0xea, 0xae, 0x6b, 0xae, 0xd9, 0xe3, 0x7a, 0x05, 0xe6, 0x67, 0x0d,
	0x9a, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x58, 0xa5, 0xb8, 0x4d, 0x60, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcServiceClient is the client API for RpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcServiceClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (RpcService_ChatClient, error)
	Push(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (RpcService_PushClient, error)
}

type rpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewRpcServiceClient(cc *grpc.ClientConn) RpcServiceClient {
	return &rpcServiceClient{cc}
}

func (c *rpcServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (RpcService_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcService_serviceDesc.Streams[0], "/chatsample.RpcService/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcServiceChatClient{stream}
	return x, nil
}

type RpcService_ChatClient interface {
	Send(*common.ChatMessage) error
	Recv() (*common.ChatMessage, error)
	grpc.ClientStream
}

type rpcServiceChatClient struct {
	grpc.ClientStream
}

func (x *rpcServiceChatClient) Send(m *common.ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcServiceChatClient) Recv() (*common.ChatMessage, error) {
	m := new(common.ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcServiceClient) Push(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (RpcService_PushClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcService_serviceDesc.Streams[1], "/chatsample.RpcService/Push", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcServicePushClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RpcService_PushClient interface {
	Recv() (*Code, error)
	grpc.ClientStream
}

type rpcServicePushClient struct {
	grpc.ClientStream
}

func (x *rpcServicePushClient) Recv() (*Code, error) {
	m := new(Code)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RpcServiceServer is the server API for RpcService service.
type RpcServiceServer interface {
	Chat(RpcService_ChatServer) error
	Push(*empty.Empty, RpcService_PushServer) error
}

// UnimplementedRpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRpcServiceServer struct {
}

func (*UnimplementedRpcServiceServer) Chat(srv RpcService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (*UnimplementedRpcServiceServer) Push(req *empty.Empty, srv RpcService_PushServer) error {
	return status.Errorf(codes.Unimplemented, "method Push not implemented")
}

func RegisterRpcServiceServer(s *grpc.Server, srv RpcServiceServer) {
	s.RegisterService(&_RpcService_serviceDesc, srv)
}

func _RpcService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServiceServer).Chat(&rpcServiceChatServer{stream})
}

type RpcService_ChatServer interface {
	Send(*common.ChatMessage) error
	Recv() (*common.ChatMessage, error)
	grpc.ServerStream
}

type rpcServiceChatServer struct {
	grpc.ServerStream
}

func (x *rpcServiceChatServer) Send(m *common.ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcServiceChatServer) Recv() (*common.ChatMessage, error) {
	m := new(common.ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RpcService_Push_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RpcServiceServer).Push(m, &rpcServicePushServer{stream})
}

type RpcService_PushServer interface {
	Send(*Code) error
	grpc.ServerStream
}

type rpcServicePushServer struct {
	grpc.ServerStream
}

func (x *rpcServicePushServer) Send(m *Code) error {
	return x.ServerStream.SendMsg(m)
}

var _RpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatsample.RpcService",
	HandlerType: (*RpcServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _RpcService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Push",
			Handler:       _RpcService_Push_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chatsample/chatsample.proto",
}