// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: server/game/game.proto

package game

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

// Api Endpoints for GameService service

func NewGameServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GameService service

type GameService interface {
	GetRemotePlayerInfo(ctx context.Context, in *GetRemotePlayerInfoRq, opts ...client.CallOption) (*GetRemotePlayerInfoRs, error)
	KickAccountOffline(ctx context.Context, in *KickAccountOfflineRq, opts ...client.CallOption) (*KickAccountOfflineRs, error)
	// test
	UpdatePlayerExp(ctx context.Context, in *UpdatePlayerExpRequest, opts ...client.CallOption) (*UpdatePlayerExpReply, error)
}

type gameService struct {
	c    client.Client
	name string
}

func NewGameService(name string, c client.Client) GameService {
	return &gameService{
		c:    c,
		name: name,
	}
}

func (c *gameService) GetRemotePlayerInfo(ctx context.Context, in *GetRemotePlayerInfoRq, opts ...client.CallOption) (*GetRemotePlayerInfoRs, error) {
	req := c.c.NewRequest(c.name, "GameService.GetRemotePlayerInfo", in)
	out := new(GetRemotePlayerInfoRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameService) KickAccountOffline(ctx context.Context, in *KickAccountOfflineRq, opts ...client.CallOption) (*KickAccountOfflineRs, error) {
	req := c.c.NewRequest(c.name, "GameService.KickAccountOffline", in)
	out := new(KickAccountOfflineRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameService) UpdatePlayerExp(ctx context.Context, in *UpdatePlayerExpRequest, opts ...client.CallOption) (*UpdatePlayerExpReply, error) {
	req := c.c.NewRequest(c.name, "GameService.UpdatePlayerExp", in)
	out := new(UpdatePlayerExpReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GameService service

type GameServiceHandler interface {
	GetRemotePlayerInfo(context.Context, *GetRemotePlayerInfoRq, *GetRemotePlayerInfoRs) error
	KickAccountOffline(context.Context, *KickAccountOfflineRq, *KickAccountOfflineRs) error
	// test
	UpdatePlayerExp(context.Context, *UpdatePlayerExpRequest, *UpdatePlayerExpReply) error
}

func RegisterGameServiceHandler(s server.Server, hdlr GameServiceHandler, opts ...server.HandlerOption) error {
	type gameService interface {
		GetRemotePlayerInfo(ctx context.Context, in *GetRemotePlayerInfoRq, out *GetRemotePlayerInfoRs) error
		KickAccountOffline(ctx context.Context, in *KickAccountOfflineRq, out *KickAccountOfflineRs) error
		UpdatePlayerExp(ctx context.Context, in *UpdatePlayerExpRequest, out *UpdatePlayerExpReply) error
	}
	type GameService struct {
		gameService
	}
	h := &gameServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GameService{h}, opts...))
}

type gameServiceHandler struct {
	GameServiceHandler
}

func (h *gameServiceHandler) GetRemotePlayerInfo(ctx context.Context, in *GetRemotePlayerInfoRq, out *GetRemotePlayerInfoRs) error {
	return h.GameServiceHandler.GetRemotePlayerInfo(ctx, in, out)
}

func (h *gameServiceHandler) KickAccountOffline(ctx context.Context, in *KickAccountOfflineRq, out *KickAccountOfflineRs) error {
	return h.GameServiceHandler.KickAccountOffline(ctx, in, out)
}

func (h *gameServiceHandler) UpdatePlayerExp(ctx context.Context, in *UpdatePlayerExpRequest, out *UpdatePlayerExpReply) error {
	return h.GameServiceHandler.UpdatePlayerExp(ctx, in, out)
}
