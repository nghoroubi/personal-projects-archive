package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"amadeus-go/pkg/endpoint"

	http1 "github.com/go-kit/kit/transport/http"
	httpTransport "github.com/go-kit/kit/transport/http"
)

func decodeHealthCheckRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	req := endpoint.HealthCheckRequest{}
	return req, nil
}
func decodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SearchRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func decodeReserveFlightRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReserveFlightRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func decodeCancelPNRRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CancelPNRRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func decodeETIssueRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AirBookRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func decodeAirRevalidateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AirRevalidateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func decodeAirRulesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetAirRulesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeHealthCheckResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// //////////////////////////////////////////////////////////////////
func makeHealthCheckHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []httpTransport.ServerOption) {
	m.Handle("/health", httpTransport.NewServer(endpoints.HealthCheckEndpoint, decodeHealthCheckRequest, encodeHealthCheckResponse, options...))
}
func makeSearchHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/flights", http1.NewServer(endpoints.SearchEndpoint, decodeSearchRequest, encodeResponse, options...))
}
func makeReserveFlightHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []httpTransport.ServerOption) {
	m.Handle("/reserve", httpTransport.NewServer(endpoints.ReserveFlightEndpoint, decodeReserveFlightRequest, encodeResponse, options...))
}
func makeAirBookHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []httpTransport.ServerOption) {
	m.Handle("/book", httpTransport.NewServer(endpoints.ETIssueEndpoint, decodeETIssueRequest, encodeResponse, options...))
}
func makeCancelPNRHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []httpTransport.ServerOption) {
	m.Handle("/cancel_pnr", httpTransport.NewServer(endpoints.CancelPNREndpoint, decodeCancelPNRRequest, encodeResponse, options...))
}
func makeAirRevalidateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []httpTransport.ServerOption) {
	m.Handle("/air_revalidate", httpTransport.NewServer(endpoints.AirRevalidateEndpoint, decodeAirRevalidateRequest, encodeResponse, options...))
}

func makeAirRulesHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []httpTransport.ServerOption) {
	m.Handle("/air_rules", httpTransport.NewServer(endpoints.AirRulesEndpoint, decodeAirRulesRequest, encodeResponse, options...))
}

// ErrorEncoder ...
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	_ = json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

// ErrorDecoder ...
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}