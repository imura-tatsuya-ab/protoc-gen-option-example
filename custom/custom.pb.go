// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: custom.proto

package custom

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type MessageCustomOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target bool `protobuf:"varint,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *MessageCustomOptions) Reset() {
	*x = MessageCustomOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCustomOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCustomOptions) ProtoMessage() {}

func (x *MessageCustomOptions) ProtoReflect() protoreflect.Message {
	mi := &file_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCustomOptions.ProtoReflect.Descriptor instead.
func (*MessageCustomOptions) Descriptor() ([]byte, []int) {
	return file_custom_proto_rawDescGZIP(), []int{0}
}

func (x *MessageCustomOptions) GetTarget() bool {
	if x != nil {
		return x.Target
	}
	return false
}

var file_custom_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*MessageCustomOptions)(nil),
		Field:         50000,
		Name:          "custom.custom",
		Tag:           "bytes,50000,opt,name=custom",
		Filename:      "custom.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional custom.MessageCustomOptions custom = 50000;
	E_Custom = &file_custom_proto_extTypes[0]
)

var File_custom_proto protoreflect.FileDescriptor

var file_custom_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x3a, 0x57, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6d, 0x75, 0x72, 0x61, 0x2d, 0x74, 0x61, 0x74, 0x73, 0x75, 0x79, 0x61, 0x2d, 0x61, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_custom_proto_rawDescOnce sync.Once
	file_custom_proto_rawDescData = file_custom_proto_rawDesc
)

func file_custom_proto_rawDescGZIP() []byte {
	file_custom_proto_rawDescOnce.Do(func() {
		file_custom_proto_rawDescData = protoimpl.X.CompressGZIP(file_custom_proto_rawDescData)
	})
	return file_custom_proto_rawDescData
}

var file_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_custom_proto_goTypes = []interface{}{
	(*MessageCustomOptions)(nil),      // 0: custom.MessageCustomOptions
	(*descriptor.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_custom_proto_depIdxs = []int32{
	1, // 0: custom.custom:extendee -> google.protobuf.MessageOptions
	0, // 1: custom.custom:type_name -> custom.MessageCustomOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_custom_proto_init() }
func file_custom_proto_init() {
	if File_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCustomOptions); i {
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
			RawDescriptor: file_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_custom_proto_goTypes,
		DependencyIndexes: file_custom_proto_depIdxs,
		MessageInfos:      file_custom_proto_msgTypes,
		ExtensionInfos:    file_custom_proto_extTypes,
	}.Build()
	File_custom_proto = out.File
	file_custom_proto_rawDesc = nil
	file_custom_proto_goTypes = nil
	file_custom_proto_depIdxs = nil
}
