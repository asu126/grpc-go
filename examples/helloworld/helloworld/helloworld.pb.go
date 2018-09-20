// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	HelloBytes
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HelloBytes struct {
	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *HelloBytes) Reset()                    { *m = HelloBytes{} }
func (m *HelloBytes) String() string            { return proto.CompactTextString(m) }
func (*HelloBytes) ProtoMessage()               {}
func (*HelloBytes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HelloBytes) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*HelloBytes)(nil), "helloworld.HelloBytes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloAgain(ctx context.Context, in *HelloBytes, opts ...grpc.CallOption) (*HelloBytes, error)
	// 服务器端流式 RPC
	ServerStream(ctx context.Context, in *HelloBytes, opts ...grpc.CallOption) (Greeter_ServerStreamClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloBytes, opts ...grpc.CallOption) (*HelloBytes, error) {
	out := new(HelloBytes)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHelloAgain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ServerStream(ctx context.Context, in *HelloBytes, opts ...grpc.CallOption) (Greeter_ServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/helloworld.Greeter/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_ServerStreamClient interface {
	Recv() (*HelloBytes, error)
	grpc.ClientStream
}

type greeterServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterServerStreamClient) Recv() (*HelloBytes, error) {
	m := new(HelloBytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(context.Context, *HelloBytes) (*HelloBytes, error)
	// 服务器端流式 RPC
	ServerStream(*HelloBytes, Greeter_ServerStreamServer) error
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
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloBytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloBytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloBytes)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).ServerStream(m, &greeterServerStreamServer{stream})
}

type Greeter_ServerStreamServer interface {
	Send(*HelloBytes) error
	grpc.ServerStream
}

type greeterServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterServerStreamServer) Send(m *HelloBytes) error {
	return x.ServerStream.SendMsg(m)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _Greeter_ServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd1, 0xbd, 0x4a, 0x04, 0x31,
	0x10, 0x07, 0xf0, 0x0b, 0x88, 0xa7, 0xc3, 0x8a, 0x32, 0x85, 0x2c, 0x67, 0x23, 0x29, 0xc4, 0x2a,
	0x1c, 0xda, 0x0b, 0x97, 0x46, 0xcb, 0xe3, 0xb6, 0xb0, 0x8e, 0x3a, 0xac, 0x07, 0xc9, 0x26, 0x4e,
	0xe2, 0x47, 0xde, 0xd2, 0x47, 0x92, 0x0d, 0xee, 0xba, 0xf8, 0x51, 0xd8, 0xcd, 0x4c, 0x7e, 0x13,
	0xfe, 0x30, 0x70, 0xf4, 0x48, 0xd6, 0xfa, 0x57, 0xcf, 0xf6, 0x41, 0x05, 0xf6, 0xc9, 0x23, 0x7c,
	0x4d, 0xa4, 0x84, 0xea, 0xa6, 0xef, 0x36, 0xf4, 0xf4, 0x4c, 0x31, 0x21, 0xc2, 0x4e, 0x67, 0x1c,
	0xd5, 0xe2, 0x54, 0x9c, 0xef, 0x6f, 0x4a, 0x2d, 0xcf, 0x00, 0x3e, 0x4d, 0xb0, 0x19, 0x6b, 0x98,
	0x3b, 0x8a, 0xd1, 0xb4, 0x03, 0x1a, 0xda, 0xd1, 0xe9, 0x9c, 0x28, 0x7e, 0x77, 0xd5, 0xe8, 0x2e,
	0xde, 0x05, 0xcc, 0xaf, 0x99, 0x28, 0x11, 0xe3, 0x15, 0xec, 0x35, 0x26, 0x97, 0x35, 0xac, 0xd5,
	0x24, 0xea, 0x34, 0xd5, 0xe2, 0xf8, 0x97, 0x97, 0x60, 0xb3, 0x9c, 0xe1, 0x0a, 0x0e, 0x86, 0xfd,
	0x55, 0x6b, 0xb6, 0x1d, 0xfe, 0xa4, 0x25, 0xce, 0xe2, 0x8f, 0xb9, 0x9c, 0xa1, 0x86, 0xaa, 0x21,
	0x7e, 0x21, 0x6e, 0x12, 0x93, 0x71, 0xff, 0xff, 0x61, 0x29, 0xf4, 0x12, 0x4e, 0xb6, 0x5e, 0xb5,
	0x1c, 0xee, 0x15, 0xbd, 0x19, 0x17, 0x2c, 0xc5, 0x89, 0xd6, 0x87, 0x85, 0xdf, 0xf6, 0xf5, 0xba,
	0x3f, 0xc1, 0x5a, 0xdc, 0xed, 0x96, 0x5b, 0x5c, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x5e,
	0xeb, 0x3c, 0x9f, 0x01, 0x00, 0x00,
}
