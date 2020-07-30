// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/bot/proto/bot.proto

package go_micro_bot

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

type HelpRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelpRequest) Reset()         { *m = HelpRequest{} }
func (m *HelpRequest) String() string { return proto.CompactTextString(m) }
func (*HelpRequest) ProtoMessage()    {}
func (*HelpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e76f3af650717012, []int{0}
}

func (m *HelpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelpRequest.Unmarshal(m, b)
}
func (m *HelpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelpRequest.Marshal(b, m, deterministic)
}
func (m *HelpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelpRequest.Merge(m, src)
}
func (m *HelpRequest) XXX_Size() int {
	return xxx_messageInfo_HelpRequest.Size(m)
}
func (m *HelpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelpRequest proto.InternalMessageInfo

type HelpResponse struct {
	Usage                string   `protobuf:"bytes,1,opt,name=usage,proto3" json:"usage,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelpResponse) Reset()         { *m = HelpResponse{} }
func (m *HelpResponse) String() string { return proto.CompactTextString(m) }
func (*HelpResponse) ProtoMessage()    {}
func (*HelpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e76f3af650717012, []int{1}
}

func (m *HelpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelpResponse.Unmarshal(m, b)
}
func (m *HelpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelpResponse.Marshal(b, m, deterministic)
}
func (m *HelpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelpResponse.Merge(m, src)
}
func (m *HelpResponse) XXX_Size() int {
	return xxx_messageInfo_HelpResponse.Size(m)
}
func (m *HelpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelpResponse proto.InternalMessageInfo

func (m *HelpResponse) GetUsage() string {
	if m != nil {
		return m.Usage
	}
	return ""
}

func (m *HelpResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ExecRequest struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecRequest) Reset()         { *m = ExecRequest{} }
func (m *ExecRequest) String() string { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()    {}
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e76f3af650717012, []int{2}
}

func (m *ExecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecRequest.Unmarshal(m, b)
}
func (m *ExecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecRequest.Marshal(b, m, deterministic)
}
func (m *ExecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecRequest.Merge(m, src)
}
func (m *ExecRequest) XXX_Size() int {
	return xxx_messageInfo_ExecRequest.Size(m)
}
func (m *ExecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecRequest proto.InternalMessageInfo

func (m *ExecRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type ExecResponse struct {
	Result               []byte   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecResponse) Reset()         { *m = ExecResponse{} }
func (m *ExecResponse) String() string { return proto.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()    {}
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e76f3af650717012, []int{3}
}

func (m *ExecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecResponse.Unmarshal(m, b)
}
func (m *ExecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecResponse.Marshal(b, m, deterministic)
}
func (m *ExecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecResponse.Merge(m, src)
}
func (m *ExecResponse) XXX_Size() int {
	return xxx_messageInfo_ExecResponse.Size(m)
}
func (m *ExecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecResponse proto.InternalMessageInfo

func (m *ExecResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ExecResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*HelpRequest)(nil), "go.micro.bot.HelpRequest")
	proto.RegisterType((*HelpResponse)(nil), "go.micro.bot.HelpResponse")
	proto.RegisterType((*ExecRequest)(nil), "go.micro.bot.ExecRequest")
	proto.RegisterType((*ExecResponse)(nil), "go.micro.bot.ExecResponse")
}

func init() { proto.RegisterFile("client/bot/proto/bot.proto", fileDescriptor_e76f3af650717012) }

var fileDescriptor_e76f3af650717012 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0xdd, 0xea, 0xba, 0xb2, 0xd3, 0x7a, 0x09, 0x22, 0xb5, 0xa7, 0x9a, 0xd3, 0x9e, 0xb2, 0xa0,
	0x57, 0xc1, 0x83, 0x28, 0x9e, 0xfb, 0x0f, 0xda, 0xee, 0x50, 0x0a, 0x6d, 0x27, 0x4e, 0x52, 0xf0,
	0x3f, 0xf8, 0xa7, 0x25, 0x1f, 0x87, 0x50, 0xbc, 0xbd, 0x97, 0x17, 0xde, 0xc7, 0x40, 0xd5, 0x4f,
	0x23, 0x2e, 0xf6, 0xdc, 0x91, 0x3d, 0x6b, 0x26, 0x4b, 0x0e, 0x29, 0x8f, 0x44, 0x31, 0x90, 0x9a,
	0xc7, 0x9e, 0x49, 0x75, 0x64, 0xe5, 0x1d, 0xe4, 0x5f, 0x38, 0xe9, 0x06, 0xbf, 0x57, 0x34, 0x56,
	0x7e, 0x42, 0x11, 0xa8, 0xd1, 0xb4, 0x18, 0x14, 0xf7, 0x70, 0xb3, 0x9a, 0x76, 0xc0, 0x32, 0xab,
	0xb3, 0xd3, 0xb1, 0x09, 0x44, 0xd4, 0x90, 0x5f, 0xd0, 0xf4, 0x3c, 0x6a, 0x3b, 0xd2, 0x52, 0x5e,
	0x79, 0x2d, 0x7d, 0x92, 0x4f, 0x90, 0x7f, 0xfc, 0x60, 0x1f, 0x6d, 0x85, 0x80, 0x7d, 0xcb, 0x83,
	0x29, 0xb3, 0xfa, 0xfa, 0x74, 0x6c, 0x3c, 0x96, 0xaf, 0x50, 0x84, 0x2f, 0x31, 0xea, 0x01, 0x0e,
	0x8c, 0x66, 0x9d, 0xac, 0xcf, 0x2a, 0x9a, 0xc8, 0x5c, 0x05, 0x64, 0x26, 0x8e, 0x31, 0x81, 0x3c,
	0xff, 0x66, 0x70, 0xfb, 0x4e, 0xf3, 0xdc, 0x2e, 0x17, 0xf1, 0x06, 0x7b, 0x57, 0x5a, 0x3c, 0xaa,
	0x74, 0x9a, 0x4a, 0x76, 0x55, 0xd5, 0x7f, 0x52, 0x08, 0x96, 0x3b, 0x67, 0xe0, 0xaa, 0x6c, 0x0d,
	0x92, 0x05, 0x5b, 0x83, 0xb4, 0xb9, 0xdc, 0x75, 0x07, 0x7f, 0xda, 0x97, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0x06, 0x95, 0xc0, 0x78, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandClient is the client API for Command service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandClient interface {
	Help(ctx context.Context, in *HelpRequest, opts ...grpc.CallOption) (*HelpResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
}

type commandClient struct {
	cc *grpc.ClientConn
}

func NewCommandClient(cc *grpc.ClientConn) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) Help(ctx context.Context, in *HelpRequest, opts ...grpc.CallOption) (*HelpResponse, error) {
	out := new(HelpResponse)
	err := c.cc.Invoke(ctx, "/go.micro.bot.Command/Help", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := c.cc.Invoke(ctx, "/go.micro.bot.Command/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServer is the server API for Command service.
type CommandServer interface {
	Help(context.Context, *HelpRequest) (*HelpResponse, error)
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
}

// UnimplementedCommandServer can be embedded to have forward compatible implementations.
type UnimplementedCommandServer struct {
}

func (*UnimplementedCommandServer) Help(ctx context.Context, req *HelpRequest) (*HelpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Help not implemented")
}
func (*UnimplementedCommandServer) Exec(ctx context.Context, req *ExecRequest) (*ExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

func RegisterCommandServer(s *grpc.Server, srv CommandServer) {
	s.RegisterService(&_Command_serviceDesc, srv)
}

func _Command_Help_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Help(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.bot.Command/Help",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Help(ctx, req.(*HelpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.bot.Command/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Command_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.bot.Command",
	HandlerType: (*CommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Help",
			Handler:    _Command_Help_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Command_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/bot/proto/bot.proto",
}
