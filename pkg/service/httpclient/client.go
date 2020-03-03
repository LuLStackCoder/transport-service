//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	`github.com/LuLStackCoder/test-service/pkg/models`
)

var (
	GetUser   = option{}
	PostOrder = option{}
	GetCount  = option{}
	GetOrder  = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service implements Service interface
type Service interface {
	GetUser(ctx context.Context, request *models.Request) (response models.Response, err error)
	PostOrder(ctx context.Context, request *models.Request) (response models.Response, err error)
	GetCount(ctx context.Context, request *models.Request) (response models.Response, err error)
	GetOrder(ctx context.Context, request *models.Request) (response models.Response, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetUser   GetUserClientTransport
	transportPostOrder PostOrderClientTransport
	transportGetCount  GetCountClientTransport
	transportGetOrder  GetOrderClientTransport
	options            map[interface{}]Option
}

// GetUser ...
func (s *client) GetUser(ctx context.Context, request *models.Request) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetUser]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetUser.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetUser.DecodeResponse(ctx, res)
}

// PostOrder ...
func (s *client) PostOrder(ctx context.Context, request *models.Request) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[PostOrder]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportPostOrder.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportPostOrder.DecodeResponse(ctx, res)
}

// GetCount ...
func (s *client) GetCount(ctx context.Context, request *models.Request) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetCount]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetCount.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetCount.DecodeResponse(ctx, res)
}

// GetOrder ...
func (s *client) GetOrder(ctx context.Context, request *models.Request) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetOrder]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetOrder.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetOrder.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,

	transportGetUser GetUserClientTransport,
	transportPostOrder PostOrderClientTransport,
	transportGetCount GetCountClientTransport,
	transportGetOrder GetOrderClientTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli: cli,

		transportGetUser:   transportGetUser,
		transportPostOrder: transportPostOrder,
		transportGetCount:  transportGetCount,
		transportGetOrder:  transportGetOrder,
		options:            options,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string,
	serverHost string,
	maxConns int,
	options map[interface{}]Option,
	errorProcessor errorProcessor,
	errorCreator errorCreator,

	uriPathGetUser string,
	uriPathPostOrder string,
	uriPathGetCount string,
	uriPathGetOrder string,

	httpMethodGetUser string,
	httpMethodPostOrder string,
	httpMethodGetCount string,
	httpMethodGetOrder string,
) Service {

	transportGetUser := NewGetUserClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetUser,
		httpMethodGetUser,
	)

	transportPostOrder := NewPostOrderClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathPostOrder,
		httpMethodPostOrder,
	)

	transportGetCount := NewGetCountClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetCount,
		httpMethodGetCount,
	)

	transportGetOrder := NewGetOrderClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetOrder,
		httpMethodGetOrder,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},

		transportGetUser,
		transportPostOrder,
		transportGetCount,
		transportGetOrder,
		options,
	)
}
