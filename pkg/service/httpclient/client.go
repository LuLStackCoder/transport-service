//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	"github.com/LuLStackCoder/test-service/pkg/models"
)

const (
	method = "http://"

	URIPrefix = "/api/v1"

	URIPathClientGetUser = URIPrefix + "/user"
	URIPathClientPostOrder = URIPrefix + "/orders"
	URIPathClientGetUserCount = URIPrefix + "/user/%s/count"
	URIPathClientGetOrders = URIPrefix + "/orders"

	HTTPMethodGetUser = "GET"
	HTTPMethodPostOrder = "POST"
	HTTPMethodGetUserCount = "GET"
	HTTPMethodGetOrders = "GET"
)

var (

	GetUser = option{}
	PostOrder = option{}
	GetUserCount = option{}
	GetOrders = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service implements Service interface
type Service interface {
	GetUser(ctx context.Context, request *models.Request) (res models.Response, err error)
	PostOrder(ctx context.Context, request *models.Request) (res models.Response, err error)
	GetCount(ctx context.Context, request *models.Request) (res models.Response, err error)
	GetOrder(ctx context.Context) (res models.Response, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetUser GetUserClientTransport
	transportPostOrder PostOrderClientTransport
	transportGetCount GetCountClientTransport
	transportGetOrder GetOrderClientTransport
	options map[interface{}]Option
}

// GetUser ...
func (s *client) GetUser(ctx context.Context, request *models.Request) (res models.Response, err error) {
	req, ress := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(ress)
	}()
	if opt, ok := s.options[GetUser]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetUser.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, ress)
	if err != nil {
		return
	}
	return s.transportGetUser.DecodeResponse(ctx, ress)
}

// PostOrder ...
func (s *client) PostOrder(ctx context.Context, request *models.Request) (res models.Response, err error) {
	req, ress := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(ress)
	}()
	if opt, ok := s.options[PostOrder]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportPostOrder.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, ress)
	if err != nil {
		return
	}
	return s.transportPostOrder.DecodeResponse(ctx, ress)
}

// GetUserCount ...
func (s *client) GetCount(ctx context.Context, request *models.Request) (res models.Response, err error) {
	req, ress := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(ress)
	}()
	if opt, ok := s.options[GetUserCount]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetCount.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, ress)
	if err != nil {
		return
	}
	return s.transportGetCount.DecodeResponse(ctx, ress)
}

// GetOrders ...
func (s *client) GetOrder(ctx context.Context) (res models.Response, err error) {
	req, ress := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(ress)
	}()
	if opt, ok := s.options[GetOrders]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetOrder.EncodeRequest(ctx, req); err != nil {
		return
	}
	err = s.cli.Do(req, ress)
	if err != nil {
		return
	}
	return s.transportGetOrder.DecodeResponse(ctx, ress)
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

		transportGetUser: transportGetUser,
		transportPostOrder: transportPostOrder,
		transportGetCount: transportGetCount,
		transportGetOrder: transportGetOrder,
		options: options,
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
	uriPathGetUserCount string,
	uriPathGetOrders string,

	httpMethodGetUser string,
	httpMethodPostOrder string,
	httpMethodGetUserCount string,
	httpMethodGetOrders string,
) Service {
	serverURL = method + serverURL
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

	transportGetUserCount := NewGetCountClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetUserCount,
		httpMethodGetUserCount,
	)

	transportGetOrders := NewGetOrderClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetOrders,
		httpMethodGetOrders,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},

		transportGetUser,
		transportPostOrder,
		transportGetUserCount,
		transportGetOrders,
		options,
	)
}
