// Copyright 2015 The Cayley Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: quads.proto

package pquads

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

// Quad is in internal representation of quad used by Cayley.
type Quad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject        string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate      string `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object         string `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	Label          string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	SubjectValue   *Value `protobuf:"bytes,5,opt,name=subject_value,json=subjectValue,proto3" json:"subject_value,omitempty"`
	PredicateValue *Value `protobuf:"bytes,6,opt,name=predicate_value,json=predicateValue,proto3" json:"predicate_value,omitempty"`
	ObjectValue    *Value `protobuf:"bytes,7,opt,name=object_value,json=objectValue,proto3" json:"object_value,omitempty"`
	LabelValue     *Value `protobuf:"bytes,8,opt,name=label_value,json=labelValue,proto3" json:"label_value,omitempty"`
}

func (x *Quad) Reset() {
	*x = Quad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quad) ProtoMessage() {}

func (x *Quad) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quad.ProtoReflect.Descriptor instead.
func (*Quad) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{0}
}

func (x *Quad) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Quad) GetPredicate() string {
	if x != nil {
		return x.Predicate
	}
	return ""
}

func (x *Quad) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *Quad) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Quad) GetSubjectValue() *Value {
	if x != nil {
		return x.SubjectValue
	}
	return nil
}

func (x *Quad) GetPredicateValue() *Value {
	if x != nil {
		return x.PredicateValue
	}
	return nil
}

func (x *Quad) GetObjectValue() *Value {
	if x != nil {
		return x.ObjectValue
	}
	return nil
}

func (x *Quad) GetLabelValue() *Value {
	if x != nil {
		return x.LabelValue
	}
	return nil
}

// WireQuad is a quad that allows any value for it's directions.
type WireQuad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject   *Value `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate *Value `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object    *Value `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	Label     *Value `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *WireQuad) Reset() {
	*x = WireQuad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WireQuad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WireQuad) ProtoMessage() {}

func (x *WireQuad) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WireQuad.ProtoReflect.Descriptor instead.
func (*WireQuad) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{1}
}

func (x *WireQuad) GetSubject() *Value {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *WireQuad) GetPredicate() *Value {
	if x != nil {
		return x.Predicate
	}
	return nil
}

func (x *WireQuad) GetObject() *Value {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *WireQuad) GetLabel() *Value {
	if x != nil {
		return x.Label
	}
	return nil
}

// WireQuadRaw is the same as WireQuad, but doesn't decode underlying values.
type WireQuadRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject   []byte `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate []byte `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object    []byte `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	Label     []byte `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *WireQuadRaw) Reset() {
	*x = WireQuadRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WireQuadRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WireQuadRaw) ProtoMessage() {}

func (x *WireQuadRaw) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WireQuadRaw.ProtoReflect.Descriptor instead.
func (*WireQuadRaw) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{2}
}

func (x *WireQuadRaw) GetSubject() []byte {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *WireQuadRaw) GetPredicate() []byte {
	if x != nil {
		return x.Predicate
	}
	return nil
}

func (x *WireQuadRaw) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *WireQuadRaw) GetLabel() []byte {
	if x != nil {
		return x.Label
	}
	return nil
}

// StrictQuad is a quad as described by RDF spec.
type StrictQuad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject   *StrictQuad_Ref `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate *StrictQuad_Ref `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object    *Value          `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	Label     *StrictQuad_Ref `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *StrictQuad) Reset() {
	*x = StrictQuad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrictQuad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrictQuad) ProtoMessage() {}

func (x *StrictQuad) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrictQuad.ProtoReflect.Descriptor instead.
func (*StrictQuad) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{3}
}

func (x *StrictQuad) GetSubject() *StrictQuad_Ref {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *StrictQuad) GetPredicate() *StrictQuad_Ref {
	if x != nil {
		return x.Predicate
	}
	return nil
}

func (x *StrictQuad) GetObject() *Value {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *StrictQuad) GetLabel() *StrictQuad_Ref {
	if x != nil {
		return x.Label
	}
	return nil
}

// StrictQuadRaw is the same as StrictQuad, but doesn't decode underlying values.
type StrictQuadRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject   []byte `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate []byte `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object    []byte `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	Label     []byte `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *StrictQuadRaw) Reset() {
	*x = StrictQuadRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrictQuadRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrictQuadRaw) ProtoMessage() {}

func (x *StrictQuadRaw) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrictQuadRaw.ProtoReflect.Descriptor instead.
func (*StrictQuadRaw) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{4}
}

