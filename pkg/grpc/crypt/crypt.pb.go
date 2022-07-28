// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: crypt.proto

package crypt

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

// A SealedMessage is an encrypted protobuf message.
type SealedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Curve25519 public key used to encrypt the data encryption key.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// The XChacha20poly1305 key used to encrypt the data,
	// itself stored encrypted by the Curve25519 public key.
	DataEncryptionKey []byte `protobuf:"bytes,2,opt,name=data_encryption_key,json=dataEncryptionKey,proto3" json:"data_encryption_key,omitempty"`
	// The message type indicates the type of the protobuf message stored encrypted in encrypted_message.
	MessageType string `protobuf:"bytes,3,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	// An arbitrary encrypted protobuf message (marshaled as protojson before encryption).
	EncryptedMessage []byte `protobuf:"bytes,4,opt,name=encrypted_message,json=encryptedMessage,proto3" json:"encrypted_message,omitempty"`
}

func (x *SealedMessage) Reset() {
	*x = SealedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SealedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SealedMessage) ProtoMessage() {}

func (x *SealedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crypt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SealedMessage.ProtoReflect.Descriptor instead.
func (*SealedMessage) Descriptor() ([]byte, []int) {
	return file_crypt_proto_rawDescGZIP(), []int{0}
}

func (x *SealedMessage) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *SealedMessage) GetDataEncryptionKey() []byte {
	if x != nil {
		return x.DataEncryptionKey
	}
	return nil
}

func (x *SealedMessage) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *SealedMessage) GetEncryptedMessage() []byte {
	if x != nil {
		return x.EncryptedMessage
	}
	return nil
}

type PublicKeyEncryptionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PublicKeyEncryptionKey) Reset() {
	*x = PublicKeyEncryptionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKeyEncryptionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKeyEncryptionKey) ProtoMessage() {}

func (x *PublicKeyEncryptionKey) ProtoReflect() protoreflect.Message {
	mi := &file_crypt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKeyEncryptionKey.ProtoReflect.Descriptor instead.
func (*PublicKeyEncryptionKey) Descriptor() ([]byte, []int) {
	return file_crypt_proto_rawDescGZIP(), []int{1}
}

func (x *PublicKeyEncryptionKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicKeyEncryptionKey) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_crypt_proto protoreflect.FileDescriptor

var file_crypt_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70,
	0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x22, 0xa6, 0x01,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x16, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65,
	0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypt_proto_rawDescOnce sync.Once
	file_crypt_proto_rawDescData = file_crypt_proto_rawDesc
)

func file_crypt_proto_rawDescGZIP() []byte {
	file_crypt_proto_rawDescOnce.Do(func() {
		file_crypt_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypt_proto_rawDescData)
	})
	return file_crypt_proto_rawDescData
}

var file_crypt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crypt_proto_goTypes = []interface{}{
	(*SealedMessage)(nil),          // 0: pomerium.crypt.SealedMessage
	(*PublicKeyEncryptionKey)(nil), // 1: pomerium.crypt.PublicKeyEncryptionKey
}
var file_crypt_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crypt_proto_init() }
func file_crypt_proto_init() {
	if File_crypt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SealedMessage); i {
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
		file_crypt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKeyEncryptionKey); i {
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
			RawDescriptor: file_crypt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypt_proto_goTypes,
		DependencyIndexes: file_crypt_proto_depIdxs,
		MessageInfos:      file_crypt_proto_msgTypes,
	}.Build()
	File_crypt_proto = out.File
	file_crypt_proto_rawDesc = nil
	file_crypt_proto_goTypes = nil
	file_crypt_proto_depIdxs = nil
}