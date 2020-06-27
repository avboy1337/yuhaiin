// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

//    BlackIcon          bool   `json:"black_icon"`
//    IsDNSOverHTTPS     bool   `json:"is_dns_over_https"`
//    DNSAcrossProxy     bool   `json:"dns_across_proxy"`
//DnsServer          string `json:"dnsServer"`
//    DnsSubNet          string `json:"dns_sub_net"`
//    Bypass             bool   `json:"bypass"`
//    HttpProxyAddress   string `json:"httpProxyAddress"`
//    Socks5ProxyAddress string `json:"socks5ProxyAddress"`
//    RedirProxyAddress  string `json:"redir_proxy_address"`
//    BypassFile         string `json:"bypassFile"`
//    SsrPath            string `json:"ssrPath"`
//    LocalAddress       string `json:"localAddress"`
//    LocalPort          string `json:"localPort"`
//
//// not use now
//    PythonPath   string `json:"pythonPath"`
//    PidFile      string `json:"pidPath"`
//    LogFile      string `json:"logPath"`
//    FastOpen     bool   `json:"fastOpen"`
//    Works        string `json:"works"`
//    TimeOut      string `json:"timeOut"`
//    HttpProxy    bool   `json:"httpProxy"`
//    UdpTrans     bool   `json:"udpTrans"`
//    AutoStartSsr bool   `json:"autoStartSsr"`
//    IsPrintLog   bool   `json:"is_print_log"`
//    UseLocalDNS  bool   `json:"use_local_dns"`
type Config struct {
	BlackIcon            bool     `protobuf:"varint,1,opt,name=BlackIcon,proto3" json:"BlackIcon,omitempty"`
	DOH                  bool     `protobuf:"varint,2,opt,name=DOH,proto3" json:"DOH,omitempty"`
	DNSProxy             bool     `protobuf:"varint,3,opt,name=DNSProxy,proto3" json:"DNSProxy,omitempty"`
	DNS                  string   `protobuf:"bytes,4,opt,name=DNS,proto3" json:"DNS,omitempty"`
	DNSSubNet            string   `protobuf:"bytes,5,opt,name=DNSSubNet,proto3" json:"DNSSubNet,omitempty"`
	Bypass               bool     `protobuf:"varint,6,opt,name=Bypass,proto3" json:"Bypass,omitempty"`
	BypassFile           string   `protobuf:"bytes,7,opt,name=BypassFile,proto3" json:"BypassFile,omitempty"`
	HTTP                 string   `protobuf:"bytes,8,opt,name=HTTP,proto3" json:"HTTP,omitempty"`
	SOCKS5               string   `protobuf:"bytes,9,opt,name=SOCKS5,proto3" json:"SOCKS5,omitempty"`
	REDIR                string   `protobuf:"bytes,10,opt,name=REDIR,proto3" json:"REDIR,omitempty"`
	SSRPath              string   `protobuf:"bytes,11,opt,name=SSRPath,proto3" json:"SSRPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetBlackIcon() bool {
	if m != nil {
		return m.BlackIcon
	}
	return false
}

func (m *Config) GetDOH() bool {
	if m != nil {
		return m.DOH
	}
	return false
}

func (m *Config) GetDNSProxy() bool {
	if m != nil {
		return m.DNSProxy
	}
	return false
}

func (m *Config) GetDNS() string {
	if m != nil {
		return m.DNS
	}
	return ""
}

func (m *Config) GetDNSSubNet() string {
	if m != nil {
		return m.DNSSubNet
	}
	return ""
}

func (m *Config) GetBypass() bool {
	if m != nil {
		return m.Bypass
	}
	return false
}

func (m *Config) GetBypassFile() string {
	if m != nil {
		return m.BypassFile
	}
	return ""
}

func (m *Config) GetHTTP() string {
	if m != nil {
		return m.HTTP
	}
	return ""
}

func (m *Config) GetSOCKS5() string {
	if m != nil {
		return m.SOCKS5
	}
	return ""
}

func (m *Config) GetREDIR() string {
	if m != nil {
		return m.REDIR
	}
	return ""
}

func (m *Config) GetSSRPath() string {
	if m != nil {
		return m.SSRPath
	}
	return ""
}

type AllGroupOrNode struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Name                 []string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllGroupOrNode) Reset()         { *m = AllGroupOrNode{} }
func (m *AllGroupOrNode) String() string { return proto.CompactTextString(m) }
func (*AllGroupOrNode) ProtoMessage()    {}
func (*AllGroupOrNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *AllGroupOrNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllGroupOrNode.Unmarshal(m, b)
}
func (m *AllGroupOrNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllGroupOrNode.Marshal(b, m, deterministic)
}
func (m *AllGroupOrNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllGroupOrNode.Merge(m, src)
}
func (m *AllGroupOrNode) XXX_Size() int {
	return xxx_messageInfo_AllGroupOrNode.Size(m)
}
func (m *AllGroupOrNode) XXX_DiscardUnknown() {
	xxx_messageInfo_AllGroupOrNode.DiscardUnknown(m)
}

var xxx_messageInfo_AllGroupOrNode proto.InternalMessageInfo

func (m *AllGroupOrNode) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AllGroupOrNode) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

type NowNodeGroupAndNode struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Group                string   `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Node                 string   `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NowNodeGroupAndNode) Reset()         { *m = NowNodeGroupAndNode{} }
func (m *NowNodeGroupAndNode) String() string { return proto.CompactTextString(m) }
func (*NowNodeGroupAndNode) ProtoMessage()    {}
func (*NowNodeGroupAndNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *NowNodeGroupAndNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NowNodeGroupAndNode.Unmarshal(m, b)
}
func (m *NowNodeGroupAndNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NowNodeGroupAndNode.Marshal(b, m, deterministic)
}
func (m *NowNodeGroupAndNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NowNodeGroupAndNode.Merge(m, src)
}
func (m *NowNodeGroupAndNode) XXX_Size() int {
	return xxx_messageInfo_NowNodeGroupAndNode.Size(m)
}
func (m *NowNodeGroupAndNode) XXX_DiscardUnknown() {
	xxx_messageInfo_NowNodeGroupAndNode.DiscardUnknown(m)
}

var xxx_messageInfo_NowNodeGroupAndNode proto.InternalMessageInfo

func (m *NowNodeGroupAndNode) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *NowNodeGroupAndNode) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *NowNodeGroupAndNode) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "config")
	proto.RegisterType((*AllGroupOrNode)(nil), "allGroupOrNode")
	proto.RegisterType((*NowNodeGroupAndNode)(nil), "NowNodeGroupAndNode")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdb, 0x6a, 0xdb, 0x40,
	0x10, 0xf5, 0x25, 0xb6, 0xac, 0x29, 0xbd, 0xb0, 0x35, 0x61, 0x71, 0x43, 0x30, 0x7a, 0xca, 0x93,
	0x02, 0x2d, 0x79, 0x29, 0x85, 0xd0, 0x58, 0xa9, 0x13, 0x5a, 0x64, 0xa3, 0x0d, 0xed, 0xf3, 0xda,
	0x9e, 0x28, 0xa2, 0xb2, 0x76, 0x91, 0x56, 0xb8, 0xfa, 0x93, 0x7e, 0x66, 0x3f, 0xa1, 0xec, 0x48,
	0x89, 0x7b, 0xb1, 0xfd, 0x36, 0xe7, 0xcc, 0x9c, 0xa3, 0xd9, 0x33, 0x02, 0x57, 0xea, 0xc4, 0xd7,
	0xb9, 0x32, 0x6a, 0xf4, 0x26, 0x56, 0x2a, 0x4e, 0xf1, 0x9c, 0xd0, 0xa2, 0xbc, 0x3f, 0xc7, 0xb5,
	0x36, 0x55, 0xd3, 0x3c, 0xfd, 0xb7, 0xb9, 0xc9, 0xa5, 0xd6, 0x98, 0x17, 0x75, 0xdf, 0xfb, 0xd9,
	0x81, 0xfe, 0x52, 0x65, 0xf7, 0x49, 0xcc, 0x4e, 0xc0, 0xbd, 0x4a, 0xe5, 0xf2, 0xfb, 0xed, 0x52,
	0x65, 0xbc, 0x3d, 0x6e, 0x9f, 0x0d, 0xa2, 0x2d, 0xc1, 0x5e, 0x41, 0x37, 0x98, 0xdd, 0xf0, 0x0e,
	0xf1, 0xb6, 0x64, 0x23, 0x18, 0x04, 0xa1, 0x98, 0xe7, 0xea, 0x47, 0xc5, 0xbb, 0x44, 0x3f, 0x61,
	0x9a, 0x0e, 0x05, 0x3f, 0x1a, 0xb7, 0xcf, 0xdc, 0xc8, 0x96, 0xd6, 0x3d, 0x08, 0x85, 0x28, 0x17,
	0x21, 0x1a, 0xde, 0x23, 0x7e, 0x4b, 0xb0, 0x63, 0xe8, 0x5f, 0x55, 0x5a, 0x16, 0x05, 0xef, 0x93,
	0x53, 0x83, 0xd8, 0x29, 0x40, 0x5d, 0x7d, 0x4a, 0x52, 0xe4, 0x0e, 0xc9, 0xfe, 0x60, 0x18, 0x83,
	0xa3, 0x9b, 0xbb, 0xbb, 0x39, 0x1f, 0x50, 0x87, 0x6a, 0xeb, 0x25, 0x66, 0x93, 0xcf, 0xe2, 0x82,
	0xbb, 0xc4, 0x36, 0x88, 0x0d, 0xa1, 0x17, 0x5d, 0x07, 0xb7, 0x11, 0x07, 0xa2, 0x6b, 0xc0, 0x38,
	0x38, 0x42, 0x44, 0x73, 0x69, 0x1e, 0xf8, 0x33, 0xe2, 0x1f, 0xa1, 0xf7, 0x01, 0x5e, 0xc8, 0x34,
	0x9d, 0xe6, 0xaa, 0xd4, 0xb3, 0x3c, 0x54, 0x2b, 0xb4, 0xce, 0x85, 0x91, 0xa6, 0x2c, 0x28, 0x9e,
	0x5e, 0xd4, 0x20, 0xbb, 0x45, 0x26, 0xd7, 0xc8, 0x3b, 0xe3, 0xae, 0xdd, 0xc2, 0xd6, 0xde, 0x37,
	0x78, 0x1d, 0xaa, 0x8d, 0x95, 0x91, 0xc3, 0xc7, 0x6c, 0x75, 0xd0, 0x62, 0x08, 0xbd, 0xd8, 0xce,
	0x51, 0xc0, 0x6e, 0x54, 0x03, 0x32, 0x56, 0x2b, 0xa4, 0x78, 0xad, 0xb1, 0x5a, 0xe1, 0xdb, 0x5f,
	0x1d, 0xe8, 0x4a, 0x9d, 0x30, 0x1f, 0x5c, 0x81, 0x66, 0x52, 0xdf, 0xce, 0xf1, 0xeb, 0x23, 0x8e,
	0x8e, 0xfd, 0xfa, 0xe0, 0xfe, 0xe3, 0xc1, 0xfd, 0x6b, 0xfb, 0x37, 0x78, 0x2d, 0x76, 0x01, 0x83,
	0x29, 0x1a, 0x5a, 0x86, 0xed, 0x99, 0x1a, 0xbd, 0xf4, 0xff, 0x7e, 0xb1, 0xd7, 0x62, 0xef, 0xc1,
	0x99, 0xa2, 0xa1, 0xdd, 0x4f, 0xfe, 0x53, 0x09, 0x93, 0x27, 0x59, 0xfc, 0x55, 0xa6, 0x25, 0xee,
	0xd2, 0x06, 0xc0, 0x48, 0xbb, 0x79, 0x8a, 0x40, 0xae, 0x71, 0xef, 0xc7, 0x87, 0xfe, 0x8e, 0xc0,
	0xbc, 0x16, 0xbb, 0x84, 0xe7, 0x93, 0x07, 0x99, 0xc5, 0xd8, 0xb4, 0xd9, 0xce, 0xc1, 0x03, 0x2f,
	0xbf, 0x04, 0xe7, 0x8b, 0x34, 0x98, 0x2d, 0xab, 0x3d, 0xd2, 0x83, 0x0f, 0xf3, 0x5a, 0x8b, 0x3e,
	0xf1, 0xef, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x90, 0xe3, 0x85, 0x75, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiClient interface {
	SetConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*empty.Empty, error)
	GetGroup(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllGroupOrNode, error)
	GetNode(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*AllGroupOrNode, error)
	GetNowGroupAndName(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NowNodeGroupAndNode, error)
	ChangeNowNode(ctx context.Context, in *NowNodeGroupAndNode, opts ...grpc.CallOption) (*empty.Empty, error)
	Latency(ctx context.Context, in *NowNodeGroupAndNode, opts ...grpc.CallOption) (*wrappers.StringValue, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) SetConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetGroup(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllGroupOrNode, error) {
	out := new(AllGroupOrNode)
	err := c.cc.Invoke(ctx, "/api/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetNode(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*AllGroupOrNode, error) {
	out := new(AllGroupOrNode)
	err := c.cc.Invoke(ctx, "/api/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetNowGroupAndName(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NowNodeGroupAndNode, error) {
	out := new(NowNodeGroupAndNode)
	err := c.cc.Invoke(ctx, "/api/GetNowGroupAndName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ChangeNowNode(ctx context.Context, in *NowNodeGroupAndNode, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api/ChangeNowNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Latency(ctx context.Context, in *NowNodeGroupAndNode, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/api/Latency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
type ApiServer interface {
	SetConfig(context.Context, *Config) (*empty.Empty, error)
	GetGroup(context.Context, *empty.Empty) (*AllGroupOrNode, error)
	GetNode(context.Context, *wrappers.StringValue) (*AllGroupOrNode, error)
	GetNowGroupAndName(context.Context, *empty.Empty) (*NowNodeGroupAndNode, error)
	ChangeNowNode(context.Context, *NowNodeGroupAndNode) (*empty.Empty, error)
	Latency(context.Context, *NowNodeGroupAndNode) (*wrappers.StringValue, error)
}

// UnimplementedApiServer can be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (*UnimplementedApiServer) SetConfig(ctx context.Context, req *Config) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (*UnimplementedApiServer) GetGroup(ctx context.Context, req *empty.Empty) (*AllGroupOrNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (*UnimplementedApiServer) GetNode(ctx context.Context, req *wrappers.StringValue) (*AllGroupOrNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (*UnimplementedApiServer) GetNowGroupAndName(ctx context.Context, req *empty.Empty) (*NowNodeGroupAndNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNowGroupAndName not implemented")
}
func (*UnimplementedApiServer) ChangeNowNode(ctx context.Context, req *NowNodeGroupAndNode) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNowNode not implemented")
}
func (*UnimplementedApiServer) Latency(ctx context.Context, req *NowNodeGroupAndNode) (*wrappers.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Latency not implemented")
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).SetConfig(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetGroup(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetNode(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetNowGroupAndName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetNowGroupAndName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/GetNowGroupAndName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetNowGroupAndName(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ChangeNowNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NowNodeGroupAndNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ChangeNowNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/ChangeNowNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ChangeNowNode(ctx, req.(*NowNodeGroupAndNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Latency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NowNodeGroupAndNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Latency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api/Latency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Latency(ctx, req.(*NowNodeGroupAndNode))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConfig",
			Handler:    _Api_SetConfig_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _Api_GetGroup_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _Api_GetNode_Handler,
		},
		{
			MethodName: "GetNowGroupAndName",
			Handler:    _Api_GetNowGroupAndName_Handler,
		},
		{
			MethodName: "ChangeNowNode",
			Handler:    _Api_ChangeNowNode_Handler,
		},
		{
			MethodName: "Latency",
			Handler:    _Api_Latency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
