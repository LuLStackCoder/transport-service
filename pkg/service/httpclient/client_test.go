package httpclient

import (
	`context`
	`log`
	`net/http`
	`reflect`
	`testing`
	`time`

	`github.com/stretchr/testify/assert`
	`github.com/valyala/fasthttp`

	`github.com/LuLStackCoder/test-service/pkg/models`
	`github.com/LuLStackCoder/test-service/pkg/service`
	`github.com/LuLStackCoder/test-service/pkg/service/httpserver`
)

const (
	serverAddr         = "localhost:8080"
	maxConns           = 512
	maxRequestBodySize = 15 * 1024 * 1024
	serverTimeout            = 1 * time.Millisecond
	serverLaunchingWaitSleep = 1 * time.Second

	methodGetUser   = "GetUser"
	methodPostOrder = "PostOrder"
	methodGetCount  = "GetCount"
	methodGetOrder  = "GetOrder"

	getUserSuccess    = "Get user success test"
)

var (
	nilError error

	testId = 12
	testError = false
	testErrorText = ""
	testData = models.DataStruct{Res: true}
	testCustomError map[string]string = nil
)

func TestNewClient(t *testing.T) {
	type args struct {
		cli                *fasthttp.HostClient
		transportGetUser   GetUserClientTransport
		transportPostOrder PostOrderClientTransport
		transportGetCount  GetCountClientTransport
		transportGetOrder  GetOrderClientTransport
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.cli, tt.args.transportGetUser, tt.args.transportPostOrder, tt.args.transportGetCount, tt.args.transportGetOrder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPreparedClient(t *testing.T) {
	type args struct {
		serverURL      string
		maxConns       int
		errorProcessor errorProcessor
		errorCreator   errorCreator
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPreparedClient(tt.args.serverURL, tt.args.maxConns, tt.args.errorProcessor, tt.args.errorCreator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPreparedClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_GetCount(t *testing.T) {
	type fields struct {
		cli                *fasthttp.HostClient
		transportGetUser   GetUserClientTransport
		transportPostOrder PostOrderClientTransport
		transportGetCount  GetCountClientTransport
		transportGetOrder  GetOrderClientTransport
	}
	type args struct {
		ctx     fasthttp.RequestCtx
		request *models.Request
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse models.Response
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                tt.fields.cli,
				transportGetUser:   tt.fields.transportGetUser,
				transportPostOrder: tt.fields.transportPostOrder,
				transportGetCount:  tt.fields.transportGetCount,
				transportGetOrder:  tt.fields.transportGetOrder,
			}
			gotResponse, err := s.GetCount(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("GetCount() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func Test_client_GetOrder(t *testing.T) {
	type fields struct {
		cli                *fasthttp.HostClient
		transportGetUser   GetUserClientTransport
		transportPostOrder PostOrderClientTransport
		transportGetCount  GetCountClientTransport
		transportGetOrder  GetOrderClientTransport
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse models.Response
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                tt.fields.cli,
				transportGetUser:   tt.fields.transportGetUser,
				transportPostOrder: tt.fields.transportPostOrder,
				transportGetCount:  tt.fields.transportGetCount,
				transportGetOrder:  tt.fields.transportGetOrder,
			}
			gotResponse, err := s.GetOrder(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("GetOrder() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func Test_client_GetUserSuccess(t *testing.T) {
	request := makeClientRequest()
	response := makeClientResponse()
	t.Run(getUserSuccess, func(t *testing.T) {
		serviceMock := new(service.Mock)
		serviceMock.On(methodGetUser, context.Background(), *request).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		_, err := client.GetUser(context.Background(), request)

		assert.NoError(t, err, "unexpected error:", err)
	})
}

func Test_client_PostOrder(t *testing.T) {

}

func makeServerClient(serverAddr string, svc Service) (server *fasthttp.Server, client Service) {
	client = NewPreparedClient(
		serverAddr,
		maxConns,
		httpserver.NewErrorProcessor(http.StatusInternalServerError, "Internal Error"),
		httpserver.NewError,
	)
	router := httpserver.NewPreparedServer(svc)
	server = &fasthttp.Server {
		Handler: router.Handler,
		MaxRequestBodySize: maxRequestBodySize,
		ReadTimeout: serverTimeout,
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

func makeClientResponse() *models.Response {
	return &models.Response{
		Error: testError,
		ErrorText: testErrorText,
		Data: &testData,
		CustomError: testCustomError,
	}
}
