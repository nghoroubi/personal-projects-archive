package repository

// IFlight ...
type IFlight interface {
	GetPNR(request *GetPNR) (*GetPNRResponse, error)
	Cancel(request *Cancel) (*CancelResponse, error)
	GetPrice(request *GetPrice) (*GetPriceResponse, error)
	BookFlight(request *BookFlight) (*BookFlightResponse, error)
	GetPNRPrice(request *GetPNRPrice) (*GetPNRPriceResponse, error)
	CheckETicket(request *CheckETicket) (*CheckETicketResponse, error)
	CreateTicket(request *CreateTicket) (*CreateTicketResponse, error)
	GetBaggageInfo(request *GetBaggageInfo) (*GetBaggageInfoResponse, error)
	GetFlightRules(request *GetFlightRules) (*GetFlightRulesResponse, error)
	SearchFlight(request *SearchFlightRequest) (*SearchFlightResponse, error)
	GetCancelPolicyText(request *GetCancelPolicyText) (*GetCancelPolicyTextResponse, error)
	GetNextFlightResults(request *GetNextFlightResults) (*GetNextFlightResultsResponse, error)
	BookFlightWithRecommendation(request *BookFlightWithRecommendation) (*BookFlightWithRecommendationResponse, error)
}
// IConnectivity ...
type IConnectivity interface {
	Ping(request *Ping) (*PingResponse, error)
	SignOut(request *SignOut) (*SignOutResponse, error)
}

// IAmadeus , main interface of amadeus service
type IAmadeus interface {
	IFlight
	IConnectivity
}
