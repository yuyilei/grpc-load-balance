// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broker.proto

package broker

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HelloRequest struct {
	ClientId             int32    `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetClientId() int32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	NodeName             string   `protobuf:"bytes,2,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type RegisterNodeRequest struct {
	NodeId               int32    `protobuf:"varint,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 string   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNodeRequest) Reset()         { *m = RegisterNodeRequest{} }
func (m *RegisterNodeRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeRequest) ProtoMessage()    {}
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{2}
}

func (m *RegisterNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeRequest.Unmarshal(m, b)
}
func (m *RegisterNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeRequest.Marshal(b, m, deterministic)
}
func (m *RegisterNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeRequest.Merge(m, src)
}
func (m *RegisterNodeRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeRequest.Size(m)
}
func (m *RegisterNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeRequest proto.InternalMessageInfo

func (m *RegisterNodeRequest) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *RegisterNodeRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *RegisterNodeRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type RegisterNodeReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNodeReply) Reset()         { *m = RegisterNodeReply{} }
func (m *RegisterNodeReply) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeReply) ProtoMessage()    {}
func (*RegisterNodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{3}
}

func (m *RegisterNodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeReply.Unmarshal(m, b)
}
func (m *RegisterNodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeReply.Marshal(b, m, deterministic)
}
func (m *RegisterNodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeReply.Merge(m, src)
}
func (m *RegisterNodeReply) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeReply.Size(m)
}
func (m *RegisterNodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeReply proto.InternalMessageInfo

func (m *RegisterNodeReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "broker.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "broker.HelloReply")
	proto.RegisterType((*RegisterNodeRequest)(nil), "broker.RegisterNodeRequest")
	proto.RegisterType((*RegisterNodeReply)(nil), "broker.RegisterNodeReply")
}

func init() { proto.RegisterFile("broker.proto", fileDescriptor_f209535e190f2bed) }

var fileDescriptor_f209535e190f2bed = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xbb, 0x4e, 0x03, 0x31,
	0x10, 0xc4, 0x3c, 0x2e, 0x61, 0x75, 0x0d, 0x0b, 0x42, 0xc7, 0xd1, 0x20, 0x57, 0x11, 0x12, 0x11,
	0x0a, 0x12, 0x05, 0x1d, 0x69, 0x08, 0x4d, 0x0a, 0x23, 0x3e, 0xe0, 0xc2, 0xad, 0x42, 0x84, 0x93,
	0x35, 0xb6, 0x29, 0xee, 0x27, 0xf8, 0x66, 0x64, 0xc7, 0x77, 0x0a, 0x28, 0xa2, 0x48, 0xe7, 0x99,
	0x1d, 0xcd, 0xec, 0x8e, 0x21, 0x9f, 0x59, 0xfe, 0x20, 0x3b, 0x34, 0x96, 0x3d, 0x63, 0xb6, 0x46,
	0xf2, 0x1a, 0xf2, 0x09, 0x69, 0xcd, 0x8a, 0x3e, 0xbf, 0xc8, 0x79, 0x2c, 0xa1, 0xff, 0xa6, 0x17,
	0xb4, 0xf2, 0xcf, 0x75, 0x21, 0xae, 0xc4, 0xe0, 0x48, 0x75, 0x58, 0x8e, 0x01, 0x92, 0xd6, 0xe8,
	0x06, 0x0b, 0xe8, 0x2d, 0xc9, 0xb9, 0x6a, 0x4e, 0x51, 0x78, 0xac, 0x5a, 0x18, 0x3c, 0x56, 0x5c,
	0xd3, 0xb4, 0x5a, 0x52, 0xb1, 0x1f, 0x47, 0x1d, 0x96, 0xaf, 0x70, 0xaa, 0x68, 0xbe, 0x70, 0x9e,
	0xec, 0x94, 0x6b, 0x6a, 0x63, 0xcf, 0x21, 0x0b, 0x92, 0x2e, 0x34, 0x21, 0x44, 0x38, 0x7c, 0x67,
	0xe7, 0x93, 0x4d, 0x7c, 0x07, 0xce, 0xb0, 0xf5, 0xc5, 0xc1, 0x9a, 0x0b, 0x6f, 0x79, 0x03, 0x27,
	0xbf, 0x6d, 0xff, 0xdd, 0x70, 0xf4, 0x2d, 0x20, 0x1b, 0xc7, 0x02, 0xf0, 0x01, 0xfa, 0x2f, 0x55,
	0x13, 0xef, 0xc2, 0xb3, 0x61, 0xea, 0x68, 0xb3, 0x92, 0x12, 0xff, 0xb0, 0x46, 0x37, 0x72, 0x6f,
	0x20, 0x6e, 0x05, 0x4e, 0x20, 0xdf, 0x4c, 0xc5, 0xcb, 0x56, 0xb9, 0xe5, 0xc4, 0xf2, 0x62, 0xfb,
	0x30, 0xba, 0x8d, 0x1e, 0xa1, 0xf7, 0x64, 0x89, 0x3c, 0x59, 0xbc, 0xdf, 0x6d, 0xa1, 0x59, 0x16,
	0x3f, 0xf6, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x16, 0xf6, 0x62, 0xe8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerClient interface {
	SayHello(ctx context.Context, opts ...grpc.CallOption) (Broker_SayHelloClient, error)
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeReply, error)
}

type brokerClient struct {
	cc *grpc.ClientConn
}

func NewBrokerClient(cc *grpc.ClientConn) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) SayHello(ctx context.Context, opts ...grpc.CallOption) (Broker_SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Broker_serviceDesc.Streams[0], "/broker.Broker/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerSayHelloClient{stream}
	return x, nil
}

type Broker_SayHelloClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type brokerSayHelloClient struct {
	grpc.ClientStream
}

func (x *brokerSayHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *brokerSayHelloClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *brokerClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeReply, error) {
	out := new(RegisterNodeReply)
	err := c.cc.Invoke(ctx, "/broker.Broker/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServer is the server API for Broker service.
type BrokerServer interface {
	SayHello(Broker_SayHelloServer) error
	RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeReply, error)
}

// UnimplementedBrokerServer can be embedded to have forward compatible implementations.
type UnimplementedBrokerServer struct {
}

func (*UnimplementedBrokerServer) SayHello(srv Broker_SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedBrokerServer) RegisterNode(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}

func RegisterBrokerServer(s *grpc.Server, srv BrokerServer) {
	s.RegisterService(&_Broker_serviceDesc, srv)
}

func _Broker_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrokerServer).SayHello(&brokerSayHelloServer{stream})
}

type Broker_SayHelloServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type brokerSayHelloServer struct {
	grpc.ServerStream
}

func (x *brokerSayHelloServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *brokerSayHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Broker_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broker.Broker/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "broker.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _Broker_RegisterNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _Broker_SayHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "broker.proto",
}

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/broker.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broker.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "broker.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "broker.proto",
}
