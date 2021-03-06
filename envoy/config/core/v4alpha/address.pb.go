// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v4alpha/address.proto

package envoy_config_core_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}

var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}

func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b7129990fbcdefb, []int{1, 0}
}

type Pipe struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Mode                 uint32   `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pipe) Reset()         { *m = Pipe{} }
func (m *Pipe) String() string { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()    {}
func (*Pipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7129990fbcdefb, []int{0}
}

func (m *Pipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipe.Unmarshal(m, b)
}
func (m *Pipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipe.Marshal(b, m, deterministic)
}
func (m *Pipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipe.Merge(m, src)
}
func (m *Pipe) XXX_Size() int {
	return xxx_messageInfo_Pipe.Size(m)
}
func (m *Pipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipe.DiscardUnknown(m)
}

var xxx_messageInfo_Pipe proto.InternalMessageInfo

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Pipe) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=envoy.config.core.v4alpha.SocketAddress_Protocol" json:"protocol,omitempty"`
	Address  string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier        isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	ResolverName         string                        `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	Ipv4Compat           bool                          `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat,proto3" json:"ipv4_compat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SocketAddress) Reset()         { *m = SocketAddress{} }
func (m *SocketAddress) String() string { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()    {}
func (*SocketAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7129990fbcdefb, []int{1}
}

func (m *SocketAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketAddress.Unmarshal(m, b)
}
func (m *SocketAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketAddress.Marshal(b, m, deterministic)
}
func (m *SocketAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketAddress.Merge(m, src)
}
func (m *SocketAddress) XXX_Size() int {
	return xxx_messageInfo_SocketAddress.Size(m)
}
func (m *SocketAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketAddress.DiscardUnknown(m)
}

var xxx_messageInfo_SocketAddress proto.InternalMessageInfo

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,proto3,oneof"`
}

type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,proto3,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}

func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

func (m *SocketAddress) GetIpv4Compat() bool {
	if m != nil {
		return m.Ipv4Compat
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

type TcpKeepalive struct {
	KeepaliveProbes      *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	KeepaliveTime        *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3" json:"keepalive_time,omitempty"`
	KeepaliveInterval    *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3" json:"keepalive_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TcpKeepalive) Reset()         { *m = TcpKeepalive{} }
func (m *TcpKeepalive) String() string { return proto.CompactTextString(m) }
func (*TcpKeepalive) ProtoMessage()    {}
func (*TcpKeepalive) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7129990fbcdefb, []int{2}
}

func (m *TcpKeepalive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpKeepalive.Unmarshal(m, b)
}
func (m *TcpKeepalive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpKeepalive.Marshal(b, m, deterministic)
}
func (m *TcpKeepalive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpKeepalive.Merge(m, src)
}
func (m *TcpKeepalive) XXX_Size() int {
	return xxx_messageInfo_TcpKeepalive.Size(m)
}
func (m *TcpKeepalive) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpKeepalive.DiscardUnknown(m)
}

var xxx_messageInfo_TcpKeepalive proto.InternalMessageInfo

func (m *TcpKeepalive) GetKeepaliveProbes() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveProbes
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveTime() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveTime
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveInterval() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveInterval
	}
	return nil
}

type BindConfig struct {
	SourceAddress        *SocketAddress      `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	Freebind             *wrappers.BoolValue `protobuf:"bytes,2,opt,name=freebind,proto3" json:"freebind,omitempty"`
	SocketOptions        []*SocketOption     `protobuf:"bytes,3,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BindConfig) Reset()         { *m = BindConfig{} }
func (m *BindConfig) String() string { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()    {}
func (*BindConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7129990fbcdefb, []int{3}
}

func (m *BindConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindConfig.Unmarshal(m, b)
}
func (m *BindConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindConfig.Marshal(b, m, deterministic)
}
func (m *BindConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindConfig.Merge(m, src)
}
func (m *BindConfig) XXX_Size() int {
	return xxx_messageInfo_BindConfig.Size(m)
}
func (m *BindConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BindConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BindConfig proto.InternalMessageInfo

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *BindConfig) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *BindConfig) GetSocketOptions() []*SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address              isAddress_Address `protobuf_oneof:"address"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7129990fbcdefb, []int{4}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3,oneof"`
}

