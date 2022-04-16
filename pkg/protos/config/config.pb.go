// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: config/config.proto

package config

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DnsDnsType int32

const (
	Dns_reserve DnsDnsType = 0
	Dns_udp     DnsDnsType = 1
	Dns_tcp     DnsDnsType = 2
	Dns_doh     DnsDnsType = 3
	Dns_dot     DnsDnsType = 4
)

// Enum value maps for DnsDnsType.
var (
	DnsDnsType_name = map[int32]string{
		0: "reserve",
		1: "udp",
		2: "tcp",
		3: "doh",
		4: "dot",
	}
	DnsDnsType_value = map[string]int32{
		"reserve": 0,
		"udp":     1,
		"tcp":     2,
		"doh":     3,
		"dot":     4,
	}
)

func (x DnsDnsType) Enum() *DnsDnsType {
	p := new(DnsDnsType)
	*p = x
	return p
}

func (x DnsDnsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DnsDnsType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_config_proto_enumTypes[0].Descriptor()
}

func (DnsDnsType) Type() protoreflect.EnumType {
	return &file_config_config_proto_enumTypes[0]
}

func (x DnsDnsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DnsDnsType.Descriptor instead.
func (DnsDnsType) EnumDescriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{4, 0}
}

type Setting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemProxy *SystemProxy `protobuf:"bytes,1,opt,name=system_proxy,proto3" json:"system_proxy,omitempty"`
	Bypass      *Bypass      `protobuf:"bytes,2,opt,name=bypass,proto3" json:"bypass,omitempty"`
	Dns         *DnsSetting  `protobuf:"bytes,4,opt,name=dns,proto3" json:"dns,omitempty"`
	Server      *Server      `protobuf:"bytes,5,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *Setting) Reset() {
	*x = Setting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *Setting) GetSystemProxy() *SystemProxy {
	if x != nil {
		return x.SystemProxy
	}
	return nil
}

func (x *Setting) GetBypass() *Bypass {
	if x != nil {
		return x.Bypass
	}
	return nil
}

func (x *Setting) GetDns() *DnsSetting {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *Setting) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type DnsSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remote *Dns `protobuf:"bytes,1,opt,name=remote,proto3" json:"remote,omitempty"`
	Local  *Dns `protobuf:"bytes,2,opt,name=local,proto3" json:"local,omitempty"`
}

func (x *DnsSetting) Reset() {
	*x = DnsSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsSetting) ProtoMessage() {}

func (x *DnsSetting) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsSetting.ProtoReflect.Descriptor instead.
func (*DnsSetting) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *DnsSetting) GetRemote() *Dns {
	if x != nil {
		return x.Remote
	}
	return nil
}

func (x *DnsSetting) GetLocal() *Dns {
	if x != nil {
		return x.Local
	}
	return nil
}

type SystemProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http   bool `protobuf:"varint,2,opt,name=http,proto3" json:"http,omitempty"`
	Socks5 bool `protobuf:"varint,3,opt,name=socks5,proto3" json:"socks5,omitempty"`
}

func (x *SystemProxy) Reset() {
	*x = SystemProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemProxy) ProtoMessage() {}

func (x *SystemProxy) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemProxy.ProtoReflect.Descriptor instead.
func (*SystemProxy) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{2}
}

func (x *SystemProxy) GetHttp() bool {
	if x != nil {
		return x.Http
	}
	return false
}

func (x *SystemProxy) GetSocks5() bool {
	if x != nil {
		return x.Socks5
	}
	return false
}

type Bypass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled    bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	BypassFile string `protobuf:"bytes,2,opt,name=bypass_file,proto3" json:"bypass_file,omitempty"`
}

func (x *Bypass) Reset() {
	*x = Bypass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bypass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bypass) ProtoMessage() {}

func (x *Bypass) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bypass.ProtoReflect.Descriptor instead.
func (*Bypass) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{3}
}

func (x *Bypass) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Bypass) GetBypassFile() string {
	if x != nil {
		return x.BypassFile
	}
	return ""
}

type Dns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host   string     `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Type   DnsDnsType `protobuf:"varint,5,opt,name=type,proto3,enum=yuhaiin.config.DnsDnsType" json:"type,omitempty"`
	Proxy  bool       `protobuf:"varint,3,opt,name=proxy,proto3" json:"proxy,omitempty"`
	Subnet string     `protobuf:"bytes,4,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *Dns) Reset() {
	*x = Dns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dns) ProtoMessage() {}

func (x *Dns) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dns.ProtoReflect.Descriptor instead.
func (*Dns) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{4}
}

func (x *Dns) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Dns) GetType() DnsDnsType {
	if x != nil {
		return x.Type
	}
	return Dns_reserve
}

func (x *Dns) GetProxy() bool {
	if x != nil {
		return x.Proxy
	}
	return false
}

func (x *Dns) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers map[string]*ServerProtocol `protobuf:"bytes,5,rep,name=servers,proto3" json:"servers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{5}
}

