// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/asset/thumb.proto

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

type ThumbInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Face                 string   `protobuf:"bytes,3,opt,name=face,proto3" json:"face,omitempty"`
	Probably             float32  `protobuf:"fixed32,4,opt,name=probably,proto3" json:"probably,omitempty"`
	Asset                string   `protobuf:"bytes,5,opt,name=asset,proto3" json:"asset,omitempty"`
	Url                  string   `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Created              int64    `protobuf:"varint,7,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,8,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,10,opt,name=operator,proto3" json:"operator,omitempty"`
	Similar              float32  `protobuf:"fixed32,11,opt,name=similar,proto3" json:"similar,omitempty"`
	Blur                 float32  `protobuf:"fixed32,12,opt,name=blur,proto3" json:"blur,omitempty"`
	Meta                 string   `protobuf:"bytes,13,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThumbInfo) Reset()         { *m = ThumbInfo{} }
func (m *ThumbInfo) String() string { return proto.CompactTextString(m) }
func (*ThumbInfo) ProtoMessage()    {}
func (*ThumbInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42b69ff0892ee7bc, []int{0}
}

func (m *ThumbInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThumbInfo.Unmarshal(m, b)
}
func (m *ThumbInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThumbInfo.Marshal(b, m, deterministic)
}
func (m *ThumbInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThumbInfo.Merge(m, src)
}
func (m *ThumbInfo) XXX_Size() int {
	return xxx_messageInfo_ThumbInfo.Size(m)
}
func (m *ThumbInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ThumbInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ThumbInfo proto.InternalMessageInfo

func (m *ThumbInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ThumbInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ThumbInfo) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *ThumbInfo) GetProbably() float32 {
	if m != nil {
		return m.Probably
	}
	return 0
}

func (m *ThumbInfo) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ThumbInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ThumbInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ThumbInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ThumbInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ThumbInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ThumbInfo) GetSimilar() float32 {
	if m != nil {
		return m.Similar
	}
	return 0
}

func (m *ThumbInfo) GetBlur() float32 {
	if m != nil {
		return m.Blur
	}
	return 0
}

func (m *ThumbInfo) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

type ReqThumbAdd struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Face                 string   `protobuf:"bytes,2,opt,name=face,proto3" json:"face,omitempty"`
	Asset                string   `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Probably             float32  `protobuf:"fixed32,5,opt,name=probably,proto3" json:"probably,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Similar              float32  `protobuf:"fixed32,7,opt,name=similar,proto3" json:"similar,omitempty"`
	Blur                 float32  `protobuf:"fixed32,8,opt,name=blur,proto3" json:"blur,omitempty"`
	Meta                 string   `protobuf:"bytes,9,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqThumbAdd) Reset()         { *m = ReqThumbAdd{} }
func (m *ReqThumbAdd) String() string { return proto.CompactTextString(m) }
func (*ReqThumbAdd) ProtoMessage()    {}
func (*ReqThumbAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_42b69ff0892ee7bc, []int{1}
}

func (m *ReqThumbAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqThumbAdd.Unmarshal(m, b)
}
func (m *ReqThumbAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqThumbAdd.Marshal(b, m, deterministic)
}
func (m *ReqThumbAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqThumbAdd.Merge(m, src)
}
func (m *ReqThumbAdd) XXX_Size() int {
	return xxx_messageInfo_ReqThumbAdd.Size(m)
}
func (m *ReqThumbAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqThumbAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqThumbAdd proto.InternalMessageInfo

func (m *ReqThumbAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqThumbAdd) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *ReqThumbAdd) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ReqThumbAdd) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ReqThumbAdd) GetProbably() float32 {
	if m != nil {
		return m.Probably
	}
	return 0
}

func (m *ReqThumbAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqThumbAdd) GetSimilar() float32 {
	if m != nil {
		return m.Similar
	}
	return 0
}

func (m *ReqThumbAdd) GetBlur() float32 {
	if m != nil {
		return m.Blur
	}
	return 0
}

func (m *ReqThumbAdd) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

type ReqThumbUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Similar              float32  `protobuf:"fixed32,3,opt,name=similar,proto3" json:"similar,omitempty"`
	Asset                string   `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqThumbUpdate) Reset()         { *m = ReqThumbUpdate{} }
