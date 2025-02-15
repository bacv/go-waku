// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: waku_swap.proto

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

type Cheque struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssuerAddress string `protobuf:"bytes,1,opt,name=issuerAddress,proto3" json:"issuerAddress,omitempty"`
	Beneficiary   []byte `protobuf:"bytes,2,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Date          uint32 `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	Amount        uint32 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Signature     []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Cheque) Reset() {
	*x = Cheque{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waku_swap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cheque) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cheque) ProtoMessage() {}

func (x *Cheque) ProtoReflect() protoreflect.Message {
	mi := &file_waku_swap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cheque.ProtoReflect.Descriptor instead.
func (*Cheque) Descriptor() ([]byte, []int) {
	return file_waku_swap_proto_rawDescGZIP(), []int{0}
}

func (x *Cheque) GetIssuerAddress() string {
	if x != nil {
		return x.IssuerAddress
	}
	return ""
}

func (x *Cheque) GetBeneficiary() []byte {
	if x != nil {
		return x.Beneficiary
	}
	return nil
}

func (x *Cheque) GetDate() uint32 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Cheque) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Cheque) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beneficiary []byte `protobuf:"bytes,1,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waku_swap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_waku_swap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_waku_swap_proto_rawDescGZIP(), []int{1}
}

func (x *Handshake) GetBeneficiary() []byte {
	if x != nil {
		return x.Beneficiary
	}
	return nil
}

var File_waku_swap_proto protoreflect.FileDescriptor

var file_waku_swap_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x61, 0x6b, 0x75, 0x5f, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x65, 0x71, 0x75, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x63, 0x69, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x62, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x2d, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waku_swap_proto_rawDescOnce sync.Once
	file_waku_swap_proto_rawDescData = file_waku_swap_proto_rawDesc
)

func file_waku_swap_proto_rawDescGZIP() []byte {
	file_waku_swap_proto_rawDescOnce.Do(func() {
		file_waku_swap_proto_rawDescData = protoimpl.X.CompressGZIP(file_waku_swap_proto_rawDescData)
	})
	return file_waku_swap_proto_rawDescData
}

var file_waku_swap_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_waku_swap_proto_goTypes = []interface{}{
	(*Cheque)(nil),    // 0: pb.Cheque
	(*Handshake)(nil), // 1: pb.Handshake
}
var file_waku_swap_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_waku_swap_proto_init() }
func file_waku_swap_proto_init() {
	if File_waku_swap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waku_swap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cheque); i {
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
		file_waku_swap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake); i {
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
			RawDescriptor: file_waku_swap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waku_swap_proto_goTypes,
		DependencyIndexes: file_waku_swap_proto_depIdxs,
		MessageInfos:      file_waku_swap_proto_msgTypes,
	}.Build()
	File_waku_swap_proto = out.File
	file_waku_swap_proto_rawDesc = nil
	file_waku_swap_proto_goTypes = nil
	file_waku_swap_proto_depIdxs = nil
}
