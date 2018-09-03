// Code generated by protoc-gen-go. DO NOT EDIT.
// source: newclientdemand.proto

package pb

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

// Configuration message
type ConfigFileReq struct {
	Keypublic            string   `protobuf:"bytes,1,opt,name=keypublic,proto3" json:"keypublic,omitempty"`
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigFileReq) Reset()         { *m = ConfigFileReq{} }
func (m *ConfigFileReq) String() string { return proto.CompactTextString(m) }
func (*ConfigFileReq) ProtoMessage()    {}
func (*ConfigFileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_newclientdemand_8d56dd2f2dc0e096, []int{0}
}
func (m *ConfigFileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileReq.Unmarshal(m, b)
}
func (m *ConfigFileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileReq.Marshal(b, m, deterministic)
}
func (dst *ConfigFileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileReq.Merge(dst, src)
}
func (m *ConfigFileReq) XXX_Size() int {
	return xxx_messageInfo_ConfigFileReq.Size(m)
}
func (m *ConfigFileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileReq proto.InternalMessageInfo

func (m *ConfigFileReq) GetKeypublic() string {
	if m != nil {
		return m.Keypublic
	}
	return ""
}

func (m *ConfigFileReq) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// Ack to webserver
type AckWeb struct {
	Ack                  bool     `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckWeb) Reset()         { *m = AckWeb{} }
func (m *AckWeb) String() string { return proto.CompactTextString(m) }
func (*AckWeb) ProtoMessage()    {}
func (*AckWeb) Descriptor() ([]byte, []int) {
	return fileDescriptor_newclientdemand_8d56dd2f2dc0e096, []int{1}
}
func (m *AckWeb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckWeb.Unmarshal(m, b)
}
func (m *AckWeb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckWeb.Marshal(b, m, deterministic)
}
func (dst *AckWeb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckWeb.Merge(dst, src)
}
func (m *AckWeb) XXX_Size() int {
	return xxx_messageInfo_AckWeb.Size(m)
}
func (m *AckWeb) XXX_DiscardUnknown() {
	xxx_messageInfo_AckWeb.DiscardUnknown(m)
}

var xxx_messageInfo_AckWeb proto.InternalMessageInfo

func (m *AckWeb) GetAck() bool {
	if m != nil {
		return m.Ack
	}
	return false
}

func init() {
	proto.RegisterType((*ConfigFileReq)(nil), "newclientdemand.ConfigFileReq")
	proto.RegisterType((*AckWeb)(nil), "newclientdemand.AckWeb")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NewClientDemandClient is the client API for NewClientDemand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NewClientDemandClient interface {
	GetClientDemand(ctx context.Context, in *ConfigFileReq, opts ...grpc.CallOption) (*AckWeb, error)
}

type newClientDemandClient struct {
	cc *grpc.ClientConn
}

func NewNewClientDemandClient(cc *grpc.ClientConn) NewClientDemandClient {
	return &newClientDemandClient{cc}
}

func (c *newClientDemandClient) GetClientDemand(ctx context.Context, in *ConfigFileReq, opts ...grpc.CallOption) (*AckWeb, error) {
	out := new(AckWeb)
	err := c.cc.Invoke(ctx, "/newclientdemand.NewClientDemand/GetClientDemand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewClientDemandServer is the server API for NewClientDemand service.
type NewClientDemandServer interface {
	GetClientDemand(context.Context, *ConfigFileReq) (*AckWeb, error)
}

func RegisterNewClientDemandServer(s *grpc.Server, srv NewClientDemandServer) {
	s.RegisterService(&_NewClientDemand_serviceDesc, srv)
}

func _NewClientDemand_GetClientDemand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewClientDemandServer).GetClientDemand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/newclientdemand.NewClientDemand/GetClientDemand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewClientDemandServer).GetClientDemand(ctx, req.(*ConfigFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NewClientDemand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "newclientdemand.NewClientDemand",
	HandlerType: (*NewClientDemandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClientDemand",
			Handler:    _NewClientDemand_GetClientDemand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "newclientdemand.proto",
}

func init() {
	proto.RegisterFile("newclientdemand.proto", fileDescriptor_newclientdemand_8d56dd2f2dc0e096)
}

var fileDescriptor_newclientdemand_8d56dd2f2dc0e096 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4b, 0x2d, 0x4f,
	0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x49, 0x49, 0xcd, 0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x56, 0xf2, 0xe4, 0xe2, 0x75, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x77,
	0xcb, 0xcc, 0x49, 0x0d, 0x4a, 0x2d, 0x14, 0x92, 0xe1, 0xe2, 0xcc, 0x4e, 0xad, 0x2c, 0x28, 0x4d,
	0xca, 0xc9, 0x4c, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x49, 0x71, 0x71,
	0x64, 0xe4, 0x17, 0x97, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x25, 0xe1, 0x7c, 0x25, 0x29,
	0x2e, 0x36, 0xc7, 0xe4, 0xec, 0xf0, 0xd4, 0x24, 0x21, 0x01, 0x2e, 0xe6, 0xc4, 0xe4, 0x6c, 0xb0,
	0x6e, 0x8e, 0x20, 0x10, 0xd3, 0x28, 0x9e, 0x8b, 0xdf, 0x2f, 0xb5, 0xdc, 0x19, 0x6c, 0xb3, 0x0b,
	0xd8, 0x66, 0x21, 0x1f, 0x2e, 0x7e, 0xf7, 0xd4, 0x12, 0x14, 0x21, 0x39, 0x3d, 0x74, 0x57, 0xa3,
	0xb8, 0x4d, 0x4a, 0x1c, 0x43, 0x1e, 0x62, 0xa1, 0x12, 0x83, 0x13, 0x4b, 0x14, 0x53, 0x41, 0x52,
	0x12, 0x1b, 0xd8, 0x97, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0x76, 0xcf, 0xaf, 0xfe,
	0x00, 0x00, 0x00,
}