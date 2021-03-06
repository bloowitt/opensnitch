// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ui.proto

package protocol

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Action int32

const (
	Action_NONE            Action = 0
	Action_LOAD_FIREWALL   Action = 1
	Action_UNLOAD_FIREWALL Action = 2
	Action_CHANGE_CONFIG   Action = 3
	Action_ENABLE_RULE     Action = 4
	Action_DISABLE_RULE    Action = 5
	Action_DELETE_RULE     Action = 6
	Action_CHANGE_RULE     Action = 7
	Action_LOG_LEVEL       Action = 8
	Action_STOP            Action = 9
)

var Action_name = map[int32]string{
	0: "NONE",
	1: "LOAD_FIREWALL",
	2: "UNLOAD_FIREWALL",
	3: "CHANGE_CONFIG",
	4: "ENABLE_RULE",
	5: "DISABLE_RULE",
	6: "DELETE_RULE",
	7: "CHANGE_RULE",
	8: "LOG_LEVEL",
	9: "STOP",
}

var Action_value = map[string]int32{
	"NONE":            0,
	"LOAD_FIREWALL":   1,
	"UNLOAD_FIREWALL": 2,
	"CHANGE_CONFIG":   3,
	"ENABLE_RULE":     4,
	"DISABLE_RULE":    5,
	"DELETE_RULE":     6,
	"CHANGE_RULE":     7,
	"LOG_LEVEL":       8,
	"STOP":            9,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{0}
}

type NotificationReplyCode int32

const (
	NotificationReplyCode_OK    NotificationReplyCode = 0
	NotificationReplyCode_ERROR NotificationReplyCode = 1
)

var NotificationReplyCode_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}

var NotificationReplyCode_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x NotificationReplyCode) String() string {
	return proto.EnumName(NotificationReplyCode_name, int32(x))
}

func (NotificationReplyCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{1}
}

type Event struct {
	Time       string      `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Connection *Connection `protobuf:"bytes,2,opt,name=connection,proto3" json:"connection,omitempty"`
	Rule       *Rule       `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{0}
}

func (m *Event) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Event) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *Event) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

