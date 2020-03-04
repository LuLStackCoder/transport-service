package main

import (
	`net/http/pprof`
	`os`
	`time`

	fasthttpprometheus `github.com/flf2ko/fasthttp-prometheus`
	`github.com/go-kit/kit/log`
	`github.com/go-kit/kit/log/level`
	`github.com/kelseyhightower/envconfig`
	`github.com/valyala/fasthttp`
	`github.com/valyala/fasthttp/fasthttpadaptor`
	_`github.com/kelseyhightower/envconfig`
	`github.com/LuLStackCoder/test-service/pkg/service`

	`github.com/LuLStackCoder/test-service/pkg/service/httpserver`
)

var (
	serviceVersion = "dev"
	methodError    = []string{"method", "error"}
)

type configuration struct {
	Port               string `envconfig:"PORT" default:"8080"`
	MaxRequestBodySize int    `envconfig:"MAX_REQUEST_BODY_SIZE" default:"10485760"` // 10 MB
	Debug              bool   `envconfig:"DEBUG" default:"false"`

	AllowedIssuers string        `envconfig:"ALLOWED_ISSUERS" default:"suppliers-portal-dev,portal-purchase-admin-dev"`
	ReadTimeout    time.Duration `envconnfig:"READ_TIMEOUT" default:"1s"`

	Subsystem    string `envconfig:"METRICS_SUBSYSTEM" default:""`
}

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = level.Info(logger).Log("msg", "initializing", "version", serviceVersion)
	svc := service.NewService()
	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		_ = level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}
	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}
	router := httpserver.NewPreparedServer(svc)
	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
	p := fasthttpprometheus.NewPrometheus(cfg.Subsystem)
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
