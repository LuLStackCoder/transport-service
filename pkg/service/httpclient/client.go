//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"
	`net/http`

	"github.com/valyala/fasthttp"

	`github.com/LuLStackCoder/test-service/pkg/models`
)

const (
	method = "http://"

	URIPathClientGetUser    = "/api/v1/user"
	URIPathClientPostOrder = "/api/v1/orders"
	URIPathClientGetCount      = "/api/v1/user/:id/count"
	URIPathClientGetOrder  = "/api/v1/orders"
)

type (
	beforeRequest func(req *fasthttp.Request) (err error)
	afterResponse func(resp *fasthttp.Response) (err error)
)


// Service implements Service interface
type Service interface {
	GetUser(ctx context.Context, request *models.Request) (response models.Response, err error)
	PostOrder(ctx context.Context, request *models.Request) (response models.Response, err error)
	GetCount(ctx fasthttp.RequestCtx, request *models.Request) (response models.Response, err error)
	GetOrder(ctx context.Context) (response models.Response, err error)
}

type client struct {
	cli *fasthttp.HostClient


	transportGetUser   GetUserClientTransport
	transportPostOrder PostOrderClientTransport
	transportGetCount  GetCountClientTransport
	transportGetOrder  GetOrderClientTransport
}

// GetUser ...
func (s *client) GetUser(ctx context.Context, request *models.Request) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
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
	defer func() {;
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
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
func (s *client) GetCount(ctx fasthttp.RequestCtx, request *models.Request) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
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
func (s *client) GetOrder(ctx context.Context) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if err = s.transportGetOrder.EncodeRequest(ctx, req,); err != nil {
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
) Service {
	return &client{
		cli: cli,

		transportGetUser:   transportGetUser,
		transportPostOrder: transportPostOrder,
		transportGetCount:  transportGetCount,
		transportGetOrder:  transportGetOrder,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string,
	maxConns int,
	errorProcessor errorProcessor,
	errorCreator errorCreator,
) Service {

	transportGetUser := NewGetUserClientTransport(
		errorProcessor,
		errorCreator,
		method+serverURL+URIPathClientGetUser,
		http.MethodGet,
	)

	transportPostOrder := NewPostOrderClientTransport(
		errorProcessor,
		errorCreator,
		method+serverURL+URIPathClientGetUser,
		http.MethodPost,
	)

	transportGetCount := NewGetCountClientTransport(
		errorProcessor,
		errorCreator,
		method+serverURL+URIPathClientGetUser,
		http.MethodGet,
	)

	transportGetOrder := NewGetOrderClientTransport(
		errorProcessor,
		errorCreator,
		method+serverURL+URIPathClientGetUser,
		http.MethodGet,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverURL,
			MaxConns: maxConns,
		},

		transportGetUser,
		transportPostOrder,
		transportGetCount,
		transportGetOrder,
	)
}