func (x *Server) GetServers() map[string]*ServerProtocol {
	if x != nil {
		return x.Servers
	}
	return nil
}

type Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Http) Reset() {
	*x = Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{6}
}

func (x *Http) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Http) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Http) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Socks5 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Socks5) Reset() {
	*x = Socks5{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Socks5) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Socks5) ProtoMessage() {}

func (x *Socks5) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Socks5.ProtoReflect.Descriptor instead.
func (*Socks5) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{7}
}

func (x *Socks5) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Socks5) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Socks5) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Redir struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Redir) Reset() {
	*x = Redir{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redir) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redir) ProtoMessage() {}

func (x *Redir) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redir.ProtoReflect.Descriptor instead.
func (*Redir) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{8}
}

func (x *Redir) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type ServerProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// Types that are assignable to Protocol:
	//	*ServerProtocol_Http
	//	*ServerProtocol_Socks5
	//	*ServerProtocol_Redir
	Protocol isServerProtocol_Protocol `protobuf_oneof:"protocol"`
}

func (x *ServerProtocol) Reset() {
	*x = ServerProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerProtocol) ProtoMessage() {}

func (x *ServerProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerProtocol.ProtoReflect.Descriptor instead.
func (*ServerProtocol) Descriptor() ([]byte, []int) {
	return file_config_config_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ServerProtocol) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerProtocol) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (m *ServerProtocol) GetProtocol() isServerProtocol_Protocol {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (x *ServerProtocol) GetHttp() *Http {
	if x, ok := x.GetProtocol().(*ServerProtocol_Http); ok {
		return x.Http
	}
	return nil
}

func (x *ServerProtocol) GetSocks5() *Socks5 {
	if x, ok := x.GetProtocol().(*ServerProtocol_Socks5); ok {
		return x.Socks5
	}
	return nil
}

func (x *ServerProtocol) GetRedir() *Redir {
	if x, ok := x.GetProtocol().(*ServerProtocol_Redir); ok {
		return x.Redir
	}
	return nil
}

type isServerProtocol_Protocol interface {
	isServerProtocol_Protocol()
}

type ServerProtocol_Http struct {
	Http *Http `protobuf:"bytes,3,opt,name=http,proto3,oneof"`
}

type ServerProtocol_Socks5 struct {
	Socks5 *Socks5 `protobuf:"bytes,4,opt,name=socks5,proto3,oneof"`
}

type ServerProtocol_Redir struct {
	Redir *Redir `protobuf:"bytes,5,opt,name=redir,proto3,oneof"`
}

func (*ServerProtocol_Http) isServerProtocol_Protocol() {}

func (*ServerProtocol_Socks5) isServerProtocol_Protocol() {}

func (*ServerProtocol_Redir) isServerProtocol_Protocol() {}

var File_config_config_proto protoreflect.FileDescriptor

var file_config_config_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x40,
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x12, 0x2e, 0x0a, 0x06, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x52, 0x06, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73,
	0x12, 0x2d, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x64,
	0x6e, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x12,
	0x2e, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22,
	0x65, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2b,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x64, 0x6e, 0x73, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x75, 0x68,
	0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x52,
	0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0x3a, 0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x63, 0x6b, 0x73, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x6f, 0x63, 0x6b,
	0x73, 0x35, 0x22, 0x44, 0x0a, 0x06, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x79, 0x70,
	0x61, 0x73, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x03, 0x64, 0x6e, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x22, 0x3b, 0x0a, 0x08, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x75, 0x64, 0x70, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x74, 0x63, 0x70, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x64, 0x6f, 0x68, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x6f, 0x74, 0x10,
	0x04, 0x22, 0xf2, 0x02, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x1a, 0xcb, 0x01, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x2a, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x48, 0x00, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x30, 0x0a, 0x06,
	0x73, 0x6f, 0x63, 0x6b, 0x73, 0x35, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x79,
	0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x6f,
	0x63, 0x6b, 0x73, 0x35, 0x48, 0x00, 0x52, 0x06, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x35, 0x12, 0x2d,
	0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x72, 0x42, 0x0a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x5b, 0x0a, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x75, 0x68,
	0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x54, 0x0a, 0x06, 0x73, 0x6f,
	0x63, 0x6b, 0x73, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x1b, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x32, 0x7e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x6f, 0x12, 0x37, 0x0a, 0x04, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x79, 0x75,
	0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x61, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x79,
	0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x75, 0x74,
	0x6f, 0x72, 0x75, 0x66, 0x61, 0x2f, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_config_proto_rawDescOnce sync.Once
	file_config_config_proto_rawDescData = file_config_config_proto_rawDesc
)

