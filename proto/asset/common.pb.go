// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/asset/common.proto

package asset

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

type ResultStatus int32

const (
	ResultStatus_Success     ResultStatus = 0
	ResultStatus_MaxLimit    ResultStatus = 1
	ResultStatus_Repeated    ResultStatus = 2
	ResultStatus_NotExisted  ResultStatus = 3
	ResultStatus_DBException ResultStatus = 4
	ResultStatus_Empty       ResultStatus = 5
)

var ResultStatus_name = map[int32]string{
	0: "Success",
	1: "MaxLimit",
	2: "Repeated",
	3: "NotExisted",
	4: "DBException",
	5: "Empty",
}

var ResultStatus_value = map[string]int32{
	"Success":     0,
	"MaxLimit":    1,
	"Repeated":    2,
	"NotExisted":  3,
	"DBException": 4,
	"Empty":       5,
}

func (x ResultStatus) String() string {
	return proto.EnumName(ResultStatus_name, int32(x))
}

func (ResultStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{0}
}

type RequestInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{0}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestFilter struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{1}
}

func (m *RequestFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFilter.Unmarshal(m, b)
}
func (m *RequestFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFilter.Marshal(b, m, deterministic)
}
func (m *RequestFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFilter.Merge(m, src)
}
func (m *RequestFilter) XXX_Size() int {
	return xxx_messageInfo_RequestFilter.Size(m)
}
func (m *RequestFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFilter proto.InternalMessageInfo

func (m *RequestFilter) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestFilter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestFilter) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RequestFilter) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ReplyStatus struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{2}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ReplyInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{3}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type RequestUpdate struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Field                string   `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Values               []string `protobuf:"bytes,6,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUpdate) Reset()         { *m = RequestUpdate{} }
func (m *RequestUpdate) String() string { return proto.CompactTextString(m) }
func (*RequestUpdate) ProtoMessage()    {}
func (*RequestUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{4}
}

func (m *RequestUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUpdate.Unmarshal(m, b)
}
func (m *RequestUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUpdate.Marshal(b, m, deterministic)
}
func (m *RequestUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUpdate.Merge(m, src)
}
func (m *RequestUpdate) XXX_Size() int {
	return xxx_messageInfo_RequestUpdate.Size(m)
}
func (m *RequestUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUpdate proto.InternalMessageInfo

func (m *RequestUpdate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestUpdate) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *RequestUpdate) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestUpdate) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterEnum("omo.msp.asset.ResultStatus", ResultStatus_name, ResultStatus_value)
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.asset.RequestInfo")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.asset.RequestFilter")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.asset.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.asset.ReplyInfo")
	proto.RegisterType((*RequestUpdate)(nil), "omo.msp.asset.RequestUpdate")
}

func init() {
	proto.RegisterFile("proto/asset/common.proto", fileDescriptor_3e70cb96243de9f1)
}

