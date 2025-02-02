// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: study_service/survey-response.proto

package api

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

type SurveyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key           string                `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ParticipantId string                `protobuf:"bytes,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	SubmittedAt   int64                 `protobuf:"varint,3,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	Responses     []*SurveyItemResponse `protobuf:"bytes,4,rep,name=responses,proto3" json:"responses,omitempty"`
	Context       map[string]string     `protobuf:"bytes,5,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // e.g. to store location and other context data
	VersionId     string                `protobuf:"bytes,6,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	OpenedAt      int64                 `protobuf:"varint,7,opt,name=opened_at,json=openedAt,proto3" json:"opened_at,omitempty"`
	Id            string                `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SurveyResponse) Reset() {
	*x = SurveyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_study_service_survey_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyResponse) ProtoMessage() {}

func (x *SurveyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_study_service_survey_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyResponse.ProtoReflect.Descriptor instead.
func (*SurveyResponse) Descriptor() ([]byte, []int) {
	return file_study_service_survey_response_proto_rawDescGZIP(), []int{0}
}

func (x *SurveyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SurveyResponse) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *SurveyResponse) GetSubmittedAt() int64 {
	if x != nil {
		return x.SubmittedAt
	}
	return 0
}

func (x *SurveyResponse) GetResponses() []*SurveyItemResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

func (x *SurveyResponse) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *SurveyResponse) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *SurveyResponse) GetOpenedAt() int64 {
	if x != nil {
		return x.OpenedAt
	}
	return 0
}

func (x *SurveyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SurveyResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*SurveyResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *SurveyResponses) Reset() {
	*x = SurveyResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_study_service_survey_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyResponses) ProtoMessage() {}

func (x *SurveyResponses) ProtoReflect() protoreflect.Message {
	mi := &file_study_service_survey_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyResponses.ProtoReflect.Descriptor instead.
func (*SurveyResponses) Descriptor() ([]byte, []int) {
	return file_study_service_survey_response_proto_rawDescGZIP(), []int{1}
}

func (x *SurveyResponses) GetResponses() []*SurveyResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

type SurveyItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // reference to the item
	Meta *ResponseMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	// for item groups:
	Items []*SurveyItemResponse `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	// for single items:
	Response         *ResponseItem `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	ConfidentialMode string        `protobuf:"bytes,5,opt,name=confidential_mode,json=confidentialMode,proto3" json:"confidential_mode,omitempty"`
}

func (x *SurveyItemResponse) Reset() {
	*x = SurveyItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_study_service_survey_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyItemResponse) ProtoMessage() {}

func (x *SurveyItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_study_service_survey_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyItemResponse.ProtoReflect.Descriptor instead.
func (*SurveyItemResponse) Descriptor() ([]byte, []int) {
	return file_study_service_survey_response_proto_rawDescGZIP(), []int{2}
}

func (x *SurveyItemResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SurveyItemResponse) GetMeta() *ResponseMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SurveyItemResponse) GetItems() []*SurveyItemResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SurveyItemResponse) GetResponse() *ResponseItem {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *SurveyItemResponse) GetConfidentialMode() string {
	if x != nil {
		return x.ConfidentialMode
	}
	return ""
}

type ResponseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Dtype string `protobuf:"bytes,3,opt,name=dtype,proto3" json:"dtype,omitempty"`
	// For response option groups:
	Items []*ResponseItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ResponseItem) Reset() {
	*x = ResponseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_study_service_survey_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseItem) ProtoMessage() {}

func (x *ResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_study_service_survey_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseItem.ProtoReflect.Descriptor instead.
func (*ResponseItem) Descriptor() ([]byte, []int) {
	return file_study_service_survey_response_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ResponseItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ResponseItem) GetDtype() string {
	if x != nil {
		return x.Dtype
	}
	return ""
}

func (x *ResponseItem) GetItems() []*ResponseItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ResponseMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position   int32  `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	LocaleCode string `protobuf:"bytes,2,opt,name=locale_code,json=localeCode,proto3" json:"locale_code,omitempty"`
	Version    int32  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// timestamps:
	Rendered  []int64 `protobuf:"varint,4,rep,packed,name=rendered,proto3" json:"rendered,omitempty"`
	Displayed []int64 `protobuf:"varint,5,rep,packed,name=displayed,proto3" json:"displayed,omitempty"`
	Responded []int64 `protobuf:"varint,6,rep,packed,name=responded,proto3" json:"responded,omitempty"`
}

