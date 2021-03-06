// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: pkg/cat/cat.proto

package cat

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

type CatInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CatInput) Reset() {
	*x = CatInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cat_cat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatInput) ProtoMessage() {}

func (x *CatInput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cat_cat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatInput.ProtoReflect.Descriptor instead.
func (*CatInput) Descriptor() ([]byte, []int) {
	return file_pkg_cat_cat_proto_rawDescGZIP(), []int{0}
}

func (x *CatInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Cat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FavoriteTreat string `protobuf:"bytes,2,opt,name=favoriteTreat,proto3" json:"favoriteTreat,omitempty"`
	Coat          string `protobuf:"bytes,3,opt,name=coat,proto3" json:"coat,omitempty"`
}

func (x *Cat) Reset() {
	*x = Cat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cat_cat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cat) ProtoMessage() {}

func (x *Cat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cat_cat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cat.ProtoReflect.Descriptor instead.
func (*Cat) Descriptor() ([]byte, []int) {
	return file_pkg_cat_cat_proto_rawDescGZIP(), []int{1}
}

func (x *Cat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cat) GetFavoriteTreat() string {
	if x != nil {
		return x.FavoriteTreat
	}
	return ""
}

func (x *Cat) GetCoat() string {
	if x != nil {
		return x.Coat
	}
	return ""
}

var File_pkg_cat_cat_proto protoreflect.FileDescriptor

var file_pkg_cat_cat_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x74, 0x2f, 0x63, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x7a, 0x72, 0x64, 0x61, 0x6c, 0x65, 0x79, 0x2e, 0x63, 0x61, 0x74,
	0x22, 0x1e, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x53, 0x0a, 0x03, 0x43, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x54, 0x72, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x54, 0x72, 0x65, 0x61,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x61, 0x74, 0x32, 0x42, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x4d, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x74,
	0x12, 0x15, 0x2e, 0x7a, 0x72, 0x64, 0x61, 0x6c, 0x65, 0x79, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x43,
	0x61, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x10, 0x2e, 0x7a, 0x72, 0x64, 0x61, 0x6c, 0x65,
	0x79, 0x2e, 0x63, 0x61, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x63, 0x61,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_cat_cat_proto_rawDescOnce sync.Once
	file_pkg_cat_cat_proto_rawDescData = file_pkg_cat_cat_proto_rawDesc
)

func file_pkg_cat_cat_proto_rawDescGZIP() []byte {
	file_pkg_cat_cat_proto_rawDescOnce.Do(func() {
		file_pkg_cat_cat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_cat_cat_proto_rawDescData)
	})
	return file_pkg_cat_cat_proto_rawDescData
}

var file_pkg_cat_cat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_cat_cat_proto_goTypes = []interface{}{
	(*CatInput)(nil), // 0: zrdaley.cat.CatInput
	(*Cat)(nil),      // 1: zrdaley.cat.Cat
}
var file_pkg_cat_cat_proto_depIdxs = []int32{
	0, // 0: zrdaley.cat.CatGenerator.MakeCat:input_type -> zrdaley.cat.CatInput
	1, // 1: zrdaley.cat.CatGenerator.MakeCat:output_type -> zrdaley.cat.Cat
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_cat_cat_proto_init() }
func file_pkg_cat_cat_proto_init() {
	if File_pkg_cat_cat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_cat_cat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatInput); i {
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
		file_pkg_cat_cat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cat); i {
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
			RawDescriptor: file_pkg_cat_cat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_cat_cat_proto_goTypes,
		DependencyIndexes: file_pkg_cat_cat_proto_depIdxs,
		MessageInfos:      file_pkg_cat_cat_proto_msgTypes,
	}.Build()
	File_pkg_cat_cat_proto = out.File
	file_pkg_cat_cat_proto_rawDesc = nil
	file_pkg_cat_cat_proto_goTypes = nil
	file_pkg_cat_cat_proto_depIdxs = nil
}
