package repository

import (
	"encoding/xml"
	"time"
)

// Ping ...
type Ping struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS Ping"`
	
	OTA_PingRQ *OTAPingRQ `xml:"OTA_PingRQ,omitempty"`
}

// OTAPingRQ ...
type OTAPingRQ struct {
	EchoData  string `xml:"EchoData,omitempty"`
	EchoToken string `xml:"EchoToken,attr,omitempty"`
	// TimeStamp string `xml:"TimeStamp,attr,omitempty"`
	Target                  string  `xml:"Target,attr,omitempty"`
	Version                 float64 `xml:"Version,attr,omitempty"`
	TransactionIdentifier   string  `xml:"TransactionIdentifier,attr,omitempty"`
	SequenceNmbr            int   `xml:"SequenceNmbr,attr,omitempty"`
	TransactionStatusCode   string  `xml:"TransactionStatusCode,attr,omitempty"`
	RetransmissionIndicator bool    `xml:"RetransmissionIndicator,attr,omitempty"`
}

// PingResponse ...
type PingResponse struct {
	XMLName   xml.Name   `xml:"http://epowerv5.amadeus.com.tr/WS PingResponse"`
	OTAPingRS *OTAPingRQ `xml:"OTA_PingRS,omitempty"`
}

// OTAPingRS ...
type OTAPingRS struct {
	SequenceNmbr            uint          `xml:"SequenceNmbr,attr,omitempty"`
	RetransmissionIndicator bool          `xml:"RetransmissionIndicator,attr,omitempty"`
	EchoToken               string        `xml:"EchoToken,attr,omitempty"`
	EchoData                string        `xml:"EchoData,omitempty"`
	TransactionIdentifier   string        `xml:"TransactionIdentifier,attr,omitempty"`
	TransactionStatusCode   string        `xml:"TransactionStatusCode,attr,omitempty"`
	Target                  string        `xml:"Target,attr,omitempty"`
	Version                 float64       `xml:"Version,attr,omitempty"`
	TimeStamp               string     `xml:"TimeStamp,attr,omitempty"`
	Errors                  *ErrorsType   `xml:"Errors,omitempty"`
	Success                 *SuccessType  `xml:"Success,omitempty"`
	Warnings                *WarningsType `xml:"Warnings,omitempty"`
}

// SignOut ...
type SignOut struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS SignOut"`
}

// SignOutResponse ...
type SignOutResponse struct {
	XMLName       xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS SignOutResponse"`
	SignOutResult bool     `xml:"SignOutResult,omitempty"`
}

// SearchFlightRequest ...
type SearchFlightRequest struct {
	XMLName               xml.Name               `xml:"http://epowerv5.amadeus.com.tr/WS SearchFlight"`
	OTAAirLowFareSearchRQ *OTAAirLowFareSearchRQ `xml:"OTA_AirLowFareSearchRQ,omitempty"`
}

// OTAAirLowFareSearchRQ ...
type OTAAirLowFareSearchRQ struct {
	Version                      int                             `xml:"Version,attr,omitempty"`
	SequenceNmbr                 string                            `xml:"SequenceNmbr,attr,omitempty"`
	MaxResponses                 uint                            `xml:"MaxResponses,attr,omitempty"`
	RetransmissionIndicator      bool                            `xml:"RetransmissionIndicator,attr,omitempty"`
	DirectFlightsOnly            bool                            `xml:"DirectFlightsOnly,attr,omitempty"`
	AvailableFlightsOnly         bool                            `xml:"AvailableFlightsOnly,attr,omitempty"`
	TimeStamp                    string                          `xml:"TimeStamp,attr,omitempty"`
	Target                       string                          `xml:"Target,attr,omitempty"`
	TransactionIdentifier        string                          `xml:"TransactionIdentifier,attr,omitempty"`
	TransactionStatusCode        string                          `xml:"TransactionStatusCode,attr,omitempty"`
	MaxPrice                     float64                         `xml:"MaxPrice,attr,omitempty"`
	AdvanceSearchInfo            *AdvanceSearchInfo              `xml:"AdvanceSearchInfo,omitempty"`
	TravelerInfoSummary          *TravelerInfoSummary            `xml:"TravelerInfoSummary,omitempty"`
	RefundableType               RefundableTypesEnum             `xml:"RefundableType,attr,omitempty"`
	ProviderType                 SearchTypeByProviderEnum        `xml:"ProviderType,attr,omitempty"`
	OriginDestinationInformation []*OriginDestinationInformation `xml:"OriginDestinationInformation,omitempty"`
	TravelPreferences            []*TravelPreferences            `xml:"TravelPreferences,omitempty"`
	SpecificFlightInfo           *SpecificFlightInfoType         `xml:"SpecificFlightInfo,omitempty"`
}

// OriginDestinationInformation ...
type OriginDestinationInformation struct {
	*OriginDestinationInformationType
	RefNumber             int32                             `xml:"RefNumber,attr,omitempty"`
	RPH                   string                            `xml:"RPH,attr,omitempty"`
	TPAExtensions         *TPAExtensionsType                `xml:"TPA_Extensions,omitempty"`
	AlternateLocationInfo *OriginDestinationInformationType `xml:"AlternateLocationInfo,omitempty"`
}

// TravelerInfoSummary ...
type TravelerInfoSummary struct {
	*TravelerInfoSummaryType
	ForcePTC             bool   `xml:"ForcePTC,attr,omitempty"`
	FamilyCard           bool   `xml:"FamilyCard,attr,omitempty"`
	FamilyType           bool   `xml:"FamilyType,attr,omitempty"`
	FamilyDiscount       bool   `xml:"FamilyDiscount,attr,omitempty"`
	SpecificPTCIndicator bool   `xml:"SpecificPTC_Indicator,attr,omitempty"`
	TicketingCountryCode string `xml:"TicketingCountryCode,attr,omitempty"`
}

// AdvanceSearchInfo ...
type AdvanceSearchInfo struct {
	IgnoreAvailability         bool                    `xml:"IgnoreAvailability,omitempty"`
	ReturnMoreOvernightFlights bool                    `xml:"ReturnMoreOvernightFlights,omitempty"`
	NumberOfRecommendation     uint                    `xml:"NumberOfRecommendation,omitempty"`
	MaxEFTPercentage           int32                   `xml:"MaxEFTPercentage,omitempty"`
	Currency                   string                  `xml:"Currency,omitempty"`
	MaxLayoverPerConnection    *Connection             `xml:"MaxLayoverPerConnection,omitempty"`
	AirlineDiversity           *AirlineDiversityType   `xml:"AirlineDiversity,omitempty"`
	ExpandedParameters         *ExpandedParametersType `xml:"ExpandedParameters,omitempty"`
	
	// DuplicatedRecommedationsBehavior *DuplicatedRecommedationsBehaviorEnum `xml:"DuplicatedRecommedationsBehavior,omitempty"`
}

// Connection ...
type Connection struct {
	Hour   int32 `xml:"Hour,attr,omitempty"`
	Minute int32 `xml:"Minute,attr,omitempty"`
}

// TravelerInfoSummaryType ...
type TravelerInfoSummaryType struct {
	SeatsRequested          []uint                       `xml:"SeatsRequested,omitempty"`
	AirTravelerAvail        []*TravelerInformationType   `xml:"AirTravelerAvail,omitempty"`
	PriceRequestInformation *PriceRequestInformationType `xml:"PriceRequestInformation,omitempty"`
}

// TravelerInformationType ...
type TravelerInformationType struct {
	PassengerTypeQuantity []*PassengerTypeQuantityType `xml:"PassengerTypeQuantity,omitempty"`
	// AirTraveler *AirTravelerType `xml:"AirTraveler,omitempty"`
}

// PassengerTypeQuantityType ...
type PassengerTypeQuantityType struct {
	Code             string `xml:"Code,attr,omitempty"`
	Quantity         int32  `xml:"Quantity,attr,omitempty"`
	Age              int32  `xml:"Age,attr,omitempty"`
	CodeContext      string `xml:"CodeContext,attr,omitempty"`
	URI              string `xml:"URI,attr,omitempty"`
	AlternativeCode1 string `xml:"AlternativeCode1,attr,omitempty"`
	AlternativeCode2 string `xml:"AlternativeCode2,attr,omitempty"`
}

// SearchFlightResponse ...
type SearchFlightResponse struct {
	XMLName               xml.Name               `xml:"http://epowerv5.amadeus.com.tr/WS SearchFlightResponse"`
	OTAAirLowFareSearchRS *OTAAirLowFareSearchRS `xml:"OTA_AirLowFareSearchRS,omitempty"`
}

