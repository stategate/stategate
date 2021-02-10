// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: schema.proto

package eventgate

import (
	context "context"
	_ "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// HistoryOpts are options when querying historical events
type HistoryOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// only return events that occurred after specified min timestamp
	Min *timestamp.Timestamp `protobuf:"bytes,2,opt,name=min,proto3" json:"min,omitempty"`
	// only return events that occurred before specified max timestamp
	Max *timestamp.Timestamp `protobuf:"bytes,3,opt,name=max,proto3" json:"max,omitempty"`
	// limit returned events
	Limit int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// offset returned events(pagination)
	Offset int64 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *HistoryOpts) Reset() {
	*x = HistoryOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryOpts) ProtoMessage() {}

func (x *HistoryOpts) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryOpts.ProtoReflect.Descriptor instead.
func (*HistoryOpts) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryOpts) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *HistoryOpts) GetMin() *timestamp.Timestamp {
	if x != nil {
		return x.Min
	}
	return nil
}

func (x *HistoryOpts) GetMax() *timestamp.Timestamp {
	if x != nil {
		return x.Max
	}
	return nil
}

func (x *HistoryOpts) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *HistoryOpts) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// ReceiveOpts filters events before they are received by a consumer
type ReceiveOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *ReceiveOpts) Reset() {
	*x = ReceiveOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveOpts) ProtoMessage() {}

func (x *ReceiveOpts) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveOpts.ProtoReflect.Descriptor instead.
func (*ReceiveOpts) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveOpts) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

// Event is a specification for describing event-sourced data
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the channel/subject to which the event will be sent
	Channel string `protobuf:"bytes,30,opt,name=channel,proto3" json:"channel,omitempty"`
	// The event payload(structured).
	Data *_struct.Struct `protobuf:"bytes,31,opt,name=data,proto3" json:"data,omitempty"`
	// Arbitrary metadata about the event
	Metadata *_struct.Struct `protobuf:"bytes,32,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Event) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Event) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// EventDetail wraps an Event with additional details.
type EventDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the event(uuid).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Identifies the channel/subject to which the event will be sent
	Channel string `protobuf:"bytes,30,opt,name=channel,proto3" json:"channel,omitempty"`
	// The event payload(structured).
	Data *_struct.Struct `protobuf:"bytes,31,opt,name=data,proto3" json:"data,omitempty"`
	// Arbitrary metadata about the event
	Metadata *_struct.Struct `protobuf:"bytes,32,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The authentication claims of the event producer.
	Claims *_struct.Struct `protobuf:"bytes,2,opt,name=claims,proto3" json:"claims,omitempty"`
	// Timestamp of when the event was received.
	Time *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *EventDetail) Reset() {
	*x = EventDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDetail) ProtoMessage() {}

func (x *EventDetail) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDetail.ProtoReflect.Descriptor instead.
func (*EventDetail) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{3}
}

func (x *EventDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventDetail) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *EventDetail) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EventDetail) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EventDetail) GetClaims() *_struct.Struct {
	if x != nil {
		return x.Claims
	}
	return nil
}

func (x *EventDetail) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

