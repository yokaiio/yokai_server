// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: server/gate/gate.proto

package gate

import (
	fmt "fmt"
	_ "github.com/east-eden/server/proto/global"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GateService service

func NewGateServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GateService service

type GateService interface {
	GetGateStatus(ctx context.Context, in *GateEmptyMessage, opts ...client.CallOption) (*GetGateStatusReply, error)
	SyncPlayerInfo(ctx context.Context, in *SyncPlayerInfoRequest, opts ...client.CallOption) (*SyncPlayerInfoReply, error)
}

type gateService struct {
	c    client.Client
	name string
}

func NewGateService(name string, c client.Client) GateService {
	return &gateService{
		c:    c,
		name: name,
	}
}

func (c *gateService) GetGateStatus(ctx context.Context, in *GateEmptyMessage, opts ...client.CallOption) (*GetGateStatusReply, error) {
	req := c.c.NewRequest(c.name, "GateService.GetGateStatus", in)
	out := new(GetGateStatusReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateService) SyncPlayerInfo(ctx context.Context, in *SyncPlayerInfoRequest, opts ...client.CallOption) (*SyncPlayerInfoReply, error) {
	req := c.c.NewRequest(c.name, "GateService.SyncPlayerInfo", in)
	out := new(SyncPlayerInfoReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GateService service

type GateServiceHandler interface {
	GetGateStatus(context.Context, *GateEmptyMessage, *GetGateStatusReply) error
	SyncPlayerInfo(context.Context, *SyncPlayerInfoRequest, *SyncPlayerInfoReply) error
}

func RegisterGateServiceHandler(s server.Server, hdlr GateServiceHandler, opts ...server.HandlerOption) error {
	type gateService interface {
		GetGateStatus(ctx context.Context, in *GateEmptyMessage, out *GetGateStatusReply) error
		SyncPlayerInfo(ctx context.Context, in *SyncPlayerInfoRequest, out *SyncPlayerInfoReply) error
	}
	type GateService struct {
		gateService
	}
	h := &gateServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GateService{h}, opts...))
}

type gateServiceHandler struct {
	GateServiceHandler
}

func (h *gateServiceHandler) GetGateStatus(ctx context.Context, in *GateEmptyMessage, out *GetGateStatusReply) error {
	return h.GateServiceHandler.GetGateStatus(ctx, in, out)
}

func (h *gateServiceHandler) SyncPlayerInfo(ctx context.Context, in *SyncPlayerInfoRequest, out *SyncPlayerInfoReply) error {
	return h.GateServiceHandler.SyncPlayerInfo(ctx, in, out)
}
