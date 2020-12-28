// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/asset/thumb.proto

package asset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ThumbService service

type ThumbService interface {
	AddOne(ctx context.Context, in *ReqThumbAdd, opts ...client.CallOption) (*ReplyThumbOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyThumbOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetList(ctx context.Context, in *ReqThumbList, opts ...client.CallOption) (*ReplyThumbList, error)
	GetByOwner(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyThumbList, error)
	UpdateBase(ctx context.Context, in *ReqThumbUpdate, opts ...client.CallOption) (*ReplyThumbOne, error)
}

type thumbService struct {
	c    client.Client
	name string
}

func NewThumbService(name string, c client.Client) ThumbService {
	return &thumbService{
		c:    c,
		name: name,
	}
}

func (c *thumbService) AddOne(ctx context.Context, in *ReqThumbAdd, opts ...client.CallOption) (*ReplyThumbOne, error) {
	req := c.c.NewRequest(c.name, "ThumbService.AddOne", in)
	out := new(ReplyThumbOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thumbService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyThumbOne, error) {
	req := c.c.NewRequest(c.name, "ThumbService.GetOne", in)
	out := new(ReplyThumbOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thumbService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ThumbService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thumbService) GetList(ctx context.Context, in *ReqThumbList, opts ...client.CallOption) (*ReplyThumbList, error) {
	req := c.c.NewRequest(c.name, "ThumbService.GetList", in)
	out := new(ReplyThumbList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thumbService) GetByOwner(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyThumbList, error) {
	req := c.c.NewRequest(c.name, "ThumbService.GetByOwner", in)
	out := new(ReplyThumbList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thumbService) UpdateBase(ctx context.Context, in *ReqThumbUpdate, opts ...client.CallOption) (*ReplyThumbOne, error) {
	req := c.c.NewRequest(c.name, "ThumbService.UpdateBase", in)
	out := new(ReplyThumbOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ThumbService service

type ThumbServiceHandler interface {
	AddOne(context.Context, *ReqThumbAdd, *ReplyThumbOne) error
	GetOne(context.Context, *RequestInfo, *ReplyThumbOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetList(context.Context, *ReqThumbList, *ReplyThumbList) error
	GetByOwner(context.Context, *RequestInfo, *ReplyThumbList) error
	UpdateBase(context.Context, *ReqThumbUpdate, *ReplyThumbOne) error
}

func RegisterThumbServiceHandler(s server.Server, hdlr ThumbServiceHandler, opts ...server.HandlerOption) error {
	type thumbService interface {
		AddOne(ctx context.Context, in *ReqThumbAdd, out *ReplyThumbOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyThumbOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetList(ctx context.Context, in *ReqThumbList, out *ReplyThumbList) error
		GetByOwner(ctx context.Context, in *RequestInfo, out *ReplyThumbList) error
		UpdateBase(ctx context.Context, in *ReqThumbUpdate, out *ReplyThumbOne) error
	}
	type ThumbService struct {
		thumbService
	}
	h := &thumbServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ThumbService{h}, opts...))
}

type thumbServiceHandler struct {
	ThumbServiceHandler
}

func (h *thumbServiceHandler) AddOne(ctx context.Context, in *ReqThumbAdd, out *ReplyThumbOne) error {
	return h.ThumbServiceHandler.AddOne(ctx, in, out)
}

func (h *thumbServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyThumbOne) error {
	return h.ThumbServiceHandler.GetOne(ctx, in, out)
}

func (h *thumbServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.ThumbServiceHandler.RemoveOne(ctx, in, out)
}

func (h *thumbServiceHandler) GetList(ctx context.Context, in *ReqThumbList, out *ReplyThumbList) error {
	return h.ThumbServiceHandler.GetList(ctx, in, out)
}

func (h *thumbServiceHandler) GetByOwner(ctx context.Context, in *RequestInfo, out *ReplyThumbList) error {
	return h.ThumbServiceHandler.GetByOwner(ctx, in, out)
}

func (h *thumbServiceHandler) UpdateBase(ctx context.Context, in *ReqThumbUpdate, out *ReplyThumbOne) error {
	return h.ThumbServiceHandler.UpdateBase(ctx, in, out)
}
