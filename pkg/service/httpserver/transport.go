//Package service http transport
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	"net/http"
	`strconv`

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	`github.com/LuLStackCoder/test-service/pkg/models`
)

type errorCreator func(status int, format string, v ...interface{}) error

// GetUserTransport transport interface
type GetUserTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.Request, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error)
}

type getUserTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getUserTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.Request, err error) {
	request.Id, err = strconv.Atoi(string(r.URI().QueryArgs().Peek("id")))
	if err != nil {
		return request, t.errorCreator(
			http.StatusBadRequest,
			"Bad request, check the fields.",
			"failed to get Id from query: %v",
			err,
		)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getUserTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetUserTransport the transport creator for http requests
func NewGetUserTransport(errorCreator errorCreator) GetUserTransport {
	return &getUserTransport{
		errorCreator: errorCreator,
	}
}

// PostOrderTransport transport interface
type PostOrderTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.Request, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error)
}

type postOrderTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *postOrderTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.Request, err error) {
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		return models.Request{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *postOrderTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewPostOrderTransport the transport creator for http requests
func NewPostOrderTransport(errorCreator errorCreator) PostOrderTransport {
	return &postOrderTransport{
		errorCreator: errorCreator,
	}
}

// GetCountTransport transport interface
type GetCountTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.Request, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error)
}

type getCountTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getCountTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.Request, err error) {
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		return models.Request{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getCountTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetCountTransport the transport creator for http requests
func NewGetCountTransport(errorCreator errorCreator) GetCountTransport {
	return &getCountTransport{
		errorCreator: errorCreator,
	}
}

// GetOrderTransport transport interface
type GetOrderTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.Request, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error)
}

type getOrderTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getOrderTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.Request, err error) {
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		return models.Request{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getOrderTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetOrderTransport the transport creator for http requests
func NewGetOrderTransport(errorCreator errorCreator) GetOrderTransport {
	return &getOrderTransport{
		errorCreator: errorCreator,
	}
}