type Statistics struct {
	DaemonVersion string            `protobuf:"bytes,1,opt,name=daemon_version,json=daemonVersion,proto3" json:"daemon_version,omitempty"`
	Rules         uint64            `protobuf:"varint,2,opt,name=rules,proto3" json:"rules,omitempty"`
	Uptime        uint64            `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	DnsResponses  uint64            `protobuf:"varint,4,opt,name=dns_responses,json=dnsResponses,proto3" json:"dns_responses,omitempty"`
	Connections   uint64            `protobuf:"varint,5,opt,name=connections,proto3" json:"connections,omitempty"`
	Ignored       uint64            `protobuf:"varint,6,opt,name=ignored,proto3" json:"ignored,omitempty"`
	Accepted      uint64            `protobuf:"varint,7,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Dropped       uint64            `protobuf:"varint,8,opt,name=dropped,proto3" json:"dropped,omitempty"`
	RuleHits      uint64            `protobuf:"varint,9,opt,name=rule_hits,json=ruleHits,proto3" json:"rule_hits,omitempty"`
	RuleMisses    uint64            `protobuf:"varint,10,opt,name=rule_misses,json=ruleMisses,proto3" json:"rule_misses,omitempty"`
	ByProto       map[string]uint64 `protobuf:"bytes,11,rep,name=by_proto,json=byProto,proto3" json:"by_proto,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ByAddress     map[string]uint64 `protobuf:"bytes,12,rep,name=by_address,json=byAddress,proto3" json:"by_address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ByHost        map[string]uint64 `protobuf:"bytes,13,rep,name=by_host,json=byHost,proto3" json:"by_host,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ByPort        map[string]uint64 `protobuf:"bytes,14,rep,name=by_port,json=byPort,proto3" json:"by_port,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ByUid         map[string]uint64 `protobuf:"bytes,15,rep,name=by_uid,json=byUid,proto3" json:"by_uid,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ByExecutable  map[string]uint64 `protobuf:"bytes,16,rep,name=by_executable,json=byExecutable,proto3" json:"by_executable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Events        []*Event          `protobuf:"bytes,17,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *Statistics) Reset()         { *m = Statistics{} }
func (m *Statistics) String() string { return proto.CompactTextString(m) }
func (*Statistics) ProtoMessage()    {}
func (*Statistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{1}
}

func (m *Statistics) GetDaemonVersion() string {
	if m != nil {
		return m.DaemonVersion
	}
	return ""
}

func (m *Statistics) GetRules() uint64 {
	if m != nil {
		return m.Rules
	}
	return 0
}

func (m *Statistics) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *Statistics) GetDnsResponses() uint64 {
	if m != nil {
		return m.DnsResponses
	}
	return 0
}

func (m *Statistics) GetConnections() uint64 {
	if m != nil {
		return m.Connections
	}
	return 0
}

func (m *Statistics) GetIgnored() uint64 {
	if m != nil {
		return m.Ignored
	}
	return 0
}

func (m *Statistics) GetAccepted() uint64 {
	if m != nil {
		return m.Accepted
	}
	return 0
}

func (m *Statistics) GetDropped() uint64 {
	if m != nil {
		return m.Dropped
	}
	return 0
}

func (m *Statistics) GetRuleHits() uint64 {
	if m != nil {
		return m.RuleHits
	}
	return 0
}

func (m *Statistics) GetRuleMisses() uint64 {
	if m != nil {
		return m.RuleMisses
	}
	return 0
}

func (m *Statistics) GetByProto() map[string]uint64 {
	if m != nil {
		return m.ByProto
	}
	return nil
}

func (m *Statistics) GetByAddress() map[string]uint64 {
	if m != nil {
		return m.ByAddress
	}
	return nil
}

func (m *Statistics) GetByHost() map[string]uint64 {
	if m != nil {
		return m.ByHost
	}
	return nil
}

func (m *Statistics) GetByPort() map[string]uint64 {
	if m != nil {
		return m.ByPort
	}
	return nil
}

func (m *Statistics) GetByUid() map[string]uint64 {
	if m != nil {
		return m.ByUid
	}
	return nil
}

func (m *Statistics) GetByExecutable() map[string]uint64 {
	if m != nil {
		return m.ByExecutable
	}
	return nil
}

func (m *Statistics) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type PingRequest struct {
	Id    uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Stats *Statistics `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{2}
}

func (m *PingRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PingRequest) GetStats() *Statistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

type PingReply struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{3}
}