// OTAAirLowFareSearchRS ...
type OTAAirLowFareSearchRS struct {
	HasMoreResult     bool                   `xml:"HasMoreResult,omitempty"`
	Version           float64                `xml:"Version,attr,omitempty"`
	TimeStamp         string              `xml:"TimeStamp,attr,omitempty"`
	Errors            *ErrorsType            `xml:"Errors,omitempty"`
	Success           *SuccessType           `xml:"Success,omitempty"`
	Warnings          *WarningsType          `xml:"Warnings,omitempty"`
	PricedItineraries *PricedItinerariesType `xml:"PricedItineraries,omitempty"`
}

// TravelPreferences ...
type TravelPreferences struct {
	*AirSearchPrefsType
	FlexWeekendIndicator  bool     `xml:"FlexWeekendIndicator,attr,omitempty"`
	FlexLevelIndicator    bool     `xml:"FlexLevelIndicator,attr,omitempty"`
	NoFareBreakIndicator  bool     `xml:"NoFareBreakIndicator,attr,omitempty"`
	FlexDatePref          string   `xml:"FlexDatePref,attr,omitempty"`
	OriginDestinationRPHs []string `xml:"OriginDestinationRPHs,attr,omitempty"`
}

// BookFlight ...
type BookFlight struct {
	XMLName      xml.Name      `xml:"http://epowerv5.amadeus.com.tr/WS BookFlight"`
	OTAAirBookRQ *OTAAirBookRQ `xml:"OTA_AirBookRQ,omitempty"`
}

// BookFlightResponse ...
type BookFlightResponse struct {
	XMLName      xml.Name      `xml:"http://epowerv5.amadeus.com.tr/WS BookFlightResponse"`
	OTAAirBookRS *OTAAirBookRS `xml:"OTA_AirBookRS,omitempty"`
}

// CreateTicket ...
type CreateTicket struct {
	XMLName      xml.Name      `xml:"http://epowerv5.amadeus.com.tr/WS CreateTicket"`
	OTAAirBookRQ *OTAAirBookRQ `xml:"OTA_AirBookRQ,omitempty"`
}

// CreateTicketResponse ...
type CreateTicketResponse struct {
	XMLName      xml.Name      `xml:"http://epowerv5.amadeus.com.tr/WS CreateTicketResponse"`
	OTAAirBookRS *OTAAirBookRS `xml:"OTA_AirBookRS,omitempty"`
}

// GetPNRPrice ...
type GetPNRPrice struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetPNRPrice"`
	PNRNo   string   `xml:"PNRNo,attr,omitempty"`
	Surname string   `xml:"Surname,attr,omitempty"`
}

// GetPNRPriceResponse ...
type GetPNRPriceResponse struct {
	XMLName      xml.Name      `xml:"http://epowerv5.amadeus.com.tr/WS GetPNRPriceResponse"`
	OTAAirBookRS *OTAAirBookRS `xml:"OTA_AirBookRS,omitempty"`
}

// CurrencyConversion ...
type CurrencyConversion struct {
	XMLName                 xml.Name                 `xml:"http://epowerv5.amadeus.com.tr/WS CurrencyConversion"`
	OTACurrencyConversionRQ *OTACurrencyConversionRQ `xml:"OTA_CurrencyConversionRQ,omitempty"`
}

// CurrencyConversionResponse ...
type CurrencyConversionResponse struct {
	XMLName                 xml.Name                 `xml:"http://epowerv5.amadeus.com.tr/WS CurrencyConversionResponse"`
	OTACurrencyConversionRS *OTACurrencyConversionRS `xml:"OTA_CurrencyConversionRS,omitempty"`
}

// Cancel ...
type Cancel struct {
	XMLName      xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS Cancel"`
	OTA_CancelRQ struct {
		PNRNo string `xml:"PNRNo,omitempty"`
		
		BookingReferenceID *UniqueIDType `xml:"BookingReferenceID,omitempty"`
		
		POS *ArrayOfSourceType `xml:"POS,omitempty"`
		
		UniqueID []struct {
			*UniqueIDType
			
			Reason string `xml:"Reason,attr,omitempty"`
		} `xml:"UniqueID,omitempty"`
		
		Verification []*VerificationType `xml:"Verification,omitempty"`
		
		Segment []struct {
			ItinSegNbr int `xml:"ItinSegNbr,attr,omitempty"`
			
			FirstItinSegNbr int `xml:"FirstItinSegNbr,attr,omitempty"`
			
			LastItinSegNbr int `xml:"LastItinSegNbr,attr,omitempty"`
		} `xml:"Segment,omitempty"`
		
		CancellationOverrides struct {
			CancellationOverride *CancelRuleType `xml:"CancellationOverride,omitempty"`
		} `xml:"CancellationOverrides,omitempty"`
		
		Reasons *ArrayOfFreeTextType `xml:"Reasons,omitempty"`
		
		TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
		
		EchoToken string `xml:"EchoToken,attr,omitempty"`
		
		TimeStamp string `xml:"TimeStamp,attr,omitempty"`
		
		Target string `xml:"Target,attr,omitempty"`
		
		Version float64 `xml:"Version,attr,omitempty"`
		
		TransactionIdentifier string `xml:"TransactionIdentifier,attr,omitempty"`
		
		SequenceNmbr int `xml:"SequenceNmbr,attr,omitempty"`
		
		TransactionStatusCode string `xml:"TransactionStatusCode,attr,omitempty"`
		
		RetransmissionIndicator bool `xml:"RetransmissionIndicator,attr,omitempty"`
		
		ReqRespVersion string `xml:"ReqRespVersion,attr,omitempty"`
		
		CancelType *TransactionActionType `xml:"CancelType,attr,omitempty"`
		
		TravelSector string `xml:"TravelSector,attr,omitempty"`
	} `xml:"OTA_CancelRQ,omitempty"`
}

// OTACancelRQ ...
type OTACancelRQ struct {
	SequenceNmbr            uint                   `xml:"SequenceNmbr,attr,omitempty"`
	RetransmissionIndicator bool                   `xml:"RetransmissionIndicator,attr,omitempty"`
	PNRNo                   string                 `xml:"PNRNo,omitempty"`
	EchoToken               string                 `xml:"EchoToken,attr,omitempty"`
	Target                  string                 `xml:"Target,attr,omitempty"`
	TransactionIdentifier   string                 `xml:"TransactionIdentifier,attr,omitempty"`
	TransactionStatusCode   string                 `xml:"TransactionStatusCode,attr,omitempty"`
	ReqRespVersion          string                 `xml:"ReqRespVersion,attr,omitempty"`
	TravelSector            string                 `xml:"TravelSector,attr,omitempty"`
	Version                 float64                `xml:"Version,attr,omitempty"`
	TimeStamp              string              `xml:"TimeStamp,attr,omitempty"`
	BookingReferenceID      *UniqueIDType          `xml:"BookingReferenceID,omitempty"`
	UniqueID                []*CancelUniqueID      `xml:"UniqueID,omitempty"`
	Segment                 []*CancelSegment       `xml:"Segment,omitempty"`
	POS                     *ArrayOfSourceType     `xml:"POS,omitempty"`
	TPAExtensions           *TPAExtensionsType     `xml:"TPA_Extensions,omitempty"`
	Verification            []*VerificationType    `xml:"Verification,omitempty"`
	Reasons                 *ArrayOfFreeTextType   `xml:"Reasons,omitempty"`
	CancelType              *TransactionActionType `xml:"CancelType,attr,omitempty"`
	CancellationOverrides   *CancellationOverrides `xml:"CancellationOverrides,omitempty"`
}

// CancellationOverrides ...
type CancellationOverrides struct {
	CancellationOverride *CancelRuleType `xml:"CancellationOverride,omitempty"`
}

// CancelUniqueID ...
type CancelUniqueID struct {
	*UniqueIDType
	Reason string `xml:"Reason,attr,omitempty"`
}

// CancelSegment ...
type CancelSegment struct {
	ItinSegNbr      int `xml:"ItinSegNbr,attr,omitempty"`
	FirstItinSegNbr int `xml:"FirstItinSegNbr,attr,omitempty"`
	LastItinSegNbr  int `xml:"LastItinSegNbr,attr,omitempty"`
}

