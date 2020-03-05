package httpclient

import (
	`context`
	`log`
	`net/http`
	`testing`
	`time`

	`github.com/stretchr/testify/assert`
	`github.com/valyala/fasthttp`

	`github.com/LuLStackCoder/test-service/pkg/models`
	`github.com/LuLStackCoder/test-service/pkg/service`
	`github.com/LuLStackCoder/test-service/pkg/service/httpserver`
)

const (
	serverAddr               = "localhost:8080"
	maxConns                 = 512
	maxRequestBodySize       = 15 * 1024 * 1024
	serverTimeout            = 1 * time.Millisecond
	serverLaunchingWaitSleep = 1 * time.Second

	methodGetUser   = "GetUser"
	methodPostOrder = "PostOrder"
	methodGetCount  = "GetCount"
	methodGetOrder  = "GetOrder"

	testId        = 12
	testError     = false
	testErrorText = ""

	getUserSuccess      = "Get user success test"
	postOrderSuccess    = "TestPostOrderSuccess"
	getUserCountSuccess = "TestGetUserCountSuccess"
	gettOrdersSuccess   = "TestGetOrdersSuccess"
)

var (
	nilError error

	testData                          = models.DataStruct{Res: true}
	testCustomError map[string]string = nil
)

func Test_client_GetUserSuccess(t *testing.T) {
	t.Run(getUserSuccess, func(t *testing.T) {
		request := makeClientRequest()
		response := makeClientResponse()
		serviceMock := new(service.Mock)
		serviceMock.On(methodGetUser, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetUser(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)
	})
}

func Test_client_PostOrder(t *testing.T) {
	t.Run(postOrderSuccess, func(t *testing.T) {
		request := makeClientRequest()
		response := makeClientResponse()
		serviceMock := new(service.Mock)
		serviceMock.On(methodPostOrder, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.PostOrder(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}

func Test_client_GetUserCount(t *testing.T) {
	t.Run(getUserCountSuccess, func(t *testing.T) {
		request := makeClientRequest()
		response := makeClientResponse()
		serviceMock := new(service.Mock)
		serviceMock.On(methodGetCount, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetCount(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}

func Test_client_GetOrders(t *testing.T) {
	t.Run(gettOrdersSuccess, func(t *testing.T) {
		response := models.Response{}
		serviceMock := new(service.Mock)
		serviceMock.On(methodGetOrder, context.Background()).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetOrder(context.Background())
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}

func makeServerClient(serverAddr string, svc Service) (server *fasthttp.Server, client Service) {
	client = NewPreparedClient(
		serverAddr,
		serverAddr,
		maxConns,
		nil,
		httpserver.NewErrorProcessor(http.StatusInternalServerError, "Internal Error"),
		httpserver.NewError,
		URIPathClientGetUser,
		URIPathClientPostOrder,
		URIPathClientGetUserCount,
		URIPathClientGetOrders,
		HTTPMethodGetUser,
		HTTPMethodPostOrder,
		HTTPMethodGetUserCount,
		HTTPMethodGetOrders,
	)
	router := httpserver.NewPreparedServer(svc)
	server = &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: maxRequestBodySize,
		ReadTimeout:        serverTimeout,
	}

	go func() {
		err := server.ListenAndServe(serverAddr)
		if err != nil {
			log.Printf("server shut down err: %v", err)
		}
	}()
	return
}

func makeClientRequest() *models.Request {
	return &models.Request{
		Id: testId,
	}
}

func makePostOrderRequest() *models.Request {
	return &models.Request{
		Id: testId,
	}
}

func makeClientResponse() models.Response {
	return models.Response{
		Data: &models.DataStruct{Res: true},
	}
}
