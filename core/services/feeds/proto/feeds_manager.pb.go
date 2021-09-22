// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: pkg/noderpc/proto/feeds_manager.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobType int32

const (
	JobType_JOB_TYPE_UNSPECIFIED  JobType = 0
	JobType_JOB_TYPE_FLUX_MONITOR JobType = 1
	JobType_JOB_TYPE_OCR          JobType = 2
)

// Enum value maps for JobType.
var (
	JobType_name = map[int32]string{
		0: "JOB_TYPE_UNSPECIFIED",
		1: "JOB_TYPE_FLUX_MONITOR",
		2: "JOB_TYPE_OCR",
	}
	JobType_value = map[string]int32{
		"JOB_TYPE_UNSPECIFIED":  0,
		"JOB_TYPE_FLUX_MONITOR": 1,
		"JOB_TYPE_OCR":          2,
	}
)

func (x JobType) Enum() *JobType {
	p := new(JobType)
	*p = x
	return p
}

func (x JobType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_noderpc_proto_feeds_manager_proto_enumTypes[0].Descriptor()
}

func (JobType) Type() protoreflect.EnumType {
	return &file_pkg_noderpc_proto_feeds_manager_proto_enumTypes[0]
}

func (x JobType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobType.Descriptor instead.
func (JobType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{0}
}

type UpdateNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobTypes           []JobType `protobuf:"varint,1,rep,packed,name=job_types,json=jobTypes,proto3,enum=cfm.JobType" json:"job_types,omitempty"`
	ChainId            int64     `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	AccountAddresses   []string  `protobuf:"bytes,3,rep,name=account_addresses,json=accountAddresses,proto3" json:"account_addresses,omitempty"`
	IsBootstrapPeer    bool      `protobuf:"varint,4,opt,name=is_bootstrap_peer,json=isBootstrapPeer,proto3" json:"is_bootstrap_peer,omitempty"`
	BootstrapMultiaddr string    `protobuf:"bytes,5,opt,name=bootstrap_multiaddr,json=bootstrapMultiaddr,proto3" json:"bootstrap_multiaddr,omitempty"`
	Version            string    `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	ChainIds           []int64   `protobuf:"varint,7,rep,packed,name=chain_ids,json=chainIds,proto3" json:"chain_ids,omitempty"`
}

func (x *UpdateNodeRequest) Reset() {
	*x = UpdateNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeRequest) ProtoMessage() {}

func (x *UpdateNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateNodeRequest) GetJobTypes() []JobType {
	if x != nil {
		return x.JobTypes
	}
	return nil
}

func (x *UpdateNodeRequest) GetChainId() int64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *UpdateNodeRequest) GetAccountAddresses() []string {
	if x != nil {
		return x.AccountAddresses
	}
	return nil
}

func (x *UpdateNodeRequest) GetIsBootstrapPeer() bool {
	if x != nil {
		return x.IsBootstrapPeer
	}
	return false
}

func (x *UpdateNodeRequest) GetBootstrapMultiaddr() string {
	if x != nil {
		return x.BootstrapMultiaddr
	}
	return ""
}

func (x *UpdateNodeRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UpdateNodeRequest) GetChainIds() []int64 {
	if x != nil {
		return x.ChainIds
	}
	return nil
}

type UpdateNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateNodeResponse) Reset() {
	*x = UpdateNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeResponse) ProtoMessage() {}

func (x *UpdateNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeResponse.ProtoReflect.Descriptor instead.
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{1}
}

type ApprovedJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ApprovedJobRequest) Reset() {
	*x = ApprovedJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovedJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovedJobRequest) ProtoMessage() {}

func (x *ApprovedJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovedJobRequest.ProtoReflect.Descriptor instead.
func (*ApprovedJobRequest) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ApprovedJobRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ApprovedJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApprovedJobResponse) Reset() {
	*x = ApprovedJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovedJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovedJobResponse) ProtoMessage() {}

func (x *ApprovedJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovedJobResponse.ProtoReflect.Descriptor instead.
func (*ApprovedJobResponse) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{3}
}

type RejectedJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *RejectedJobRequest) Reset() {
	*x = RejectedJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectedJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectedJobRequest) ProtoMessage() {}

func (x *RejectedJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectedJobRequest.ProtoReflect.Descriptor instead.
func (*RejectedJobRequest) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{4}
}

func (x *RejectedJobRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type RejectedJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RejectedJobResponse) Reset() {
	*x = RejectedJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectedJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectedJobResponse) ProtoMessage() {}

func (x *RejectedJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectedJobResponse.ProtoReflect.Descriptor instead.
func (*RejectedJobResponse) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{5}
}

type CancelledJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CancelledJobRequest) Reset() {
	*x = CancelledJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelledJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelledJobRequest) ProtoMessage() {}

func (x *CancelledJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelledJobRequest.ProtoReflect.Descriptor instead.
func (*CancelledJobRequest) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{6}
}

func (x *CancelledJobRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CancelledJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelledJobResponse) Reset() {
	*x = CancelledJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelledJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelledJobResponse) ProtoMessage() {}

func (x *CancelledJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelledJobResponse.ProtoReflect.Descriptor instead.
func (*CancelledJobResponse) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{7}
}

type ProposeJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       string   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Multiaddrs []string `protobuf:"bytes,3,rep,name=multiaddrs,proto3" json:"multiaddrs,omitempty"`
}

func (x *ProposeJobRequest) Reset() {
	*x = ProposeJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeJobRequest) ProtoMessage() {}

func (x *ProposeJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeJobRequest.ProtoReflect.Descriptor instead.
func (*ProposeJobRequest) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{8}
}