// CancelResponse ...
type CancelResponse struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS CancelResponse"`
	
	OTA_CancelRS struct {
		// OTA_HotelResRS *OTA_HotelResRS `xml:"OTA_HotelResRS,omitempty"`
		
		OTA_AirBookRS struct {
			NewPrice *AirItineraryPricingInfoType `xml:"NewPrice,omitempty"`
			
			OldPrice *AirItineraryPricingInfoType `xml:"OldPrice,omitempty"`
			
			AirReservation struct {
				*AirReservationType
			} `xml:"AirReservation,omitempty"`
			
			BookingToBeCancelled *BookingToBeCancelled `xml:"BookingToBeCancelled,omitempty"`
			
			Errors *ErrorsType `xml:"Errors,omitempty"`
			
			Success *SuccessType `xml:"Success,omitempty"`
			
			Warnings *WarningsType `xml:"Warnings,omitempty"`
			
			ReferenceNumber string `xml:"ReferenceNumber,attr,omitempty"`
			
			PaymentRequestURL string `xml:"PaymentRequestURL,attr,omitempty"`
			
			Cancel bool `xml:"Cancel,attr,omitempty"`
			
			EchoToken string `xml:"EchoToken,attr,omitempty"`
			
			TimeStamp string `xml:"TimeStamp,attr,omitempty"`
			
			Target string `xml:"Target,attr,omitempty"`
			
			Version float64 `xml:"Version,attr,omitempty"`
			
			TransactionIdentifier string `xml:"TransactionIdentifier,attr,omitempty"`
			
			SequenceNmbr int `xml:"SequenceNmbr,attr,omitempty"`
			
			TransactionStatusCode string `xml:"TransactionStatusCode,attr,omitempty"`
			
			RetransmissionIndicator bool `xml:"RetransmissionIndicator,attr,omitempty"`
			
			PriceInd bool `xml:"PriceInd,attr,omitempty"`
		} `xml:"OTA_AirBookRS,omitempty"`
		
		// OTA_VehResRS *OTA_VehResRS `xml:"OTA_VehResRS,omitempty"`
		
		Errors *ErrorsType `xml:"Errors,omitempty"`
		
		Success *SuccessType `xml:"Success,omitempty"`
		
		Warnings *WarningsType `xml:"Warnings,omitempty"`
		
		EchoToken string `xml:"EchoToken,attr,omitempty"`
		
		TimeStamp string `xml:"TimeStamp,attr,omitempty"`
		
		Target string `xml:"Target,attr,omitempty"`
		
		Version float64 `xml:"Version,attr,omitempty"`
		
		TransactionIdentifier string `xml:"TransactionIdentifier,attr,omitempty"`
		
		SequenceNmbr int `xml:"SequenceNmbr,attr,omitempty"`
		
		TransactionStatusCode string `xml:"TransactionStatusCode,attr,omitempty"`
		
		RetransmissionIndicator bool `xml:"RetransmissionIndicator,attr,omitempty"`
		
		Status *TransactionStatusType `xml:"Status,attr,omitempty"`
	} `xml:"OTA_CancelRS,omitempty"`
}

// Queue ...
type Queue struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS Queue"`
	QueueRQ *QueueRQ `xml:"QueueRQ,omitempty"`
}

// QueueResponse ...
type QueueResponse struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS QueueResponse"`
	QueueRS *QueueRS `xml:"QueueRS,omitempty"`
}

// GetServiceFee ...
type GetServiceFee struct {
	XMLName          xml.Name          `xml:"http://epowerv5.amadeus.com.tr/WS GetServiceFee"`
	ServiceFeeInfoRQ *ServiceFeeInfoRQ `xml:"ServiceFeeInfoRQ,omitempty"`
}

// GetServiceFeeResponse ...
type GetServiceFeeResponse struct {
	XMLName          xml.Name          `xml:"http://epowerv5.amadeus.com.tr/WS GetServiceFeeResponse"`
	ServiceFeeInfoRS *ServiceFeeInfoRS `xml:"ServiceFeeInfoRS,omitempty"`
}

// GetBaggageInfo ...
type GetBaggageInfo struct {
	XMLName       xml.Name       `xml:"http://epowerv5.amadeus.com.tr/WS GetBaggageInfo"`
	BaggageInfoRQ *BaggageInfoRQ `xml:"BaggageInfoRQ,omitempty"`
}

// GetBaggageInfoResponse ...
type GetBaggageInfoResponse struct {
	XMLName       xml.Name       `xml:"http://epowerv5.amadeus.com.tr/WS GetBaggageInfoResponse"`
	BaggageInfoRS *BaggageInfoRS `xml:"BaggageInfoRS,omitempty"`
}

// GetServiceFeeV2 ...
type GetServiceFeeV2 struct {
	XMLName          xml.Name          `xml:"http://epowerv5.amadeus.com.tr/WS GetServiceFeeV2"`
	ServiceFeeInfoRQ *ServiceFeeInfoRQ `xml:"ServiceFeeInfoRQ,omitempty"`
}

// GetServiceFeeV2Response ...
type GetServiceFeeV2Response struct {
	XMLName          xml.Name          `xml:"http://epowerv5.amadeus.com.tr/WS GetServiceFeeV2Response"`
	ServiceFeeInfoRS *ServiceFeeInfoRS `xml:"ServiceFeeInfoRS,omitempty"`
}

// GetPrice ...
type GetPrice struct {
	XMLName     xml.Name     `xml:"http://epowerv5.amadeus.com.tr/WS GetPrice"`
	PriceInfoRQ *PriceInfoRQ `xml:"PriceInfoRQ,omitempty"`
}

// GetPriceResponse ...
type GetPriceResponse struct {
	XMLName               xml.Name               `xml:"http://epowerv5.amadeus.com.tr/WS GetPriceResponse"`
	OTAAirLowFareSearchRS *OTAAirLowFareSearchRS `xml:"OTA_AirLowFareSearchRS,omitempty"`
}

// GetRequiredParametersForLCC ...
type GetRequiredParametersForLCC struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetRequiredParametersForLCC"`
	
	RequiredParametersRQ *RequiredParametersRQ `xml:"RequiredParametersRQ,omitempty"`
}

// GetRequiredParametersForLCCResponse ...
type GetRequiredParametersForLCCResponse struct {
	XMLName              xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetRequiredParametersForLCCResponse"`
	RequiredParametersRS struct {
		Errors                 *ErrorsType `xml:"Errors,omitempty"`
		RequiredParametersData struct {
			RequiredParametersOnProcessDetails []*LCCRequiredParameter `xml:"RequiredParametersOnProcessDetails,omitempty"`
			
			LCCSupportedCardList *ArrayOfLCCSupportedCard `xml:"LCCSupportedCardList,omitempty"`
			
			LCCBaggageFeeList *ArrayOfLCCBaggageFee `xml:"LCCBaggageFeeList,omitempty"`
			
			OutwardLuggageOptions *ArrayOfLCCBaggageFee `xml:"OutwardLuggageOptions,omitempty"`
			
			ReturnLuggageOptions *ArrayOfLCCBaggageFee `xml:"ReturnLuggageOptions,omitempty"`
			
			LCCHandLuggageFeeList *ArrayOfLCCBaggageFee `xml:"LCCHandLuggageFeeList,omitempty"`
			
			LCCRuleLink string `xml:"LCCRuleLink,omitempty"`
			
			BaggageAllowanceNote string `xml:"BaggageAllowanceNote,omitempty"`
			
			SupplierName string `xml:"SupplierName,omitempty"`
			
			SupplierCode string `xml:"SupplierCode,omitempty"`
			
			ItineraryIndex int32 `xml:"ItineraryIndex,attr,omitempty"`
		} `xml:"RequiredParametersData,omitempty"`
		Success *SuccessType `xml:"Success,omitempty"`
	} `xml:"RequiredParametersRS,omitempty"`
}

// GetNextFlightResults ...
type GetNextFlightResults struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetNextFlightResults"`
}

// GetNextFlightResultsResponse ...
type GetNextFlightResultsResponse struct {
	XMLName               xml.Name               `xml:"http://epowerv5.amadeus.com.tr/WS GetNextFlightResultsResponse"`
	OTAAirLowFareSearchRS *OTAAirLowFareSearchRS `xml:"OTA_AirLowFareSearchRS,omitempty"`
}

// SearchFlightCalendar ...
type SearchFlightCalendar struct {
	XMLName               xml.Name               `xml:"http://epowerv5.amadeus.com.tr/WS SearchFlightCalendar"`
	OTAAirLowFareSearchRQ *OTAAirLowFareSearchRQ `xml:"OTA_AirLowFareSearchRQ,omitempty"`
}

// SearchFlightCalendarResponse ...
type SearchFlightCalendarResponse struct {
	XMLName               xml.Name               `xml:"http://epowerv5.amadeus.com.tr/WS SearchFlightCalendarResponse"`
	OTAAirLowFareSearchRS *OTAAirLowFareSearchRS `xml:"OTA_AirLowFareSearchRS,omitempty"`
}

