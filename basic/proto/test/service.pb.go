// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic/proto/test/service.proto

package test

import (
	context "context"
	fmt "fmt"
	sub "github.com/flutterWang/learningRPC/basic/proto/sub"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type DocumentExtType int32

const (
	DocumentExtType_HTML DocumentExtType = 0
	DocumentExtType_XML  DocumentExtType = 1
	DocumentExtType_JSON DocumentExtType = 2
	DocumentExtType_PDF  DocumentExtType = 3
)

var DocumentExtType_name = map[int32]string{
	0: "HTML",
	1: "XML",
	2: "JSON",
	3: "PDF",
}

var DocumentExtType_value = map[string]int32{
	"HTML": 0,
	"XML":  1,
	"JSON": 2,
	"PDF":  3,
}

func (x DocumentExtType) String() string {
	return proto.EnumName(DocumentExtType_name, int32(x))
}

func (DocumentExtType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62acc759e806be61, []int{0}
}

type TestRequest struct {
	Query    *wrappers.StringValue `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Type     DocumentExtType       `protobuf:"varint,2,opt,name=type,proto3,enum=test.DocumentExtType" json:"type,omitempty"`
	TestMap  map[int32]string      `protobuf:"bytes,3,rep,name=testMap,proto3" json:"testMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Child    *TestRequest_Child    `protobuf:"bytes,4,opt,name=child,proto3" json:"child,omitempty"`
	Char     []byte                `protobuf:"bytes,5,opt,name=char,proto3" json:"char,omitempty"`
	Snippets []string              `protobuf:"bytes,6,rep,name=snippets,proto3" json:"snippets,omitempty"`
	// Types that are valid to be assigned to TestOneof:
	//	*TestRequest_Name
	//	*TestRequest_SubName
	TestOneof            isTestRequest_TestOneof `protobuf_oneof:"test_oneof"`
	Submessage           *sub.SubMessage         `protobuf:"bytes,9,opt,name=submessage,proto3" json:"submessage,omitempty"`
	Details              []*any.Any              `protobuf:"bytes,10,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62acc759e806be61, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetQuery() *wrappers.StringValue {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *TestRequest) GetType() DocumentExtType {
	if m != nil {
		return m.Type
	}
	return DocumentExtType_HTML
}

func (m *TestRequest) GetTestMap() map[int32]string {
	if m != nil {
		return m.TestMap
	}
	return nil
}

func (m *TestRequest) GetChild() *TestRequest_Child {
	if m != nil {
		return m.Child
	}
	return nil
}

func (m *TestRequest) GetChar() []byte {
	if m != nil {
		return m.Char
	}
	return nil
}

func (m *TestRequest) GetSnippets() []string {
	if m != nil {
		return m.Snippets
	}
	return nil
}

type isTestRequest_TestOneof interface {
	isTestRequest_TestOneof()
}

type TestRequest_Name struct {
	Name string `protobuf:"bytes,7,opt,name=name,proto3,oneof"`
}

type TestRequest_SubName struct {
	SubName string `protobuf:"bytes,8,opt,name=sub_name,json=subName,proto3,oneof"`
}

func (*TestRequest_Name) isTestRequest_TestOneof() {}

func (*TestRequest_SubName) isTestRequest_TestOneof() {}

func (m *TestRequest) GetTestOneof() isTestRequest_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (m *TestRequest) GetName() string {
	if x, ok := m.GetTestOneof().(*TestRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (m *TestRequest) GetSubName() string {
	if x, ok := m.GetTestOneof().(*TestRequest_SubName); ok {
		return x.SubName
	}
	return ""
}

func (m *TestRequest) GetSubmessage() *sub.SubMessage {
	if m != nil {
		return m.Submessage
	}
	return nil
}

func (m *TestRequest) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestRequest_Name)(nil),
		(*TestRequest_SubName)(nil),
	}
}

type TestRequest_Child struct {
	Name                 int32    `protobuf:"varint,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest_Child) Reset()         { *m = TestRequest_Child{} }
func (m *TestRequest_Child) String() string { return proto.CompactTextString(m) }
func (*TestRequest_Child) ProtoMessage()    {}
func (*TestRequest_Child) Descriptor() ([]byte, []int) {
	return fileDescriptor_62acc759e806be61, []int{0, 1}
}

func (m *TestRequest_Child) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest_Child.Unmarshal(m, b)
}
func (m *TestRequest_Child) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest_Child.Marshal(b, m, deterministic)
}
func (m *TestRequest_Child) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest_Child.Merge(m, src)
}
func (m *TestRequest_Child) XXX_Size() int {
	return xxx_messageInfo_TestRequest_Child.Size(m)
}
func (m *TestRequest_Child) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest_Child.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest_Child proto.InternalMessageInfo

