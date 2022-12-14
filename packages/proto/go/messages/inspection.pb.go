// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: messages/inspection.proto

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

type InspectionArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *InspectionArgs) Reset() {
	*x = InspectionArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_inspection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InspectionArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InspectionArgs) ProtoMessage() {}

func (x *InspectionArgs) ProtoReflect() protoreflect.Message {
	mi := &file_messages_inspection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InspectionArgs.ProtoReflect.Descriptor instead.
func (*InspectionArgs) Descriptor() ([]byte, []int) {
	return file_messages_inspection_proto_rawDescGZIP(), []int{0}
}

func (x *InspectionArgs) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type InspectionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url        string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Screenshot string   `protobuf:"bytes,2,opt,name=screenshot,proto3" json:"screenshot,omitempty"`
	Malicious  bool     `protobuf:"varint,3,opt,name=malicious,proto3" json:"malicious,omitempty"`
	Reasons    []string `protobuf:"bytes,4,rep,name=reasons,proto3" json:"reasons,omitempty"`
	Locations  []string `protobuf:"bytes,5,rep,name=locations,proto3" json:"locations,omitempty"`
	Hosts      []string `protobuf:"bytes,6,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Urls       []string `protobuf:"bytes,7,rep,name=urls,proto3" json:"urls,omitempty"`
	SendingTo  []string `protobuf:"bytes,8,rep,name=sending_to,json=sendingTo,proto3" json:"sending_to,omitempty"`
}

func (x *InspectionResult) Reset() {
	*x = InspectionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_inspection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InspectionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InspectionResult) ProtoMessage() {}

func (x *InspectionResult) ProtoReflect() protoreflect.Message {
	mi := &file_messages_inspection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InspectionResult.ProtoReflect.Descriptor instead.
func (*InspectionResult) Descriptor() ([]byte, []int) {
	return file_messages_inspection_proto_rawDescGZIP(), []int{1}
}

func (x *InspectionResult) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *InspectionResult) GetScreenshot() string {
	if x != nil {
		return x.Screenshot
	}
	return ""
}

func (x *InspectionResult) GetMalicious() bool {
	if x != nil {
		return x.Malicious
	}
	return false
}

func (x *InspectionResult) GetReasons() []string {
	if x != nil {
		return x.Reasons
	}
	return nil
}

func (x *InspectionResult) GetLocations() []string {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *InspectionResult) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *InspectionResult) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *InspectionResult) GetSendingTo() []string {
	if x != nil {
		return x.SendingTo
	}
	return nil
}

type HandlerArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string          `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Args      *InspectionArgs `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
}

func (x *HandlerArgs) Reset() {
	*x = HandlerArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_inspection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandlerArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandlerArgs) ProtoMessage() {}

func (x *HandlerArgs) ProtoReflect() protoreflect.Message {
	mi := &file_messages_inspection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandlerArgs.ProtoReflect.Descriptor instead.
func (*HandlerArgs) Descriptor() ([]byte, []int) {
	return file_messages_inspection_proto_rawDescGZIP(), []int{2}
}

func (x *HandlerArgs) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *HandlerArgs) GetArgs() *InspectionArgs {
	if x != nil {
		return x.Args
	}
	return nil
}

type HandlerResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string            `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Result    *InspectionResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Error     *string           `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *HandlerResult) Reset() {
	*x = HandlerResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_inspection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandlerResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandlerResult) ProtoMessage() {}

func (x *HandlerResult) ProtoReflect() protoreflect.Message {
	mi := &file_messages_inspection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandlerResult.ProtoReflect.Descriptor instead.
func (*HandlerResult) Descriptor() ([]byte, []int) {
	return file_messages_inspection_proto_rawDescGZIP(), []int{3}
}

func (x *HandlerResult) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *HandlerResult) GetResult() *InspectionResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *HandlerResult) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_messages_inspection_proto protoreflect.FileDescriptor

var file_messages_inspection_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xe3, 0x01, 0x0a, 0x10,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68,
	0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6c, 0x69, 0x63, 0x69, 0x6f, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x61, 0x6c, 0x69, 0x63, 0x69, 0x6f, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72,
	0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x22, 0x5c, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22,
	0x89, 0x01, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_inspection_proto_rawDescOnce sync.Once
	file_messages_inspection_proto_rawDescData = file_messages_inspection_proto_rawDesc
)

func file_messages_inspection_proto_rawDescGZIP() []byte {
	file_messages_inspection_proto_rawDescOnce.Do(func() {
		file_messages_inspection_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_inspection_proto_rawDescData)
	})
	return file_messages_inspection_proto_rawDescData
}

var file_messages_inspection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_inspection_proto_goTypes = []interface{}{
	(*InspectionArgs)(nil),   // 0: inspection.InspectionArgs
	(*InspectionResult)(nil), // 1: inspection.InspectionResult
	(*HandlerArgs)(nil),      // 2: inspection.HandlerArgs
	(*HandlerResult)(nil),    // 3: inspection.HandlerResult
}
var file_messages_inspection_proto_depIdxs = []int32{
	0, // 0: inspection.HandlerArgs.args:type_name -> inspection.InspectionArgs
	1, // 1: inspection.HandlerResult.result:type_name -> inspection.InspectionResult
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_messages_inspection_proto_init() }
func file_messages_inspection_proto_init() {
	if File_messages_inspection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_inspection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InspectionArgs); i {
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
		file_messages_inspection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InspectionResult); i {
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
		file_messages_inspection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandlerArgs); i {
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
		file_messages_inspection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandlerResult); i {
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
	file_messages_inspection_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_inspection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_inspection_proto_goTypes,
		DependencyIndexes: file_messages_inspection_proto_depIdxs,
		MessageInfos:      file_messages_inspection_proto_msgTypes,
	}.Build()
	File_messages_inspection_proto = out.File
	file_messages_inspection_proto_rawDesc = nil
	file_messages_inspection_proto_goTypes = nil
	file_messages_inspection_proto_depIdxs = nil
}
