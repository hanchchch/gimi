// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: messages/runtimefilter.proto

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

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Os  string `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_runtimefilter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_runtimefilter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_messages_runtimefilter_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StartRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_runtimefilter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_runtimefilter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_messages_runtimefilter_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetResultRequest) Reset() {
	*x = GetResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_runtimefilter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResultRequest) ProtoMessage() {}

func (x *GetResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_runtimefilter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResultRequest.ProtoReflect.Descriptor instead.
func (*GetResultRequest) Descriptor() ([]byte, []int) {
	return file_messages_runtimefilter_proto_rawDescGZIP(), []int{2}
}

func (x *GetResultRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Stdout string `protobuf:"bytes,3,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr string `protobuf:"bytes,4,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (x *GetResultResponse) Reset() {
	*x = GetResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_runtimefilter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResultResponse) ProtoMessage() {}

func (x *GetResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_runtimefilter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResultResponse.ProtoReflect.Descriptor instead.
func (*GetResultResponse) Descriptor() ([]byte, []int) {
	return file_messages_runtimefilter_proto_rawDescGZIP(), []int{3}
}

func (x *GetResultResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetResultResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetResultResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *GetResultResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

type SubResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timeout *int32 `protobuf:"varint,2,opt,name=timeout,proto3,oneof" json:"timeout,omitempty"`
}

func (x *SubResultRequest) Reset() {
	*x = SubResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_runtimefilter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubResultRequest) ProtoMessage() {}

func (x *SubResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_runtimefilter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubResultRequest.ProtoReflect.Descriptor instead.
func (*SubResultRequest) Descriptor() ([]byte, []int) {
	return file_messages_runtimefilter_proto_rawDescGZIP(), []int{4}
}

func (x *SubResultRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubResultRequest) GetTimeout() int32 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

var File_messages_runtimefilter_proto protoreflect.FileDescriptor

var file_messages_runtimefilter_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x30, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x22,
	0x1f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64,
	0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x22, 0x4d, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x32, 0x80, 0x02, 0x0a, 0x14, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_runtimefilter_proto_rawDescOnce sync.Once
	file_messages_runtimefilter_proto_rawDescData = file_messages_runtimefilter_proto_rawDesc
)

func file_messages_runtimefilter_proto_rawDescGZIP() []byte {
	file_messages_runtimefilter_proto_rawDescOnce.Do(func() {
		file_messages_runtimefilter_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_runtimefilter_proto_rawDescData)
	})
	return file_messages_runtimefilter_proto_rawDescData
}

var file_messages_runtimefilter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_messages_runtimefilter_proto_goTypes = []interface{}{
	(*StartRequest)(nil),      // 0: runtimefilter.StartRequest
	(*StartResponse)(nil),     // 1: runtimefilter.StartResponse
	(*GetResultRequest)(nil),  // 2: runtimefilter.GetResultRequest
	(*GetResultResponse)(nil), // 3: runtimefilter.GetResultResponse
	(*SubResultRequest)(nil),  // 4: runtimefilter.SubResultRequest
}
var file_messages_runtimefilter_proto_depIdxs = []int32{
	0, // 0: runtimefilter.RuntimeFilterService.Start:input_type -> runtimefilter.StartRequest
	2, // 1: runtimefilter.RuntimeFilterService.GetResult:input_type -> runtimefilter.GetResultRequest
	4, // 2: runtimefilter.RuntimeFilterService.SubResult:input_type -> runtimefilter.SubResultRequest
	1, // 3: runtimefilter.RuntimeFilterService.Start:output_type -> runtimefilter.StartResponse
	3, // 4: runtimefilter.RuntimeFilterService.GetResult:output_type -> runtimefilter.GetResultResponse
	3, // 5: runtimefilter.RuntimeFilterService.SubResult:output_type -> runtimefilter.GetResultResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_runtimefilter_proto_init() }
func file_messages_runtimefilter_proto_init() {
	if File_messages_runtimefilter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_runtimefilter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_messages_runtimefilter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_messages_runtimefilter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResultRequest); i {
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
		file_messages_runtimefilter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResultResponse); i {
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
		file_messages_runtimefilter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubResultRequest); i {
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
	file_messages_runtimefilter_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_runtimefilter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messages_runtimefilter_proto_goTypes,
		DependencyIndexes: file_messages_runtimefilter_proto_depIdxs,
		MessageInfos:      file_messages_runtimefilter_proto_msgTypes,
	}.Build()
	File_messages_runtimefilter_proto = out.File
	file_messages_runtimefilter_proto_rawDesc = nil
	file_messages_runtimefilter_proto_goTypes = nil
	file_messages_runtimefilter_proto_depIdxs = nil
}
