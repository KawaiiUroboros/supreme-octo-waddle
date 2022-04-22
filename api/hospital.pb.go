// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: hospital.proto

package api_pb

import (
	context "context"
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

//create team request
type CreateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails []string `protobuf:"bytes,1,rep,name=emails,proto3" json:"emails,omitempty"`
	City   string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *CreateTeamRequest) Reset() {
	*x = CreateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamRequest) ProtoMessage() {}

func (x *CreateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTeamRequest) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *CreateTeamRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

//create team response
type CreateTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateTeamResponse) Reset() {
	*x = CreateTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamResponse) ProtoMessage() {}

func (x *CreateTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamResponse.ProtoReflect.Descriptor instead.
func (*CreateTeamResponse) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTeamResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateTeamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

//create incident request
type CreateIncidentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	City        string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *CreateIncidentRequest) Reset() {
	*x = CreateIncidentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIncidentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIncidentRequest) ProtoMessage() {}

func (x *CreateIncidentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIncidentRequest.ProtoReflect.Descriptor instead.
func (*CreateIncidentRequest) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{2}
}

func (x *CreateIncidentRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateIncidentRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateIncidentRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

//create incident response
type CreateIncidentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateIncidentResponse) Reset() {
	*x = CreateIncidentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIncidentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIncidentResponse) ProtoMessage() {}

func (x *CreateIncidentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIncidentResponse.ProtoReflect.Descriptor instead.
func (*CreateIncidentResponse) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{3}
}

func (x *CreateIncidentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateIncidentResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

//get city and address by external user id request
type GetCityAndAddressByExternalUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalUserId string `protobuf:"bytes,1,opt,name=external_user_id,json=externalUserId,proto3" json:"external_user_id,omitempty"`
}

func (x *GetCityAndAddressByExternalUserIdRequest) Reset() {
	*x = GetCityAndAddressByExternalUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityAndAddressByExternalUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityAndAddressByExternalUserIdRequest) ProtoMessage() {}

func (x *GetCityAndAddressByExternalUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityAndAddressByExternalUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetCityAndAddressByExternalUserIdRequest) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{4}
}

func (x *GetCityAndAddressByExternalUserIdRequest) GetExternalUserId() string {
	if x != nil {
		return x.ExternalUserId
	}
	return ""
}