// EventDetails is an array of event details
type EventDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*EventDetail `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventDetails) Reset() {
	*x = EventDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDetails) ProtoMessage() {}

func (x *EventDetails) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDetails.ProtoReflect.Descriptor instead.
func (*EventDetails) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{4}
}

func (x *EventDetails) GetEvents() []*EventDetail {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_schema_proto protoreflect.FileDescriptor

var file_schema_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b,
	0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x0b, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x12, 0x1c, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x2f, 0x0a, 0x0b, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x93, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xa3, 0x02, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2,
	0xdf, 0x1f, 0x03, 0x90, 0x01, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x33, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12,
	0x36, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20,
	0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xf3, 0x01, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x47, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x04,
	0x53, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x4d, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4f,
	0x70, 0x74, 0x73, 0x1a, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x30, 0x01, 0x12,
	0x4c, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x70,
	0x74, 0x73, 0x1a, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x0b, 0x5a,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x67, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_schema_proto_rawDescOnce sync.Once
	file_schema_proto_rawDescData = file_schema_proto_rawDesc
)

func file_schema_proto_rawDescGZIP() []byte {
	file_schema_proto_rawDescOnce.Do(func() {
		file_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_proto_rawDescData)
	})
	return file_schema_proto_rawDescData
}

var file_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_schema_proto_goTypes = []interface{}{
	(*HistoryOpts)(nil),         // 0: eventgate.HistoryOpts
	(*ReceiveOpts)(nil),         // 1: eventgate.ReceiveOpts
	(*Event)(nil),               // 2: eventgate.Event
	(*EventDetail)(nil),         // 3: eventgate.EventDetail
	(*EventDetails)(nil),        // 4: eventgate.EventDetails
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*_struct.Struct)(nil),      // 6: google.protobuf.Struct
	(*empty.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_schema_proto_depIdxs = []int32{
	5,  // 0: eventgate.HistoryOpts.min:type_name -> google.protobuf.Timestamp
	5,  // 1: eventgate.HistoryOpts.max:type_name -> google.protobuf.Timestamp
	6,  // 2: eventgate.Event.data:type_name -> google.protobuf.Struct
	6,  // 3: eventgate.Event.metadata:type_name -> google.protobuf.Struct
	6,  // 4: eventgate.EventDetail.data:type_name -> google.protobuf.Struct
	6,  // 5: eventgate.EventDetail.metadata:type_name -> google.protobuf.Struct
	6,  // 6: eventgate.EventDetail.claims:type_name -> google.protobuf.Struct
	5,  // 7: eventgate.EventDetail.time:type_name -> google.protobuf.Timestamp
	3,  // 8: eventgate.EventDetails.events:type_name -> eventgate.EventDetail
	2,  // 9: eventgate.EventGateService.Send:input_type -> eventgate.Event
	1,  // 10: eventgate.EventGateService.Receive:input_type -> eventgate.ReceiveOpts
	0,  // 11: eventgate.EventGateService.History:input_type -> eventgate.HistoryOpts
	7,  // 12: eventgate.EventGateService.Send:output_type -> google.protobuf.Empty
	3,  // 13: eventgate.EventGateService.Receive:output_type -> eventgate.EventDetail
	4,  // 14: eventgate.EventGateService.History:output_type -> eventgate.EventDetails
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_schema_proto_init() }
func file_schema_proto_init() {
	if File_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryOpts); i {
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
		file_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveOpts); i {
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
		file_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDetail); i {
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
		file_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDetails); i {
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
			RawDescriptor: file_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_proto_goTypes,
		DependencyIndexes: file_schema_proto_depIdxs,
		MessageInfos:      file_schema_proto_msgTypes,
	}.Build()
	File_schema_proto = out.File
	file_schema_proto_rawDesc = nil
	file_schema_proto_goTypes = nil
	file_schema_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventGateServiceClient is the client API for EventGateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventGateServiceClient interface {
	// Send broadcasts an event to all consumers on a given channel. Producers invoke this method.
	Send(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error)
	// Receive creates an event stream/subscription to a given channel until fn returns false OR the context cancels.
	// Event Consumers invoke this method.
	Receive(ctx context.Context, in *ReceiveOpts, opts ...grpc.CallOption) (EventGateService_ReceiveClient, error)
	// History returns an array of immutable historical events.
	// Event Consumers invoke this method.
	History(ctx context.Context, in *HistoryOpts, opts ...grpc.CallOption) (*EventDetails, error)
}

type eventGateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventGateServiceClient(cc grpc.ClientConnInterface) EventGateServiceClient {
	return &eventGateServiceClient{cc}
}

func (c *eventGateServiceClient) Send(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/eventgate.EventGateService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventGateServiceClient) Receive(ctx context.Context, in *ReceiveOpts, opts ...grpc.CallOption) (EventGateService_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventGateService_serviceDesc.Streams[0], "/eventgate.EventGateService/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventGateServiceReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventGateService_ReceiveClient interface {
	Recv() (*EventDetail, error)
	grpc.ClientStream
}

type eventGateServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *eventGateServiceReceiveClient) Recv() (*EventDetail, error) {
	m := new(EventDetail)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventGateServiceClient) History(ctx context.Context, in *HistoryOpts, opts ...grpc.CallOption) (*EventDetails, error) {
	out := new(EventDetails)
	err := c.cc.Invoke(ctx, "/eventgate.EventGateService/History", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventGateServiceServer is the server API for EventGateService service.
type EventGateServiceServer interface {
	// Send broadcasts an event to all consumers on a given channel. Producers invoke this method.
	Send(context.Context, *Event) (*empty.Empty, error)
	// Receive creates an event stream/subscription to a given channel until fn returns false OR the context cancels.
	// Event Consumers invoke this method.
	Receive(*ReceiveOpts, EventGateService_ReceiveServer) error
	// History returns an array of immutable historical events.
	// Event Consumers invoke this method.
	History(context.Context, *HistoryOpts) (*EventDetails, error)
}

// UnimplementedEventGateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventGateServiceServer struct {
}

func (*UnimplementedEventGateServiceServer) Send(context.Context, *Event) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedEventGateServiceServer) Receive(*ReceiveOpts, EventGateService_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (*UnimplementedEventGateServiceServer) History(context.Context, *HistoryOpts) (*EventDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}

func RegisterEventGateServiceServer(s *grpc.Server, srv EventGateServiceServer) {
	s.RegisterService(&_EventGateService_serviceDesc, srv)
}

func _EventGateService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventGateServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventgate.EventGateService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventGateServiceServer).Send(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventGateService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveOpts)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventGateServiceServer).Receive(m, &eventGateServiceReceiveServer{stream})
}

type EventGateService_ReceiveServer interface {
	Send(*EventDetail) error
	grpc.ServerStream
}

type eventGateServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *eventGateServiceReceiveServer) Send(m *EventDetail) error {
	return x.ServerStream.SendMsg(m)
}

func _EventGateService_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventGateServiceServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventgate.EventGateService/History",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventGateServiceServer).History(ctx, req.(*HistoryOpts))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventGateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventgate.EventGateService",
	HandlerType: (*EventGateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _EventGateService_Send_Handler,
		},
		{
			MethodName: "History",
			Handler:    _EventGateService_History_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _EventGateService_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "schema.proto",
}
