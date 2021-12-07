package repository

import (
	"github.com/hooklift/gowsdl/soap"
)

type impl struct {
	client *soap.Client
}

// NewEPowerServiceSoap ...
func NewEPowerServiceSoap(client *soap.Client) IAmadeus {
	return &impl{
		client: client,
	}
}

// Ping , pings the server with at least an echo data
func (service *impl) Ping(request *Ping) (*PingResponse, error) {
	response := new(PingResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/Ping", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// SignOut , revokes the session and signs out
func (service *impl) SignOut(request *SignOut) (*SignOutResponse, error) {
	response := new(SignOutResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/SignOut", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// Cancel , cancel any amadeus services which had booked
func (service *impl) Cancel(request *Cancel) (*CancelResponse, error) {
	response := new(CancelResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/Cancel", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// SearchFlight , normal search for low fare flights
func (service *impl) SearchFlight(request *SearchFlightRequest) (*SearchFlightResponse, error) {
	response := new(SearchFlightResponse)
	
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/SearchFlight", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetBaggageInfo , get baggage info for various airlines
func (service *impl) GetBaggageInfo(request *GetBaggageInfo) (*GetBaggageInfoResponse, error) {
	response := new(GetBaggageInfoResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/GetBaggageInfo", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// GetNextFlightResults , the next page of search results if first page hasMoreResult filed be true
func (service *impl) GetNextFlightResults(request *GetNextFlightResults) (*GetNextFlightResultsResponse, error) {
	response := new(GetNextFlightResultsResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/GetNextFlightResults", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// GetFlightRules , gets the rules of flight per each airline
func (service *impl) GetFlightRules(request *GetFlightRules) (*GetFlightRulesResponse, error) {
	response := new(GetFlightRulesResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/GetFlightRules", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// BookFlight , reservation progress for selected flight
func (service *impl) BookFlight(request *BookFlight) (*BookFlightResponse, error) {
	response := new(BookFlightResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/BookFlight", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// BookFlightWithRecommendation , reserve progress for recommended flights
func (service *impl) BookFlightWithRecommendation(request *BookFlightWithRecommendation) (*BookFlightWithRecommendationResponse, error) {
	response := new(BookFlightWithRecommendationResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/BookFlightWithRecommendation", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// CreateTicket , finalizing booked flight and get the ticket(s)
func (service *impl) CreateTicket(request *CreateTicket) (*CreateTicketResponse, error) {
	response := new(CreateTicketResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/CreateTicket", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// GetPNR , get reserved flight's information
func (service *impl) GetPNR(request *GetPNR) (*GetPNRResponse, error) {
	response := new(GetPNRResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/GetPNR", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// GetPNRPrice , get reserved flight's price
func (service *impl) GetPNRPrice(request *GetPNRPrice) (*GetPNRPriceResponse, error) {
	response := new(GetPNRPriceResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/GetPNRPrice", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// GetCancelPolicyText , get cancel policies for selected airline
func (service *impl) GetCancelPolicyText(request *GetCancelPolicyText) (*GetCancelPolicyTextResponse, error) {
	response := new(GetCancelPolicyTextResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/GetCancelPolicyText", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// CheckETicket , gets the ticket issued for a flight via bookFlight API
func (service *impl) CheckETicket(request *CheckETicket) (*CheckETicketResponse, error) {
	response := new(CheckETicketResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/CheckETicket", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}

// GetPrice , gets the price for some provides them responses dont contain price ( Turkish market )
func (service *impl) GetPrice(request *GetPrice) (*GetPriceResponse, error) {
	response := new(GetPriceResponse)
	err := service.client.Call("http://epowerv5.amadeus.com.tr/WS/GetPrice", request, response)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}
