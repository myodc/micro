// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/config/proto/config.proto

package config

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

type ChangeSet struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Checksum             string   `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Format               string   `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	Source               string   `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeSet) Reset()         { *m = ChangeSet{} }
func (m *ChangeSet) String() string { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()    {}
func (*ChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{0}
}

func (m *ChangeSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeSet.Unmarshal(m, b)
}
func (m *ChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeSet.Marshal(b, m, deterministic)
}
func (m *ChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeSet.Merge(m, src)
}
func (m *ChangeSet) XXX_Size() int {
	return xxx_messageInfo_ChangeSet.Size(m)
}
func (m *ChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeSet proto.InternalMessageInfo

func (m *ChangeSet) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ChangeSet) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *ChangeSet) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *ChangeSet) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ChangeSet) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Change struct {
	Namespace            string     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path                 string     `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	ChangeSet            *ChangeSet `protobuf:"bytes,3,opt,name=changeSet,proto3" json:"changeSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{1}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Change) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Change) GetChangeSet() *ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

type CreateRequest struct {
	Change               *Change  `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type UpdateRequest struct {
	Change               *Change  `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{5}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Change               *Change  `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{7}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{8}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListResponse struct {
	Values               []*Change `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{9}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetValues() []*Change {
	if m != nil {
		return m.Values
	}
	return nil
}

type ReadRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{10}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ReadRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ReadResponse struct {
	Change               *Change  `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{11}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type WatchRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{12}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WatchRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type WatchResponse struct {
	Namespace            string     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ChangeSet            *ChangeSet `protobuf:"bytes,2,opt,name=changeSet,proto3" json:"changeSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f3d36580b48e31, []int{13}
}

func (m *WatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchResponse.Unmarshal(m, b)
}
func (m *WatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchResponse.Marshal(b, m, deterministic)
}
func (m *WatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchResponse.Merge(m, src)
}
func (m *WatchResponse) XXX_Size() int {
	return xxx_messageInfo_WatchResponse.Size(m)
}
func (m *WatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchResponse proto.InternalMessageInfo

func (m *WatchResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WatchResponse) GetChangeSet() *ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeSet)(nil), "config.ChangeSet")
	proto.RegisterType((*Change)(nil), "config.Change")
	proto.RegisterType((*CreateRequest)(nil), "config.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "config.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "config.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "config.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "config.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "config.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "config.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "config.ListResponse")
	proto.RegisterType((*ReadRequest)(nil), "config.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "config.ReadResponse")
	proto.RegisterType((*WatchRequest)(nil), "config.WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "config.WatchResponse")
}

func init() { proto.RegisterFile("service/config/proto/config.proto", fileDescriptor_10f3d36580b48e31) }

var fileDescriptor_10f3d36580b48e31 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x6d, 0xda, 0x6e, 0xb0, 0xb7, 0xbb, 0x8b, 0x8e, 0xee, 0x12, 0x8a, 0x0f, 0x35, 0x0f, 0xb2,
	0x20, 0x34, 0xd2, 0xa2, 0xab, 0xf8, 0xa0, 0x58, 0x1f, 0x7d, 0x8a, 0x88, 0xe0, 0x83, 0x30, 0x3b,
	0xbd, 0xdb, 0x84, 0xdd, 0x64, 0x62, 0x66, 0xd2, 0xff, 0xe0, 0x8f, 0xf4, 0xbf, 0xc8, 0x7c, 0x24,
	0x93, 0x84, 0xb2, 0x96, 0xbe, 0x94, 0xb9, 0x1f, 0xe7, 0x9c, 0x7b, 0x7b, 0x0f, 0x81, 0x17, 0x02,
	0xcb, 0x5d, 0xca, 0x30, 0x62, 0x3c, 0xbf, 0x4d, 0xb7, 0x51, 0x51, 0x72, 0xc9, 0x6d, 0xb0, 0xd0,
	0x01, 0xf1, 0x4d, 0x14, 0xfe, 0xf1, 0x60, 0xb2, 0x4e, 0x68, 0xbe, 0xc5, 0x6f, 0x28, 0x09, 0x81,
	0xf1, 0x86, 0x4a, 0x1a, 0x78, 0x73, 0xef, 0x6a, 0x12, 0xeb, 0x37, 0x99, 0xc1, 0x23, 0x96, 0x20,
	0xbb, 0x13, 0x55, 0x16, 0x0c, 0x75, 0xbe, 0x89, 0xc9, 0x25, 0xf8, 0xb7, 0xbc, 0xcc, 0xa8, 0x0c,
	0x46, 0xba, 0x62, 0x23, 0x95, 0x17, 0xbc, 0x2a, 0x19, 0x06, 0x63, 0x93, 0x37, 0x11, 0x79, 0x0e,
	0x13, 0x99, 0x66, 0x28, 0x24, 0xcd, 0x8a, 0xe0, 0x64, 0xee, 0x5d, 0x8d, 0x62, 0x97, 0x08, 0xef,
	0xc0, 0x37, 0xa3, 0xa8, 0xbe, 0x9c, 0x66, 0x28, 0x0a, 0xca, 0xd0, 0x0e, 0xe3, 0x12, 0x6a, 0xca,
	0x82, 0xca, 0xc4, 0x4e, 0xa3, 0xdf, 0x24, 0x82, 0x09, 0xab, 0xd7, 0xd0, 0xc3, 0x4c, 0x97, 0x4f,
	0x16, 0x76, 0xe3, 0x66, 0xbf, 0xd8, 0xf5, 0x84, 0xd7, 0x70, 0xb6, 0x2e, 0x91, 0x4a, 0x8c, 0xf1,
	0x77, 0x85, 0x42, 0x92, 0x97, 0xe0, 0x9b, 0xaa, 0x16, 0x9c, 0x2e, 0xcf, 0xbb, 0xf0, 0xd8, 0x56,
	0xc3, 0xc7, 0x70, 0x5e, 0x03, 0x45, 0xc1, 0x73, 0x81, 0x8a, 0xea, 0x7b, 0xb1, 0x39, 0x8e, 0xaa,
	0x06, 0x3a, 0xaa, 0x2f, 0x78, 0x8f, 0x47, 0x51, 0xd5, 0x40, 0x4b, 0xf5, 0x0a, 0xa6, 0x5f, 0x53,
	0x21, 0x6b, 0xa2, 0x07, 0xff, 0xd2, 0xf0, 0x2d, 0x9c, 0x9a, 0x66, 0x03, 0x56, 0xb2, 0x3b, 0x7a,
	0x5f, 0xa1, 0x08, 0xbc, 0xf9, 0x68, 0x9f, 0xac, 0xa9, 0x86, 0x1f, 0x61, 0x1a, 0x23, 0xdd, 0x1c,
	0x24, 0xb2, 0xef, 0x6e, 0x4a, 0xd8, 0x10, 0x38, 0xe1, 0x83, 0xf6, 0xfd, 0x04, 0xa7, 0x3f, 0xa8,
	0x64, 0xc9, 0xf1, 0xca, 0xbf, 0xe0, 0xcc, 0x32, 0x58, 0xe9, 0x87, 0x29, 0x3a, 0x06, 0x1b, 0xfe,
	0xdf, 0x60, 0xcb, 0xbf, 0x43, 0xf0, 0xd7, 0xba, 0x4e, 0xde, 0x83, 0x6f, 0x2c, 0x43, 0x2e, 0x1a,
	0x48, 0xdb, 0x7b, 0xb3, 0xcb, 0x7e, 0xda, 0xde, 0x70, 0xa0, 0xa0, 0xc6, 0x22, 0x0e, 0xda, 0xf1,
	0x9a, 0x83, 0xf6, 0x9c, 0xa4, 0xa1, 0xc6, 0x12, 0x0e, 0xda, 0xf1, 0x96, 0x83, 0xf6, 0x9c, 0x33,
	0x20, 0x2b, 0x18, 0x2b, 0x3b, 0x90, 0xa7, 0x75, 0x47, 0xcb, 0x49, 0xb3, 0x67, 0xdd, 0x64, 0x1b,
	0xa4, 0x4e, 0xe9, 0x40, 0x2d, 0x67, 0x38, 0x50, 0xfb, 0xda, 0xe1, 0x80, 0xbc, 0x83, 0x13, 0x7d,
	0x05, 0xd2, 0x34, 0xb4, 0xcf, 0x3a, 0xbb, 0xe8, 0x65, 0x6b, 0xdc, 0x6b, 0xef, 0xf3, 0xf5, 0xcf,
	0x37, 0xdb, 0x54, 0x26, 0xd5, 0xcd, 0x82, 0xf1, 0x2c, 0xca, 0x52, 0x56, 0x72, 0xfb, 0xbb, 0x5b,
	0x45, 0xfb, 0x3e, 0x80, 0x1f, 0x4c, 0x70, 0xe3, 0xeb, 0x68, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0xa0, 0x43, 0x27, 0x57, 0x26, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigClient is the client API for Config service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Config_WatchClient, error)
}

type configClient struct {
	cc *grpc.ClientConn
}

func NewConfigClient(cc *grpc.ClientConn) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/config.Config/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/config.Config/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/config.Config/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/config.Config/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/config.Config/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Config_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Config_serviceDesc.Streams[0], "/config.Config/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &configWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Config_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type configWatchClient struct {
	grpc.ClientStream
}

func (x *configWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConfigServer is the server API for Config service.
type ConfigServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Watch(*WatchRequest, Config_WatchServer) error
}

// UnimplementedConfigServer can be embedded to have forward compatible implementations.
type UnimplementedConfigServer struct {
}

func (*UnimplementedConfigServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedConfigServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedConfigServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedConfigServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedConfigServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedConfigServer) Watch(req *WatchRequest, srv Config_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterConfigServer(s *grpc.Server, srv ConfigServer) {
	s.RegisterService(&_Config_serviceDesc, srv)
}

func _Config_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Config/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Config/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Config/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Config/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Config/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigServer).Watch(m, &configWatchServer{stream})
}

type Config_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type configWatchServer struct {
	grpc.ServerStream
}

func (x *configWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Config_serviceDesc = grpc.ServiceDesc{
	ServiceName: "config.Config",
	HandlerType: (*ConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Config_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Config_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Config_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Config_List_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Config_Read_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Config_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/config/proto/config.proto",
}
