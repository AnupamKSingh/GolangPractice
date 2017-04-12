// Code generated by protoc-gen-go.
// source: manymany.proto
// DO NOT EDIT!

/*
Package manymany is a generated protocol buffer package.

It is generated from these files:
	manymany.proto

It has these top-level messages:
	Point
	Rect
*/
package manymany

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

type Point struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Point) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type Rect struct {
	A *Point `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B *Point `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
}

func (m *Rect) Reset()                    { *m = Rect{} }
func (m *Rect) String() string            { return proto.CompactTextString(m) }
func (*Rect) ProtoMessage()               {}
func (*Rect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rect) GetA() *Point {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Rect) GetB() *Point {
	if m != nil {
		return m.B
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "manymany.Point")
	proto.RegisterType((*Rect)(nil), "manymany.Rect")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Manymany service

type ManymanyClient interface {
	ServerSendMany(ctx context.Context, in *Point, opts ...grpc.CallOption) (Manymany_ServerSendManyClient, error)
}

type manymanyClient struct {
	cc *grpc.ClientConn
}

func NewManymanyClient(cc *grpc.ClientConn) ManymanyClient {
	return &manymanyClient{cc}
}

func (c *manymanyClient) ServerSendMany(ctx context.Context, in *Point, opts ...grpc.CallOption) (Manymany_ServerSendManyClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Manymany_serviceDesc.Streams[0], c.cc, "/manymany.manymany/ServerSendMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &manymanyServerSendManyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manymany_ServerSendManyClient interface {
	Recv() (*Rect, error)
	grpc.ClientStream
}

type manymanyServerSendManyClient struct {
	grpc.ClientStream
}

func (x *manymanyServerSendManyClient) Recv() (*Rect, error) {
	m := new(Rect)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Manymany service

type ManymanyServer interface {
	ServerSendMany(*Point, Manymany_ServerSendManyServer) error
}

func RegisterManymanyServer(s *grpc.Server, srv ManymanyServer) {
	s.RegisterService(&_Manymany_serviceDesc, srv)
}

func _Manymany_ServerSendMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Point)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManymanyServer).ServerSendMany(m, &manymanyServerSendManyServer{stream})
}

type Manymany_ServerSendManyServer interface {
	Send(*Rect) error
	grpc.ServerStream
}

type manymanyServerSendManyServer struct {
	grpc.ServerStream
}

func (x *manymanyServerSendManyServer) Send(m *Rect) error {
	return x.ServerStream.SendMsg(m)
}

var _Manymany_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manymany.manymany",
	HandlerType: (*ManymanyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSendMany",
			Handler:       _Manymany_ServerSendMany_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manymany.proto",
}

func init() { proto.RegisterFile("manymany.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0xcc, 0xab,
	0x04, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x0e, 0x18, 0x5f, 0x49, 0x99, 0x8b, 0x35,
	0x20, 0x3f, 0x33, 0xaf, 0x44, 0x88, 0x87, 0x8b, 0x31, 0x51, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35,
	0x88, 0x31, 0x11, 0xc4, 0x4b, 0x92, 0x60, 0x82, 0xf0, 0x92, 0x94, 0x5c, 0xb8, 0x58, 0x82, 0x52,
	0x93, 0x4b, 0x84, 0x64, 0x61, 0x6a, 0xb8, 0x8d, 0xf8, 0xf5, 0xe0, 0x46, 0x82, 0xf5, 0x83, 0x34,
	0xc9, 0xc2, 0x34, 0x61, 0x93, 0x4e, 0x32, 0x72, 0xe4, 0x82, 0x5b, 0x2b, 0x64, 0xca, 0xc5, 0x17,
	0x9c, 0x5a, 0x54, 0x96, 0x5a, 0x14, 0x9c, 0x9a, 0x97, 0xe2, 0x0b, 0x12, 0x41, 0xd7, 0x21, 0xc5,
	0x87, 0x10, 0x00, 0x59, 0xae, 0xc4, 0x60, 0xc0, 0xe8, 0x64, 0x80, 0x30, 0xc2, 0x49, 0x29, 0x33,
	0x5f, 0x2f, 0xbd, 0xa8, 0x20, 0x59, 0x2f, 0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27, 0xb5, 0x18, 0xa1,
	0x1e, 0xc6, 0x08, 0x60, 0x4c, 0x62, 0x03, 0x7b, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x26,
	0xca, 0x39, 0xf5, 0x02, 0x01, 0x00, 0x00,
}
