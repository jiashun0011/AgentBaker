// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pkg/proto/nbcontract/v1/customcloudconfig.proto

package nbcontractv1

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

type CustomCloudConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status shows whether CustomCloudConfig is enabled or not
	Status *FeatureState `protobuf:"varint,1,opt,name=status,proto3,enum=nbcontract.v1.FeatureState,oneof" json:"status,omitempty"`
	// InitFilePath is the path to the file that contains the init script
	InitFilePath *string `protobuf:"bytes,2,opt,name=init_file_path,json=initFilePath,proto3,oneof" json:"init_file_path,omitempty"`
	// RepoDepotEndpoint is the endpoint of the repo depot
	RepoDepotEndpoint    *string `protobuf:"bytes,3,opt,name=repo_depot_endpoint,json=repoDepotEndpoint,proto3,oneof" json:"repo_depot_endpoint,omitempty"`
	TargetEnvironment    string  `protobuf:"bytes,4,opt,name=target_environment,json=targetEnvironment,proto3" json:"target_environment,omitempty"`
	TargetCloud          string  `protobuf:"bytes,5,opt,name=target_cloud,json=targetCloud,proto3" json:"target_cloud,omitempty"`                                // can probably get rid of this, analyze more
	CustomEnvJsonContent string  `protobuf:"bytes,6,opt,name=custom_env_json_content,json=customEnvJsonContent,proto3" json:"custom_env_json_content,omitempty"` // can be generated on the VHD, also rename
}

func (x *CustomCloudConfig) Reset() {
	*x = CustomCloudConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nbcontract_v1_customcloudconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomCloudConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomCloudConfig) ProtoMessage() {}

func (x *CustomCloudConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nbcontract_v1_customcloudconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomCloudConfig.ProtoReflect.Descriptor instead.
func (*CustomCloudConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDescGZIP(), []int{0}
}

func (x *CustomCloudConfig) GetStatus() FeatureState {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return FeatureState_FEATURE_STATE_UNSPECIFIED
}

func (x *CustomCloudConfig) GetInitFilePath() string {
	if x != nil && x.InitFilePath != nil {
		return *x.InitFilePath
	}
	return ""
}

func (x *CustomCloudConfig) GetRepoDepotEndpoint() string {
	if x != nil && x.RepoDepotEndpoint != nil {
		return *x.RepoDepotEndpoint
	}
	return ""
}

func (x *CustomCloudConfig) GetTargetEnvironment() string {
	if x != nil {
		return x.TargetEnvironment
	}
	return ""
}

func (x *CustomCloudConfig) GetTargetCloud() string {
	if x != nil {
		return x.TargetCloud
	}
	return ""
}

func (x *CustomCloudConfig) GetCustomEnvJsonContent() string {
	if x != nil {
		return x.CustomEnvJsonContent
	}
	return ""
}

var File_pkg_proto_nbcontract_v1_customcloudconfig_proto protoreflect.FileDescriptor

var file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a,
	0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e,
	0x69, 0x6e, 0x69, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x5f,
	0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x70, 0x6f,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x12,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x35,
	0x0a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x76, 0x4a, 0x73, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x64, 0x65, 0x70,
	0x6f, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0xc2, 0x01, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x4e, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0e, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDescOnce sync.Once
	file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDescData = file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDesc
)

func file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDescGZIP() []byte {
	file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDescOnce.Do(func() {
		file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDescData)
	})
	return file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDescData
}

var file_pkg_proto_nbcontract_v1_customcloudconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_nbcontract_v1_customcloudconfig_proto_goTypes = []interface{}{
	(*CustomCloudConfig)(nil), // 0: nbcontract.v1.CustomCloudConfig
	(FeatureState)(0),         // 1: nbcontract.v1.FeatureState
}
var file_pkg_proto_nbcontract_v1_customcloudconfig_proto_depIdxs = []int32{
	1, // 0: nbcontract.v1.CustomCloudConfig.status:type_name -> nbcontract.v1.FeatureState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_nbcontract_v1_customcloudconfig_proto_init() }
func file_pkg_proto_nbcontract_v1_customcloudconfig_proto_init() {
	if File_pkg_proto_nbcontract_v1_customcloudconfig_proto != nil {
		return
	}
	file_pkg_proto_nbcontract_v1_featurestate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_nbcontract_v1_customcloudconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomCloudConfig); i {
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
	file_pkg_proto_nbcontract_v1_customcloudconfig_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_nbcontract_v1_customcloudconfig_proto_goTypes,
		DependencyIndexes: file_pkg_proto_nbcontract_v1_customcloudconfig_proto_depIdxs,
		MessageInfos:      file_pkg_proto_nbcontract_v1_customcloudconfig_proto_msgTypes,
	}.Build()
	File_pkg_proto_nbcontract_v1_customcloudconfig_proto = out.File
	file_pkg_proto_nbcontract_v1_customcloudconfig_proto_rawDesc = nil
	file_pkg_proto_nbcontract_v1_customcloudconfig_proto_goTypes = nil
	file_pkg_proto_nbcontract_v1_customcloudconfig_proto_depIdxs = nil
}
