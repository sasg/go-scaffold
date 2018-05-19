// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/statestorage.proto

package statestorage // import "github.com/orbs-network/go-experiment/types/services/statestorage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WriteKeyInput struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteKeyInput) Reset()         { *m = WriteKeyInput{} }
func (m *WriteKeyInput) String() string { return proto.CompactTextString(m) }
func (*WriteKeyInput) ProtoMessage()    {}
func (*WriteKeyInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_statestorage_30ff1f66b12808fb, []int{0}
}
func (m *WriteKeyInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteKeyInput.Unmarshal(m, b)
}
func (m *WriteKeyInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteKeyInput.Marshal(b, m, deterministic)
}
func (dst *WriteKeyInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteKeyInput.Merge(dst, src)
}
func (m *WriteKeyInput) XXX_Size() int {
	return xxx_messageInfo_WriteKeyInput.Size(m)
}
func (m *WriteKeyInput) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteKeyInput.DiscardUnknown(m)
}

var xxx_messageInfo_WriteKeyInput proto.InternalMessageInfo

func (m *WriteKeyInput) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *WriteKeyInput) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type WriteKeyOutput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteKeyOutput) Reset()         { *m = WriteKeyOutput{} }
func (m *WriteKeyOutput) String() string { return proto.CompactTextString(m) }
func (*WriteKeyOutput) ProtoMessage()    {}
func (*WriteKeyOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_statestorage_30ff1f66b12808fb, []int{1}
}
func (m *WriteKeyOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteKeyOutput.Unmarshal(m, b)
}
func (m *WriteKeyOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteKeyOutput.Marshal(b, m, deterministic)
}
func (dst *WriteKeyOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteKeyOutput.Merge(dst, src)
}
func (m *WriteKeyOutput) XXX_Size() int {
	return xxx_messageInfo_WriteKeyOutput.Size(m)
}
func (m *WriteKeyOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteKeyOutput.DiscardUnknown(m)
}

var xxx_messageInfo_WriteKeyOutput proto.InternalMessageInfo

type ReadKeyInput struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadKeyInput) Reset()         { *m = ReadKeyInput{} }
func (m *ReadKeyInput) String() string { return proto.CompactTextString(m) }
func (*ReadKeyInput) ProtoMessage()    {}
func (*ReadKeyInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_statestorage_30ff1f66b12808fb, []int{2}
}
func (m *ReadKeyInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadKeyInput.Unmarshal(m, b)
}
func (m *ReadKeyInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadKeyInput.Marshal(b, m, deterministic)
}
func (dst *ReadKeyInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadKeyInput.Merge(dst, src)
}
func (m *ReadKeyInput) XXX_Size() int {
	return xxx_messageInfo_ReadKeyInput.Size(m)
}
func (m *ReadKeyInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadKeyInput.DiscardUnknown(m)
}

var xxx_messageInfo_ReadKeyInput proto.InternalMessageInfo

func (m *ReadKeyInput) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ReadKeyOutput struct {
	Value                int32    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadKeyOutput) Reset()         { *m = ReadKeyOutput{} }
func (m *ReadKeyOutput) String() string { return proto.CompactTextString(m) }
func (*ReadKeyOutput) ProtoMessage()    {}
func (*ReadKeyOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_statestorage_30ff1f66b12808fb, []int{3}
}
func (m *ReadKeyOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadKeyOutput.Unmarshal(m, b)
}
func (m *ReadKeyOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadKeyOutput.Marshal(b, m, deterministic)
}
func (dst *ReadKeyOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadKeyOutput.Merge(dst, src)
}
func (m *ReadKeyOutput) XXX_Size() int {
	return xxx_messageInfo_ReadKeyOutput.Size(m)
}
func (m *ReadKeyOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadKeyOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ReadKeyOutput proto.InternalMessageInfo

func (m *ReadKeyOutput) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*WriteKeyInput)(nil), "statestorage.WriteKeyInput")
	proto.RegisterType((*WriteKeyOutput)(nil), "statestorage.WriteKeyOutput")
	proto.RegisterType((*ReadKeyInput)(nil), "statestorage.ReadKeyInput")
	proto.RegisterType((*ReadKeyOutput)(nil), "statestorage.ReadKeyOutput")
}

func init() {
	proto.RegisterFile("services/statestorage.proto", fileDescriptor_statestorage_30ff1f66b12808fb)
}

var fileDescriptor_statestorage_30ff1f66b12808fb = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c,
	0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x16, 0x53, 0x32, 0xe7, 0xe2, 0x0d,
	0x2f, 0xca, 0x2c, 0x49, 0xf5, 0x4e, 0xad, 0xf4, 0xcc, 0x2b, 0x28, 0x2d, 0x11, 0x12, 0xe0, 0x62,
	0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58,
	0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0x25, 0x01,
	0x2e, 0x3e, 0x98, 0x46, 0xff, 0xd2, 0x92, 0x82, 0xd2, 0x12, 0x25, 0x05, 0x2e, 0x9e, 0xa0, 0xd4,
	0xc4, 0x14, 0xdc, 0x26, 0x29, 0xa9, 0x72, 0xf1, 0x42, 0x55, 0x40, 0xb4, 0x20, 0x8c, 0x66, 0x44,
	0x32, 0xda, 0x68, 0x26, 0x23, 0x17, 0x4f, 0x30, 0xc8, 0x91, 0xc1, 0x10, 0x47, 0x0a, 0xb9, 0x72,
	0x71, 0xc0, 0xec, 0x12, 0x92, 0xd6, 0x43, 0xf1, 0x13, 0x8a, 0xe3, 0xa5, 0x64, 0xb0, 0x4b, 0x42,
	0x6d, 0x73, 0xe2, 0x62, 0x87, 0x5a, 0x2f, 0x24, 0x85, 0xaa, 0x10, 0xd9, 0xdd, 0x52, 0xd2, 0x58,
	0xe5, 0x20, 0x66, 0x38, 0x39, 0x47, 0x39, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0xe7, 0x17, 0x25, 0x15, 0xeb, 0xe6, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0xeb, 0xa7,
	0xe7, 0xeb, 0xa6, 0x56, 0x14, 0xa4, 0x16, 0x65, 0xe6, 0xa6, 0xe6, 0x95, 0xe8, 0x97, 0x54, 0x16,
	0x80, 0xc2, 0x1f, 0x5b, 0x44, 0x24, 0xb1, 0x81, 0x63, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xdf, 0x57, 0x47, 0x13, 0xa8, 0x01, 0x00, 0x00,
}
