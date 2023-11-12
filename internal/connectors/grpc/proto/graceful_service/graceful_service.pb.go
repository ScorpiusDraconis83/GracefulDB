// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: config/proto/graceful_service.proto

package graceful_service

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

// VQL
type VRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instruction string `protobuf:"bytes,1,opt,name=instruction,proto3" json:"instruction,omitempty"`
}

func (x *VRequest) Reset() {
	*x = VRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_graceful_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VRequest) ProtoMessage() {}

func (x *VRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_graceful_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VRequest.ProtoReflect.Descriptor instead.
func (*VRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_graceful_service_proto_rawDescGZIP(), []int{0}
}

func (x *VRequest) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

type VResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VResponse) Reset() {
	*x = VResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_graceful_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VResponse) ProtoMessage() {}

func (x *VResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_graceful_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VResponse.ProtoReflect.Descriptor instead.
func (*VResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_graceful_service_proto_rawDescGZIP(), []int{1}
}

func (x *VResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// SQL
type SRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instruction string   `protobuf:"bytes,1,opt,name=instruction,proto3" json:"instruction,omitempty"`
	Placeholder []string `protobuf:"bytes,2,rep,name=placeholder,proto3" json:"placeholder,omitempty"`
}

func (x *SRequest) Reset() {
	*x = SRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_graceful_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SRequest) ProtoMessage() {}

func (x *SRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_graceful_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SRequest.ProtoReflect.Descriptor instead.
func (*SRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_graceful_service_proto_rawDescGZIP(), []int{2}
}

func (x *SRequest) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

func (x *SRequest) GetPlaceholder() []string {
	if x != nil {
		return x.Placeholder
	}
	return nil
}

type SResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SResponse) Reset() {
	*x = SResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_graceful_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SResponse) ProtoMessage() {}

func (x *SResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_graceful_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SResponse.ProtoReflect.Descriptor instead.
func (*SResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_graceful_service_proto_rawDescGZIP(), []int{3}
}

func (x *SResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_config_proto_graceful_service_proto protoreflect.FileDescriptor

var file_config_proto_graceful_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x08, 0x56, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x09, 0x56, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x08,
	0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x09,
	0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x97, 0x01, 0x0a, 0x0f, 0x47, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x56, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a,
	0x31, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_graceful_service_proto_rawDescOnce sync.Once
	file_config_proto_graceful_service_proto_rawDescData = file_config_proto_graceful_service_proto_rawDesc
)

func file_config_proto_graceful_service_proto_rawDescGZIP() []byte {
	file_config_proto_graceful_service_proto_rawDescOnce.Do(func() {
		file_config_proto_graceful_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_graceful_service_proto_rawDescData)
	})
	return file_config_proto_graceful_service_proto_rawDescData
}

var file_config_proto_graceful_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_proto_graceful_service_proto_goTypes = []interface{}{
	(*VRequest)(nil),  // 0: graceful_service.VRequest
	(*VResponse)(nil), // 1: graceful_service.VResponse
	(*SRequest)(nil),  // 2: graceful_service.SRequest
	(*SResponse)(nil), // 3: graceful_service.SResponse
}
var file_config_proto_graceful_service_proto_depIdxs = []int32{
	0, // 0: graceful_service.GracefulService.VQuery:input_type -> graceful_service.VRequest
	2, // 1: graceful_service.GracefulService.SQuery:input_type -> graceful_service.SRequest
	1, // 2: graceful_service.GracefulService.VQuery:output_type -> graceful_service.VResponse
	3, // 3: graceful_service.GracefulService.SQuery:output_type -> graceful_service.SResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_proto_graceful_service_proto_init() }
func file_config_proto_graceful_service_proto_init() {
	if File_config_proto_graceful_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_graceful_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VRequest); i {
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
		file_config_proto_graceful_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VResponse); i {
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
		file_config_proto_graceful_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SRequest); i {
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
		file_config_proto_graceful_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SResponse); i {
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
			RawDescriptor: file_config_proto_graceful_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_graceful_service_proto_goTypes,
		DependencyIndexes: file_config_proto_graceful_service_proto_depIdxs,
		MessageInfos:      file_config_proto_graceful_service_proto_msgTypes,
	}.Build()
	File_config_proto_graceful_service_proto = out.File
	file_config_proto_graceful_service_proto_rawDesc = nil
	file_config_proto_graceful_service_proto_goTypes = nil
	file_config_proto_graceful_service_proto_depIdxs = nil
}