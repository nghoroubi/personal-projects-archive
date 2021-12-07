package models

// AirBookingRequest ...
type AirBookingRequest struct {
	SessionId string `validate:"required" json:"session_id"`
	UniqueId  string `validate:"required" json:"unique_id"`
	Email     string `validate:"required" json:"email"`
}

// AirBookingResponse ...
type AirBookingResponse struct {
	StdError []*StdError     `json:"errors"`
	Data     *AirBookingData `json:"data"`
	Meta     *Meta           `json:"meta"`
	Links    *Link           `json:"links"`
}

// AirBookingData ...
type AirBookingData struct {
	ID            string                       `json:"id"`
	Type          string                       `json:"type"`
	Attributes    *AirBookingDataAttributes    `json:"attributes"`
	Relationships *AirBookingDataRelationships `json:"relationships"`
}

// AirBookingDataAttributes ...
type AirBookingDataAttributes struct {
	UniqueId      string  `json:"unique_id"`
	BookedBy      string  `json:"booked_by"`
	OrderBy       string  `json:"order_by"`
	ClientBalance float32 `json:"client_balance"`
	Success       bool    `json:"success"`
	IssueDate     string  `json:"issue_date"`
	TktTimeLimit  string  `json:"tkt_time_limit"`
	Category      int32   `json:"category"`
	Status        int32   `json:"status"`
	RefundMethod  int32   `json:"refund_method"`
}

// AirBookingDataRelationships ...
type AirBookingDataRelationships struct {
	TravelItinerary *TravelItinerary
}

// TravelItinerary ...
type TravelItinerary struct {
	Relationships *TravelItineraryRelationships `json:"relationships"`
}

// TravelItineraryRelationships ...
type TravelItineraryRelationships struct {
	ItineraryInfo *ItineraryInfo  `json:"itinerary_info"`
	BookingNotes  []*BookingNotes `json:"booking_notes"`
	Services      []*Services     `json:"services"`
}

// BookingNotes ...
type BookingNotes struct {
	Note string `json:"note"`
	Date string `json:"data"`
}

// ItineraryInfo ...
type ItineraryInfo struct {
	Relationships *ItineraryInfoRelationships `json:"relationships"`
}

// ItineraryInfoRelationships ...
type ItineraryInfoRelationships struct {
	ItineraryPricing            *ItineraryPricing
	CustomerInfoes              []*CustomerInfoes
	ReservationItems            []*ReservationItems
	TripDetailPtcFareBreakdowns []*TripDetailPtcFareBreakdowns
}

// ItineraryPricing ...
type ItineraryPricing struct {
	BaseFare        float32
	ServiceTax      float32
	TotalTax        float32
	TotalFare       float32
	TotalCommission float32
	Currency        string
}

// CustomerInfoes ...
type CustomerInfoes struct {
	ID            string                       `json:"id"`
	Type          string                       `json:"Type"`
	Attributes    *CustomerInfoesAttributes    `json:"attributes"`
	Relationships *CustomerInfoesRelationships `json:"relationships"`
}

// CustomerInfoesAttributes ...
type CustomerInfoesAttributes struct {
	ETickets string
}

// CustomerInfoesRelationships ...
type CustomerInfoesRelationships struct {
	Customer       *Customer
	ETicketNumbers []*Eticket
}

// Customer ...
type Customer struct {
	PassengerType      int32 `validate:"required"` // 1 => Adult , 2 => Child  3 => Infant
	PassportNumber     string
	NationalId         string
	DateOfBirth        string
	PassportExpireDate string
	PaxName            *PassengerName
}

// ReservationItems ...
type ReservationItems struct {
	ID            string                         `json:"id"`
	Type          string                         `json:"type"`
	Attributes    *ReservationItemsAttributes    `json:"attributes"`
	Relationships *ReservationItemsRelationships `json:"relationships"`
}

// ReservationItemsAttributes ...
type ReservationItemsAttributes struct {
	AirEquipmentType             string
	AirlinePnr                   string
	ArrivalAirportLocationCode   string
	ArrivalDateTime              string
	ArrivalTerminal              string
	Baggage                      string
	DepartureAirportLocationCode string
	DepartureDateTime            string
	DepartureTerminal            string
	FlightNumber                 string
	JourneyDuration              string
	JourneyDurationPerMinute     int32
	MarketingAirlineCode         string
	OperatingAirlineCode         string
	ResBookDesignCode            string
	StopQuantity                 int32
	IsCharter                    bool
}

// ReservationItemsRelationships ...
type ReservationItemsRelationships struct {
	TechnicalStops []*TechnicalStops
}

// TripDetailPtcFareBreakdowns ...
type TripDetailPtcFareBreakdowns struct {
	Relationships *TripDetailPtcFareBreakdownsRelationships `json:"relationships"`
}

// TripDetailPtcFareBreakdownsRelationships ...
type TripDetailPtcFareBreakdownsRelationships struct {
	PassengerTypeQuantity   *PassengerTypeQuantity
	TripDetailPassengerFare *TripDetailPassengerFare
}

// PassengerTypeQuantity ...
type PassengerTypeQuantity struct {
	PassengerType int32
	Quantity      int32
}

// TripDetailPassengerFare ...
type TripDetailPassengerFare struct {
	BaseFare   float32
	ServiceTax float32
	Tax        float32
	TotalFare  float32
	Commission float32
	Currency   string
}

// Eticket ...
type Eticket struct {
	ETicketNumber string
	DateOfIssue   string
	IsRefunded    bool
	AirlinePnr    string
	TotalRefund   float32
}

// Services ...
type Services struct {
	PassengerFirstName  string
	PassengerMiddleName string
	PassengerLastName   string
	Description         string
	ServiceId           string
	Behavior            int32
	Amount              float32
	Currency            string
}