// CheckETicket ...
type CheckETicket struct {
	XMLName              xml.Name              `xml:"http://epowerv5.amadeus.com.tr/WS CheckETicket"`
	OTAAirCheckETicketRQ *OTAAirCheckETicketRQ `xml:"OTA_AirCheckETicketRQ,omitempty"`
}

// CheckETicketResponse ...
type CheckETicketResponse struct {
	XMLName              xml.Name              `xml:"http://epowerv5.amadeus.com.tr/WS CheckETicketResponse"`
	OTAAirCheckETicketRS *OTAAirCheckETicketRS `xml:"OTA_AirCheckETicketRS,omitempty"`
}

// GetFlightRules ...
type GetFlightRules struct {
	XMLName       xml.Name       `xml:"http://epowerv5.amadeus.com.tr/WS GetFlightRules"`
	OTAAirRulesRQ *OTAAirRulesRQ `xml:"OTA_AirRulesRQ,omitempty"`
}

// GetFlightRulesResponse ...
type GetFlightRulesResponse struct {
	XMLName       xml.Name       `xml:"http://epowerv5.amadeus.com.tr/WS GetFlightRulesResponse"`
	OTAAirRulesRS *OTAAirRulesRS `xml:"OTA_AirRulesRS,omitempty"`
}

// OTAAirBookRQ ...
type OTAAirBookRQ struct {
	*GenericFlightRQ
	Version                 int                      `xml:"Version,attr,omitempty"`
	SequenceNmbr            uint                         `xml:"SequenceNmbr,attr,omitempty"`
	RetransmissionIndicator bool                         `xml:"RetransmissionIndicator,attr,omitempty"`
	PriceInd                bool                         `xml:"PriceInd,attr,omitempty"`
	ReferenceNumber         string                       `xml:"ReferenceNumber,attr,omitempty"`
	ControlNumber           string                       `xml:"ControlNumber,attr,omitempty"`
	CorporatePinNumber      string                       `xml:"CorporatePinNumber,attr,omitempty"`
	EchoToken               string                       `xml:"EchoToken,attr,omitempty"`
	Target                  string                       `xml:"Target,attr,omitempty"`
	TransactionIdentifier   string                       `xml:"TransactionIdentifier,attr,omitempty"`
	TransactionStatusCode   string                       `xml:"TransactionStatusCode,attr,omitempty"`
	TimeStamp               string                    `xml:"TimeStamp,attr,omitempty"`
	Queue                   *BookQueue                   `xml:"Queue,omitempty"`
	Fulfillment             *Fulfillment                 `xml:"Fulfillment,omitempty"`
	BookingReferenceID      *UniqueIDType                `xml:"BookingReferenceID,omitempty"`
	PriceInfo               *BookPriceInfo               `xml:"PriceInfo,omitempty"`
	InsurancePlan           *InsurancePlan               `xml:"InsurancePlan,omitempty"`
	RelatedMember           *RelatedMember               `xml:"RelatedMember,omitempty"`
	TravelerInfo            *TravelerInfoType            `xml:"TravelerInfo,omitempty"`
	AirItinerary            *AirItineraryType            `xml:"AirItinerary,omitempty"`
	TPAExtensions           *TPAExtensionsType           `xml:"TPA_Extensions,omitempty"`
	POS                     *ArrayOfSourceType           `xml:"POS,omitempty"`
	Ticketing               []*BookTicketingInfo         `xml:"Ticketing,omitempty"`
	PNRRemarks              *ArrayOfPNRRemarkInfo        `xml:"PNRRemarks,omitempty"`
	ResStatus               *TransactionActionType       `xml:"ResStatus,attr,omitempty"`
	AirItineraryPricingInfo *AirItineraryPricingInfoType `xml:"AirItineraryPricingInfo,omitempty"`
}

// BookQueue ...
type BookQueue struct {
	PseudoCityCode string `xml:"PseudoCityCode,attr,omitempty"`
	QueueNumber    string `xml:"QueueNumber,attr,omitempty"`
	QueueCategory  string `xml:"QueueCategory,attr,omitempty"`
	SystemCode     string `xml:"SystemCode,attr,omitempty"`
	QueueID        string `xml:"QueueID,attr,omitempty"`
	DateTime       string `xml:"DateTime,attr,omitempty"`
	Text           string `xml:"Text,attr,omitempty"`
}

// BookTicketingInfo ...
type BookTicketingInfo struct {
	*TicketingInfoType
	TicketingVendor *TicketingVendor `xml:"TicketingVendor,omitempty"`
	PricingSystem   *PricingSystem   `xml:"PricingSystem,omitempty"`
}

// PricingSystem ...
type PricingSystem struct {
	CompanyShortName string `xml:"CompanyShortName,attr,omitempty"`
	
	TravelSector string `xml:"TravelSector,attr,omitempty"`
	
	Code string `xml:"Code,attr,omitempty"`
	
	CodeContext string `xml:"CodeContext,attr,omitempty"`
}

// TicketingVendor ...
type TicketingVendor struct {
	CompanyShortName string `xml:"CompanyShortName,attr,omitempty"`
	TravelSector     string `xml:"TravelSector,attr,omitempty"`
	Code             string `xml:"Code,attr,omitempty"`
	CodeContext      string `xml:"CodeContext,attr,omitempty"`
}

// BookPriceInfo ...
type BookPriceInfo struct {
	*BookingPriceInfoType
	
	PricingPref []struct {
		Type string `xml:"Type,attr,omitempty"`
		
		ExcludeInd bool `xml:"ExcludeInd,attr,omitempty"`
	} `xml:"PricingPref,omitempty"`
	
	ParticipationLevel string `xml:"ParticipationLevel,attr,omitempty"`
}

// OTAAirBookRS ...
type OTAAirBookRS struct {
	SequenceNmbr            uint                         `xml:"SequenceNmbr,attr,omitempty"`
	RetransmissionIndicator bool                         `xml:"RetransmissionIndicator,attr,omitempty"`
	PriceInd                bool                         `xml:"PriceInd,attr,omitempty"`
	Cancel                  bool                         `xml:"Cancel,attr,omitempty"`
	PaymentRequestURL       string                       `xml:"PaymentRequestURL,attr,omitempty"`
	EchoToken               string                       `xml:"EchoToken,attr,omitempty"`
	Target                  string                       `xml:"Target,attr,omitempty"`
	ReferenceNumber         string                       `xml:"ReferenceNumber,attr,omitempty"`
	TransactionIdentifier   string                       `xml:"TransactionIdentifier,attr,omitempty"`
	TransactionStatusCode   string                       `xml:"TransactionStatusCode,attr,omitempty"`
	Version                 float64                      `xml:"Version,attr,omitempty"`
	TimeStamp               time.Time                    `xml:"TimeStamp,attr,omitempty"`
	Errors                  *ErrorsType                  `xml:"Errors,omitempty"`
	Success                 *SuccessType                 `xml:"Success,omitempty"`
	Warnings                *WarningsType                `xml:"Warnings,omitempty"`
	AirReservation          *AirReservation              `xml:"AirReservation,omitempty"`
	BookingToBeCancelled    *BookingToBeCancelled        `xml:"BookingToBeCancelled,omitempty"`
	NewPrice                *AirItineraryPricingInfoType `xml:"NewPrice,omitempty"`
	OldPrice                *AirItineraryPricingInfoType `xml:"OldPrice,omitempty"`
}

// AirReservation ...
type AirReservation struct {
	*AirReservationType
}

// BookFlightWithRecommendation ...
type BookFlightWithRecommendation struct {
	XMLName      xml.Name      `xml:"http://epowerv5.amadeus.com.tr/WS BookFlightWithRecommendation"`
	OTAAirBookRQ *OTAAirBookRQ `xml:"OTA_AirBookRQ,omitempty"`
}

// BookFlightWithRecommendationResponse ...
type BookFlightWithRecommendationResponse struct {
	XMLName      xml.Name      `xml:"http://epowerv5.amadeus.com.tr/WS BookFlightWithRecommendationResponse"`
	OTAAirBookRS *OTAAirBookRS `xml:"OTA_AirBookRS,omitempty"`
}

