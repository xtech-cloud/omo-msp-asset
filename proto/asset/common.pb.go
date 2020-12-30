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
	return fileDescriptor_3e70cb96243de9f1, []int{1}
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
	return fileDescriptor_3e70cb96243de9f1, []int{2}
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

func init() {
	proto.RegisterEnum("omo.msp.asset.ResultStatus", ResultStatus_name, ResultStatus_value)
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.asset.RequestInfo")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.asset.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.asset.ReplyInfo")
}

func init() {
	proto.RegisterFile("proto/asset/common.proto", fileDescriptor_3e70cb96243de9f1)
}

var fileDescriptor_3e70cb96243de9f1 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0x47, 0xe9, 0x9f, 0x94, 0xe6, 0xd2, 0x80, 0x65, 0x31, 0x44, 0x9d, 0x50, 0x26, 0xc4, 0x90,
	0x4a, 0x65, 0x60, 0xaf, 0xc8, 0x80, 0x04, 0x48, 0x75, 0x37, 0xb6, 0xe0, 0x1c, 0x52, 0xa4, 0x3a,
	0x67, 0xec, 0x8b, 0x48, 0xbf, 0x3d, 0x8a, 0xa9, 0xaa, 0x32, 0xb0, 0xdd, 0xbb, 0xf3, 0x3d, 0xff,
	0x6c, 0xc8, 0xac, 0x23, 0xa6, 0x55, 0xe5, 0x3d, 0xf2, 0x4a, 0x93, 0x31, 0xd4, 0x16, 0xa1, 0x25,
	0x53, 0x32, 0x54, 0x18, 0x6f, 0x8b, 0x30, 0xcb, 0xb7, 0x90, 0x28, 0xfc, 0xea, 0xd0, 0xf3, 0x73,
	0xfb, 0x49, 0x52, 0xc0, 0xa4, 0x6b, 0xea, 0x6c, 0x74, 0x3b, 0xba, 0x8b, 0xd5, 0x50, 0xca, 0x1b,
	0x88, 0xe8, 0xbb, 0x45, 0x97, 0x8d, 0x43, 0xef, 0x17, 0xe4, 0x12, 0xe6, 0x64, 0xd1, 0x55, 0x4c,
	0x2e, 0x9b, 0x84, 0xc1, 0x89, 0xf3, 0xc7, 0x41, 0x69, 0xf7, 0x87, 0x1d, 0x57, 0xdc, 0x79, 0x29,
	0x61, 0xaa, 0xa9, 0xc6, 0xe0, 0x4c, 0x55, 0xa8, 0x07, 0x29, 0x3a, 0x47, 0x27, 0x69, 0x80, 0x7c,
	0x0b, 0x71, 0x58, 0xfc, 0x27, 0xc9, 0x1a, 0x66, 0x3e, 0x28, 0xc3, 0x56, 0xb2, 0x5e, 0x16, 0x7f,
	0x9e, 0x52, 0x9c, 0x5d, 0xaa, 0x8e, 0x27, 0xef, 0x35, 0x2c, 0x14, 0xfa, 0x6e, 0xcf, 0xc7, 0x30,
	0x09, 0x5c, 0xee, 0x3a, 0xad, 0xd1, 0x7b, 0x71, 0x21, 0x17, 0x30, 0x7f, 0xad, 0xfa, 0x97, 0xc6,
	0x34, 0x2c, 0x46, 0x03, 0x29, 0xb4, 0x58, 0x31, 0xd6, 0x62, 0x2c, 0xaf, 0x00, 0xde, 0x88, 0xcb,
	0xbe, 0xf1, 0x03, 0x4f, 0xe4, 0x35, 0x24, 0x4f, 0x9b, 0xb2, 0xd7, 0x68, 0xb9, 0xa1, 0x56, 0x4c,
	0x65, 0x0c, 0x51, 0x69, 0x2c, 0x1f, 0x44, 0xb4, 0x49, 0xdf, 0x93, 0xb3, 0xef, 0xfe, 0x98, 0x05,
	0x78, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x19, 0x38, 0xcf, 0xca, 0x84, 0x01, 0x00, 0x00,
}