type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,proto3,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}

func (*Address_Pipe) isAddress_Address() {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Address) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

type CidrRange struct {
	AddressPrefix        string                `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix,proto3" json:"address_prefix,omitempty"`
	PrefixLen            *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CidrRange) Reset()         { *m = CidrRange{} }
func (m *CidrRange) String() string { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()    {}
func (*CidrRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7129990fbcdefb, []int{5}
}

func (m *CidrRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidrRange.Unmarshal(m, b)
}
func (m *CidrRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidrRange.Marshal(b, m, deterministic)
}
func (m *CidrRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidrRange.Merge(m, src)
}
func (m *CidrRange) XXX_Size() int {
	return xxx_messageInfo_CidrRange.Size(m)
}
func (m *CidrRange) XXX_DiscardUnknown() {
	xxx_messageInfo_CidrRange.DiscardUnknown(m)
}

var xxx_messageInfo_CidrRange proto.InternalMessageInfo

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.core.v4alpha.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
	proto.RegisterType((*Pipe)(nil), "envoy.config.core.v4alpha.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.config.core.v4alpha.SocketAddress")
	proto.RegisterType((*TcpKeepalive)(nil), "envoy.config.core.v4alpha.TcpKeepalive")
	proto.RegisterType((*BindConfig)(nil), "envoy.config.core.v4alpha.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.config.core.v4alpha.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.config.core.v4alpha.CidrRange")
}

func init() {
	proto.RegisterFile("envoy/config/core/v4alpha/address.proto", fileDescriptor_9b7129990fbcdefb)
}