func (x *StrictQuadRaw) GetSubject() []byte {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *StrictQuadRaw) GetPredicate() []byte {
	if x != nil {
		return x.Predicate
	}
	return nil
}

func (x *StrictQuadRaw) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *StrictQuadRaw) GetLabel() []byte {
	if x != nil {
		return x.Label
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*Value_Raw
	//	*Value_Str
	//	*Value_Iri
	//	*Value_Bnode
	//	*Value_TypedStr
	//	*Value_LangStr
	//	*Value_Int
	//	*Value_Float
	//	*Value_Boolean
	//	*Value_Time
	Value isValue_Value `protobuf_oneof:"value"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{5}
}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Value) GetRaw() []byte {
	if x, ok := x.GetValue().(*Value_Raw); ok {
		return x.Raw
	}
	return nil
}

func (x *Value) GetStr() string {
	if x, ok := x.GetValue().(*Value_Str); ok {
		return x.Str
	}
	return ""
}

func (x *Value) GetIri() string {
	if x, ok := x.GetValue().(*Value_Iri); ok {
		return x.Iri
	}
	return ""
}

func (x *Value) GetBnode() string {
	if x, ok := x.GetValue().(*Value_Bnode); ok {
		return x.Bnode
	}
	return ""
}

func (x *Value) GetTypedStr() *Value_TypedString {
	if x, ok := x.GetValue().(*Value_TypedStr); ok {
		return x.TypedStr
	}
	return nil
}

func (x *Value) GetLangStr() *Value_LangString {
	if x, ok := x.GetValue().(*Value_LangStr); ok {
		return x.LangStr
	}
	return nil
}

func (x *Value) GetInt() int64 {
	if x, ok := x.GetValue().(*Value_Int); ok {
		return x.Int
	}
	return 0
}

func (x *Value) GetFloat() float64 {
	if x, ok := x.GetValue().(*Value_Float); ok {
		return x.Float
	}
	return 0
}

func (x *Value) GetBoolean() bool {
	if x, ok := x.GetValue().(*Value_Boolean); ok {
		return x.Boolean
	}
	return false
}

func (x *Value) GetTime() *Value_Timestamp {
	if x, ok := x.GetValue().(*Value_Time); ok {
		return x.Time
	}
	return nil
}

type isValue_Value interface {
	isValue_Value()
}

type Value_Raw struct {
	Raw []byte `protobuf:"bytes,1,opt,name=raw,proto3,oneof"`
}

type Value_Str struct {
	Str string `protobuf:"bytes,2,opt,name=str,proto3,oneof"`
}

type Value_Iri struct {
	Iri string `protobuf:"bytes,3,opt,name=iri,proto3,oneof"`
}

type Value_Bnode struct {
	Bnode string `protobuf:"bytes,4,opt,name=bnode,proto3,oneof"`
}

type Value_TypedStr struct {
	TypedStr *Value_TypedString `protobuf:"bytes,5,opt,name=typed_str,json=typedStr,proto3,oneof"`
}

type Value_LangStr struct {
	LangStr *Value_LangString `protobuf:"bytes,6,opt,name=lang_str,json=langStr,proto3,oneof"`
}

type Value_Int struct {
	Int int64 `protobuf:"varint,7,opt,name=int,proto3,oneof"`
}

type Value_Float struct {
	Float float64 `protobuf:"fixed64,8,opt,name=float,proto3,oneof"`
}

type Value_Boolean struct {
	Boolean bool `protobuf:"varint,9,opt,name=boolean,proto3,oneof"`
}

type Value_Time struct {
	Time *Value_Timestamp `protobuf:"bytes,10,opt,name=time,proto3,oneof"`
}

func (*Value_Raw) isValue_Value() {}

func (*Value_Str) isValue_Value() {}

func (*Value_Iri) isValue_Value() {}

func (*Value_Bnode) isValue_Value() {}

func (*Value_TypedStr) isValue_Value() {}

func (*Value_LangStr) isValue_Value() {}

func (*Value_Int) isValue_Value() {}

func (*Value_Float) isValue_Value() {}

func (*Value_Boolean) isValue_Value() {}

func (*Value_Time) isValue_Value() {}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full is set if encoder always writes every quad directions instead of
	// skipping duplicated values on each direction (except label) for subsequent quads.
	Full bool `protobuf:"varint,1,opt,name=full,proto3" json:"full,omitempty"`
	// NotStrict is set if encoder emits WireQuad instead of StrictQuad messages.
	NotStrict bool `protobuf:"varint,2,opt,name=not_strict,json=notStrict,proto3" json:"not_strict,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{6}
}

func (x *Header) GetFull() bool {
	if x != nil {
		return x.Full
	}
	return false
}

func (x *Header) GetNotStrict() bool {
	if x != nil {
		return x.NotStrict
	}
	return false
}

type StrictQuad_Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*StrictQuad_Ref_BnodeLabel
	//	*StrictQuad_Ref_Iri
	Value isStrictQuad_Ref_Value `protobuf_oneof:"value"`
}

