package service

import (
	"amadeus-go/models"
	"context"
	
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(Service) Service

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (l loggingMiddleware) GenerateControlNumber(refNum string) (string, error) {
	panic("implement me")
}

// LoggingMiddleware takes a logger as a dependency
// and returns a IranairtourService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{logger, next}
	}
	
}

func (l loggingMiddleware) Search(ctx context.Context, req *models.SearchRequest) (m0 *models.SearchResponse) {
	defer func() {
		l.logger.Log("method", "Search", "req", req, "m0", m0)
	}()
	return l.next.Search(ctx, req)
}

func (l loggingMiddleware) ReserveFlight(ctx context.Context, req *models.ReserveRequest) (m0 *models.ReserveResponse, e1 error) {
	defer func() {
		l.logger.Log("method", "ReserveFlight", "req", req, "m0", m0, "e1", e1)
	}()
	return l.next.ReserveFlight(ctx, req)
}
func (l loggingMiddleware) CancelPNR(ctx context.Context, req *models.CancelPNRRequest) (m0 *models.CancelPNRResponse, e1 error) {
	defer func() {
		l.logger.Log("method", "CancelPNR", "req", req, "m0", m0, "e1", e1)
	}()
	return l.next.CancelPNR(ctx, req)
}
func (l loggingMiddleware) AirBook(ctx context.Context, req *models.AirBookingRequest) (m0 *models.AirBookingResponse, e1 error) {
	defer func() {
		l.logger.Log("method", "AirBookingData", "req", req, "m0", m0, "e1", e1)
	}()
	return l.next.AirBook(ctx, req)
}
func (l loggingMiddleware) HealthCheck(ctx context.Context) (b0 bool) {
	defer func() {
		l.logger.Log("method", "HealthCheck", "b0", b0)
	}()
	return l.next.HealthCheck(ctx)
}
func (l loggingMiddleware) AirRules(ctx context.Context, req *models.AirRulesRequest) (m0 *models.AirRulesResponse) {
	defer func() {
		l.logger.Log("method", "AirRules", "m0", m0)
	}()
	return l.next.AirRules(ctx,req)
}
