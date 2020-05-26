// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chunk.proto

package chunk

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c46bd41e8571bd, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Chunk struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c46bd41e8571bd, []int{1}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "chunk.Empty")
	proto.RegisterType((*Chunk)(nil), "chunk.Chunk")
}

func init() {
	proto.RegisterFile("chunk.proto", fileDescriptor_67c46bd41e8571bd)
}

var fileDescriptor_67c46bd41e8571bd = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xce, 0x28, 0xcd,
	0xcb, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xd8, 0xb9, 0x58, 0x5d,
	0x73, 0x0b, 0x4a, 0x2a, 0x95, 0x64, 0xb9, 0x58, 0x9d, 0x41, 0x22, 0x42, 0x22, 0x5c, 0x10, 0x29,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x08, 0xc7, 0xc8, 0x84, 0x8b, 0x1d, 0x2c, 0x9d, 0x5a,
	0x24, 0xa4, 0x89, 0x60, 0xf2, 0xe8, 0x41, 0x8c, 0x04, 0x1b, 0x21, 0x05, 0xe3, 0x81, 0x65, 0x95,
	0x18, 0x0c, 0x18, 0x9d, 0xcc, 0xa2, 0x4c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0xd3, 0x72, 0x4a, 0x4b, 0x4a, 0x52, 0x8b, 0xc2, 0x13, 0xf3, 0xd2, 0xf5, 0x73, 0x52,
	0x13, 0x8b, 0xf2, 0x32, 0xf3, 0xd2, 0x83, 0x02, 0x9c, 0xf5, 0x93, 0x12, 0x8b, 0x33, 0x93, 0xf5,
	0xc1, 0xce, 0xd2, 0x07, 0x9b, 0x90, 0xc4, 0x06, 0xe6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x6a, 0xf8, 0xb8, 0xb2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChunkerClient is the client API for Chunker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChunkerClient interface {
	Chunker(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chunker_ChunkerClient, error)
}

type chunkerClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkerClient(cc grpc.ClientConnInterface) ChunkerClient {
	return &chunkerClient{cc}
}

func (c *chunkerClient) Chunker(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chunker_ChunkerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chunker_serviceDesc.Streams[0], "/chunk.Chunker/Chunker", opts...)
	if err != nil {
		return nil, err
	}
	x := &chunkerChunkerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chunker_ChunkerClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type chunkerChunkerClient struct {
	grpc.ClientStream
}

func (x *chunkerChunkerClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChunkerServer is the server API for Chunker service.
type ChunkerServer interface {
	Chunker(*Empty, Chunker_ChunkerServer) error
}

// UnimplementedChunkerServer can be embedded to have forward compatible implementations.
type UnimplementedChunkerServer struct {
}

func (*UnimplementedChunkerServer) Chunker(req *Empty, srv Chunker_ChunkerServer) error {
	return status.Errorf(codes.Unimplemented, "method Chunker not implemented")
}

func RegisterChunkerServer(s *grpc.Server, srv ChunkerServer) {
	s.RegisterService(&_Chunker_serviceDesc, srv)
}

func _Chunker_Chunker_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChunkerServer).Chunker(m, &chunkerChunkerServer{stream})
}

type Chunker_ChunkerServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type chunkerChunkerServer struct {
	grpc.ServerStream
}

func (x *chunkerChunkerServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Chunker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chunk.Chunker",
	HandlerType: (*ChunkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chunker",
			Handler:       _Chunker_Chunker_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chunk.proto",
}