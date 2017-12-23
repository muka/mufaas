// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	FunctionInfo
	AddRequest
	AddResponse
	RemoveRequest
	RemoveResponse
	ListRequest
	ListResponse
	RunRequest
	RunResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type FunctionInfo struct {
	ID    string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Type  string `protobuf:"bytes,3,opt,name=Type" json:"Type,omitempty"`
	Cmd   string `protobuf:"bytes,4,opt,name=Cmd" json:"Cmd,omitempty"`
	Error string `protobuf:"bytes,5,opt,name=Error" json:"Error,omitempty"`
}

func (m *FunctionInfo) Reset()                    { *m = FunctionInfo{} }
func (m *FunctionInfo) String() string            { return proto.CompactTextString(m) }
func (*FunctionInfo) ProtoMessage()               {}
func (*FunctionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FunctionInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *FunctionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FunctionInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FunctionInfo) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *FunctionInfo) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type AddRequest struct {
	Info   *FunctionInfo `protobuf:"bytes,1,opt,name=Info" json:"Info,omitempty"`
	Source []byte        `protobuf:"bytes,2,opt,name=Source,proto3" json:"Source,omitempty"`
	Image  string        `protobuf:"bytes,3,opt,name=Image" json:"Image,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddRequest) GetInfo() *FunctionInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *AddRequest) GetSource() []byte {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AddRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type AddResponse struct {
	Info *FunctionInfo `protobuf:"bytes,1,opt,name=Info" json:"Info,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddResponse) GetInfo() *FunctionInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type RemoveRequest struct {
	Filter []string `protobuf:"bytes,1,rep,name=Filter" json:"Filter,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RemoveRequest) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

type RemoveResponse struct {
	Functions []*FunctionInfo `protobuf:"bytes,1,rep,name=Functions" json:"Functions,omitempty"`
}

func (m *RemoveResponse) Reset()                    { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()               {}
func (*RemoveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveResponse) GetFunctions() []*FunctionInfo {
	if m != nil {
		return m.Functions
	}
	return nil
}

type ListRequest struct {
	Filter []string `protobuf:"bytes,1,rep,name=Filter" json:"Filter,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListRequest) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

type ListResponse struct {
	Functions []*FunctionInfo `protobuf:"bytes,1,rep,name=Functions" json:"Functions,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListResponse) GetFunctions() []*FunctionInfo {
	if m != nil {
		return m.Functions
	}
	return nil
}

type RunRequest struct {
	Name  string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Args  [][]byte `protobuf:"bytes,2,rep,name=Args,proto3" json:"Args,omitempty"`
	Env   []string `protobuf:"bytes,3,rep,name=Env" json:"Env,omitempty"`
	Stdin []byte   `protobuf:"bytes,4,opt,name=Stdin,proto3" json:"Stdin,omitempty"`
}

func (m *RunRequest) Reset()                    { *m = RunRequest{} }
func (m *RunRequest) String() string            { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()               {}
func (*RunRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RunRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RunRequest) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *RunRequest) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *RunRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

type RunResponse struct {
	Output []byte `protobuf:"bytes,1,opt,name=Output,proto3" json:"Output,omitempty"`
	Err    []byte `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (m *RunResponse) Reset()                    { *m = RunResponse{} }
func (m *RunResponse) String() string            { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()               {}
func (*RunResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RunResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *RunResponse) GetErr() []byte {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*FunctionInfo)(nil), "api.FunctionInfo")
	proto.RegisterType((*AddRequest)(nil), "api.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "api.AddResponse")
	proto.RegisterType((*RemoveRequest)(nil), "api.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "api.RemoveResponse")
	proto.RegisterType((*ListRequest)(nil), "api.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "api.ListResponse")
	proto.RegisterType((*RunRequest)(nil), "api.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "api.RunResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MufaasService service

type MufaasServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type mufaasServiceClient struct {
	cc *grpc.ClientConn
}

func NewMufaasServiceClient(cc *grpc.ClientConn) MufaasServiceClient {
	return &mufaasServiceClient{cc}
}

func (c *mufaasServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/api.MufaasService/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mufaasServiceClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := grpc.Invoke(ctx, "/api.MufaasService/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mufaasServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/api.MufaasService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mufaasServiceClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := grpc.Invoke(ctx, "/api.MufaasService/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MufaasService service

type MufaasServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Run(context.Context, *RunRequest) (*RunResponse, error)
}

func RegisterMufaasServiceServer(s *grpc.Server, srv MufaasServiceServer) {
	s.RegisterService(&_MufaasService_serviceDesc, srv)
}

func _MufaasService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MufaasServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MufaasService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MufaasServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MufaasService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MufaasServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MufaasService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MufaasServiceServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MufaasService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MufaasServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MufaasService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MufaasServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MufaasService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MufaasServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MufaasService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MufaasServiceServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MufaasService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MufaasService",
	HandlerType: (*MufaasServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MufaasService_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _MufaasService_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MufaasService_List_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _MufaasService_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x76, 0x6a, 0x96, 0x13, 0x27, 0x6d, 0x95, 0x2d, 0x78, 0x61, 0x17, 0x41, 0x50, 0x16,
	0x7a, 0x51, 0xb3, 0x6e, 0x30, 0xe8, 0x4d, 0x17, 0xfa, 0xc3, 0x02, 0xfb, 0x01, 0x65, 0x97, 0xbb,
	0xd1, 0x62, 0x35, 0x18, 0x12, 0xc9, 0x93, 0xe5, 0x40, 0x19, 0xbb, 0xd9, 0x2b, 0xec, 0x09, 0xf6,
	0x4c, 0x7b, 0x85, 0x3d, 0xc8, 0xd0, 0x91, 0x4c, 0x12, 0x08, 0x6c, 0xf4, 0xee, 0x9c, 0x4f, 0x47,
	0xdf, 0xf7, 0xe9, 0x7c, 0x36, 0xb4, 0x79, 0x59, 0x9c, 0x95, 0x5a, 0x19, 0x45, 0x22, 0x5e, 0x16,
	0xc3, 0x67, 0x0b, 0xa5, 0x16, 0x4b, 0x91, 0xf1, 0xb2, 0xc8, 0xb8, 0x94, 0xca, 0x70, 0x53, 0x28,
	0x59, 0xb9, 0x11, 0x2a, 0x21, 0xb9, 0xad, 0xe5, 0xdc, 0x42, 0x53, 0x79, 0xa7, 0x48, 0x0f, 0xc2,
	0xe9, 0x75, 0x1a, 0x8c, 0x82, 0x71, 0x9b, 0x85, 0xd3, 0x6b, 0x42, 0xa0, 0xf5, 0x81, 0xaf, 0x44,
	0x1a, 0x22, 0x82, 0xb5, 0xc5, 0x3e, 0xdd, 0x97, 0x22, 0x8d, 0x1c, 0x66, 0x6b, 0x72, 0x04, 0xd1,
	0xd5, 0x2a, 0x4f, 0x5b, 0x08, 0xd9, 0x92, 0x3c, 0x86, 0x83, 0x1b, 0xad, 0x95, 0x4e, 0x0f, 0x10,
	0x73, 0x0d, 0xe5, 0x00, 0x93, 0x3c, 0x67, 0xe2, 0x6b, 0x2d, 0x2a, 0x43, 0x4e, 0xa0, 0x65, 0x55,
	0x51, 0xaf, 0x73, 0x7e, 0x7c, 0x66, 0xad, 0x6f, 0xdb, 0x61, 0x78, 0x4c, 0x06, 0x10, 0xcf, 0x54,
	0xad, 0xe7, 0xce, 0x46, 0xc2, 0x7c, 0x67, 0x25, 0xa6, 0x2b, 0xbe, 0x68, 0x9c, 0xb8, 0x86, 0xbe,
	0x82, 0x0e, 0x4a, 0x54, 0xa5, 0x92, 0x95, 0xf8, 0x4f, 0x0d, 0xfa, 0x1c, 0xba, 0x4c, 0xac, 0xd4,
	0x5a, 0x34, 0xde, 0x06, 0x10, 0xdf, 0x16, 0x4b, 0x23, 0x74, 0x1a, 0x8c, 0xa2, 0x71, 0x9b, 0xf9,
	0x8e, 0x4e, 0xa0, 0xd7, 0x0c, 0x7a, 0x85, 0x0c, 0xda, 0x0d, 0x61, 0x85, 0xc3, 0x7b, 0x65, 0x36,
	0x33, 0xf4, 0x04, 0x3a, 0xef, 0x8a, 0xca, 0xfc, 0x4b, 0xe9, 0x12, 0x12, 0x37, 0xf6, 0x50, 0x9d,
	0xcf, 0x00, 0xac, 0x96, 0x8d, 0x4c, 0x13, 0x65, 0xb0, 0x1b, 0xe5, 0x44, 0x2f, 0xaa, 0x34, 0x1c,
	0x45, 0xe3, 0x84, 0x61, 0x6d, 0xa3, 0xbc, 0x91, 0xeb, 0x34, 0x42, 0x2f, 0xb6, 0xb4, 0x7b, 0x9e,
	0x99, 0xbc, 0x90, 0x18, 0x6f, 0xc2, 0x5c, 0x43, 0x5f, 0x43, 0x07, 0xd9, 0xbd, 0xbb, 0x01, 0xc4,
	0x1f, 0x6b, 0x53, 0xd6, 0x06, 0x05, 0x12, 0xe6, 0x3b, 0xa4, 0xd3, 0xda, 0x27, 0x67, 0xcb, 0xf3,
	0x5f, 0x21, 0x74, 0xdf, 0xd7, 0x77, 0x9c, 0x57, 0x33, 0xa1, 0xd7, 0xc5, 0x5c, 0x90, 0x4b, 0x88,
	0x26, 0x79, 0x4e, 0x0e, 0xf1, 0x35, 0x9b, 0xef, 0x63, 0x78, 0xb4, 0x01, 0x9c, 0x0a, 0x7d, 0xf2,
	0xe3, 0xf7, 0x9f, 0x9f, 0xe1, 0x21, 0x85, 0x6c, 0xfd, 0x22, 0xcb, 0x45, 0xb9, 0x54, 0xf7, 0x17,
	0xc1, 0x29, 0x79, 0x0b, 0xb1, 0x0b, 0x85, 0x10, 0xbc, 0xb2, 0x13, 0xe5, 0xb0, 0xbf, 0x83, 0xed,
	0x63, 0xd2, 0x78, 0x66, 0x99, 0xde, 0x40, 0xcb, 0x2e, 0x9d, 0x38, 0xe9, 0xad, 0x98, 0x86, 0xc7,
	0x5b, 0x88, 0xe7, 0xe8, 0x23, 0x47, 0x97, 0x3e, 0xb2, 0x1c, 0xcb, 0xa2, 0x32, 0x96, 0xe1, 0x0a,
	0x22, 0x56, 0x4b, 0xff, 0x98, 0xcd, 0xfe, 0xfd, 0x63, 0xb6, 0x56, 0x46, 0x9f, 0xe2, 0xf5, 0x3e,
	0xed, 0xa1, 0x85, 0x5a, 0x66, 0xdf, 0x6c, 0x28, 0xdf, 0x2f, 0x82, 0xd3, 0x2f, 0x31, 0xfe, 0x9e,
	0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xab, 0xb9, 0xd6, 0xce, 0x03, 0x00, 0x00,
}
