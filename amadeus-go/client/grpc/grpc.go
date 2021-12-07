package grpc

import (
	"context"
	"errors"

	"amadeus-go/models"
	endpointSvc "amadeus-go/pkg/endpoint"
	"amadeus-go/pkg/service"
	types "amadeus-go/pkg/transport/grpc/pb/types"

	"github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.Service, error) {
	var searchEndpoint endpoint.Endpoint
	{
		searchEndpoint = grpc1.NewClient(conn, "search.struct.Service", "AirLowFareSearch", encodeSearchRequest, decodeSearchResponse, types.SearchReply{}, options["Search"]...).Endpoint()
	}
	
	var reserveFlightEndpoint endpoint.Endpoint
	{
		reserveFlightEndpoint = grpc1.NewClient(conn, "search.struct.Service", "AirBook", encodeReserveFlightRequest, decodeReserveFlightResponse, types.AirBookReply{}, options["AirBook"]...).Endpoint()
	}
	
	var cancelPNREndpoint endpoint.Endpoint
	{
		cancelPNREndpoint = grpc1.NewClient(conn, "search.struct.Service", "CancelPNR", encodeCancelPNRRequest, decodeCancelPNRResponse, types.CancelPNRReply{}, options["CancelPNR"]...).Endpoint()
	}
	
	var eTIssueEndpoint endpoint.Endpoint
	{
		eTIssueEndpoint = grpc1.NewClient(conn, "search.struct.Service", "AirBookingData", encodeETIssueRequest, decodeETIssueResponse, types.AirBookingDataReply{}, options["AirBookingData"]...).Endpoint()
	}
	
	var healthCheckEndpoint endpoint.Endpoint
	{
		healthCheckEndpoint = grpc1.NewClient(conn, "search.struct.Service", "HealthCheck", encodeHealthCheckRequest, decodeHealthCheckResponse, types.HealthCheckReply{}, options["HealthCheck"]...).Endpoint()
	}
	
	var airRulesEndpoint endpoint.Endpoint
	{
		airRulesEndpoint = grpc1.NewClient(conn, "search.struct.Service", "AirRules", encodeAirRulesRequest, decodeAirRulesResponse, types.AirRulesResponse{}, options["AirRules"]...).Endpoint()
	}
	
	return endpointSvc.Endpoints{
		CancelPNREndpoint:     cancelPNREndpoint,
		AirBookEndpoint:       eTIssueEndpoint,
		HealthCheckEndpoint:   healthCheckEndpoint,
		ReserveFlightEndpoint: reserveFlightEndpoint,
		SearchEndpoint:        searchEndpoint,
		AirRulesEndpoint:      airRulesEndpoint,
	}, nil
}

func decodeAirRulesResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*types.AirRulesResponse)
	if !ok {
		return nil, errors.New("conversion error")
	}
	fareRules := make([]*models.FareRules, 0)
	
	for _, fr := range resp.Data.Relationships.FareRules {
		rulesDetail := make([]*models.RuleDetails, 0)
		for _, rd := range fr.Relationships.RuleDetails {
			ruleDet := &models.RuleDetails{
				Category: rd.Category,
				Rules:    rd.Rules,
			}
			rulesDetail = append(rulesDetail, ruleDet)
		}
		fRule := &models.FareRules{
			ID:   fr.ID,
			Type: fr.Type,
			Attributes: &models.FareRulesAttributes{
				Airline:  fr.Attributes.Airline,
				CityPair: fr.Attributes.CityPair,
			},
			Relationships: &models.FareRulesRelationships{
				RuleDetails: rulesDetail,
			},
		}
		fareRules = append(fareRules, fRule)
	}
	
	return endpointSvc.AirRulesResponse{Response: &models.AirRulesResponse{
		StdError: nil,
		Data: &models.AirRulesData{
			ID:   resp.Data.ID,
			Type: resp.Data.Type,
			Attributes: &models.AirRulesAttributes{
				Success:  resp.Data.Attributes.Success,
				FareType: resp.Data.Attributes.FareType,
			},
			Relationships: &models.AirRulesRelationships{FareRules: fareRules},
		},
	},
	}, nil
}

func encodeAirRulesRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(endpointSvc.AirRulesRequest)
	return &types.AirRulesRequest{
		SessionId:      "",
		FareSourceCode: "",
		UniqueId:       "",
	}, nil
}

// encodeSearchRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Search request to a gRPC request.
func encodeSearchRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpointSvc.SearchRequest)
	return &types.SearchRequest{
		Origin:        req.Origin,
		Destination:   req.Destination,
		DepartureTime: req.DepartureTime,
		ReturnTime:    req.ReturnTime,
		AdultCount:    req.AdultCount,
		ChildCount:    req.ChildCount,
		InfantCount:   req.InfantCount,
	}, nil
}

// decodeSearchResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSearchResponse(_ context.Context, reply interface{}) (interface{}, error) {
	var (
		data    = make([]*models.Data, 0)
		Errs    = make([]*models.StdError, 0)
		resp, _ = reply.(*types.SearchReply)
		meta    *models.Meta
		link    *models.Link
	)
	
	for _, d := range resp.Data {
		_data := &models.Data{
			ID:   d.ID,
			Type: d.Type,
			Attributes: &models.Attributes{
				Currency: d.Attributes.Currency,
				Stop:     d.Attributes.Stop,
				Date: &models.Date{
					ArrivalDate:   d.Attributes.Date.ArivalDate,
					DepartureDate: d.Attributes.Date.DepartureDate,
				},
			},
			Relationships: &models.Relationships{
				OfferItems: []*models.OfferItems{
					{
						&models.OfferItemRelationship{
							Price: &models.Price{
								Total:    d.Relationships.OfferItems[0].Relationships.Price.Total,
								TotalTax: "",
							},
							PricePerAdult: &models.Price{
								Total:    d.Relationships.OfferItems[0].Relationships.PricePerAdult.Total,
								TotalTax: "",
							},
							PricePerChild: &models.Price{
								Total:    d.Relationships.OfferItems[0].Relationships.PricePerChild.Total,
								TotalTax: "",
							},
							PricePerInfant: &models.Price{
								Total:    d.Relationships.OfferItems[0].Relationships.PricePerInfant.Total,
								TotalTax: "",
							},
							Service: []*models.Service{
								{
									ServiceRelationships: &models.ServiceRelationships{
										Segment: []*models.Segment{
											{
												ID:   d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].ID,
												Type: d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Type,
												Attributes: &models.SegmentAtt{
													Duration:     d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.Duration,
													TravelClass:  d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.TravelClass,
													FareClass:    d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.FareClass,
													Number:       d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.Number,
													AircraftCode: d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.AircraftCode,
													Arrival: &models.DateDetail{
														IataCode: d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.Arrival.IataCode,
														At:       d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.Arrival.At,
													},
													Departure: &models.DateDetail{
														IataCode: d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.Departure.IataCode,
														At:       d.Relationships.OfferItems[0].Relationships.Service[0].Relationships.Segment[0].Attributes.Departure.At,
													},
												},
											},
										},
									},
								},
							},
						},
						&models.Link{Self: ""},
					},
				},
				Links: nil,
			},
		}
		data = append(data, _data)
	}
	
	for _, e := range resp.Errors {
		err := &models.StdError{
			ID:     e.ID,
			Status: e.Status,
			Code:   e.Code,
			Title:  e.Title,
			Detail: e.Detail,
		}
		Errs = append(Errs, err)
	}
	if resp.Links != nil {
		link = &models.Link{Self: resp.Links.Self}
	}
	
	if resp.Meta != nil {
		meta = &models.Meta{
			Pages: &models.Page{
				TotalPages: resp.Meta.Pages.TotalPages,
				From:       resp.Meta.Pages.From,
				LastPage:   resp.Meta.Pages.LastPage,
				To:         resp.Meta.Pages.To,
				Total:      resp.Meta.Pages.Total,
			},
			Currency:     "",
			Defaults:     nil,
			Dictionaries: nil,
		}
	}
	
	return endpointSvc.SearchResponse{
		Errors: Errs,
		Data:   data,
		Meta:   meta,
		Links:  link,
	}, nil
	
}

// encodeReserveFlightRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain AirBook request to a gRPC request.
func encodeReserveFlightRequest(_ context.Context, request interface{}) (interface{}, error) {
	passengers := make([]*types.AirTraveler, 0)
	req, ok := request.(endpointSvc.ReserveFlightRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	// Fill Passengers
	for _, p := range req.TravelerInfo.AirTravelers {
		passenger := &types.AirTraveler{
			DateOfBirth:   p.DateOfBirth,
			Gender:        p.Gender,
			PassengerType: p.PassengerType,
			PassengerName: &types.PassengerName{
				PassengerFirstName:        p.PassengerName.PassengerFirstName,
				PassengerLastName:         p.PassengerName.PassengerLastName,
				PassengerMiddleName:       p.PassengerName.PassengerMiddleName,
				PassengerTitle:            p.PassengerName.PassengerTitle,
				PassengerFirstNamePersian: p.PassengerName.PersianFirstName,
				PassengerLastNamePersian:  p.PassengerName.PersianLastName,
			},
			Passport: &types.Passport{
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
			passenger.Visa = &types.Visa{
				VisaType:    p.Visa.VisaType,
				VisaCountry: p.Visa.VisaCountry,
			}
		}
		passengers = append(passengers, passenger)
	}
	
	return &types.AirBookRequest{
		SessionId:       req.SessionId,
		FareSourceCode:  req.FareSourceCode,
		ClientUniqueId:  req.ClientUniqueId,
		MarkupForAdult:  float32(req.MarkupForAdult),
		MarkupForChild:  float32(req.MarkupForChild),
		MarkupForInfant: float32(req.MarkupForInfant),
		TravelerInfo: &types.TravelerInfo{
			PhoneNumber:  req.TravelerInfo.PhoneNumber,
			Email:        req.TravelerInfo.Email,
			AirTravelers: passengers,
		},
		TravelInfo: &types.TravelInfo{
			Origin:        req.TravelInfo.Origin,
			Destination:   req.TravelInfo.Destination,
			FlightClass:   req.TravelInfo.FlightClass,
			FlightNo:      req.TravelInfo.FlightNo,
			DepartureDate: req.TravelInfo.DepartureDate,
		},
	}, nil
	
}

// decodeReserveFlightResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeReserveFlightResponse(_ context.Context, reply interface{}) (interface{}, error) {
	var (
		Errors = make([]*models.StdError, 0)
	)
	
	resp, ok := reply.(*types.AirBookReply)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	if len(resp.StdError) > 0 {
		for _, v := range resp.StdError {
			err := &models.StdError{
				ID:     v.ID,
				Status: v.Status,
				Code:   v.Code,
				Title:  v.Title,
				Detail: v.Detail,
			}
			Errors = append(Errors, err)
		}
	}
	
	return endpointSvc.ReserveFlightResponse{
		Response: &models.ReserveResponse{
			Errors: Errors,
			Data: &models.ReserveData{
				ID:   resp.Data.ID,
				Type: resp.Data.Type,
				Attributes: &models.ReserveAttributes{
					Success:      resp.Data.Attributes.Success,
					TktTimeLimit: resp.Data.Attributes.TktTimeLimit,
					Category:     resp.Data.Attributes.Category,
					Status:       resp.Data.Attributes.Status,
					UniqueId:     resp.Data.Attributes.UniqueId,
				},
			},
			Meta:  nil,
			Links: nil,
		},
	}, nil
}

// encodeCancelPNRRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain CancelPNR request to a gRPC request.
func encodeCancelPNRRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(endpointSvc.CancelPNRRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return &types.CancelPNRRequest{
		PNR: req.PNR,
	}, nil
}

// decodeCancelPNRResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCancelPNRResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*types.CancelPNRReply)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return endpointSvc.CancelPNRResponse{
		Resp:  &models.CancelPNRResponse{Error: resp.Error},
		Error: nil,
	}, nil
}

// encodeETIssueRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain AirBookingData request to a gRPC request.
func encodeETIssueRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*endpointSvc.AirBookRequest)
	if !ok {
		return nil, errors.New("conversion error")
	}
	
	return &types.AirBookingDataRequest{
		UniqueId: req.UniqueId,
	}, nil
	
}

// decodeETIssueResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeETIssueResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*types.AirBookingDataReply)
	if !ok {
		return nil, errors.New("conversion error")
	}
	var (
		customerInfo = make([]*models.CustomerInfoes, 0)
		
		resItems = make([]*models.ReservationItems, 0)
		
		trDetails = make([]*models.TripDetailPtcFareBreakdowns, 0)
		
		bNotes = make([]*models.BookingNotes, 0)
		svc    = make([]*models.Services, 0)
	)
	
	// CustomerInfo
	for _, ci := range resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.CustomerInfoes {
		ETickets := make([]*models.Eticket, 0)
		for _, et := range ci.Relationships.ETicketNumbers {
			eTicket := &models.Eticket{
				ETicketNumber: et.ETicketNumber,
				DateOfIssue:   et.DateOfIssue,
				IsRefunded:    et.IsRefunded,
				AirlinePnr:    et.AirlinePnr,
				TotalRefund:   et.TotalRefund,
			}
			ETickets = append(ETickets, eTicket)
		}
		
		cInfo := &models.CustomerInfoes{
			ID:   ci.ID,
			Type: ci.Type,
			Attributes: &models.CustomerInfoesAttributes{
				ETickets: ci.Attributes.ETickets,
			},
			Relationships: &models.CustomerInfoesRelationships{
				Customer: &models.Customer{
					PassengerType:      ci.Relationships.Customer.PassengerType,
					PassportNumber:     ci.Relationships.Customer.PassportNumber,
					NationalId:         ci.Relationships.Customer.NationalId,
					DateOfBirth:        ci.Relationships.Customer.DateOfBirth,
					PassportExpireDate: ci.Relationships.Customer.PassportExpireDate,
					PaxName: &models.PassengerName{
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
	for _, ri := range resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ReservationItems {
		tStops := make([]*models.TechnicalStops, 0)
		for _, ts := range ri.Relationships.TechnicalStops {
			stop := &models.TechnicalStops{
				Data: &models.TechnicalStopsData{
					ArrivalAirport:    ts.Data.ArrivalAirport,
					ArrivalDateTime:   ts.Data.ArrivalDateTime,
					DepartureDateTime: ts.Data.DepartureDateTime,
				},
			}
			tStops = append(tStops, stop)
		}
		
		rItem := &models.ReservationItems{
			ID:   ri.ID,
			Type: ri.Type,
			Attributes: &models.ReservationItemsAttributes{
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
			Relationships: &models.ReservationItemsRelationships{
				TechnicalStops: tStops,
			},
		}
		resItems = append(resItems, rItem)
	}
	
	// TripDetail
	for _, bd := range resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.TripDetailPtcFareBreakdowns {
		bDown := &models.TripDetailPtcFareBreakdowns{
			Relationships: &models.TripDetailPtcFareBreakdownsRelationships{
				PassengerTypeQuantity: &models.PassengerTypeQuantity{
					PassengerType: bd.Relationships.PassengerTypeQuantity.PassengerType,
					Quantity:      bd.Relationships.PassengerTypeQuantity.Quantity,
				},
				TripDetailPassengerFare: &models.TripDetailPassengerFare{
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
	for _, bn := range resp.Data.Relationships.TravelItinerary.Relationships.BookingNotes {
		b := &models.BookingNotes{
			Note: bn.Note,
			Date: bn.Date,
		}
		bNotes = append(bNotes, b)
	}
	// Services
	for _, s := range resp.Data.Relationships.TravelItinerary.Relationships.Services {
		srv := &models.Services{
			PassengerFirstName:  s.PassengerFirstName,
			PassengerMiddleName: s.PassengerMiddleName,
			PassengerLastName:   s.PassengerLastName,
			Description:         s.Description,
			ServiceId:           s.ServiceId,
			Behavior:            s.Behavior,
			Amount:              s.Amount,
			Currency:            s.Currency,
		}
		svc = append(svc, srv)
	}
	
	return endpointSvc.AirBookResponse{
		Response: &models.AirBookingResponse{
			StdError: nil,
			Data: &models.AirBookingData{
				ID:   resp.Data.ID,
				Type: resp.Data.Type,
				Attributes: &models.AirBookingDataAttributes{
					UniqueId:      resp.Data.Attributes.UniqueId,
					BookedBy:      resp.Data.Attributes.BookedBy,
					OrderBy:       resp.Data.Attributes.OrderBy,
					ClientBalance: resp.Data.Attributes.ClientBalance,
					Success:       resp.Data.Attributes.Success,
					TktTimeLimit:  resp.Data.Attributes.TktTimeLimit,
					Category:      resp.Data.Attributes.Category,
					Status:        resp.Data.Attributes.Status,
					RefundMethod:  resp.Data.Attributes.RefundMethod,
				},
				Relationships: &models.AirBookingDataRelationships{TravelItinerary: &models.TravelItinerary{Relationships: &models.TravelItineraryRelationships{
					ItineraryInfo: &models.ItineraryInfo{Relationships: &models.ItineraryInfoRelationships{
						ItineraryPricing: &models.ItineraryPricing{
							BaseFare:        resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.BaseFare,
							ServiceTax:      resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.ServiceTax,
							TotalTax:        resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.TotalTax,
							TotalFare:       resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.TotalFare,
							TotalCommission: resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.TotalCommission,
							Currency:        resp.Data.Relationships.TravelItinerary.Relationships.ItineraryInfo.Relationships.ItineraryPricing.Currency,
						},
						CustomerInfoes:              customerInfo,
						ReservationItems:            resItems,
						TripDetailPtcFareBreakdowns: trDetails,
					}},
					BookingNotes: bNotes,
					Services:     svc,
				}}},
			},
			Meta:  nil,
			Links: nil,
		},
		Error: nil,
	}, nil
}

// encodeHealthCheckRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain HealthCheck request to a gRPC request.
func encodeHealthCheckRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return nil, nil
}

// decodeHealthCheckResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeHealthCheckResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*types.HealthCheckReply)
	if !ok {
		return nil, errors.New("conversion error")
	}
	return endpointSvc.HealthCheckResponse{Status: resp.Ok}, nil
}