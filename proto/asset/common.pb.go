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

type PairInt struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                uint32   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Count                uint32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairInt) Reset()         { *m = PairInt{} }
func (m *PairInt) String() string { return proto.CompactTextString(m) }
func (*PairInt) ProtoMessage()    {}
func (*PairInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{0}
}

func (m *PairInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairInt.Unmarshal(m, b)
}
func (m *PairInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairInt.Marshal(b, m, deterministic)
}
func (m *PairInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairInt.Merge(m, src)
}
func (m *PairInt) XXX_Size() int {
	return xxx_messageInfo_PairInt.Size(m)
}
func (m *PairInt) XXX_DiscardUnknown() {
	xxx_messageInfo_PairInt.DiscardUnknown(m)
}

var xxx_messageInfo_PairInt proto.InternalMessageInfo

func (m *PairInt) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PairInt) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PairInt) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
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
	return fileDescriptor_3e70cb96243de9f1, []int{1}
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
	Uid                  string   `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{2}
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

func (m *RequestFilter) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type RequestList struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []string `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{3}
}

func (m *RequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestList.Unmarshal(m, b)
}
func (m *RequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestList.Marshal(b, m, deterministic)
}
func (m *RequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestList.Merge(m, src)
}
func (m *RequestList) XXX_Size() int {
	return xxx_messageInfo_RequestList.Size(m)
}
func (m *RequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestList proto.InternalMessageInfo

func (m *RequestList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestList) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
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
	return fileDescriptor_3e70cb96243de9f1, []int{4}
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
	return fileDescriptor_3e70cb96243de9f1, []int{5}
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
	return fileDescriptor_3e70cb96243de9f1, []int{6}
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

type ReplyStatistic struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Key                  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Owner                string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Count                uint32       `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	List                 []*PairInt   `protobuf:"bytes,10,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStatistic) Reset()         { *m = ReplyStatistic{} }
func (m *ReplyStatistic) String() string { return proto.CompactTextString(m) }
func (*ReplyStatistic) ProtoMessage()    {}
func (*ReplyStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e70cb96243de9f1, []int{7}
}

func (m *ReplyStatistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatistic.Unmarshal(m, b)
}
func (m *ReplyStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatistic.Marshal(b, m, deterministic)
}
func (m *ReplyStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatistic.Merge(m, src)
}
func (m *ReplyStatistic) XXX_Size() int {
	return xxx_messageInfo_ReplyStatistic.Size(m)
}
func (m *ReplyStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatistic proto.InternalMessageInfo

func (m *ReplyStatistic) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyStatistic) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyStatistic) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReplyStatistic) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReplyStatistic) GetList() []*PairInt {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("omo.msp.asset.ResultStatus", ResultStatus_name, ResultStatus_value)
	proto.RegisterType((*PairInt)(nil), "omo.msp.asset.PairInt")
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.asset.RequestInfo")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.asset.RequestFilter")
	proto.RegisterType((*RequestList)(nil), "omo.msp.asset.RequestList")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.asset.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.asset.ReplyInfo")
	proto.RegisterType((*RequestUpdate)(nil), "omo.msp.asset.RequestUpdate")
	proto.RegisterType((*ReplyStatistic)(nil), "omo.msp.asset.ReplyStatistic")
}

func init() {
	proto.RegisterFile("proto/asset/common.proto", fileDescriptor_3e70cb96243de9f1)
}

