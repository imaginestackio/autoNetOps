// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: networkstuff.proto

package networkstuff

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

type RouterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RouterRequest) Reset() {
	*x = RouterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkstuff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterRequest) ProtoMessage() {}

func (x *RouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_networkstuff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterRequest.ProtoReflect.Descriptor instead.
func (*RouterRequest) Descriptor() ([]byte, []int) {
	return file_networkstuff_proto_rawDescGZIP(), []int{0}
}

func (x *RouterRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Router struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hostname   string       `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Interfaces []*Interface `protobuf:"bytes,3,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
}

func (x *Router) Reset() {
	*x = Router{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkstuff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Router) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Router) ProtoMessage() {}

func (x *Router) ProtoReflect() protoreflect.Message {
	mi := &file_networkstuff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Router.ProtoReflect.Descriptor instead.
func (*Router) Descriptor() ([]byte, []int) {
	return file_networkstuff_proto_rawDescGZIP(), []int{1}
}

func (x *Router) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Router) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Router) GetInterfaces() []*Interface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

type Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Interface) Reset() {
	*x = Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkstuff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface) ProtoMessage() {}

func (x *Interface) ProtoReflect() protoreflect.Message {
	mi := &file_networkstuff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface.ProtoReflect.Descriptor instead.
func (*Interface) Descriptor() ([]byte, []int) {
	return file_networkstuff_proto_rawDescGZIP(), []int{2}
}

func (x *Interface) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Interface) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_networkstuff_proto protoreflect.FileDescriptor

var file_networkstuff_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x74, 0x75,
	0x66, 0x66, 0x22, 0x1f, 0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x22, 0x3d, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0x4f, 0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12,
	0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x74, 0x75, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_networkstuff_proto_rawDescOnce sync.Once
	file_networkstuff_proto_rawDescData = file_networkstuff_proto_rawDesc
)

func file_networkstuff_proto_rawDescGZIP() []byte {
	file_networkstuff_proto_rawDescOnce.Do(func() {
		file_networkstuff_proto_rawDescData = protoimpl.X.CompressGZIP(file_networkstuff_proto_rawDescData)
	})
	return file_networkstuff_proto_rawDescData
}

var file_networkstuff_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_networkstuff_proto_goTypes = []interface{}{
	(*RouterRequest)(nil), // 0: networkstuff.RouterRequest
	(*Router)(nil),        // 1: networkstuff.Router
	(*Interface)(nil),     // 2: networkstuff.Interface
}
var file_networkstuff_proto_depIdxs = []int32{
	2, // 0: networkstuff.Router.interfaces:type_name -> networkstuff.Interface
	0, // 1: networkstuff.RouterService.GetRouter:input_type -> networkstuff.RouterRequest
	1, // 2: networkstuff.RouterService.GetRouter:output_type -> networkstuff.Router
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_networkstuff_proto_init() }
func file_networkstuff_proto_init() {
	if File_networkstuff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_networkstuff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterRequest); i {
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
		file_networkstuff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Router); i {
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
		file_networkstuff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface); i {
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
			RawDescriptor: file_networkstuff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_networkstuff_proto_goTypes,
		DependencyIndexes: file_networkstuff_proto_depIdxs,
		MessageInfos:      file_networkstuff_proto_msgTypes,
	}.Build()
	File_networkstuff_proto = out.File
	file_networkstuff_proto_rawDesc = nil
	file_networkstuff_proto_goTypes = nil
	file_networkstuff_proto_depIdxs = nil
}
