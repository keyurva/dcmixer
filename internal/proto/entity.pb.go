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

//  **** IMPORTANT NOTE ****
//
//  The proto of BT data has to match exactly the g3 proto, including tag
//  number.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: entity.proto

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

// Basic info for a node (subject or object).
type EntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Types        []string `protobuf:"bytes,2,rep,name=types,proto3" json:"types,omitempty"`
	Dcid         string   `protobuf:"bytes,3,opt,name=dcid,proto3" json:"dcid,omitempty"`
	ProvenanceId string   `protobuf:"bytes,4,opt,name=provenance_id,json=provenanceId,proto3" json:"provenance_id,omitempty"`
	Value        string   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"` // Only for object value.
}

func (x *EntityInfo) Reset() {
	*x = EntityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityInfo) ProtoMessage() {}

func (x *EntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityInfo.ProtoReflect.Descriptor instead.
func (*EntityInfo) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{0}
}

func (x *EntityInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EntityInfo) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *EntityInfo) GetDcid() string {
	if x != nil {
		return x.Dcid
	}
	return ""
}

func (x *EntityInfo) GetProvenanceId() string {
	if x != nil {
		return x.ProvenanceId
	}
	return ""
}

func (x *EntityInfo) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// A page of nodes. The page number starts from 0, and is in the cache key.
// Page size is set by ::datacommons::prophet::kPageSize.
type PagedNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of EntityInfo messages for PagedPropVal{In|Out} cache result.
	Nodes          []*EntityInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	TotalPageCount float64       `protobuf:"fixed64,2,opt,name=total_page_count,json=totalPageCount,proto3" json:"total_page_count,omitempty"`
}

func (x *PagedNodes) Reset() {
	*x = PagedNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedNodes) ProtoMessage() {}

func (x *PagedNodes) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedNodes.ProtoReflect.Descriptor instead.
func (*PagedNodes) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{1}
}

func (x *PagedNodes) GetNodes() []*EntityInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *PagedNodes) GetTotalPageCount() float64 {
	if x != nil {
		return x.TotalPageCount
	}
	return 0
}

// Basic info for a collection of nodes.
type EntityInfoCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*EntityInfo `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	Total    float64       `protobuf:"fixed64,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *EntityInfoCollection) Reset() {
	*x = EntityInfoCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityInfoCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityInfoCollection) ProtoMessage() {}

func (x *EntityInfoCollection) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityInfoCollection.ProtoReflect.Descriptor instead.
func (*EntityInfoCollection) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{2}
}

func (x *EntityInfoCollection) GetEntities() []*EntityInfo {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *EntityInfoCollection) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type IdWithProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prop string `protobuf:"bytes,1,opt,name=prop,proto3" json:"prop,omitempty"`
	Val  string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *IdWithProperty) Reset() {
	*x = IdWithProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdWithProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdWithProperty) ProtoMessage() {}

func (x *IdWithProperty) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdWithProperty.ProtoReflect.Descriptor instead.
func (*IdWithProperty) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{3}
}

func (x *IdWithProperty) GetProp() string {
	if x != nil {
		return x.Prop
	}
	return ""
}

func (x *IdWithProperty) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type EntityIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []*IdWithProperty `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *EntityIds) Reset() {
	*x = EntityIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityIds) ProtoMessage() {}

func (x *EntityIds) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityIds.ProtoReflect.Descriptor instead.
func (*EntityIds) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{4}
}

func (x *EntityIds) GetIds() []*IdWithProperty {
	if x != nil {
		return x.Ids
	}
	return nil
}

// An entity is represented by a subgraph, which contains itself and its
// neighbors.
type EntitySubGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED: source_id must be a key within `sub_graph.nodes`, or one of the
	// `ids`.
	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	// Types that are assignable to GraphRepresentation:
	//
	//	*EntitySubGraph_SubGraph
	//	*EntitySubGraph_EntityIds
	GraphRepresentation isEntitySubGraph_GraphRepresentation `protobuf_oneof:"graph_representation"`
}

func (x *EntitySubGraph) Reset() {
	*x = EntitySubGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntitySubGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitySubGraph) ProtoMessage() {}

func (x *EntitySubGraph) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitySubGraph.ProtoReflect.Descriptor instead.
func (*EntitySubGraph) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{5}
}

func (x *EntitySubGraph) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (m *EntitySubGraph) GetGraphRepresentation() isEntitySubGraph_GraphRepresentation {
	if m != nil {
		return m.GraphRepresentation
	}
	return nil
}

func (x *EntitySubGraph) GetSubGraph() *McfGraph {
	if x, ok := x.GetGraphRepresentation().(*EntitySubGraph_SubGraph); ok {
		return x.SubGraph
	}
	return nil
}

func (x *EntitySubGraph) GetEntityIds() *EntityIds {
	if x, ok := x.GetGraphRepresentation().(*EntitySubGraph_EntityIds); ok {
		return x.EntityIds
	}
	return nil
}

type isEntitySubGraph_GraphRepresentation interface {
	isEntitySubGraph_GraphRepresentation()
}

type EntitySubGraph_SubGraph struct {
	SubGraph *McfGraph `protobuf:"bytes,2,opt,name=sub_graph,json=subGraph,proto3,oneof"`
}

type EntitySubGraph_EntityIds struct {
	EntityIds *EntityIds `protobuf:"bytes,3,opt,name=entity_ids,json=entityIds,proto3,oneof"`
}

func (*EntitySubGraph_SubGraph) isEntitySubGraph_GraphRepresentation() {}

func (*EntitySubGraph_EntityIds) isEntitySubGraph_GraphRepresentation() {}

type EntityPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityOne *EntitySubGraph `protobuf:"bytes,1,opt,name=entity_one,json=entityOne,proto3" json:"entity_one,omitempty"`
	EntityTwo *EntitySubGraph `protobuf:"bytes,2,opt,name=entity_two,json=entityTwo,proto3" json:"entity_two,omitempty"`
}

func (x *EntityPair) Reset() {
	*x = EntityPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityPair) ProtoMessage() {}

func (x *EntityPair) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityPair.ProtoReflect.Descriptor instead.
func (*EntityPair) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{6}
}

func (x *EntityPair) GetEntityOne() *EntitySubGraph {
	if x != nil {
		return x.EntityOne
	}
	return nil
}

func (x *EntityPair) GetEntityTwo() *EntitySubGraph {
	if x != nil {
		return x.EntityTwo
	}
	return nil
}

type NodePropertyValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node   string        `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Values []*EntityInfo `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *NodePropertyValues) Reset() {
	*x = NodePropertyValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePropertyValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePropertyValues) ProtoMessage() {}

func (x *NodePropertyValues) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePropertyValues.ProtoReflect.Descriptor instead.
func (*NodePropertyValues) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{7}
}

func (x *NodePropertyValues) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *NodePropertyValues) GetValues() []*EntityInfo {
	if x != nil {
		return x.Values
	}
	return nil
}

type PropertyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property string `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PropertyValue) Reset() {
	*x = PropertyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyValue) ProtoMessage() {}

