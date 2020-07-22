// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: test/dep-test/dep-test-service/proto/dep/dep.proto

package go_micro_service_dep

import (
	fmt "fmt"
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

// Api Endpoints for Dep service

func NewDepEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Dep service

type DepService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Dep_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Dep_PingPongService, error)
}

type depService struct {
	c    client.Client
	name string
}

func NewDepService(name string, c client.Client) DepService {
	return &depService{
		c:    c,
		name: name,
	}
}

func (c *depService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Dep.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Dep_StreamService, error) {
	req := c.c.NewRequest(c.name, "Dep.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &depServiceStream{stream}, nil
}

type Dep_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type depServiceStream struct {
	stream client.Stream
}

func (x *depServiceStream) Close() error {
	return x.stream.Close()
}

func (x *depServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *depServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *depServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *depServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *depService) PingPong(ctx context.Context, opts ...client.CallOption) (Dep_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Dep.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &depServicePingPong{stream}, nil
}

type Dep_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type depServicePingPong struct {
	stream client.Stream
}

func (x *depServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *depServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *depServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *depServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *depServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *depServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Dep service

type DepHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Dep_StreamStream) error
	PingPong(context.Context, Dep_PingPongStream) error
}

func RegisterDepHandler(s server.Server, hdlr DepHandler, opts ...server.HandlerOption) error {
	type dep interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Dep struct {
		dep
	}
	h := &depHandler{hdlr}
	return s.Handle(s.NewHandler(&Dep{h}, opts...))
}

type depHandler struct {
	DepHandler
}

func (h *depHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.DepHandler.Call(ctx, in, out)
}

func (h *depHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DepHandler.Stream(ctx, m, &depStreamStream{stream})
}

type Dep_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type depStreamStream struct {
	stream server.Stream
}

func (x *depStreamStream) Close() error {
	return x.stream.Close()
}

func (x *depStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *depStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *depStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *depStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *depHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.DepHandler.PingPong(ctx, &depPingPongStream{stream})
}

type Dep_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type depPingPongStream struct {
	stream server.Stream
}

func (x *depPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *depPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *depPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *depPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *depPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *depPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
