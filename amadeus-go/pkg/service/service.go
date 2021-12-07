package service

import (
	"amadeus-go/config/env"
	"amadeus-go/helpers"
	"amadeus-go/models"
	"amadeus-go/pkg/repository"
	"context"
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Service describes the service.
type Service interface {
	HealthCheck(ctx context.Context) bool
	GenerateControlNumber(refNum string) (string, error)
	AirRules(ctx context.Context, req *models.AirRulesRequest) *models.AirRulesResponse
	Search(ctx context.Context, req *models.SearchRequest) *models.SearchResponse
	ReserveFlight(ctx context.Context, req *models.ReserveRequest) (*models.ReserveResponse, error)
	CancelPNR(ctx context.Context, req *models.CancelPNRRequest) (*models.CancelPNRResponse, error)
	AirBook(ctx context.Context, req *models.AirBookingRequest) (*models.AirBookingResponse, error)
}

// ServiceConfig , common configuration of amadeus api,
// which most of methods would use
type ServiceConfig struct {
	Version              int
	OnlyDirectFlights    bool
	OnlyAvailableFlights bool
	Currency             string
	Timestamp            string
	RefundMode           string
	Username             string
	Password             string
	ProviderType         string
	TicketingType        repository.TicketType
}

// Implementer of Service Interface
type impl struct {
	repo repository.IAmadeus
	*ServiceConfig
}

// newAmadeusService returns a naive, stateless implementation of Service.
func newAmadeusService(repo repository.IAmadeus) Service {
	return &impl{
		repo: repo,
		ServiceConfig: &ServiceConfig{
			Version:              env.GetInt("api.version"),
			OnlyDirectFlights:    env.GetBool("api.config.search.direct_flights_only"),
			OnlyAvailableFlights: env.GetBool("api.config.search.available_flights_only"),
			Currency:             env.GetString("api.currency"),
			Timestamp:            env.GetString("api.timestamp"),
			RefundMode:           env.GetString("api.config.search.refund_mode"),
			ProviderType:         env.GetString("api.config.search.provider_type"),
			TicketingType:        getTicketType(env.GetString("api.config.book.ticketing_type")),
			Username:             env.GetString("api.auth.user"),
			Password:             env.GetString("api.auth.password"),
		},
	}
}

// New returns a IranairtourService with all of the expected middleware wired in.
func New(middleware []Middleware, amadeus repository.IAmadeus) Service {
	var svc = newAmadeusService(amadeus)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

// Search ...
func (svc *impl) Search(ctx context.Context, req *models.SearchRequest) *models.SearchResponse {
	var (
		data                  = make([]*models.Data, 0)
		Errs                  = make([]*models.StdError, 0)
		AirTravelerAvail      = make([]*repository.TravelerInformationType, 0)
		originDestinationInfo = make([]*repository.OriginDestinationInformation, 0)
	)
	// Check if there any response for the same request has already cached, if not, do a fresh search
	if exist := repository.Cache().Exists(key(req)).Val(); exist == 1 {
		var cachedData []*models.Data
		err := repository.Cache().ZRange(key(req), 0, 100).ScanSlice(&cachedData)
		if err != nil {
			log.Println(err)
			return &models.SearchResponse{
				Data: data,
				Errors: []*models.StdError{
					{
						Code:   err.Error(),
						Title:  err.Error(),
						Detail: err.Error(),
					}},
			}
		}
		if len(cachedData) == 1 && cachedData[0].Attributes == nil {
			return &models.SearchResponse{
				Data: data,
				Errors: []*models.StdError{
					{
						Code:   "Route Not Found in this Airline",
						Title:  "Route Not Found in this Airline",
						Detail: "Route Not Found in this Airline",
					}},
			}
		}
		return &models.SearchResponse{
			Data:   cachedData,
			Errors: Errs,
			Meta: &models.Meta{
				Pages: &models.Page{
					TotalPages: 1,
					From:       1,
					LastPage:   1,
					PerPage:    10,
					To:         1,
					Total:      int32(len(cachedData)),
				},
				Currency:     "IRR",
				Defaults:     nil,
				Dictionaries: nil,
			},
			Links: nil,
		}
	} else {
		AirTravelerAvail = travelerInfoConstructor(req)
		// Making departure time format to amadeus timestamp format
		t, err := time.Parse("2006-01-02T15:04:05", req.DepartureTime)
		if err != nil {
			return &models.SearchResponse{
				Data: data,
				Errors: []*models.StdError{
					{
						Code:   err.Error(),
						Title:  err.Error(),
						Detail: err.Error(),
					},
				},
			}
		}
		depTime := t.Format("2006-01-02T15:04:05")
		originDestinationInfo = append(originDestinationInfo, &repository.OriginDestinationInformation{
			OriginDestinationInformationType: &repository.OriginDestinationInformationType{
				TravelDateTimeType: &repository.TravelDateTimeType{
					DepartureDateTime: depTime,
				},
				OriginLocation: &repository.Location{
					LocationCode:        req.Origin,
					MultiAirportCityInd: env.GetBool("api.config.search.multi_airport_city"),
				},
				DestinationLocation: &repository.Location{
					LocationCode:        req.Destination,
					MultiAirportCityInd: env.GetBool("api.config.search.multi_airport_city"),
				},
			},
		})
		
		// Check the RoundTrip and/or multi-leg flights and construct extra segments of the flight
		if req.ReturnTime != "" {
			rt, err := time.Parse("2006-01-02T15:04:05", req.ReturnTime)
			if err != nil {
				return &models.SearchResponse{
					Data: data,
					Errors: []*models.StdError{
						{
							Code:   err.Error(),
							Title:  err.Error(),
							Detail: err.Error(),
						},
					},
				}
			}
			retTime := rt.Format("2006-01-02T15:04:05")
			returnSegment := &repository.OriginDestinationInformation{
				OriginDestinationInformationType: &repository.OriginDestinationInformationType{
					TravelDateTimeType: &repository.TravelDateTimeType{
						DepartureDateTime: retTime,
					},
					OriginLocation: &repository.Location{
						LocationCode:        req.Destination,
						MultiAirportCityInd: env.GetBool("api.config.search.multi_airport_city"),
					},
					DestinationLocation: &repository.Location{
						LocationCode:        req.Origin,
						MultiAirportCityInd: env.GetBool("api.config.search.multi_airport_city"),
					},
				},
			}
			originDestinationInfo = append(originDestinationInfo, returnSegment)
		}
		
		resp, err := svc.repo.SearchFlight(&repository.SearchFlightRequest{
			OTAAirLowFareSearchRQ: &repository.OTAAirLowFareSearchRQ{
				TimeStamp:            svc.Timestamp,
				Version:              svc.Version,
				MaxResponses:         uint(env.GetInt("api.config.search.maximum_response_count")),
				DirectFlightsOnly:    svc.OnlyDirectFlights,
				AvailableFlightsOnly: svc.OnlyAvailableFlights,
				AdvanceSearchInfo: &repository.AdvanceSearchInfo{
					NumberOfRecommendation: uint(env.GetInt("api.config.search.maximum_response_count")),
					Currency:               svc.Currency,
				},
				TravelerInfoSummary: &repository.TravelerInfoSummary{
					TravelerInfoSummaryType: &repository.TravelerInfoSummaryType{
						AirTravelerAvail: AirTravelerAvail,
					},
				},
				RefundableType:               getRefundableType(svc.RefundMode),
				ProviderType:                 getProviderType(svc.ProviderType),
				OriginDestinationInformation: originDestinationInfo,
			},
		})
		if err != nil {
			return &models.SearchResponse{
				Data: data,
				Errors: []*models.StdError{
					{
						Code:   err.Error(),
						Title:  err.Error(),
						Detail: err.Error(),
					},
				},
			}
		}
		
		if respErr := resp.OTAAirLowFareSearchRS.Errors; respErr != nil {
			if len(respErr.Error) > 0 {
				return &models.SearchResponse{
					Errors: []*models.StdError{
						{
							ID:     respErr.Error[0].Code,
							Status: respErr.Error[0].Status,
							Code:   respErr.Error[0].ErrorCode,
							Title:  respErr.Error[0].ShortText,
							Detail: respErr.Error[0].Value,
						},
					},
				}
			}
		}
		
		// Splitting all of originDestinationOptions of each pricedItinerary in a map
		// which have capacity equal to flight legs
		for _, itin := range resp.OTAAirLowFareSearchRS.PricedItineraries.PricedItinerary {
			originDestinationOptions := helpers.SeparateODOByDirection(itin.AirItinerary.OriginDestinationOptions.OriginDestinationOption)
			
			// each index of this 2d array will be send to parser, in a loop
			combinedODs := helpers.DetectCombinableOriginDestinations(originDestinationOptions, itin.AirItinerary.OriginDestinationCombinations.OriginDestinationCombination)
			
			// send an array which contains one or more originDestinationOption to the parser
			for _, CODO := range combinedODs {
				d := helpers.ParseSearchResponse(itin, CODO.OriginDestinationOptions, resp.OTAAirLowFareSearchRS.PricedItineraries, CODO.CombinationId)
				data = append(data, d)
			}
		}
		
		log.Printf("\nresponse count : %v Flight(s)\n", len(data))
		if err = cacheTheSearchResult(req, data); err != nil {
			return &models.SearchResponse{
				Data: data,
				Errors: []*models.StdError{
					{
						Code:   err.Error(),
						Title:  err.Error(),
						Detail: err.Error(),
					}},
			}
		}
		return &models.SearchResponse{
			Data:   data,
			Errors: Errs,
			Meta: &models.Meta{
				Pages: &models.Page{
					TotalPages: 1,
					From:       1,
					LastPage:   1,
					PerPage:    10,
					To:         1,
					Total:      int32(len(data)),
				},
				Currency:     "IRR",
				Defaults:     nil,
				Dictionaries: nil,
			},
			Links: nil,
		}
	}
	
}

// ReserveFlight ...
func (svc *impl) ReserveFlight(ctx context.Context, req *models.ReserveRequest) (*models.ReserveResponse, error) {
	// Construct and fill the passengers info
	bookRQ := svc.constructReserveRQ(req)
	
	// Book - temp reservation
	response, err := svc.repo.BookFlight(bookRQ)
	if err != nil {
		return &models.ReserveResponse{
			Errors: []*models.StdError{
				{
					Status: "failed",
					Code:   "reserve_error",
					Title:  err.Error(),
					Detail: err.Error(),
				},},
			Data:  nil,
			Meta:  nil,
			Links: nil,
		}, err
	}
	
	if er := response.OTAAirBookRS.Errors; er != nil {
		// TODO : handle the error in a separate function
		if er.Error[0].ShortText == "SearchedFlightRecommendations can not be null or empty" {
			return &models.ReserveResponse{
				Errors: []*models.StdError{{
					ID:     "401",
					Status: "forbidden",
					Code:   er.Error[0].Code,
					Title:  er.Error[0].ShortText,
					Detail: er.Error[0].ShortText,
				}},
			}, errors.New("session expired,please search again")
		}
		return &models.ReserveResponse{
			Errors: []*models.StdError{{
				ID:     "",
				Status: "failed",
				Code:   response.OTAAirBookRS.Errors.Error[0].Code,
				Title:  response.OTAAirBookRS.Errors.Error[0].ShortText,
				Detail: response.OTAAirBookRS.Errors.Error[0].ShortText,
			}},
		}, err
	}
	
	// Cache in successful case
	if pnr := response.OTAAirBookRS.AirReservation.BookingReferenceID[0].IDContext; pnr != "" {
		if cacheErr := cacheTheReserveRequest(pnr, response.OTAAirBookRS.AirReservation.TravelerInfo.AirTraveler[0].PersonName.Surname); cacheErr != nil {
			return &models.ReserveResponse{
				Errors: []*models.StdError{
					{
						ID:     "",
						Status: "failed to cache",
						Code:   cacheErr.Error(),
						Title:  cacheErr.Error(),
						Detail: cacheErr.Error(),
					},
				},
			}, cacheErr
		}
	}
	
	// caching book request for using in book request in unsuccessful case
	if errs := response.OTAAirBookRS.Errors; errs != nil && len(errs.Error) > 0 && priceHasChanged(errs.Error) {
		if cacheErr := cacheTheReserveRequest(response.OTAAirBookRS.ReferenceNumber, req.TravelerInfo.AirTravelers[0].PassengerName.PassengerLastName); cacheErr != nil {
			return &models.ReserveResponse{
				Errors: []*models.StdError{
					{
						ID:     "",
						Status: "failed to cache",
						Code:   cacheErr.Error(),
						Title:  cacheErr.Error(),
						Detail: cacheErr.Error(),
					},
				},
			}, cacheErr
		}
		return &models.ReserveResponse{
				Errors: []*models.StdError{
					{
						ID:     response.OTAAirBookRS.ReferenceNumber,
						Status: "priced changed",
						Code:   "A000",
						Title:  "do booking progress with given reference id",
						Detail: fmt.Sprintf("new price:%f", response.OTAAirBookRS.NewPrice.ItinTotalFare.TotalFare.Amount),
					},
				},
				Data:  nil,
				Meta:  nil,
				Links: nil,
			},
			errors.New("price changed")
	}
	
	// return successful response
	return &models.ReserveResponse{
		Errors: nil,
		Data: &models.ReserveData{
			ID:   "",
			Type: env.GetString("service.name"),
			Attributes: &models.ReserveAttributes{
				Success:      response.OTAAirBookRS.Success != nil,
				TktTimeLimit: response.OTAAirBookRS.AirReservation.Ticketing[0].TicketTimeLimit,
				Status:       1,
				UniqueId:     response.OTAAirBookRS.AirReservation.BookingReferenceID[0].IDContext,
			},
		},
	}, nil
}

// AirBook ...
func (svc *impl) AirBook(ctx context.Context, req *models.AirBookingRequest) (*models.AirBookingResponse, error) {
	var (
		strKey         string
		confirmSurname string
	)
	if req.UniqueId != "" {
		strKey = req.UniqueId
	} else {
		strKey = req.SessionId
	}
	
	// Get travelers from cache
	if exist := repository.Cache().Exists(strKey).Val(); exist == 1 {
		confirmSurname = repository.Cache().Get(strKey).Val()
		if confirmSurname == "" {
			log.Println("error getting required value from cache")
			
			return &models.AirBookingResponse{
				StdError: []*models.StdError{
					{
						ID:     "",
						Status: "failed create ticket",
						Code:   "",
						Title:  "failed create ticket",
						Detail: "error getting required value from cache",
					},
				},
			}, errors.New("error getting required value from cache")
		}
	}
	
	// Get response from webService
	request := &repository.CreateTicket{OTAAirBookRQ: &repository.OTAAirBookRQ{
		Version:         svc.Version,
		ReferenceNumber: "",
		ControlNumber:   "",
		TimeStamp:       svc.Timestamp,
		BookingReferenceID: &repository.UniqueIDType{
			IDContext: req.UniqueId,
		},
		TravelerInfo: &repository.TravelerInfoType{
			AirTraveler: []*repository.AirTraveler{
				{
					
					AirTravelerType: &repository.AirTravelerType{
						PersonName: &repository.PersonNameType{
							Surname: confirmSurname,
						},
					},
				},
			},
		},
	}}
	
	resp, err := svc.repo.CreateTicket(request)
	if err != nil {
		return &models.AirBookingResponse{
			StdError: []*models.StdError{
				{
					ID:     "",
					Status: "failed create ticket",
					Code:   "",
					Title:  "failed create ticket",
					Detail: err.Error(),
				},
			},
		}, err
	}
	if err := resp.OTAAirBookRS.Errors; err != nil && len(err.Error) > 0 {
		return &models.AirBookingResponse{
			StdError: []*models.StdError{
				{
					ID:     "",
					Status: "failed create ticket",
					Code:   "",
					Title:  "failed create ticket",
					Detail: err.Error[0].ShortText,
				},
			},
		}, errors.New(err.Error[0].ShortText)
	}
	var (
		customersInfo                = make([]*models.CustomerInfoes, 0)
		eTickets                     = make([]*models.Eticket, 0)
		tripDetailsPTCFareBreakDowns = make([]*models.TripDetailPtcFareBreakdowns, 0)
		baseFare                     = resp.OTAAirBookRS.AirReservation.PriceInfo.ItinTotalFare.BaseFare.Amount
		totalFare                    = resp.OTAAirBookRS.AirReservation.PriceInfo.ItinTotalFare.TotalFare.Amount
	)
	
	for _, ptc := range resp.OTAAirBookRS.AirReservation.PriceInfo.PTCFareBreakdowns.PTCFareBreakdown {
		p := &models.TripDetailPtcFareBreakdowns{Relationships: &models.TripDetailPtcFareBreakdownsRelationships{
			PassengerTypeQuantity: &models.PassengerTypeQuantity{
				PassengerType: getIntPassengerType(ptc.PassengerTypeQuantity.Code),
				Quantity:      ptc.PassengerTypeQuantity.Quantity,
			},
			TripDetailPassengerFare: &models.TripDetailPassengerFare{
				BaseFare:   float32(ptc.PassengerFare.BaseFare.Amount),
				ServiceTax: 0,
				Tax:        0,
				TotalFare:  float32(ptc.PassengerFare.TotalFare.Amount),
				Commission: 0,
				Currency:   svc.Currency,
			},
		}}
		tripDetailsPTCFareBreakDowns = append(tripDetailsPTCFareBreakDowns, p)
	}
	
	for pIndex, v := range resp.OTAAirBookRS.AirReservation.TravelerInfo.AirTraveler {
		for _, et := range v.ETicketDocuments.ETicketInfo {
			t := &models.Eticket{
				ETicketNumber: et.TicketNumber,
				DateOfIssue:   et.TicketingDate.Format("2006-01-02T15:04:05"),
				IsRefunded:    false,
				AirlinePnr:    req.UniqueId,
				TotalRefund:   0,
			}
			eTickets = append(eTickets, t)
		}
		c := &models.CustomerInfoes{
			ID:         "",
			Type:       env.GetString("service.name"),
			Attributes: &models.CustomerInfoesAttributes{ETickets: v.ETicketNumber},
			Relationships: &models.CustomerInfoesRelationships{
				Customer: &models.Customer{
					PassengerType: tripDetailsPTCFareBreakDowns[pIndex].Relationships.PassengerTypeQuantity.PassengerType,
					PaxName: &models.PassengerName{
						PassengerFirstName: v.PersonName.GivenName,
						PassengerLastName:  v.PersonName.Surname,
					},
					PassportNumber:     getPassportNumber(v.Document),
					PassportExpireDate: getPassportExpireDate(v.Document),
					NationalId:         getNationalID(v.Document),
				},
				ETicketNumbers: eTickets,
			},
		}
		customersInfo = append(customersInfo, c)
	}
	
	return &models.AirBookingResponse{
		StdError: nil,
		Data: &models.AirBookingData{
			ID:   strconv.Itoa(int(resp.OTAAirBookRS.SequenceNmbr)),
			Type: env.GetString("service.name"),
			Attributes: &models.AirBookingDataAttributes{
				UniqueId:      req.UniqueId,
				BookedBy:      "",
				OrderBy:       "",
				ClientBalance: 0,
				Success:       resp.OTAAirBookRS.Errors != nil && resp.OTAAirBookRS.Warnings != nil,
				IssueDate:     time.Now().Format("2006-01-02T15:04:05"),
				TktTimeLimit:  resp.OTAAirBookRS.AirReservation.Ticketing[0].TicketTimeLimit,
				Category:      21,
				Status:        1,
				RefundMethod:  0,
			},
			Relationships: &models.AirBookingDataRelationships{
				TravelItinerary: &models.TravelItinerary{
					Relationships: &models.TravelItineraryRelationships{
						ItineraryInfo: &models.ItineraryInfo{
							Relationships: &models.ItineraryInfoRelationships{
								ItineraryPricing: &models.ItineraryPricing{
									BaseFare:  float32(baseFare),
									TotalFare: float32(totalFare),
									Currency:  svc.Currency,
								},
								CustomerInfoes:              customersInfo,
								ReservationItems:            nil,
								TripDetailPtcFareBreakdowns: tripDetailsPTCFareBreakDowns,
							}},
						BookingNotes: nil,
						Services:     nil,
					}}},
		},
	}, nil
}

// CancelPNR ...
func (svc *impl) CancelPNR(ctx context.Context, req *models.CancelPNRRequest) (resp *models.CancelPNRResponse, err error) {
	// TODO implement the business logic of CancelPNR
	return resp, err
}

// AirRules ...
func (svc *impl) AirRules(ctx context.Context, req *models.AirRulesRequest) *models.AirRulesResponse {
	// Extract required params from request
	recommendationID, combinationID := getBookingIDs(req.FareSourceCode)
	
	// Call method
	resp, err := svc.repo.GetFlightRules(&repository.GetFlightRules{
		XMLName: xml.Name{},
		OTAAirRulesRQ: &repository.OTAAirRulesRQ{
			GenericFlightRQ:     &repository.GenericFlightRQ{RecommendationID: recommendationID, CombinationID: combinationID},
			MiniRuleEnabled:     1,
			PriceMessageEnabled: 0,
			FlightRuleEnabled:   1,
			PassengerType:       "ADT",
			TimeStamp:           svc.Timestamp,
			Version:             float64(svc.Version),
		},
	})
	if err != nil {
		return &models.AirRulesResponse{
			StdError: []*models.StdError{{
				ID:     "",
				Status: "failed",
				Code:   "get_rules",
				Title:  "get rules failed",
				Detail: "get rules failed",
			}},
		}
	}
	
	// Check for API error (extra errors which API server returns in failing cases
	if e := resp.OTAAirRulesRS.Errors; e != nil && len(e.Error) > 0 {
		return &models.AirRulesResponse{
			StdError: []*models.StdError{
				{
					ID:     e.Error[0].Code,
					Status: e.Error[0].Status,
					Code:   e.Error[0].Code,
					Title:  e.Error[0].ShortText,
					Detail: e.Error[0].ShortText,
				},
			},
		}
	}
	
	// In successful case
	fareRules := parseRuleResponse(resp)
	return &models.AirRulesResponse{
		StdError: nil,
		Data: &models.AirRulesData{
			ID:   "",
			Type: "iranAirTour",
			Attributes: &models.AirRulesAttributes{
				Success:  true,
				FareType: 0,
			},
			Relationships: &models.AirRulesRelationships{FareRules: fareRules},
		},
	}
}

// HealthCheck ....
func (svc *impl) HealthCheck(ctx context.Context) (status bool) {
	return true
}

// create repository's required type for book request
func (svc *impl) constructReserveRQ(req *models.ReserveRequest) *repository.BookFlight {
	recommendationID, combinationID := getBookingIDs(req.FareSourceCode)
	return &repository.BookFlight{
		OTAAirBookRQ: &repository.OTAAirBookRQ{
			GenericFlightRQ: &repository.GenericFlightRQ{
				RecommendationID: recommendationID,
				CombinationID:    combinationID,
			},
			Version: svc.Version,
			TravelerInfo: &repository.TravelerInfoType{
				AirTraveler: passengerConstructor(req),
			},
			Fulfillment: &repository.Fulfillment{
				PaymentDetails: &repository.ArrayOfPaymentDetailType{
					PaymentDetail: []*repository.PaymentDetailType{
						{
							PaymentFormType: &repository.PaymentFormType{
								PaymentType:   repository.PaymentTypesNone,
								PaymentFPType: repository.PaymentFPTypesFPCA,
							},
						},
					},
				},
			},
			Ticketing: []*repository.BookTicketingInfo{
				{
					TicketingInfoType: &repository.TicketingInfoType{
						TicketType: svc.TicketingType,
					},
				},
			},
		},
	}
}

// /////////////////// Service HELPERS ////////////////////

// GenerateControlNumber
// generating hash of reference Id for acknowledge in changed price book request
func (svc *impl) GenerateControlNumber(refNum string) (string, error) {
	/*
		Reference number: fef578e5-90d1-46cc-a19f-0592ba1e3515
		Example web service security number provided by Amadeus: 1234
		Concatenated string: fef578e5-90d1-46cc-a19f-0592ba1e35151234
		Generated control number: KWHr9o+yfaATmTnAQBkPMmlpif0=
	*/
	source := fmt.Sprintf("%s%s", refNum, svc.Password)
	
	hash := sha1.New()
	hash.Write([]byte(source))
	h := hash.Sum(nil)
	return fmt.Sprintf("%x", h), nil
}

// gets FlightRules response and parses to the native structure
func parseRuleResponse(resp *repository.GetFlightRulesResponse) []*models.FareRules {
	var (
		ruleDetails = make([]*models.RuleDetails, 0)
		fareRules   = make([]*models.FareRules, 0)
	)
	
	for _, rule := range resp.OTAAirRulesRS.FareRuleResponseInfo.FareRuleInfo {
		fr := &models.FareRules{
			ID:   rule.FareReference[0].Value,
			Type: env.GetString("service.name"),
			Attributes: &models.FareRulesAttributes{
				Airline: rule.MarketingAirline[0].Code,
				CityPair: fmt.Sprintf("%s-%s",
					rule.DepartureAirport.LocationCode,
					rule.ArrivalAirport.LocationCode),
			},
			Relationships: &models.FareRulesRelationships{RuleDetails: nil},
		}
		
		for _, txt := range rule.FareRules.SubSection {
			rd := &models.RuleDetails{
				Category: txt.SubTitle,
				Rules:    txt.Paragraph.Text,
			}
			ruleDetails = append(ruleDetails, rd)
		}
		
		fr.Relationships.RuleDetails = ruleDetails
		fareRules = append(fareRules, fr)
	}
	return fareRules
}

// get integer type for each passenger
func getIntPassengerType(s string) int32 {
	switch s {
	case "CHD":
		return 2
	case "INF":
		return 3
	default:
		return 1
	}
}

// Checks the errors of API response and returns true if the price has changed
func priceHasChanged(errs []*repository.ErrorType) bool {
	for _, err := range errs {
		if strings.EqualFold(strings.ToLower(err.ShortText), strings.ToLower(env.GetString("api.critical_errors.price_changed"))) {
			return true
		}
	}
	return false
}

// sample errors and warnings
// <Error Type="EpowerInternalError" ShortText="Prices have changed" Code="A000" NodeList="BOOK_WITH_PRICECHANGE" />