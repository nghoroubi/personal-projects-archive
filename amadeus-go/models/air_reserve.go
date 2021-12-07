package models

import "encoding/json"

// ReserveReq ...
type ReserveReq struct {
	Request *ReserveRequest `json:"request"`
}

// ReserveResp ...
type ReserveResp struct {
	Response *ReserveResponse `json:"response"`
}

// ReserveRequest ...
type ReserveRequest struct {
	SessionId       string        `json:"session_id"`
	ServiceId       int32         `json:"service_id"`
	FareSourceCode  string        `json:"fare_source_code"`
	ClientUniqueId  string        `json:"client_unique_id"`
	MarkupForAdult  uint64        `json:"markup_for_adult"`
	MarkupForChild  uint64        `json:"markup_for_child"`
	MarkupForInfant uint64        `json:"markup_for_infant"`
	TravelerInfo    *TravelerInfo `json:"traveler_info"`
	TravelInfo      *TravelInfo   `json:"travel_info"`
}
// MarshalBinary ...
func (data *ReserveRequest) MarshalBinary() ([]byte, error) {
	return json.Marshal(data)
}

// UnmarshalBinary ...
func (data *ReserveRequest) UnmarshalBinary(in []byte) error {
	return json.Unmarshal(in, &data)
}

// TravelInfo ...
type TravelInfo struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	FlightClass   string `json:"flight_class"`
	FlightNo      string `json:"flight_no"`
	DepartureDate string `json:"departure_date"`
}

// TravelerInfo ...
type TravelerInfo struct {
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone_number"`
	AirTravelers []*AirTraveler `json:"air_travelers"`
}

// AirTraveler ...
type AirTraveler struct {
	DateOfBirth         string         `json:"date_of_birth" validate:"required"`
	Gender              int32          `json:"gender" validate:"required"`         // 0 =>male  , 1 => fmale
	PassengerType       int32          `json:"passenger_type" validate:"required"` // 1 => Adult , 2 => Child  3 => Infant
	PassengerName       *PassengerName `json:"passenger_name"`
	Passport            *Passport      `json:"passport"`
	Visa                *Visa          `json:"visa"`
	NationalId          string         `json:"national_id"`
	Nationality         string         `json:"nationality"`
	ExtraServiceId      []string       `json:"extra_service_id"`
	FrequentFlyerNumber string         `json:"frequent_flyer_number"`
	SeatPreference      int32          `json:"seat_preference"` // 0 =>any , 1=>Aisle , 2=>Window
	MealPreference      int32          `json:"meal_preference"`
	Wheelchair          bool           `json:"wheelchair"`
}

// PassengerName ...
type PassengerName struct {
	PassengerFirstName  string `json:"passenger_first_name" validate:"required"`
	PersianFirstName    string `json:"persian_first_name"`
	PassengerMiddleName string `json:"passenger_middle_name" validate:"required"`
	PassengerLastName   string `json:"passenger_last_name" validate:"required"`
	PersianLastName     string `json:"persian_last_name"`
	PassengerTitle      int32  `json:"passenger_title" validate:"required"` //  Adult - Male=>0 , -Adult - FeMale =>1 , Adule - Female=>2 ,Child,Infant - Female =>4 , Child,Infant - Male =>5
}

// Passport ...
type Passport struct {
	Country        string `json:"country"`
	ExpiryDate     string `json:"expiry_date"`
	IssueDate      string `json:"issue_date"`
	PassportNumber string `json:"passport_number"`
}

// Visa ...
type Visa struct {
	VisaType    string `json:"visa_type"`
	VisaCountry string `json:"visa_country"`
}

// ReserveResponse ...
type ReserveResponse struct {
	Errors []*StdError  `json:"errors"`
	Data   *ReserveData `json:"data"`
	Meta   *Meta        `json:"meta"`
	Links  *Link        `json:"links"`
}

// ReserveData ...
type ReserveData struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes *ReserveAttributes `json:"attributes"`
}

// ReserveAttributes ...
type ReserveAttributes struct {
	Success      bool   `json:"success"`
	TktTimeLimit string `json:"tkt_time_limit"`
	Category     int32  `json:"category"`
	Status       int32  `json:"status"`
	UniqueId     string `json:"unique_id"`
}

// AirBookRequest ...
type AirBookRequest struct {
	PNR   string `json:"PNR"`
	Email string `json:"Email"`
}

// AirBookResponse ...
type AirBookResponse struct {
	Tickets []FinalTicket
}

// FinalTicket ...
type FinalTicket struct {
	PassengerFirstName string
	PassengerLastName  string
	TicketNo           string
}
