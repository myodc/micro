// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cmd/protoc-gen-micro/examples/greeter/greeter.proto

package greeter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
	microClient "github.com/micro/micro/v3/service/client"
	microServer "github.com/micro/micro/v3/service/server"
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
var _ = microServer.Handle
var _ = microClient.Call

// Api Endpoints for Greeter service

func NewGreeterEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Greeter.Hello",
			Path:    []string{"/hello"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Greeter.Stream",
			Path:    []string{"/stream"},
			Method:  []string{"GET"},
			Stream:  true,
			Handler: "rpc",
		},
	}
}

// Client API for Greeter service

type GreeterService interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, opts ...client.CallOption) (Greeter_StreamService, error)
}

type greeterService struct {
	name string
}

func NewGreeterService(name string) GreeterService {
	return &greeterService{name: name}
}

func (c *greeterService) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := microClient.NewRequest(c.name, "Greeter.Hello", in)
	out := new(Response)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterService) Stream(ctx context.Context, opts ...client.CallOption) (Greeter_StreamService, error) {
	req := microClient.NewRequest(c.name, "Greeter.Stream", &Request{})
	stream, err := microClient.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &greeterServiceStream{stream}, nil
}

type Greeter_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type greeterServiceStream struct {
	stream client.Stream
}

func (x *greeterServiceStream) Close() error {
	return x.stream.Close()
}

func (x *greeterServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *greeterServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greeterServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greeterServiceStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *greeterServiceStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Hello(context.Context, *Request, *Response) error
	Stream(context.Context, Greeter_StreamStream) error
}

func RegisterGreeterHandler(hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Hello(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Greeter.Hello",
		Path:    []string{"/hello"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Greeter.Stream",
		Path:    []string{"/stream"},
		Method:  []string{"GET"},
		Stream:  true,
		Handler: "rpc",
	}))
	return microServer.Handle(microServer.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.GreeterHandler.Hello(ctx, in, out)
}

func (h *greeterHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.GreeterHandler.Stream(ctx, &greeterStreamStream{stream})
}

type Greeter_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type greeterStreamStream struct {
	stream server.Stream
}

func (x *greeterStreamStream) Close() error {
	return x.stream.Close()
}

func (x *greeterStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *greeterStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greeterStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greeterStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *greeterStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