func (m *ReqThumbUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqThumbUpdate) ProtoMessage()    {}
func (*ReqThumbUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_42b69ff0892ee7bc, []int{2}
}

func (m *ReqThumbUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqThumbUpdate.Unmarshal(m, b)
}
func (m *ReqThumbUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqThumbUpdate.Marshal(b, m, deterministic)
}
func (m *ReqThumbUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqThumbUpdate.Merge(m, src)
}
func (m *ReqThumbUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqThumbUpdate.Size(m)
}
func (m *ReqThumbUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqThumbUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqThumbUpdate proto.InternalMessageInfo

func (m *ReqThumbUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqThumbUpdate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqThumbUpdate) GetSimilar() float32 {
	if m != nil {
		return m.Similar
	}
	return 0
}

func (m *ReqThumbUpdate) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ReqThumbUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyThumbOne struct {
	Info                 *ThumbInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyThumbOne) Reset()         { *m = ReplyThumbOne{} }
func (m *ReplyThumbOne) String() string { return proto.CompactTextString(m) }
func (*ReplyThumbOne) ProtoMessage()    {}
func (*ReplyThumbOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_42b69ff0892ee7bc, []int{3}
}

func (m *ReplyThumbOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyThumbOne.Unmarshal(m, b)
}
func (m *ReplyThumbOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyThumbOne.Marshal(b, m, deterministic)
}
func (m *ReplyThumbOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyThumbOne.Merge(m, src)
}
func (m *ReplyThumbOne) XXX_Size() int {
	return xxx_messageInfo_ReplyThumbOne.Size(m)
}
func (m *ReplyThumbOne) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyThumbOne.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyThumbOne proto.InternalMessageInfo

func (m *ReplyThumbOne) GetInfo() *ThumbInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyThumbOne) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqThumbList struct {
	Asset                string   `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	List                 []string `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqThumbList) Reset()         { *m = ReqThumbList{} }
func (m *ReqThumbList) String() string { return proto.CompactTextString(m) }
func (*ReqThumbList) ProtoMessage()    {}
func (*ReqThumbList) Descriptor() ([]byte, []int) {
	return fileDescriptor_42b69ff0892ee7bc, []int{4}
}

