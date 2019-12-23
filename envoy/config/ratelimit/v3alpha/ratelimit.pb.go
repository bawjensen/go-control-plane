// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/ratelimit/v3alpha/ratelimit.proto

package envoy_config_ratelimit_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type RateLimitDescriptor struct {
	Entries              []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitDescriptor) Reset()         { *m = RateLimitDescriptor{} }
func (m *RateLimitDescriptor) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor) ProtoMessage()    {}
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_b69d4caa297e0eb9, []int{0}
}

func (m *RateLimitDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor.Unmarshal(m, b)
}
func (m *RateLimitDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor.Merge(m, src)
}
func (m *RateLimitDescriptor) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor.Size(m)
}
func (m *RateLimitDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor proto.InternalMessageInfo

func (m *RateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RateLimitDescriptor_Entry struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitDescriptor_Entry) Reset()         { *m = RateLimitDescriptor_Entry{} }
func (m *RateLimitDescriptor_Entry) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor_Entry) ProtoMessage()    {}
func (*RateLimitDescriptor_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b69d4caa297e0eb9, []int{0, 0}
}

func (m *RateLimitDescriptor_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Unmarshal(m, b)
}
func (m *RateLimitDescriptor_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptor_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor_Entry.Merge(m, src)
}
func (m *RateLimitDescriptor_Entry) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Size(m)
}
func (m *RateLimitDescriptor_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor_Entry proto.InternalMessageInfo

func (m *RateLimitDescriptor_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptor_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*RateLimitDescriptor)(nil), "envoy.config.ratelimit.v3alpha.RateLimitDescriptor")
	proto.RegisterType((*RateLimitDescriptor_Entry)(nil), "envoy.config.ratelimit.v3alpha.RateLimitDescriptor.Entry")
}

func init() {
	proto.RegisterFile("envoy/config/ratelimit/v3alpha/ratelimit.proto", fileDescriptor_b69d4caa297e0eb9)
}

var fileDescriptor_b69d4caa297e0eb9 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x18, 0xc4, 0xe5, 0x94, 0xd0, 0x62, 0x24, 0x86, 0x30, 0x10, 0x22, 0x51, 0x05, 0xa6, 0x0a, 0x55,
	0x36, 0xa4, 0x03, 0xa2, 0x13, 0xb2, 0x60, 0x63, 0xa8, 0xb2, 0x33, 0x98, 0xd6, 0x94, 0x4f, 0x04,
	0xdb, 0x72, 0x5c, 0x8b, 0xf0, 0x04, 0xcc, 0x8c, 0xac, 0x3c, 0x0b, 0x2f, 0xd5, 0x09, 0x25, 0x26,
	0x2a, 0x48, 0xfc, 0xdb, 0x3e, 0xdd, 0xdd, 0x77, 0x3a, 0xfd, 0x30, 0x11, 0xd2, 0xa9, 0x8a, 0x4e,
	0x95, 0xbc, 0x81, 0x39, 0x35, 0xdc, 0x8a, 0x02, 0xee, 0xc1, 0x52, 0x37, 0xe2, 0x85, 0xbe, 0xe5,
	0x2b, 0x85, 0x68, 0xa3, 0xac, 0x8a, 0xfa, 0x4d, 0x9e, 0xf8, 0x3c, 0x59, 0xb9, 0x1f, 0xf9, 0x64,
	0x7f, 0x31, 0xd3, 0x9c, 0x72, 0x29, 0x95, 0xe5, 0x16, 0x94, 0x2c, 0xa9, 0x13, 0xa6, 0x04, 0x25,
	0x41, 0xce, 0x7d, 0x45, 0xb2, 0xe3, 0x78, 0x01, 0x33, 0x6e, 0x05, 0x6d, 0x0f, 0x6f, 0x1c, 0xbc,
	0x06, 0x78, 0x3b, 0xe7, 0x56, 0x5c, 0xd6, 0x8d, 0xe7, 0xa2, 0x9c, 0x1a, 0xd0, 0x56, 0x99, 0xe8,
	0x0a, 0x77, 0x85, 0xb4, 0x06, 0x44, 0x19, 0xa3, 0xb4, 0x33, 0xd8, 0xcc, 0x4e, 0xc9, 0xef, 0x2b,
	0xc8, 0x37, 0x2d, 0xe4, 0x42, 0x5a, 0x53, 0xb1, 0xde, 0x92, 0x85, 0xcf, 0x28, 0xe8, 0xa1, 0xbc,
	0xed, 0x4c, 0x1e, 0x71, 0xd8, 0x78, 0xd1, 0x2e, 0xee, 0xdc, 0x89, 0x2a, 0x46, 0x29, 0x1a, 0x6c,
	0xb0, 0xee, 0x92, 0xad, 0x99, 0x20, 0x45, 0x79, 0xad, 0x45, 0x7b, 0x38, 0x74, 0xbc, 0x58, 0x88,
	0x38, 0xf8, 0x6a, 0x7a, 0x75, 0x7c, 0xf2, 0xf2, 0xf6, 0xd4, 0xcf, 0xf0, 0x91, 0x9f, 0xc5, 0x35,
	0x10, 0x97, 0x7d, 0x9a, 0xf5, 0xe3, 0x9c, 0xf1, 0x71, 0xfd, 0x38, 0xc4, 0x87, 0xff, 0x7f, 0x64,
	0x67, 0x78, 0x08, 0xca, 0x03, 0xd0, 0x46, 0x3d, 0x54, 0x7f, 0xb0, 0x60, 0x5b, 0x79, 0x2b, 0x4d,
	0x6a, 0xca, 0x13, 0x74, 0xbd, 0xde, 0xe0, 0x1e, 0xbd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x07, 0x16,
	0x3f, 0xe8, 0xfc, 0x01, 0x00, 0x00,
}