func (m *PingReply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Connection struct {
	Protocol    string            `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SrcIp       string            `protobuf:"bytes,2,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	SrcPort     uint32            `protobuf:"varint,3,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstIp       string            `protobuf:"bytes,4,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	DstHost     string            `protobuf:"bytes,5,opt,name=dst_host,json=dstHost,proto3" json:"dst_host,omitempty"`
	DstPort     uint32            `protobuf:"varint,6,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	UserId      uint32            `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProcessId   uint32            `protobuf:"varint,8,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	ProcessPath string            `protobuf:"bytes,9,opt,name=process_path,json=processPath,proto3" json:"process_path,omitempty"`
	ProcessCwd  string            `protobuf:"bytes,10,opt,name=process_cwd,json=processCwd,proto3" json:"process_cwd,omitempty"`
	ProcessArgs []string          `protobuf:"bytes,11,rep,name=process_args,json=processArgs,proto3" json:"process_args,omitempty"`
	ProcessEnv  map[string]string `protobuf:"bytes,12,rep,name=process_env,json=processEnv,proto3" json:"process_env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{4}
}

func (m *Connection) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Connection) GetSrcIp() string {
	if m != nil {
		return m.SrcIp
	}
	return ""
}

func (m *Connection) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *Connection) GetDstIp() string {
	if m != nil {
		return m.DstIp
	}
	return ""
}

func (m *Connection) GetDstHost() string {
	if m != nil {
		return m.DstHost
	}
	return ""
}

func (m *Connection) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *Connection) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Connection) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *Connection) GetProcessPath() string {
	if m != nil {
		return m.ProcessPath
	}
	return ""
}

func (m *Connection) GetProcessCwd() string {
	if m != nil {
		return m.ProcessCwd
	}
	return ""
}

func (m *Connection) GetProcessArgs() []string {
	if m != nil {
		return m.ProcessArgs
	}
	return nil
}

func (m *Connection) GetProcessEnv() map[string]string {
	if m != nil {
		return m.ProcessEnv
	}
	return nil
}

type Operator struct {
	Type      string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Operand   string `protobuf:"bytes,2,opt,name=operand,proto3" json:"operand,omitempty"`
	Data      string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Sensitive bool   `protobuf:"varint,4,opt,name=sensitive,proto3" json:"sensitive,omitempty"`
}

func (m *Operator) Reset()         { *m = Operator{} }
func (m *Operator) String() string { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()    {}
func (*Operator) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{5}
}

func (m *Operator) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Operator) GetOperand() string {
	if m != nil {
		return m.Operand
	}
	return ""
}

func (m *Operator) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Operator) GetSensitive() bool {
	if m != nil {
		return m.Sensitive
	}
	return false
}

type Rule struct {
	Name       string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled    bool      `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Precedence bool      `protobuf:"varint,3,opt,name=precedence,proto3" json:"precedence,omitempty"`
	Action     string    `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	Duration   string    `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	Operator   *Operator `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{6}
}

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Rule) GetPrecedence() bool {
	if m != nil {
		return m.Precedence
	}
	return false
}

func (m *Rule) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Rule) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Rule) GetOperator() *Operator {
	if m != nil {
		return m.Operator
	}
	return nil
}

// client configuration sent on Subscribe()
type ClientConfig struct {
	Id                uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version           string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	IsFirewallRunning bool   `protobuf:"varint,4,opt,name=isFirewallRunning,proto3" json:"isFirewallRunning,omitempty"`
	// daemon configuration as json string
	Config   string  `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	LogLevel uint32  `protobuf:"varint,6,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
	Rules    []*Rule `protobuf:"bytes,7,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (m *ClientConfig) Reset()         { *m = ClientConfig{} }
func (m *ClientConfig) String() string { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()    {}
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{7}
}

func (m *ClientConfig) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ClientConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClientConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ClientConfig) GetIsFirewallRunning() bool {
	if m != nil {
		return m.IsFirewallRunning
	}
	return false
}

func (m *ClientConfig) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *ClientConfig) GetLogLevel() uint32 {
	if m != nil {
		return m.LogLevel
	}
	return 0
}

func (m *ClientConfig) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// notification sent to the clients (daemons)
type Notification struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientName string `protobuf:"bytes,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
	ServerName string `protobuf:"bytes,3,opt,name=serverName,proto3" json:"serverName,omitempty"`
	// CHANGE_CONFIG: 2, data: {"default_timeout": 1, ...}
	Type  Action  `protobuf:"varint,4,opt,name=type,proto3,enum=protocol.Action" json:"type,omitempty"`
	Data  string  `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Rules []*Rule `protobuf:"bytes,6,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{8}
}

func (m *Notification) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Notification) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Notification) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *Notification) GetType() Action {
	if m != nil {
		return m.Type
	}
	return Action_NONE
}

func (m *Notification) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Notification) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// notification reply sent to the server (GUI)
type NotificationReply struct {
	Id   uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code NotificationReplyCode `protobuf:"varint,2,opt,name=code,proto3,enum=protocol.NotificationReplyCode" json:"code,omitempty"`
	Data string                `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *NotificationReply) Reset()         { *m = NotificationReply{} }
func (m *NotificationReply) String() string { return proto.CompactTextString(m) }
func (*NotificationReply) ProtoMessage()    {}
func (*NotificationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{9}
}

func (m *NotificationReply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NotificationReply) GetCode() NotificationReplyCode {
	if m != nil {
		return m.Code
	}
	return NotificationReplyCode_OK
}

func (m *NotificationReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterEnum("protocol.Action", Action_name, Action_value)
	proto.RegisterEnum("protocol.NotificationReplyCode", NotificationReplyCode_name, NotificationReplyCode_value)
	proto.RegisterType((*Event)(nil), "protocol.Event")
	proto.RegisterType((*Statistics)(nil), "protocol.Statistics")
	proto.RegisterType((*PingRequest)(nil), "protocol.PingRequest")
	proto.RegisterType((*PingReply)(nil), "protocol.PingReply")
	proto.RegisterType((*Connection)(nil), "protocol.Connection")
	proto.RegisterType((*Operator)(nil), "protocol.Operator")
	proto.RegisterType((*Rule)(nil), "protocol.Rule")
	proto.RegisterType((*ClientConfig)(nil), "protocol.ClientConfig")
	proto.RegisterType((*Notification)(nil), "protocol.Notification")
	proto.RegisterType((*NotificationReply)(nil), "protocol.NotificationReply")
}

