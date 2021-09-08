// Code generated by protoc-gen-go. DO NOT EDIT.
// source: PostRestoreAction copy.proto

package generated

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

type PostRestoreActionExecuteRequest struct {
	Plugin  string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	Restore []byte `protobuf:"bytes,2,opt,name=restore,proto3" json:"restore,omitempty"`
}

func (m *PostRestoreActionExecuteRequest) Reset()         { *m = PostRestoreActionExecuteRequest{} }
func (m *PostRestoreActionExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*PostRestoreActionExecuteRequest) ProtoMessage()    {}
func (*PostRestoreActionExecuteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0}
}

func (m *PostRestoreActionExecuteRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *PostRestoreActionExecuteRequest) GetRestore() []byte {
	if m != nil {
		return m.Restore
	}
	return nil
}

func init() {
	proto.RegisterType((*PostRestoreActionExecuteRequest)(nil), "generated.PostRestoreActionExecuteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PostRestoreAction service

type PostRestoreActionClient interface {
	Execute(ctx context.Context, in *PostRestoreActionExecuteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type postRestoreActionClient struct {
	cc *grpc.ClientConn
}

func NewPostRestoreActionClient(cc *grpc.ClientConn) PostRestoreActionClient {
	return &postRestoreActionClient{cc}
}

func (c *postRestoreActionClient) Execute(ctx context.Context, in *PostRestoreActionExecuteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/generated.PostRestoreAction/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostRestoreAction service

type PostRestoreActionServer interface {
	Execute(context.Context, *PostRestoreActionExecuteRequest) (*Empty, error)
}

func RegisterPostRestoreActionServer(s *grpc.Server, srv PostRestoreActionServer) {
	s.RegisterService(&_PostRestoreAction_serviceDesc, srv)
}

func _PostRestoreAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRestoreActionExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostRestoreActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.PostRestoreAction/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostRestoreActionServer).Execute(ctx, req.(*PostRestoreActionExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostRestoreAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.PostRestoreAction",
	HandlerType: (*PostRestoreActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _PostRestoreAction_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "PostRestoreAction copy.proto",
}

func init() { proto.RegisterFile("PostRestoreAction copy.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x09, 0xc8, 0x2f, 0x2e,
	0x09, 0x4a, 0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0x75, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x53, 0x48, 0xce,
	0x2f, 0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a,
	0x2c, 0x49, 0x4d, 0x91, 0xe2, 0x09, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0x81, 0x48, 0x28, 0x05, 0x73,
	0xc9, 0x63, 0x68, 0x74, 0xad, 0x48, 0x4d, 0x2e, 0x2d, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x24, 0xb8, 0xd8, 0x8b, 0x20, 0xda, 0x24, 0x98, 0x14, 0x18, 0x35,
	0x78, 0x82, 0x60, 0x5c, 0xa3, 0x18, 0x2e, 0x41, 0x0c, 0x43, 0x85, 0xdc, 0xb9, 0xd8, 0xa1, 0x06,
	0x0b, 0x69, 0xe9, 0xc1, 0x9d, 0xa3, 0x47, 0xc0, 0x76, 0x29, 0x01, 0x24, 0xb5, 0xae, 0xb9, 0x05,
	0x25, 0x95, 0x49, 0x6c, 0x60, 0x97, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x93, 0xa0, 0x21,
	0xe0, 0xf2, 0x00, 0x00, 0x00,
}
