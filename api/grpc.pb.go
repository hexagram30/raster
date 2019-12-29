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

func init() {
	proto.RegisterType((*GenericData)(nil), "api.GenericData")
	proto.RegisterType((*HealthRequest)(nil), "api.HealthRequest")
	proto.RegisterType((*HealthReply)(nil), "api.HealthReply")
	proto.RegisterType((*PingRequest)(nil), "api.PingRequest")
	proto.RegisterType((*PingReply)(nil), "api.PingReply")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0x1b, 0x0d, 0x81, 0x4e, 0xb1, 0xe8, 0x5c, 0x48, 0xe9, 0x85, 0x3f, 0x11, 0x41, 0x44,
	0x42, 0xb1, 0x4f, 0xd0, 0x62, 0xd1, 0xde, 0x2d, 0xf1, 0x09, 0xd6, 0x75, 0x48, 0x16, 0xda, 0xec,
	0x3a, 0x59, 0x45, 0xdf, 0xc4, 0xc7, 0x95, 0xdd, 0x46, 0x58, 0x31, 0xf4, 0x6e, 0xf6, 0xcc, 0xb7,
	0x33, 0xe7, 0x30, 0x00, 0x15, 0x5b, 0x55, 0x58, 0x36, 0xce, 0xe0, 0xa1, 0xb4, 0x3a, 0xbf, 0x84,
	0xd1, 0x23, 0x35, 0xc4, 0x5a, 0x3d, 0x48, 0x27, 0x11, 0x21, 0x7d, 0x95, 0x4e, 0x4e, 0x92, 0x8b,
	0xe4, 0x66, 0x58, 0x86, 0x3a, 0xbf, 0x82, 0xa3, 0x27, 0x92, 0x1b, 0x57, 0x97, 0xf4, 0xf6, 0x4e,
	0xad, 0xeb, 0x85, 0x56, 0x30, 0xfa, 0x85, 0xec, 0xe6, 0x0b, 0xcf, 0x00, 0x94, 0xd9, 0x5a, 0xd3,
	0x50, 0xe3, 0xda, 0x0e, 0x8c, 0x14, 0x3c, 0x85, 0x8c, 0x98, 0x0d, 0xb7, 0x93, 0x83, 0xd0, 0xeb,
	0x5e, 0xde, 0x8e, 0xd0, 0x4d, 0xb5, 0x6f, 0xd3, 0x39, 0x0c, 0x77, 0x88, 0xdf, 0xd3, 0x03, 0xdc,
	0x7f, 0x27, 0x00, 0xcf, 0xc4, 0x1f, 0x5a, 0xd1, 0x42, 0xac, 0xf1, 0x0e, 0xd2, 0x95, 0xaa, 0x0d,
	0x1e, 0x17, 0xd2, 0xea, 0x22, 0x0a, 0x3b, 0xfd, 0xa7, 0xe4, 0x03, 0x9c, 0x41, 0xb6, 0xcb, 0x81,
	0x18, 0xba, 0x7f, 0x92, 0x77, 0x3f, 0xa2, 0xa0, 0xf9, 0x00, 0x6f, 0x21, 0xf5, 0x7e, 0xba, 0xf9,
	0x91, 0xfb, 0xe9, 0x38, 0x52, 0x02, 0xbb, 0xbc, 0x86, 0x93, 0xfa, 0xb3, 0xda, 0xce, 0x67, 0x05,
	0xcb, 0xd6, 0x11, 0x7b, 0x60, 0x39, 0x2e, 0x43, 0xbd, 0x10, 0x6b, 0xe1, 0xef, 0x22, 0x92, 0x97,
	0x2c, 0x1c, 0x68, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x2d, 0x58, 0x10, 0xae, 0x01, 0x00,
	0x00,
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

// ServiceAPIServer is the server API for ServiceAPI service.
type ServiceAPIServer interface {
	// Respond with what we got
	Echo(context.Context, *GenericData) (*GenericData, error)
	// Return server health status
	Health(context.Context, *HealthRequest) (*HealthReply, error)
	// Return "pong" for "ping"
	Ping(context.Context, *PingRequest) (*PingReply, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}