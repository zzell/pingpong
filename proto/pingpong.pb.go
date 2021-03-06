// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pingpong.proto

package pingpong

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

type Ping struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfbf639ab46154b, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Ping) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Pong struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfbf639ab46154b, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "pingpong.Ping")
	proto.RegisterType((*Pong)(nil), "pingpong.Pong")
}

func init() { proto.RegisterFile("pingpong.proto", fileDescriptor_1cfbf639ab46154b) }

var fileDescriptor_1cfbf639ab46154b = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc8, 0xcc, 0x4b,
	0x2f, 0xc8, 0xcf, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x74,
	0xb8, 0x58, 0x02, 0x32, 0xf3, 0xd2, 0x85, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0x49, 0x82, 0x8b, 0x25, 0x20, 0x1f, 0x9b, 0x6a, 0x23,
	0x0b, 0x2e, 0x8e, 0x00, 0xa8, 0x99, 0x42, 0x20, 0x33, 0x73, 0x12, 0x2b, 0x85, 0xf8, 0xf4, 0xe0,
	0xd6, 0x82, 0xe4, 0xa4, 0x90, 0xf9, 0x20, 0xbb, 0x19, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0,
	0x4e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xae, 0x19, 0xa5, 0xa4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingpongClient is the client API for Pingpong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingpongClient interface {
	Play(ctx context.Context, opts ...grpc.CallOption) (Pingpong_PlayClient, error)
}

type pingpongClient struct {
	cc *grpc.ClientConn
}

func NewPingpongClient(cc *grpc.ClientConn) PingpongClient {
	return &pingpongClient{cc}
}

func (c *pingpongClient) Play(ctx context.Context, opts ...grpc.CallOption) (Pingpong_PlayClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Pingpong_serviceDesc.Streams[0], "/pingpong.Pingpong/Play", opts...)
	if err != nil {
		return nil, err
	}
	x := &pingpongPlayClient{stream}
	return x, nil
}

type Pingpong_PlayClient interface {
	Send(*Ping) error
	Recv() (*Pong, error)
	grpc.ClientStream
}

type pingpongPlayClient struct {
	grpc.ClientStream
}

func (x *pingpongPlayClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pingpongPlayClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingpongServer is the server API for Pingpong service.
type PingpongServer interface {
	Play(Pingpong_PlayServer) error
}

// UnimplementedPingpongServer can be embedded to have forward compatible implementations.
type UnimplementedPingpongServer struct {
}

func (*UnimplementedPingpongServer) Play(srv Pingpong_PlayServer) error {
	return status.Errorf(codes.Unimplemented, "method Play not implemented")
}

func RegisterPingpongServer(s *grpc.Server, srv PingpongServer) {
	s.RegisterService(&_Pingpong_serviceDesc, srv)
}

func _Pingpong_Play_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PingpongServer).Play(&pingpongPlayServer{stream})
}

type Pingpong_PlayServer interface {
	Send(*Pong) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type pingpongPlayServer struct {
	grpc.ServerStream
}

func (x *pingpongPlayServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pingpongPlayServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Pingpong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.Pingpong",
	HandlerType: (*PingpongServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Play",
			Handler:       _Pingpong_Play_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pingpong.proto",
}
