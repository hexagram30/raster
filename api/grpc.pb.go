// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package api

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

// GenericData ...
type GenericData struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericData) Reset()         { *m = GenericData{} }
func (m *GenericData) String() string { return proto.CompactTextString(m) }
func (*GenericData) ProtoMessage()    {}
func (*GenericData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *GenericData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericData.Unmarshal(m, b)
}
func (m *GenericData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericData.Marshal(b, m, deterministic)
}
func (m *GenericData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericData.Merge(m, src)
}
func (m *GenericData) XXX_Size() int {
	return xxx_messageInfo_GenericData.Size(m)
}
func (m *GenericData) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericData.DiscardUnknown(m)
}

var xxx_messageInfo_GenericData proto.InternalMessageInfo

func (m *GenericData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// HealthRequest ...
type HealthRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

func (m *HealthRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// HealthReply ...
type HealthReply struct {
	Components           string   `protobuf:"bytes,1,opt,name=components,proto3" json:"components,omitempty"`
	Errors               string   `protobuf:"bytes,2,opt,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthReply) Reset()         { *m = HealthReply{} }
func (m *HealthReply) String() string { return proto.CompactTextString(m) }
func (*HealthReply) ProtoMessage()    {}
func (*HealthReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{2}
}

func (m *HealthReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthReply.Unmarshal(m, b)
}
func (m *HealthReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthReply.Marshal(b, m, deterministic)
}
func (m *HealthReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthReply.Merge(m, src)
}
func (m *HealthReply) XXX_Size() int {
	return xxx_messageInfo_HealthReply.Size(m)
}
func (m *HealthReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthReply.DiscardUnknown(m)
}

var xxx_messageInfo_HealthReply proto.InternalMessageInfo

func (m *HealthReply) GetComponents() string {
	if m != nil {
		return m.Components
	}
	return ""
}

func (m *HealthReply) GetErrors() string {
	if m != nil {
		return m.Errors
	}
	return ""
}

// PingRequest ...
type PingRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{3}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// PingReply ...
type PingReply struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{4}
}

func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// VersionRequest ...
type VersionRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{5}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func (m *VersionRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// VersionReply ...
type VersionReply struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	BuildDate            string   `protobuf:"bytes,2,opt,name=buildDate,proto3" json:"buildDate,omitempty"`
	GitCommit            string   `protobuf:"bytes,3,opt,name=gitCommit,proto3" json:"gitCommit,omitempty"`
	GitBranch            string   `protobuf:"bytes,4,opt,name=gitBranch,proto3" json:"gitBranch,omitempty"`
	GitSummary           string   `protobuf:"bytes,5,opt,name=gitSummary,proto3" json:"gitSummary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionReply) Reset()         { *m = VersionReply{} }
func (m *VersionReply) String() string { return proto.CompactTextString(m) }
func (*VersionReply) ProtoMessage()    {}
func (*VersionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{6}
}

func (m *VersionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionReply.Unmarshal(m, b)
}
func (m *VersionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionReply.Marshal(b, m, deterministic)
}
func (m *VersionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionReply.Merge(m, src)
}
func (m *VersionReply) XXX_Size() int {
	return xxx_messageInfo_VersionReply.Size(m)
}
func (m *VersionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionReply.DiscardUnknown(m)
}

var xxx_messageInfo_VersionReply proto.InternalMessageInfo

func (m *VersionReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionReply) GetBuildDate() string {
	if m != nil {
		return m.BuildDate
	}
	return ""
}

func (m *VersionReply) GetGitCommit() string {
	if m != nil {
		return m.GitCommit
	}
	return ""
}

func (m *VersionReply) GetGitBranch() string {
	if m != nil {
		return m.GitBranch
	}
	return ""
}

func (m *VersionReply) GetGitSummary() string {
	if m != nil {
		return m.GitSummary
	}
	return ""
}

func init() {
	proto.RegisterType((*GenericData)(nil), "api.GenericData")
	proto.RegisterType((*HealthRequest)(nil), "api.HealthRequest")
	proto.RegisterType((*HealthReply)(nil), "api.HealthReply")
	proto.RegisterType((*PingRequest)(nil), "api.PingRequest")
	proto.RegisterType((*PingReply)(nil), "api.PingReply")
	proto.RegisterType((*VersionRequest)(nil), "api.VersionRequest")
	proto.RegisterType((*VersionReply)(nil), "api.VersionReply")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xed, 0x6a, 0xfa, 0x30,
	0x18, 0xc5, 0xed, 0xdf, 0xfe, 0x15, 0x1f, 0x37, 0x99, 0xcf, 0x60, 0x14, 0x19, 0x7b, 0xe9, 0x36,
	0x18, 0x63, 0x14, 0x99, 0x57, 0xa0, 0x53, 0x36, 0xbf, 0x15, 0x85, 0x7d, 0x8f, 0x35, 0xb4, 0x01,
	0xdb, 0x74, 0x69, 0x94, 0x79, 0x3b, 0xbb, 0x9f, 0xdd, 0xd3, 0xc8, 0x8b, 0x5a, 0x99, 0xf8, 0x2d,
	0x39, 0xe7, 0x97, 0x27, 0xe1, 0x9c, 0x00, 0xc4, 0x22, 0x8f, 0x82, 0x5c, 0x70, 0xc9, 0xb1, 0x4a,
	0x72, 0xe6, 0xdf, 0x42, 0xf3, 0x8d, 0x66, 0x54, 0xb0, 0x68, 0x48, 0x24, 0x41, 0x04, 0x77, 0x4e,
	0x24, 0xf1, 0x9c, 0x1b, 0xe7, 0xb1, 0x31, 0xd1, 0x6b, 0xff, 0x0e, 0x4e, 0xdf, 0x29, 0x59, 0xc8,
	0x64, 0x42, 0x3f, 0x97, 0xb4, 0x90, 0x07, 0xa1, 0x11, 0x34, 0x37, 0x50, 0xbe, 0x58, 0xe3, 0x15,
	0x40, 0xc4, 0xd3, 0x9c, 0x67, 0x34, 0x93, 0x85, 0x05, 0x4b, 0x0a, 0x5e, 0x40, 0x8d, 0x0a, 0xc1,
	0x45, 0xe1, 0xfd, 0xd3, 0x9e, 0xdd, 0xa9, 0xe7, 0x84, 0x2c, 0x8b, 0x8f, 0xdd, 0x74, 0x0d, 0x0d,
	0x83, 0xa8, 0x7b, 0x0e, 0x01, 0xf7, 0xd0, 0xfa, 0xa0, 0xa2, 0x60, 0x3c, 0x3b, 0x36, 0xe6, 0xdb,
	0x81, 0x93, 0x2d, 0xa6, 0x46, 0x79, 0x50, 0x5f, 0x99, 0xbd, 0xe5, 0x36, 0x5b, 0xbc, 0x84, 0xc6,
	0x6c, 0xc9, 0x16, 0xf3, 0x21, 0x91, 0xd4, 0xbe, 0x77, 0x27, 0x28, 0x37, 0x66, 0xf2, 0x95, 0xa7,
	0x29, 0x93, 0x5e, 0xd5, 0xb8, 0x5b, 0xc1, 0xba, 0x03, 0x41, 0xb2, 0x28, 0xf1, 0xdc, 0xad, 0x6b,
	0x04, 0x15, 0x53, 0xcc, 0xe4, 0x74, 0x99, 0xa6, 0x44, 0xac, 0xbd, 0xff, 0x26, 0xa6, 0x9d, 0xf2,
	0xf2, 0xe3, 0x00, 0x4c, 0xa9, 0x58, 0xb1, 0x88, 0xf6, 0xc3, 0x31, 0x3e, 0x83, 0x3b, 0x8a, 0x12,
	0x8e, 0x67, 0x01, 0xc9, 0x59, 0x50, 0xea, 0xad, 0xf3, 0x47, 0xf1, 0x2b, 0xd8, 0x85, 0x9a, 0xa9,
	0x04, 0x51, 0xbb, 0x7b, 0x25, 0xda, 0x13, 0xa5, 0xce, 0xfc, 0x0a, 0x3e, 0x81, 0xab, 0xa2, 0xb5,
	0xf3, 0x4b, 0x45, 0x74, 0x5a, 0x25, 0xc5, 0xb0, 0x3d, 0xa8, 0xdb, 0xf8, 0xf0, 0x5c, 0x9b, 0xfb,
	0x99, 0x77, 0xda, 0xfb, 0xa2, 0x3e, 0x34, 0x78, 0x80, 0x76, 0xf2, 0x15, 0xa7, 0xbd, 0x6e, 0x20,
	0x48, 0x21, 0xa9, 0x50, 0xcc, 0xa0, 0x35, 0xd1, 0xeb, 0x7e, 0x38, 0x0e, 0xd5, 0xbf, 0x0c, 0x9d,
	0x59, 0x4d, 0x7f, 0xd0, 0xde, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xd9, 0x47, 0x97, 0xae,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceAPIClient is the client API for ServiceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceAPIClient interface {
	// Respond with what we got
	Echo(ctx context.Context, in *GenericData, opts ...grpc.CallOption) (*GenericData, error)
	// Return server health status
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthReply, error)
	// Return "pong" for "ping"
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	// Get version data
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error)
}

type serviceAPIClient struct {
	cc *grpc.ClientConn
}

func NewServiceAPIClient(cc *grpc.ClientConn) ServiceAPIClient {
	return &serviceAPIClient{cc}
}

func (c *serviceAPIClient) Echo(ctx context.Context, in *GenericData, opts ...grpc.CallOption) (*GenericData, error) {
	out := new(GenericData)
	err := c.cc.Invoke(ctx, "/api.ServiceAPI/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAPIClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthReply, error) {
	out := new(HealthReply)
	err := c.cc.Invoke(ctx, "/api.ServiceAPI/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAPIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/api.ServiceAPI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAPIClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error) {
	out := new(VersionReply)
	err := c.cc.Invoke(ctx, "/api.ServiceAPI/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAPIServer is the server API for ServiceAPI service.
type ServiceAPIServer interface {
	// Respond with what we got
	Echo(context.Context, *GenericData) (*GenericData, error)
	// Return server health status
	Health(context.Context, *HealthRequest) (*HealthReply, error)
	// Return "pong" for "ping"
	Ping(context.Context, *PingRequest) (*PingReply, error)
	// Get version data
	Version(context.Context, *VersionRequest) (*VersionReply, error)
}

// UnimplementedServiceAPIServer can be embedded to have forward compatible implementations.
type UnimplementedServiceAPIServer struct {
}

func (*UnimplementedServiceAPIServer) Echo(ctx context.Context, req *GenericData) (*GenericData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedServiceAPIServer) Health(ctx context.Context, req *HealthRequest) (*HealthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedServiceAPIServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedServiceAPIServer) Version(ctx context.Context, req *VersionRequest) (*VersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterServiceAPIServer(s *grpc.Server, srv ServiceAPIServer) {
	s.RegisterService(&_ServiceAPI_serviceDesc, srv)
}

func _ServiceAPI_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAPIServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceAPI/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAPIServer).Echo(ctx, req.(*GenericData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAPI_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAPIServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceAPI/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAPIServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAPI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAPIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceAPI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAPIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceAPI/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAPIServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ServiceAPI",
	HandlerType: (*ServiceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _ServiceAPI_Echo_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _ServiceAPI_Health_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ServiceAPI_Ping_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _ServiceAPI_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
