// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: agent.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AgentResponseStatus int32

const (
	AgentResponseStatus_FAIL    AgentResponseStatus = 0 /// Success
	AgentResponseStatus_SUCCESS AgentResponseStatus = 1 /// Failed
)

// Enum value maps for AgentResponseStatus.
var (
	AgentResponseStatus_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
	}
	AgentResponseStatus_value = map[string]int32{
		"FAIL":    0,
		"SUCCESS": 1,
	}
)

func (x AgentResponseStatus) Enum() *AgentResponseStatus {
	p := new(AgentResponseStatus)
	*p = x
	return p
}

func (x AgentResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_agent_proto_enumTypes[0].Descriptor()
}

func (AgentResponseStatus) Type() protoreflect.EnumType {
	return &file_agent_proto_enumTypes[0]
}

func (x AgentResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentResponseStatus.Descriptor instead.
func (AgentResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IpControl          string `protobuf:"bytes,2,opt,name=ip_control,proto3" json:"ip_control,omitempty"`
	IpReceiveMulticast string `protobuf:"bytes,3,opt,name=ip_receive_multicast,proto3" json:"ip_receive_multicast,omitempty"`
	Cpu                int64  `protobuf:"varint,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram                int64  `protobuf:"varint,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Disk               int64  `protobuf:"varint,6,opt,name=disk,proto3" json:"disk,omitempty"`
	Location           string `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	IsMonitor          bool   `protobuf:"varint,8,opt,name=is_monitor,proto3" json:"is_monitor,omitempty"`
	IsAlarm            bool   `protobuf:"varint,9,opt,name=is_alarm,proto3" json:"is_alarm,omitempty"`
	SignalMonitor      bool   `protobuf:"varint,10,opt,name=signal_monitor,proto3" json:"signal_monitor,omitempty"`
	VideoMonitor       bool   `protobuf:"varint,11,opt,name=video_monitor,proto3" json:"video_monitor,omitempty"`
	AudioMonitor       bool   `protobuf:"varint,12,opt,name=audio_monitor,proto3" json:"audio_monitor,omitempty"`
	RunThread          int64  `protobuf:"varint,13,opt,name=run_thread,proto3" json:"run_thread,omitempty"`
	DateCreate         string `protobuf:"bytes,14,opt,name=date_create,proto3" json:"date_create,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

func (x *Agent) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Agent) GetIpControl() string {
	if x != nil {
		return x.IpControl
	}
	return ""
}

func (x *Agent) GetIpReceiveMulticast() string {
	if x != nil {
		return x.IpReceiveMulticast
	}
	return ""
}

func (x *Agent) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Agent) GetRam() int64 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *Agent) GetDisk() int64 {
	if x != nil {
		return x.Disk
	}
	return 0
}

func (x *Agent) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Agent) GetIsMonitor() bool {
	if x != nil {
		return x.IsMonitor
	}
	return false
}

func (x *Agent) GetIsAlarm() bool {
	if x != nil {
		return x.IsAlarm
	}
	return false
}

func (x *Agent) GetSignalMonitor() bool {
	if x != nil {
		return x.SignalMonitor
	}
	return false
}

func (x *Agent) GetVideoMonitor() bool {
	if x != nil {
		return x.VideoMonitor
	}
	return false
}

func (x *Agent) GetAudioMonitor() bool {
	if x != nil {
		return x.AudioMonitor
	}
	return false
}

func (x *Agent) GetRunThread() int64 {
	if x != nil {
		return x.RunThread
	}
	return 0
}

func (x *Agent) GetDateCreate() string {
	if x != nil {
		return x.DateCreate
	}
	return ""
}

