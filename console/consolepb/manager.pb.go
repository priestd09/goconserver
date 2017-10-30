// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manager.proto

/*
Package consolepb is a generated protocol buffer package.

It is generated from these files:
	manager.proto

It has these top-level messages:
	NodeName
	Node
*/
package consolepb

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

type NodeName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *NodeName) Reset()                    { *m = NodeName{} }
func (m *NodeName) String() string            { return proto.CompactTextString(m) }
func (*NodeName) ProtoMessage()               {}
func (*NodeName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NodeName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Node struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Driver   string            `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	Params   map[string]string `protobuf:"bytes,3,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ondemand bool              `protobuf:"varint,4,opt,name=ondemand" json:"ondemand,omitempty"`
	Status   int32             `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Node) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Node) GetOndemand() bool {
	if m != nil {
		return m.Ondemand
	}
	return false
}

func (m *Node) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeName)(nil), "consolepb.NodeName")
	proto.RegisterType((*Node)(nil), "consolepb.Node")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ConsoleManager service

type ConsoleManagerClient interface {
	ShowNode(ctx context.Context, in *NodeName, opts ...grpc.CallOption) (*Node, error)
}

type consoleManagerClient struct {
	cc *grpc.ClientConn
}

func NewConsoleManagerClient(cc *grpc.ClientConn) ConsoleManagerClient {
	return &consoleManagerClient{cc}
}

func (c *consoleManagerClient) ShowNode(ctx context.Context, in *NodeName, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/consolepb.ConsoleManager/ShowNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsoleManager service

type ConsoleManagerServer interface {
	ShowNode(context.Context, *NodeName) (*Node, error)
}

func RegisterConsoleManagerServer(s *grpc.Server, srv ConsoleManagerServer) {
	s.RegisterService(&_ConsoleManager_serviceDesc, srv)
}

func _ConsoleManager_ShowNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleManagerServer).ShowNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consolepb.ConsoleManager/ShowNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleManagerServer).ShowNode(ctx, req.(*NodeName))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsoleManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consolepb.ConsoleManager",
	HandlerType: (*ConsoleManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowNode",
			Handler:    _ConsoleManager_ShowNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}

func init() { proto.RegisterFile("manager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0x6d, 0xfe, 0x90, 0x4e, 0x79, 0x55, 0x46, 0x91, 0x25, 0x82, 0x84, 0x9c, 0x72,
	0xca, 0x21, 0xbd, 0xa8, 0x57, 0xf5, 0x68, 0x91, 0xf8, 0x09, 0xa6, 0x66, 0x50, 0xb1, 0xbb, 0x1b,
	0x36, 0x69, 0xa5, 0x5f, 0xd4, 0xcf, 0x23, 0xbb, 0x59, 0x8a, 0x14, 0x6f, 0xf3, 0xcc, 0x0c, 0xf3,
	0xfc, 0xe6, 0x81, 0xff, 0x8a, 0x34, 0xbd, 0xb1, 0xad, 0x7b, 0x6b, 0x46, 0x83, 0xf3, 0x57, 0xa3,
	0x07, 0xb3, 0xe1, 0x7e, 0x5d, 0x5e, 0x43, 0xb6, 0x32, 0x1d, 0xaf, 0x48, 0x31, 0x22, 0xc4, 0x9a,
	0x14, 0x4b, 0x51, 0x88, 0x6a, 0xde, 0xfa, 0xba, 0xfc, 0x16, 0x10, 0xbb, 0x85, 0xbf, 0x86, 0x78,
	0x09, 0x69, 0x67, 0x3f, 0x76, 0x6c, 0xe5, 0xcc, 0x77, 0x83, 0xc2, 0x25, 0xa4, 0x3d, 0x59, 0x52,
	0x83, 0x8c, 0x8a, 0xa8, 0x5a, 0x34, 0x57, 0xf5, 0xc1, 0xb0, 0x76, 0xc7, 0xea, 0x67, 0x3f, 0x7d,
	0xd4, 0xa3, 0xdd, 0xb7, 0x61, 0x15, 0x73, 0xc8, 0x8c, 0xee, 0x58, 0x91, 0xee, 0x64, 0x5c, 0x88,
	0x2a, 0x6b, 0x0f, 0xda, 0x19, 0x0d, 0x23, 0x8d, 0xdb, 0x41, 0x26, 0x85, 0xa8, 0x92, 0x36, 0xa8,
	0xfc, 0x16, 0x16, 0xbf, 0x4e, 0xe1, 0x19, 0x44, 0x9f, 0xbc, 0x0f, 0x88, 0xae, 0xc4, 0x0b, 0x48,
	0x76, 0xb4, 0xd9, 0x72, 0x00, 0x9c, 0xc4, 0xdd, 0xec, 0x46, 0x34, 0x0f, 0x70, 0x72, 0x3f, 0x41,
	0x3d, 0x4d, 0xd9, 0x60, 0x03, 0xd9, 0xcb, 0xbb, 0xf9, 0xf2, 0xdf, 0x9e, 0x1f, 0x11, 0xbb, 0x7c,
	0xf2, 0xd3, 0xa3, 0x66, 0xf9, 0x6f, 0x9d, 0xfa, 0x40, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3a, 0x93, 0x7d, 0xf6, 0x61, 0x01, 0x00, 0x00,
}