func (x *StrictQuad_Ref) Reset() {
	*x = StrictQuad_Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrictQuad_Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrictQuad_Ref) ProtoMessage() {}

func (x *StrictQuad_Ref) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrictQuad_Ref.ProtoReflect.Descriptor instead.
func (*StrictQuad_Ref) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{3, 0}
}

func (m *StrictQuad_Ref) GetValue() isStrictQuad_Ref_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *StrictQuad_Ref) GetBnodeLabel() string {
	if x, ok := x.GetValue().(*StrictQuad_Ref_BnodeLabel); ok {
		return x.BnodeLabel
	}
	return ""
}

func (x *StrictQuad_Ref) GetIri() string {
	if x, ok := x.GetValue().(*StrictQuad_Ref_Iri); ok {
		return x.Iri
	}
	return ""
}

type isStrictQuad_Ref_Value interface {
	isStrictQuad_Ref_Value()
}

type StrictQuad_Ref_BnodeLabel struct {
	BnodeLabel string `protobuf:"bytes,2,opt,name=bnode_label,json=bnodeLabel,proto3,oneof"`
}

type StrictQuad_Ref_Iri struct {
	Iri string `protobuf:"bytes,3,opt,name=iri,proto3,oneof"`
}

func (*StrictQuad_Ref_BnodeLabel) isStrictQuad_Ref_Value() {}

func (*StrictQuad_Ref_Iri) isStrictQuad_Ref_Value() {}

type Value_TypedString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Value_TypedString) Reset() {
	*x = Value_TypedString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value_TypedString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value_TypedString) ProtoMessage() {}

func (x *Value_TypedString) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value_TypedString.ProtoReflect.Descriptor instead.
func (*Value_TypedString) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Value_TypedString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Value_TypedString) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Value_LangString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Lang  string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *Value_LangString) Reset() {
	*x = Value_LangString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value_LangString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value_LangString) ProtoMessage() {}

func (x *Value_LangString) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value_LangString.ProtoReflect.Descriptor instead.
func (*Value_LangString) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{5, 1}
}

func (x *Value_LangString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Value_LangString) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

// From https://github.com/golang/protobuf/blob/master/ptypes/timestamp/timestamp.proto
type Value_Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Value_Timestamp) Reset() {
	*x = Value_Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quads_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value_Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value_Timestamp) ProtoMessage() {}

