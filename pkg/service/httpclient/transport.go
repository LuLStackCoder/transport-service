//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"bufio"
	"context"
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/LuLStackCoder/test-service/pkg/models"
)

type errorCreator func(status int, format string, v ...interface{}) error

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
	Decode(r *fasthttp.Response) error
}

// GetUserClientTransport transport interface
type GetUserClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.Request) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error)
}

type getUserClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getUserClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.Request) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	r.URI().QueryArgs().Add("id", string(request.Id))
	return
}

// DecodeResponse method for decoding response on client side
func (t *getUserClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetUserClientTransport the transport creator for http requests
func NewGetUserClientTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) GetUserClientTransport {
	return &getUserClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

// PostOrderClientTransport transport interface
type PostOrderClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.Request) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error)
}

type postOrderClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *postOrderClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.Request) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	r.SetBodyStreamWriter(func(w *bufio.Writer) {
		if err = json.NewEncoder(w).Encode(request); err != nil {
			return
		}
	})
	return
}

// DecodeResponse method for decoding response on client side
func (t *postOrderClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewPostOrderClientTransport the transport creator for http requests
func NewPostOrderClientTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) PostOrderClientTransport {
	return &postOrderClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

// GetCountClientTransport transport interface
type GetCountClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.Request) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error)
}

type getCountClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getCountClientTransport) EncodeRequest(ctx fasthttp.RequestCtx, r *fasthttp.Request, request *models.Request) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	ctx.SetUserValue("id", string(request.Id))
	return
}

// DecodeResponse method for decoding response on client side
func (t *getCountClientTransport) DecodeResponse(ctx fasthttp.RequestCtx, r *fasthttp.Response) (response models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetCountClientTransport the transport creator for http requests
func NewGetCountClientTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) GetCountClientTransport {
	return &getCountClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

// GetOrderClientTransport transport interface
type GetOrderClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.Request) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error)
}

type getOrderClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getOrderClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.Request) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	return
}

// DecodeResponse method for decoding response on client side
func (t *getOrderClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetOrderClientTransport the transport creator for http requests
func NewGetOrderClientTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) GetOrderClientTransport {
	return &getOrderClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
