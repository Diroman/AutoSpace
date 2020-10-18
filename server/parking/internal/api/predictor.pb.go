// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: predictor.proto

package predictor

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predictor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type Boxes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Boxes) Reset() {
	*x = Boxes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predictor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Boxes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Boxes) ProtoMessage() {}

func (x *Boxes) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Boxes.ProtoReflect.Descriptor instead.
func (*Boxes) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{1}
}

func (x *Boxes) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Boxes) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Area  int64    `protobuf:"varint,1,opt,name=area,proto3" json:"area,omitempty"`
	Boxes []*Boxes `protobuf:"bytes,2,rep,name=boxes,proto3" json:"boxes,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predictor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{2}
}

func (x *Row) GetArea() int64 {
	if x != nil {
		return x.Area
	}
	return 0
}

func (x *Row) GetBoxes() []*Boxes {
	if x != nil {
		return x.Boxes
	}
	return nil
}

type Rows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Row `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Rows) Reset() {
	*x = Rows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predictor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rows) ProtoMessage() {}

func (x *Rows) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rows.ProtoReflect.Descriptor instead.
func (*Rows) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{3}
}

func (x *Rows) GetData() []*Row {
	if x != nil {
		return x.Data
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classes map[string]*Rows `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predictor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetClasses() map[string]*Rows {
	if x != nil {
		return x.Classes
	}
	return nil
}

var File_predictor_proto protoreflect.FileDescriptor

var file_predictor_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1d, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x23, 0x0a, 0x05, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x37, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61,
	0x12, 0x1c, 0x0a, 0x05, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x05, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x22, 0x20,
	0x0a, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x7b, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x1a, 0x41, 0x0a, 0x0c, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x6f,
	0x77, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x29, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x07,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x06, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a,
	0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_predictor_proto_rawDescOnce sync.Once
	file_predictor_proto_rawDescData = file_predictor_proto_rawDesc
)

func file_predictor_proto_rawDescGZIP() []byte {
	file_predictor_proto_rawDescOnce.Do(func() {
		file_predictor_proto_rawDescData = protoimpl.X.CompressGZIP(file_predictor_proto_rawDescData)
	})
	return file_predictor_proto_rawDescData
}

var file_predictor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_predictor_proto_goTypes = []interface{}{
	(*Image)(nil),  // 0: Image
	(*Boxes)(nil),  // 1: Boxes
	(*Row)(nil),    // 2: Row
	(*Rows)(nil),   // 3: Rows
	(*Result)(nil), // 4: Result
	nil,            // 5: Result.ClassesEntry
}
var file_predictor_proto_depIdxs = []int32{
	1, // 0: Row.boxes:type_name -> Boxes
	2, // 1: Rows.data:type_name -> Row
	5, // 2: Result.classes:type_name -> Result.ClassesEntry
	3, // 3: Result.ClassesEntry.value:type_name -> Rows
	0, // 4: CarDetector.predict:input_type -> Image
	4, // 5: CarDetector.predict:output_type -> Result
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_predictor_proto_init() }
func file_predictor_proto_init() {
	if File_predictor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_predictor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_predictor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Boxes); i {
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
		file_predictor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_predictor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rows); i {
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
		file_predictor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_predictor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_predictor_proto_goTypes,
		DependencyIndexes: file_predictor_proto_depIdxs,
		MessageInfos:      file_predictor_proto_msgTypes,
	}.Build()
	File_predictor_proto = out.File
	file_predictor_proto_rawDesc = nil
	file_predictor_proto_goTypes = nil
	file_predictor_proto_depIdxs = nil
}