func (x *ResponseMeta) Reset() {
	*x = ResponseMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_study_service_survey_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMeta) ProtoMessage() {}

func (x *ResponseMeta) ProtoReflect() protoreflect.Message {
	mi := &file_study_service_survey_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMeta.ProtoReflect.Descriptor instead.
func (*ResponseMeta) Descriptor() ([]byte, []int) {
	return file_study_service_survey_response_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseMeta) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *ResponseMeta) GetLocaleCode() string {
	if x != nil {
		return x.LocaleCode
	}
	return ""
}

func (x *ResponseMeta) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ResponseMeta) GetRendered() []int64 {
	if x != nil {
		return x.Rendered
	}
	return nil
}

func (x *ResponseMeta) GetDisplayed() []int64 {
	if x != nil {
		return x.Displayed
	}
	return nil
}

func (x *ResponseMeta) GetResponded() []int64 {
	if x != nil {
		return x.Responded
	}
	return nil
}

var File_study_service_survey_response_proto protoreflect.FileDescriptor

var file_study_service_survey_response_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61,
	0x6e, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x95, 0x03, 0x0a, 0x0e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x4c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61, 0x6e,
	0x65, 0x74, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x51,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61, 0x6e, 0x65, 0x74, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x3a, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a, 0x0f, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61, 0x6e, 0x65, 0x74, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x9d, 0x02, 0x0a, 0x12, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x44, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69,
	0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x44, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a,
	0x61, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61,
	0x6e, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x08, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61, 0x6e, 0x65, 0x74,
	0x2f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_study_service_survey_response_proto_rawDescOnce sync.Once
	file_study_service_survey_response_proto_rawDescData = file_study_service_survey_response_proto_rawDesc
)

func file_study_service_survey_response_proto_rawDescGZIP() []byte {
	file_study_service_survey_response_proto_rawDescOnce.Do(func() {
		file_study_service_survey_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_study_service_survey_response_proto_rawDescData)
	})
	return file_study_service_survey_response_proto_rawDescData
}

var file_study_service_survey_response_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_study_service_survey_response_proto_goTypes = []interface{}{
	(*SurveyResponse)(nil),     // 0: influenzanet.study_service.SurveyResponse
	(*SurveyResponses)(nil),    // 1: influenzanet.study_service.SurveyResponses
	(*SurveyItemResponse)(nil), // 2: influenzanet.study_service.SurveyItemResponse
	(*ResponseItem)(nil),       // 3: influenzanet.study_service.ResponseItem
	(*ResponseMeta)(nil),       // 4: influenzanet.study_service.ResponseMeta
	nil,                        // 5: influenzanet.study_service.SurveyResponse.ContextEntry
}
var file_study_service_survey_response_proto_depIdxs = []int32{
	2, // 0: influenzanet.study_service.SurveyResponse.responses:type_name -> influenzanet.study_service.SurveyItemResponse
	5, // 1: influenzanet.study_service.SurveyResponse.context:type_name -> influenzanet.study_service.SurveyResponse.ContextEntry
	0, // 2: influenzanet.study_service.SurveyResponses.responses:type_name -> influenzanet.study_service.SurveyResponse
	4, // 3: influenzanet.study_service.SurveyItemResponse.meta:type_name -> influenzanet.study_service.ResponseMeta
	2, // 4: influenzanet.study_service.SurveyItemResponse.items:type_name -> influenzanet.study_service.SurveyItemResponse
	3, // 5: influenzanet.study_service.SurveyItemResponse.response:type_name -> influenzanet.study_service.ResponseItem
	3, // 6: influenzanet.study_service.ResponseItem.items:type_name -> influenzanet.study_service.ResponseItem
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_study_service_survey_response_proto_init() }
func file_study_service_survey_response_proto_init() {
	if File_study_service_survey_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_study_service_survey_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyResponse); i {
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
		file_study_service_survey_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyResponses); i {
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
		file_study_service_survey_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyItemResponse); i {
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
		file_study_service_survey_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseItem); i {
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
		file_study_service_survey_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMeta); i {
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
			RawDescriptor: file_study_service_survey_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_study_service_survey_response_proto_goTypes,
		DependencyIndexes: file_study_service_survey_response_proto_depIdxs,
		MessageInfos:      file_study_service_survey_response_proto_msgTypes,
	}.Build()
	File_study_service_survey_response_proto = out.File
	file_study_service_survey_response_proto_rawDesc = nil
	file_study_service_survey_response_proto_goTypes = nil
	file_study_service_survey_response_proto_depIdxs = nil
}