func (x *PropertyValue) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyValue.ProtoReflect.Descriptor instead.
func (*PropertyValue) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{8}
}

func (x *PropertyValue) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *PropertyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_entity_proto protoreflect.FileDescriptor

var file_entity_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x09, 0x6d, 0x63, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x63, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x65,
	0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x67, 0x0a, 0x14, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x36,
	0x0a, 0x0e, 0x49, 0x64, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x72, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x3a, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x49,
	0x64, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x63, 0x66, 0x47, 0x72, 0x61, 0x70, 0x68, 0x48, 0x00, 0x52, 0x08,
	0x73, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x73, 0x48, 0x00, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x73, 0x42, 0x16, 0x0a, 0x14, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x61, 0x69, 0x72, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x53, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4f, 0x6e, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74,
	0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x77, 0x6f,
	0x22, 0x59, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x78, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_proto_rawDescOnce sync.Once
	file_entity_proto_rawDescData = file_entity_proto_rawDesc
)

func file_entity_proto_rawDescGZIP() []byte {
	file_entity_proto_rawDescOnce.Do(func() {
		file_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_proto_rawDescData)
	})
	return file_entity_proto_rawDescData
}

var file_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_entity_proto_goTypes = []interface{}{
	(*EntityInfo)(nil),           // 0: datacommons.EntityInfo
	(*PagedNodes)(nil),           // 1: datacommons.PagedNodes
	(*EntityInfoCollection)(nil), // 2: datacommons.EntityInfoCollection
	(*IdWithProperty)(nil),       // 3: datacommons.IdWithProperty
	(*EntityIds)(nil),            // 4: datacommons.EntityIds
	(*EntitySubGraph)(nil),       // 5: datacommons.EntitySubGraph
	(*EntityPair)(nil),           // 6: datacommons.EntityPair
	(*NodePropertyValues)(nil),   // 7: datacommons.NodePropertyValues
	(*PropertyValue)(nil),        // 8: datacommons.PropertyValue
	(*McfGraph)(nil),             // 9: datacommons.McfGraph
}
var file_entity_proto_depIdxs = []int32{
	0, // 0: datacommons.PagedNodes.nodes:type_name -> datacommons.EntityInfo
	0, // 1: datacommons.EntityInfoCollection.entities:type_name -> datacommons.EntityInfo
	3, // 2: datacommons.EntityIds.ids:type_name -> datacommons.IdWithProperty
	9, // 3: datacommons.EntitySubGraph.sub_graph:type_name -> datacommons.McfGraph
	4, // 4: datacommons.EntitySubGraph.entity_ids:type_name -> datacommons.EntityIds
	5, // 5: datacommons.EntityPair.entity_one:type_name -> datacommons.EntitySubGraph
	5, // 6: datacommons.EntityPair.entity_two:type_name -> datacommons.EntitySubGraph
	0, // 7: datacommons.NodePropertyValues.values:type_name -> datacommons.EntityInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_entity_proto_init() }
func file_entity_proto_init() {
	if File_entity_proto != nil {
		return
	}
	file_mcf_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityInfo); i {
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
		file_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedNodes); i {
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
		file_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityInfoCollection); i {
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
		file_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdWithProperty); i {
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
		file_entity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityIds); i {
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
		file_entity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntitySubGraph); i {
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
		file_entity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityPair); i {
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
		file_entity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePropertyValues); i {
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
		file_entity_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyValue); i {
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
	file_entity_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*EntitySubGraph_SubGraph)(nil),
		(*EntitySubGraph_EntityIds)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_proto_goTypes,
		DependencyIndexes: file_entity_proto_depIdxs,
		MessageInfos:      file_entity_proto_msgTypes,
	}.Build()
	File_entity_proto = out.File
	file_entity_proto_rawDesc = nil
	file_entity_proto_goTypes = nil
	file_entity_proto_depIdxs = nil
}
