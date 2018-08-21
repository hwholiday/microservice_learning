// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: db_agent_server.proto

package dbagent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DbAgentServer service

type DbAgentServerService interface {
	GetOneTestUser(ctx context.Context, in *StringValue, opts ...client.CallOption) (*TestUser, error)
	GetAllTestUser(ctx context.Context, in *StringValue, opts ...client.CallOption) (*ListUser, error)
}

type dbAgentServerService struct {
	c    client.Client
	name string
}

func NewDbAgentServerService(name string, c client.Client) DbAgentServerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dbagent"
	}
	return &dbAgentServerService{
		c:    c,
		name: name,
	}
}

func (c *dbAgentServerService) GetOneTestUser(ctx context.Context, in *StringValue, opts ...client.CallOption) (*TestUser, error) {
	req := c.c.NewRequest(c.name, "DbAgentServer.GetOneTestUser", in)
	out := new(TestUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbAgentServerService) GetAllTestUser(ctx context.Context, in *StringValue, opts ...client.CallOption) (*ListUser, error) {
	req := c.c.NewRequest(c.name, "DbAgentServer.GetAllTestUser", in)
	out := new(ListUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DbAgentServer service

type DbAgentServerHandler interface {
	GetOneTestUser(context.Context, *StringValue, *TestUser) error
	GetAllTestUser(context.Context, *StringValue, *ListUser) error
}

func RegisterDbAgentServerHandler(s server.Server, hdlr DbAgentServerHandler, opts ...server.HandlerOption) error {
	type dbAgentServer interface {
		GetOneTestUser(ctx context.Context, in *StringValue, out *TestUser) error
		GetAllTestUser(ctx context.Context, in *StringValue, out *ListUser) error
	}
	type DbAgentServer struct {
		dbAgentServer
	}
	h := &dbAgentServerHandler{hdlr}
	return s.Handle(s.NewHandler(&DbAgentServer{h}, opts...))
}

type dbAgentServerHandler struct {
	DbAgentServerHandler
}

func (h *dbAgentServerHandler) GetOneTestUser(ctx context.Context, in *StringValue, out *TestUser) error {
	return h.DbAgentServerHandler.GetOneTestUser(ctx, in, out)
}

func (h *dbAgentServerHandler) GetAllTestUser(ctx context.Context, in *StringValue, out *ListUser) error {
	return h.DbAgentServerHandler.GetAllTestUser(ctx, in, out)
}