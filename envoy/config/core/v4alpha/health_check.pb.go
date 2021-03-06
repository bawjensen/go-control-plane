// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v4alpha/health_check.proto

package envoy_config_core_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v4alpha"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type HealthStatus int32

const (
	HealthStatus_UNKNOWN   HealthStatus = 0
	HealthStatus_HEALTHY   HealthStatus = 1
	HealthStatus_UNHEALTHY HealthStatus = 2
	HealthStatus_DRAINING  HealthStatus = 3
	HealthStatus_TIMEOUT   HealthStatus = 4
	HealthStatus_DEGRADED  HealthStatus = 5
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
	5: "DEGRADED",
}

var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
	"DEGRADED":  5,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}

func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0}
}

type HealthCheck struct {
	Timeout               *duration.Duration    `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Interval              *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	InitialJitter         *duration.Duration    `protobuf:"bytes,20,opt,name=initial_jitter,json=initialJitter,proto3" json:"initial_jitter,omitempty"`
	IntervalJitter        *duration.Duration    `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter,proto3" json:"interval_jitter,omitempty"`
	IntervalJitterPercent uint32                `protobuf:"varint,18,opt,name=interval_jitter_percent,json=intervalJitterPercent,proto3" json:"interval_jitter_percent,omitempty"`
	UnhealthyThreshold    *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	HealthyThreshold      *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	AltPort               *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=alt_port,json=altPort,proto3" json:"alt_port,omitempty"`
	ReuseConnection       *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection,proto3" json:"reuse_connection,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_GrpcHealthCheck_
	//	*HealthCheck_CustomHealthCheck_
	HealthChecker                isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
	NoTrafficInterval            *duration.Duration          `protobuf:"bytes,12,opt,name=no_traffic_interval,json=noTrafficInterval,proto3" json:"no_traffic_interval,omitempty"`
	UnhealthyInterval            *duration.Duration          `protobuf:"bytes,14,opt,name=unhealthy_interval,json=unhealthyInterval,proto3" json:"unhealthy_interval,omitempty"`
	UnhealthyEdgeInterval        *duration.Duration          `protobuf:"bytes,15,opt,name=unhealthy_edge_interval,json=unhealthyEdgeInterval,proto3" json:"unhealthy_edge_interval,omitempty"`
	HealthyEdgeInterval          *duration.Duration          `protobuf:"bytes,16,opt,name=healthy_edge_interval,json=healthyEdgeInterval,proto3" json:"healthy_edge_interval,omitempty"`
	EventLogPath                 string                      `protobuf:"bytes,17,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	EventService                 *EventServiceConfig         `protobuf:"bytes,22,opt,name=event_service,json=eventService,proto3" json:"event_service,omitempty"`
	AlwaysLogHealthCheckFailures bool                        `protobuf:"varint,19,opt,name=always_log_health_check_failures,json=alwaysLogHealthCheckFailures,proto3" json:"always_log_health_check_failures,omitempty"`
	TlsOptions                   *HealthCheck_TlsOptions     `protobuf:"bytes,21,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty"`
	TransportSocketMatchCriteria *_struct.Struct             `protobuf:"bytes,23,opt,name=transport_socket_match_criteria,json=transportSocketMatchCriteria,proto3" json:"transport_socket_match_criteria,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                    `json:"-"`
	XXX_unrecognized             []byte                      `json:"-"`
	XXX_sizecache                int32                       `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetInitialJitter() *duration.Duration {
	if m != nil {
		return m.InitialJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *duration.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitterPercent() uint32 {
	if m != nil {
		return m.IntervalJitterPercent
	}
	return 0
}

func (m *HealthCheck) GetUnhealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *wrappers.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *wrappers.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,proto3,oneof"`
}

type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,proto3,oneof"`
}