//get city and address by external user id response
type GetCityAndAddressByExternalUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City    string `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetCityAndAddressByExternalUserIdResponse) Reset() {
	*x = GetCityAndAddressByExternalUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityAndAddressByExternalUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityAndAddressByExternalUserIdResponse) ProtoMessage() {}

func (x *GetCityAndAddressByExternalUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityAndAddressByExternalUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetCityAndAddressByExternalUserIdResponse) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{5}
}

func (x *GetCityAndAddressByExternalUserIdResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetCityAndAddressByExternalUserIdResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_hospital_proto protoreflect.FileDescriptor

var file_hospital_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x44, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x63, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x28, 0x47, 0x65,
	0x74, 0x43, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x59, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xe5, 0x03, 0x0a, 0x0f,
	0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x72, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x27, 0x2e,
	0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b,
	0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x82, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b,
	0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61,
	0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x69, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xd8, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74,
	0x43, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e,
	0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x64,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f,
	0x2e, 0x61, 0x63, 0x68, 0x69, 0x72, 0x69, 0x6b, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x64,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x20, 0x5a, 0x1e, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2d, 0x6f,
	0x63, 0x74, 0x6f, 0x2d, 0x77, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61,
	0x70, 0x69, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hospital_proto_rawDescOnce sync.Once
	file_hospital_proto_rawDescData = file_hospital_proto_rawDesc
)

func file_hospital_proto_rawDescGZIP() []byte {
	file_hospital_proto_rawDescOnce.Do(func() {
		file_hospital_proto_rawDescData = protoimpl.X.CompressGZIP(file_hospital_proto_rawDescData)
	})
	return file_hospital_proto_rawDescData
}

var file_hospital_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hospital_proto_goTypes = []interface{}{
	(*CreateTeamRequest)(nil),                         // 0: achirikova.emergence.CreateTeamRequest
	(*CreateTeamResponse)(nil),                        // 1: achirikova.emergence.CreateTeamResponse
	(*CreateIncidentRequest)(nil),                     // 2: achirikova.emergence.CreateIncidentRequest
	(*CreateIncidentResponse)(nil),                    // 3: achirikova.emergence.CreateIncidentResponse
	(*GetCityAndAddressByExternalUserIdRequest)(nil),  // 4: achirikova.emergence.GetCityAndAddressByExternalUserIdRequest
	(*GetCityAndAddressByExternalUserIdResponse)(nil), // 5: achirikova.emergence.GetCityAndAddressByExternalUserIdResponse
}
var file_hospital_proto_depIdxs = []int32{
	0, // 0: achirikova.emergence.HospitalService.CreateTeam:input_type -> achirikova.emergence.CreateTeamRequest
	2, // 1: achirikova.emergence.HospitalService.CreateIncident:input_type -> achirikova.emergence.CreateIncidentRequest
	4, // 2: achirikova.emergence.HospitalService.GetCityAndAddressByExternalUserId:input_type -> achirikova.emergence.GetCityAndAddressByExternalUserIdRequest
	1, // 3: achirikova.emergence.HospitalService.CreateTeam:output_type -> achirikova.emergence.CreateTeamResponse
	3, // 4: achirikova.emergence.HospitalService.CreateIncident:output_type -> achirikova.emergence.CreateIncidentResponse
	5, // 5: achirikova.emergence.HospitalService.GetCityAndAddressByExternalUserId:output_type -> achirikova.emergence.GetCityAndAddressByExternalUserIdResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hospital_proto_init() }
func file_hospital_proto_init() {
	if File_hospital_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hospital_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamRequest); i {
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
		file_hospital_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamResponse); i {
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
		file_hospital_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIncidentRequest); i {
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
		file_hospital_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIncidentResponse); i {
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
		file_hospital_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCityAndAddressByExternalUserIdRequest); i {
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
		file_hospital_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCityAndAddressByExternalUserIdResponse); i {
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
			RawDescriptor: file_hospital_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hospital_proto_goTypes,
		DependencyIndexes: file_hospital_proto_depIdxs,
		MessageInfos:      file_hospital_proto_msgTypes,
	}.Build()
	File_hospital_proto = out.File
	file_hospital_proto_rawDesc = nil
	file_hospital_proto_goTypes = nil
	file_hospital_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HospitalServiceClient is the client API for HospitalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HospitalServiceClient interface {
	//create team
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error)
	//create incident
	CreateIncident(ctx context.Context, in *CreateIncidentRequest, opts ...grpc.CallOption) (*CreateIncidentResponse, error)
	//Get city and address by external user id
	GetCityAndAddressByExternalUserId(ctx context.Context, in *GetCityAndAddressByExternalUserIdRequest, opts ...grpc.CallOption) (*GetCityAndAddressByExternalUserIdResponse, error)
}

type hospitalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalServiceClient(cc grpc.ClientConnInterface) HospitalServiceClient {
	return &hospitalServiceClient{cc}
}

func (c *hospitalServiceClient) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error) {
	out := new(CreateTeamResponse)
	err := c.cc.Invoke(ctx, "/achirikova.emergence.HospitalService/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalServiceClient) CreateIncident(ctx context.Context, in *CreateIncidentRequest, opts ...grpc.CallOption) (*CreateIncidentResponse, error) {
	out := new(CreateIncidentResponse)
	err := c.cc.Invoke(ctx, "/achirikova.emergence.HospitalService/CreateIncident", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalServiceClient) GetCityAndAddressByExternalUserId(ctx context.Context, in *GetCityAndAddressByExternalUserIdRequest, opts ...grpc.CallOption) (*GetCityAndAddressByExternalUserIdResponse, error) {
	out := new(GetCityAndAddressByExternalUserIdResponse)
	err := c.cc.Invoke(ctx, "/achirikova.emergence.HospitalService/GetCityAndAddressByExternalUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServiceServer is the server API for HospitalService service.
type HospitalServiceServer interface {
	//create team
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error)
	//create incident
	CreateIncident(context.Context, *CreateIncidentRequest) (*CreateIncidentResponse, error)
	//Get city and address by external user id
	GetCityAndAddressByExternalUserId(context.Context, *GetCityAndAddressByExternalUserIdRequest) (*GetCityAndAddressByExternalUserIdResponse, error)
}

// UnimplementedHospitalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHospitalServiceServer struct {
}

func (*UnimplementedHospitalServiceServer) CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (*UnimplementedHospitalServiceServer) CreateIncident(context.Context, *CreateIncidentRequest) (*CreateIncidentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIncident not implemented")
}
func (*UnimplementedHospitalServiceServer) GetCityAndAddressByExternalUserId(context.Context, *GetCityAndAddressByExternalUserIdRequest) (*GetCityAndAddressByExternalUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCityAndAddressByExternalUserId not implemented")
}

func RegisterHospitalServiceServer(s *grpc.Server, srv HospitalServiceServer) {
	s.RegisterService(&_HospitalService_serviceDesc, srv)
}

func _HospitalService_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/achirikova.emergence.HospitalService/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).CreateTeam(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HospitalService_CreateIncident_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIncidentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).CreateIncident(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/achirikova.emergence.HospitalService/CreateIncident",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).CreateIncident(ctx, req.(*CreateIncidentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HospitalService_GetCityAndAddressByExternalUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCityAndAddressByExternalUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).GetCityAndAddressByExternalUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/achirikova.emergence.HospitalService/GetCityAndAddressByExternalUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).GetCityAndAddressByExternalUserId(ctx, req.(*GetCityAndAddressByExternalUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HospitalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "achirikova.emergence.HospitalService",
	HandlerType: (*HospitalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _HospitalService_CreateTeam_Handler,
		},
		{
			MethodName: "CreateIncident",
			Handler:    _HospitalService_CreateIncident_Handler,
		},
		{
			MethodName: "GetCityAndAddressByExternalUserId",
			Handler:    _HospitalService_GetCityAndAddressByExternalUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hospital.proto",
}