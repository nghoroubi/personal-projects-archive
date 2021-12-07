package endpoint

import (
	"amadeus-go/models"
	"amadeus-go/pkg/service"
	"context"
	"fmt"
	
	"github.com/go-kit/kit/endpoint"
)

// HealthCheckRequest ...
type HealthCheckRequest struct{}

// HealthCheckResponse ...
type HealthCheckResponse struct {
	Status bool `json:"status"`
}

// ReserveFlightRequest ...
type ReserveFlightRequest struct {
	SessionId       string               `json:"session_id"`
	ServiceId       int32                `json:"service_id"`
	FareSourceCode  string               `json:"fare_source_code"`
	ClientUniqueId  string               `json:"client_unique_id"`
	MarkupForAdult  uint64               `json:"markup_for_adult"`
	MarkupForChild  uint64               `json:"markup_for_child"`
	MarkupForInfant uint64               `json:"markup_for_infant"`
	TravelerInfo    *models.TravelerInfo `json:"traveler_info"`
	TravelInfo      *models.TravelInfo   `json:"travel_info"`
}

// ReserveFlightResponse ...
type ReserveFlightResponse struct {
	Response *models.ReserveResponse `json:"response"`
}

// CancelPNRRequest ...
type CancelPNRRequest struct {
	PNR string `json:"pnr"`
}

// CancelPNRResponse ...
type CancelPNRResponse struct {
	Resp  *models.CancelPNRResponse `json:"resp"`
	Error error                     `json:"err"`
}

// AirBookRequest ...
type AirBookRequest struct {
	SessionId string `validate:"required" json:"session_id"`
	UniqueId  string `validate:"required" json:"unique_id"`
	Email     string `validate:"required" json:"email"`
}

// AirBookResponse ...
type AirBookResponse struct {
	Response *models.AirBookingResponse `json:"response"`
	Error    error                      `json:"error"`
}

// SearchRequest ...
type SearchRequest struct {
	Origin        string         `json:"origin"`         // e.g THR
	Destination   string         `json:"destination"`    // e.g MHD
	DepartureTime string         `json:"departure_time"` // e.g 2000-12-26T01:15:06+04:20
	ReturnTime    string         `json:"return_time"`    // e.g 2000-12-26T01:15:06+04:20
	AdultCount    int32          `json:"adult_count"`    // e.g 2
	ChildCount    int32          `json:"child_count"`    // e.g 0
	InfantCount   int32          `json:"infant_count"`   // e.g 1
	Airline       string         `json:"airline"`
	AirClass      string         `json:"air_class"`
	Sort          *models.Sort   `json:"sort"`
	Filter        *models.Filter `json:"filter"`
}

// SearchResponse ...
type SearchResponse struct {
	Errors []*models.StdError `json:"errors"`
	Data   []*models.Data     `json:"data"`
	Meta   *models.Meta       `json:"meta"`
	Links  *models.Link       `json:"links"`
}

// AirRevalidateRequest ...
type AirRevalidateRequest struct {
	Origin         string  `json:"origin"`
	Destination    string  `json:"destination"`
	DepartureTime  string  `json:"departure_time"`
	FlightClass    string  `json:"flight_class"`
	FlightNo       string  `json:"flight_no"`
	AdultCount     int32   `json:"adult_count"`
	ChildCount     int32   `json:"child_count"`
	InfantCount    int32   `json:"infant_count"`
	AdultPrice     float64 `json:"adult_price"`
	ChildPrice     float64 `json:"child_price"`
	InfantPrice    float64 `json:"infant_price"`
	AvailableSeats int32   `json:"available_seats"`
}

// AirRevalidateResponse ...
type AirRevalidateResponse struct {
	Response *models.AirRevalidateResponse `json:"response"`
	Error    error                         `json:"error"`
}

// GetAirRulesRequest ...
type GetAirRulesRequest struct {
	ServiceId      int32  `validate:"required" json:"service_id"`
	SessionId      string `validate:"required" json:"session_id"`
	FareSourceCode string `validate:"required" json:"fare_source_code"`
	UniqueId       string `validate:"required" json:"unique_id"`
}

// AirRulesResponse ...
type AirRulesResponse struct {
	Response *models.AirRulesResponse `json:"response"`
}

// //////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////

// MakeHealthCheckEndpoint ...
func MakeHealthCheckEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status := s.HealthCheck(ctx)
		return HealthCheckResponse{Status: status}, nil
	}
}

// MakeSearchEndpoint ...
func MakeSearchEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchRequest)
		r := models.SearchRequest(req)
		resp := s.Search(ctx, &r)
		return SearchResponse(*resp), nil
	}
}

// MakeReserveFlightEndpoint ...
func MakeReserveFlightEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReserveFlightRequest)
		r := models.ReserveRequest(req)
		resp, err := s.ReserveFlight(ctx, &r)
		if err != nil {
			return ReserveFlightResponse{
				Response: &models.ReserveResponse{
					Errors: []*models.StdError{{
						Code:   "500",
						Title:  err.Error(),
						Detail: err.Error(),
					}},
					Data:  nil,
					Meta:  nil,
					Links: nil,
				},
			}, err
		}
		return ReserveFlightResponse{
			Response: resp,
		}, err
	}
}