func (x *ProposeJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProposeJobRequest) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *ProposeJobRequest) GetMultiaddrs() []string {
	if x != nil {
		return x.Multiaddrs
	}
	return nil
}

type ProposeJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProposeJobResponse) Reset() {
	*x = ProposeJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeJobResponse) ProtoMessage() {}

func (x *ProposeJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeJobResponse.ProtoReflect.Descriptor instead.
func (*ProposeJobResponse) Descriptor() ([]byte, []int) {
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP(), []int{9}
}

func (x *ProposeJobResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pkg_noderpc_proto_feeds_manager_proto protoreflect.FileDescriptor

var file_pkg_noderpc_proto_feeds_manager_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x66, 0x6d, 0x22, 0x9a, 0x02, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x4a, 0x6f, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x69, 0x73, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x65, 0x65,
	0x72, 0x12, 0x2f, 0x0a, 0x13, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x28, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x28, 0x0a, 0x12, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72, 0x73, 0x22, 0x24, 0x0a,
	0x12, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x2a, 0x50, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4a, 0x4f, 0x42, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x55, 0x58, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f,
	0x52, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4a, 0x4f, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4f, 0x43, 0x52, 0x10, 0x02, 0x32, 0x96, 0x02, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x73, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x18, 0x2e, 0x63, 0x66, 0x6d, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4c,
	0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x16, 0x2e, 0x63, 0x66,
	0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x66, 0x65, 0x65, 0x64,
	0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_noderpc_proto_feeds_manager_proto_rawDescOnce sync.Once
	file_pkg_noderpc_proto_feeds_manager_proto_rawDescData = file_pkg_noderpc_proto_feeds_manager_proto_rawDesc
)

func file_pkg_noderpc_proto_feeds_manager_proto_rawDescGZIP() []byte {
	file_pkg_noderpc_proto_feeds_manager_proto_rawDescOnce.Do(func() {
		file_pkg_noderpc_proto_feeds_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_noderpc_proto_feeds_manager_proto_rawDescData)
	})
	return file_pkg_noderpc_proto_feeds_manager_proto_rawDescData
}

var file_pkg_noderpc_proto_feeds_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_noderpc_proto_feeds_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_noderpc_proto_feeds_manager_proto_goTypes = []interface{}{
	(JobType)(0),                 // 0: cfm.JobType
	(*UpdateNodeRequest)(nil),    // 1: cfm.UpdateNodeRequest
	(*UpdateNodeResponse)(nil),   // 2: cfm.UpdateNodeResponse
	(*ApprovedJobRequest)(nil),   // 3: cfm.ApprovedJobRequest
	(*ApprovedJobResponse)(nil),  // 4: cfm.ApprovedJobResponse
	(*RejectedJobRequest)(nil),   // 5: cfm.RejectedJobRequest
	(*RejectedJobResponse)(nil),  // 6: cfm.RejectedJobResponse
	(*CancelledJobRequest)(nil),  // 7: cfm.CancelledJobRequest
	(*CancelledJobResponse)(nil), // 8: cfm.CancelledJobResponse
	(*ProposeJobRequest)(nil),    // 9: cfm.ProposeJobRequest
	(*ProposeJobResponse)(nil),   // 10: cfm.ProposeJobResponse
}
var file_pkg_noderpc_proto_feeds_manager_proto_depIdxs = []int32{
	0,  // 0: cfm.UpdateNodeRequest.job_types:type_name -> cfm.JobType
	3,  // 1: cfm.FeedsManager.ApprovedJob:input_type -> cfm.ApprovedJobRequest
	1,  // 2: cfm.FeedsManager.UpdateNode:input_type -> cfm.UpdateNodeRequest
	5,  // 3: cfm.FeedsManager.RejectedJob:input_type -> cfm.RejectedJobRequest
	7,  // 4: cfm.FeedsManager.CancelledJob:input_type -> cfm.CancelledJobRequest
	9,  // 5: cfm.NodeService.ProposeJob:input_type -> cfm.ProposeJobRequest
	4,  // 6: cfm.FeedsManager.ApprovedJob:output_type -> cfm.ApprovedJobResponse
	2,  // 7: cfm.FeedsManager.UpdateNode:output_type -> cfm.UpdateNodeResponse
	6,  // 8: cfm.FeedsManager.RejectedJob:output_type -> cfm.RejectedJobResponse
	8,  // 9: cfm.FeedsManager.CancelledJob:output_type -> cfm.CancelledJobResponse
	10, // 10: cfm.NodeService.ProposeJob:output_type -> cfm.ProposeJobResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_noderpc_proto_feeds_manager_proto_init() }
func file_pkg_noderpc_proto_feeds_manager_proto_init() {
	if File_pkg_noderpc_proto_feeds_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeRequest); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeResponse); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovedJobRequest); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovedJobResponse); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectedJobRequest); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectedJobResponse); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelledJobRequest); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelledJobResponse); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeJobRequest); i {
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
		file_pkg_noderpc_proto_feeds_manager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeJobResponse); i {
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
			RawDescriptor: file_pkg_noderpc_proto_feeds_manager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_noderpc_proto_feeds_manager_proto_goTypes,
		DependencyIndexes: file_pkg_noderpc_proto_feeds_manager_proto_depIdxs,
		EnumInfos:         file_pkg_noderpc_proto_feeds_manager_proto_enumTypes,
		MessageInfos:      file_pkg_noderpc_proto_feeds_manager_proto_msgTypes,
	}.Build()
	File_pkg_noderpc_proto_feeds_manager_proto = out.File
	file_pkg_noderpc_proto_feeds_manager_proto_rawDesc = nil
	file_pkg_noderpc_proto_feeds_manager_proto_goTypes = nil
	file_pkg_noderpc_proto_feeds_manager_proto_depIdxs = nil
}
