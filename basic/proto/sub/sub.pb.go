// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic/proto/sub/sub.proto

package sub

import (
	fmt "fmt"
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

type SubMessage struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubMessage) Reset()         { *m = SubMessage{} }
func (m *SubMessage) String() string { return proto.CompactTextString(m) }
func (*SubMessage) ProtoMessage()    {}
func (*SubMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_347e29edcdae8834, []int{0}
}

func (m *SubMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubMessage.Unmarshal(m, b)
}
func (m *SubMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubMessage.Marshal(b, m, deterministic)
}
func (m *SubMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubMessage.Merge(m, src)
}
func (m *SubMessage) XXX_Size() int {
	return xxx_messageInfo_SubMessage.Size(m)
}
func (m *SubMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubMessage proto.InternalMessageInfo

func (m *SubMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*SubMessage)(nil), "sub.SubMessage")
}

func init() {
	proto.RegisterFile("basic/proto/sub/sub.proto", fileDescriptor_347e29edcdae8834)
}

var fileDescriptor_347e29edcdae8834 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4a, 0x2c, 0xce,
	0x4c, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0x4d, 0x02, 0x61, 0x3d, 0x30, 0x4f,
	0x88, 0xb9, 0xb8, 0x34, 0x49, 0x49, 0x8e, 0x8b, 0x2b, 0xb8, 0x34, 0xc9, 0x37, 0xb5, 0xb8, 0x38,
	0x31, 0x3d, 0x55, 0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x5d, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xc4, 0x74, 0x32, 0x89, 0x32, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcb, 0x29, 0x2d, 0x29, 0x49, 0x2d, 0x0a, 0x4f, 0xcc, 0x4b, 0xd7, 0xcf, 0x49, 0x4d,
	0x2c, 0xca, 0xcb, 0xcc, 0x4b, 0x0f, 0x0a, 0x70, 0xd6, 0x47, 0xb3, 0x24, 0x89, 0x0d, 0xcc, 0x34,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x85, 0x58, 0xc1, 0x7e, 0x00, 0x00, 0x00,
}