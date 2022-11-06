// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        (unknown)
// source: FingerprintProtocol.proto

package pb

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

type LogicalFingerprint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"` //  optional bytes identifier = 2;
}

func (x *LogicalFingerprint) Reset() {
	*x = LogicalFingerprint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FingerprintProtocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogicalFingerprint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicalFingerprint) ProtoMessage() {}

func (x *LogicalFingerprint) ProtoReflect() protoreflect.Message {
	mi := &file_FingerprintProtocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicalFingerprint.ProtoReflect.Descriptor instead.
func (*LogicalFingerprint) Descriptor() ([]byte, []int) {
	return file_FingerprintProtocol_proto_rawDescGZIP(), []int{0}
}

func (x *LogicalFingerprint) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type CombinedFingerprints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version           *uint32             `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	LocalFingerprint  *LogicalFingerprint `protobuf:"bytes,2,opt,name=localFingerprint" json:"localFingerprint,omitempty"`
	RemoteFingerprint *LogicalFingerprint `protobuf:"bytes,3,opt,name=remoteFingerprint" json:"remoteFingerprint,omitempty"`
}

func (x *CombinedFingerprints) Reset() {
	*x = CombinedFingerprints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FingerprintProtocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombinedFingerprints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombinedFingerprints) ProtoMessage() {}

func (x *CombinedFingerprints) ProtoReflect() protoreflect.Message {
	mi := &file_FingerprintProtocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombinedFingerprints.ProtoReflect.Descriptor instead.
func (*CombinedFingerprints) Descriptor() ([]byte, []int) {
	return file_FingerprintProtocol_proto_rawDescGZIP(), []int{1}
}

func (x *CombinedFingerprints) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *CombinedFingerprints) GetLocalFingerprint() *LogicalFingerprint {
	if x != nil {
		return x.LocalFingerprint
	}
	return nil
}

func (x *CombinedFingerprints) GetRemoteFingerprint() *LogicalFingerprint {
	if x != nil {
		return x.RemoteFingerprint
	}
	return nil
}

var File_FingerprintProtocol_proto protoreflect.FileDescriptor

var file_FingerprintProtocol_proto_rawDesc = []byte{
	0x0a, 0x19, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x25, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x65, 0x64, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x0f, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x12, 0x30, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x12, 0x31, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x43, 0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x68, 0x69,
	0x73, 0x70, 0x65, 0x72, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x6c, 0x69, 0x62, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x42, 0x11, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
}

var (
	file_FingerprintProtocol_proto_rawDescOnce sync.Once
	file_FingerprintProtocol_proto_rawDescData = file_FingerprintProtocol_proto_rawDesc
)

func file_FingerprintProtocol_proto_rawDescGZIP() []byte {
	file_FingerprintProtocol_proto_rawDescOnce.Do(func() {
		file_FingerprintProtocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_FingerprintProtocol_proto_rawDescData)
	})
	return file_FingerprintProtocol_proto_rawDescData
}

var file_FingerprintProtocol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FingerprintProtocol_proto_goTypes = []interface{}{
	(*LogicalFingerprint)(nil),   // 0: pb.LogicalFingerprint
	(*CombinedFingerprints)(nil), // 1: pb.CombinedFingerprints
}
var file_FingerprintProtocol_proto_depIdxs = []int32{
	0, // 0: pb.CombinedFingerprints.localFingerprint:type_name -> pb.LogicalFingerprint
	0, // 1: pb.CombinedFingerprints.remoteFingerprint:type_name -> pb.LogicalFingerprint
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FingerprintProtocol_proto_init() }
func file_FingerprintProtocol_proto_init() {
	if File_FingerprintProtocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FingerprintProtocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogicalFingerprint); i {
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
		file_FingerprintProtocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombinedFingerprints); i {
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
			RawDescriptor: file_FingerprintProtocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FingerprintProtocol_proto_goTypes,
		DependencyIndexes: file_FingerprintProtocol_proto_depIdxs,
		MessageInfos:      file_FingerprintProtocol_proto_msgTypes,
	}.Build()
	File_FingerprintProtocol_proto = out.File
	file_FingerprintProtocol_proto_rawDesc = nil
	file_FingerprintProtocol_proto_goTypes = nil
	file_FingerprintProtocol_proto_depIdxs = nil
}