//*
// Represents the params to identify agent.
type AgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IpControl          string `protobuf:"bytes,2,opt,name=ip_control,proto3" json:"ip_control,omitempty"`
	IpReceiveMulticast string `protobuf:"bytes,3,opt,name=ip_receive_multicast,proto3" json:"ip_receive_multicast,omitempty"`
	Cpu                int64  `protobuf:"varint,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram                int64  `protobuf:"varint,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Disk               int64  `protobuf:"varint,6,opt,name=disk,proto3" json:"disk,omitempty"`
	Location           string `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	IsMonitor          bool   `protobuf:"varint,8,opt,name=is_monitor,proto3" json:"is_monitor,omitempty"`
	IsAlarm            bool   `protobuf:"varint,9,opt,name=is_alarm,proto3" json:"is_alarm,omitempty"`
	SignalMonitor      bool   `protobuf:"varint,10,opt,name=signal_monitor,proto3" json:"signal_monitor,omitempty"`
	VideoMonitor       bool   `protobuf:"varint,11,opt,name=video_monitor,proto3" json:"video_monitor,omitempty"`
	AudioMonitor       bool   `protobuf:"varint,12,opt,name=audio_monitor,proto3" json:"audio_monitor,omitempty"`
	RunThread          int64  `protobuf:"varint,13,opt,name=run_thread,proto3" json:"run_thread,omitempty"`
}

func (x *AgentRequest) Reset() {
	*x = AgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRequest) ProtoMessage() {}

func (x *AgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRequest.ProtoReflect.Descriptor instead.
func (*AgentRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AgentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AgentRequest) GetIpControl() string {
	if x != nil {
		return x.IpControl
	}
	return ""
}

func (x *AgentRequest) GetIpReceiveMulticast() string {
	if x != nil {
		return x.IpReceiveMulticast
	}
	return ""
}

func (x *AgentRequest) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *AgentRequest) GetRam() int64 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *AgentRequest) GetDisk() int64 {
	if x != nil {
		return x.Disk
	}
	return 0
}

func (x *AgentRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AgentRequest) GetIsMonitor() bool {
	if x != nil {
		return x.IsMonitor
	}
	return false
}

func (x *AgentRequest) GetIsAlarm() bool {
	if x != nil {
		return x.IsAlarm
	}
	return false
}

func (x *AgentRequest) GetSignalMonitor() bool {
	if x != nil {
		return x.SignalMonitor
	}
	return false
}

func (x *AgentRequest) GetVideoMonitor() bool {
	if x != nil {
		return x.VideoMonitor
	}
	return false
}

func (x *AgentRequest) GetAudioMonitor() bool {
	if x != nil {
		return x.AudioMonitor
	}
	return false
}

func (x *AgentRequest) GetRunThread() int64 {
	if x != nil {
		return x.RunThread
	}
	return 0
}

type AgentFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IpControl string `protobuf:"bytes,2,opt,name=ip_control,proto3" json:"ip_control,omitempty"`
	Location  string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *AgentFilter) Reset() {
	*x = AgentFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentFilter) ProtoMessage() {}

func (x *AgentFilter) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentFilter.ProtoReflect.Descriptor instead.
func (*AgentFilter) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *AgentFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AgentFilter) GetIpControl() string {
	if x != nil {
		return x.IpControl
	}
	return ""
}

func (x *AgentFilter) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type AgentGetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AgentGetAll) Reset() {
	*x = AgentGetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentGetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentGetAll) ProtoMessage() {}

func (x *AgentGetAll) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentGetAll.ProtoReflect.Descriptor instead.
func (*AgentGetAll) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

type AgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Status
	Status AgentResponseStatus `protobuf:"varint,1,opt,name=Status,json=status,proto3,enum=proto.AgentResponseStatus" json:"Status,omitempty"`
	//*
	// Slice of agent object
	Agents []*Agent `protobuf:"bytes,2,rep,name=agents,json=data,proto3" json:"agents,omitempty"`
}

func (x *AgentResponse) Reset() {
	*x = AgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentResponse) ProtoMessage() {}

func (x *AgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentResponse.ProtoReflect.Descriptor instead.
func (*AgentResponse) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{4}
}

func (x *AgentResponse) GetStatus() AgentResponseStatus {
	if x != nil {
		return x.Status
	}
	return AgentResponseStatus_FAIL
}

func (x *AgentResponse) GetAgents() []*Agent {
	if x != nil {
		return x.Agents
	}
	return nil
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x03, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x32,
	0x0a, 0x14, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x70,
	0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x63, 0x70, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x5f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x22, 0x96, 0x03, 0x0a, 0x0c, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x70, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x32, 0x0a, 0x14, 0x69, 0x70, 0x5f,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x61,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x26, 0x0a,
	0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x22, 0x59, 0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0d, 0x0a, 0x0b,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x22, 0x67, 0x0a, 0x0d, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x22, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x2a, 0x2c, 0x0a, 0x13, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x32, 0x99, 0x02, 0x0a, 0x0f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x54, 0x4c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x47, 0x65, 0x74, 0x73, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x70, 0x74, 0x76, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_agent_proto_goTypes = []interface{}{
	(AgentResponseStatus)(0), // 0: proto.AgentResponseStatus
	(*Agent)(nil),            // 1: proto.Agent
	(*AgentRequest)(nil),     // 2: proto.AgentRequest
	(*AgentFilter)(nil),      // 3: proto.AgentFilter
	(*AgentGetAll)(nil),      // 4: proto.AgentGetAll
	(*AgentResponse)(nil),    // 5: proto.AgentResponse
}
var file_agent_proto_depIdxs = []int32{
	0, // 0: proto.AgentResponse.Status:type_name -> proto.AgentResponseStatus
	1, // 1: proto.AgentResponse.agents:type_name -> proto.Agent
	4, // 2: proto.AgentCTLService.Gets:input_type -> proto.AgentGetAll
	3, // 3: proto.AgentCTLService.Get:input_type -> proto.AgentFilter
	2, // 4: proto.AgentCTLService.Add:input_type -> proto.AgentRequest
	2, // 5: proto.AgentCTLService.Update:input_type -> proto.AgentRequest
	3, // 6: proto.AgentCTLService.Delete:input_type -> proto.AgentFilter
	5, // 7: proto.AgentCTLService.Gets:output_type -> proto.AgentResponse
	5, // 8: proto.AgentCTLService.Get:output_type -> proto.AgentResponse
	5, // 9: proto.AgentCTLService.Add:output_type -> proto.AgentResponse
	5, // 10: proto.AgentCTLService.Update:output_type -> proto.AgentResponse
	5, // 11: proto.AgentCTLService.Delete:output_type -> proto.AgentResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
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
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRequest); i {
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
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentFilter); i {
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
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentGetAll); i {
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
		file_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		EnumInfos:         file_agent_proto_enumTypes,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentCTLServiceClient is the client API for AgentCTLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentCTLServiceClient interface {
	//*
	// Get All Agent
	Gets(ctx context.Context, in *AgentGetAll, opts ...grpc.CallOption) (*AgentResponse, error)
	//*
	// Get Agent
	Get(ctx context.Context, in *AgentFilter, opts ...grpc.CallOption) (*AgentResponse, error)
	//*
	// Add Agent
	Add(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*AgentResponse, error)
	//*
	// Update Agent
	Update(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*AgentResponse, error)
	//*
	// Delete Agent
	Delete(ctx context.Context, in *AgentFilter, opts ...grpc.CallOption) (*AgentResponse, error)
}

type agentCTLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentCTLServiceClient(cc grpc.ClientConnInterface) AgentCTLServiceClient {
	return &agentCTLServiceClient{cc}
}

func (c *agentCTLServiceClient) Gets(ctx context.Context, in *AgentGetAll, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/Gets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) Get(ctx context.Context, in *AgentFilter, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) Add(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) Update(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentCTLServiceClient) Delete(ctx context.Context, in *AgentFilter, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentCTLService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentCTLServiceServer is the server API for AgentCTLService service.
type AgentCTLServiceServer interface {
	//*
	// Get All Agent
	Gets(context.Context, *AgentGetAll) (*AgentResponse, error)
	//*
	// Get Agent
	Get(context.Context, *AgentFilter) (*AgentResponse, error)
	//*
	// Add Agent
	Add(context.Context, *AgentRequest) (*AgentResponse, error)
	//*
	// Update Agent
	Update(context.Context, *AgentRequest) (*AgentResponse, error)
	//*
	// Delete Agent
	Delete(context.Context, *AgentFilter) (*AgentResponse, error)
}

// UnimplementedAgentCTLServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentCTLServiceServer struct {
}

func (*UnimplementedAgentCTLServiceServer) Gets(context.Context, *AgentGetAll) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gets not implemented")
}
func (*UnimplementedAgentCTLServiceServer) Get(context.Context, *AgentFilter) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAgentCTLServiceServer) Add(context.Context, *AgentRequest) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedAgentCTLServiceServer) Update(context.Context, *AgentRequest) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAgentCTLServiceServer) Delete(context.Context, *AgentFilter) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAgentCTLServiceServer(s *grpc.Server, srv AgentCTLServiceServer) {
	s.RegisterService(&_AgentCTLService_serviceDesc, srv)
}

func _AgentCTLService_Gets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentGetAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).Gets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).Gets(ctx, req.(*AgentGetAll))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).Get(ctx, req.(*AgentFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).Add(ctx, req.(*AgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).Update(ctx, req.(*AgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentCTLService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentCTLServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentCTLService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentCTLServiceServer).Delete(ctx, req.(*AgentFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentCTLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AgentCTLService",
	HandlerType: (*AgentCTLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Gets",
			Handler:    _AgentCTLService_Gets_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AgentCTLService_Get_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _AgentCTLService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AgentCTLService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AgentCTLService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