// GetPNR ...
type GetPNR struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetPNR"`
	
	OTA_ReadRQ struct {
		POS *ArrayOfSourceType `xml:"POS,omitempty"`
		
		BookingReferenceID *UniqueIDType `xml:"BookingReferenceID,omitempty"`
		
		ReadRequests struct {
			AirReadRequest struct {
				POS *ArrayOfSourceType `xml:"POS,omitempty"`
				
				Airline *CompanyNameType `xml:"Airline,omitempty"`
				
				FlightNumber string `xml:"FlightNumber,omitempty"`
				
				DepartureAirport *LocationType `xml:"DepartureAirport,omitempty"`
				
				DepartureDate time.Time `xml:"DepartureDate,omitempty"`
				
				Name *PersonNameType `xml:"Name,omitempty"`
				
				Telephone struct {
				} `xml:"Telephone,omitempty"`
				
				CustLoyalty struct {
					ProgramID string `xml:"ProgramID,attr,omitempty"`
					
					MembershipID string `xml:"MembershipID,attr,omitempty"`
					
					TravelSector string `xml:"TravelSector,attr,omitempty"`
					
					RPH string `xml:"RPH,attr,omitempty"`
					
					VendorCode string `xml:"VendorCode,attr,omitempty"`
				} `xml:"CustLoyalty,omitempty"`
				
				CreditCardInfo *PaymentCardType `xml:"CreditCardInfo,omitempty"`
				
				TicketNumber *TicketingInfoRSType `xml:"TicketNumber,omitempty"`
				
				QueueInfo struct {
					Queue []struct {
						PseudoCityCode string `xml:"PseudoCityCode,attr,omitempty"`
						
						QueueNumber string `xml:"QueueNumber,attr,omitempty"`
						
						QueueCategory string `xml:"QueueCategory,attr,omitempty"`
						
						SystemCode string `xml:"SystemCode,attr,omitempty"`
						
						QueueID string `xml:"QueueID,attr,omitempty"`
					} `xml:"Queue,omitempty"`
					
					FirstItemOnlyInd bool `xml:"FirstItemOnlyInd,attr,omitempty"`
					
					RemoveFromQueueInd bool `xml:"RemoveFromQueueInd,attr,omitempty"`
					
					FullDataInd bool `xml:"FullDataInd,attr,omitempty"`
					
					StartDate string `xml:"StartDate,attr,omitempty"`
					
					EndDate string `xml:"EndDate,attr,omitempty"`
				} `xml:"QueueInfo,omitempty"`
				
				TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
				
				SeatNumber string `xml:"SeatNumber,attr,omitempty"`
				
				IncludeFF_EquivPartnerLev bool `xml:"IncludeFF_EquivPartnerLev,attr,omitempty"`
				
				ReturnFF_Number bool `xml:"ReturnFF_Number,attr,omitempty"`
				
				ReturnDownlineSeg bool `xml:"ReturnDownlineSeg,attr,omitempty"`
				
				InfoToReturn string `xml:"InfoToReturn,attr,omitempty"`
				
				FF_RequestCriteria string `xml:"FF_RequestCriteria,attr,omitempty"`
			} `xml:"AirReadRequest,omitempty"`
			
			PkgReadRequest struct {
				Name *PersonNameType `xml:"Name,omitempty"`
				
				ArrivalLocation *LocationType `xml:"ArrivalLocation,omitempty"`
				
				DepartureLocation *LocationType `xml:"DepartureLocation,omitempty"`
				
				TravelCode string `xml:"TravelCode,attr,omitempty"`
				
				TourCode string `xml:"TourCode,attr,omitempty"`
				
				PackageID string `xml:"PackageID,attr,omitempty"`
				
				Start string `xml:"Start,attr,omitempty"`
				
				Duration string `xml:"Duration,attr,omitempty"`
				
				End string `xml:"End,attr,omitempty"`
			} `xml:"PkgReadRequest,omitempty"`
			
			ReadRequest struct {
				BookingReferenceID *UniqueIDType `xml:"BookingReferenceID,omitempty"`
				
				Verification *VerificationType `xml:"Verification,omitempty"`
				
				HistoryRequestedInd bool `xml:"HistoryRequestedInd,attr,omitempty"`
			} `xml:"ReadRequest,omitempty"`
			
			/*	HotelReadRequest struct {
				CityName string `xml:"CityName,omitempty"`
			
				Airport struct {
					LocationCode string `xml:"LocationCode,attr,omitempty"`
			
					CodeContext string `xml:"CodeContext,attr,omitempty"`
			
					AirportName string `xml:"AirportName,attr,omitempty"`
				} `xml:"Airport,omitempty"`
			
				UserID struct {
					*UniqueID_Type
			
					PinNumber string `xml:"PinNumber,attr,omitempty"`
				} `xml:"UserID,omitempty"`
			
				Verification *VerificationType `xml:"Verification,omitempty"`
			
				TPA_Extensions *TPA_ExtensionsType `xml:"TPA_Extensions,omitempty"`
			
				ChainCode string `xml:"ChainCode,attr,omitempty"`
			
				BrandCode string `xml:"BrandCode,attr,omitempty"`
			
				HotelCode string `xml:"HotelCode,attr,omitempty"`
			
				HotelCityCode string `xml:"HotelCityCode,attr,omitempty"`
			
				HotelName string `xml:"HotelName,attr,omitempty"`
			
				HotelCodeContext string `xml:"HotelCodeContext,attr,omitempty"`
			
				ChainName string `xml:"ChainName,attr,omitempty"`
			
				BrandName string `xml:"BrandName,attr,omitempty"`
			} `xml:"HotelReadRequest,omitempty"`*/
			
			GlobalReservationReadRequest struct {
				TravelerName *PersonNameType `xml:"TravelerName,omitempty"`
				
				Start string `xml:"Start,attr,omitempty"`
				
				Duration string `xml:"Duration,attr,omitempty"`
				
				End string `xml:"End,attr,omitempty"`
			} `xml:"GlobalReservationReadRequest,omitempty"`
		} `xml:"ReadRequests,omitempty"`
		
		EchoToken string `xml:"EchoToken,attr,omitempty"`
		
		TimeStamp string `xml:"TimeStamp,attr,omitempty"`
		
		Target string `xml:"Target,attr,omitempty"`
		
		Version float64 `xml:"Version,attr,omitempty"`
		
		TransactionIdentifier string `xml:"TransactionIdentifier,attr,omitempty"`
		
		SequenceNmbr int `xml:"SequenceNmbr,attr,omitempty"`
		
		TransactionStatusCode string `xml:"TransactionStatusCode,attr,omitempty"`
		
		RetransmissionIndicator bool `xml:"RetransmissionIndicator,attr,omitempty"`
		
		ReqRespVersion string `xml:"ReqRespVersion,attr,omitempty"`
		
		ReservationType string `xml:"ReservationType,attr,omitempty"`
		
		ReturnListIndicator bool `xml:"ReturnListIndicator,attr,omitempty"`
		
		MoreIndicator bool `xml:"MoreIndicator,attr,omitempty"`
		
		MoreDataEchoToken string `xml:"MoreDataEchoToken,attr,omitempty"`
	} `xml:"OTA_ReadRQ,omitempty"`
}