func (x *Value_Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_quads_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value_Timestamp.ProtoReflect.Descriptor instead.
func (*Value_Timestamp) Descriptor() ([]byte, []int) {
	return file_quads_proto_rawDescGZIP(), []int{5, 2}
}

func (x *Value_Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Value_Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_quads_proto protoreflect.FileDescriptor

var file_quads_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x71, 0x75, 0x61, 0x64, 0x73, 0x22, 0xba, 0x02, 0x0a, 0x04, 0x51, 0x75, 0x61, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x32, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x71,
	0x75, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x36, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x30, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x57, 0x69, 0x72, 0x65, 0x51, 0x75, 0x61, 0x64, 0x12,
	0x27, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x71,
	0x75, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x71,
	0x75, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x22, 0x73, 0x0a, 0x0b, 0x57, 0x69, 0x72, 0x65, 0x51, 0x75, 0x61, 0x64, 0x52, 0x61, 0x77,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x96, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x51, 0x75, 0x61, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x51, 0x75, 0x61, 0x64, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x71, 0x75,
	0x61, 0x64, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x51, 0x75, 0x61, 0x64, 0x2e, 0x52,
	0x65, 0x66, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x51, 0x75, 0x61, 0x64, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x1a, 0x4b, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0a, 0x62, 0x6e, 0x6f, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x03,
	0x69, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x69, 0x72, 0x69,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22,
	0x75, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x51, 0x75, 0x61, 0x64, 0x52, 0x61, 0x77,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xfa, 0x03, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x03, 0x72, 0x61, 0x77, 0x12, 0x12, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x03, 0x69, 0x72, 0x69, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x69, 0x72, 0x69, 0x12, 0x16, 0x0a, 0x05,
	0x62, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x62,
	0x6e, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x48, 0x00, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72, 0x12, 0x35,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x61,
	0x6e, 0x67, 0x53, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x12, 0x1a, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x2d, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x71,
	0x75, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x37, 0x0a, 0x0b,
	0x54, 0x79, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x36, 0x0a, 0x0a, 0x4c, 0x61, 0x6e, 0x67, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x1a, 0x3b, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x3b, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x75, 0x6c,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x75, 0x63, 0x65, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x63, 0x61, 0x79, 0x6c, 0x65, 0x79, 0x2f, 0x71,
	0x75, 0x61, 0x64, 0x2f, 0x70, 0x71, 0x75, 0x61, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_quads_proto_rawDescOnce sync.Once
	file_quads_proto_rawDescData = file_quads_proto_rawDesc
)

func file_quads_proto_rawDescGZIP() []byte {
	file_quads_proto_rawDescOnce.Do(func() {
		file_quads_proto_rawDescData = protoimpl.X.CompressGZIP(file_quads_proto_rawDescData)
	})
	return file_quads_proto_rawDescData
}

var file_quads_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_quads_proto_goTypes = []interface{}{
	(*Quad)(nil),              // 0: pquads.Quad
	(*WireQuad)(nil),          // 1: pquads.WireQuad
	(*WireQuadRaw)(nil),       // 2: pquads.WireQuadRaw
	(*StrictQuad)(nil),        // 3: pquads.StrictQuad
	(*StrictQuadRaw)(nil),     // 4: pquads.StrictQuadRaw
	(*Value)(nil),             // 5: pquads.Value
	(*Header)(nil),            // 6: pquads.Header
	(*StrictQuad_Ref)(nil),    // 7: pquads.StrictQuad.Ref
	(*Value_TypedString)(nil), // 8: pquads.Value.TypedString
	(*Value_LangString)(nil),  // 9: pquads.Value.LangString
	(*Value_Timestamp)(nil),   // 10: pquads.Value.Timestamp
}
var file_quads_proto_depIdxs = []int32{
	5,  // 0: pquads.Quad.subject_value:type_name -> pquads.Value
	5,  // 1: pquads.Quad.predicate_value:type_name -> pquads.Value
	5,  // 2: pquads.Quad.object_value:type_name -> pquads.Value
	5,  // 3: pquads.Quad.label_value:type_name -> pquads.Value
	5,  // 4: pquads.WireQuad.subject:type_name -> pquads.Value
	5,  // 5: pquads.WireQuad.predicate:type_name -> pquads.Value
	5,  // 6: pquads.WireQuad.object:type_name -> pquads.Value
	5,  // 7: pquads.WireQuad.label:type_name -> pquads.Value
	7,  // 8: pquads.StrictQuad.subject:type_name -> pquads.StrictQuad.Ref
	7,  // 9: pquads.StrictQuad.predicate:type_name -> pquads.StrictQuad.Ref
	5,  // 10: pquads.StrictQuad.object:type_name -> pquads.Value
	7,  // 11: pquads.StrictQuad.label:type_name -> pquads.StrictQuad.Ref
	8,  // 12: pquads.Value.typed_str:type_name -> pquads.Value.TypedString
	9,  // 13: pquads.Value.lang_str:type_name -> pquads.Value.LangString
	10, // 14: pquads.Value.time:type_name -> pquads.Value.Timestamp
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_quads_proto_init() }
func file_quads_proto_init() {
	if File_quads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quad); i {
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
		file_quads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WireQuad); i {
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
		file_quads_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WireQuadRaw); i {
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
		file_quads_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrictQuad); i {
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
		file_quads_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrictQuadRaw); i {
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
		file_quads_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_quads_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_quads_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrictQuad_Ref); i {
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
		file_quads_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value_TypedString); i {
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
		file_quads_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value_LangString); i {
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
		file_quads_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value_Timestamp); i {
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
	file_quads_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Value_Raw)(nil),
		(*Value_Str)(nil),
		(*Value_Iri)(nil),
		(*Value_Bnode)(nil),
		(*Value_TypedStr)(nil),
		(*Value_LangStr)(nil),
		(*Value_Int)(nil),
		(*Value_Float)(nil),
		(*Value_Boolean)(nil),
		(*Value_Time)(nil),
	}
	file_quads_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*StrictQuad_Ref_BnodeLabel)(nil),
		(*StrictQuad_Ref_Iri)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_quads_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quads_proto_goTypes,
		DependencyIndexes: file_quads_proto_depIdxs,
		MessageInfos:      file_quads_proto_msgTypes,
	}.Build()
	File_quads_proto = out.File
	file_quads_proto_rawDesc = nil
	file_quads_proto_goTypes = nil
	file_quads_proto_depIdxs = nil
}
