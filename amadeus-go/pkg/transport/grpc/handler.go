package grpc

import (
	"context"
	"errors"

	"amadeus-go/config/env"
	"amadeus-go/models"
	"amadeus-go/pkg/endpoint"
	"amadeus-go/pkg/transport/grpc/pb/types"

	"github.com/go-kit/kit/transport/grpc"
)

/*
Encoders &
Decoders
*/
func decodeSearchRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*search_type.SearchRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return endpoint.SearchRequest{
			Origin:        req.Origin,
			Destination:   req.Destination,
			DepartureTime: req.DepartureTime,
			ReturnTime:    "",
			AdultCount:    req.AdultCount,
			ChildCount:    req.ChildCount,
			InfantCount:   req.InfantCount,
		},
		nil
}
func encodeSearchResponse(_ context.Context, r interface{}) (interface{}, error) {
	var data = make([]*search_type.Data, 0)
	var Errs = make([]*search_type.StdError, 0)
	resp, ok := r.(endpoint.SearchResponse)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	if len(resp.Data) > 0 {
		for _, d := range resp.Data {
			offerItems := make([]*search_type.OfferItems, 0)
			
			for _, oi := range d.Relationships.OfferItems {
				services := make([]*search_type.Service, 0)
				
				for _, s := range oi.Relationships.Service {
					segments := make([]*search_type.Segment, 0)
					
					for _, sg := range s.ServiceRelationships.Segment {
						segment := &search_type.Segment{
							ID:   sg.ID,
							Type: sg.Type,
						}
						if sg.Attributes != nil {
							segment.Attributes = &search_type.SegmentAtt{
								TravelClass:    sg.Attributes.TravelClass,
								FareClass:      sg.Attributes.FareClass,
								Availability:   sg.Attributes.Availability,
								FareBasis:      sg.Attributes.FareBasis,
								Duration:       sg.Attributes.Duration,
								DurationPerMin: sg.Attributes.DurationPerMin,
								Arrival: &search_type.DateDetail{
									IataCode: sg.Attributes.Arrival.IataCode,
									At:       sg.Attributes.Arrival.At,
								},
								Departure: &search_type.DateDetail{
									IataCode: sg.Attributes.Departure.IataCode,
									At:       sg.Attributes.Departure.At,
								},
								CarrierCode:  sg.Attributes.CarrierCode,
								Number:       sg.Attributes.Number,
								AircraftCode: sg.Attributes.AircraftCode,
								Baggage:      sg.Attributes.Baggage,
								IsCharter:    sg.Attributes.IsCharter,
							}
						}
						if sg.Relationships != nil {
							
							segment.Relationships = &search_type.SegmentRelationships{
								Operating: &search_type.Operating{
									CarrierCode: sg.Relationships.Operating.CarrierCode,
									Number:      sg.Relationships.Operating.Number,
								},
							}
						}
						segments = append(segments, segment)
					}
					
					service := &search_type.Service{
						Attributes: &search_type.ServiceAttributes{},
						Relationships: &search_type.ServiceRelationships{
							Segment: segments,
						},
					}
					if s.Attributes != nil {
						service.Attributes.Duration = s.Attributes.Duration
						service.Attributes.DurationPerMin = s.Attributes.DurationPerMin
					}
					services = append(services, service)
				}
				offerItem := &search_type.OfferItems{
					Relationships: &search_type.OfferItemRelationship{
						Price: &search_type.Price{
							Total:    oi.Relationships.Price.Total,
							TotalTax: oi.Relationships.Price.TotalTax,
						},
						PricePerAdult: &search_type.Price{
							Total:    oi.Relationships.PricePerAdult.Total,
							TotalTax: oi.Relationships.PricePerAdult.TotalTax,
						},
						PricePerChild: &search_type.Price{
							Total:    oi.Relationships.PricePerChild.Total,
							TotalTax: oi.Relationships.PricePerChild.TotalTax,
						},
						PricePerInfant: &search_type.Price{
							Total:    oi.Relationships.PricePerInfant.Total,
							TotalTax: oi.Relationships.PricePerInfant.TotalTax,
						},
						Service: services,
					},
					Links: nil,
				}
				offerItems = append(offerItems, offerItem)
			}
			
			_data := &search_type.Data{
				ID:   d.ID,
				Type: d.Type,
				Attributes: &search_type.Attributes{
					Currency: d.Attributes.Currency,
					Stop:     d.Attributes.Stop,
					Date: &search_type.Date{
						ArivalDate:    d.Attributes.Date.ArrivalDate,
						DepartureDate: d.Attributes.Date.DepartureDate,
					},
				},
				Relationships: &search_type.Relationships{
					OfferItems: offerItems,
					Links:      nil,
				},
			}
			data = append(data, _data)
		}
	}
	if len(resp.Errors) > 0 {
		for _, err := range resp.Errors {
			Err := &search_type.StdError{
				ID:     err.ID,
				Status: err.Status,
				Code:   err.Code,
				Title:  err.Title,
				Detail: err.Detail,
			}
			Errs = append(Errs, Err)
		}
		
		return &search_type.SearchReply{
				Errors: Errs,
				Data:   nil,
				Meta:   nil,
				Links:  nil,
			},
			errors.New(resp.Errors[0].Detail)
	}
	
	return &search_type.SearchReply{
			Errors: Errs,
			Data:   data,
			Meta: &search_type.Meta{
				Pages: &search_type.Pages{
					TotalPages: resp.Meta.Pages.TotalPages,
					From:       resp.Meta.Pages.From,
					LastPage:   resp.Meta.Pages.LastPage,
					To:         resp.Meta.Pages.To,
					Total:      resp.Meta.Pages.Total,
				},
				Currency:     resp.Meta.Currency,
				Defaults:     nil,
				Dictionaries: nil,
			},
			Links: nil,
		},
		nil
}
func decodeHealthCheckRequest(_ context.Context, r interface{}) (interface{}, error) {
	_, ok := r.(*search_type.HealthCheckRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	return endpoint.HealthCheckRequest{}, nil
}
func encodeHealthCheckResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.HealthCheckResponse)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return &search_type.HealthCheckReply{
		Ok: resp.Status,
	}, nil
}
func decodeReserveFlightRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*search_type.AirBookRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	passengers := make([]*models.AirTraveler, 0)
	for _, p := range req.TravelerInfo.AirTravelers {
		passenger := &models.AirTraveler{
			DateOfBirth:   p.DateOfBirth,
			Gender:        p.Gender,
			PassengerType: p.PassengerType,
			PassengerName: &models.PassengerName{
				PassengerFirstName:  p.PassengerName.PassengerFirstName,
				PersianFirstName:    p.PassengerName.PassengerFirstName,
				PassengerMiddleName: p.PassengerName.PassengerMiddleName,
				PassengerLastName:   p.PassengerName.PassengerLastName,
				PersianLastName:     p.PassengerName.PassengerLastName,
				PassengerTitle:      p.PassengerName.PassengerTitle,
			},
			Passport: &models.Passport{
				Country:        p.Passport.Country,
				ExpiryDate:     p.Passport.ExpiryDate,
				IssueDate:      p.Passport.IssueDate,
				PassportNumber: p.Passport.PassportNumber,
			},
			NationalId:          p.NationalId,
			Nationality:         p.Nationality,
			ExtraServiceId:      p.ExtraServiceId,
			FrequentFlyerNumber: p.FrequentFlyerNumber,
			SeatPreference:      p.SeatPreference,
			MealPreference:      p.MealPreference,
			Wheelchair:          p.Wheelchair,
		}
		if p.Visa != nil {
			passenger.Visa = &models.Visa{
				VisaType:    p.Visa.VisaType,
				VisaCountry: p.Visa.VisaCountry,
			}
		}
		passengers = append(passengers, passenger)
	}
	
	return endpoint.ReserveFlightRequest{
		SessionId: req.GetSessionId(),
		
		FareSourceCode:  req.GetFareSourceCode(),
		ClientUniqueId:  req.GetClientUniqueId(),
		MarkupForAdult:  uint64(req.GetMarkupForAdult()),
		MarkupForChild:  uint64(req.GetMarkupForChild()),
		MarkupForInfant: uint64(req.GetMarkupForInfant()),
		TravelerInfo: &models.TravelerInfo{
			Email:        req.TravelerInfo.GetEmail(),
			PhoneNumber:  req.TravelerInfo.GetPhoneNumber(),
			AirTravelers: passengers,
		},
		TravelInfo: &models.TravelInfo{
			Origin:        req.TravelInfo.GetOrigin(),
			Destination:   req.TravelInfo.GetDestination(),
			FlightClass:   req.TravelInfo.GetFlightClass(),
			FlightNo:      req.TravelInfo.GetFlightNo(),
			DepartureDate: req.TravelInfo.GetDepartureDate(),
		},
	}, nil
}
func encodeReserveFlightResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.ReserveFlightResponse)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return &search_type.AirBookReply{
		StdError: []*search_type.StdError{},
		Data: &search_type.AirBookResponseData{
			ID:   resp.Response.Data.ID,
			Type: resp.Response.Data.Type,
			Attributes: &search_type.AirBookAtt{
				Success:      resp.Response.Data.Attributes.Success,
				TktTimeLimit: resp.Response.Data.Attributes.TktTimeLimit,
				Category:     resp.Response.Data.Attributes.Category,
				Status:       resp.Response.Data.Attributes.Status,
				UniqueId:     resp.Response.Data.Attributes.UniqueId,
			},
		},
	}, nil
}
func decodeCancelPNRRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*search_type.CancelPNRRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return endpoint.CancelPNRRequest{PNR: req.PNR}, nil
}
func encodeCancelPNRResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.CancelPNRResponse)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return &search_type.CancelPNRReply{
		Error: resp.Resp.Error,
	}, nil
}
func decodeETIssueRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*search_type.AirBookingDataRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return endpoint.AirBookRequest{
		UniqueId: req.UniqueId,
		Email:    env.GetString("rules.booking_email"),
	}, nil
}
func encodeETIssueResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.AirBookResponse)
	if !ok || resp.Response == nil {
		return nil, errors.New("conversion error")
	}
	
	var (
		customerInfo = make([]*search_type.CustomerInfoes, 0)
		
		resItems = make([]*search_type.ReservationItems, 0)
		
		trDetails = make([]*search_type.TripDetailPtcFareBreakdowns, 0)
		
		bNotes = make([]*search_type.BookingNotes, 0)
		svc    = make([]*search_type.Services, 0)
	)
	
	// CustomerInfo
	if resp.Response != nil && resp.Response.Data != nil {
		for _, ci := range resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.CustomerInfoes {
			ETickets := make([]*search_type.ETicket, 0)
			for _, et := range ci.Relationships.ETicketNumbers {
				eTicket := &search_type.ETicket{
					ETicketNumber: et.ETicketNumber,
					DateOfIssue:   et.DateOfIssue,
					IsRefunded:    et.IsRefunded,
					AirlinePnr:    et.AirlinePnr,
					TotalRefund:   float32(et.TotalRefund),
				}
				ETickets = append(ETickets, eTicket)
			}
			
			cInfo := &search_type.CustomerInfoes{
				ID:   ci.ID,
				Type: ci.Type,
				Attributes: &search_type.CustomerInfoesAttributes{
					ETickets: ci.Attributes.ETickets,
				},
				Relationships: &search_type.CustomerInfoesRelationships{
					Customer: &search_type.Customer{
						PassengerType:      ci.Relationships.Customer.PassengerType,
						PassportNumber:     ci.Relationships.Customer.PassportNumber,
						NationalId:         ci.Relationships.Customer.NationalId,
						DateOfBirth:        ci.Relationships.Customer.DateOfBirth,
						PassportExpireDate: ci.Relationships.Customer.PassportExpireDate,
						PaxName: &search_type.PassengerName{
							PassengerFirstName:  ci.Relationships.Customer.PaxName.PassengerFirstName,
							PassengerLastName:   ci.Relationships.Customer.PaxName.PassengerLastName,
							PassengerMiddleName: ci.Relationships.Customer.PaxName.PassengerMiddleName,
							PassengerTitle:      ci.Relationships.Customer.PaxName.PassengerTitle,
						},
					},
					ETicketNumbers: ETickets,
				},
			}
			customerInfo = append(customerInfo, cInfo)
		}
		
		// ReservationItems
		for _, ri := range resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ReservationItems {
			tStops := make([]*search_type.TechnicalStops, 0)
			for _, ts := range ri.Relationships.TechnicalStops {
				stop := &search_type.TechnicalStops{
					Data: &search_type.TechnicalStopsData{
						ArrivalAirport:    ts.Data.ArrivalAirport,
						ArrivalDateTime:   ts.Data.ArrivalDateTime,
						DepartureDateTime: ts.Data.DepartureDateTime,
					},
				}
				tStops = append(tStops, stop)
			}
			
			rItem := &search_type.ReservationItems{
				ID:   ri.ID,
				Type: ri.Type,
				Attributes: &search_type.ReservationItemsAttributes{
					AirEquipmentType:             ri.Attributes.AirEquipmentType,
					AirlinePnr:                   ri.Attributes.AirlinePnr,
					ArrivalAirportLocationCode:   ri.Attributes.ArrivalAirportLocationCode,
					ArrivalDateTime:              ri.Attributes.ArrivalDateTime,
					ArrivalTerminal:              ri.Attributes.ArrivalTerminal,
					Baggage:                      ri.Attributes.Baggage,
					DepartureAirportLocationCode: ri.Attributes.DepartureAirportLocationCode,
					DepartureDateTime:            ri.Attributes.DepartureDateTime,
					DepartureTerminal:            ri.Attributes.DepartureTerminal,
					FlightNumber:                 ri.Attributes.FlightNumber,
					JourneyDuration:              ri.Attributes.JourneyDuration,
					JourneyDurationPerMinute:     ri.Attributes.JourneyDurationPerMinute,
					MarketingAirlineCode:         ri.Attributes.MarketingAirlineCode,
					OperatingAirlineCode:         ri.Attributes.OperatingAirlineCode,
					ResBookDesignCode:            ri.Attributes.ResBookDesignCode,
					StopQuantity:                 ri.Attributes.StopQuantity,
					IsCharter:                    ri.Attributes.IsCharter,
				},
				Relationships: &search_type.ReservationItemsRelationships{
					TechnicalStops: tStops,
				},
			}
			resItems = append(resItems, rItem)
		}
		
		// TripDetail
		for _, bd := range resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.TripDetailPtcFareBreakdowns {
			bDown := &search_type.TripDetailPtcFareBreakdowns{
				Relationships: &search_type.TripDetailPtcFareBreakdownsRelationships{
					PassengerTypeQuantity: &search_type.PassengerTypeQuantity{
						PassengerType: bd.Relationships.PassengerTypeQuantity.PassengerType,
						Quantity:      bd.Relationships.PassengerTypeQuantity.Quantity,
					},
					TripDetailPassengerFare: &search_type.TripDetailPassengerFare{
						BaseFare:   bd.Relationships.TripDetailPassengerFare.BaseFare,
						ServiceTax: bd.Relationships.TripDetailPassengerFare.ServiceTax,
						Tax:        bd.Relationships.TripDetailPassengerFare.Tax,
						TotalFare:  bd.Relationships.TripDetailPassengerFare.TotalFare,
						Commission: bd.Relationships.TripDetailPassengerFare.Commission,
						Currency:   bd.Relationships.TripDetailPassengerFare.Currency,
					},
				},
			}
			trDetails = append(trDetails, bDown)
		}
		// bookingNotes
		for _, bn := range resp.Response.Data.Relationships.TravelItinerary.Relationships.BookingNotes {
			b := &search_type.BookingNotes{
				Note: bn.Note,
				Date: bn.Date,
			}
			bNotes = append(bNotes, b)
		}
		// Services
		for _, s := range resp.Response.Data.Relationships.TravelItinerary.Relationships.Services {
			service := &search_type.Services{
				PassengerFirstName:  s.PassengerFirstName,
				PassengerMiddleName: s.PassengerMiddleName,
				PassengerLastName:   s.PassengerLastName,
				Description:         s.Description,
				ServiceId:           s.ServiceId,
				Behavior:            s.Behavior,
				Amount:              s.Amount,
				Currency:            s.Currency,
			}
			svc = append(svc, service)
		}
	}
	return &search_type.AirBookingDataReply{
		StdError: nil,
		Data: &search_type.AirBookingData{
			ID:   "",
			Type: env.GetString("service.name"),
			Attributes: &search_type.AirBookingDataAttributes{
				UniqueId:      resp.Response.Data.Attributes.UniqueId,
				BookedBy:      resp.Response.Data.Attributes.BookedBy,
				OrderBy:       resp.Response.Data.Attributes.OrderBy,
				ClientBalance: resp.Response.Data.Attributes.ClientBalance,
				Success:       resp.Response.Data.Attributes.Success,
				TktTimeLimit:  resp.Response.Data.Attributes.TktTimeLimit,
				Category:      resp.Response.Data.Attributes.Category,
				Status:        resp.Response.Data.Attributes.Status,
				RefundMethod:  resp.Response.Data.Attributes.RefundMethod,
			},
			Relationships: &search_type.AirBookingDataRelationships{
				TravelItinerary: &search_type.TravelItinerary{
					Relationships: &search_type.TravelItineraryRelationships{
						ItineraryInfo: &search_type.ItineraryInfo{
							Relationships: &search_type.ItineraryInfoRelationships{
								ItineraryPricing: &search_type.ItineraryPricing{
									BaseFare:        resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.BaseFare,
									ServiceTax:      resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.ServiceTax,
									TotalTax:        resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.TotalTax,
									TotalFare:       resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.TotalFare,
									TotalCommission: resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.TotalCommission,
									Currency:        resp.Response.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.Currency,
								},
								CustomerInfoes:              customerInfo,
								ReservationItems:            resItems,
								TripDetailPtcFareBreakdowns: trDetails,
							},
						},
						BookingNotes: bNotes,
						Services:     svc,
					},
				},
			},
		},
		Meta:  nil,
		Links: nil,
	}, nil
}
func decodeAirRevalidateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*search_type.AirRevalidateRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return endpoint.AirRevalidateRequest{
		Origin:         req.GetOrigin(),
		Destination:    req.GetDestination(),
		DepartureTime:  req.GetDepartureTime(),
		FlightClass:    req.GetFlightClass(),
		FlightNo:       req.GetFlightNo(),
		AdultCount:     req.GetAdultCount(),
		ChildCount:     req.GetChildCount(),
		InfantCount:    req.GetInfantCount(),
		AdultPrice:     float64(req.GetAdultPrice()),
		ChildPrice:     float64(req.GetChildPrice()),
		InfantPrice:    float64(req.GetInfantPrice()),
		AvailableSeats: req.GetAvailableSeats(),
	}, nil
}
func encodeAirRevalidateResponse(_ context.Context, r interface{}) (interface{}, error) {
	_, ok := r.(endpoint.AirRevalidateResponse)
	if !ok {
		return nil, errors.New("conversion error")
	}
	return &search_type.AirRevalidateReply{
	}, nil
}
func encodeAirRulesResponse(_ context.Context, r interface{}) (response interface{}, err error) {
	resp, ok := r.(endpoint.AirRulesResponse)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	rules := make([]*search_type.FareRules, 0)
	
	for _, fareRules := range resp.Response.Data.Relationships.FareRules {
		ruleDetails := make([]*search_type.RuleDetails, 0)
		for _, rd := range fareRules.Relationships.RuleDetails {
			rDet := &search_type.RuleDetails{
				Category: rd.Category,
				Rules:    rd.Rules,
			}
			ruleDetails = append(ruleDetails, rDet)
		}
		fr := &search_type.FareRules{
			ID:   fareRules.ID,
			Type: fareRules.Type,
			Attributes: &search_type.FareRulesAttributes{
				Airline:  fareRules.Attributes.Airline,
				CityPair: fareRules.Attributes.Airline,
			},
			Relationships: &search_type.FareRulesRelationships{
				RuleDetails: ruleDetails,
			},
		}
		rules = append(rules, fr)
	}
	
	return &search_type.AirRulesResponse{
		StdError: nil,
		Data: &search_type.AirRulesData{
			ID:   resp.Response.Data.ID,
			Type: resp.Response.Data.Type,
			Attributes: &search_type.AirRulesAttributes{
				Success:  resp.Response.Data.Attributes.Success,
				FareType: resp.Response.Data.Attributes.FareType,
			},
			Relationships: &search_type.AirRulesRelationships{
				FareRules: rules,
			},
		},
	}, nil
}
func decodeAirRulesRequest(_ context.Context, r interface{}) (request interface{}, err error) {
	req, ok := r.(*search_type.AirRulesRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	return endpoint.GetAirRulesRequest{
		SessionId:      req.SessionId,
		FareSourceCode: req.FareSourceCode,
		UniqueId:       req.UniqueId,
	}, nil
}

/*
Handler Makers
*/
func makeSearchHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SearchEndpoint, decodeSearchRequest, encodeSearchResponse, options...)
}
func makeHealthCheckHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.HealthCheckEndpoint, decodeHealthCheckRequest, encodeHealthCheckResponse, options...)
}
func makeAirBookHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ReserveFlightEndpoint, decodeReserveFlightRequest, encodeReserveFlightResponse, options...)
}
func makeAirBookingDataHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ETIssueEndpoint, decodeETIssueRequest, encodeETIssueResponse, options...)
}
func makeCancelPNRHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CancelPNREndpoint, decodeCancelPNRRequest, encodeCancelPNRResponse, options...)
}
func makeAirRevalidateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AirRevalidateEndpoint, decodeAirRevalidateRequest, encodeAirRevalidateResponse, options...)
}
func makeAirRulesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AirRulesEndpoint, decodeAirRulesRequest, encodeAirRulesResponse, options...)
}

