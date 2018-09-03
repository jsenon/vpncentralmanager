// Code generated by protoc-gen-go. DO NOT EDIT.
// source: retrieveconfig.proto

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

// id could be vpn server or client config id
// type will be choose in order to retrieve server or client configuration
type AllConfigFileReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllConfigFileReq) Reset()         { *m = AllConfigFileReq{} }
func (m *AllConfigFileReq) String() string { return proto.CompactTextString(m) }
func (*AllConfigFileReq) ProtoMessage()    {}
func (*AllConfigFileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieveconfig_8fe25e6281480e7a, []int{0}
}
func (m *AllConfigFileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllConfigFileReq.Unmarshal(m, b)
}
func (m *AllConfigFileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllConfigFileReq.Marshal(b, m, deterministic)
}
func (dst *AllConfigFileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllConfigFileReq.Merge(dst, src)
}
func (m *AllConfigFileReq) XXX_Size() int {
	return xxx_messageInfo_AllConfigFileReq.Size(m)
}
func (m *AllConfigFileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AllConfigFileReq.DiscardUnknown(m)
}

var xxx_messageInfo_AllConfigFileReq proto.InternalMessageInfo

func (m *AllConfigFileReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AllConfigFileReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// All configuration
// could be client or vpn server configuration
type AllConfigFileResp struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllConfigFileResp) Reset()         { *m = AllConfigFileResp{} }
func (m *AllConfigFileResp) String() string { return proto.CompactTextString(m) }
func (*AllConfigFileResp) ProtoMessage()    {}
func (*AllConfigFileResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieveconfig_8fe25e6281480e7a, []int{1}
}
func (m *AllConfigFileResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllConfigFileResp.Unmarshal(m, b)
}
func (m *AllConfigFileResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllConfigFileResp.Marshal(b, m, deterministic)
}
func (dst *AllConfigFileResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllConfigFileResp.Merge(dst, src)
}
func (m *AllConfigFileResp) XXX_Size() int {
	return xxx_messageInfo_AllConfigFileResp.Size(m)
}
func (m *AllConfigFileResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AllConfigFileResp.DiscardUnknown(m)
}

var xxx_messageInfo_AllConfigFileResp proto.InternalMessageInfo