func init() {
	proto.RegisterFile("ui.proto", fileDescriptor_63867a62624c1283)
}

var fileDescriptor_63867a62624c1283 = []byte{
	// 1292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdf, 0x73, 0xd3, 0xc6,
	0x13, 0x8f, 0x1d, 0x5b, 0x96, 0xd6, 0x76, 0xe2, 0x1c, 0x84, 0xaf, 0xbe, 0xa6, 0x85, 0x60, 0x68,
	0x9b, 0xc9, 0x74, 0x32, 0x6d, 0x60, 0x3a, 0xc0, 0xc0, 0x74, 0x8c, 0x11, 0xe0, 0xe2, 0xda, 0x99,
	0x0b, 0xa1, 0x8f, 0x1a, 0x49, 0x77, 0x38, 0x57, 0x84, 0xa4, 0xea, 0x4e, 0xa6, 0xfe, 0x9b, 0xfa,
	0x0f, 0xf4, 0xa9, 0xaf, 0x7d, 0xe9, 0x63, 0xff, 0x98, 0x3e, 0x76, 0xee, 0x4e, 0xb2, 0x94, 0x5f,
	0x74, 0xf2, 0x64, 0xed, 0x67, 0xf7, 0xb3, 0xb7, 0xbb, 0xb7, 0xb7, 0x6b, 0x30, 0x33, 0xb6, 0x9f,
	0xa4, 0xb1, 0x88, 0x91, 0xa9, 0x7e, 0x82, 0x38, 0x1c, 0x64, 0xd0, 0x74, 0x16, 0x34, 0x12, 0x08,
	0x41, 0x43, 0xb0, 0x0f, 0xd4, 0xae, 0xed, 0xd4, 0x76, 0x2d, 0xac, 0xbe, 0xd1, 0x03, 0x80, 0x20,
	0x8e, 0x22, 0x1a, 0x08, 0x16, 0x47, 0x76, 0x7d, 0xa7, 0xb6, 0xdb, 0x3e, 0xb8, 0xbe, 0x5f, 0x70,
	0xf7, 0x47, 0x2b, 0x1d, 0xae, 0xd8, 0xa1, 0x01, 0x34, 0xd2, 0x2c, 0xa4, 0xf6, 0xba, 0xb2, 0xdf,
	0x28, 0xed, 0x71, 0x16, 0x52, 0xac, 0x74, 0x83, 0x3f, 0x4d, 0x80, 0x23, 0xe1, 0x09, 0xc6, 0x05,
	0x0b, 0x38, 0xfa, 0x02, 0x36, 0x88, 0x47, 0x3f, 0xc4, 0x91, 0xbb, 0xa0, 0x29, 0x97, 0x87, 0xe9,
	0x30, 0xba, 0x1a, 0x7d, 0xab, 0x41, 0x74, 0x1d, 0x9a, 0x92, 0xcd, 0x55, 0x28, 0x0d, 0xac, 0x05,
	0x74, 0x03, 0x8c, 0x2c, 0x51, 0xb1, 0xaf, 0x2b, 0x38, 0x97, 0xd0, 0x5d, 0xe8, 0x92, 0x88, 0xbb,
	0x29, 0xe5, 0x49, 0x1c, 0x71, 0xca, 0xed, 0x86, 0x52, 0x77, 0x48, 0xc4, 0x71, 0x81, 0xa1, 0x1d,
	0x68, 0x97, 0xa1, 0x73, 0xbb, 0xa9, 0x4c, 0xaa, 0x10, 0xb2, 0xa1, 0xc5, 0xe6, 0x51, 0x9c, 0x52,
	0x62, 0x1b, 0x4a, 0x5b, 0x88, 0xa8, 0x0f, 0xa6, 0x17, 0x04, 0x34, 0x11, 0x94, 0xd8, 0x2d, 0xa5,
	0x5a, 0xc9, 0x92, 0x45, 0xd2, 0x38, 0x49, 0x28, 0xb1, 0x4d, 0xcd, 0xca, 0x45, 0x74, 0x13, 0x2c,
	0x19, 0xb7, 0x7b, 0xc2, 0x04, 0xb7, 0x2d, 0x4d, 0x93, 0xc0, 0x2b, 0x26, 0x38, 0xba, 0x0d, 0x6d,
	0xa5, 0xfc, 0xc0, 0xb8, 0x8c, 0x18, 0x94, 0x1a, 0x24, 0xf4, 0xa3, 0x42, 0xd0, 0x13, 0x30, 0xfd,
	0xa5, 0xab, 0x4a, 0x6a, 0xb7, 0x77, 0xd6, 0x77, 0xdb, 0x07, 0x77, 0xca, 0x02, 0x97, 0x15, 0xdd,
	0x7f, 0xb6, 0x3c, 0x94, 0xa8, 0x13, 0x89, 0x74, 0x89, 0x5b, 0xbe, 0x96, 0xd0, 0x33, 0x00, 0x7f,
	0xe9, 0x7a, 0x84, 0xa4, 0x94, 0x73, 0xbb, 0xa3, 0xf8, 0x77, 0x2f, 0xe1, 0x0f, 0xb5, 0x95, 0xf6,
	0x60, 0xf9, 0x85, 0x8c, 0x1e, 0x41, 0xcb, 0x5f, 0xba, 0x27, 0x31, 0x17, 0x76, 0x57, 0x39, 0xd8,
	0xb9, 0xc4, 0xc1, 0xab, 0x98, 0x0b, 0xcd, 0x36, 0x7c, 0x25, 0xe4, 0xd4, 0x24, 0x4e, 0x85, 0xbd,
	0xf1, 0x49, 0xea, 0x61, 0x9c, 0x96, 0x54, 0x29, 0xa0, 0xef, 0xc0, 0xf0, 0x97, 0x6e, 0xc6, 0x88,
	0xbd, 0xa9, 0x98, 0xb7, 0x2f, 0x61, 0x1e, 0x33, 0xa2, 0x89, 0x4d, 0x5f, 0x7e, 0xa3, 0xd7, 0xd0,
	0xf5, 0x97, 0x2e, 0xfd, 0x95, 0x06, 0x99, 0xf0, 0xfc, 0x90, 0xda, 0x3d, 0x45, 0xff, 0xf2, 0x12,
	0xba, 0xb3, 0x32, 0xd4, 0x5e, 0x3a, 0x7e, 0x05, 0x42, 0x5f, 0x81, 0x41, 0xe5, 0x63, 0xe1, 0xf6,
	0x96, 0xf2, 0xb2, 0x59, 0x7a, 0x51, 0x8f, 0x08, 0xe7, 0xea, 0xfe, 0x63, 0xe8, 0x54, 0x2f, 0x00,
	0xf5, 0x60, 0xfd, 0x3d, 0x5d, 0xe6, 0x4d, 0x2d, 0x3f, 0x65, 0x2b, 0x2f, 0xbc, 0x30, 0xa3, 0x45,
	0x2b, 0x2b, 0xe1, 0x71, 0xfd, 0x61, 0xad, 0xff, 0x04, 0x36, 0x4e, 0x17, 0xff, 0x4a, 0xec, 0x47,
	0xd0, 0xae, 0x54, 0xfe, 0xea, 0xd4, 0x55, 0xe5, 0xaf, 0x44, 0x7d, 0x08, 0x50, 0x96, 0xfe, 0x4a,
	0xcc, 0xef, 0x61, 0xeb, 0x5c, 0xd5, 0xaf, 0xe2, 0x60, 0x30, 0x86, 0xf6, 0x21, 0x8b, 0xe6, 0x98,
	0xfe, 0x92, 0x51, 0x2e, 0xd0, 0x06, 0xd4, 0x19, 0x51, 0xcc, 0x06, 0xae, 0x33, 0x82, 0xf6, 0xa0,
	0xc9, 0x85, 0x27, 0xf8, 0xf9, 0xe9, 0x55, 0xde, 0x3b, 0xd6, 0x26, 0x83, 0x9b, 0x60, 0x69, 0x57,
	0x49, 0xb8, 0x3c, 0xeb, 0x68, 0xf0, 0xd7, 0x3a, 0x40, 0x39, 0xf0, 0xe4, 0xdb, 0x2f, 0x3c, 0xe5,
	0x71, 0xae, 0x64, 0xb4, 0x0d, 0x06, 0x4f, 0x03, 0x97, 0x25, 0xea, 0x50, 0x0b, 0x37, 0x79, 0x1a,
	0x8c, 0x13, 0xf4, 0x7f, 0x30, 0x25, 0xac, 0xda, 0x5f, 0x4e, 0xaa, 0x2e, 0x6e, 0xf1, 0x34, 0x50,
	0xdd, 0xbd, 0x0d, 0x06, 0xe1, 0x42, 0x32, 0x1a, 0x9a, 0x41, 0xb8, 0xd0, 0x0c, 0x09, 0xab, 0xb7,
	0xd6, 0x54, 0x8a, 0x16, 0xe1, 0x42, 0x3d, 0xa5, 0x5c, 0xa5, 0x9c, 0x19, 0xda, 0x19, 0xe1, 0x42,
	0x39, 0xfb, 0x1f, 0xb4, 0x32, 0x4e, 0x53, 0x97, 0xe9, 0xa9, 0xd4, 0xc5, 0x86, 0x14, 0xc7, 0x04,
	0x7d, 0x0e, 0x90, 0xa4, 0x71, 0x40, 0x39, 0x97, 0x3a, 0x53, 0xe9, 0xac, 0x1c, 0x19, 0x13, 0x74,
	0x07, 0x3a, 0x85, 0x3a, 0xf1, 0xc4, 0x89, 0x9a, 0x4d, 0x16, 0x6e, 0xe7, 0xd8, 0xa1, 0x27, 0x4e,
	0xe4, 0x78, 0x2a, 0x4c, 0x82, 0x8f, 0x44, 0x8d, 0x27, 0x0b, 0x17, 0x4e, 0x47, 0x1f, 0x4f, 0xf9,
	0xf0, 0xd2, 0x39, 0x57, 0x23, 0xaa, 0xf4, 0x31, 0x4c, 0xe7, 0x1c, 0x39, 0xa5, 0x0f, 0x1a, 0x2d,
	0xf2, 0x21, 0x74, 0xef, 0xa2, 0xad, 0xb2, 0x7f, 0xa8, 0xed, 0x9c, 0x68, 0xa1, 0x5f, 0x63, 0x71,
	0x92, 0x13, 0x2d, 0xfa, 0x4f, 0x61, 0xf3, 0x8c, 0xfa, 0xbf, 0xda, 0xc6, 0xaa, 0xb6, 0xcd, 0xcf,
	0x60, 0xce, 0x12, 0x9a, 0x7a, 0x22, 0x4e, 0xd5, 0xea, 0x5b, 0x26, 0xe5, 0xea, 0x5b, 0x26, 0x54,
	0xce, 0xef, 0x58, 0xea, 0x23, 0x92, 0x73, 0x0b, 0x51, 0x5a, 0x13, 0x4f, 0x78, 0xea, 0x0a, 0x2d,
	0xac, 0xbe, 0xd1, 0x67, 0x60, 0x71, 0x1a, 0x71, 0x26, 0xd8, 0x82, 0xaa, 0x2b, 0x34, 0x71, 0x09,
	0x0c, 0x7e, 0xaf, 0x41, 0x43, 0xee, 0x3e, 0x49, 0x8d, 0xbc, 0x72, 0xc7, 0xca, 0x6f, 0x79, 0x10,
	0x8d, 0x64, 0xeb, 0xeb, 0x83, 0x4c, 0x5c, 0x88, 0xe8, 0x96, 0xbc, 0x2e, 0x1a, 0x50, 0x42, 0xa3,
	0x40, 0xef, 0x36, 0x13, 0x57, 0x10, 0xb9, 0xf7, 0x3c, 0xbd, 0x99, 0x75, 0xd3, 0xe4, 0x92, 0x6c,
	0x4d, 0x92, 0xa5, 0x9e, 0xd2, 0xe8, 0xae, 0x59, 0xc9, 0x68, 0x1f, 0xcc, 0x38, 0x4f, 0x5b, 0xb5,
	0x4d, 0xfb, 0x00, 0x95, 0x95, 0x2f, 0x0a, 0x82, 0x57, 0x36, 0x83, 0xbf, 0x6b, 0xd0, 0x19, 0x85,
	0x8c, 0x46, 0x62, 0x14, 0x47, 0xef, 0xd8, 0xfc, 0xdc, 0xfb, 0x2a, 0x52, 0xaa, 0x9f, 0x4e, 0xa9,
	0x58, 0xe3, 0xba, 0x48, 0x85, 0x88, 0xbe, 0x86, 0x2d, 0xc6, 0x5f, 0xb0, 0x94, 0x7e, 0xf4, 0xc2,
	0x10, 0x67, 0x51, 0xc4, 0xa2, 0x79, 0x5e, 0xaf, 0xf3, 0x0a, 0x99, 0x60, 0xa0, 0x4e, 0xcd, 0xd3,
	0xc8, 0x25, 0x99, 0x60, 0x18, 0xcf, 0x27, 0x74, 0x41, 0xc3, 0xbc, 0xf7, 0x57, 0x32, 0xba, 0x57,
	0xfc, 0x45, 0x68, 0xa9, 0xbe, 0x3a, 0xfb, 0xef, 0x43, 0x2b, 0x07, 0x7f, 0xd4, 0xa0, 0x33, 0x8d,
	0x05, 0x7b, 0xc7, 0x02, 0x5d, 0x97, 0xb3, 0x69, 0xdd, 0x02, 0x08, 0x54, 0xda, 0xd3, 0x32, 0xb9,
	0x0a, 0x22, 0xf5, 0x9c, 0xa6, 0x0b, 0x9a, 0x2a, 0xbd, 0xce, 0xb2, 0x82, 0xa0, 0x7b, 0x79, 0x4b,
	0xc9, 0xdc, 0x36, 0x0e, 0x7a, 0x65, 0x14, 0x43, 0xfd, 0x7f, 0x49, 0x37, 0x59, 0xd1, 0x4a, 0xcd,
	0x4a, 0x2b, 0xad, 0x12, 0x30, 0x3e, 0x95, 0x40, 0x08, 0x5b, 0xd5, 0xf8, 0x2f, 0x1c, 0x59, 0xe8,
	0x3e, 0x34, 0x82, 0x98, 0xe8, 0xf0, 0x37, 0xaa, 0x1b, 0xf3, 0x1c, 0x75, 0x14, 0x13, 0x8a, 0x95,
	0xf1, 0x45, 0xed, 0xbd, 0xf7, 0x5b, 0x0d, 0x0c, 0x1d, 0x38, 0x32, 0xa1, 0x31, 0x9d, 0x4d, 0x9d,
	0xde, 0x1a, 0xda, 0x82, 0xee, 0x64, 0x36, 0x7c, 0xee, 0xbe, 0x18, 0x63, 0xe7, 0xa7, 0xe1, 0x64,
	0xd2, 0xab, 0xa1, 0x6b, 0xb0, 0x79, 0x3c, 0x3d, 0x0d, 0xd6, 0xa5, 0xdd, 0xe8, 0xd5, 0x70, 0xfa,
	0xd2, 0x71, 0x47, 0xb3, 0xe9, 0x8b, 0xf1, 0xcb, 0xde, 0x3a, 0xda, 0x84, 0xb6, 0x33, 0x1d, 0x3e,
	0x9b, 0x38, 0x2e, 0x3e, 0x9e, 0x38, 0xbd, 0x06, 0xea, 0x41, 0xe7, 0xf9, 0xf8, 0xa8, 0x44, 0x9a,
	0xd2, 0xe4, 0xb9, 0x33, 0x71, 0xde, 0xe4, 0x80, 0x21, 0x81, 0xdc, 0x8d, 0x02, 0x5a, 0xa8, 0x0b,
	0xd6, 0x64, 0xf6, 0xd2, 0x9d, 0x38, 0x6f, 0x9d, 0x49, 0xcf, 0x94, 0x81, 0x1d, 0xbd, 0x99, 0x1d,
	0xf6, 0xac, 0xbd, 0x3d, 0xd8, 0xbe, 0x30, 0x41, 0x64, 0x40, 0x7d, 0xf6, 0xba, 0xb7, 0x86, 0x2c,
	0x68, 0x3a, 0x18, 0xcf, 0x70, 0xaf, 0x76, 0xf0, 0x4f, 0x0d, 0xea, 0xc7, 0x63, 0xf4, 0x00, 0x1a,
	0x72, 0xf2, 0xa3, 0xed, 0xb2, 0x46, 0x95, 0xa5, 0xd2, 0xbf, 0x76, 0x16, 0x4e, 0xc2, 0xe5, 0x60,
	0x0d, 0x7d, 0x0b, 0xad, 0x21, 0x7f, 0xaf, 0x5e, 0xf6, 0x85, 0xff, 0x8a, 0xfb, 0x67, 0x2e, 0x6f,
	0xb0, 0x86, 0x9e, 0x82, 0x75, 0x94, 0xf9, 0x3c, 0x48, 0x99, 0x4f, 0xd1, 0x8d, 0x0a, 0xa9, 0xf2,
	0xc6, 0xfa, 0x97, 0xe0, 0x83, 0x35, 0xf4, 0x03, 0x74, 0xab, 0xa9, 0x71, 0x74, 0xf3, 0x13, 0x97,
	0x5a, 0xf5, 0x53, 0x55, 0x0e, 0xd6, 0x76, 0x6b, 0xdf, 0xd4, 0x7c, 0x43, 0x29, 0xef, 0xff, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x75, 0xf5, 0x8c, 0x39, 0x16, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UIClient is the client API for UI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UIClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	AskRule(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Rule, error)
	Subscribe(ctx context.Context, in *ClientConfig, opts ...grpc.CallOption) (*ClientConfig, error)
	Notifications(ctx context.Context, opts ...grpc.CallOption) (UI_NotificationsClient, error)
}