var fileDescriptor_9b7129990fbcdefb = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0xae, 0xe3, 0x6c, 0x9b, 0xbc, 0x36, 0x21, 0x8c, 0x84, 0x30, 0xdd, 0x90, 0x26, 0x29, 0xd0,
	0x50, 0x09, 0x5b, 0xa4, 0x0b, 0x87, 0xdc, 0xea, 0x80, 0xe8, 0x6a, 0xd1, 0x62, 0x4c, 0x17, 0xc4,
	0xc9, 0x9a, 0xd8, 0x93, 0xec, 0x68, 0x1d, 0xcf, 0x68, 0xec, 0x98, 0xdd, 0xdb, 0x8a, 0x13, 0xe2,
	0xca, 0x01, 0x89, 0x7f, 0xc0, 0x5f, 0xe0, 0x8e, 0xb4, 0x57, 0xfe, 0x0d, 0xea, 0x81, 0xa2, 0x19,
	0x7b, 0x9c, 0x2e, 0x9b, 0xd2, 0x6a, 0x6f, 0xf6, 0x7b, 0xdf, 0xfb, 0xe6, 0x7b, 0xdf, 0x7c, 0x1a,
	0x38, 0x22, 0x49, 0xce, 0x9e, 0x39, 0x21, 0x4b, 0xe6, 0x74, 0xe1, 0x84, 0x4c, 0x10, 0x27, 0xbf,
	0x87, 0x63, 0xfe, 0x18, 0x3b, 0x38, 0x8a, 0x04, 0x49, 0x53, 0x9b, 0x0b, 0x96, 0x31, 0xf4, 0x8e,
	0x02, 0xda, 0x05, 0xd0, 0x96, 0x40, 0xbb, 0x04, 0xee, 0x7f, 0x74, 0x3d, 0x47, 0xca, 0xc2, 0x27,
	0x24, 0x0b, 0x18, 0xcf, 0x28, 0x4b, 0x0a, 0xa6, 0xfd, 0xde, 0x82, 0xb1, 0x45, 0x4c, 0x1c, 0xf5,
	0x37, 0x5b, 0xcd, 0x9d, 0x1f, 0x04, 0xe6, 0x9c, 0x88, 0xf2, 0xa4, 0xfd, 0x77, 0x57, 0x11, 0xc7,
	0x0e, 0x4e, 0x12, 0x96, 0x61, 0x39, 0x96, 0x3a, 0x69, 0x86, 0xb3, 0x95, 0x6e, 0x0f, 0x5e, 0x69,
	0xe7, 0x44, 0xa4, 0x94, 0x25, 0x34, 0x59, 0x94, 0x90, 0xb7, 0x73, 0x1c, 0xd3, 0x08, 0x67, 0xc4,
	0xd1, 0x1f, 0x45, 0x63, 0x18, 0x42, 0xdd, 0xa3, 0x9c, 0xa0, 0xbb, 0x50, 0xe7, 0x38, 0x7b, 0x6c,
	0x19, 0x7d, 0x63, 0xd4, 0x74, 0x77, 0x2e, 0xdc, 0xba, 0xa8, 0xf5, 0x0d, 0x5f, 0x15, 0x51, 0x17,
	0xea, 0x4b, 0x16, 0x11, 0xab, 0xd6, 0x37, 0x46, 0x2d, 0xb7, 0x71, 0xe1, 0xde, 0x39, 0x36, 0xad,
	0x4b, 0xd3, 0x57, 0xd5, 0x49, 0xff, 0xb7, 0x3f, 0x7f, 0xea, 0xdd, 0x85, 0x4d, 0x76, 0x9c, 0xd8,
	0x92, 0x7c, 0xf8, 0x4f, 0x0d, 0x5a, 0xdf, 0xa8, 0xbd, 0x4f, 0x0b, 0x07, 0xd1, 0x77, 0xd0, 0x50,
	0xe7, 0x87, 0x2c, 0x56, 0x47, 0xb6, 0xc7, 0x1f, 0xdb, 0xd7, 0xda, 0x69, 0xbf, 0x34, 0x6b, 0x7b,
	0xe5, 0xa0, 0x12, 0xf2, 0xa3, 0x51, 0xeb, 0x18, 0x7e, 0x45, 0x86, 0x06, 0xb0, 0x53, 0xde, 0x92,
	0x52, 0x7b, 0x65, 0x15, 0x5d, 0x47, 0xc7, 0x00, 0x9c, 0x89, 0x2c, 0xc8, 0x71, 0xbc, 0x22, 0x96,
	0xa9, 0x76, 0x6a, 0x5e, 0xb8, 0xdb, 0xc7, 0x75, 0xeb, 0xf2, 0xd2, 0x3c, 0xdb, 0xf2, 0x9b, 0xb2,
	0xfd, 0xad, 0xec, 0xa2, 0x03, 0x80, 0x04, 0x2f, 0x49, 0x14, 0xc8, 0x92, 0x55, 0x97, 0x8c, 0x12,
	0xa0, 0x6a, 0x1e, 0x13, 0x19, 0x3a, 0x84, 0x96, 0x20, 0x29, 0x8b, 0x73, 0x22, 0x02, 0x59, 0xb5,
	0xee, 0x48, 0x8c, 0xbf, 0xa7, 0x8b, 0x0f, 0xf1, 0x52, 0xb2, 0xec, 0x52, 0x9e, 0xdf, 0x0b, 0x42,
	0xb6, 0xe4, 0x38, 0xb3, 0xb6, 0xfb, 0xc6, 0xa8, 0xe1, 0x83, 0x2c, 0x4d, 0x55, 0x65, 0xd8, 0x85,
	0x86, 0xde, 0x0a, 0xed, 0x80, 0x79, 0x3e, 0xf5, 0x3a, 0x5b, 0xf2, 0xe3, 0xd1, 0x67, 0x5e, 0xc7,
	0x98, 0x7c, 0x28, 0x0d, 0x7e, 0x0f, 0x86, 0x1b, 0x0d, 0x7e, 0xc9, 0x1b, 0xf7, 0x2d, 0x68, 0xab,
	0xdd, 0x52, 0x4e, 0x42, 0x3a, 0xa7, 0x44, 0x20, 0xf3, 0x6f, 0xd7, 0x18, 0xfe, 0x52, 0x83, 0xbd,
	0xf3, 0x90, 0x3f, 0x20, 0x84, 0xe3, 0x98, 0xe6, 0x04, 0x7d, 0x01, 0x9d, 0x27, 0xfa, 0x27, 0xe0,
	0x82, 0xcd, 0x48, 0xaa, 0xee, 0x61, 0x77, 0xdc, 0xb5, 0x8b, 0x30, 0xda, 0x3a, 0x8c, 0xf6, 0xa3,
	0xfb, 0x49, 0x76, 0x32, 0x56, 0x7e, 0xf8, 0x6f, 0x54, 0x53, 0x9e, 0x1a, 0x42, 0x53, 0x68, 0xaf,
	0x89, 0x32, 0xba, 0x2c, 0x42, 0x72, 0x13, 0x4d, 0xab, 0x9a, 0x39, 0xa7, 0x4b, 0x82, 0x1e, 0x00,
	0x5a, 0x93, 0xd0, 0x24, 0x23, 0x22, 0xc7, 0xb1, 0xba, 0x99, 0x9b, 0x88, 0xde, 0xac, 0xe6, 0xee,
	0x97, 0x63, 0x93, 0x91, 0x74, 0xeb, 0x10, 0x06, 0x1b, 0xdd, 0xba, 0x6a, 0xc2, 0xf0, 0xd7, 0x1a,
	0x80, 0x4b, 0x93, 0x68, 0xaa, 0x30, 0xe8, 0x7b, 0x68, 0xa7, 0x6c, 0x25, 0x42, 0x12, 0xe8, 0x04,
	0x15, 0x8e, 0x8c, 0x6e, 0x9b, 0x4c, 0x15, 0xc8, 0x9f, 0x55, 0x20, 0x5b, 0x05, 0x93, 0x8e, 0xfb,
	0xa7, 0xd0, 0x98, 0x0b, 0x42, 0x66, 0x34, 0x89, 0x4a, 0x7f, 0xf6, 0x5f, 0x59, 0xcb, 0x65, 0x2c,
	0x2e, 0x96, 0xaa, 0xb0, 0xe8, 0xa1, 0x94, 0x74, 0xe5, 0xbd, 0x48, 0x2d, 0xb3, 0x6f, 0x8e, 0x76,
	0xc7, 0x47, 0x37, 0x4a, 0xfa, 0x4a, 0xe1, 0xa5, 0x8e, 0xf5, 0x5f, 0x3a, 0xf9, 0x40, 0x7a, 0x33,
	0x80, 0x83, 0x8d, 0xde, 0xac, 0xad, 0x18, 0xbe, 0x30, 0x60, 0x47, 0x6b, 0xff, 0xba, 0xd2, 0xf0,
	0x9a, 0xb6, 0x9c, 0x6d, 0x69, 0x19, 0x9a, 0xf2, 0x13, 0xa8, 0x73, 0xca, 0x75, 0x54, 0x0e, 0xfe,
	0x87, 0x48, 0x3e, 0x1f, 0x67, 0x5b, 0xbe, 0x82, 0x4f, 0x0e, 0xa5, 0xfa, 0x1e, 0x74, 0x37, 0xaa,
	0xd7, 0x77, 0xd0, 0xae, 0x1e, 0x80, 0x22, 0xfa, 0xbf, 0x1b, 0xd0, 0x9c, 0xd2, 0x48, 0xf8, 0x38,
	0x59, 0x10, 0x64, 0x43, 0xbb, 0xec, 0x06, 0x5c, 0x90, 0x39, 0x7d, 0xfa, 0xdf, 0x07, 0xaf, 0x55,
	0xb6, 0x3d, 0xd5, 0x45, 0x9f, 0x03, 0x14, 0xb8, 0x20, 0x26, 0xc9, 0x6d, 0xa2, 0xad, 0x5f, 0xc7,
	0xe7, 0x86, 0xdf, 0x2c, 0x26, 0xbf, 0x24, 0xc9, 0xe4, 0x7d, 0xa9, 0xbc, 0x0f, 0xbd, 0x8d, 0xca,
	0x2b, 0x75, 0xee, 0xe9, 0x1f, 0xcf, 0x5f, 0xfc, 0xb5, 0x5d, 0xeb, 0x98, 0x70, 0x44, 0x59, 0xe1,
	0x0a, 0x17, 0xec, 0xe9, 0xb3, 0xeb, 0x0d, 0x72, 0xf7, 0x4e, 0xb5, 0x5e, 0x96, 0x31, 0xcf, 0x98,
	0x6d, 0x2b, 0x51, 0x27, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x4c, 0xe4, 0x38, 0xc6, 0x06,
	0x00, 0x00,
}