type HealthCheck_GrpcHealthCheck_ struct {
	GrpcHealthCheck *HealthCheck_GrpcHealthCheck `protobuf:"bytes,11,opt,name=grpc_health_check,json=grpcHealthCheck,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_ struct {
	CustomHealthCheck *HealthCheck_CustomHealthCheck `protobuf:"bytes,13,opt,name=custom_health_check,json=customHealthCheck,proto3,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_GrpcHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_CustomHealthCheck_) isHealthCheck_HealthChecker() {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetGrpcHealthCheck() *HealthCheck_GrpcHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_GrpcHealthCheck_); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetCustomHealthCheck() *HealthCheck_CustomHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_CustomHealthCheck_); ok {
		return x.CustomHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetNoTrafficInterval() *duration.Duration {
	if m != nil {
		return m.NoTrafficInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetHealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.HealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

func (m *HealthCheck) GetEventService() *EventServiceConfig {
	if m != nil {
		return m.EventService
	}
	return nil
}

func (m *HealthCheck) GetAlwaysLogHealthCheckFailures() bool {
	if m != nil {
		return m.AlwaysLogHealthCheckFailures
	}
	return false
}

func (m *HealthCheck) GetTlsOptions() *HealthCheck_TlsOptions {
	if m != nil {
		return m.TlsOptions
	}
	return nil
}

func (m *HealthCheck) GetTransportSocketMatchCriteria() *_struct.Struct {
	if m != nil {
		return m.TransportSocketMatchCriteria
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_GrpcHealthCheck_)(nil),
		(*HealthCheck_CustomHealthCheck_)(nil),
	}
}

type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload              isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HealthCheck_Payload) Reset()         { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()    {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0, 0}
}

func (m *HealthCheck_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Payload.Unmarshal(m, b)
}
func (m *HealthCheck_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Payload.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Payload.Merge(m, src)
}
func (m *HealthCheck_Payload) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Payload.Size(m)
}
func (m *HealthCheck_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Payload proto.InternalMessageInfo

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload() {}

func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

type HealthCheck_HttpHealthCheck struct {
	Host                   string                 `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                   string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Send                   *HealthCheck_Payload   `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	Receive                *HealthCheck_Payload   `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	RequestHeadersToAdd    []*HeaderValueOption   `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove []string               `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	ExpectedStatuses       []*v3.Int64Range       `protobuf:"bytes,9,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	CodecClientType        v3.CodecClientType     `protobuf:"varint,10,opt,name=codec_client_type,json=codecClientType,proto3,enum=envoy.type.v3.CodecClientType" json:"codec_client_type,omitempty"`
	ServiceNameMatcher     *v4alpha.StringMatcher `protobuf:"bytes,11,opt,name=service_name_matcher,json=serviceNameMatcher,proto3" json:"service_name_matcher,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}               `json:"-"`
	XXX_unrecognized       []byte                 `json:"-"`
	XXX_sizecache          int32                  `json:"-"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()         { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0, 1}
}

func (m *HealthCheck_HttpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Size(m)
}
func (m *HealthCheck_HttpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_HttpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetExpectedStatuses() []*v3.Int64Range {
	if m != nil {
		return m.ExpectedStatuses
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetCodecClientType() v3.CodecClientType {
	if m != nil {
		return m.CodecClientType
	}
	return v3.CodecClientType_HTTP1
}

func (m *HealthCheck_HttpHealthCheck) GetServiceNameMatcher() *v4alpha.StringMatcher {
	if m != nil {
		return m.ServiceNameMatcher
	}
	return nil
}

type HealthCheck_TcpHealthCheck struct {
	Send                 *HealthCheck_Payload   `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	Receive              []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()         { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0, 2}
}

func (m *HealthCheck_TcpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Size(m)
}
func (m *HealthCheck_TcpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TcpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_RedisHealthCheck) Reset()         { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()    {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0, 3}
}

func (m *HealthCheck_RedisHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.Merge(m, src)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Size(m)
}
func (m *HealthCheck_RedisHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_RedisHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_RedisHealthCheck) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type HealthCheck_GrpcHealthCheck struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_GrpcHealthCheck) Reset()         { *m = HealthCheck_GrpcHealthCheck{} }
func (m *HealthCheck_GrpcHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_GrpcHealthCheck) ProtoMessage()    {}
func (*HealthCheck_GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0, 4}
}

func (m *HealthCheck_GrpcHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.Merge(m, src)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Size(m)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_GrpcHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_GrpcHealthCheck) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

type HealthCheck_CustomHealthCheck struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*HealthCheck_CustomHealthCheck_TypedConfig
	ConfigType           isHealthCheck_CustomHealthCheck_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *HealthCheck_CustomHealthCheck) Reset()         { *m = HealthCheck_CustomHealthCheck{} }
func (m *HealthCheck_CustomHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_CustomHealthCheck) ProtoMessage()    {}
func (*HealthCheck_CustomHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0, 5}
}

func (m *HealthCheck_CustomHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.Merge(m, src)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Size(m)
}
func (m *HealthCheck_CustomHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_CustomHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_CustomHealthCheck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isHealthCheck_CustomHealthCheck_ConfigType interface {
	isHealthCheck_CustomHealthCheck_ConfigType()
}

type HealthCheck_CustomHealthCheck_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*HealthCheck_CustomHealthCheck_TypedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (m *HealthCheck_CustomHealthCheck) GetConfigType() isHealthCheck_CustomHealthCheck_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_CustomHealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_CustomHealthCheck_TypedConfig)(nil),
	}
}

type HealthCheck_TlsOptions struct {
	AlpnProtocols        []string `protobuf:"bytes,1,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_TlsOptions) Reset()         { *m = HealthCheck_TlsOptions{} }
func (m *HealthCheck_TlsOptions) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TlsOptions) ProtoMessage()    {}
func (*HealthCheck_TlsOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c892f9ff37cb3f95, []int{0, 6}
}

func (m *HealthCheck_TlsOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TlsOptions.Unmarshal(m, b)
}
func (m *HealthCheck_TlsOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TlsOptions.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TlsOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TlsOptions.Merge(m, src)
}
func (m *HealthCheck_TlsOptions) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TlsOptions.Size(m)
}
func (m *HealthCheck_TlsOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TlsOptions.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TlsOptions proto.InternalMessageInfo

func (m *HealthCheck_TlsOptions) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.core.v4alpha.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.core.v4alpha.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.config.core.v4alpha.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.config.core.v4alpha.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.config.core.v4alpha.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.config.core.v4alpha.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.config.core.v4alpha.HealthCheck.GrpcHealthCheck")
	proto.RegisterType((*HealthCheck_CustomHealthCheck)(nil), "envoy.config.core.v4alpha.HealthCheck.CustomHealthCheck")
	proto.RegisterType((*HealthCheck_TlsOptions)(nil), "envoy.config.core.v4alpha.HealthCheck.TlsOptions")
}

func init() {
	proto.RegisterFile("envoy/config/core/v4alpha/health_check.proto", fileDescriptor_c892f9ff37cb3f95)
}

var fileDescriptor_c892f9ff37cb3f95 = []byte{
	// 1594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0x4f, 0x73, 0xdb, 0xc6,
	0x15, 0x17, 0x48, 0x8a, 0x7f, 0x1e, 0x45, 0x0a, 0x5a, 0x59, 0x16, 0xc4, 0x38, 0x0e, 0xed, 0xa6,
	0x63, 0xc5, 0x4d, 0xc8, 0x9a, 0x72, 0x9d, 0x44, 0x97, 0x56, 0x90, 0x69, 0x53, 0x8a, 0x2d, 0x6b,
	0x56, 0x74, 0x3d, 0x9d, 0xcc, 0x04, 0x5d, 0x03, 0x2b, 0x12, 0x31, 0x84, 0x45, 0x17, 0x4b, 0xc6,
	0xbc, 0x75, 0xda, 0x4b, 0xa7, 0xc7, 0x9c, 0x3a, 0xfd, 0x08, 0xf9, 0x00, 0x3d, 0xf4, 0xd4, 0x99,
	0x4e, 0x67, 0x72, 0xed, 0xa5, 0x9f, 0xa1, 0xdf, 0xa0, 0x1d, 0x9d, 0x3a, 0xbb, 0x00, 0x48, 0x82,
	0x94, 0x4c, 0xa9, 0x93, 0x1b, 0xf1, 0xde, 0xef, 0xf7, 0x7b, 0x6f, 0x77, 0xdf, 0xdb, 0xb7, 0x84,
	0x8f, 0xa9, 0x3f, 0x64, 0xa3, 0xa6, 0xcd, 0xfc, 0x53, 0xb7, 0xd7, 0xb4, 0x19, 0xa7, 0xcd, 0xe1,
	0x43, 0xe2, 0x05, 0x7d, 0xd2, 0xec, 0x53, 0xe2, 0x89, 0xbe, 0x65, 0xf7, 0xa9, 0xfd, 0xa6, 0x11,
	0x70, 0x26, 0x18, 0xda, 0x52, 0xe8, 0x46, 0x84, 0x6e, 0x48, 0x74, 0x23, 0x46, 0xd7, 0x3e, 0xbc,
	0x5c, 0xe8, 0x35, 0x09, 0x69, 0x24, 0x50, 0x7b, 0x78, 0x39, 0x8a, 0x0e, 0xa9, 0x2f, 0xac, 0x90,
	0xf2, 0xa1, 0x6b, 0x53, 0x2b, 0x0e, 0x11, 0xb1, 0xee, 0x45, 0x2c, 0x31, 0x0a, 0x68, 0xf3, 0x8c,
	0x08, 0xbb, 0x4f, 0xf9, 0x98, 0x16, 0x0a, 0xee, 0xfa, 0x09, 0xd0, 0x98, 0x02, 0x0e, 0x77, 0x9a,
	0x7d, 0x21, 0x82, 0xd8, 0xb3, 0x95, 0xf6, 0x70, 0xe2, 0xf7, 0x92, 0x9c, 0xb6, 0x7a, 0x8c, 0xf5,
	0x3c, 0xda, 0x54, 0x5f, 0xaf, 0x07, 0xa7, 0x4d, 0xe2, 0x8f, 0x62, 0xd7, 0xed, 0x59, 0x97, 0x33,
	0xe0, 0x44, 0xb8, 0xcc, 0x8f, 0xfd, 0xb7, 0x66, 0xfd, 0xa1, 0xe0, 0x03, 0x5b, 0x5c, 0xc6, 0xfe,
	0x86, 0x93, 0x20, 0xa0, 0x3c, 0x8c, 0xfd, 0x3f, 0x8a, 0x72, 0x22, 0xbe, 0xcf, 0x84, 0x52, 0x0d,
	0x9b, 0x0e, 0x0d, 0x38, 0xb5, 0xa7, 0x43, 0xbc, 0x3f, 0x70, 0x02, 0x92, 0xc2, 0x84, 0x82, 0x88,
	0x41, 0xa2, 0x71, 0x67, 0xce, 0x3d, 0xa4, 0x3c, 0x74, 0x99, 0x3f, 0xd9, 0x94, 0xcd, 0x21, 0xf1,
	0x5c, 0x87, 0x08, 0xda, 0x4c, 0x7e, 0x44, 0x8e, 0xbb, 0x7f, 0x79, 0x0f, 0xca, 0x1d, 0x75, 0xc8,
	0xfb, 0xf2, 0x8c, 0xd1, 0xcf, 0xa1, 0x20, 0xdc, 0x33, 0xca, 0x06, 0xc2, 0xd0, 0xea, 0xda, 0x76,
	0xb9, 0xb5, 0xd5, 0x88, 0x56, 0xd0, 0x48, 0x56, 0xd0, 0x78, 0x1c, 0xaf, 0xdf, 0x84, 0x73, 0xb3,
	0xf0, 0x9d, 0x96, 0x2b, 0x6a, 0xf7, 0x97, 0x70, 0xc2, 0x42, 0x7b, 0x50, 0x74, 0x7d, 0x41, 0xf9,
	0x90, 0x78, 0x46, 0xe6, 0x3a, 0x0a, 0x63, 0x1a, 0xfa, 0x05, 0x54, 0x5d, 0xdf, 0x15, 0x2e, 0xf1,
	0xac, 0xaf, 0x5d, 0x21, 0x28, 0x37, 0x6e, 0x2c, 0x10, 0xc2, 0x95, 0x98, 0x70, 0xa8, 0xf0, 0xc8,
	0x84, 0xd5, 0x44, 0x2d, 0x91, 0xc8, 0x2e, 0x92, 0xa8, 0x26, 0x8c, 0x58, 0xe3, 0x11, 0x6c, 0xce,
	0x68, 0x58, 0x01, 0xe5, 0x36, 0xf5, 0x85, 0x81, 0xea, 0xda, 0x76, 0x05, 0x6f, 0xa4, 0x09, 0xc7,
	0x91, 0x13, 0xbd, 0x82, 0xf5, 0x81, 0x1f, 0xf5, 0xcd, 0xc8, 0x12, 0x7d, 0x4e, 0xc3, 0x3e, 0xf3,
	0x1c, 0x23, 0xa7, 0xe2, 0xdf, 0x9a, 0x8b, 0xff, 0xf2, 0xc0, 0x17, 0x3b, 0xad, 0x5f, 0x12, 0x6f,
	0x40, 0xcd, 0xe2, 0xb9, 0xb9, 0xfc, 0x47, 0x2d, 0xa3, 0x6b, 0x18, 0x8d, 0x25, 0xba, 0x89, 0x02,
	0x3a, 0x81, 0xb5, 0x79, 0xd9, 0xe5, 0x6b, 0xc9, 0xea, 0x73, 0xa2, 0x9f, 0x42, 0x91, 0x78, 0xc2,
	0x0a, 0x18, 0x17, 0x46, 0x7e, 0xb1, 0x16, 0x2e, 0x10, 0x4f, 0x1c, 0x33, 0x2e, 0x50, 0x1b, 0x74,
	0x4e, 0x07, 0xa1, 0xea, 0x52, 0x9f, 0xda, 0x72, 0x0b, 0x8d, 0x82, 0x12, 0xa8, 0xcd, 0x09, 0x98,
	0x8c, 0x79, 0x11, 0x7d, 0x55, 0x71, 0xf6, 0xc7, 0x14, 0xe4, 0xc0, 0x9a, 0xec, 0x50, 0x6b, 0xfa,
	0xa2, 0x31, 0x8a, 0x4a, 0xe7, 0x51, 0xe3, 0xd2, 0x9b, 0xa6, 0x31, 0x55, 0xb2, 0x8d, 0x8e, 0x10,
	0xc1, 0xd4, 0x77, 0x67, 0x09, 0xaf, 0xf6, 0xd3, 0x26, 0x44, 0x40, 0x17, 0xf6, 0x4c, 0x90, 0x92,
	0x0a, 0xf2, 0xb3, 0x2b, 0x06, 0xe9, 0xda, 0x33, 0x31, 0xaa, 0x22, 0x65, 0x91, 0x0b, 0xe9, 0xf1,
	0xc0, 0x4e, 0xc7, 0x28, 0x5f, 0x6b, 0x21, 0x4f, 0x79, 0x60, 0xcf, 0x2c, 0xa4, 0x97, 0x36, 0xa1,
	0xaf, 0x61, 0xdd, 0x1e, 0x84, 0x82, 0x9d, 0xa5, 0xe3, 0x54, 0x54, 0x9c, 0xcf, 0xae, 0x18, 0x67,
	0x5f, 0x29, 0xa4, 0x23, 0xad, 0xd9, 0xb3, 0x46, 0x74, 0x02, 0xeb, 0x3e, 0xb3, 0x04, 0x27, 0xa7,
	0xa7, 0xae, 0x6d, 0x8d, 0x9b, 0x7a, 0x65, 0x51, 0x53, 0xcb, 0x72, 0xfb, 0x4e, 0xcb, 0xdc, 0x5f,
	0xc2, 0x6b, 0x3e, 0xeb, 0x46, 0xf4, 0x83, 0xa4, 0xb7, 0x31, 0x4c, 0x4a, 0x7b, 0xa2, 0x59, 0xbd,
	0x86, 0xe6, 0x98, 0x3e, 0xd6, 0xfc, 0x12, 0x36, 0x27, 0x9a, 0xd4, 0xe9, 0xd1, 0x89, 0xf0, 0xea,
	0xd5, 0x85, 0x37, 0xc6, 0x1a, 0x6d, 0xa7, 0x47, 0xc7, 0xe2, 0xaf, 0x60, 0xe3, 0x62, 0x69, 0xfd,
	0xea, 0xd2, 0xeb, 0x17, 0x09, 0x7f, 0x08, 0xd5, 0x68, 0xdc, 0x79, 0xac, 0x67, 0x05, 0x44, 0xf4,
	0x8d, 0xb5, 0xba, 0xb6, 0x5d, 0xc2, 0x2b, 0xca, 0xfa, 0x8c, 0xf5, 0x8e, 0x89, 0xe8, 0x23, 0x0c,
	0x95, 0xd4, 0x50, 0x34, 0x6e, 0xaa, 0xb0, 0x9f, 0xbc, 0xe3, 0xa8, 0xdb, 0x12, 0x7f, 0x12, 0xc1,
	0xf7, 0x95, 0x3f, 0xd6, 0x8c, 0x6d, 0xe8, 0x09, 0xd4, 0x89, 0xf7, 0x0d, 0x19, 0x85, 0x2a, 0xf4,
	0x74, 0x21, 0x59, 0xa7, 0xc4, 0xf5, 0x06, 0x9c, 0x86, 0xc6, 0x7a, 0x5d, 0xdb, 0x2e, 0xe2, 0x5b,
	0x11, 0xee, 0x19, 0xeb, 0x4d, 0x15, 0xc6, 0x93, 0x18, 0x83, 0x30, 0x94, 0x85, 0x17, 0x5a, 0x2c,
	0x50, 0x43, 0xc7, 0xd8, 0x50, 0x99, 0x3d, 0xb8, 0x6a, 0x43, 0x79, 0xe1, 0x8b, 0x88, 0x88, 0x41,
	0x8c, 0x7f, 0xa3, 0xaf, 0xe0, 0x03, 0xc1, 0x89, 0x1f, 0xca, 0x0b, 0xc9, 0x0a, 0x99, 0xfd, 0x86,
	0x0a, 0x4b, 0x8d, 0x7b, 0xcb, 0xe6, 0xae, 0xa0, 0xdc, 0x25, 0xc6, 0xa6, 0x8a, 0xb3, 0x39, 0xb7,
	0xf1, 0x27, 0x6a, 0xee, 0xe2, 0x5b, 0x63, 0xfe, 0x89, 0xa2, 0x3f, 0x97, 0xec, 0xfd, 0x98, 0x5c,
	0xfb, 0xbd, 0x06, 0x85, 0x63, 0x32, 0xf2, 0x18, 0x71, 0xd0, 0xfb, 0x90, 0x13, 0xf4, 0x6d, 0x34,
	0xe8, 0x4a, 0x66, 0xe1, 0xdc, 0xcc, 0xf1, 0x4c, 0x5d, 0xeb, 0x2c, 0x61, 0x65, 0x46, 0x06, 0xe4,
	0x5f, 0xbb, 0x3e, 0xe1, 0x23, 0x35, 0xc7, 0x56, 0x3a, 0x4b, 0x38, 0xfe, 0xde, 0x6d, 0xfe, 0xf9,
	0x1f, 0x7f, 0xb8, 0x7d, 0x1f, 0xb6, 0x2f, 0x58, 0xe9, 0x4e, 0x6a, 0x91, 0x71, 0x24, 0xb3, 0x0a,
	0x85, 0x20, 0x0e, 0x9a, 0xfd, 0xaf, 0xa9, 0xd5, 0xfe, 0xb3, 0x0c, 0xab, 0x33, 0xd7, 0x16, 0xfa,
	0x00, 0x72, 0x7d, 0x16, 0x26, 0xd9, 0x94, 0xcf, 0xcd, 0x22, 0xcf, 0xff, 0x4d, 0xcb, 0x7c, 0xaf,
	0x2d, 0x61, 0xe5, 0x40, 0x77, 0x20, 0xa7, 0xca, 0x24, 0xa3, 0x00, 0x95, 0x73, 0x13, 0x78, 0xb1,
	0xae, 0x25, 0x10, 0xe9, 0x42, 0x26, 0xe4, 0x42, 0xea, 0x3b, 0xf1, 0xb0, 0x6b, 0x5c, 0xf1, 0x28,
	0xe2, 0x2c, 0xb1, 0xe2, 0xa2, 0x0e, 0x14, 0x38, 0xb5, 0xa9, 0x3b, 0xa4, 0xf1, 0xcc, 0xba, 0xae,
	0x4c, 0x42, 0x47, 0x1e, 0xdc, 0xe4, 0xf4, 0x37, 0x03, 0x1a, 0x0a, 0x59, 0x64, 0x0e, 0xe5, 0xa1,
	0x25, 0x98, 0x45, 0x1c, 0xc7, 0xc8, 0xd7, 0xb3, 0xdb, 0xe5, 0xd6, 0xc7, 0xef, 0x16, 0x76, 0x28,
	0x57, 0x43, 0x23, 0x2a, 0x0d, 0xb3, 0x74, 0x6e, 0xe6, 0xbf, 0xd5, 0xb2, 0xfa, 0xbf, 0x0b, 0x78,
	0x3d, 0x96, 0x8d, 0x40, 0x61, 0x97, 0xed, 0x39, 0x0e, 0xfa, 0x02, 0xb6, 0x2e, 0x88, 0xc6, 0xe9,
	0x19, 0x1b, 0x52, 0xa3, 0x58, 0xcf, 0x6e, 0x97, 0x4c, 0xfd, 0xdc, 0xac, 0x7c, 0xab, 0xc1, 0x5d,
	0xb5, 0xb3, 0x9a, 0xdc, 0xb6, 0x9b, 0xb3, 0x4a, 0x58, 0xe1, 0xd1, 0x13, 0x58, 0xa3, 0x6f, 0x03,
	0x6a, 0x0b, 0xea, 0x58, 0xd1, 0x5b, 0x8b, 0x86, 0x46, 0x49, 0x65, 0xbd, 0x15, 0x67, 0x2d, 0x9f,
	0x91, 0xf2, 0xbc, 0x0f, 0x7c, 0xf1, 0xe8, 0x21, 0x96, 0x6f, 0x49, 0xac, 0x27, 0x9c, 0x93, 0x98,
	0x82, 0xba, 0xb0, 0x66, 0x33, 0x87, 0xda, 0x96, 0xed, 0xb9, 0xb2, 0x8b, 0x25, 0xc9, 0x80, 0xba,
	0xb6, 0x5d, 0x6d, 0xdd, 0x9e, 0xd1, 0xd9, 0x97, 0xb8, 0x7d, 0x05, 0xeb, 0x8e, 0x82, 0x68, 0x6a,
	0xff, 0x4e, 0x4d, 0xed, 0x55, 0x3b, 0xed, 0x42, 0x5f, 0xc2, 0x8d, 0xe4, 0x8d, 0xec, 0x93, 0x33,
	0x6a, 0xc5, 0xef, 0xe1, 0x78, 0xdc, 0x7c, 0x34, 0x2d, 0x1c, 0xbb, 0xc6, 0xfb, 0x7a, 0xa2, 0x9e,
	0xca, 0xcf, 0x23, 0x2b, 0x46, 0xb1, 0xcc, 0x11, 0x39, 0xa3, 0xb1, 0x6d, 0xf7, 0x53, 0x59, 0xdc,
	0x2d, 0xf8, 0xe9, 0xc2, 0xe2, 0x9e, 0x29, 0xe0, 0xc3, 0x5c, 0x71, 0x59, 0xcf, 0x1f, 0xe6, 0x8a,
	0x05, 0xbd, 0x88, 0x57, 0xa6, 0xf3, 0xc3, 0x25, 0xf9, 0x4e, 0x90, 0x33, 0xb9, 0x55, 0xfb, 0x97,
	0x06, 0xd5, 0xf4, 0x2c, 0x1d, 0x17, 0xad, 0xf6, 0xc3, 0x14, 0x6d, 0x46, 0x9d, 0xd2, 0xff, 0x5b,
	0xb4, 0xbb, 0x8f, 0xe4, 0xf2, 0x1f, 0x40, 0x73, 0xe1, 0xf2, 0xd3, 0xab, 0xa8, 0x7d, 0x05, 0x3a,
	0xa6, 0x8e, 0x1b, 0x4e, 0xaf, 0x4c, 0x87, 0xec, 0x1b, 0x3a, 0x8a, 0x3a, 0x1a, 0xcb, 0x9f, 0xbb,
	0x9f, 0x49, 0xf5, 0x1d, 0x78, 0xb0, 0x50, 0x7d, 0x56, 0xab, 0xf6, 0x27, 0x0d, 0x56, 0x67, 0x1e,
	0x08, 0xe8, 0x0e, 0xa4, 0xf6, 0x39, 0x0e, 0x54, 0x9e, 0x3a, 0x54, 0xf4, 0x11, 0x94, 0xc8, 0x40,
	0xf4, 0x19, 0x77, 0xc5, 0x28, 0xbe, 0x39, 0x52, 0x57, 0xcb, 0xc4, 0x7b, 0xf5, 0x83, 0x9f, 0x49,
	0xa3, 0xf6, 0x77, 0x0d, 0xd6, 0xe6, 0xde, 0x14, 0xe8, 0x3d, 0xc8, 0x4d, 0x92, 0x1a, 0xdf, 0xae,
	0x58, 0x19, 0xd1, 0xe7, 0xb0, 0x22, 0xcb, 0xd3, 0x89, 0xff, 0xe3, 0xc5, 0x17, 0xd6, 0x8d, 0xb9,
	0x3b, 0x7d, 0xcf, 0x1f, 0x75, 0x96, 0x70, 0x59, 0x61, 0xa3, 0x59, 0xb6, 0xfb, 0xb9, 0x4c, 0xf3,
	0x21, 0xb4, 0x16, 0xa6, 0x39, 0x97, 0x92, 0x59, 0x81, 0x72, 0x84, 0x57, 0x7d, 0x78, 0x98, 0x2b,
	0x66, 0xf4, 0x2c, 0xce, 0x47, 0xa6, 0x5a, 0x0f, 0x60, 0x32, 0x93, 0xd0, 0x8f, 0xa1, 0x4a, 0xbc,
	0xc0, 0xb7, 0x54, 0x26, 0x36, 0xf3, 0x42, 0x43, 0x93, 0x57, 0x08, 0xae, 0x48, 0xeb, 0x71, 0x62,
	0xdc, 0x6d, 0xc9, 0x64, 0x3e, 0x81, 0x9f, 0x2c, 0xae, 0x96, 0xb1, 0xf4, 0xee, 0x3d, 0xc9, 0xb9,
	0x0b, 0xf5, 0x45, 0x1c, 0x73, 0x03, 0xaa, 0xd3, 0xc3, 0x99, 0x72, 0x35, 0x3c, 0x0e, 0x73, 0x45,
	0xd0, 0xcb, 0xf7, 0x7f, 0x0d, 0x2b, 0x11, 0x36, 0xba, 0x6b, 0x50, 0x19, 0x0a, 0x2f, 0x8f, 0xbe,
	0x38, 0x7a, 0xf1, 0xea, 0x48, 0x5f, 0x92, 0x1f, 0x9d, 0xf6, 0xde, 0xb3, 0x6e, 0xe7, 0x57, 0xba,
	0x86, 0x2a, 0x50, 0x7a, 0x79, 0x94, 0x7c, 0x66, 0xd0, 0x0a, 0x14, 0x1f, 0xe3, 0xbd, 0x83, 0xa3,
	0x83, 0xa3, 0xa7, 0x7a, 0x56, 0x22, 0xbb, 0x07, 0xcf, 0xdb, 0x2f, 0x5e, 0x76, 0xf5, 0x9c, 0x72,
	0xb5, 0x9f, 0xe2, 0xbd, 0xc7, 0xed, 0xc7, 0xfa, 0xb2, 0xd9, 0xfe, 0xeb, 0x6f, 0xbf, 0xff, 0x67,
	0x3e, 0xa3, 0x67, 0xe1, 0x9e, 0xcb, 0xa2, 0x76, 0x0a, 0x38, 0x7b, 0x3b, 0xba, 0xbc, 0xb3, 0x4c,
	0x7d, 0x2a, 0x7d, 0xb5, 0x49, 0xc7, 0xda, 0xeb, 0xbc, 0xda, 0xc2, 0x9d, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x73, 0x99, 0xfa, 0xf5, 0x66, 0x10, 0x00, 0x00,
}
