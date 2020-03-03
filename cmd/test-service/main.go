package main

import (
	`fmt`
	`net/http/pprof`
	`os`
	`time`

	fasthttpprometheus `github.com/flf2ko/fasthttp-prometheus`
	`github.com/go-kit/kit/log`
	`github.com/go-kit/kit/log/level`
	`github.com/valyala/fasthttp`
	`github.com/valyala/fasthttp/fasthttpadaptor`

	`github.com/LuLStackCoder/test-service/pkg/service`

	`github.com/LuLStackCoder/test-service/pkg/service/httpserver`
)

var (
	serviceVersion = "dev"
	methodError    = []string{"method", "error"}
)

type configuration struct {
	Port               string `envconfig:"PORT" required:"true"`
	MaxRequestBodySize int    `envconfig:"MAX_REQUEST_BODY_SIZE" default:"10485760"` // 10 MB
	Debug              bool   `envconfig:"DEBUG" default:"false"`

	AllowedIssuers string        `envconfig:"ALLOWED_ISSUERS" default:"suppliers-portal-dev,portal-purchase-admin-dev"`
	ReadTimeout    time.Duration `envconnfig:"READ_TIMEOUT" default:"1s"`

	StorageAuthSource             string        `envconfig:"STORAGE_AUTH_SOURCE" required:"true"`
	StorageUsername               string        `envconfig:"STORAGE_USERNAME" required:"true"`
	StoragePassword               string        `envconfig:"STORAGE_PASSWORD" required:"true"`
	StorageBarcodeCountCollection string        `envconfig:"STORAGE_BARCODE_COUNT_COLLECTION" required:"true"`
	StorageBarcodeCollection      string        `envconfig:"STORAGE_BARCODE_COLLECTION" required:"true"`
	StorageNotError               string        `envconfig:"STORAGE_NOT_ERROR" required:"true"`
	StorageClusterAddress         string        `envconfig:"STORAGE_CLUSTER_ADDRESS" required:"true"`
	StorageDatabase               string        `envconfig:"STORAGE_DATABASE_NAME" required:"true"`
	StorageConnectionTimeout      time.Duration `envconfig:"STORAGE_CONNECTION_TIMEOUT" default:"60s"`

	MetricsNamespace    string `envconfig:"METRICS_NAMESPACE" default:"wb"`
	MetricsSubsystem    string `envconfig:"METRICS_SUBSYSTEM" default:"barcode_service"`
	MetricsNameCount    string `envconfig:"METRICS_NAME_COUNT" default:"request_count"`
	MetricsNameDuration string `envconfig:"METRICS_NAME_DURATION" default:"request_duration"`
	MetricsHelpCount    string `envconfig:"METRICS_HELP_COUNT" default:"Request count"`
	MetricsHelpDuration string `envconfig:"METRICS_HELP_DURATION" default:"Request duration"`
}

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = level.Info(logger).Log("msg", "initializing", "version", serviceVersion)
	svc := service.NewService()
	var cfg configuration
	cfg.Port = "8080"
	router := httpserver.NewPreparedServer(svc)
	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
	router.Handle("GET", "/user", func(ctx *fasthttp.RequestCtx) {
		fmt.Printf("->id:%s",ctx.URI().QueryArgs().Peek("id"))
	})
	p := fasthttpprometheus.NewPrometheus(cfg.MetricsSubsystem)
	fasthttpServer := &fasthttp.Server{
		Handler:            p.WrapHandler(router),
		MaxRequestBodySize: cfg.MaxRequestBodySize,
		ReadTimeout:        cfg.ReadTimeout,
	}
	go func() {
		_ = level.Info(logger).Log("msg", "starting http server", "port", cfg.Port)
		if err := fasthttpServer.ListenAndServe(":" + cfg.Port); err != nil {
			_ = level.Error(logger).Log("msg", "server run failure", "err", err)
			os.Exit(1)
		}
	}()
	c := make(chan os.Signal, 1)
	//signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		_ = level.Info(logger).Log("msg", "received signal, exiting", "signal", sig)

		if err := fasthttpServer.Shutdown(); err != nil {
			_ = level.Error(logger).Log("msg", "server shutdown failure", "err", err)
		}

		_ = level.Info(logger).Log("msg", "goodbye")
	}(<-c)

}