// MakeETIssueEndpoint ...
func MakeETIssueEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AirBookRequest)
		resp, err := s.AirBook(ctx, &models.AirBookingRequest{
			SessionId: req.SessionId,
			UniqueId:  req.UniqueId,
			Email:     req.Email,
		})
		return AirBookResponse{
			Error:    err,
			Response: resp,
		}, nil
	}
}

// MakeCancelPNREndpoint ...
func MakeCancelPNREndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CancelPNRRequest)
		resp, err := s.CancelPNR(ctx, &models.CancelPNRRequest{PNR: req.PNR})
		return CancelPNRResponse{
			Error: err,
			Resp:  resp,
		}, nil
	}
}

// MakeAirRulesEndpoint ...
func MakeAirRulesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetAirRulesRequest)
		
		rules := s.AirRules(ctx, &models.AirRulesRequest{
			ServiceId:      req.ServiceId,
			SessionId:      req.SessionId,
			FareSourceCode: req.FareSourceCode,
			UniqueId:       req.UniqueId,
		})
		return AirRulesResponse{
			Response: &models.AirRulesResponse{
				Data: rules.Data,
			},
		}, nil
	}
}

// ///////////////////////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////

// Failure ...
type Failure interface {
	Failed() error
}

// Failed ...
func (r AirBookResponse) Failed() error {
	return r.Error
}

// Failed ...
func (r ReserveFlightResponse) Failed() error {
	if len(r.Response.Errors) > 0 {
		return fmt.Errorf("%s", r.Response.Errors[0].Detail)
	}
	return nil
}

// Failed ...
func (r CancelPNRResponse) Failed() error {
	return r.Error
}

// Failed ...
func (r SearchResponse) Failed() error {
	if len(r.Errors) > 0 {
		
		return fmt.Errorf(r.Errors[0].Title)
	}
	return nil
}

// Failed ...
func (r AirRevalidateResponse) Failed() error {
	return r.Error
}

// //////////////////////////////////////////////////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////
// Search ...
func (e Endpoints) Search(ctx context.Context, req *models.SearchRequest) *models.SearchResponse {
	// var resp *model.SearchResponse
	var (
		data = make([]*models.Data, 0)
	)
	request := SearchRequest(*req)
	response, err := e.SearchEndpoint(ctx, request)
	if err != nil {
		return &models.SearchResponse{
			Errors: []*models.StdError{
				{
					ID:     err.Error(),
					Status: err.Error(),
					Code:   err.Error(),
					Title:  err.Error(),
					Detail: err.Error(),
				},
			},
			Data:  data,
			Meta:  nil,
			Links: nil,
		}
	}
	resp := models.SearchResponse(response.(SearchResponse)) // Assertion response to endpoint.SearResponse , then cast that into model.SearchResponse
	return &resp                                             //  and finally return it as a pointer
	
}

// AirBook ...
func (e Endpoints) ReserveFlight(ctx context.Context, req *models.ReserveRequest) (resp *models.ReserveResponse, err error) {
	request := ReserveFlightRequest(*req)
	response, err := e.ReserveFlightEndpoint(ctx, request)
	
	return response.(ReserveFlightResponse).Response, err
}

// CancelPNR ...
func (e Endpoints) CancelPNR(ctx context.Context, req *models.CancelPNRRequest) (resp *models.CancelPNRResponse, err error) {
	request := &CancelPNRRequest{
		PNR: req.PNR,
	}
	response, err := e.CancelPNREndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CancelPNRResponse).Resp, response.(CancelPNRResponse).Error
}

// AirBookingData ...
func (e Endpoints) AirBook(ctx context.Context, req *models.AirBookingRequest) (resp *models.AirBookingResponse, err error) {
	request := &AirBookRequest{
		UniqueId: req.UniqueId,
		Email:    req.Email,
	}
	response, err := e.ETIssueEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AirBookResponse).Response, response.(AirBookResponse).Error
}

// HealthCheck ...
func (e Endpoints) HealthCheck(ctx context.Context) bool {
	request := HealthCheckRequest{}
	response, err := e.HealthCheckEndpoint(ctx, request)
	if err != nil {
		return false
	}
	return response.(HealthCheckResponse).Status
}

// AirRevalidate ...
func (e Endpoints) AirRevalidate(ctx context.Context, req *models.AirRevalidateRequest) (resp *models.AirRevalidateResponse, err error) {
	request := &AirRevalidateRequest{
		Origin:         req.Origin,
		Destination:    req.Destination,
		DepartureTime:  req.DepartureTime,
		FlightClass:    req.FlightClass,
		FlightNo:       req.FlightNo,
		AdultCount:     req.AdultCount,
		ChildCount:     req.ChildCount,
		InfantCount:    req.InfantCount,
		AdultPrice:     req.AdultPrice,
		ChildPrice:     req.ChildPrice,
		InfantPrice:    req.InfantPrice,
		AvailableSeats: req.AvailableSeats,
	}
	response, err := e.AirRevalidateEndpoint(ctx, request)
	return response.(*AirRevalidateResponse).Response, response.(*AirRevalidateResponse).Error
}

// AirRules ...
func (e Endpoints) AirRules(ctx context.Context) *models.AirRulesResponse {
	request := &GetAirRulesRequest{}
	response, _ := e.AirRulesEndpoint(ctx, request)
	return response.(*AirRulesResponse).Response
}
