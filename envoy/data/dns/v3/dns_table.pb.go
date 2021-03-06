// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/dns/v3/dns_table.proto

package envoy_data_dns_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type DnsTable struct {
	ExternalRetryCount   uint32                       `protobuf:"varint,1,opt,name=external_retry_count,json=externalRetryCount,proto3" json:"external_retry_count,omitempty"`
	VirtualDomains       []*DnsTable_DnsVirtualDomain `protobuf:"bytes,2,rep,name=virtual_domains,json=virtualDomains,proto3" json:"virtual_domains,omitempty"`
	KnownSuffixes        []*v3.StringMatcher          `protobuf:"bytes,3,rep,name=known_suffixes,json=knownSuffixes,proto3" json:"known_suffixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DnsTable) Reset()         { *m = DnsTable{} }
func (m *DnsTable) String() string { return proto.CompactTextString(m) }
func (*DnsTable) ProtoMessage()    {}
func (*DnsTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0588926abfe2398, []int{0}
}

func (m *DnsTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable.Unmarshal(m, b)
}
func (m *DnsTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable.Marshal(b, m, deterministic)
}
func (m *DnsTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable.Merge(m, src)
}
func (m *DnsTable) XXX_Size() int {
	return xxx_messageInfo_DnsTable.Size(m)
}
func (m *DnsTable) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable proto.InternalMessageInfo

func (m *DnsTable) GetExternalRetryCount() uint32 {
	if m != nil {
		return m.ExternalRetryCount
	}
	return 0
}

func (m *DnsTable) GetVirtualDomains() []*DnsTable_DnsVirtualDomain {
	if m != nil {
		return m.VirtualDomains
	}
	return nil
}

func (m *DnsTable) GetKnownSuffixes() []*v3.StringMatcher {
	if m != nil {
		return m.KnownSuffixes
	}
	return nil
}

type DnsTable_AddressList struct {
	Address              []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnsTable_AddressList) Reset()         { *m = DnsTable_AddressList{} }
func (m *DnsTable_AddressList) String() string { return proto.CompactTextString(m) }
func (*DnsTable_AddressList) ProtoMessage()    {}
func (*DnsTable_AddressList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0588926abfe2398, []int{0, 0}
}

func (m *DnsTable_AddressList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_AddressList.Unmarshal(m, b)
}
func (m *DnsTable_AddressList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_AddressList.Marshal(b, m, deterministic)
}
func (m *DnsTable_AddressList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_AddressList.Merge(m, src)
}
func (m *DnsTable_AddressList) XXX_Size() int {
	return xxx_messageInfo_DnsTable_AddressList.Size(m)
}
func (m *DnsTable_AddressList) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_AddressList.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_AddressList proto.InternalMessageInfo

func (m *DnsTable_AddressList) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

type DnsTable_DnsEndpoint struct {
	// Types that are valid to be assigned to EndpointConfig:
	//	*DnsTable_DnsEndpoint_AddressList
	//	*DnsTable_DnsEndpoint_ClusterName
	EndpointConfig       isDnsTable_DnsEndpoint_EndpointConfig `protobuf_oneof:"endpoint_config"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DnsTable_DnsEndpoint) Reset()         { *m = DnsTable_DnsEndpoint{} }
func (m *DnsTable_DnsEndpoint) String() string { return proto.CompactTextString(m) }
func (*DnsTable_DnsEndpoint) ProtoMessage()    {}
func (*DnsTable_DnsEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0588926abfe2398, []int{0, 1}
}

