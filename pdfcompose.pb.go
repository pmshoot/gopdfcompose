// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: pdfcompose.proto

package pdfcompose

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

type FileCompose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *FileCompose) Reset() {
	*x = FileCompose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdfcompose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileCompose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileCompose) ProtoMessage() {}

func (x *FileCompose) ProtoReflect() protoreflect.Message {
	mi := &file_pdfcompose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileCompose.ProtoReflect.Descriptor instead.
func (*FileCompose) Descriptor() ([]byte, []int) {
	return file_pdfcompose_proto_rawDescGZIP(), []int{0}
}

func (x *FileCompose) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type ComposeFromFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*FileCompose `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ComposeFromFilesRequest) Reset() {
	*x = ComposeFromFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdfcompose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeFromFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeFromFilesRequest) ProtoMessage() {}

func (x *ComposeFromFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pdfcompose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeFromFilesRequest.ProtoReflect.Descriptor instead.
func (*ComposeFromFilesRequest) Descriptor() ([]byte, []int) {
	return file_pdfcompose_proto_rawDescGZIP(), []int{1}
}

func (x *ComposeFromFilesRequest) GetFiles() []*FileCompose {
	if x != nil {
		return x.Files
	}
	return nil
}

type ComposeFromFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outfile *FileCompose `protobuf:"bytes,1,opt,name=outfile,proto3,oneof" json:"outfile,omitempty"`
	Success bool         `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Error   *string      `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *ComposeFromFilesResponse) Reset() {
	*x = ComposeFromFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdfcompose_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeFromFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeFromFilesResponse) ProtoMessage() {}

func (x *ComposeFromFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pdfcompose_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeFromFilesResponse.ProtoReflect.Descriptor instead.
func (*ComposeFromFilesResponse) Descriptor() ([]byte, []int) {
	return file_pdfcompose_proto_rawDescGZIP(), []int{2}
}

func (x *ComposeFromFilesResponse) GetOutfile() *FileCompose {
	if x != nil {
		return x.Outfile
	}
	return nil
}

func (x *ComposeFromFilesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ComposeFromFilesResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_pdfcompose_proto protoreflect.FileDescriptor

var file_pdfcompose_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x21,
	0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x48, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x64,
	0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x18,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x64, 0x66, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x6c,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x6d, 0x0a, 0x0a, 0x50,
	0x64, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x2e,
	0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f,
	0x3b, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pdfcompose_proto_rawDescOnce sync.Once
	file_pdfcompose_proto_rawDescData = file_pdfcompose_proto_rawDesc
)

func file_pdfcompose_proto_rawDescGZIP() []byte {
	file_pdfcompose_proto_rawDescOnce.Do(func() {
		file_pdfcompose_proto_rawDescData = protoimpl.X.CompressGZIP(file_pdfcompose_proto_rawDescData)
	})
	return file_pdfcompose_proto_rawDescData
}

var file_pdfcompose_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pdfcompose_proto_goTypes = []interface{}{
	(*FileCompose)(nil),              // 0: pdfcompose.FileCompose
	(*ComposeFromFilesRequest)(nil),  // 1: pdfcompose.ComposeFromFilesRequest
	(*ComposeFromFilesResponse)(nil), // 2: pdfcompose.ComposeFromFilesResponse
}
var file_pdfcompose_proto_depIdxs = []int32{
	0, // 0: pdfcompose.ComposeFromFilesRequest.files:type_name -> pdfcompose.FileCompose
	0, // 1: pdfcompose.ComposeFromFilesResponse.outfile:type_name -> pdfcompose.FileCompose
	1, // 2: pdfcompose.PdfCompose.ComposeFromFiles:input_type -> pdfcompose.ComposeFromFilesRequest
	2, // 3: pdfcompose.PdfCompose.ComposeFromFiles:output_type -> pdfcompose.ComposeFromFilesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pdfcompose_proto_init() }
func file_pdfcompose_proto_init() {
	if File_pdfcompose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pdfcompose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileCompose); i {
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
		file_pdfcompose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeFromFilesRequest); i {
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
		file_pdfcompose_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeFromFilesResponse); i {
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
	file_pdfcompose_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pdfcompose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pdfcompose_proto_goTypes,
		DependencyIndexes: file_pdfcompose_proto_depIdxs,
		MessageInfos:      file_pdfcompose_proto_msgTypes,
	}.Build()
	File_pdfcompose_proto = out.File
	file_pdfcompose_proto_rawDesc = nil
	file_pdfcompose_proto_goTypes = nil
	file_pdfcompose_proto_depIdxs = nil
}