// OTAReadRQ ...
type OTAReadRQ struct {
	POS *ArrayOfSourceType `xml:"POS,omitempty"`
	
	BookingReferenceID *UniqueIDType `xml:"BookingReferenceID,omitempty"`
	
	ReadRequests struct {
		AirReadRequest struct {
			POS *ArrayOfSourceType `xml:"POS,omitempty"`
			
			Airline *CompanyNameType `xml:"Airline,omitempty"`
			
			FlightNumber string `xml:"FlightNumber,omitempty"`
			
			DepartureAirport *LocationType `xml:"DepartureAirport,omitempty"`
			
			DepartureDate time.Time `xml:"DepartureDate,omitempty"`
			
			Name *PersonNameType `xml:"Name,omitempty"`
			
			Telephone struct {
			} `xml:"Telephone,omitempty"`
			
			CustLoyalty struct {
				ProgramID string `xml:"ProgramID,attr,omitempty"`
				
				MembershipID string `xml:"MembershipID,attr,omitempty"`
				
				TravelSector string `xml:"TravelSector,attr,omitempty"`
				
				RPH string `xml:"RPH,attr,omitempty"`
				
				VendorCode string `xml:"VendorCode,attr,omitempty"`
			} `xml:"CustLoyalty,omitempty"`
			
			CreditCardInfo *PaymentCardType `xml:"CreditCardInfo,omitempty"`
			
			TicketNumber *TicketingInfoRSType `xml:"TicketNumber,omitempty"`
			
			QueueInfo struct {
				Queue []struct {
					PseudoCityCode string `xml:"PseudoCityCode,attr,omitempty"`
					
					QueueNumber string `xml:"QueueNumber,attr,omitempty"`
					
					QueueCategory string `xml:"QueueCategory,attr,omitempty"`
					
					SystemCode string `xml:"SystemCode,attr,omitempty"`
					
					QueueID string `xml:"QueueID,attr,omitempty"`
				} `xml:"Queue,omitempty"`
				
				FirstItemOnlyInd bool `xml:"FirstItemOnlyInd,attr,omitempty"`
				
				RemoveFromQueueInd bool `xml:"RemoveFromQueueInd,attr,omitempty"`
				
				FullDataInd bool `xml:"FullDataInd,attr,omitempty"`
				
				StartDate string `xml:"StartDate,attr,omitempty"`
				
				EndDate string `xml:"EndDate,attr,omitempty"`
			} `xml:"QueueInfo,omitempty"`
			
			TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
			
			SeatNumber string `xml:"SeatNumber,attr,omitempty"`
			
			IncludeFF_EquivPartnerLev bool `xml:"IncludeFF_EquivPartnerLev,attr,omitempty"`
			
			ReturnFF_Number bool `xml:"ReturnFF_Number,attr,omitempty"`
			
			ReturnDownlineSeg bool `xml:"ReturnDownlineSeg,attr,omitempty"`
			
			InfoToReturn string `xml:"InfoToReturn,attr,omitempty"`
			
			FF_RequestCriteria string `xml:"FF_RequestCriteria,attr,omitempty"`
		} `xml:"AirReadRequest,omitempty"`
		
		PkgReadRequest struct {
			Name *PersonNameType `xml:"Name,omitempty"`
			
			ArrivalLocation *LocationType `xml:"ArrivalLocation,omitempty"`
			
			DepartureLocation *LocationType `xml:"DepartureLocation,omitempty"`
			
			TravelCode string `xml:"TravelCode,attr,omitempty"`
			
			TourCode string `xml:"TourCode,attr,omitempty"`
			
			PackageID string `xml:"PackageID,attr,omitempty"`
			
			Start string `xml:"Start,attr,omitempty"`
			
			Duration string `xml:"Duration,attr,omitempty"`
			
			End string `xml:"End,attr,omitempty"`
		} `xml:"PkgReadRequest,omitempty"`
		
		ReadRequest struct {
			BookingReferenceID *UniqueIDType `xml:"BookingReferenceID,omitempty"`
			
			Verification *VerificationType `xml:"Verification,omitempty"`
			
			HistoryRequestedInd bool `xml:"HistoryRequestedInd,attr,omitempty"`
		} `xml:"ReadRequest,omitempty"`
		
		/*	HotelReadRequest struct {
			CityName string `xml:"CityName,omitempty"`
		
			Airport struct {
				LocationCode string `xml:"LocationCode,attr,omitempty"`
		
				CodeContext string `xml:"CodeContext,attr,omitempty"`
		
				AirportName string `xml:"AirportName,attr,omitempty"`
			} `xml:"Airport,omitempty"`
		
			UserID struct {
				*UniqueID_Type
		
				PinNumber string `xml:"PinNumber,attr,omitempty"`
			} `xml:"UserID,omitempty"`
		
			Verification *VerificationType `xml:"Verification,omitempty"`
		
			TPA_Extensions *TPA_ExtensionsType `xml:"TPA_Extensions,omitempty"`
		
			ChainCode string `xml:"ChainCode,attr,omitempty"`
		
			BrandCode string `xml:"BrandCode,attr,omitempty"`
		
			HotelCode string `xml:"HotelCode,attr,omitempty"`
		
			HotelCityCode string `xml:"HotelCityCode,attr,omitempty"`
		
			HotelName string `xml:"HotelName,attr,omitempty"`
		
			HotelCodeContext string `xml:"HotelCodeContext,attr,omitempty"`
		
			ChainName string `xml:"ChainName,attr,omitempty"`
		
			BrandName string `xml:"BrandName,attr,omitempty"`
		} `xml:"HotelReadRequest,omitempty"`*/
		
		GlobalReservationReadRequest struct {
			TravelerName *PersonNameType `xml:"TravelerName,omitempty"`
			
			Start string `xml:"Start,attr,omitempty"`
			
			Duration string `xml:"Duration,attr,omitempty"`
			
			End string `xml:"End,attr,omitempty"`
		} `xml:"GlobalReservationReadRequest,omitempty"`
	} `xml:"ReadRequests,omitempty"`
	
	EchoToken string `xml:"EchoToken,attr,omitempty"`
	
	TimeStamp string `xml:"TimeStamp,attr,omitempty"`
	
	Target string `xml:"Target,attr,omitempty"`
	
	Version float64 `xml:"Version,attr,omitempty"`
	
	TransactionIdentifier string `xml:"TransactionIdentifier,attr,omitempty"`
	
	SequenceNmbr int `xml:"SequenceNmbr,attr,omitempty"`
	
	TransactionStatusCode string `xml:"TransactionStatusCode,attr,omitempty"`
	
	RetransmissionIndicator bool `xml:"RetransmissionIndicator,attr,omitempty"`
	
	ReqRespVersion string `xml:"ReqRespVersion,attr,omitempty"`
	
	ReservationType string `xml:"ReservationType,attr,omitempty"`
	
	ReturnListIndicator bool `xml:"ReturnListIndicator,attr,omitempty"`
	
	MoreIndicator bool `xml:"MoreIndicator,attr,omitempty"`
	
	MoreDataEchoToken string `xml:"MoreDataEchoToken,attr,omitempty"`
}

// GetPNRResponse ...
type GetPNRResponse struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetPNRResponse"`
	
	OTA_AirBookRS struct {
		NewPrice *AirItineraryPricingInfoType `xml:"NewPrice,omitempty"`
		
		OldPrice *AirItineraryPricingInfoType `xml:"OldPrice,omitempty"`
		
		AirReservation struct {
			*AirReservationType
		} `xml:"AirReservation,omitempty"`
		
		BookingToBeCancelled *BookingToBeCancelled `xml:"BookingToBeCancelled,omitempty"`
		
		Errors *ErrorsType `xml:"Errors,omitempty"`
		
		Success *SuccessType `xml:"Success,omitempty"`
		
		Warnings *WarningsType `xml:"Warnings,omitempty"`
		
		ReferenceNumber string `xml:"ReferenceNumber,attr,omitempty"`
		
		PaymentRequestURL string `xml:"PaymentRequestURL,attr,omitempty"`
		
		Cancel bool `xml:"Cancel,attr,omitempty"`
		
		EchoToken string `xml:"EchoToken,attr,omitempty"`
		
		TimeStamp string `xml:"TimeStamp,attr,omitempty"`
		
		Target string `xml:"Target,attr,omitempty"`
		
		Version float64 `xml:"Version,attr,omitempty"`
		
		TransactionIdentifier string `xml:"TransactionIdentifier,attr,omitempty"`
		
		SequenceNmbr int `xml:"SequenceNmbr,attr,omitempty"`
		
		TransactionStatusCode string `xml:"TransactionStatusCode,attr,omitempty"`
		
		RetransmissionIndicator bool `xml:"RetransmissionIndicator,attr,omitempty"`
		
		PriceInd bool `xml:"PriceInd,attr,omitempty"`
	} `xml:"OTA_AirBookRS,omitempty"`
}

// EditPNR ...
type EditPNR struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS EditPNR"`
	
	EditPNRRQ struct {
		BookingReferenceID *UniqueIDType `xml:"BookingReferenceID,omitempty"`
		
		SurName string `xml:"SurName,omitempty"`
		
		AgentSignature string `xml:"AgentSignature,omitempty"`
		
		PassengerDetailChanges *PassengerDetailChanges `xml:"PassengerDetailChanges,omitempty"`
		
		PNRRemarkChanges *ArrayOfPNRRemarkChange `xml:"PNRRemarkChanges,omitempty"`
		
		OSIElementChanges *ArrayOfOSIElementChange `xml:"OSIElementChanges,omitempty"`
	} `xml:"EditPNRRQ,omitempty"`
}

// EditPNRResponse ...
type EditPNRResponse struct {
	XMLName   xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS EditPNRResponse"`
	EditPnrRS struct {
		AirReservation struct {
			*AirReservationType
		} `xml:"AirReservation,omitempty"`
		Errors   *ErrorsType   `xml:"Errors,omitempty"`
		Success  *SuccessType  `xml:"Success,omitempty"`
		Warnings *WarningsType `xml:"Warnings,omitempty"`
	} `xml:"EditPnrRS,omitempty"`
}

// GetLastTicketingDate ...
type GetLastTicketingDate struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetLastTicketingDate"`
	PnrNo   string   `xml:"PnrNo,omitempty"`
}

