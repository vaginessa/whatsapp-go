// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: NoiseHandshake.proto

package pb

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

type PatternStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S1 []byte `protobuf:"bytes,1,opt,name=S1" json:"S1,omitempty"`
	S2 []byte `protobuf:"bytes,2,opt,name=S2" json:"S2,omitempty"`
	S3 []byte `protobuf:"bytes,3,opt,name=S3" json:"S3,omitempty"`
}

func (x *PatternStep) Reset() {
	*x = PatternStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NoiseHandshake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatternStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatternStep) ProtoMessage() {}

func (x *PatternStep) ProtoReflect() protoreflect.Message {
	mi := &file_NoiseHandshake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatternStep.ProtoReflect.Descriptor instead.
func (*PatternStep) Descriptor() ([]byte, []int) {
	return file_NoiseHandshake_proto_rawDescGZIP(), []int{0}
}

func (x *PatternStep) GetS1() []byte {
	if x != nil {
		return x.S1
	}
	return nil
}

func (x *PatternStep) GetS2() []byte {
	if x != nil {
		return x.S2
	}
	return nil
}

func (x *PatternStep) GetS3() []byte {
	if x != nil {
		return x.S3
	}
	return nil
}

type PatternXX struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	E         *PatternStep `protobuf:"bytes,2,opt,name=E" json:"E,omitempty"`                             // e
	E_EE_S_ES *PatternStep `protobuf:"bytes,3,opt,name=E_EE_S_ES,json=EEESES" json:"E_EE_S_ES,omitempty"` // e, ee, s, es
	S_SE      *PatternStep `protobuf:"bytes,4,opt,name=S_SE,json=SSE" json:"S_SE,omitempty"`              // s, se
}

func (x *PatternXX) Reset() {
	*x = PatternXX{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NoiseHandshake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatternXX) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatternXX) ProtoMessage() {}

func (x *PatternXX) ProtoReflect() protoreflect.Message {
	mi := &file_NoiseHandshake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatternXX.ProtoReflect.Descriptor instead.
func (*PatternXX) Descriptor() ([]byte, []int) {
	return file_NoiseHandshake_proto_rawDescGZIP(), []int{1}
}

func (x *PatternXX) GetE() *PatternStep {
	if x != nil {
		return x.E
	}
	return nil
}

func (x *PatternXX) GetE_EE_S_ES() *PatternStep {
	if x != nil {
		return x.E_EE_S_ES
	}
	return nil
}

func (x *PatternXX) GetS_SE() *PatternStep {
	if x != nil {
		return x.S_SE
	}
	return nil
}

type PatternIK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	E_EE_S_ES *PatternStep `protobuf:"bytes,2,opt,name=E_EE_S_ES,json=EEESES" json:"E_EE_S_ES,omitempty"` // e, es, s, ss ->
	E_EE_SE   *PatternStep `protobuf:"bytes,3,opt,name=E_EE_SE,json=EEESE" json:"E_EE_SE,omitempty"`      // <- e, ee, se
}

func (x *PatternIK) Reset() {
	*x = PatternIK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NoiseHandshake_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatternIK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatternIK) ProtoMessage() {}

func (x *PatternIK) ProtoReflect() protoreflect.Message {
	mi := &file_NoiseHandshake_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatternIK.ProtoReflect.Descriptor instead.
func (*PatternIK) Descriptor() ([]byte, []int) {
	return file_NoiseHandshake_proto_rawDescGZIP(), []int{2}
}

func (x *PatternIK) GetE_EE_S_ES() *PatternStep {
	if x != nil {
		return x.E_EE_S_ES
	}
	return nil
}

func (x *PatternIK) GetE_EE_SE() *PatternStep {
	if x != nil {
		return x.E_EE_SE
	}
	return nil
}

var File_NoiseHandshake_proto protoreflect.FileDescriptor

var file_NoiseHandshake_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x53, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x53, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x53, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x53, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x53, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x53, 0x33, 0x22, 0x71, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x58, 0x58, 0x12, 0x1a, 0x0a, 0x01, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x01, 0x45, 0x12, 0x27,
	0x0a, 0x09, 0x45, 0x5f, 0x45, 0x45, 0x5f, 0x53, 0x5f, 0x45, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x52,
	0x06, 0x45, 0x45, 0x45, 0x53, 0x45, 0x53, 0x12, 0x1f, 0x0a, 0x04, 0x53, 0x5f, 0x53, 0x45, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x03, 0x53, 0x53, 0x45, 0x22, 0x5a, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x49, 0x4b, 0x12, 0x27, 0x0a, 0x09, 0x45, 0x5f, 0x45, 0x45, 0x5f, 0x53, 0x5f,
	0x45, 0x53, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x06, 0x45, 0x45, 0x45, 0x53, 0x45, 0x53, 0x12, 0x24,
	0x0a, 0x07, 0x45, 0x5f, 0x45, 0x45, 0x5f, 0x53, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x45,
	0x45, 0x45, 0x53, 0x45, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
}

var (
	file_NoiseHandshake_proto_rawDescOnce sync.Once
	file_NoiseHandshake_proto_rawDescData = file_NoiseHandshake_proto_rawDesc
)

func file_NoiseHandshake_proto_rawDescGZIP() []byte {
	file_NoiseHandshake_proto_rawDescOnce.Do(func() {
		file_NoiseHandshake_proto_rawDescData = protoimpl.X.CompressGZIP(file_NoiseHandshake_proto_rawDescData)
	})
	return file_NoiseHandshake_proto_rawDescData
}

var file_NoiseHandshake_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_NoiseHandshake_proto_goTypes = []interface{}{
	(*PatternStep)(nil), // 0: PatternStep
	(*PatternXX)(nil),   // 1: PatternXX
	(*PatternIK)(nil),   // 2: PatternIK
}
var file_NoiseHandshake_proto_depIdxs = []int32{
	0, // 0: PatternXX.E:type_name -> PatternStep
	0, // 1: PatternXX.E_EE_S_ES:type_name -> PatternStep
	0, // 2: PatternXX.S_SE:type_name -> PatternStep
	0, // 3: PatternIK.E_EE_S_ES:type_name -> PatternStep
	0, // 4: PatternIK.E_EE_SE:type_name -> PatternStep
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_NoiseHandshake_proto_init() }
func file_NoiseHandshake_proto_init() {
	if File_NoiseHandshake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NoiseHandshake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatternStep); i {
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
		file_NoiseHandshake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatternXX); i {
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
		file_NoiseHandshake_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatternIK); i {
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
			RawDescriptor: file_NoiseHandshake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NoiseHandshake_proto_goTypes,
		DependencyIndexes: file_NoiseHandshake_proto_depIdxs,
		MessageInfos:      file_NoiseHandshake_proto_msgTypes,
	}.Build()
	File_NoiseHandshake_proto = out.File
	file_NoiseHandshake_proto_rawDesc = nil
	file_NoiseHandshake_proto_goTypes = nil
	file_NoiseHandshake_proto_depIdxs = nil
}