func (m *DnsTable_DnsEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Unmarshal(m, b)
}
func (m *DnsTable_DnsEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Marshal(b, m, deterministic)
}
func (m *DnsTable_DnsEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_DnsEndpoint.Merge(m, src)
}
func (m *DnsTable_DnsEndpoint) XXX_Size() int {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Size(m)
}
func (m *DnsTable_DnsEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_DnsEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_DnsEndpoint proto.InternalMessageInfo

type isDnsTable_DnsEndpoint_EndpointConfig interface {
	isDnsTable_DnsEndpoint_EndpointConfig()
}

type DnsTable_DnsEndpoint_AddressList struct {
	AddressList *DnsTable_AddressList `protobuf:"bytes,1,opt,name=address_list,json=addressList,proto3,oneof"`
}

type DnsTable_DnsEndpoint_ClusterName struct {
	ClusterName string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3,oneof"`
}

func (*DnsTable_DnsEndpoint_AddressList) isDnsTable_DnsEndpoint_EndpointConfig() {}

func (*DnsTable_DnsEndpoint_ClusterName) isDnsTable_DnsEndpoint_EndpointConfig() {}

func (m *DnsTable_DnsEndpoint) GetEndpointConfig() isDnsTable_DnsEndpoint_EndpointConfig {
	if m != nil {
		return m.EndpointConfig
	}
	return nil
}

func (m *DnsTable_DnsEndpoint) GetAddressList() *DnsTable_AddressList {
	if x, ok := m.GetEndpointConfig().(*DnsTable_DnsEndpoint_AddressList); ok {
		return x.AddressList
	}
	return nil
}

func (m *DnsTable_DnsEndpoint) GetClusterName() string {
	if x, ok := m.GetEndpointConfig().(*DnsTable_DnsEndpoint_ClusterName); ok {
		return x.ClusterName
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DnsTable_DnsEndpoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DnsTable_DnsEndpoint_AddressList)(nil),
		(*DnsTable_DnsEndpoint_ClusterName)(nil),
	}
}

type DnsTable_DnsVirtualDomain struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Endpoint             *DnsTable_DnsEndpoint `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AnswerTtl            *duration.Duration    `protobuf:"bytes,3,opt,name=answer_ttl,json=answerTtl,proto3" json:"answer_ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DnsTable_DnsVirtualDomain) Reset()         { *m = DnsTable_DnsVirtualDomain{} }
func (m *DnsTable_DnsVirtualDomain) String() string { return proto.CompactTextString(m) }
func (*DnsTable_DnsVirtualDomain) ProtoMessage()    {}
func (*DnsTable_DnsVirtualDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0588926abfe2398, []int{0, 2}
}

func (m *DnsTable_DnsVirtualDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Unmarshal(m, b)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Marshal(b, m, deterministic)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_DnsVirtualDomain.Merge(m, src)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Size() int {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Size(m)
}
func (m *DnsTable_DnsVirtualDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_DnsVirtualDomain.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_DnsVirtualDomain proto.InternalMessageInfo

func (m *DnsTable_DnsVirtualDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsTable_DnsVirtualDomain) GetEndpoint() *DnsTable_DnsEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *DnsTable_DnsVirtualDomain) GetAnswerTtl() *duration.Duration {
	if m != nil {
		return m.AnswerTtl
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsTable)(nil), "envoy.data.dns.v3.DnsTable")
	proto.RegisterType((*DnsTable_AddressList)(nil), "envoy.data.dns.v3.DnsTable.AddressList")
	proto.RegisterType((*DnsTable_DnsEndpoint)(nil), "envoy.data.dns.v3.DnsTable.DnsEndpoint")
	proto.RegisterType((*DnsTable_DnsVirtualDomain)(nil), "envoy.data.dns.v3.DnsTable.DnsVirtualDomain")
}

func init() { proto.RegisterFile("envoy/data/dns/v3/dns_table.proto", fileDescriptor_c0588926abfe2398) }

