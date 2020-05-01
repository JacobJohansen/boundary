// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: controller/api/resources/hosts/v1/awsec2_host_catalog.proto

package hosts

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type AwsEc2HostCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Canonical path of the resource from the API's base URI
	// Output only.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The type of the resource, to help differentiate schemas
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Friendly name, if set
	FriendlyName *wrappers.StringValue `protobuf:"bytes,3,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	// The time this host was created
	// Output only.
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// The time this host was last updated
	// Output only.
	UpdatedTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// Whether the catalog is disabled
	Disabled *wrappers.BoolValue `protobuf:"bytes,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// TODO: Figure out how to provide relevant information regarding Host and
	// Set lists on a catalog.
	// The AWS regions from which this catalog will retrieve the EC2 instances.
	Regions []string `protobuf:"bytes,7,rep,name=regions,proto3" json:"regions,omitempty"`
	// The access key used for authenticating with AWS when retrieving EC2 instance details.
	AccessKey *wrappers.StringValue `protobuf:"bytes,8,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// Input only.
	SecretKey *wrappers.StringValue `protobuf:"bytes,9,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	// This value will never be returned in a response.
	Rotate *wrappers.BoolValue `protobuf:"bytes,10,opt,name=rotate,proto3" json:"rotate,omitempty"`
}

func (x *AwsEc2HostCatalog) Reset() {
	*x = AwsEc2HostCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsEc2HostCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsEc2HostCatalog) ProtoMessage() {}

func (x *AwsEc2HostCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsEc2HostCatalog.ProtoReflect.Descriptor instead.
func (*AwsEc2HostCatalog) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *AwsEc2HostCatalog) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AwsEc2HostCatalog) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AwsEc2HostCatalog) GetFriendlyName() *wrappers.StringValue {
	if x != nil {
		return x.FriendlyName
	}
	return nil
}

func (x *AwsEc2HostCatalog) GetCreatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *AwsEc2HostCatalog) GetUpdatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *AwsEc2HostCatalog) GetDisabled() *wrappers.BoolValue {
	if x != nil {
		return x.Disabled
	}
	return nil
}

func (x *AwsEc2HostCatalog) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *AwsEc2HostCatalog) GetAccessKey() *wrappers.StringValue {
	if x != nil {
		return x.AccessKey
	}
	return nil
}

func (x *AwsEc2HostCatalog) GetSecretKey() *wrappers.StringValue {
	if x != nil {
		return x.SecretKey
	}
	return nil
}

func (x *AwsEc2HostCatalog) GetRotate() *wrappers.BoolValue {
	if x != nil {
		return x.Rotate
	}
	return nil
}

var File_controller_api_resources_hosts_v1_awsec2_host_catalog_proto protoreflect.FileDescriptor

var file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x65, 0x63, 0x32, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfc, 0x03, 0x0a, 0x11, 0x41, 0x77, 0x73, 0x45, 0x63, 0x32, 0x48, 0x6f, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x41, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x36, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12,
	0x3b, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x06,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x74, 0x6f,
	0x77, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x3b,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDescOnce sync.Once
	file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDescData = file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDesc
)

func file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDescGZIP() []byte {
	file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDescData)
	})
	return file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDescData
}

var file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_goTypes = []interface{}{
	(*AwsEc2HostCatalog)(nil),    // 0: controller.api.resources.hosts.v1.AwsEc2HostCatalog
	(*wrappers.StringValue)(nil), // 1: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*wrappers.BoolValue)(nil),   // 3: google.protobuf.BoolValue
}
var file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_depIdxs = []int32{
	1, // 0: controller.api.resources.hosts.v1.AwsEc2HostCatalog.friendly_name:type_name -> google.protobuf.StringValue
	2, // 1: controller.api.resources.hosts.v1.AwsEc2HostCatalog.created_time:type_name -> google.protobuf.Timestamp
	2, // 2: controller.api.resources.hosts.v1.AwsEc2HostCatalog.updated_time:type_name -> google.protobuf.Timestamp
	3, // 3: controller.api.resources.hosts.v1.AwsEc2HostCatalog.disabled:type_name -> google.protobuf.BoolValue
	1, // 4: controller.api.resources.hosts.v1.AwsEc2HostCatalog.access_key:type_name -> google.protobuf.StringValue
	1, // 5: controller.api.resources.hosts.v1.AwsEc2HostCatalog.secret_key:type_name -> google.protobuf.StringValue
	3, // 6: controller.api.resources.hosts.v1.AwsEc2HostCatalog.rotate:type_name -> google.protobuf.BoolValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_init() }
func file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_init() {
	if File_controller_api_resources_hosts_v1_awsec2_host_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsEc2HostCatalog); i {
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
			RawDescriptor: file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_msgTypes,
	}.Build()
	File_controller_api_resources_hosts_v1_awsec2_host_catalog_proto = out.File
	file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_rawDesc = nil
	file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_goTypes = nil
	file_controller_api_resources_hosts_v1_awsec2_host_catalog_proto_depIdxs = nil
}