func (m *TestRequest_Child) GetName() int32 {
	if m != nil {
		return m.Name
	}
	return 0
}

type TestResponse struct {
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Details              *any.Any `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62acc759e806be61, []int{1}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *TestResponse) GetDetails() *any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

type TestDetail struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestDetail) Reset()         { *m = TestDetail{} }
func (m *TestDetail) String() string { return proto.CompactTextString(m) }
func (*TestDetail) ProtoMessage()    {}
func (*TestDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_62acc759e806be61, []int{2}
}

func (m *TestDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestDetail.Unmarshal(m, b)
}
func (m *TestDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestDetail.Marshal(b, m, deterministic)
}
func (m *TestDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestDetail.Merge(m, src)
}
func (m *TestDetail) XXX_Size() int {
	return xxx_messageInfo_TestDetail.Size(m)
}
func (m *TestDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TestDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TestDetail proto.InternalMessageInfo

func (m *TestDetail) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("test.DocumentExtType", DocumentExtType_name, DocumentExtType_value)
	proto.RegisterType((*TestRequest)(nil), "test.TestRequest")
	proto.RegisterMapType((map[int32]string)(nil), "test.TestRequest.TestMapEntry")
	proto.RegisterType((*TestRequest_Child)(nil), "test.TestRequest.Child")
	proto.RegisterType((*TestResponse)(nil), "test.TestResponse")
	proto.RegisterType((*TestDetail)(nil), "test.TestDetail")
}

func init() {
	proto.RegisterFile("basic/proto/test/service.proto", fileDescriptor_62acc759e806be61)
}

var fileDescriptor_62acc759e806be61 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdf, 0x6b, 0xdb, 0x30,
	0x10, 0xae, 0x13, 0xa7, 0x49, 0xae, 0x65, 0xcd, 0x44, 0xc7, 0x5c, 0x77, 0x04, 0x93, 0x27, 0x6f,
	0x30, 0x1b, 0x52, 0x46, 0x4b, 0xdf, 0xd6, 0x1f, 0xa3, 0x8c, 0xa6, 0x2b, 0x6a, 0xd8, 0xc6, 0x5e,
	0x8a, 0xec, 0x5e, 0x5d, 0x33, 0x47, 0x76, 0x2d, 0xa9, 0x9b, 0xdf, 0xf6, 0xa7, 0x0f, 0x49, 0xc9,
	0x30, 0x0d, 0xec, 0xc1, 0xf8, 0xee, 0xbe, 0xef, 0xce, 0xdf, 0x7d, 0x92, 0x61, 0x9c, 0x30, 0x91,
	0xa7, 0x71, 0x55, 0x97, 0xb2, 0x8c, 0x25, 0x0a, 0x19, 0x0b, 0xac, 0x9f, 0xf2, 0x14, 0x23, 0x53,
	0x22, 0xae, 0xae, 0xf9, 0x7b, 0x59, 0x59, 0x66, 0x05, 0x5a, 0x5a, 0xa2, 0xee, 0x63, 0xc6, 0x1b,
	0x4b, 0xf0, 0xf7, 0xda, 0x03, 0x84, 0x4a, 0xf4, 0xb3, 0x84, 0xc6, 0xcf, 0xbb, 0x7e, 0xd5, 0xac,
	0xaa, 0xb0, 0x16, 0x16, 0x9f, 0xfc, 0x71, 0x61, 0x6b, 0x8e, 0x42, 0x52, 0x7c, 0x54, 0x28, 0x24,
	0x99, 0x42, 0xef, 0x51, 0x61, 0xdd, 0x78, 0x4e, 0xe0, 0x84, 0x5b, 0xd3, 0x37, 0x91, 0xed, 0x8f,
	0x56, 0xfd, 0xd1, 0x8d, 0xac, 0x73, 0x9e, 0x7d, 0x65, 0x85, 0x42, 0x6a, 0xa9, 0xe4, 0x2d, 0xb8,
	0xb2, 0xa9, 0xd0, 0xeb, 0x04, 0x4e, 0xf8, 0x62, 0xfa, 0x2a, 0xd2, 0x72, 0xa3, 0xb3, 0x32, 0x55,
	0x0b, 0xe4, 0xf2, 0xfc, 0xb7, 0x9c, 0x37, 0x15, 0x52, 0x43, 0x21, 0x47, 0xd0, 0xd7, 0xe8, 0x8c,
	0x55, 0x5e, 0x37, 0xe8, 0x86, 0x5b, 0xd3, 0xb1, 0x65, 0xb7, 0x24, 0x98, 0x78, 0xc6, 0xaa, 0x73,
	0x2e, 0xeb, 0x86, 0xae, 0xe8, 0xe4, 0x3d, 0xf4, 0xd2, 0x87, 0xbc, 0xb8, 0xf3, 0x5c, 0x23, 0xec,
	0xf5, 0x7a, 0xdf, 0xa9, 0x86, 0xa9, 0x65, 0x11, 0x02, 0x6e, 0xfa, 0xc0, 0x6a, 0xaf, 0x17, 0x38,
	0xe1, 0x36, 0x35, 0x31, 0xf1, 0x61, 0x20, 0x78, 0x5e, 0x55, 0x28, 0x85, 0xb7, 0x19, 0x74, 0xc3,
	0x21, 0xfd, 0x97, 0x93, 0x5d, 0x70, 0x39, 0x5b, 0xa0, 0xd7, 0x0f, 0x9c, 0x70, 0x78, 0xb1, 0x41,
	0x4d, 0x46, 0xf6, 0x61, 0x20, 0x54, 0x72, 0x6b, 0x90, 0xc1, 0x12, 0xe9, 0x0b, 0x95, 0x5c, 0x69,
	0x30, 0x06, 0x10, 0x2a, 0x59, 0xa0, 0x10, 0x2c, 0x43, 0x6f, 0x68, 0x64, 0xed, 0x44, 0xda, 0xfa,
	0x1b, 0x95, 0xcc, 0x6c, 0x99, 0xb6, 0x28, 0x24, 0x82, 0xfe, 0x1d, 0x4a, 0x96, 0x17, 0xc2, 0x03,
	0xb3, 0xfc, 0xee, 0x9a, 0xbb, 0x1f, 0x79, 0x43, 0x57, 0x24, 0xff, 0x18, 0xb6, 0xdb, 0x5e, 0x90,
	0x11, 0x74, 0x7f, 0xa2, 0x3d, 0x99, 0x1e, 0xd5, 0x21, 0xd9, 0x85, 0xde, 0x93, 0x3e, 0x09, 0x63,
	0xfd, 0x90, 0xda, 0xe4, 0xb8, 0x73, 0xe4, 0xf8, 0xfb, 0xd0, 0x3b, 0x5d, 0x19, 0x61, 0xe4, 0xdb,
	0x2e, 0x13, 0x9f, 0x6c, 0x03, 0x68, 0xf7, 0x6e, 0x4b, 0x8e, 0xe5, 0xfd, 0x64, 0x6e, 0x3f, 0x43,
	0x51, 0x54, 0x25, 0x17, 0xa8, 0x87, 0x0a, 0xc9, 0xa4, 0x6d, 0x19, 0x52, 0x9b, 0xb4, 0xc5, 0x77,
	0xcc, 0xaa, 0xff, 0x17, 0x3f, 0x19, 0x03, 0xe8, 0xa9, 0x67, 0x26, 0xd5, 0xd2, 0x17, 0x22, 0x5b,
	0x4e, 0xd4, 0xe1, 0xbb, 0x43, 0xd8, 0x79, 0x76, 0x45, 0xc8, 0x00, 0xdc, 0x8b, 0xf9, 0xec, 0x72,
	0xb4, 0x41, 0xfa, 0xd0, 0xfd, 0x3e, 0xbb, 0x1c, 0x39, 0xba, 0xf4, 0xf9, 0xe6, 0xcb, 0xd5, 0xa8,
	0xa3, 0x4b, 0xd7, 0x67, 0x9f, 0x46, 0xdd, 0xe9, 0x21, 0xb8, 0x7a, 0x30, 0x89, 0x97, 0xef, 0x97,
	0x6b, 0x37, 0xc1, 0x27, 0xed, 0x92, 0xdd, 0x6a, 0xb2, 0x71, 0xf2, 0xe1, 0xc7, 0x41, 0x96, 0xcb,
	0x07, 0x95, 0x44, 0x69, 0xb9, 0x88, 0xef, 0x0b, 0x25, 0x25, 0xd6, 0xdf, 0x18, 0xcf, 0xe2, 0x02,
	0x59, 0xcd, 0x73, 0x9e, 0xd1, 0xeb, 0xd3, 0xf8, 0xf9, 0xbf, 0x98, 0x6c, 0x9a, 0xf8, 0xe0, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xa3, 0xbd, 0x92, 0xa6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	// Sends a greeting
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/test.Test/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	// Sends a greeting
	Test(context.Context, *TestRequest) (*TestResponse, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) Test(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Test/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Test_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic/proto/test/service.proto",
}