type uIClient struct {
	cc *grpc.ClientConn
}

func NewUIClient(cc *grpc.ClientConn) UIClient {
	return &uIClient{cc}
}

func (c *uIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/protocol.UI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) AskRule(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, "/protocol.UI/AskRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) Subscribe(ctx context.Context, in *ClientConfig, opts ...grpc.CallOption) (*ClientConfig, error) {
	out := new(ClientConfig)
	err := c.cc.Invoke(ctx, "/protocol.UI/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) Notifications(ctx context.Context, opts ...grpc.CallOption) (UI_NotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UI_serviceDesc.Streams[0], "/protocol.UI/Notifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &uINotificationsClient{stream}
	return x, nil
}

type UI_NotificationsClient interface {
	Send(*NotificationReply) error
	Recv() (*Notification, error)
	grpc.ClientStream
}

type uINotificationsClient struct {
	grpc.ClientStream
}

func (x *uINotificationsClient) Send(m *NotificationReply) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uINotificationsClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UIServer is the server API for UI service.
type UIServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	AskRule(context.Context, *Connection) (*Rule, error)
	Subscribe(context.Context, *ClientConfig) (*ClientConfig, error)
	Notifications(UI_NotificationsServer) error
}

// UnimplementedUIServer can be embedded to have forward compatible implementations.
type UnimplementedUIServer struct {
}

