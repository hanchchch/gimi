// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: messages/dbfilter.proto

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

type Blacklist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url        string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	DetectedAt string `protobuf:"bytes,2,opt,name=detectedAt,proto3" json:"detectedAt,omitempty"`
}

func (x *Blacklist) Reset() {
	*x = Blacklist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_dbfilter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blacklist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blacklist) ProtoMessage() {}

func (x *Blacklist) ProtoReflect() protoreflect.Message {
	mi := &file_messages_dbfilter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blacklist.ProtoReflect.Descriptor instead.
func (*Blacklist) Descriptor() ([]byte, []int) {
	return file_messages_dbfilter_proto_rawDescGZIP(), []int{0}
}

func (x *Blacklist) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Blacklist) GetDetectedAt() string {
	if x != nil {
		return x.DetectedAt
	}
	return ""
}

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_dbfilter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_dbfilter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_messages_dbfilter_proto_rawDescGZIP(), []int{1}
}

func (x *FindRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type FindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found     bool       `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Blacklist *Blacklist `protobuf:"bytes,2,opt,name=blacklist,proto3,oneof" json:"blacklist,omitempty"`
}

func (x *FindResponse) Reset() {
	*x = FindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_dbfilter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindResponse) ProtoMessage() {}

func (x *FindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_dbfilter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindResponse.ProtoReflect.Descriptor instead.
func (*FindResponse) Descriptor() ([]byte, []int) {
	return file_messages_dbfilter_proto_rawDescGZIP(), []int{2}
}

func (x *FindResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *FindResponse) GetBlacklist() *Blacklist {
	if x != nil {
		return x.Blacklist
	}
	return nil
}

var File_messages_dbfilter_proto protoreflect.FileDescriptor

var file_messages_dbfilter_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x64, 0x62, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x62, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x09, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x1f, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x6a, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x62, 0x6c, 0x61,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64,
	0x62, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x32,
	0x4a, 0x0a, 0x0f, 0x44, 0x62, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x64, 0x62, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x64, 0x62, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_dbfilter_proto_rawDescOnce sync.Once
	file_messages_dbfilter_proto_rawDescData = file_messages_dbfilter_proto_rawDesc
)

func file_messages_dbfilter_proto_rawDescGZIP() []byte {
	file_messages_dbfilter_proto_rawDescOnce.Do(func() {
		file_messages_dbfilter_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_dbfilter_proto_rawDescData)
	})
	return file_messages_dbfilter_proto_rawDescData
}

var file_messages_dbfilter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messages_dbfilter_proto_goTypes = []interface{}{
	(*Blacklist)(nil),    // 0: dbfilter.Blacklist
	(*FindRequest)(nil),  // 1: dbfilter.FindRequest
	(*FindResponse)(nil), // 2: dbfilter.FindResponse
}
var file_messages_dbfilter_proto_depIdxs = []int32{
	0, // 0: dbfilter.FindResponse.blacklist:type_name -> dbfilter.Blacklist
	1, // 1: dbfilter.DbFilterService.Find:input_type -> dbfilter.FindRequest
	2, // 2: dbfilter.DbFilterService.Find:output_type -> dbfilter.FindResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_dbfilter_proto_init() }
func file_messages_dbfilter_proto_init() {
	if File_messages_dbfilter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_dbfilter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blacklist); i {
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
		file_messages_dbfilter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
		file_messages_dbfilter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindResponse); i {
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
	file_messages_dbfilter_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_dbfilter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messages_dbfilter_proto_goTypes,
		DependencyIndexes: file_messages_dbfilter_proto_depIdxs,
		MessageInfos:      file_messages_dbfilter_proto_msgTypes,
	}.Build()
	File_messages_dbfilter_proto = out.File
	file_messages_dbfilter_proto_rawDesc = nil
	file_messages_dbfilter_proto_goTypes = nil
	file_messages_dbfilter_proto_depIdxs = nil
}