var fileDescriptor_c0588926abfe2398 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x18, 0xc5, 0x49, 0xd9, 0x5a, 0x67, 0x3f, 0x4a, 0x84, 0x20, 0x54, 0x62, 0xeb, 0x06, 0x82, 0x8a,
	0x1f, 0x09, 0xca, 0x0e, 0x88, 0x8a, 0x0b, 0x59, 0x91, 0x90, 0x18, 0x68, 0xca, 0x06, 0xd7, 0xc8,
	0x6b, 0xdc, 0xce, 0x5a, 0x6a, 0x47, 0xb6, 0x93, 0xad, 0x37, 0x8e, 0x9c, 0x39, 0xf2, 0x27, 0xf0,
	0x27, 0x70, 0xe2, 0x02, 0xe2, 0xc0, 0x85, 0x7f, 0x06, 0xa1, 0x9e, 0x90, 0x9d, 0x84, 0x85, 0x55,
	0x62, 0x3b, 0xc5, 0xfe, 0xbe, 0xef, 0x3d, 0xbf, 0xf7, 0xa2, 0x0f, 0x6e, 0x60, 0x9a, 0xb3, 0xa9,
	0x17, 0x23, 0x89, 0xbc, 0x98, 0x0a, 0x2f, 0xdf, 0x52, 0x9f, 0x48, 0xa2, 0x83, 0x04, 0xbb, 0x29,
	0x67, 0x92, 0xd9, 0x57, 0xf4, 0x88, 0xab, 0x46, 0xdc, 0x98, 0x0a, 0x37, 0xdf, 0xea, 0x6c, 0x16,
	0x28, 0x39, 0x4d, 0xb1, 0x37, 0x41, 0x72, 0x78, 0x88, 0xb9, 0x42, 0x0a, 0xc9, 0x09, 0x1d, 0x17,
	0xb0, 0xce, 0xda, 0x98, 0xb1, 0x71, 0x82, 0x3d, 0x7d, 0x3b, 0xc8, 0x46, 0x5e, 0x9c, 0x71, 0x24,
	0x09, 0xa3, 0x65, 0xff, 0x66, 0x16, 0xa7, 0xc8, 0x43, 0x94, 0x32, 0xa9, 0xcb, 0xc2, 0x13, 0x12,
	0xc9, 0x4c, 0x94, 0xed, 0x8d, 0xb9, 0x76, 0x8e, 0xb9, 0x20, 0x8c, 0x9e, 0xbe, 0x70, 0x3d, 0x47,
	0x09, 0x89, 0x91, 0xc4, 0x5e, 0x75, 0x28, 0x1a, 0x9b, 0xdf, 0x16, 0x60, 0x73, 0x40, 0xc5, 0xbe,
	0x32, 0x61, 0x3f, 0x81, 0x57, 0xf1, 0x89, 0xc4, 0x9c, 0xa2, 0x24, 0xe2, 0x58, 0xf2, 0x69, 0x34,
	0x64, 0x19, 0x95, 0x0e, 0xe8, 0x82, 0xde, 0x72, 0xb0, 0x38, 0x0b, 0x1a, 0xf7, 0x0c, 0xc7, 0x0c,
	0xed, 0x6a, 0x28, 0x54, 0x33, 0xdb, 0x6a, 0xc4, 0x7e, 0x03, 0x57, 0x73, 0xc2, 0x65, 0x86, 0x92,
	0x28, 0x66, 0x13, 0x44, 0xa8, 0x70, 0x8c, 0xae, 0xd9, 0xb3, 0xfc, 0x07, 0xee, 0x5c, 0x26, 0x6e,
	0xf5, 0xa0, 0x3a, 0xbc, 0x2d, 0x50, 0x03, 0x0d, 0x0a, 0x57, 0xf2, 0xfa, 0x55, 0xd8, 0x2f, 0xe1,
	0xca, 0x11, 0x65, 0xc7, 0x34, 0x12, 0xd9, 0x68, 0x44, 0x4e, 0xb0, 0x70, 0x4c, 0xcd, 0x7a, 0xbb,
	0x64, 0x55, 0xb1, 0xba, 0x65, 0xac, 0x8a, 0x79, 0x4f, 0xc7, 0xfa, 0xaa, 0x28, 0x84, 0xcb, 0x1a,
	0xbb, 0x57, 0x42, 0x3b, 0x47, 0xd0, 0x7a, 0x16, 0xc7, 0x1c, 0x0b, 0xb1, 0x43, 0x84, 0xb4, 0x7b,
	0x70, 0x11, 0x15, 0x57, 0x07, 0x74, 0xcd, 0x5e, 0x2b, 0x58, 0x99, 0x05, 0xd6, 0x07, 0xd0, 0x6c,
	0x82, 0xcd, 0x06, 0x37, 0xda, 0x66, 0x58, 0xb5, 0xfb, 0xfe, 0xc7, 0xaf, 0xef, 0xd7, 0x1e, 0xc2,
	0xfb, 0x67, 0x9d, 0xf8, 0x28, 0x49, 0x0f, 0xd1, 0xa9, 0x9d, 0x1a, 0x7b, 0xe7, 0x07, 0x80, 0xd6,
	0x80, 0x8a, 0xe7, 0x34, 0x4e, 0x19, 0xa1, 0xd2, 0xde, 0x81, 0x4b, 0x25, 0x5d, 0x94, 0x10, 0x51,
	0x64, 0x6a, 0xf9, 0x77, 0xff, 0x97, 0x4e, 0x8d, 0xee, 0xc5, 0xa5, 0xd0, 0x42, 0x35, 0xed, 0xb7,
	0xe0, 0xd2, 0x30, 0xc9, 0x84, 0xc4, 0x3c, 0xa2, 0x68, 0x82, 0x1d, 0xa3, 0x0b, 0x7a, 0x2d, 0x35,
	0x54, 0x56, 0x5f, 0xa3, 0x09, 0xbe, 0xb0, 0xec, 0x9a, 0xcc, 0xe0, 0x1a, 0x5c, 0xc5, 0xe5, 0x39,
	0x1a, 0x32, 0x3a, 0x22, 0x63, 0xdb, 0xfc, 0x1d, 0x80, 0xce, 0x2f, 0x00, 0xdb, 0x67, 0xff, 0x96,
	0xbd, 0x06, 0x1b, 0xfa, 0x75, 0xe5, 0xa5, 0x15, 0xc0, 0x59, 0xb0, 0xc8, 0x2f, 0xb7, 0x8d, 0x2f,
	0x00, 0x84, 0xba, 0x6e, 0x6f, 0xc3, 0x66, 0x45, 0xa6, 0x15, 0x9e, 0xe3, 0xb7, 0xa6, 0x23, 0xfc,
	0x0b, 0xb4, 0x07, 0x10, 0x22, 0x2a, 0x8e, 0x31, 0x8f, 0xa4, 0x4c, 0x1c, 0x53, 0xd3, 0xdc, 0x70,
	0x8b, 0x8d, 0x71, 0xab, 0x8d, 0x71, 0x07, 0xe5, 0xc6, 0x68, 0x15, 0x9f, 0x40, 0xc3, 0x37, 0x9a,
	0x4f, 0xc3, 0x56, 0x01, 0xdc, 0x97, 0x49, 0xff, 0xb1, 0xca, 0xc2, 0x87, 0x8f, 0x2e, 0x90, 0xc5,
	0x3f, 0x1e, 0xfb, 0x77, 0x14, 0x70, 0x03, 0xae, 0x9f, 0x03, 0x0c, 0xfa, 0x9f, 0xdf, 0x7d, 0xff,
	0xb9, 0x60, 0xb4, 0x0d, 0xb8, 0x4e, 0x58, 0xe1, 0x32, 0xe5, 0xec, 0x64, 0x3a, 0x6f, 0x38, 0x58,
	0xae, 0x40, 0xbb, 0x4a, 0xfd, 0x2e, 0x38, 0x58, 0xd0, 0x36, 0xb6, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x46, 0x3c, 0x47, 0xf6, 0x62, 0x04, 0x00, 0x00,
}