// GetLastTicketingDateResponse ...
type GetLastTicketingDateResponse struct {
	XMLName             xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetLastTicketingDateResponse"`
	LastTicketingDateRS struct {
		Success               *SuccessType                  `xml:"Success,omitempty"`
		Errors                *ErrorsType                   `xml:"Errors,omitempty"`
		LastTicketingDateData *ArrayOfLastTicketingDateData `xml:"LastTicketingDateData,omitempty"`
	} `xml:"LastTicketingDateRS,omitempty"`
}

type GetCancelPolicyText struct {
	XMLName   xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetCancelPolicyText"`
	PCityCode string   `xml:"p_CityCode,omitempty"`
}

// GetCancelPolicyTextResponse ...
type GetCancelPolicyTextResponse struct {
	XMLName                   xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetCancelPolicyTextResponse"`
	GetCancelPolicyTextResult string   `xml:"GetCancelPolicyTextResult,omitempty"`
}

// Book ...
type Book struct {
	XMLName      xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS Book"`
	BookBasketRQ struct {
		FlightReservation struct {
			*GenericFlightRQ
		} `xml:"FlightReservation,omitempty"`
		
		HotelReservation struct {
			HotelID int32 `xml:"HotelID,attr,omitempty"`
			
			RoomID int32 `xml:"RoomID,attr,omitempty"`
		} `xml:"HotelReservation,omitempty"`
		
		CarReservation struct {
			VendorID int32 `xml:"VendorID,attr,omitempty"`
			
			VehicleID int32 `xml:"VehicleID,attr,omitempty"`
		} `xml:"CarReservation,omitempty"`
		
		InsuranceReservation struct {
			PlanID int32 `xml:"PlanID,attr,omitempty"`
		} `xml:"InsuranceReservation,omitempty"`
		
		TravelerInfo *TravelerInfoType `xml:"TravelerInfo,omitempty"`
		
		Fulfillment struct {
			PaymentDetails *ArrayOfPaymentDetailType `xml:"PaymentDetails,omitempty"`
			
			DeliveryAddress *AddressType `xml:"DeliveryAddress,omitempty"`
			
			Name *PersonNameType `xml:"Name,omitempty"`
			
			Receipt struct {
				DistribType string `xml:"DistribType,attr,omitempty"`
			} `xml:"Receipt,omitempty"`
			
			PaymentText []struct {
				*FormattedTextTextType
				
				Text string `xml:"Text,attr,omitempty"`
				
				Name string `xml:"Name,attr,omitempty"`
			} `xml:"PaymentText,omitempty"`
			
			LCCUserLogin struct {
				UserName string `xml:"UserName,attr,omitempty"`
				
				Password string `xml:"Password,attr,omitempty"`
			} `xml:"LCCUserLogin,omitempty"`
		} `xml:"Fulfillment,omitempty"`
		
		Ticketing []*TicketingInfoType `xml:"Ticketing,omitempty"`
		
		PNRRemarks *ArrayOfPNRRemarkInfo `xml:"PNRRemarks,omitempty"`
		
		ReferenceNumber string `xml:"ReferenceNumber,attr,omitempty"`
		
		ControlNumber string `xml:"ControlNumber,attr,omitempty"`
		
		CorporatePinNumber string `xml:"CorporatePinNumber,attr,omitempty"`
	} `xml:"BookBasketRQ,omitempty"`
}

// BookResponse ...
type BookResponse struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS BookResponse"`
	
	BookBasketRS struct {
		AirReservation struct {
			*AirReservationType
		} `xml:"AirReservation,omitempty"`
		
		// InsuranceReservation *OTA_InsuranceBookRSPlanForBookRS `xml:"InsuranceReservation,omitempty"`
		
		// HotelReservation *OTA_HotelResRS `xml:"HotelReservation,omitempty"`
		// CarReservation *OTA_VehResRS `xml:"CarReservation,omitempty"`
		
		TravelerInfo *TravelerInfoType `xml:"TravelerInfo,omitempty"`
		
		FullFilment struct {
			PaymentDetails *ArrayOfPaymentDetailType `xml:"PaymentDetails,omitempty"`
			
			DeliveryAddress *AddressType `xml:"DeliveryAddress,omitempty"`
		} `xml:"FullFilment,omitempty"`
		
		BookingReferenceID []*UniqueIDType `xml:"BookingReferenceID,omitempty"`
		
		NewPrice *AirItineraryPricingInfoType `xml:"NewPrice,omitempty"`
		
		Errors *ErrorsType `xml:"Errors,omitempty"`
		
		Success *SuccessType `xml:"Success,omitempty"`
		
		Warnings *WarningsType `xml:"Warnings,omitempty"`
		
		Cancel bool `xml:"Cancel,attr,omitempty"`
		
		ReferenceNumber string `xml:"ReferenceNumber,attr,omitempty"`
	} `xml:"BookBasketRS,omitempty"`
}