var fileDescriptor_3e70cb96243de9f1 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x8e, 0x9b, 0x30,
	0x10, 0x87, 0x4b, 0xf8, 0xd3, 0x30, 0x84, 0x16, 0x59, 0x55, 0x85, 0x72, 0x8a, 0x38, 0x45, 0x3d,
	0x10, 0x29, 0x3d, 0xf4, 0x1e, 0x35, 0x95, 0x2a, 0xb5, 0x95, 0xe2, 0xa8, 0x97, 0xde, 0x1c, 0x98,
	0x54, 0x68, 0x01, 0x7b, 0x6d, 0xb3, 0x1b, 0x9e, 0x64, 0x5f, 0x77, 0x65, 0x87, 0x4d, 0xc8, 0x21,
	0xb7, 0xf9, 0xcc, 0xcc, 0x6f, 0xe4, 0xcf, 0x40, 0x2a, 0x24, 0xd7, 0x7c, 0xc5, 0x94, 0x42, 0xbd,
	0x2a, 0x78, 0xd3, 0xf0, 0x36, 0xb7, 0x47, 0x24, 0xe6, 0x0d, 0xcf, 0x1b, 0x25, 0x72, 0xfb, 0x2d,
	0xdb, 0x41, 0x44, 0xf1, 0xb1, 0x43, 0xa5, 0x7f, 0xb6, 0x47, 0x4e, 0x12, 0x70, 0xbb, 0xaa, 0x4c,
	0x9d, 0x85, 0xb3, 0x0c, 0xa9, 0x29, 0xc9, 0x27, 0xf0, 0xf9, 0x73, 0x8b, 0x32, 0x9d, 0xd8, 0xb3,
	0x33, 0x90, 0x39, 0x4c, 0xb9, 0x40, 0xc9, 0x34, 0x97, 0xa9, 0x6b, 0x3f, 0x5c, 0x38, 0xeb, 0x21,
	0x1e, 0x22, 0x7f, 0x54, 0xb5, 0x46, 0x79, 0x8d, 0x70, 0xc6, 0x11, 0x09, 0xb8, 0x0f, 0xd8, 0x0f,
	0xb1, 0xa6, 0x34, 0x7d, 0x4f, 0xac, 0xee, 0x70, 0x48, 0x3c, 0x03, 0x21, 0xe0, 0x09, 0xf6, 0x1f,
	0x53, 0x6f, 0xe1, 0x2c, 0x63, 0x6a, 0x6b, 0xf2, 0x19, 0x82, 0xb6, 0x6b, 0x0e, 0x28, 0x53, 0xdf,
	0x9e, 0x0e, 0x94, 0x7d, 0x33, 0xb7, 0x11, 0x75, 0xbf, 0xd7, 0x4c, 0x77, 0xca, 0x8c, 0x16, 0xbc,
	0x44, 0xbb, 0x37, 0xa6, 0xb6, 0x36, 0x4b, 0x50, 0x4a, 0x7e, 0xb9, 0x8f, 0x85, 0x6c, 0x07, 0xa1,
	0x1d, 0xbc, 0x23, 0x61, 0x0d, 0x81, 0xb2, 0x91, 0x76, 0x2a, 0x5a, 0xcf, 0xf3, 0x1b, 0x8b, 0xf9,
	0x68, 0x29, 0x1d, 0x3a, 0xb3, 0x17, 0xe7, 0xe2, 0xe1, 0xaf, 0x28, 0x99, 0xc6, 0xfb, 0x1e, 0xcc,
	0xb6, 0xc9, 0x8d, 0xf2, 0x63, 0x85, 0x75, 0xf9, 0xe6, 0xc1, 0xc2, 0xd5, 0x8e, 0x37, 0xb6, 0x33,
	0x7e, 0x08, 0xff, 0xf6, 0x21, 0x8c, 0x25, 0xdb, 0xa4, 0xd2, 0x60, 0xe1, 0x2e, 0x43, 0x3a, 0xd0,
	0x97, 0x02, 0x66, 0x14, 0x55, 0x57, 0xeb, 0x41, 0x53, 0x04, 0xef, 0xf7, 0x5d, 0x51, 0xa0, 0x52,
	0xc9, 0x3b, 0x32, 0x83, 0xe9, 0x6f, 0x76, 0xfa, 0x55, 0x35, 0x95, 0x4e, 0x1c, 0x43, 0x14, 0x05,
	0x32, 0x8d, 0x65, 0x32, 0x21, 0x1f, 0x00, 0xfe, 0x70, 0xbd, 0x3d, 0x55, 0xca, 0xb0, 0x4b, 0x3e,
	0x42, 0xf4, 0x7d, 0xb3, 0x3d, 0x15, 0x28, 0x74, 0xc5, 0xdb, 0xc4, 0x23, 0x21, 0xf8, 0xdb, 0x46,
	0xe8, 0x3e, 0xf1, 0x37, 0xf1, 0xbf, 0x68, 0xf4, 0x0f, 0x1e, 0x02, 0x0b, 0x5f, 0x5f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x7c, 0x88, 0x0e, 0x97, 0x99, 0x02, 0x00, 0x00,
}
