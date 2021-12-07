package models

import "encoding/json"

type TravelClass int

const (
	// TravelClass_ECONOMY ...
	TravelClass_ECONOMY TravelClass = iota
	// TravelClass_PREMIUM_ECONOMY ...
	TravelClass_PREMIUM_ECONOMY
	// TravelClass_BUSINESS ...
	TravelClass_BUSINESS
	// TravelClass_FIRST ...
	TravelClass_FIRST
)

// SearchResponse ...
type SearchResponse struct {
	Errors []*StdError `json:"errors"`
	Data   []*Data     `json:"data"`
	Meta   *Meta       `json:"meta"`
	Links  *Link       `json:"links"`
}

// Errors ...
type StdError struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// Data ...
type Data struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Attributes    *Attributes    `json:"attributes"`
	Relationships *Relationships `json:"relationships"`
}

// MarshalBinary ...
func (data *Data) MarshalBinary() ([]byte, error) {
	return json.Marshal(data)
}

// UnmarshalBinary ...
func (data *Data) UnmarshalBinary(in []byte) error {
	return json.Unmarshal(in, &data)
}

// Attributes ...
type Attributes struct {
	Currency string `json:"currency"`
	Stop     bool   `json:"stop"`
	Date     *Date  `json:"date"`
}

// Date ...
type Date struct {
	ArrivalDate   string `json:"arrival_date"`
	DepartureDate string `json:"departure_date"`
}

// Relationships ...
type Relationships struct {
	OfferItems []*OfferItems `json:"offer_items"`
	Links      *Link         `json:"links"`
}

// OfferItems ...
type OfferItems struct {
	Relationships *OfferItemRelationship `json:"relationships"`
	Links         *Link                  `json:"links"`
}

// Price ...
type Price struct {
	Total    string `json:"total"`
	TotalTax string `json:"total_tax"`
}

// OfferItemRelationship ...
type OfferItemRelationship struct {
	Price          *Price     `json:"price"`
	PricePerAdult  *Price     `json:"price_per_adult"`
	PricePerChild  *Price     `json:"price_per_child"`
	PricePerInfant *Price     `json:"price_per_infant"`
	Service        []*Service `json:"service"`
}

// Service ...
type Service struct {
	Attributes           *ServiceAttributes    `json:"attributes"`
	ServiceRelationships *ServiceRelationships `json:"relationships"`
}

// ServiceAttributes ...
type ServiceAttributes struct {
	Duration       string `json:"duration"`
	DurationPerMin int32  `json:"duration_per_min"`
}

// ServiceRelationships ...
type ServiceRelationships struct {
	Segment []*Segment `json:"segment"`
}

// Segment ...
type Segment struct {
	ID            string                `json:"id"`
	Type          string                `json:"type"`
	Attributes    *SegmentAtt           `json:"attributes"`
	Relationships *SegmentRelationships `json:"relationships"`
}

// SegmentAtt ...
type SegmentAtt struct {
	TravelClass    string      `json:"travel_class"`
	FareClass      string      `json:"fare_class"`
	Availability   int32       `json:"availability"`
	FareBasis      string      `json:"fare_basis"`
	Duration       string      `json:"duration"`
	DurationPerMin int32       `json:"duration_per_min"`
	Arrival        *DateDetail `json:"arrival"`
	Departure      *DateDetail `json:"departure"`
	CarrierCode    string      `json:"carrier_code"`
	Number         string      `json:"number"`
	AircraftCode   string      `json:"aircraft_code"`
	Baggage        string      `json:"baggage"`
	IsCharter      bool        `json:"is_charter"`
}

// DateDetail ...
type DateDetail struct {
	IataCode string `json:"iata_code"`
	At       string `json:"at"`
}

// SegmentRelationships ...
type SegmentRelationships struct {
	Operating      *Operating        `json:"operating"`
	TechnicalStops []*TechnicalStops `json:"technical_stops"`
}

// TechnicalStops ...
type TechnicalStops struct {
	Data *TechnicalStopsData `json:"data"`
}

// TechnicalStopsData ..
type TechnicalStopsData struct {
	ArrivalAirport    string `json:"arrival_airport"`
	ArrivalDateTime   string `json:"arrival_dateTime"`
	DepartureDateTime string `json:"departure_dateTime"`
}

