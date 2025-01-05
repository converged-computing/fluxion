// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: pkg/fluxion-grpc/fluxion.proto

package fluxion_grpc

import (
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

type InitResponse_ResultType int32

const (
	InitResponse_INIT_SUCCESS InitResponse_ResultType = 0
	InitResponse_INIT_ERROR   InitResponse_ResultType = 1
	InitResponse_INIT_DENIED  InitResponse_ResultType = 2
)

// Enum value maps for InitResponse_ResultType.
var (
	InitResponse_ResultType_name = map[int32]string{
		0: "INIT_SUCCESS",
		1: "INIT_ERROR",
		2: "INIT_DENIED",
	}
	InitResponse_ResultType_value = map[string]int32{
		"INIT_SUCCESS": 0,
		"INIT_ERROR":   1,
		"INIT_DENIED":  2,
	}
)

func (x InitResponse_ResultType) Enum() *InitResponse_ResultType {
	p := new(InitResponse_ResultType)
	*p = x
	return p
}

func (x InitResponse_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InitResponse_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_fluxion_grpc_fluxion_proto_enumTypes[0].Descriptor()
}

func (InitResponse_ResultType) Type() protoreflect.EnumType {
	return &file_pkg_fluxion_grpc_fluxion_proto_enumTypes[0]
}

func (x InitResponse_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InitResponse_ResultType.Descriptor instead.
func (InitResponse_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{1, 0}
}

type MatchResponse_ResultType int32

const (
	MatchResponse_MATCH_SUCCESS MatchResponse_ResultType = 0
	MatchResponse_MATCH_ERROR   MatchResponse_ResultType = 1
	MatchResponse_MATCH_DENIED  MatchResponse_ResultType = 2
)

// Enum value maps for MatchResponse_ResultType.
var (
	MatchResponse_ResultType_name = map[int32]string{
		0: "MATCH_SUCCESS",
		1: "MATCH_ERROR",
		2: "MATCH_DENIED",
	}
	MatchResponse_ResultType_value = map[string]int32{
		"MATCH_SUCCESS": 0,
		"MATCH_ERROR":   1,
		"MATCH_DENIED":  2,
	}
)

func (x MatchResponse_ResultType) Enum() *MatchResponse_ResultType {
	p := new(MatchResponse_ResultType)
	*p = x
	return p
}

func (x MatchResponse_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchResponse_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_fluxion_grpc_fluxion_proto_enumTypes[1].Descriptor()
}

func (MatchResponse_ResultType) Type() protoreflect.EnumType {
	return &file_pkg_fluxion_grpc_fluxion_proto_enumTypes[1]
}

func (x MatchResponse_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchResponse_ResultType.Descriptor instead.
func (MatchResponse_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{3, 0}
}

type CancelResponse_ResultType int32

const (
	CancelResponse_CANCEL_SUCCESS       CancelResponse_ResultType = 0
	CancelResponse_CANCEL_REQUEST_ERROR CancelResponse_ResultType = 1
	CancelResponse_CANCEL_ERROR         CancelResponse_ResultType = 2
	CancelResponse_CANCEL_DENIED        CancelResponse_ResultType = 3
)

// Enum value maps for CancelResponse_ResultType.
var (
	CancelResponse_ResultType_name = map[int32]string{
		0: "CANCEL_SUCCESS",
		1: "CANCEL_REQUEST_ERROR",
		2: "CANCEL_ERROR",
		3: "CANCEL_DENIED",
	}
	CancelResponse_ResultType_value = map[string]int32{
		"CANCEL_SUCCESS":       0,
		"CANCEL_REQUEST_ERROR": 1,
		"CANCEL_ERROR":         2,
		"CANCEL_DENIED":        3,
	}
)

func (x CancelResponse_ResultType) Enum() *CancelResponse_ResultType {
	p := new(CancelResponse_ResultType)
	*p = x
	return p
}

func (x CancelResponse_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CancelResponse_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_fluxion_grpc_fluxion_proto_enumTypes[2].Descriptor()
}

func (CancelResponse_ResultType) Type() protoreflect.EnumType {
	return &file_pkg_fluxion_grpc_fluxion_proto_enumTypes[2]
}

func (x CancelResponse_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CancelResponse_ResultType.Descriptor instead.
func (CancelResponse_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{5, 0}
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy string `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	Jgf    string `protobuf:"bytes,2,opt,name=jgf,proto3" json:"jgf,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *InitRequest) GetJgf() string {
	if x != nil {
		return x.Jgf
	}
	return ""
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status InitResponse_ResultType `protobuf:"varint,1,opt,name=status,proto3,enum=fluxion.InitResponse_ResultType" json:"status,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetStatus() InitResponse_ResultType {
	if x != nil {
		return x.Status
	}
	return InitResponse_INIT_SUCCESS
}

// The Match request message (allocate, allocate_orelse_reserve)
type MatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobspec     string `protobuf:"bytes,1,opt,name=jobspec,proto3" json:"jobspec,omitempty"`
	Request     string `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Reservation bool   `protobuf:"varint,3,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *MatchRequest) Reset() {
	*x = MatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest) ProtoMessage() {}

func (x *MatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRequest.ProtoReflect.Descriptor instead.
func (*MatchRequest) Descriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{2}
}

func (x *MatchRequest) GetJobspec() string {
	if x != nil {
		return x.Jobspec
	}
	return ""
}

func (x *MatchRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *MatchRequest) GetReservation() bool {
	if x != nil {
		return x.Reservation
	}
	return false
}

// The Match response message
type MatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allocation string                   `protobuf:"bytes,1,opt,name=allocation,proto3" json:"allocation,omitempty"`
	Jobid      int64                    `protobuf:"varint,2,opt,name=jobid,proto3" json:"jobid,omitempty"`
	Reserved   bool                     `protobuf:"varint,3,opt,name=reserved,proto3" json:"reserved,omitempty"`
	At         int64                    `protobuf:"varint,4,opt,name=at,proto3" json:"at,omitempty"`
	Overhead   float32                  `protobuf:"fixed32,5,opt,name=overhead,proto3" json:"overhead,omitempty"`
	Status     MatchResponse_ResultType `protobuf:"varint,6,opt,name=status,proto3,enum=fluxion.MatchResponse_ResultType" json:"status,omitempty"`
}

func (x *MatchResponse) Reset() {
	*x = MatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResponse) ProtoMessage() {}

func (x *MatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResponse.ProtoReflect.Descriptor instead.
func (*MatchResponse) Descriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{3}
}

func (x *MatchResponse) GetAllocation() string {
	if x != nil {
		return x.Allocation
	}
	return ""
}

func (x *MatchResponse) GetJobid() int64 {
	if x != nil {
		return x.Jobid
	}
	return 0
}

func (x *MatchResponse) GetReserved() bool {
	if x != nil {
		return x.Reserved
	}
	return false
}

func (x *MatchResponse) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *MatchResponse) GetOverhead() float32 {
	if x != nil {
		return x.Overhead
	}
	return 0
}

func (x *MatchResponse) GetStatus() MatchResponse_ResultType {
	if x != nil {
		return x.Status
	}
	return MatchResponse_MATCH_SUCCESS
}

type CancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID int64 `protobuf:"varint,2,opt,name=jobID,proto3" json:"jobID,omitempty"`
}

func (x *CancelRequest) Reset() {
	*x = CancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRequest) ProtoMessage() {}

func (x *CancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRequest.ProtoReflect.Descriptor instead.
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{4}
}

func (x *CancelRequest) GetJobID() int64 {
	if x != nil {
		return x.JobID
	}
	return 0
}

// The Match response message
type CancelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID  int64                     `protobuf:"varint,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Error  int32                     `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	Status CancelResponse_ResultType `protobuf:"varint,3,opt,name=status,proto3,enum=fluxion.CancelResponse_ResultType" json:"status,omitempty"`
}

func (x *CancelResponse) Reset() {
	*x = CancelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelResponse) ProtoMessage() {}

func (x *CancelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluxion_grpc_fluxion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelResponse.ProtoReflect.Descriptor instead.
func (*CancelResponse) Descriptor() ([]byte, []int) {
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP(), []int{5}
}

func (x *CancelResponse) GetJobID() int64 {
	if x != nil {
		return x.JobID
	}
	return 0
}

func (x *CancelResponse) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *CancelResponse) GetStatus() CancelResponse_ResultType {
	if x != nil {
		return x.Status
	}
	return CancelResponse_CANCEL_SUCCESS
}

var File_pkg_fluxion_grpc_fluxion_proto protoreflect.FileDescriptor

var file_pkg_fluxion_grpc_fluxion_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0b, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6a, 0x67, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a,
	0x67, 0x66, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a,
	0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x49,
	0x4e, 0x49, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x02, 0x22, 0x64,
	0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6a, 0x6f, 0x62, 0x73, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x02, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72,
	0x68, 0x65, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72,
	0x68, 0x65, 0x61, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x42, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a,
	0x0d, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45,
	0x44, 0x10, 0x02, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x22, 0xd9, 0x01, 0x0a, 0x0e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6a, 0x6f,
	0x62, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x75, 0x78,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5f, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x44, 0x45,
	0x4e, 0x49, 0x45, 0x44, 0x10, 0x03, 0x32, 0xbe, 0x01, 0x0a, 0x0e, 0x46, 0x6c, 0x75, 0x78, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x15, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x6c, 0x75, 0x78,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x16, 0x2e,
	0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x69,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x67, 0x65, 0x64, 0x2d,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x69, 0x6f, 0x6e, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_fluxion_grpc_fluxion_proto_rawDescOnce sync.Once
	file_pkg_fluxion_grpc_fluxion_proto_rawDescData = file_pkg_fluxion_grpc_fluxion_proto_rawDesc
)

func file_pkg_fluxion_grpc_fluxion_proto_rawDescGZIP() []byte {
	file_pkg_fluxion_grpc_fluxion_proto_rawDescOnce.Do(func() {
		file_pkg_fluxion_grpc_fluxion_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_fluxion_grpc_fluxion_proto_rawDescData)
	})
	return file_pkg_fluxion_grpc_fluxion_proto_rawDescData
}

var file_pkg_fluxion_grpc_fluxion_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_fluxion_grpc_fluxion_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_fluxion_grpc_fluxion_proto_goTypes = []interface{}{
	(InitResponse_ResultType)(0),   // 0: fluxion.InitResponse.ResultType
	(MatchResponse_ResultType)(0),  // 1: fluxion.MatchResponse.ResultType
	(CancelResponse_ResultType)(0), // 2: fluxion.CancelResponse.ResultType
	(*InitRequest)(nil),            // 3: fluxion.InitRequest
	(*InitResponse)(nil),           // 4: fluxion.InitResponse
	(*MatchRequest)(nil),           // 5: fluxion.MatchRequest
	(*MatchResponse)(nil),          // 6: fluxion.MatchResponse
	(*CancelRequest)(nil),          // 7: fluxion.CancelRequest
	(*CancelResponse)(nil),         // 8: fluxion.CancelResponse
}
var file_pkg_fluxion_grpc_fluxion_proto_depIdxs = []int32{
	0, // 0: fluxion.InitResponse.status:type_name -> fluxion.InitResponse.ResultType
	1, // 1: fluxion.MatchResponse.status:type_name -> fluxion.MatchResponse.ResultType
	2, // 2: fluxion.CancelResponse.status:type_name -> fluxion.CancelResponse.ResultType
	5, // 3: fluxion.FluxionService.Match:input_type -> fluxion.MatchRequest
	7, // 4: fluxion.FluxionService.Cancel:input_type -> fluxion.CancelRequest
	3, // 5: fluxion.FluxionService.Init:input_type -> fluxion.InitRequest
	6, // 6: fluxion.FluxionService.Match:output_type -> fluxion.MatchResponse
	8, // 7: fluxion.FluxionService.Cancel:output_type -> fluxion.CancelResponse
	4, // 8: fluxion.FluxionService.Init:output_type -> fluxion.InitResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_fluxion_grpc_fluxion_proto_init() }
func file_pkg_fluxion_grpc_fluxion_proto_init() {
	if File_pkg_fluxion_grpc_fluxion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_fluxion_grpc_fluxion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_pkg_fluxion_grpc_fluxion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
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
		file_pkg_fluxion_grpc_fluxion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRequest); i {
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
		file_pkg_fluxion_grpc_fluxion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchResponse); i {
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
		file_pkg_fluxion_grpc_fluxion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelRequest); i {
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
		file_pkg_fluxion_grpc_fluxion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelResponse); i {
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
			RawDescriptor: file_pkg_fluxion_grpc_fluxion_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_fluxion_grpc_fluxion_proto_goTypes,
		DependencyIndexes: file_pkg_fluxion_grpc_fluxion_proto_depIdxs,
		EnumInfos:         file_pkg_fluxion_grpc_fluxion_proto_enumTypes,
		MessageInfos:      file_pkg_fluxion_grpc_fluxion_proto_msgTypes,
	}.Build()
	File_pkg_fluxion_grpc_fluxion_proto = out.File
	file_pkg_fluxion_grpc_fluxion_proto_rawDesc = nil
	file_pkg_fluxion_grpc_fluxion_proto_goTypes = nil
	file_pkg_fluxion_grpc_fluxion_proto_depIdxs = nil
}
