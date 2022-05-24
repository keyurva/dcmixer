// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: v1/property_values.proto

package proto

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

type PropertyValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property string `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Entity   string `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	// [Optional]
	// The limit of the number of values to return. The maximium limit is 1000.
	// If not specified, the default limit is 1000.
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// [Optional]
	// The pagination token for getting the next set of entries. This is empty
	// for the first request and needs to be set in the subsequent request.
	// This is the value returned from a prior call to PropertyValuesRequest
	NextToken string `protobuf:"bytes,4,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
	// Direction can only be "in" and "out"
	Direction string `protobuf:"bytes,5,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *PropertyValuesRequest) Reset() {
	*x = PropertyValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_property_values_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyValuesRequest) ProtoMessage() {}

func (x *PropertyValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_property_values_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyValuesRequest.ProtoReflect.Descriptor instead.
func (*PropertyValuesRequest) Descriptor() ([]byte, []int) {
	return file_v1_property_values_proto_rawDescGZIP(), []int{0}
}

func (x *PropertyValuesRequest) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *PropertyValuesRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *PropertyValuesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PropertyValuesRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

func (x *PropertyValuesRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type PropertyValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*EntityInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// The pagination token for getting the next set of entries.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *PropertyValuesResponse) Reset() {
	*x = PropertyValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_property_values_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyValuesResponse) ProtoMessage() {}

func (x *PropertyValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_property_values_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyValuesResponse.ProtoReflect.Descriptor instead.
func (*PropertyValuesResponse) Descriptor() ([]byte, []int) {
	return file_v1_property_values_proto_rawDescGZIP(), []int{1}
}

func (x *PropertyValuesResponse) GetData() []*EntityInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PropertyValuesResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type BulkPropertyValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property string   `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Entities []string `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	// [Optional]
	// The limit of the number of values to return. The maximium limit is 1000.
	// If not specified, the default limit is 1000.
	// This limit is applied for each entity, instead of the bulk result.
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// [Optional]
	// The pagination token for getting the next set of entries. This is empty
	// for the first request and needs to be set in the subsequent request.
	// This is the value returned from a prior call to PropertyValuesResponse
	NextToken string `protobuf:"bytes,4,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
	// Direction can only be "in" and "out"
	Direction string `protobuf:"bytes,5,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *BulkPropertyValuesRequest) Reset() {
	*x = BulkPropertyValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_property_values_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkPropertyValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkPropertyValuesRequest) ProtoMessage() {}

func (x *BulkPropertyValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_property_values_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkPropertyValuesRequest.ProtoReflect.Descriptor instead.
func (*BulkPropertyValuesRequest) Descriptor() ([]byte, []int) {
	return file_v1_property_values_proto_rawDescGZIP(), []int{2}
}

func (x *BulkPropertyValuesRequest) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *BulkPropertyValuesRequest) GetEntities() []string {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *BulkPropertyValuesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BulkPropertyValuesRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

func (x *BulkPropertyValuesRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type BulkPropertyValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BulkPropertyValuesResponse_EntityPropertyValues `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// The pagination token for getting the next set of entries.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *BulkPropertyValuesResponse) Reset() {
	*x = BulkPropertyValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_property_values_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkPropertyValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkPropertyValuesResponse) ProtoMessage() {}

func (x *BulkPropertyValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_property_values_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkPropertyValuesResponse.ProtoReflect.Descriptor instead.
func (*BulkPropertyValuesResponse) Descriptor() ([]byte, []int) {
	return file_v1_property_values_proto_rawDescGZIP(), []int{3}
}

func (x *BulkPropertyValuesResponse) GetData() []*BulkPropertyValuesResponse_EntityPropertyValues {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BulkPropertyValuesResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type BulkPropertyValuesResponse_EntityPropertyValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity string        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Values []*EntityInfo `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *BulkPropertyValuesResponse_EntityPropertyValues) Reset() {
	*x = BulkPropertyValuesResponse_EntityPropertyValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_property_values_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkPropertyValuesResponse_EntityPropertyValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkPropertyValuesResponse_EntityPropertyValues) ProtoMessage() {}

func (x *BulkPropertyValuesResponse_EntityPropertyValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_property_values_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkPropertyValuesResponse_EntityPropertyValues.ProtoReflect.Descriptor instead.
func (*BulkPropertyValuesResponse_EntityPropertyValues) Descriptor() ([]byte, []int) {
	return file_v1_property_values_proto_rawDescGZIP(), []int{3, 0}
}

func (x *BulkPropertyValuesResponse_EntityPropertyValues) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *BulkPropertyValuesResponse_EntityPropertyValues) GetValues() []*EntityInfo {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_v1_property_values_proto protoreflect.FileDescriptor

var file_v1_property_values_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x0c, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x64, 0x0a, 0x16, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xa6, 0x01, 0x0a, 0x19, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf1, 0x01, 0x0a, 0x1a, 0x42, 0x75, 0x6c,
	0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x5f, 0x0a, 0x14, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_property_values_proto_rawDescOnce sync.Once
	file_v1_property_values_proto_rawDescData = file_v1_property_values_proto_rawDesc
)

func file_v1_property_values_proto_rawDescGZIP() []byte {
	file_v1_property_values_proto_rawDescOnce.Do(func() {
		file_v1_property_values_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_property_values_proto_rawDescData)
	})
	return file_v1_property_values_proto_rawDescData
}

var file_v1_property_values_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_property_values_proto_goTypes = []interface{}{
	(*PropertyValuesRequest)(nil),                           // 0: datacommons.v1.PropertyValuesRequest
	(*PropertyValuesResponse)(nil),                          // 1: datacommons.v1.PropertyValuesResponse
	(*BulkPropertyValuesRequest)(nil),                       // 2: datacommons.v1.BulkPropertyValuesRequest
	(*BulkPropertyValuesResponse)(nil),                      // 3: datacommons.v1.BulkPropertyValuesResponse
	(*BulkPropertyValuesResponse_EntityPropertyValues)(nil), // 4: datacommons.v1.BulkPropertyValuesResponse.EntityPropertyValues
	(*EntityInfo)(nil),                                      // 5: datacommons.EntityInfo
}
var file_v1_property_values_proto_depIdxs = []int32{
	5, // 0: datacommons.v1.PropertyValuesResponse.data:type_name -> datacommons.EntityInfo
	4, // 1: datacommons.v1.BulkPropertyValuesResponse.data:type_name -> datacommons.v1.BulkPropertyValuesResponse.EntityPropertyValues
	5, // 2: datacommons.v1.BulkPropertyValuesResponse.EntityPropertyValues.values:type_name -> datacommons.EntityInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_property_values_proto_init() }
func file_v1_property_values_proto_init() {
	if File_v1_property_values_proto != nil {
		return
	}
	file_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_property_values_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyValuesRequest); i {
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
		file_v1_property_values_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyValuesResponse); i {
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
		file_v1_property_values_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkPropertyValuesRequest); i {
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
		file_v1_property_values_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkPropertyValuesResponse); i {
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
		file_v1_property_values_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkPropertyValuesResponse_EntityPropertyValues); i {
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
			RawDescriptor: file_v1_property_values_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_property_values_proto_goTypes,
		DependencyIndexes: file_v1_property_values_proto_depIdxs,
		MessageInfos:      file_v1_property_values_proto_msgTypes,
	}.Build()
	File_v1_property_values_proto = out.File
	file_v1_property_values_proto_rawDesc = nil
	file_v1_property_values_proto_goTypes = nil
	file_v1_property_values_proto_depIdxs = nil
}