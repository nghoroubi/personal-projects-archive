package models

// AirRevalidateRequest ...
type AirRevalidateRequest struct {
	Origin         string  `json:"origin"`
	Destination    string  `json:"destination"`
	DepartureTime  string  `json:"departure_time"`
	FlightClass    string  `json:"flight_class"`
	FlightNo       string   `json:"flight_no"`
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
	StdError []*StdError        `json:"errors"`
	Data     *AirRevalidateData `json:"data"`
	Meta     *Meta              `json:"meta"`
	Links    *Link              `json:"links"`
}

// AirRevalidateData ...
type AirRevalidateData struct {
	ID            string                     `json:"id"`
	Type          string                     `json:"type"`
	Attributes    *AirRevalidateAttributes   `json:"attributes"`
	Relationships *AirRevalidateRelationship `json:"relationships"`
}

// AirRevalidateAttributes ...
type AirRevalidateAttributes struct {
	ClientBalance float32 `json:"client_balance"`
}

// AirRevalidateRelationship ...
type AirRevalidateRelationship struct {
	Service         *AirRevalidateService `json:"service"`
	PricedItinerary *PricedItinerary      `json:"priced_itinerary"`
}

// AirRevalidateService ...
type AirRevalidateService struct {
	Data []*AirRevalidateServiceData `json:"data"`
}

// AirRevalidateServiceData ...
type AirRevalidateServiceData struct {
	PassengerFirstName  string `json:"passenger_first_name"`
	PassengerMiddleName string `json:"passenger_middle_name"`
	PassengerLastName   string `json:"passenger_last_name"`
	Description         string `json:"description"`
	ServiceId           string `json:"service_id"`
	Behavior            int32  `json:"behavior"`
	Amount              uint64 `json:"amount"`
	Currency            string `json:"currency"`
}

// PricedItinerary ...
type PricedItinerary struct {
	Data *PricedItineraryData `json:"data"`
}

// PricedItineraryData ...
type PricedItineraryData struct {
	ID            string                        `json:"id"`
	Type          string                        `json:"type"`
	Attributes    *PricedItineraryAttributes    `json:"attributes"`
	Relationships *PricedItineraryRelationships `json:"relationships"`
}

// PricedItineraryAttributes ...
type PricedItineraryAttributes struct {
	IsPassportMandatory          bool   `json:"is_passport_mandatory"`
	IsPassportIssueDateMandatory bool   `json:"is_passport_issue_date_mandatory"`
	DirectionInd                 int32  `json:"directionInd"`
	RefundMethod                 int32  `json:"refund_method"`
	ValidatingAirlineCode        string `json:"validating_airline_code"`
	FareSourceCode               string `json:"fare_source_code"`
}

// PricedItineraryRelationships ...
type PricedItineraryRelationships struct {
	AirItineraryPricingInfo  *AirItineraryPricingInfo  `json:"air_itinerary_pricing_info"`
	OriginDestinationOptions *OriginDestinationOptions `json:"origin_destination_options"`
}

// AirItineraryPricingInfo ...
type AirItineraryPricingInfo struct {
	Data *AirItineraryPricingInfoData `json:"data"`
}

// AirItineraryPricingInfoData ...
type AirItineraryPricingInfoData struct {
	ID            string                                `json:"id"`
	Type          string                                `json:"type"`
	Attributes    *AirItineraryPricingInfoAttributes    `json:"attributes"`
	Relationships *AirItineraryPricingInfoRelationships `json:"relationships"`
}

// AirItineraryPricingInfoAttributes ...
type AirItineraryPricingInfoAttributes struct {
	FareType        int32   `json:"fare_type"`
	BaseFare        float32 `json:"base_fare"`
	TotalFare       float32 `json:"total_fare"`
	TotalCommission float32 `json:"total_commission"`
	TotalTax        float32 `json:"total_tax"`
	ServiceTax      float32 `json:"service_tax"`
	Currency        string  `json:"currency"`
}

