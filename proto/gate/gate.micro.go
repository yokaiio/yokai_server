// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: gate/gate.proto

package gate

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for GateService service

type GateService interface {
	GetGateStatus(ctx context.Context, in *GetGateStatusRequest, opts ...client.CallOption) (*GetGateStatusReply, error)
}

type gateService struct {
	c    client.Client
	name string
}

func NewGateService(name string, c client.Client) GateService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "yokai_gate"
	}
	return &gateService{
		c:    c,
		name: name,
	}
}

func (c *gateService) GetGateStatus(ctx context.Context, in *GetGateStatusRequest, opts ...client.CallOption) (*GetGateStatusReply, error) {
	req := c.c.NewRequest(c.name, "GateService.GetGateStatus", in)
	out := new(GetGateStatusReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GateService service

type GateServiceHandler interface {
	GetGateStatus(context.Context, *GetGateStatusRequest, *GetGateStatusReply) error
}

func RegisterGateServiceHandler(s server.Server, hdlr GateServiceHandler, opts ...server.HandlerOption) error {
	type gateService interface {
		GetGateStatus(ctx context.Context, in *GetGateStatusRequest, out *GetGateStatusReply) error
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

func (h *gateServiceHandler) GetGateStatus(ctx context.Context, in *GetGateStatusRequest, out *GetGateStatusReply) error {
	return h.GateServiceHandler.GetGateStatus(ctx, in, out)
}
