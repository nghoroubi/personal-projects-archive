package service

import (
	"flag"
	"fmt"

	"amadeus-go/config/env"
	"amadeus-go/pkg/endpoint"
	"amadeus-go/pkg/repository"
	"amadeus-go/pkg/service"
	"amadeus-go/pkg/transport/grpc"
	pb "amadeus-go/pkg/transport/grpc/pb"
	svchttp "amadeus-go/pkg/transport/http"

	kitEndpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/hooklift/gowsdl/soap"
	"github.com/oklog/oklog/pkg/group"
	openTracingGo "github.com/opentracing/opentracing-go"
	prometheus1 "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	grpc1 "google.golang.org/grpc"

	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var tracer openTracingGo.Tracer
var logger log.Logger

var fs = flag.NewFlagSet("iranairtour", flag.ExitOnError)
var debugAddr = fs.String("debug.addr", ":19999", "Debug and metrics listen address")
var httpAddr = fs.String("svchttp-addr", ":49989", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":49889", "gRPC listen address")
var confFile = fs.String("config", "./config.json", "config file path")

// Run , entry point of service
func Run() {
	fs.Parse(os.Args[1:])
	
	env.Init(*confFile)
	
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	
	logger.Log("tracer", "none")
	tracer = openTracingGo.GlobalTracer()
	
	logger.Log("tracer", "none")
	tracer = openTracingGo.GlobalTracer()
	
	uri := env.GetString("api.base_url")
	client := soap.NewClient(uri)
	
	client.AddHeader(
		&repository.AuthenticationSoapHeader{
			WSUserName: env.GetString("api.auth.user"),
			WSPassword: env.GetString("api.auth.password"),
		},
	
	)
	
	soapSvc := repository.NewEPowerServiceSoap(client)
	svc := service.New(getServiceMiddleware(logger),soapSvc)
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	
	// Setup service discovery
	// Registering service in discovery service
	registrar := service.Register(env.GetString("service.discovery_address"),
		env.GetString("service.advertise_address"),
		env.GetString("service.http_port"),
		env.GetString("service.grpc_port"))
	
	// register and deregister service
	registrar.Register()
	defer registrar.Deregister()
	
	// Init cache
	repository.InitRedis(env.GetString("service.cache.address"))
	if err := repository.Cache().Ping().Err(); err != nil {
		logger.Log("redis_connection", err)
		panic(err)
	}
	
	defer repository.Cache().Close()
	
	// Start the service
	logger.Log("exit", g.Run())
	
}

// initHttpHandler ...
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger, tracer)
	
	// if the port be configured in config file,
	// The priority of config_file values are higher than the others
	if port := env.GetString("service.http_port"); port != "" {
		addr := fmt.Sprintf(":%s", port)
		httpAddr = &addr
	}
	
	httpHandler := svchttp.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})
	
}

// getServiceMiddleware ...
func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(logger, mw)
	mw = append(mw, service.Recover())
	return
}

// getEndpointMiddleware ...
func getEndpointMiddleware(logger log.Logger) (mw map[string][]kitEndpoint.Middleware) {
	mw = map[string][]kitEndpoint.Middleware{}
	duration := prometheus.NewSummaryFrom(prometheus1.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "iranairtour",
	}, []string{"method", "success"})
	addDefaultEndpointMiddleware(logger, duration, mw)
	
	return
}

// initMetricsEndpoint ...
func initMetricsEndpoint(g *group.Group) {
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http.Serve(debugListener, http.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}

// initCancelInterrupt ...
func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

// initGRPCHandler ...
func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, nil)
	
	// if the port be configured in config file,
	// The priority of config_file values are higher than the others
	if port := env.GetString("service.grpc_port"); port != "" {
		addr := fmt.Sprintf(":%s", port)
		grpcAddr = &addr
	}
	
	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := grpc1.NewServer()
		pb.RegisterServiceServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})
}