// AirItineraryPricingInfoRelationships ...
type AirItineraryPricingInfoRelationships struct {
	PtcFareBreakdown *PtcFareBreakdown `json:"ptc_fare_breakdown"`
	FareInfoes       []string          `json:"fare_infoes"`
}

// PtcFareBreakdown ...
type PtcFareBreakdown struct {
	Data []*PtcFareBreakdownData `json:"data"`
}

// PtcFareBreakdownData ...
type PtcFareBreakdownData struct {
	ID            string                         `json:"id"`
	Type          string                         `json:"type"`
	Attributes    *PtcFareBreakdownAttributes    `json:"attributes"`
	Relationships *PtcFareBreakdownRelationships `json:"relationships"`
}

// PtcFareBreakdownAttributes ...
type PtcFareBreakdownAttributes struct {
	BaseFare   float32 `json:"base_fare"`
	TotalFare  float32 `json:"total_fare"`
	Commission float32 `json:"total_commission"`
	ServiceTax float32 `json:"service_tax"`
	Currency   string  `json:"currency"`
}

// PtcFareBreakdownRelationships ...
type PtcFareBreakdownRelationships struct {
	Taxes *Tax `json:"taxes"`
}

// Tax ...
type Tax struct {
	Data []*TaxData `json:"data"`
}

// TaxData ...
type TaxData struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"Currency"`
}

// OriginDestinationOptions ...
type OriginDestinationOptions struct {
	Data []*OriginDestinationOptionsData `json:"data"`
}

// OriginDestinationOptionsData ...
type OriginDestinationOptionsData struct {
	ID            string                                 `json:"id"`
	Type          string                                 `json:"type"`
	Attributes    *OriginDestinationOptionsAttributes    `json:"attributes"`
	Relationships *OriginDestinationOptionsRelationships `json:"relationships"`
}

// OriginDestinationOptionsAttributes ...
type OriginDestinationOptionsAttributes struct {
	JourneyDurationPerMinute int32 `json:"journey_duration_per_minute"`
	ConnectionTimePerMinute  int32 `json:"connection_time_per_minute"`
}

// OriginDestinationOptionsRelationships ...
type OriginDestinationOptionsRelationships struct {
	FlightSegments *FlightSegments `json:"flight_segments"`
}

// FlightSegments ...
type FlightSegments struct {
	Data []*FlightSegmentData `json:"data"`
}

// FlightSegmentData ...
type FlightSegmentData struct {
	ID            string                      `json:"id"`
	Type          string                      `json:"type"`
	Attributes    *FlightSegmentsAttributes   `json:"attributes"`
	Relationships *FlightSegmentsRelationship `json:"relationships"`
}

// FlightSegmentsAttributes ...
type FlightSegmentsAttributes struct {
	DepartureDateTime            string `json:"departure_dateTime"`
	ArrivalDateTime              string `json:"arrival_dateTime"`
	StopQuantity                 int32  `json:"stop_quantity"`
	FlightNumber                 string `json:"flight_number"`
	ResBookDesignCode             string `json:"res_book_design_code"`
	JourneyDuration              string `json:"journey_duration"`
	JourneyDurationPerMinute     int32  `json:"journey_duration_per_minute"`
	ConnectionTimePerMinute      int32  `json:"connection_time_per_minute"`
	DepartureAirportLocationCode string `json:"departure_airport_location_code"`
	ArrivalAirportLocationCode   string `json:"arrival_airport_location_code"`
	MarketingAirlineCode         string `json:"marketing_airline_code"`
	CabinClassCode               int32  `json:"cabin_class_code"`
	SeatsRemaining               int32  `json:"seats_remaining"`
	IsCharter                    bool   `json:"is_charter"`
	IsReturn                     bool   `json:"is_return"`
	Baggage                      string `json:"baggage"`
}

// FlightSegmentsRelationship ...
type FlightSegmentsRelationship struct {
	OperatingAirline *Operating        `json:"operating_airline"`
	TechnicalStops   []*TechnicalStops `json:"technical_stops"`
}