/*
Handlers
*/

// AirLowFareSearch ...
func (g *grpcServer) AirLowFareSearch(ctx context.Context, req *search_type.SearchRequest) (*search_type.SearchReply, error) {
	_, rep, err := g.search.ServeGRPC(ctx, req)
	if rep != nil {
		return rep.(*search_type.SearchReply), err
	} else {
		return nil, err
	}
}

// HealthCheck ...
func (g *grpcServer) HealthCheck(ctx context.Context, req *search_type.HealthCheckRequest) (*search_type.HealthCheckReply, error) {
	_, rep, err := g.healthCheck.ServeGRPC(ctx, req)
	if rep != nil {
		return rep.(*search_type.HealthCheckReply), err
	} else {
		return nil, err
	}
}

// AirBook ...
func (g *grpcServer) AirBook(ctx context.Context, req *search_type.AirBookRequest) (*search_type.AirBookReply, error) {
	_, resp, err := g.airBook.ServeGRPC(ctx, req)
	if resp != nil {
		return resp.(*search_type.AirBookReply), err
	} else {
		return nil, err
	}
}

// CancelPNR ...
func (g *grpcServer) CancelPNR(ctx context.Context, req *search_type.CancelPNRRequest) (*search_type.CancelPNRReply, error) {
	_, rep, err := g.cancelPNR.ServeGRPC(ctx, req)
	if rep != nil {
		return rep.(*search_type.CancelPNRReply), err
	} else {
		return nil, err
	}
}

// AirBookingData ...
func (g *grpcServer) AirBookingData(ctx context.Context, req *search_type.AirBookingDataRequest) (*search_type.AirBookingDataReply, error) {
	_, rep, err := g.airBookingData.ServeGRPC(ctx, req)
	if rep != nil {
		return rep.(*search_type.AirBookingDataReply), err
	} else {
		return nil, err
	}
}

// AirRevalidate ...
func (g *grpcServer) AirRevalidate(ctx context.Context, req *search_type.AirRevalidateRequest) (*search_type.AirRevalidateReply, error) {
	_, rep, err := g.airRevalidate.ServeGRPC(ctx, req)
	if rep != nil {
		return rep.(*search_type.AirRevalidateReply), err
	} else {
		return nil, err
	}
}

// AirRules ...
func (g *grpcServer) AirRules(ctx context.Context, req *search_type.AirRulesRequest) (*search_type.AirRulesResponse, error) {
	_, rep, err := g.aieRules.ServeGRPC(ctx, req)
	if rep != nil {
		return rep.(*search_type.AirRulesResponse), nil
	} else {
		return nil, err
	}
}