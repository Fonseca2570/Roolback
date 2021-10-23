// Copyright 2015 gRPC authors.
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: proto/service1.proto

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

// The request message containing the user's name.
type UpdateFirstNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersId   int64  `protobuf:"varint,1,opt,name=users_id,json=usersId,proto3" json:"users_id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
}

func (x *UpdateFirstNameRequest) Reset() {
	*x = UpdateFirstNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFirstNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFirstNameRequest) ProtoMessage() {}

func (x *UpdateFirstNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFirstNameRequest.ProtoReflect.Descriptor instead.
func (*UpdateFirstNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_service1_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateFirstNameRequest) GetUsersId() int64 {
	if x != nil {
		return x.UsersId
	}
	return 0
}

func (x *UpdateFirstNameRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

// The response message containing the greetings
type UpdateFirstNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error     string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	TableName string `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	RowId     int64  `protobuf:"varint,3,opt,name=row_id,json=rowId,proto3" json:"row_id,omitempty"`
}

func (x *UpdateFirstNameResponse) Reset() {
	*x = UpdateFirstNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFirstNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFirstNameResponse) ProtoMessage() {}

func (x *UpdateFirstNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFirstNameResponse.ProtoReflect.Descriptor instead.
func (*UpdateFirstNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_service1_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateFirstNameResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *UpdateFirstNameResponse) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *UpdateFirstNameResponse) GetRowId() int64 {
	if x != nil {
		return x.RowId
	}
	return 0
}

type RollbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Id        int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RollbackRequest) Reset() {
	*x = RollbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackRequest) ProtoMessage() {}

func (x *RollbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackRequest.ProtoReflect.Descriptor instead.
func (*RollbackRequest) Descriptor() ([]byte, []int) {
	return file_proto_service1_proto_rawDescGZIP(), []int{2}
}

func (x *RollbackRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *RollbackRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RollbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RollbackResponse) Reset() {
	*x = RollbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackResponse) ProtoMessage() {}

func (x *RollbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackResponse.ProtoReflect.Descriptor instead.
func (*RollbackResponse) Descriptor() ([]byte, []int) {
	return file_proto_service1_proto_rawDescGZIP(), []int{3}
}

func (x *RollbackResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_service1_proto protoreflect.FileDescriptor

var file_proto_service1_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x65, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x72, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x6f,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x9c, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5a, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service1_proto_rawDescOnce sync.Once
	file_proto_service1_proto_rawDescData = file_proto_service1_proto_rawDesc
)

func file_proto_service1_proto_rawDescGZIP() []byte {
	file_proto_service1_proto_rawDescOnce.Do(func() {
		file_proto_service1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service1_proto_rawDescData)
	})
	return file_proto_service1_proto_rawDescData
}

var file_proto_service1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_service1_proto_goTypes = []interface{}{
	(*UpdateFirstNameRequest)(nil),  // 0: proto.UpdateFirstNameRequest
	(*UpdateFirstNameResponse)(nil), // 1: proto.UpdateFirstNameResponse
	(*RollbackRequest)(nil),         // 2: proto.RollbackRequest
	(*RollbackResponse)(nil),        // 3: proto.RollbackResponse
}
var file_proto_service1_proto_depIdxs = []int32{
	0, // 0: proto.Service.UpdateFirstName:input_type -> proto.UpdateFirstNameRequest
	2, // 1: proto.Service.Rollback:input_type -> proto.RollbackRequest
	1, // 2: proto.Service.UpdateFirstName:output_type -> proto.UpdateFirstNameResponse
	3, // 3: proto.Service.Rollback:output_type -> proto.RollbackResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service1_proto_init() }
func file_proto_service1_proto_init() {
	if File_proto_service1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFirstNameRequest); i {
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
		file_proto_service1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFirstNameResponse); i {
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
		file_proto_service1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackRequest); i {
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
		file_proto_service1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackResponse); i {
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
			RawDescriptor: file_proto_service1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service1_proto_goTypes,
		DependencyIndexes: file_proto_service1_proto_depIdxs,
		MessageInfos:      file_proto_service1_proto_msgTypes,
	}.Build()
	File_proto_service1_proto = out.File
	file_proto_service1_proto_rawDesc = nil
	file_proto_service1_proto_goTypes = nil
	file_proto_service1_proto_depIdxs = nil
}
