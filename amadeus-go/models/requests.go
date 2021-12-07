package models

// SearchRequest ...
type SearchRequest struct {
    Origin        string  `json:"origin"`         // e.g THR
    Destination   string  `json:"destination"`    // e.g MHD
    DepartureTime string  `json:"departure_time"` // e.g 2000-12-26T01:15:06+04:20
    ReturnTime    string  `json:"return_time"`    // e.g 2000-12-26T01:15:06+04:20
    AdultCount    int32   `json:"adult_count"`    // e.g 2
    ChildCount    int32   `json:"child_count"`    // e.g 0
    InfantCount   int32   `json:"infant_count"`   // e.g 1
    Airline       string  `json:"airline"`
    AirClass      string  `json:"air_class"`
    Sort          *Sort   `json:"sort"`
    Filter        *Filter `json:"filter"`
}

// Filter ...
type Filter struct {
    Duration           IntLimit `json:"duration"`
    MaxStopDuration    IntLimit `json:"max_stop_duration"`
    StopCount          int32    `json:"stop_count"`
    AvailabilityTicket bool     `json:"availability_ticket"`
    Airport            string   `json:"airport"`
    WentHours          StrLimit `json:"went_hours"`
    ReturnHours        StrLimit `json:"return_hours"`
    Type               string   `json:"type"`
}

// Sort ...
type Sort struct {
    Price string `json:"price"`
}

// IntLimit ...
type IntLimit struct {
    Max int32 `json:"max"`
    Min int32 `json:"min"`
}

// StrLimit ...
type StrLimit struct {
    Max string `json:"max"`
    Min string `json:"min"`
}

// FareRequest ...
type FareRequest struct {
    Route string `json:"Route"` // e.g THR-MHD
    RBD   string `json:"RBD"`   // e.g Y or Z
}

// RTRequest ...
type RTRequest struct {
    PNR string
}

// CancelPNRRequest ...
type CancelPNRRequest struct {
    PNR string
}

// CancelSeatRequest ...
type CancelSeatRequest struct {
    PNR               string
    PassengerName     string
    PassengerLastName string
    DepartureDate     string
    FlightNo          string
}


// ETRefundRequest ...
type ETRefundRequest struct {
    TicketNo string
    Fare     int32
    KU       int32
    LP       int32
    Penalty  int32
}

// ETRRequest ...
type ETRRequest struct {
    TicketNo string
}
