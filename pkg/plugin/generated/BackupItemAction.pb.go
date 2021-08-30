// Code generated by protoc-gen-go. DO NOT EDIT.
// source: BackupItemAction.proto

/*
Package generated is a generated protocol buffer package.

It is generated from these files:
	BackupItemAction.proto
	DeleteItemAction.proto
	ObjectStore.proto
	PluginLister.proto
	PreBackupAction.proto
	RestoreItemAction.proto
	Shared.proto
	VolumeSnapshotter.proto

It has these top-level messages:
	ExecuteRequest
	ExecuteResponse
	BackupItemActionAppliesToRequest
	BackupItemActionAppliesToResponse
	DeleteItemActionExecuteRequest
	DeleteItemActionAppliesToRequest
	DeleteItemActionAppliesToResponse
	PutObjectRequest
	ObjectExistsRequest
	ObjectExistsResponse
	GetObjectRequest
	Bytes
	ListCommonPrefixesRequest
	ListCommonPrefixesResponse
	ListObjectsRequest
	ListObjectsResponse
	DeleteObjectRequest
	CreateSignedURLRequest
	CreateSignedURLResponse
	ObjectStoreInitRequest
	PluginIdentifier
	ListPluginsResponse
	PreBackupActionExecuteRequest
	RestoreItemActionExecuteRequest
	RestoreItemActionExecuteResponse
	RestoreItemActionAppliesToRequest
	RestoreItemActionAppliesToResponse
	Empty
	Stack
	StackFrame
	ResourceIdentifier
	ResourceSelector
	CreateVolumeRequest
	CreateVolumeResponse
	GetVolumeInfoRequest
	GetVolumeInfoResponse
	CreateSnapshotRequest
	CreateSnapshotResponse
	DeleteSnapshotRequest
	GetVolumeIDRequest
	GetVolumeIDResponse
	SetVolumeIDRequest
	SetVolumeIDResponse
	VolumeSnapshotterInitRequest
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExecuteRequest struct {
	Plugin string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	Item   []byte `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Backup []byte `protobuf:"bytes,3,opt,name=backup,proto3" json:"backup,omitempty"`
}

func (m *ExecuteRequest) Reset()                    { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()               {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExecuteRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *ExecuteRequest) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ExecuteRequest) GetBackup() []byte {
	if m != nil {
		return m.Backup
	}
	return nil
}

type ExecuteResponse struct {
	Item            []byte                `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	AdditionalItems []*ResourceIdentifier `protobuf:"bytes,2,rep,name=additionalItems" json:"additionalItems,omitempty"`
}

func (m *ExecuteResponse) Reset()                    { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()               {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExecuteResponse) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ExecuteResponse) GetAdditionalItems() []*ResourceIdentifier {
	if m != nil {
		return m.AdditionalItems
	}
	return nil
}

type BackupItemActionAppliesToRequest struct {
	Plugin string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
}

func (m *BackupItemActionAppliesToRequest) Reset()         { *m = BackupItemActionAppliesToRequest{} }
func (m *BackupItemActionAppliesToRequest) String() string { return proto.CompactTextString(m) }
func (*BackupItemActionAppliesToRequest) ProtoMessage()    {}
func (*BackupItemActionAppliesToRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

func (m *BackupItemActionAppliesToRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

type BackupItemActionAppliesToResponse struct {
	ResourceSelector *ResourceSelector `protobuf:"bytes,1,opt,name=ResourceSelector" json:"ResourceSelector,omitempty"`
}

func (m *BackupItemActionAppliesToResponse) Reset()         { *m = BackupItemActionAppliesToResponse{} }
func (m *BackupItemActionAppliesToResponse) String() string { return proto.CompactTextString(m) }
func (*BackupItemActionAppliesToResponse) ProtoMessage()    {}
func (*BackupItemActionAppliesToResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

func (m *BackupItemActionAppliesToResponse) GetResourceSelector() *ResourceSelector {
	if m != nil {
		return m.ResourceSelector
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteRequest)(nil), "generated.ExecuteRequest")
	proto.RegisterType((*ExecuteResponse)(nil), "generated.ExecuteResponse")
	proto.RegisterType((*BackupItemActionAppliesToRequest)(nil), "generated.BackupItemActionAppliesToRequest")
	proto.RegisterType((*BackupItemActionAppliesToResponse)(nil), "generated.BackupItemActionAppliesToResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BackupItemAction service

type BackupItemActionClient interface {
	AppliesTo(ctx context.Context, in *BackupItemActionAppliesToRequest, opts ...grpc.CallOption) (*BackupItemActionAppliesToResponse, error)
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type backupItemActionClient struct {
	cc *grpc.ClientConn
}

func NewBackupItemActionClient(cc *grpc.ClientConn) BackupItemActionClient {
	return &backupItemActionClient{cc}
}

func (c *backupItemActionClient) AppliesTo(ctx context.Context, in *BackupItemActionAppliesToRequest, opts ...grpc.CallOption) (*BackupItemActionAppliesToResponse, error) {
	out := new(BackupItemActionAppliesToResponse)
	err := grpc.Invoke(ctx, "/generated.BackupItemAction/AppliesTo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupItemActionClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := grpc.Invoke(ctx, "/generated.BackupItemAction/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BackupItemAction service

type BackupItemActionServer interface {
	AppliesTo(context.Context, *BackupItemActionAppliesToRequest) (*BackupItemActionAppliesToResponse, error)
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

func RegisterBackupItemActionServer(s *grpc.Server, srv BackupItemActionServer) {
	s.RegisterService(&_BackupItemAction_serviceDesc, srv)
}

func _BackupItemAction_AppliesTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupItemActionAppliesToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).AppliesTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BackupItemAction/AppliesTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).AppliesTo(ctx, req.(*BackupItemActionAppliesToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupItemAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BackupItemAction/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupItemAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.BackupItemAction",
	HandlerType: (*BackupItemActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppliesTo",
			Handler:    _BackupItemAction_AppliesTo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _BackupItemAction_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BackupItemAction.proto",
}

func init() { proto.RegisterFile("BackupItemAction.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x2b, 0x95, 0x4e, 0x8b, 0x2d, 0x7b, 0x28, 0x31, 0x22, 0xc4, 0x9c, 0x02, 0x4a,
	0x0e, 0xf1, 0xe6, 0xc9, 0x0a, 0x52, 0x7a, 0xdd, 0xf6, 0x05, 0xd2, 0x64, 0x5a, 0x17, 0xd3, 0xdd,
	0x75, 0x77, 0x03, 0x3e, 0x9c, 0x0f, 0x27, 0xd9, 0x6e, 0x43, 0x8c, 0xc5, 0x7a, 0xcb, 0x64, 0xe6,
	0xff, 0xe7, 0xfb, 0xd9, 0x81, 0xd9, 0x4b, 0x96, 0xbf, 0x57, 0x72, 0x69, 0x70, 0x3f, 0xcf, 0x0d,
	0x13, 0x3c, 0x91, 0x4a, 0x18, 0x41, 0x86, 0x3b, 0xe4, 0xa8, 0x32, 0x83, 0x45, 0x30, 0x5e, 0xbd,
	0x65, 0x0a, 0x8b, 0x43, 0x23, 0x5a, 0xc3, 0xd5, 0xeb, 0x27, 0xe6, 0x95, 0x41, 0x8a, 0x1f, 0x15,
	0x6a, 0x43, 0x66, 0x30, 0x90, 0x65, 0xb5, 0x63, 0xdc, 0xf7, 0x42, 0x2f, 0x1e, 0x52, 0x57, 0x11,
	0x02, 0x17, 0xcc, 0xe0, 0xde, 0xef, 0x85, 0x5e, 0x3c, 0xa6, 0xf6, 0xbb, 0x9e, 0xdd, 0xd8, 0x85,
	0x7e, 0xdf, 0xfe, 0x75, 0x55, 0xc4, 0x61, 0xd2, 0xb8, 0x6a, 0x29, 0xb8, 0xc6, 0x46, 0xee, 0xb5,
	0xe4, 0x0b, 0x98, 0x64, 0x45, 0xc1, 0x6a, 0xce, 0xac, 0xac, 0x99, 0xb5, 0xdf, 0x0b, 0xfb, 0xf1,
	0x28, 0xbd, 0x4d, 0x1a, 0xde, 0x84, 0xa2, 0x16, 0x95, 0xca, 0x71, 0x59, 0x20, 0x37, 0x6c, 0xcb,
	0x50, 0xd1, 0xae, 0x2a, 0x7a, 0x82, 0xb0, 0x1b, 0x7c, 0x2e, 0x65, 0xc9, 0x50, 0xaf, 0xc5, 0x99,
	0x5c, 0x51, 0x09, 0x77, 0x7f, 0x68, 0x1d, 0xfd, 0x02, 0xa6, 0x47, 0x8e, 0x15, 0x96, 0x98, 0x1b,
	0xa1, 0xac, 0xcd, 0x28, 0xbd, 0x39, 0x81, 0x7a, 0x1c, 0xa1, 0xbf, 0x44, 0xe9, 0x97, 0x07, 0xd3,
	0xee, 0x3a, 0xb2, 0x85, 0x61, 0xb3, 0x92, 0xdc, 0xb7, 0x0c, 0xcf, 0x85, 0x0a, 0x1e, 0xfe, 0x37,
	0xec, 0x52, 0x3c, 0xc3, 0xa5, 0x7b, 0x16, 0x72, 0xdd, 0x12, 0xfe, 0x3c, 0x80, 0x20, 0x38, 0xd5,
	0x3a, 0x38, 0x6c, 0x06, 0xf6, 0x6a, 0x1e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x45, 0xdb, 0x5d,
	0x9f, 0x68, 0x02, 0x00, 0x00,
}