// GetBasketPNR ...
type GetBasketPNR struct {
	XMLName    xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetBasketPNR"`
	OTA_ReadRQ struct {
		POS                *ArrayOfSourceType `xml:"POS,omitempty"`
		BookingReferenceID *UniqueIDType      `xml:"BookingReferenceID,omitempty"`
		ReadRequests       struct {
			AirReadRequest struct {
				POS              *ArrayOfSourceType `xml:"POS,omitempty"`
				Airline          *CompanyNameType   `xml:"Airline,omitempty"`
				FlightNumber     string             `xml:"FlightNumber,omitempty"`
				DepartureAirport *LocationType      `xml:"DepartureAirport,omitempty"`
				DepartureDate    time.Time          `xml:"DepartureDate,omitempty"`
				Name             *PersonNameType    `xml:"Name,omitempty"`
				Telephone        struct {
				} `xml:"Telephone,omitempty"`
				CustLoyalty struct {
					ProgramID    string   `xml:"ProgramID,attr,omitempty"`
					MembershipID string   `xml:"MembershipID,attr,omitempty"`
					TravelSector string   `xml:"TravelSector,attr,omitempty"`
					RPH          string   `xml:"RPH,attr,omitempty"`
					VendorCode   []string `xml:"VendorCode,attr,omitempty"`
				} `xml:"CustLoyalty,omitempty"`
				CreditCardInfo *PaymentCardType     `xml:"CreditCardInfo,omitempty"`
				TicketNumber   *TicketingInfoRSType `xml:"TicketNumber,omitempty"`
				QueueInfo      struct {
					Queue []struct {
						PseudoCityCode string `xml:"PseudoCityCode,attr,omitempty"`
						QueueNumber    string `xml:"QueueNumber,attr,omitempty"`
						QueueCategory  string `xml:"QueueCategory,attr,omitempty"`
						SystemCode     string `xml:"SystemCode,attr,omitempty"`
						QueueID        string `xml:"QueueID,attr,omitempty"`
					} `xml:"Queue,omitempty"`
					FirstItemOnlyInd   bool   `xml:"FirstItemOnlyInd,attr,omitempty"`
					RemoveFromQueueInd bool   `xml:"RemoveFromQueueInd,attr,omitempty"`
					FullDataInd        bool   `xml:"FullDataInd,attr,omitempty"`
					StartDate          string `xml:"StartDate,attr,omitempty"`
					EndDate            string `xml:"EndDate,attr,omitempty"`
				} `xml:"QueueInfo,omitempty"`
				TPAExtensions            *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
				SeatNumber               string             `xml:"SeatNumber,attr,omitempty"`
				IncludeFFEquivPartnerLev bool               `xml:"IncludeFF_EquivPartnerLev,attr,omitempty"`
				ReturnFFNumber           bool               `xml:"ReturnFF_Number,attr,omitempty"`
				ReturnDownlineSeg        bool               `xml:"ReturnDownlineSeg,attr,omitempty"`
				InfoToReturn             string             `xml:"InfoToReturn,attr,omitempty"`
				FFRequestCriteria        string             `xml:"FF_RequestCriteria,attr,omitempty"`
			} `xml:"AirReadRequest,omitempty"`
			PkgReadRequest struct {
				Name              *PersonNameType `xml:"Name,omitempty"`
				ArrivalLocation   *LocationType   `xml:"ArrivalLocation,omitempty"`
				DepartureLocation *LocationType   `xml:"DepartureLocation,omitempty"`
				TravelCode        string          `xml:"TravelCode,attr,omitempty"`
				TourCode          string          `xml:"TourCode,attr,omitempty"`
				PackageID         string          `xml:"PackageID,attr,omitempty"`
				Start             string          `xml:"Start,attr,omitempty"`
				Duration          string          `xml:"Duration,attr,omitempty"`
				End               string          `xml:"End,attr,omitempty"`
			} `xml:"PkgReadRequest,omitempty"`
			ReadRequest struct {
				BookingReferenceID  *UniqueIDType     `xml:"BookingReferenceID,omitempty"`
				Verification        *VerificationType `xml:"Verification,omitempty"`
				HistoryRequestedInd bool              `xml:"HistoryRequestedInd,attr,omitempty"`
			} `xml:"ReadRequest,omitempty"`
			GolfReadRequest struct {
				Membership []struct {
					ProgramID    string   `xml:"ProgramID,attr,omitempty"`
					MembershipID string   `xml:"MembershipID,attr,omitempty"`
					TravelSector string   `xml:"TravelSector,attr,omitempty"`
					RPH          string   `xml:"RPH,attr,omitempty"`
					VendorCode   []string `xml:"VendorCode,attr,omitempty"`
				} `xml:"Membership,omitempty"`
				Name         *PersonNameType `xml:"Name,omitempty"`
				ID           string          `xml:"ID,attr,omitempty"`
				RoundID      int           `xml:"RoundID,attr,omitempty"`
				PlayDateTime string          `xml:"PlayDateTime,attr,omitempty"`
				PackageID    string          `xml:"PackageID,attr,omitempty"`
			} `xml:"GolfReadRequest,omitempty"`
			CruiseReadRequest struct {
				SelectedSailing struct {
					VoyageID  string `xml:"VoyageID,attr,omitempty"`
					Status    string `xml:"Status,attr,omitempty"`
					GroupCode string `xml:"GroupCode,attr,omitempty"`
				} `xml:"SelectedSailing,omitempty"`
				GuestInfo           *PersonNameType `xml:"GuestInfo,omitempty"`
				HistoryRequestedInd bool            `xml:"HistoryRequestedInd,attr,omitempty"`
			} `xml:"CruiseReadRequest,omitempty"`
			HotelReadRequest struct {
				CityName string `xml:"CityName,omitempty"`
				Airport  struct {
					LocationCode string `xml:"LocationCode,attr,omitempty"`
					
					CodeContext string `xml:"CodeContext,attr,omitempty"`
					
					AirportName string `xml:"AirportName,attr,omitempty"`
				} `xml:"Airport,omitempty"`
				
				UserID struct {
					*UniqueIDType
					
					PinNumber string `xml:"PinNumber,attr,omitempty"`
				} `xml:"UserID,omitempty"`
				
				Verification *VerificationType `xml:"Verification,omitempty"`
				
				TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
				
				ChainCode string `xml:"ChainCode,attr,omitempty"`
				
				BrandCode string `xml:"BrandCode,attr,omitempty"`
				
				HotelCode string `xml:"HotelCode,attr,omitempty"`
				
				HotelCityCode string `xml:"HotelCityCode,attr,omitempty"`
				
				HotelName string `xml:"HotelName,attr,omitempty"`
				
				HotelCodeContext string `xml:"HotelCodeContext,attr,omitempty"`
				
				ChainName string `xml:"ChainName,attr,omitempty"`
				
				BrandName string `xml:"BrandName,attr,omitempty"`
			} `xml:"HotelReadRequest,omitempty"`
			
			GlobalReservationReadRequest struct {
				TravelerName *PersonNameType `xml:"TravelerName,omitempty"`
				
				Start string `xml:"Start,attr,omitempty"`
				
				Duration string `xml:"Duration,attr,omitempty"`
				
				End string `xml:"End,attr,omitempty"`
			} `xml:"GlobalReservationReadRequest,omitempty"`
		} `xml:"ReadRequests,omitempty"`
		
		EchoToken string `xml:"EchoToken,attr,omitempty"`
		
		TimeStamp string `xml:"TimeStamp,attr,omitempty"`
		
		Target string `xml:"Target,attr,omitempty"`
		
		Version float64 `xml:"Version,attr,omitempty"`
		
		TransactionIdentifier string `xml:"TransactionIdentifier,attr,omitempty"`
		
		SequenceNmbr int `xml:"SequenceNmbr,attr,omitempty"`
		
		TransactionStatusCode string `xml:"TransactionStatusCode,attr,omitempty"`
		
		RetransmissionIndicator bool `xml:"RetransmissionIndicator,attr,omitempty"`
		
		ReqRespVersion string `xml:"ReqRespVersion,attr,omitempty"`
		
		ReservationType string `xml:"ReservationType,attr,omitempty"`
		
		ReturnListIndicator bool `xml:"ReturnListIndicator,attr,omitempty"`
		
		MoreIndicator bool `xml:"MoreIndicator,attr,omitempty"`
		
		MoreDataEchoToken string `xml:"MoreDataEchoToken,attr,omitempty"`
	} `xml:"OTA_ReadRQ,omitempty"`
}

// GetBasketPNRResponse ...
type GetBasketPNRResponse struct {
	XMLName      xml.Name        `xml:"http://epowerv5.amadeus.com.tr/WS GetBasketPNRResponse"`
	OTAAirBookRS *AirReservation `xml:"OTA_AirBookRS,omitempty"`
}

// ExecuteCommand ...
type ExecuteCommand struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS ExecuteCommand"`
	
	AmadeusAPICommandRQ struct {
		Command string `xml:"Command,attr,omitempty"`
		
		Server string `xml:"Server,attr,omitempty"`
		
		PortNumber string `xml:"PortNumber,attr,omitempty"`
		
		CorporateId string `xml:"CorporateId,attr,omitempty"`
		
		UserId string `xml:"UserId,attr,omitempty"`
		
		Password string `xml:"Password,attr,omitempty"`
	} `xml:"AmadeusAPICommandRQ,omitempty"`
}

// ExecuteCommandResponse ...
type ExecuteCommandResponse struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS ExecuteCommandResponse"`
	
	AmadeusAPICommandRS struct {
		Errors *ErrorsType `xml:"Errors,omitempty"`
		
		Success *SuccessType `xml:"Success,omitempty"`
		
		Warnings *WarningsType `xml:"Warnings,omitempty"`
		
		Result string `xml:"Result,attr,omitempty"`
	} `xml:"AmadeusAPICommandRS,omitempty"`
}

// GetAPISRules ...
type GetAPISRules struct {
	XMLName     xml.Name     `xml:"http://epowerv5.amadeus.com.tr/WS GetAPISRules"`
	APISRulesRQ *APISRulesRQ `xml:"APISRulesRQ,omitempty"`
}

// GetAPISRulesResponse ...
type GetAPISRulesResponse struct {
	XMLName xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS GetAPISRulesResponse"`
	
	APISRulesRS *APISRulesRS `xml:"APISRulesRS,omitempty"`
}

// GetFlightFareFamilies ...
type GetFlightFareFamilies struct {
	XMLName          xml.Name          `xml:"http://epowerv5.amadeus.com.tr/WS GetFlightFareFamilies"`
	FareFamilyInfoRQ *FareFamilyInfoRQ `xml:"FareFamilyInfoRQ,omitempty"`
}

// GetFlightFareFamiliesResponse ...
type GetFlightFareFamiliesResponse struct {
	XMLName          xml.Name          `xml:"http://epowerv5.amadeus.com.tr/WS GetFlightFareFamiliesResponse"`
	FareFamilyInfoRS *FareFamilyInfoRS `xml:"FareFamilyInfoRS,omitempty"`
}

// OTACurrencyConversionRQ ...
type OTACurrencyConversionRQ struct {
	FromCurrency string  `xml:"FromCurrency,attr,omitempty"`
	ToCurrency   string  `xml:"ToCurrency,attr,omitempty"`
	Amount       float64 `xml:"Amount,attr,omitempty"`
}

// OTACurrencyConversionRS ...
type OTACurrencyConversionRS struct {
	Errors             *ErrorsType  `xml:"Errors,omitempty"`
	Success            *SuccessType `xml:"Success,omitempty"`
	Amount             float64      `xml:"Amount,attr,omitempty"`
	TruncatedAmount    float64      `xml:"TruncatedAmount,attr,omitempty"`
	OtherChargesAmount float64      `xml:"OtherChargesAmount,attr,omitempty"`
}

// AuthenticationSoapHeader ...
type AuthenticationSoapHeader struct {
	XMLName       xml.Name `xml:"http://epowerv5.amadeus.com.tr/WS AuthenticationSoapHeader"`
	WSUserName    string   `xml:"WSUserName,omitempty"`
	WSPassword    string   `xml:"WSPassword,omitempty"`
	WSCultureInfo string   `xml:"WSCultureInfo,omitempty"`
}