func (*UnimplementedUIServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedUIServer) AskRule(ctx context.Context, req *Connection) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskRule not implemented")
}
func (*UnimplementedUIServer) Subscribe(ctx context.Context, req *ClientConfig) (*ClientConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedUIServer) Notifications(srv UI_NotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method Notifications not implemented")
}

func RegisterUIServer(s *grpc.Server, srv UIServer) {
	s.RegisterService(&_UI_serviceDesc, srv)
}

func _UI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.UI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_AskRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).AskRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.UI/AskRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).AskRule(ctx, req.(*Connection))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.UI/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Subscribe(ctx, req.(*ClientConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_Notifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UIServer).Notifications(&uINotificationsServer{stream})
}

type UI_NotificationsServer interface {
	Send(*Notification) error
	Recv() (*NotificationReply, error)
	grpc.ServerStream
}

type uINotificationsServer struct {
	grpc.ServerStream
}

func (x *uINotificationsServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uINotificationsServer) Recv() (*NotificationReply, error) {
	m := new(NotificationReply)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.UI",
	HandlerType: (*UIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UI_Ping_Handler,
		},
		{
			MethodName: "AskRule",
			Handler:    _UI_AskRule_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _UI_Subscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Notifications",
			Handler:       _UI_Notifications_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ui.proto",
}
