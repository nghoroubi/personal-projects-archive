package service

import (
	"amadeus-go/pkg/endpoint"
	"amadeus-go/pkg/service"
	http1 "amadeus-go/pkg/transport/http"

	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
	"github.com/oklog/oklog/pkg/group"
	openTracingGo "github.com/opentracing/opentracing-go"
)

// createService ...
func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	initGRPCHandler(endpoints, g)
	return g
}

// defaultHttpOptions ...
func defaultHttpOptions(logger log.Logger, tracer openTracingGo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"CancelPNR":      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CancelPNR", logger))},
		"AirBookingData": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AirBookingData", logger))},
		"HealthCheck":    {http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "HealthCheck", logger))},
		"AirBook":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AirBook", logger))},
		"Search":         {http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Search", logger))},
	}
	return options
}

// defaultGRPCOptions ...
func defaultGRPCOptions(logger log.Logger, tracer openTracingGo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"CancelPNR":      {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CancelPNR", logger))},
		"AirBookingData": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "AirBookingData", logger))},
		"HealthCheck":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "HealthCheck", logger))},
		"AirBook":        {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "AirBook", logger))},
		"Search":         {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "Search", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Search"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Search")), endpoint.InstrumentingMiddleware(duration.With("method", "Search"))}
	mw["AirBookingData"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "AirBookingData")), endpoint.InstrumentingMiddleware(duration.With("method", "AirBookingData"))}
	mw["CancelPNR"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CancelPNR")), endpoint.InstrumentingMiddleware(duration.With("method", "CancelPNR"))}
	mw["AirBookingData"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "AirBookingData")), endpoint.InstrumentingMiddleware(duration.With("method", "AirBookingData"))}
	mw["HealthCheck"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "HealthCheck")), endpoint.InstrumentingMiddleware(duration.With("method", "HealthCheck"))}
}

// addDefaultServiceMiddleware ...
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}

// addEndpointMiddlewareToAllMethods ...
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Search", "AirBookingData", "CancelPNR", "AirBookingData", "HealthCheck"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}