// Operating ...
type Operating struct {
	CarrierCode string `json:"carrier_code"`
	Number      string `json:"number"`
}

// Link ...
type Link struct {
	Self string `json:"self"`
}

// Meta ...
type Meta struct {
	Pages        *Page       `json:"pages"`
	Currency     string      `json:"currency"`
	Defaults     *Default    `json:"defaults"`
	Dictionaries *Dictionary `json:"dictionaries"`
}

// Dictionary ...
type Dictionary struct {
	Carriers   map[string]string `json:"carriers"`
	Currencies map[string]string `json:"currencies"`
	AirCrafts  map[string]string `json:"aircraft"`
	Locations  map[string]string `json:"locations"`
}

// Page ...
type Page struct {
	TotalPages int32 `json:"total_pages"`
	From       int32 `json:"from"`
	LastPage   int32 `json:"last-page"`
	PerPage    int32 `json:"per_page"`
	To         int32 `json:"to"`
	Total      int32 `json:"total"`
}

// Default ...
type Default struct {
	Stop   bool   `json:"stop"`
	Adults string `json:"adults"`
}

// UrlLink ...
type UrlLink struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

// RTResponse ...
type RTResponse struct {
	Status     string         `json:"Status"`
	Contact    string         `json:"Contact"`
	Passengers []PNRPassenger `json:"Passengers"`
	Segments   []Segment      `json:"Segments"`
	Tickets    []Ticket       `json:"Tickets"`
	AdultQty   int32          `json:"AdultQty"`
	ChildQty   int32          `json:"ChildQty"`
	InfantQty  int32          `json:"InfantQty"`
	AdultTP    int32          `json:"AdultTP"`
	ChildTP    int32          `json:"ChildTP"`
	InfantTP   int32          `json:"InfantTP"`
	TotalPrice int32          `json:"TotalPrice"`
}

// Tickets ...
type Ticket struct {
	PassengerET string `json:"PassengerET"`
}

// PNRPassenger ...
type PNRPassenger struct {
	PassengerFirstName string `json:"PassengerFirstName"`
	PassengerLastName  string `json:"PassengerLastName"`
	PassengerAgeType   string `json:"PassengerAgeType"`
}

// RTSegment ...
type RTSegment struct {
	DepartureDT string `json:"DepartureDT"` // flight date
	FlightClass string `json:"FlightClass"`
	Origin      string `json:"Origin"`
	FlightNo    string `json:"FlightNo"`
	Destination string `json:"Destination"`
}

// CancelPNRResponse ...
type CancelPNRResponse struct {
	Error string `json:"Error"`
}

/*{
"AirCancelPNR": [
{
"Error": "No Error "
}
]
}*/
// AirCancelPNR ...
type AirCancelPNR struct {
	Error string `json:"Error"`
}

// CancelSeatResponse ...
type CancelSeatResponse struct {
	AirCancelSeat AirCancelSeat `json:"AirCancelSeat"`
}

// AirCancelSeat ...
type AirCancelSeat struct {
	Done string `json:"Done"`
}

/*{
"AirNRSTICKETS": [
{
"Tickets":"GH/NIMA=0002442634350
"}],"Message": "محل صدور  بليت الکترونيک: جزيره کيش"}*/

// ETRefundResponse ...
type ETRefundResponse struct {
	AirNRSRefund AirNRSRefund `json:"AirNRSRefund"`
}

// AirNRSRefund ...
type AirNRSRefund struct {
	Done string `json:"Done"`
}

// ETRResponse ...
type ETRResponse struct {
	PassengerFullName string    `json:"PassengerFullName"`
	Fare              string    `json:"Fare"`
	Commission        string    `json:"Commission"`
	PAX               string    `json:"PAX"`
	Coupons           []*Coupon `json:"COUPONS"`
	TotalPrice        int32     `json:"TotalPrice"`
	TicketNo          string    `json:"TicketNo"`
	Taxes             []*TAX    `json:"TAXES"`
}

// Tax ...
type TAX struct {
	TaxCode   string `json:"TaxCode"`
	TaxAmount string `json:"TaxAmount"`
}

// Coupon ...
type Coupon struct {
	Status      string `json:"Status"`
	Departure   string `json:"Departure"`
	FlightClass string `json:"FlightClass"`
	FlightNo    string `json:"FlightNo"`
	Origin      string `json:"Origin"`
	Destination string `json:"Destination"`
}