func file_config_config_proto_rawDescGZIP() []byte {
	file_config_config_proto_rawDescOnce.Do(func() {
		file_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_config_proto_rawDescData)
	})
	return file_config_config_proto_rawDescData
}

var file_config_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_config_config_proto_goTypes = []interface{}{
	(DnsDnsType)(0),        // 0: yuhaiin.config.dns.dns_type
	(*Setting)(nil),        // 1: yuhaiin.config.setting
	(*DnsSetting)(nil),     // 2: yuhaiin.config.dns_setting
	(*SystemProxy)(nil),    // 3: yuhaiin.config.system_proxy
	(*Bypass)(nil),         // 4: yuhaiin.config.bypass
	(*Dns)(nil),            // 5: yuhaiin.config.dns
	(*Server)(nil),         // 6: yuhaiin.config.server
	(*Http)(nil),           // 7: yuhaiin.config.http
	(*Socks5)(nil),         // 8: yuhaiin.config.socks5
	(*Redir)(nil),          // 9: yuhaiin.config.redir
	(*ServerProtocol)(nil), // 10: yuhaiin.config.server.protocol
	nil,                    // 11: yuhaiin.config.server.ServersEntry
	(*emptypb.Empty)(nil),  // 12: google.protobuf.Empty
}
var file_config_config_proto_depIdxs = []int32{
	3,  // 0: yuhaiin.config.setting.system_proxy:type_name -> yuhaiin.config.system_proxy
	4,  // 1: yuhaiin.config.setting.bypass:type_name -> yuhaiin.config.bypass
	2,  // 2: yuhaiin.config.setting.dns:type_name -> yuhaiin.config.dns_setting
	6,  // 3: yuhaiin.config.setting.server:type_name -> yuhaiin.config.server
	5,  // 4: yuhaiin.config.dns_setting.remote:type_name -> yuhaiin.config.dns
	5,  // 5: yuhaiin.config.dns_setting.local:type_name -> yuhaiin.config.dns
	0,  // 6: yuhaiin.config.dns.type:type_name -> yuhaiin.config.dns.dns_type
	11, // 7: yuhaiin.config.server.servers:type_name -> yuhaiin.config.server.ServersEntry
	7,  // 8: yuhaiin.config.server.protocol.http:type_name -> yuhaiin.config.http
	8,  // 9: yuhaiin.config.server.protocol.socks5:type_name -> yuhaiin.config.socks5
	9,  // 10: yuhaiin.config.server.protocol.redir:type_name -> yuhaiin.config.redir
	10, // 11: yuhaiin.config.server.ServersEntry.value:type_name -> yuhaiin.config.server.protocol
	12, // 12: yuhaiin.config.config_dao.load:input_type -> google.protobuf.Empty
	1,  // 13: yuhaiin.config.config_dao.save:input_type -> yuhaiin.config.setting
	1,  // 14: yuhaiin.config.config_dao.load:output_type -> yuhaiin.config.setting
	12, // 15: yuhaiin.config.config_dao.save:output_type -> google.protobuf.Empty
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_config_config_proto_init() }
func file_config_config_proto_init() {
	if File_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Setting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsSetting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemProxy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bypass); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dns); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Socks5); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redir); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerProtocol); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_config_config_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*ServerProtocol_Http)(nil),
		(*ServerProtocol_Socks5)(nil),
		(*ServerProtocol_Redir)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_config_proto_goTypes,
		DependencyIndexes: file_config_config_proto_depIdxs,
		EnumInfos:         file_config_config_proto_enumTypes,
		MessageInfos:      file_config_config_proto_msgTypes,
	}.Build()
	File_config_config_proto = out.File
	file_config_config_proto_rawDesc = nil
	file_config_config_proto_goTypes = nil
	file_config_config_proto_depIdxs = nil
}
