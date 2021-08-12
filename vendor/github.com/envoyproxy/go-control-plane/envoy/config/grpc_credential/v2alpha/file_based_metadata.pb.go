// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/grpc_credential/v2alpha/file_based_metadata.proto

package envoy_config_grpc_credential_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FileBasedMetadataConfig struct {
	SecretData           *core.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	HeaderKey            string           `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	HeaderPrefix         string           `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FileBasedMetadataConfig) Reset()         { *m = FileBasedMetadataConfig{} }
func (m *FileBasedMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*FileBasedMetadataConfig) ProtoMessage()    {}
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f2b21de9d357383, []int{0}
}

func (m *FileBasedMetadataConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileBasedMetadataConfig.Unmarshal(m, b)
}
func (m *FileBasedMetadataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileBasedMetadataConfig.Marshal(b, m, deterministic)
}
func (m *FileBasedMetadataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBasedMetadataConfig.Merge(m, src)
}
func (m *FileBasedMetadataConfig) XXX_Size() int {
	return xxx_messageInfo_FileBasedMetadataConfig.Size(m)
}
func (m *FileBasedMetadataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBasedMetadataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FileBasedMetadataConfig proto.InternalMessageInfo

func (m *FileBasedMetadataConfig) GetSecretData() *core.DataSource {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *FileBasedMetadataConfig) GetHeaderKey() string {
	if m != nil {
		return m.HeaderKey
	}
	return ""
}

func (m *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if m != nil {
		return m.HeaderPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*FileBasedMetadataConfig)(nil), "envoy.config.grpc_credential.v2alpha.FileBasedMetadataConfig")
}

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v2alpha/file_based_metadata.proto", fileDescriptor_0f2b21de9d357383)
}

var fileDescriptor_0f2b21de9d357383 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0x3b, 0x31,
	0x10, 0xc7, 0x49, 0x7f, 0x50, 0x68, 0xfa, 0x13, 0x64, 0x0f, 0x5a, 0x8a, 0x85, 0xa2, 0x1e, 0x7a,
	0x4a, 0x60, 0xbd, 0x7b, 0x58, 0xa5, 0x17, 0x11, 0x4a, 0x05, 0xaf, 0xcb, 0x34, 0x3b, 0x6d, 0x83,
	0x6b, 0x12, 0x92, 0xec, 0xd2, 0xbd, 0x79, 0xf7, 0x35, 0xf4, 0x19, 0x14, 0x9f, 0xc0, 0xab, 0x6f,
	0x24, 0xd9, 0xac, 0x17, 0x7b, 0xf1, 0xfa, 0xfd, 0x33, 0xf3, 0x99, 0xa1, 0x97, 0xa8, 0x6a, 0xdd,
	0x70, 0xa1, 0xd5, 0x5a, 0x6e, 0xf8, 0xc6, 0x1a, 0x91, 0x0b, 0x8b, 0x05, 0x2a, 0x2f, 0xa1, 0xe4,
	0x75, 0x0a, 0xa5, 0xd9, 0x02, 0x5f, 0xcb, 0x12, 0xf3, 0x15, 0x38, 0x2c, 0xf2, 0x47, 0xf4, 0x50,
	0x80, 0x07, 0x66, 0xac, 0xf6, 0x3a, 0x39, 0x6f, 0xfb, 0x2c, 0xf6, 0xd9, 0xaf, 0x3e, 0xeb, 0xfa,
	0xe3, 0x93, 0xb8, 0x05, 0x8c, 0xe4, 0x75, 0xca, 0x85, 0xb6, 0xc8, 0xc3, 0xb4, 0x38, 0x63, 0x3c,
	0xad, 0x0a, 0x03, 0x1c, 0x94, 0xd2, 0x1e, 0xbc, 0xd4, 0xca, 0x71, 0x87, 0xca, 0x49, 0x2f, 0xeb,
	0x9f, 0xc4, 0x64, 0x3f, 0xe1, 0xc1, 0x57, 0x2e, 0xda, 0xa7, 0xaf, 0x84, 0x1e, 0xcf, 0x65, 0x89,
	0x59, 0x20, 0xbc, 0xed, 0x00, 0xaf, 0x5a, 0xa4, 0x64, 0x4e, 0x87, 0x0e, 0x85, 0x45, 0x9f, 0x07,
	0x71, 0x44, 0xa6, 0x64, 0x36, 0x4c, 0x27, 0x2c, 0x62, 0x83, 0x91, 0xac, 0x4e, 0x59, 0x00, 0x62,
	0xd7, 0xe0, 0xe1, 0x4e, 0x57, 0x56, 0x60, 0xd6, 0x7f, 0x7f, 0x7b, 0x7e, 0xe9, 0x91, 0x25, 0x8d,
	0xcd, 0xe0, 0x24, 0x13, 0x4a, 0xb7, 0x08, 0x05, 0xda, 0xfc, 0x01, 0x9b, 0x51, 0x6f, 0x4a, 0x66,
	0x83, 0xe5, 0x20, 0x2a, 0x37, 0xd8, 0x24, 0x67, 0xf4, 0xa0, 0xb3, 0x8d, 0xc5, 0xb5, 0xdc, 0x8d,
	0xfe, 0xb5, 0x89, 0xff, 0x51, 0x5c, 0xb4, 0x5a, 0x76, 0xff, 0xf1, 0xf4, 0xf9, 0xd5, 0xef, 0x1d,
	0x12, 0x9a, 0x4a, 0x1d, 0x11, 0x8c, 0xd5, 0xbb, 0x86, 0xfd, 0xe5, 0x89, 0xd9, 0xd1, 0xde, 0x89,
	0x8b, 0x70, 0xfd, 0x82, 0xac, 0xfa, 0xed, 0x1b, 0x2e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd7,
	0xd1, 0x73, 0x94, 0xcd, 0x01, 0x00, 0x00,
}