func (m *ReqThumbList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqThumbList.Unmarshal(m, b)
}
func (m *ReqThumbList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqThumbList.Marshal(b, m, deterministic)
}
func (m *ReqThumbList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqThumbList.Merge(m, src)
}
func (m *ReqThumbList) XXX_Size() int {
	return xxx_messageInfo_ReqThumbList.Size(m)
}
func (m *ReqThumbList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqThumbList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqThumbList proto.InternalMessageInfo

func (m *ReqThumbList) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ReqThumbList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyThumbList struct {
	Owner                string       `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	List                 []*ThumbInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyThumbList) Reset()         { *m = ReplyThumbList{} }
func (m *ReplyThumbList) String() string { return proto.CompactTextString(m) }
func (*ReplyThumbList) ProtoMessage()    {}
func (*ReplyThumbList) Descriptor() ([]byte, []int) {
	return fileDescriptor_42b69ff0892ee7bc, []int{5}
}

func (m *ReplyThumbList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyThumbList.Unmarshal(m, b)
}
func (m *ReplyThumbList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyThumbList.Marshal(b, m, deterministic)
}
func (m *ReplyThumbList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyThumbList.Merge(m, src)
}
func (m *ReplyThumbList) XXX_Size() int {
	return xxx_messageInfo_ReplyThumbList.Size(m)
}
func (m *ReplyThumbList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyThumbList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyThumbList proto.InternalMessageInfo

func (m *ReplyThumbList) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReplyThumbList) GetList() []*ThumbInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyThumbList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ThumbInfo)(nil), "omo.msp.asset.ThumbInfo")
	proto.RegisterType((*ReqThumbAdd)(nil), "omo.msp.asset.ReqThumbAdd")
	proto.RegisterType((*ReqThumbUpdate)(nil), "omo.msp.asset.ReqThumbUpdate")
	proto.RegisterType((*ReplyThumbOne)(nil), "omo.msp.asset.ReplyThumbOne")
	proto.RegisterType((*ReqThumbList)(nil), "omo.msp.asset.ReqThumbList")
	proto.RegisterType((*ReplyThumbList)(nil), "omo.msp.asset.ReplyThumbList")
}

func init() {
	proto.RegisterFile("proto/asset/thumb.proto", fileDescriptor_42b69ff0892ee7bc)
}

var fileDescriptor_42b69ff0892ee7bc = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4f, 0x6b, 0xdb, 0x4e,
	0x10, 0xb5, 0x2c, 0xf9, 0xdf, 0xc8, 0x0e, 0x3f, 0x96, 0x1f, 0x74, 0x71, 0x13, 0x30, 0x3a, 0xf9,
	0x10, 0x1c, 0x70, 0x2f, 0xbd, 0xc6, 0x2d, 0x35, 0x69, 0x03, 0x01, 0xa5, 0xbd, 0xf4, 0x26, 0x5b,
	0x63, 0x2a, 0x90, 0xbc, 0xca, 0x6a, 0x95, 0xe2, 0x7b, 0x0f, 0xfd, 0x46, 0xfd, 0x2c, 0xed, 0xa7,
	0x29, 0x3b, 0xb2, 0x94, 0x75, 0x2a, 0x5b, 0xb9, 0xed, 0xec, 0x9b, 0x79, 0xfb, 0xde, 0xf3, 0x18,
	0xc1, 0xab, 0x54, 0x0a, 0x25, 0xae, 0x82, 0x2c, 0x43, 0x75, 0xa5, 0xbe, 0xe5, 0xc9, 0x6a, 0x46,
	0x37, 0x6c, 0x24, 0x12, 0x31, 0x4b, 0xb2, 0x74, 0x46, 0xd0, 0x98, 0x9b, 0x7d, 0x6b, 0x91, 0x24,
	0x62, 0x5b, 0x34, 0x7a, 0xbf, 0xda, 0x30, 0xf8, 0xac, 0x07, 0x6f, 0xb6, 0x1b, 0xc1, 0xfe, 0x03,
	0x3b, 0x8f, 0x42, 0x6e, 0x4d, 0xac, 0xe9, 0xc0, 0xd7, 0x47, 0xf6, 0x3f, 0x74, 0xc4, 0xf7, 0x2d,
	0x4a, 0xde, 0xa6, 0xbb, 0xa2, 0x60, 0x0c, 0x9c, 0x4d, 0xb0, 0x46, 0x6e, 0xd3, 0x25, 0x9d, 0xd9,
	0x18, 0xfa, 0xa9, 0x14, 0xab, 0x60, 0x15, 0xef, 0xb8, 0x33, 0xb1, 0xa6, 0x6d, 0xbf, 0xaa, 0x35,
	0x0b, 0xbd, 0xcd, 0x3b, 0x05, 0x0b, 0x15, 0xf4, 0x9a, 0x8c, 0x79, 0x77, 0xff, 0x9a, 0x8c, 0x19,
	0x87, 0xde, 0x5a, 0x62, 0xa0, 0x30, 0xe4, 0xbd, 0x89, 0x35, 0xb5, 0xfd, 0xb2, 0xd4, 0x48, 0x9e,
	0x86, 0x84, 0xf4, 0x0b, 0x64, 0x5f, 0x56, 0x33, 0x42, 0xf2, 0x01, 0x31, 0x95, 0xa5, 0x56, 0x24,
	0x52, 0x94, 0x04, 0x01, 0x41, 0x55, 0xad, 0xa7, 0xb2, 0x28, 0x89, 0xe2, 0x40, 0x72, 0x97, 0xc4,
	0x96, 0xa5, 0xf6, 0xb6, 0x8a, 0x73, 0xc9, 0x87, 0x74, 0x4d, 0x67, 0x7d, 0x97, 0xa0, 0x0a, 0xf8,
	0xa8, 0xf0, 0xab, 0xcf, 0xde, 0x6f, 0x0b, 0x5c, 0x1f, 0x1f, 0x28, 0xbc, 0xeb, 0xd0, 0x48, 0xca,
	0xaa, 0x4b, 0xaa, 0x6d, 0x24, 0x55, 0xa5, 0x61, 0xd7, 0xa4, 0xe1, 0x3c, 0xa5, 0x61, 0x26, 0xda,
	0x79, 0x96, 0xa8, 0xe9, 0xad, 0x7b, 0xdc, 0x5b, 0xaf, 0xde, 0x5b, 0xbf, 0xc6, 0xdb, 0xc0, 0xf0,
	0xf6, 0xc3, 0x82, 0xb3, 0xd2, 0xdb, 0x17, 0xca, 0xf9, 0xc5, 0xab, 0x61, 0x3c, 0x6e, 0x1f, 0x3e,
	0x5e, 0xd9, 0x76, 0x4c, 0xdb, 0xa6, 0x91, 0xce, 0xa1, 0x11, 0xef, 0x01, 0x46, 0x3e, 0xa6, 0xf1,
	0x8e, 0x74, 0xdc, 0x6d, 0x91, 0x5d, 0x82, 0x13, 0x6d, 0x37, 0x82, 0x54, 0xb8, 0x73, 0x3e, 0x3b,
	0xd8, 0xf2, 0x59, 0xb5, 0xc7, 0x3e, 0x75, 0xb1, 0x39, 0x74, 0x33, 0x15, 0xa8, 0x3c, 0x23, 0x85,
	0xee, 0x7c, 0xfc, 0xac, 0x9f, 0xb8, 0xef, 0xa9, 0xc3, 0xdf, 0x77, 0x7a, 0x6f, 0x61, 0x58, 0x1a,
	0xbf, 0x8d, 0x32, 0xf5, 0x24, 0xda, 0x32, 0x45, 0x33, 0x70, 0xe2, 0x28, 0x53, 0xbc, 0x3d, 0xb1,
	0x75, 0x66, 0xfa, 0xec, 0xfd, 0xa4, 0xcc, 0x4a, 0xb5, 0xe5, 0x70, 0xcd, 0x4a, 0x5c, 0x1a, 0xc3,
	0x27, 0x4d, 0xe8, 0x2e, 0xc3, 0x84, 0xfd, 0x52, 0x13, 0xf3, 0x3f, 0x0e, 0x0c, 0x89, 0xe7, 0x1e,
	0xe5, 0x63, 0xb4, 0x46, 0xf6, 0x1e, 0xba, 0xd7, 0x61, 0xa8, 0x13, 0xfc, 0x77, 0xbc, 0xda, 0xe0,
	0xf1, 0x79, 0x1d, 0x75, 0x99, 0xbd, 0xd7, 0xd2, 0x2c, 0x4b, 0x54, 0x47, 0x58, 0x72, 0xcc, 0x94,
	0x96, 0xdd, 0xc8, 0xf2, 0x0e, 0x06, 0x3e, 0x26, 0xe2, 0x11, 0x9b, 0x88, 0x78, 0x1d, 0x91, 0x46,
	0xbc, 0x16, 0x5b, 0x42, 0x6f, 0x89, 0x8a, 0x42, 0x7e, 0x7d, 0xc4, 0x91, 0x06, 0xc7, 0x17, 0x47,
	0xc5, 0x68, 0xd8, 0x6b, 0xb1, 0x1b, 0x80, 0x25, 0xaa, 0xc5, 0xee, 0x8e, 0x7e, 0x9a, 0x53, 0x72,
	0x1a, 0xa9, 0x3e, 0x01, 0x14, 0xff, 0x95, 0x45, 0x90, 0x21, 0xbb, 0x38, 0x22, 0xab, 0x68, 0x69,
	0x4c, 0xe9, 0x23, 0x9c, 0xed, 0xc9, 0x76, 0x1f, 0xa2, 0x58, 0xa1, 0x64, 0xe7, 0xf5, 0xda, 0xf6,
	0x7c, 0xa7, 0xc2, 0xba, 0x05, 0x97, 0x3c, 0x9e, 0x26, 0x2a, 0xd0, 0x46, 0x9b, 0x8b, 0xd1, 0x57,
	0xd7, 0xf8, 0x9a, 0xac, 0xba, 0x54, 0xbc, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x09, 0x69, 0x58,
	0x8f, 0x8b, 0x06, 0x00, 0x00,
}