var fileDescriptor_3e70cb96243de9f1 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0x9b, 0x40,
	0x10, 0x2d, 0xe6, 0x23, 0x61, 0x08, 0x29, 0x5a, 0x55, 0x16, 0xca, 0xc9, 0xe2, 0x64, 0xe5, 0x40,
	0x24, 0xf7, 0xd0, 0x7b, 0x54, 0xb7, 0x8a, 0x94, 0x7e, 0x64, 0xa3, 0x5e, 0x7a, 0x23, 0x30, 0xa9,
	0x56, 0x05, 0x96, 0xb2, 0x4b, 0x6b, 0xff, 0x86, 0xfe, 0x80, 0xfe, 0x8b, 0xfe, 0xc6, 0x6a, 0x07,
	0x8c, 0x71, 0xe4, 0x48, 0xbd, 0xcd, 0x1b, 0x76, 0xdf, 0xcc, 0x7b, 0x6f, 0x81, 0xb8, 0x69, 0xa5,
	0x96, 0x57, 0x99, 0x52, 0xa8, 0xaf, 0x72, 0x59, 0x55, 0xb2, 0x4e, 0xa9, 0xc5, 0x42, 0x59, 0xc9,
	0xb4, 0x52, 0x4d, 0x4a, 0xdf, 0x92, 0xf7, 0x70, 0xf2, 0x39, 0x13, 0xed, 0x4d, 0xad, 0x59, 0x04,
	0xf6, 0x77, 0xdc, 0xc6, 0xd6, 0xc2, 0x5a, 0xfa, 0xdc, 0x94, 0xec, 0x15, 0xb8, 0x3f, 0xb3, 0xb2,
	0xc3, 0x78, 0xb6, 0xb0, 0x96, 0x21, 0xef, 0x81, 0xe9, 0xe6, 0xb2, 0xab, 0x75, 0x6c, 0xf7, 0x5d,
	0x02, 0xc9, 0x1d, 0x04, 0x1c, 0x7f, 0x74, 0xa8, 0xf4, 0x4d, 0xfd, 0x28, 0x0d, 0x59, 0x27, 0x8a,
	0x1d, 0x59, 0x27, 0x0a, 0x73, 0x4d, 0xfe, 0xaa, 0xb1, 0x25, 0x32, 0x9f, 0xf7, 0x80, 0x5d, 0xc0,
	0xa9, 0x6c, 0xb0, 0xcd, 0xb4, 0x6c, 0x89, 0xcf, 0xe7, 0x23, 0x4e, 0x7e, 0x5b, 0x10, 0x0e, 0x9c,
	0xef, 0x44, 0xa9, 0xb1, 0xdd, 0x73, 0x58, 0x53, 0x8e, 0x61, 0xf1, 0xd9, 0x91, 0xc5, 0x7b, 0xca,
	0x61, 0x71, 0x06, 0x4e, 0x93, 0x7d, 0xc3, 0xd8, 0xa1, 0xbd, 0xa9, 0x66, 0x73, 0xf0, 0xea, 0xae,
	0x7a, 0xc0, 0x36, 0x76, 0xa9, 0x3b, 0xa0, 0xdd, 0xfe, 0xde, 0xb8, 0x7f, 0xf2, 0x69, 0x14, 0x78,
	0x2b, 0x94, 0x3e, 0x22, 0x70, 0x2a, 0x65, 0x76, 0x28, 0xc5, 0x8c, 0x2e, 0x85, 0x32, 0x96, 0xd9,
	0x4b, 0x9f, 0x53, 0x9d, 0xbc, 0x31, 0x84, 0x4d, 0xb9, 0xbd, 0xd7, 0x99, 0xee, 0x94, 0x39, 0x92,
	0xcb, 0x02, 0x89, 0x31, 0xe4, 0x54, 0x1b, 0x1d, 0xd8, 0xb6, 0x23, 0x5f, 0x0f, 0x92, 0x3b, 0xf0,
	0xe9, 0xe2, 0x33, 0x46, 0xaf, 0xc0, 0x53, 0x44, 0x49, 0xb7, 0x82, 0xd5, 0x45, 0x7a, 0x10, 0x79,
	0x3a, 0x19, 0xca, 0x87, 0x93, 0xc9, 0x9f, 0xbd, 0xd5, 0x5f, 0x9a, 0x22, 0xd3, 0xf8, 0xbc, 0xd5,
	0x66, 0xda, 0xec, 0x20, 0xd6, 0x47, 0x81, 0x65, 0xb1, 0xb3, 0x9a, 0xc0, 0x3e, 0x00, 0x67, 0x1a,
	0xc0, 0xd4, 0x21, 0xf7, 0x89, 0x43, 0x73, 0xf0, 0xe8, 0x90, 0x8a, 0x3d, 0xf2, 0x68, 0x40, 0xc9,
	0x5f, 0x0b, 0xce, 0xc7, 0x8d, 0x85, 0xd2, 0x22, 0x9f, 0x08, 0xb4, 0xfe, 0x57, 0xe0, 0xf1, 0x37,
	0xd2, 0x0b, 0xb4, 0xa7, 0x02, 0xc7, 0xc7, 0xed, 0x4c, 0x1e, 0x37, 0xbb, 0x1c, 0xe2, 0x83, 0x85,
	0xbd, 0x0c, 0x56, 0xf3, 0x27, 0xf3, 0x86, 0x1f, 0xa8, 0x8f, 0xf5, 0x32, 0x87, 0x33, 0x8e, 0xaa,
	0x2b, 0xf5, 0x90, 0x6b, 0x00, 0x27, 0xf7, 0x5d, 0x9e, 0xa3, 0x52, 0xd1, 0x0b, 0x76, 0x06, 0xa7,
	0x1f, 0xb2, 0xcd, 0xad, 0xa8, 0x84, 0x8e, 0x2c, 0x83, 0x38, 0x36, 0x98, 0x69, 0x2c, 0xa2, 0x19,
	0x3b, 0x07, 0xf8, 0x28, 0xf5, 0x7a, 0x23, 0x94, 0xc1, 0x36, 0x7b, 0x09, 0xc1, 0xdb, 0xeb, 0xf5,
	0x26, 0xc7, 0x46, 0x0b, 0x59, 0x47, 0x0e, 0xf3, 0xc1, 0x5d, 0x57, 0x8d, 0xde, 0x46, 0xee, 0x75,
	0xf8, 0x35, 0x98, 0xfc, 0xe1, 0x0f, 0x1e, 0x81, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0x81, 0x63, 0x5e, 0xf7, 0x03, 0x00, 0x00,
}
