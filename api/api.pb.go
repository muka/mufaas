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
	Name  []string `protobuf:"bytes,1,rep,name=Name" json:"Name,omitempty"`
	Force bool     `protobuf:"varint,2,opt,name=Force" json:"Force,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RemoveRequest) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RemoveRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
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
	Args  []string `protobuf:"bytes,2,rep,name=Args" json:"Args,omitempty"`
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

func (m *RunRequest) GetArgs() []string {
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
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x76, 0x62, 0x9a, 0x13, 0x27, 0x6d, 0x95, 0x2e, 0x78, 0x61, 0x17, 0x45, 0x50, 0x18,
	0xbd, 0xa8, 0x59, 0x37, 0x18, 0xeb, 0x4d, 0x17, 0xda, 0x86, 0x05, 0xf6, 0x03, 0xca, 0x2e, 0x77,
	0xa3, 0xc5, 0x6a, 0x30, 0x24, 0x92, 0x27, 0xcb, 0x81, 0x32, 0x76, 0xb3, 0x57, 0xd8, 0x13, 0xec,
	0x99, 0xf6, 0x0a, 0x7b, 0x90, 0xa1, 0x23, 0x79, 0x49, 0x58, 0x2f, 0x46, 0xef, 0xce, 0xf9, 0x24,
	0x7d, 0xdf, 0x77, 0xce, 0x67, 0x43, 0x87, 0x97, 0xc5, 0x59, 0xa9, 0x95, 0x51, 0x24, 0xe2, 0x65,
	0x31, 0x7a, 0xb2, 0x50, 0x6a, 0xb1, 0x14, 0x19, 0x2f, 0x8b, 0x8c, 0x4b, 0xa9, 0x0c, 0x37, 0x85,
	0x92, 0x95, 0xbb, 0x42, 0x25, 0x24, 0x93, 0x5a, 0xce, 0x2d, 0x34, 0x95, 0xb7, 0x8a, 0xf4, 0x21,
	0x9c, 0x5e, 0xa7, 0xc1, 0x71, 0xf0, 0xb4, 0xc3, 0xc2, 0xe9, 0x35, 0x21, 0xd0, 0x7a, 0xcf, 0x57,
	0x22, 0x0d, 0x11, 0xc1, 0xda, 0x62, 0x1f, 0xef, 0x4a, 0x91, 0x46, 0x0e, 0xb3, 0x35, 0x39, 0x80,
	0xe8, 0x6a, 0x95, 0xa7, 0x2d, 0x84, 0x6c, 0x49, 0x8e, 0xa0, 0x7d, 0xa3, 0xb5, 0xd2, 0x69, 0x1b,
	0x31, 0xd7, 0x50, 0x0e, 0x30, 0xce, 0x73, 0x26, 0xbe, 0xd4, 0xa2, 0x32, 0xe4, 0x04, 0x5a, 0x56,
	0x15, 0xf5, 0xba, 0xe7, 0x87, 0x67, 0xd6, 0xfa, 0xb6, 0x1d, 0x86, 0xc7, 0x64, 0x08, 0xf1, 0x4c,
	0xd5, 0x7a, 0xee, 0x6c, 0x24, 0xcc, 0x77, 0x56, 0x62, 0xba, 0xe2, 0x8b, 0xc6, 0x89, 0x6b, 0xe8,
	0x0b, 0xe8, 0xa2, 0x44, 0x55, 0x2a, 0x59, 0x89, 0xff, 0xd4, 0xa0, 0xaf, 0xa0, 0xc7, 0xc4, 0x4a,
	0xad, 0x45, 0xe3, 0xad, 0x99, 0x3c, 0x38, 0x8e, 0xfe, 0x4e, 0x7e, 0x04, 0xed, 0x89, 0x6a, 0x7c,
	0xec, 0x31, 0xd7, 0xd0, 0x31, 0xf4, 0x9b, 0xa7, 0x5e, 0x33, 0x83, 0x4e, 0x23, 0x51, 0x21, 0xc1,
	0xbd, 0xc2, 0x9b, 0x3b, 0xf4, 0x04, 0xba, 0x6f, 0x8b, 0xca, 0x34, 0xda, 0x43, 0x88, 0x27, 0xc5,
	0xd2, 0x08, 0xed, 0xd5, 0x7d, 0x47, 0x2f, 0x21, 0x71, 0xd7, 0x1e, 0xaa, 0xf3, 0x09, 0x80, 0xd5,
	0xf2, 0xdf, 0x11, 0x77, 0xc2, 0x1d, 0xeb, 0x45, 0x95, 0x86, 0x6e, 0x6c, 0x5b, 0xdb, 0x70, 0x6f,
	0xe4, 0x3a, 0x8d, 0x10, 0xb2, 0xa5, 0x5d, 0xc4, 0xcc, 0xe4, 0x85, 0xc4, 0xc0, 0x13, 0xe6, 0x1a,
	0xfa, 0x12, 0xba, 0xc8, 0xee, 0xdd, 0x0d, 0x21, 0xfe, 0x50, 0x9b, 0xb2, 0x36, 0x28, 0x90, 0x30,
	0xdf, 0x21, 0x9d, 0xd6, 0x3e, 0x4b, 0x5b, 0x9e, 0xff, 0x0c, 0xa1, 0xf7, 0xae, 0xbe, 0xe5, 0xbc,
	0x9a, 0x09, 0xbd, 0x2e, 0xe6, 0x82, 0x5c, 0x42, 0x34, 0xce, 0x73, 0xb2, 0x8f, 0xd3, 0x6c, 0xbe,
	0x98, 0xd1, 0xc1, 0x06, 0x70, 0x2a, 0xf4, 0xd1, 0xf7, 0x5f, 0xbf, 0x7f, 0x84, 0xfb, 0x14, 0xb2,
	0xf5, 0xb3, 0x2c, 0x17, 0xe5, 0x52, 0xdd, 0x5d, 0x04, 0xa7, 0xe4, 0x0d, 0xc4, 0x2e, 0x14, 0x42,
	0xf0, 0xc9, 0x4e, 0xb8, 0xa3, 0xc1, 0x0e, 0x76, 0x1f, 0x93, 0xc6, 0x33, 0xcb, 0xf4, 0x1a, 0x5a,
	0x76, 0xe9, 0xc4, 0x49, 0x6f, 0xc5, 0x34, 0x3a, 0xdc, 0x42, 0x3c, 0xc7, 0x00, 0x39, 0x7a, 0x74,
	0xcf, 0x72, 0x2c, 0x8b, 0xca, 0x58, 0x86, 0x2b, 0x88, 0x58, 0x2d, 0xfd, 0x30, 0x9b, 0xfd, 0xfb,
	0x61, 0xb6, 0x56, 0x46, 0x1f, 0xe3, 0xf3, 0x01, 0xed, 0xa3, 0x85, 0x5a, 0x66, 0x5f, 0x6d, 0x28,
	0xdf, 0x2e, 0x82, 0xd3, 0xcf, 0x31, 0xfe, 0xb0, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0x25, 0x80, 0xc2, 0xe0, 0x03, 0x00, 0x00,
}
