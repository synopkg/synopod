// Copyright (c) 2024 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/experimental/v1/scm.proto

package v1

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

type GetSuggestedRepoURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSuggestedRepoURLsRequest) Reset() {
	*x = GetSuggestedRepoURLsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_scm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuggestedRepoURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuggestedRepoURLsRequest) ProtoMessage() {}

func (x *GetSuggestedRepoURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_scm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuggestedRepoURLsRequest.ProtoReflect.Descriptor instead.
func (*GetSuggestedRepoURLsRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_scm_proto_rawDescGZIP(), []int{0}
}

type GetSuggestedRepoURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repos []string `protobuf:"bytes,1,rep,name=repos,proto3" json:"repos,omitempty"`
}

func (x *GetSuggestedRepoURLsResponse) Reset() {
	*x = GetSuggestedRepoURLsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_scm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuggestedRepoURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuggestedRepoURLsResponse) ProtoMessage() {}

func (x *GetSuggestedRepoURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_scm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuggestedRepoURLsResponse.ProtoReflect.Descriptor instead.
func (*GetSuggestedRepoURLsResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_scm_proto_rawDescGZIP(), []int{1}
}

func (x *GetSuggestedRepoURLsResponse) GetRepos() []string {
	if x != nil {
		return x.Repos
	}
	return nil
}

var File_gitpod_experimental_v1_scm_proto protoreflect.FileDescriptor

var file_gitpod_experimental_v1_scm_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x52,
	0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x52, 0x4c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x32,
	0x92, 0x01, 0x0a, 0x0a, 0x53, 0x43, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x33, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70,
	0x6f, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x6b, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_experimental_v1_scm_proto_rawDescOnce sync.Once
	file_gitpod_experimental_v1_scm_proto_rawDescData = file_gitpod_experimental_v1_scm_proto_rawDesc
)

func file_gitpod_experimental_v1_scm_proto_rawDescGZIP() []byte {
	file_gitpod_experimental_v1_scm_proto_rawDescOnce.Do(func() {
		file_gitpod_experimental_v1_scm_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_experimental_v1_scm_proto_rawDescData)
	})
	return file_gitpod_experimental_v1_scm_proto_rawDescData
}

var file_gitpod_experimental_v1_scm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gitpod_experimental_v1_scm_proto_goTypes = []interface{}{
	(*GetSuggestedRepoURLsRequest)(nil),  // 0: gitpod.experimental.v1.GetSuggestedRepoURLsRequest
	(*GetSuggestedRepoURLsResponse)(nil), // 1: gitpod.experimental.v1.GetSuggestedRepoURLsResponse
}
var file_gitpod_experimental_v1_scm_proto_depIdxs = []int32{
	0, // 0: gitpod.experimental.v1.SCMService.GetSuggestedRepoURLs:input_type -> gitpod.experimental.v1.GetSuggestedRepoURLsRequest
	1, // 1: gitpod.experimental.v1.SCMService.GetSuggestedRepoURLs:output_type -> gitpod.experimental.v1.GetSuggestedRepoURLsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gitpod_experimental_v1_scm_proto_init() }
func file_gitpod_experimental_v1_scm_proto_init() {
	if File_gitpod_experimental_v1_scm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitpod_experimental_v1_scm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuggestedRepoURLsRequest); i {
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
		file_gitpod_experimental_v1_scm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuggestedRepoURLsResponse); i {
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
			RawDescriptor: file_gitpod_experimental_v1_scm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gitpod_experimental_v1_scm_proto_goTypes,
		DependencyIndexes: file_gitpod_experimental_v1_scm_proto_depIdxs,
		MessageInfos:      file_gitpod_experimental_v1_scm_proto_msgTypes,
	}.Build()
	File_gitpod_experimental_v1_scm_proto = out.File
	file_gitpod_experimental_v1_scm_proto_rawDesc = nil
	file_gitpod_experimental_v1_scm_proto_goTypes = nil
	file_gitpod_experimental_v1_scm_proto_depIdxs = nil
}