func (m *AllConfigFileResp) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type Item struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Addressvpn           string   `protobuf:"bytes,3,opt,name=addressvpn,proto3" json:"addressvpn,omitempty"`
	Addresspub           string   `protobuf:"bytes,4,opt,name=addresspub,proto3" json:"addresspub,omitempty"`
	Publikey             string   `protobuf:"bytes,5,opt,name=publikey,proto3" json:"publikey,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_retrieveconfig_8fe25e6281480e7a, []int{2}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetAddressvpn() string {
	if m != nil {
		return m.Addressvpn
	}
	return ""
}

func (m *Item) GetAddresspub() string {
	if m != nil {
		return m.Addresspub
	}
	return ""
}

func (m *Item) GetPublikey() string {
	if m != nil {
		return m.Publikey
	}
	return ""
}

func (m *Item) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*AllConfigFileReq)(nil), "retrieveconfig.AllConfigFileReq")
	proto.RegisterType((*AllConfigFileResp)(nil), "retrieveconfig.AllConfigFileResp")
	proto.RegisterType((*Item)(nil), "retrieveconfig.Item")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RetrieveConfigClient is the client API for RetrieveConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RetrieveConfigClient interface {
	GetAllConfig(ctx context.Context, in *AllConfigFileReq, opts ...grpc.CallOption) (*AllConfigFileResp, error)
}

type retrieveConfigClient struct {
	cc *grpc.ClientConn
}

func NewRetrieveConfigClient(cc *grpc.ClientConn) RetrieveConfigClient {
	return &retrieveConfigClient{cc}
}

func (c *retrieveConfigClient) GetAllConfig(ctx context.Context, in *AllConfigFileReq, opts ...grpc.CallOption) (*AllConfigFileResp, error) {
	out := new(AllConfigFileResp)
	err := c.cc.Invoke(ctx, "/retrieveconfig.RetrieveConfig/GetAllConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrieveConfigServer is the server API for RetrieveConfig service.
type RetrieveConfigServer interface {
	GetAllConfig(context.Context, *AllConfigFileReq) (*AllConfigFileResp, error)
}

func RegisterRetrieveConfigServer(s *grpc.Server, srv RetrieveConfigServer) {
	s.RegisterService(&_RetrieveConfig_serviceDesc, srv)
}

func _RetrieveConfig_GetAllConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllConfigFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveConfigServer).GetAllConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/retrieveconfig.RetrieveConfig/GetAllConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveConfigServer).GetAllConfig(ctx, req.(*AllConfigFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RetrieveConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "retrieveconfig.RetrieveConfig",
	HandlerType: (*RetrieveConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllConfig",
			Handler:    _RetrieveConfig_GetAllConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "retrieveconfig.proto",
}

func init() {
	proto.RegisterFile("retrieveconfig.proto", fileDescriptor_retrieveconfig_8fe25e6281480e7a)
}

var fileDescriptor_retrieveconfig_8fe25e6281480e7a = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x9a, 0x06, 0x1d, 0x25, 0xe8, 0x50, 0x64, 0xe9, 0x41, 0x62, 0x4e, 0xc5, 0x43,
	0x0f, 0x15, 0xbc, 0x8a, 0x0a, 0x8a, 0xd7, 0x80, 0x17, 0x6f, 0x59, 0x33, 0x96, 0xc5, 0xfc, 0x19,
	0xb3, 0x93, 0x42, 0xbf, 0x8c, 0x9f, 0x55, 0xba, 0x2d, 0xa1, 0x5d, 0x04, 0x6f, 0xf3, 0xde, 0x6f,
	0x96, 0x99, 0xb7, 0x03, 0x93, 0x8e, 0xa4, 0x33, 0xb4, 0xa2, 0x8f, 0xb6, 0xf9, 0x34, 0xcb, 0x39,
	0x77, 0xad, 0xb4, 0x98, 0x1c, 0xba, 0xd9, 0x1d, 0x9c, 0x3f, 0x54, 0xd5, 0x93, 0x13, 0xcf, 0xa6,
	0xa2, 0x9c, 0xbe, 0x31, 0x81, 0xd0, 0x94, 0x2a, 0x48, 0x83, 0xd9, 0x49, 0x1e, 0x9a, 0x12, 0x11,
	0x22, 0x59, 0x33, 0xa9, 0xd0, 0x39, 0xae, 0xce, 0xee, 0xe1, 0xc2, 0x7b, 0x67, 0x19, 0x6f, 0x60,
	0x6c, 0x84, 0x6a, 0xab, 0x82, 0x74, 0x34, 0x3b, 0x5d, 0x4c, 0xe6, 0xde, 0x0a, 0xaf, 0x42, 0x75,
	0xbe, 0x6d, 0xc9, 0x7e, 0x02, 0x88, 0x36, 0xfa, 0xaf, 0x69, 0x4d, 0x51, 0x0f, 0xd3, 0x36, 0x35,
	0x5e, 0x01, 0x14, 0x65, 0xd9, 0x91, 0xb5, 0x2b, 0x6e, 0xd4, 0xc8, 0x91, 0x3d, 0x67, 0x8f, 0x73,
	0xaf, 0x55, 0x74, 0xc0, 0xb9, 0xd7, 0x38, 0x85, 0x63, 0xee, 0x75, 0x65, 0xbe, 0x68, 0xad, 0xc6,
	0x8e, 0x0e, 0x1a, 0x2f, 0x21, 0xb6, 0x52, 0x48, 0x6f, 0x55, 0xec, 0xc8, 0x4e, 0x2d, 0x96, 0x90,
	0xe4, 0xbb, 0xf5, 0xb7, 0x31, 0xf1, 0x0d, 0xce, 0x5e, 0x48, 0x86, 0xd8, 0x98, 0xfa, 0xf9, 0xfc,
	0x9f, 0x9c, 0x5e, 0xff, 0xd3, 0x61, 0x39, 0x3b, 0x7a, 0x8c, 0xde, 0x43, 0xd6, 0x3a, 0x76, 0xf7,
	0xb9, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x19, 0xf4, 0x28, 0x28, 0xb7, 0x01, 0x00, 0x00,
}