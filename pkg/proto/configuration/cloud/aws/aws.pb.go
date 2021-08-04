// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.1
// source: pkg/proto/configuration/cloud/aws/aws.proto

package aws

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StaticCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKeyId     string `protobuf:"bytes,1,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	SecretAccessKey string `protobuf:"bytes,2,opt,name=secret_access_key,json=secretAccessKey,proto3" json:"secret_access_key,omitempty"`
}

func (x *StaticCredentials) Reset() {
	*x = StaticCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticCredentials) ProtoMessage() {}

func (x *StaticCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticCredentials.ProtoReflect.Descriptor instead.
func (*StaticCredentials) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescGZIP(), []int{0}
}

func (x *StaticCredentials) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *StaticCredentials) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

type SessionConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint          string             `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Region            string             `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	DisableSsl        bool               `protobuf:"varint,3,opt,name=disable_ssl,json=disableSsl,proto3" json:"disable_ssl,omitempty"`
	S3ForcePathStyle  bool               `protobuf:"varint,4,opt,name=s3_force_path_style,json=s3ForcePathStyle,proto3" json:"s3_force_path_style,omitempty"`
	StaticCredentials *StaticCredentials `protobuf:"bytes,5,opt,name=static_credentials,json=staticCredentials,proto3" json:"static_credentials,omitempty"`
}

func (x *SessionConfiguration) Reset() {
	*x = SessionConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionConfiguration) ProtoMessage() {}

func (x *SessionConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionConfiguration.ProtoReflect.Descriptor instead.
func (*SessionConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescGZIP(), []int{1}
}

func (x *SessionConfiguration) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *SessionConfiguration) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *SessionConfiguration) GetDisableSsl() bool {
	if x != nil {
		return x.DisableSsl
	}
	return false
}

func (x *SessionConfiguration) GetS3ForcePathStyle() bool {
	if x != nil {
		return x.S3ForcePathStyle
	}
	return false
}

func (x *SessionConfiguration) GetStaticCredentials() *StaticCredentials {
	if x != nil {
		return x.StaticCredentials
	}
	return nil
}

var File_pkg_proto_configuration_cloud_aws_aws_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_cloud_aws_aws_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x77, 0x73,
	0x22, 0x63, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x22, 0xff, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x73,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x73, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x73, 0x33, 0x5f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x10, 0x73, 0x33, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x53, 0x74, 0x79,
	0x6c, 0x65, 0x12, 0x63, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x77, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f,
	0x62, 0x62, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescData = file_pkg_proto_configuration_cloud_aws_aws_proto_rawDesc
)

func file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescData)
	})
	return file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescData
}

var file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_configuration_cloud_aws_aws_proto_goTypes = []interface{}{
	(*StaticCredentials)(nil),    // 0: buildbarn.configuration.cloud.aws.StaticCredentials
	(*SessionConfiguration)(nil), // 1: buildbarn.configuration.cloud.aws.SessionConfiguration
}
var file_pkg_proto_configuration_cloud_aws_aws_proto_depIdxs = []int32{
	0, // 0: buildbarn.configuration.cloud.aws.SessionConfiguration.static_credentials:type_name -> buildbarn.configuration.cloud.aws.StaticCredentials
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_cloud_aws_aws_proto_init() }
func file_pkg_proto_configuration_cloud_aws_aws_proto_init() {
	if File_pkg_proto_configuration_cloud_aws_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticCredentials); i {
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
		file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionConfiguration); i {
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
			RawDescriptor: file_pkg_proto_configuration_cloud_aws_aws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_cloud_aws_aws_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_cloud_aws_aws_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_cloud_aws_aws_proto = out.File
	file_pkg_proto_configuration_cloud_aws_aws_proto_rawDesc = nil
	file_pkg_proto_configuration_cloud_aws_aws_proto_goTypes = nil
	file_pkg_proto_configuration_cloud_aws_aws_proto_depIdxs = nil
}
