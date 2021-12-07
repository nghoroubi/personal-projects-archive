package repository

import (
	"encoding/xml"
	"time"
)

// Header ...
type Header struct {
	XMLName                  xml.Name                  `xml:"Header"`
	AuthenticationSoapHeader *AuthenticationSoapHeader `xml:"AuthenticationSoapHeader"`
}

// GenericSoapHeader ...
type GenericSoapHeader struct {
	SubSiteCode string `xml:"SubSiteCode,omitempty"`
}

// UniqueIDType ...
type UniqueIDType struct {
	URL                string           `xml:"URL,attr,omitempty"`
	Type               string           `xml:"Type,attr,omitempty"`
	Instance           string           `xml:"Instance,attr,omitempty"`
	IDContext          string           `xml:"ID_Context,attr,omitempty"`
	SecondarySplitPNRs string           `xml:"SecondarySplitPNRs,attr,omitempty"`
	CompanyName        *CompanyNameType `xml:"CompanyName,omitempty"`
}

// CompanyNameType ...
type CompanyNameType struct {
	Value              string
	LogoUrl            string `xml:"LogoUrl,attr,omitempty"`
	TermsAndConditions string `xml:"TermsAndConditions,attr,omitempty"`
	CompanyShortName   string `xml:"CompanyShortName,attr,omitempty"`
	TravelSector       string `xml:"TravelSector,attr,omitempty"`
	Code               string `xml:"Code,attr,omitempty"`
	CodeContext        string `xml:"CodeContext,attr,omitempty"`
	Division           string `xml:"Division,attr,omitempty"`
	Department         string `xml:"Department,attr,omitempty"`
}

// TravelArrangerType ...
type TravelArrangerType struct {
	*CompanyNameType
	DefaultInd         bool   `xml:"DefaultInd,attr,omitempty"`
	ShareSynchInd      string `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd     string `xml:"ShareMarketInd,attr,omitempty"`
	TravelArrangerType string `xml:"TravelArrangerType,attr,omitempty"`
	RPH                string `xml:"RPH,attr,omitempty"`
}

// CompanyNamePrefType ...
type CompanyNamePrefType struct {
	*CompanyNameType
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

// OperatingAirlineType ...
type OperatingAirlineType struct {
	*CompanyNameType
	FlightNumber     string `xml:"FlightNumber,attr,omitempty"`
	ResBookDesigCode string `xml:"ResBookDesigCode,attr,omitempty"`
	Connection       string `xml:"Connection,attr,omitempty"`
}

// ArrayOfSourceType ...
type ArrayOfSourceType struct {
	Source []*SourceType `xml:"Source,omitempty"`
}

// Position ...
type Position struct {
	Latitude                  string `xml:"Latitude,attr,omitempty"`
	Longitude                 string `xml:"Longitude,attr,omitempty"`
	Altitude                  string `xml:"Altitude,attr,omitempty"`
	AltitudeUnitOfMeasureCode string `xml:"AltitudeUnitOfMeasureCode,attr,omitempty"`
}

// BookingChannel ...
type BookingChannel struct {
	Primary     bool             `xml:"Primary,attr,omitempty"`
	Type        string           `xml:"Type,attr,omitempty"`
	CompanyName *CompanyNameType `xml:"CompanyName,omitempty"`
}

// SourceType ...
type SourceType struct {
	RequestorID struct {
		*UniqueIDType
		MessagePassword string `xml:"MessagePassword,attr,omitempty"`
	} `xml:"RequestorID,omitempty"`
	AgentSine        string          `xml:"AgentSine,attr,omitempty"`
	PseudoCityCode   string          `xml:"PseudoCityCode,attr,omitempty"`
	ISOCountry       string          `xml:"ISOCountry,attr,omitempty"`
	ISOCurrency      string          `xml:"ISOCurrency,attr,omitempty"`
	AgentDutyCode    string          `xml:"AgentDutyCode,attr,omitempty"`
	AirlineVendorID  string          `xml:"AirlineVendorID,attr,omitempty"`
	AirportCode      string          `xml:"AirportCode,attr,omitempty"`
	FirstDepartPoint string          `xml:"FirstDepartPoint,attr,omitempty"`
	ERSPUserID       string          `xml:"ERSP_UserID,attr,omitempty"`
	TerminalID       string          `xml:"TerminalID,attr,omitempty"`
	Position         *Position       `xml:"Position,omitempty"`
	BookingChannel   *BookingChannel `xml:"BookingChannel,omitempty"`
}

// CustomerLoyalty ...
type CustomerLoyalty struct {
	ProgramID    string `xml:"ProgramID,attr,omitempty"`
	MembershipID string `xml:"MembershipID,attr,omitempty"`
	TravelSector string `xml:"TravelSector,attr,omitempty"`
	RPH          string `xml:"RPH,attr,omitempty"`
	VendorCode   string `xml:"VendorCode,attr,omitempty"`
}

// VerificationType ...
type VerificationType struct {
	PersonName struct {
		*PersonNameType
		PartialName bool `xml:"PartialName,attr,omitempty"`
	} `xml:"PersonName,omitempty"`
	Email         *EmailType `xml:"Email,omitempty"`
	TelephoneInfo struct {
		RPH string `xml:"RPH,attr,omitempty"`
	} `xml:"TelephoneInfo,omitempty"`
	PaymentCard         *PaymentCardType   `xml:"PaymentCard,omitempty"`
	AddressInfo         *AddressInfoType   `xml:"AddressInfo,omitempty"`
	CustLoyalty         []*CustomerLoyalty `xml:"CustLoyalty,omitempty"`
	Vendor              []*CompanyNameType `xml:"Vendor,omitempty"`
	ReservationTimeSpan *TimeSpan          `xml:"ReservationTimeSpan,omitempty"`
	AssociatedQuantity  []struct {
		URI string `xml:"URI,attr,omitempty"`
	} `xml:"AssociatedQuantity,omitempty"`
	StartLocation struct {
		*LocationType
		AssociatedDateTime time.Time `xml:"AssociatedDateTime,attr,omitempty"`
	} `xml:"StartLocation,omitempty"`
	
	EndLocation struct {
		*LocationType
		AssociatedDateTime time.Time `xml:"AssociatedDateTime,attr,omitempty"`
	} `xml:"EndLocation,omitempty"`
	TPAExtensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
}

// TimeSpan ...
type TimeSpan struct {
	Start    string `xml:"Start,attr,omitempty"`
	Duration string `xml:"Duration,attr,omitempty"`
	End      string `xml:"End,attr,omitempty"`
}

// PersonNameType ...
type PersonNameType struct {
	NamePrefix     string `xml:"NamePrefix,omitempty"`
	GivenName      string `xml:"GivenName,omitempty"`
	MiddleName     string `xml:"MiddleName,omitempty"`
	SurnamePrefix  string `xml:"SurnamePrefix,omitempty"`
	Surname        string `xml:"Surname,omitempty"`
	NameSuffix     string `xml:"NameSuffix,omitempty"`
	NameTitle      string `xml:"NameTitle,omitempty"`
	ShareSynchInd  string `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	NameType       string `xml:"NameType,attr,omitempty"`
}

// EmailType ...
type EmailType struct {
	Value          string
	ShareSynchInd  string `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	DefaultInd     bool   `xml:"DefaultInd,attr,omitempty"`
	EmailType      string `xml:"EmailType,attr,omitempty"`
	RPH            string `xml:"RPH,attr,omitempty"`
}

// PaymentCardType ...
type PaymentCardType struct {
	CardHolderName string `xml:"CardHolderName,omitempty"`
	CardIssuerName struct {
		BankID string `xml:"BankID,attr,omitempty"`
	} `xml:"CardIssuerName,omitempty"`
	Address                *AddressType `xml:"Address,omitempty"`
	ShareSynchInd          string       `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd         string       `xml:"ShareMarketInd,attr,omitempty"`
	CardType               string       `xml:"CardType,attr,omitempty"`
	SecurityNumber         string       `xml:"SecurityNumber,attr,omitempty"`
	CardCode               string       `xml:"CardCode,attr,omitempty"`
	CardNumber             string       `xml:"CardNumber,attr,omitempty"`
	SeriesCode             string       `xml:"SeriesCode,attr,omitempty"`
	EffectiveDate          string       `xml:"EffectiveDate,attr,omitempty"`
	ExpireDate             time.Time    `xml:"ExpireDate,attr,omitempty"`
	MaskedCardNumber       string       `xml:"MaskedCardNumber,attr,omitempty"`
	CardHolderRPH          string       `xml:"CardHolderRPH,attr,omitempty"`
	ExtendPaymentIndicator bool         `xml:"ExtendPaymentIndicator,attr,omitempty"`
}

// AddressType ...
type AddressType struct {
	ContactName      string        `xml:"ContactName,omitempty"`
	Location         string        `xml:"Location,omitempty"`
	AreaLocationCode string        `xml:"AreaLocationCode,omitempty"`
	TaxNumber        string        `xml:"TaxNumber,omitempty"`
	TaxOffice        string        `xml:"TaxOffice,omitempty"`
	Company          string        `xml:"Company,omitempty"`
	StreetNumber     *StreetNumber `xml:"StreetNmbr,omitempty"`
	BldgRoom         []struct {
		Value             string
		BldgNameIndicator bool `xml:"BldgNameIndicator,attr,omitempty"`
	} `xml:"BldgRoom,omitempty"`
	AddressLine    []string         `xml:"AddressLine,omitempty"`
	CityName       string           `xml:"CityName,omitempty"`
	PostalCode     string           `xml:"PostalCode,omitempty"`
	County         string           `xml:"County,omitempty"`
	StateProv      *StateProvType   `xml:"StateProv,omitempty"`
	CountryName    *CountryNameType `xml:"CountryName,omitempty"`
	FormattedInd   bool             `xml:"FormattedInd,attr,omitempty"`
	ShareSynchInd  string           `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd string           `xml:"ShareMarketInd,attr,omitempty"`
	Type           string           `xml:"Type,attr,omitempty"`
}

// StreetNumber ...
type StreetNumber struct {
	*StreetNmbrType
	StreetNmbrSuffix string `xml:"StreetNmbrSuffix,attr,omitempty"`
	StreetDirection  string `xml:"StreetDirection,attr,omitempty"`
	RuralRouteNmbr   string `xml:"RuralRouteNmbr,attr,omitempty"`
}

// StreetNmbrType ...
type StreetNmbrType struct {
	Value  string
	PO_Box string `xml:"PO_Box,attr,omitempty"`
}

// StateProvType ...
type StateProvType struct {
	Value     string
	StateCode string `xml:"StateCode,attr,omitempty"`
}

// CountryNameType ...
type CountryNameType struct {
	Value string
	Code  string `xml:"Code,attr,omitempty"`
}

// OffLocationServiceCoreTypeAddress ...
type OffLocationServiceCoreTypeAddress struct {
	*AddressType
	SiteID   string `xml:"SiteID,attr,omitempty"`
	SiteName string `xml:"SiteName,attr,omitempty"`
}

// VehicleVendorAddressType ...
type VehicleVendorAddressType struct {
	*AddressType
	OpeningHour string `xml:"OpeningHour,omitempty"`
	ClosingHour string `xml:"ClosingHour,omitempty"`
	FaxNumber   string `xml:"FaxNumber,omitempty"`
	OfficePhone string `xml:"OfficePhone,omitempty"`
}

// AddressInfoType ...
type AddressInfoType struct {
	*AddressType
	DefaultInd bool   `xml:"DefaultInd,attr,omitempty"`
	UseType    string `xml:"UseType,attr,omitempty"`
	RPH        string `xml:"RPH,attr,omitempty"`
}

// LocationType ...
type LocationType struct {
	Value        string
	LocationCode string `xml:"LocationCode,attr,omitempty"`
	CodeContext  string `xml:"CodeContext,attr,omitempty"`
}

// AirportPrefType ...
type AirportPrefType struct {
	*LocationType
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

// TPAExtensionsType ...
type TPAExtensionsType struct {
}

// CancelRuleType ...
type CancelRuleType struct {
	PaymentCard  *PaymentCardType `xml:"PaymentCard,omitempty"`
	CancelByDate string           `xml:"CancelByDate,attr,omitempty"`
	Amount       float64          `xml:"Amount,attr,omitempty"`
	Percent      float64          `xml:"Percent,attr,omitempty"`
	Type         string           `xml:"Type,attr,omitempty"`
}

// ArrayOfFreeTextType ...
type ArrayOfFreeTextType struct {
	Reason []*FreeTextType `xml:"Reason,omitempty"`
}

// ArrayOfWarningType ...
type ArrayOfWarningType struct {
	WarningType []*WarningType `xml:"WarningType,omitempty"`
}

// ParagraphType ...
type ParagraphType struct {
	// Text               *FormattedTextTextType `xml:"Text,omitempty"`
	Text               string `xml:"Text,omitempty"`
	ListItem           *ParagraphTypeListItem `xml:"ListItem,omitempty"`
	Image              string                 `xml:"Image,omitempty"`
	Name               string                 `xml:"Name,attr,omitempty"`
	ParagraphNumber    int                    `xml:"ParagraphNumber,attr,omitempty"`
	CreateDateTime     time.Time              `xml:"CreateDateTime,attr,omitempty"`
	CreatorID          string                 `xml:"CreatorID,attr,omitempty"`
	LastModifyDateTime time.Time              `xml:"LastModifyDateTime,attr,omitempty"`
	LastModifierID     string                 `xml:"LastModifierID,attr,omitempty"`
	PurgeDate          time.Time              `xml:"PurgeDate,attr,omitempty"`
	Language           *Language              `xml:"Language,attr,omitempty"`
}

// FormattedTextTextType ...
type FormattedTextTextType struct {
	Value     string
	Formatted bool      `xml:"Formatted,attr,omitempty"`
	Language  *Language `xml:"Language,attr,omitempty"`
}

// CoverageDetailsType ...
type CoverageDetailsType struct {
	*FormattedTextTextType
	CoverageTextType *CoverageTextType `xml:"CoverageTextType,attr,omitempty"`
}

// RateQualifierTypeRateComment ...
type RateQualifierTypeRateComment struct {
	*FormattedTextTextType
	Name string `xml:"Name,attr,omitempty"`
}

// ParagraphTypeListItem ...
type ParagraphTypeListItem struct {
	*FormattedTextTextType
	ListItem int32 `xml:"ListItem,attr,omitempty"`
}

// ArrayOfAdditionalDetailType ...
type ArrayOfAdditionalDetailType struct {
	AdditionalDetail []*AdditionalDetailType `xml:"AdditionalDetail,omitempty"`
}

// AdditionalDetailType ...
type AdditionalDetailType struct {
	DetailDescription *ParagraphType `xml:"DetailDescription,omitempty"`
	Type              string         `xml:"Type,attr,omitempty"`
	Code              string         `xml:"Code,attr,omitempty"`
	Amount            float64        `xml:"Amount,attr,omitempty"`
}

// ArrayOfRatePlanType ...
type ArrayOfRatePlanType struct {
	RatePlan []*RatePlanType `xml:"RatePlan,omitempty"`
}

// RatePlanInclusion ...
type RatePlanInclusion struct {
	IsDaily                     bool           `xml:"IsDaily,omitempty"`
	IsPackage                   bool           `xml:"IsPackage,omitempty"`
	IsPerPeriod                 bool           `xml:"IsPerPeriod,omitempty"`
	IsPerPerson                 bool           `xml:"IsPerPerson,omitempty"`
	IsPerRoom                   bool           `xml:"IsPerRoom,omitempty"`
	IsTaxIncluded               bool           `xml:"IsTaxIncluded,omitempty"`
	TaxInclusive                bool           `xml:"TaxInclusive,attr,omitempty"`
	ServiceFeeInclusive         bool           `xml:"ServiceFeeInclusive,attr,omitempty"`
	CategoryType                string         `xml:"CategoryType,omitempty"`
	Code                        string         `xml:"Code,omitempty"`
	Currency                    string         `xml:"Currency,omitempty"`
	Percentage                  float64        `xml:"Percentage,omitempty"`
	Amount                      float64        `xml:"Amount,omitempty"`
	RatePlanInclusionDesciption *ParagraphType `xml:"RatePlanInclusionDesciption,omitempty"`
}

// MealsIncluded ...
type MealsIncluded struct {
	Breakfast         bool     `xml:"Breakfast,attr,omitempty"`
	Lunch             bool     `xml:"Lunch,attr,omitempty"`
	Dinner            bool     `xml:"Dinner,attr,omitempty"`
	MealPlanIndicator bool     `xml:"MealPlanIndicator,attr,omitempty"`
	MealPlanCodes     []string `xml:"MealPlanCodes,attr,omitempty"`
}

// RatePlanType ...
type RatePlanType struct {
	IDRequiredInd       bool                         `xml:"ID_RequiredInd,attr,omitempty"`
	PriceViewableInd    bool                         `xml:"PriceViewableInd,attr,omitempty"`
	RatePlanType        string                       `xml:"RatePlanType,attr,omitempty"`
	RatePlanID          string                       `xml:"RatePlanID,attr,omitempty"`
	RatePlanName        string                       `xml:"RatePlanName,attr,omitempty"`
	MarketCode          string                       `xml:"MarketCode,attr,omitempty"`
	QualificationType   string                       `xml:"QualificationType,attr,omitempty"`
	BookingCode         string                       `xml:"BookingCode,attr,omitempty"`
	RatePlanCode        string                       `xml:"RatePlanCode,attr,omitempty"`
	EffectiveDate       time.Time                    `xml:"EffectiveDate,attr,omitempty"`
	ExpireDate          time.Time                    `xml:"ExpireDate,attr,omitempty"`
	RatePlanDescription *ParagraphType               `xml:"RatePlanDescription,omitempty"`
	MealsIncluded       *MealsIncluded               `xml:"MealsIncluded,omitempty"`
	Commission          *CommissionType              `xml:"Commission,omitempty"`
	Guarantee           []*GuaranteeType             `xml:"Guarantee,omitempty"`
	RateIndicator       *RateIndicatorType           `xml:"RateIndicator,attr,omitempty"`
	RestrictionStatus   *RestrictionStatus           `xml:"RestrictionStatus,omitempty"`
	AvailabilityStatus  *RateIndicatorType           `xml:"AvailabilityStatus,attr,omitempty"`
	CancelPenalties     *CancelPenaltiesType         `xml:"CancelPenalties,omitempty"`
	RatePlanInclusion   []*RatePlanInclusion         `xml:"RatePlanInclusion,omitempty"`
	AdditionalDetails   *ArrayOfAdditionalDetailType `xml:"AdditionalDetails,omitempty"`
}

// Deadline ...
type Deadline struct {
	OffsetUnitMultiplier int32         `xml:"OffsetUnitMultiplier,attr,omitempty"`
	AbsoluteDeadline     string        `xml:"AbsoluteDeadline,attr,omitempty"`
	OffsetDropTime       string        `xml:"OffsetDropTime,attr,omitempty"`
	Deadline             time.Time     `xml:"Deadline,attr,omitempty"`
	OffsetTimeUnit       *TimeUnitType `xml:"OffsetTimeUnit,attr,omitempty"`
}

// GuaranteeType ...
type GuaranteeType struct {
	RetributionType      string                                 `xml:"RetributionType,attr,omitempty"`
	GuaranteeCode        string                                 `xml:"GuaranteeCode,attr,omitempty"`
	GuaranteeType        string                                 `xml:"GuaranteeType,attr,omitempty"`
	HoldTime             time.Time                              `xml:"HoldTime,attr,omitempty"`
	Deadline             *Deadline                              `xml:"Deadline,omitempty"`
	GuaranteeDescription []*ParagraphType                       `xml:"GuaranteeDescription,omitempty"`
	Comments             *ArrayOfCommentTypeComment             `xml:"Comments,omitempty"`
	GuaranteesAccepted   *ArrayOfGuaranteeTypeGuaranteeAccepted `xml:"GuaranteesAccepted,omitempty"`
}

// ArrayOfGuaranteeTypeGuaranteeAccepted ...
type ArrayOfGuaranteeTypeGuaranteeAccepted struct {
	GuaranteeAccepted []struct {
		*PaymentFormType
		
		Default bool `xml:"Default,attr,omitempty"`
		
		NoCardHolderInfoReqInd bool `xml:"NoCardHolderInfoReqInd,attr,omitempty"`
		
		NameReqInd bool `xml:"NameReqInd,attr,omitempty"`
		
		AddressReqInd bool `xml:"AddressReqInd,attr,omitempty"`
		
		PhoneReqInd bool `xml:"PhoneReqInd,attr,omitempty"`
		
		InterbankNbrReqInd bool `xml:"InterbankNbrReqInd,attr,omitempty"`
		
		BookingSourceAllowedInd bool `xml:"BookingSourceAllowedInd,attr,omitempty"`
		
		CorpDiscountNbrAllowedInd bool `xml:"CorpDiscountNbrAllowedInd,attr,omitempty"`
		
		GuaranteeTypeCode string `xml:"GuaranteeTypeCode,attr,omitempty"`
		
		GuaranteeID string `xml:"GuaranteeID,attr,omitempty"`
	} `xml:"GuaranteeAccepted,omitempty"`
}

// PaymentFormType ...
type PaymentFormType struct {
	PaymentType   PaymentTypes   `xml:"PaymentType,attr,omitempty"`
	PaymentFPType PaymentFPTypes `xml:"PaymentFPType,attr,omitempty"`
}

// DirectBillType ...
type DirectBillType struct {
	CompanyName struct {
		*CompanyNameType
		ContactName string `xml:"ContactName,attr,omitempty"`
	} `xml:"CompanyName,omitempty"`
	Address        *AddressInfoType `xml:"Address,omitempty"`
	ShareSynchInd  string           `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd string           `xml:"ShareMarketInd,attr,omitempty"`
	DirectBill_ID  string           `xml:"DirectBill_ID,attr,omitempty"`
	BillingNumber  string           `xml:"BillingNumber,attr,omitempty"`
}

// BankAcctType ...
type BankAcctType struct {
	BankAcctName      string `xml:"BankAcctName,omitempty"`
	ShareSynchInd     string `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd    string `xml:"ShareMarketInd,attr,omitempty"`
	BankID            string `xml:"BankID,attr,omitempty"`
	AcctType          string `xml:"AcctType,attr,omitempty"`
	BankAcctNumber    string `xml:"BankAcctNumber,attr,omitempty"`
	ChecksAcceptedInd bool   `xml:"ChecksAcceptedInd,attr,omitempty"`
}

// PaymentDetailType ...
type PaymentDetailType struct {
	*PaymentFormType
}

// ArrayOfCommentTypeComment ...
type ArrayOfCommentTypeComment struct {
	Comment []struct {
		*ParagraphType
		CommentOriginatorCode string `xml:"CommentOriginatorCode,attr,omitempty"`
		GuestViewable         bool   `xml:"GuestViewable,attr,omitempty"`
	} `xml:"Comment,omitempty"`
}

// CancelPenaltiesType ...
type CancelPenaltiesType struct {
	CancelPenalty         []*CancelPenaltyType `xml:"CancelPenalty,omitempty"`
	CancelPolicyIndicator bool                 `xml:"CancelPolicyIndicator,attr,omitempty"`
}

// CancelPenaltyType ...
type CancelPenaltyType struct {
	NonRefundable      bool               `xml:"NonRefundable,attr,omitempty"`
	Mon                bool               `xml:"Mon,attr,omitempty"`
	Tue                bool               `xml:"Tue,attr,omitempty"`
	Weds               bool               `xml:"Weds,attr,omitempty"`
	Thur               bool               `xml:"Thur,attr,omitempty"`
	Fri                bool               `xml:"Fri,attr,omitempty"`
	Sat                bool               `xml:"Sat,attr,omitempty"`
	Sun                bool               `xml:"Sun,attr,omitempty"`
	ConfirmClassCode   string             `xml:"ConfirmClassCode,attr,omitempty"`
	PolicyCode         string             `xml:"PolicyCode,attr,omitempty"`
	Start              string             `xml:"Start,attr,omitempty"`
	Duration           string             `xml:"Duration,attr,omitempty"`
	End                string             `xml:"End,attr,omitempty"`
	RoomTypeCode       string             `xml:"RoomTypeCode,attr,omitempty"`
	Deadline           *Deadline          `xml:"Deadline,omitempty"`
	AmountPercent      *AmountPercentType `xml:"AmountPercent,omitempty"`
	PenaltyDescription []*ParagraphType   `xml:"PenaltyDescription,omitempty"`
}

// AmountPercentType ...
type AmountPercentType struct {
	Taxes         *TaxesType `xml:"Taxes,omitempty"`
	TaxInclusive  bool       `xml:"TaxInclusive,attr,omitempty"`
	FeesInclusive bool       `xml:"FeesInclusive,attr,omitempty"`
	NmbrOfNights  int        `xml:"NmbrOfNights,attr,omitempty"`
	BasisType     string     `xml:"BasisType,attr,omitempty"`
	Percent       float64    `xml:"Percent,attr,omitempty"`
	Amount        float64    `xml:"Amount,attr,omitempty"`
	ApplyAs       string     `xml:"ApplyAs,attr,omitempty"`
}

// TaxesType ...
type TaxesType struct {
	Tax    []*TaxType `xml:"Tax,omitempty"`
	Amount float64    `xml:"Amount,attr,omitempty"`
}

// TaxType ...
type TaxType struct {
	ChargeUnitExempt          uint                     `xml:"ChargeUnitExempt,attr,omitempty"`
	ChargeFrequencyExempt     uint                     `xml:"ChargeFrequencyExempt,attr,omitempty"`
	MaxChargeUnitApplies      uint                     `xml:"MaxChargeUnitApplies,attr,omitempty"`
	MaxChargeFrequencyApplies uint                     `xml:"MaxChargeFrequencyApplies,attr,omitempty"`
	ChargeUnit                string                   `xml:"ChargeUnit,attr,omitempty"`
	ChargeFrequency           string                   `xml:"ChargeFrequency,attr,omitempty"`
	Code                      string                   `xml:"Code,attr,omitempty"`
	Percent                   float64                  `xml:"Percent,attr,omitempty"`
	Amount                    float64                  `xml:"Amount,attr,omitempty"`
	EffectiveDate             time.Time                `xml:"EffectiveDate,attr,omitempty"`
	ExpireDate                time.Time                `xml:"ExpireDate,attr,omitempty"`
	TaxDescription            []*ParagraphType         `xml:"TaxDescription,omitempty"`
	Type                      *AmountDeterminationType `xml:"Type,attr,omitempty"`
}

// CommissionType ...
type CommissionType struct {
	UniqueID             *UniqueIDType `xml:"UniqueID,omitempty"`
	CommissionableAmount struct {
		Amount                float64 `xml:"Amount,attr,omitempty"`
		TaxInclusiveIndicator bool    `xml:"TaxInclusiveIndicator,attr,omitempty"`
	} `xml:"CommissionableAmount,omitempty"`
	PrepaidAmount struct {
		Amount float64 `xml:"Amount,attr,omitempty"`
	} `xml:"PrepaidAmount,omitempty"`
	FlatCommission struct {
		Amount float64 `xml:"Amount,attr,omitempty"`
	} `xml:"FlatCommission,omitempty"`
	CommissionPayableAmount struct {
		Amount float64 `xml:"Amount,attr,omitempty"`
	} `xml:"CommissionPayableAmount,omitempty"`
	Comment                  *ParagraphType `xml:"Comment,omitempty"`
	StatusType               string         `xml:"StatusType,attr,omitempty"`
	Percent                  float64        `xml:"Percent,attr,omitempty"`
	CurrencyCode             string         `xml:"CurrencyCode,attr,omitempty"`
	DecimalPlaces            int            `xml:"DecimalPlaces,attr,omitempty"`
	ReasonCode               string         `xml:"ReasonCode,attr,omitempty"`
	BillToID                 string         `xml:"BillToID,attr,omitempty"`
	Frequency                string         `xml:"Frequency,attr,omitempty"`
	MaxCommissionUnitApplies int            `xml:"MaxCommissionUnitApplies,attr,omitempty"`
	CapAmount                float64        `xml:"CapAmount,attr,omitempty"`
}

// AmountType ...
type AmountType struct {
	Base                   *TotalType `xml:"Base,omitempty"`
	AdditionalGuestAmounts struct {
		AdditionalGuestAmount []*AdditionalGuestAmountType `xml:"AdditionalGuestAmount,omitempty"`
		AmountBeforeTax       float64                      `xml:"AmountBeforeTax,attr,omitempty"`
		AmountAfterTax        float64                      `xml:"AmountAfterTax,attr,omitempty"`
		CurrencyCode          string                       `xml:"CurrencyCode,attr,omitempty"`
		DecimalPlaces         int                          `xml:"DecimalPlaces,attr,omitempty"`
	} `xml:"AdditionalGuestAmounts,omitempty"`
	Fees            *ArrayOfFeeType                              `xml:"Fees,omitempty"`
	CancelPolicies  *CancelPenaltiesType                         `xml:"CancelPolicies,omitempty"`
	PaymentPolicies *ArrayOfRequiredPaymentsTypeGuaranteePayment `xml:"PaymentPolicies,omitempty"`
	Discount        []struct {
		*DiscountType
		AppliesTo string `xml:"AppliesTo,attr,omitempty"`
		ItemRPH   string `xml:"ItemRPH,attr,omitempty"`
	} `xml:"Discount,omitempty"`
	Total             *TotalType     `xml:"Total,omitempty"`
	RateDescription   *ParagraphType `xml:"RateDescription,omitempty"`
	AdditionalCharges struct {
		AdditionalCharge []struct {
			Amount          *TotalType `xml:"Amount,omitempty"`
			RoomAmenityCode string     `xml:"RoomAmenityCode,attr,omitempty"`
			Quantity        int        `xml:"Quantity,attr,omitempty"`
		} `xml:"AdditionalCharge,omitempty"`
		AmountBeforeTax float64 `xml:"AmountBeforeTax,attr,omitempty"`
		AmountAfterTax  float64 `xml:"AmountAfterTax,attr,omitempty"`
		CurrencyCode    string  `xml:"CurrencyCode,attr,omitempty"`
		DecimalPlaces   int     `xml:"DecimalPlaces,attr,omitempty"`
	} `xml:"AdditionalCharges,omitempty"`
	AdvanceBookingRestriction struct {
		Start                   string        `xml:"Start,attr,omitempty"`
		Duration                string        `xml:"Duration,attr,omitempty"`
		End                     string        `xml:"End,attr,omitempty"`
		MinAdvanceBookingOffset time.Duration `xml:"MinAdvanceBookingOffset,attr,omitempty"`
		MaxAdvanceBookingOffset time.Duration `xml:"MaxAdvanceBookingOffset,attr,omitempty"`
		Mon                     bool          `xml:"Mon,attr,omitempty"`
		Tue                     bool          `xml:"Tue,attr,omitempty"`
		Weds                    bool          `xml:"Weds,attr,omitempty"`
		Thur                    bool          `xml:"Thur,attr,omitempty"`
		Fri                     bool          `xml:"Fri,attr,omitempty"`
		Sat                     bool          `xml:"Sat,attr,omitempty"`
		Sun                     bool          `xml:"Sun,attr,omitempty"`
	} `xml:"AdvanceBookingRestriction,omitempty"`
	EffectiveDate        time.Time      `xml:"EffectiveDate,attr,omitempty"`
	ExpireDate           time.Time      `xml:"ExpireDate,attr,omitempty"`
	AgeQualifyingCode    string         `xml:"AgeQualifyingCode,attr,omitempty"`
	MinAge               int32          `xml:"MinAge,attr,omitempty"`
	MaxAge               int32          `xml:"MaxAge,attr,omitempty"`
	AgeTimeUnit          *TimeUnitType  `xml:"AgeTimeUnit,attr,omitempty"`
	GuaranteedInd        bool           `xml:"GuaranteedInd,attr,omitempty"`
	NumberOfUnits        int32          `xml:"NumberOfUnits,attr,omitempty"`
	RateTimeUnit         *TimeUnitType  `xml:"RateTimeUnit,attr,omitempty"`
	UnitMultiplier       int            `xml:"UnitMultiplier,attr,omitempty"`
	MinGuestApplicable   int            `xml:"MinGuestApplicable,attr,omitempty"`
	MaxGuestApplicable   int            `xml:"MaxGuestApplicable,attr,omitempty"`
	MinLOS               int            `xml:"MinLOS,attr,omitempty"`
	MaxLOS               int            `xml:"MaxLOS,attr,omitempty"`
	StayOverDate         *DayOfWeekType `xml:"StayOverDate,attr,omitempty"`
	AlternateCurrencyInd bool           `xml:"AlternateCurrencyInd,attr,omitempty"`
}

// TotalType ...
type TotalType struct {
	DecimalPlaces                   uint       `xml:"DecimalPlaces,attr,omitempty"`
	AdditionalFeesExcludedIndicator bool       `xml:"AdditionalFeesExcludedIndicator,attr,omitempty"`
	IsDailyPrice                    bool       `xml:"IsDailyPrice,omitempty"`
	IsRateChange                    bool       `xml:"IsRateChange,omitempty"`
	AmountBeforeTax                 float64    `xml:"AmountBeforeTax,attr,omitempty"`
	AmountAfterTax                  float64    `xml:"AmountAfterTax,attr,omitempty"`
	MarkupAmount                    float64    `xml:"MarkupAmount,attr,omitempty"`
	CurrencyCode                    string     `xml:"CurrencyCode,attr,omitempty"`
	Taxes                           *TaxesType `xml:"Taxes,omitempty"`
}

// DiscountType ...
type DiscountType struct {
	*TotalType
	RestrictedDisplayIndicator bool           `xml:"RestrictedDisplayIndicator,attr,omitempty"`
	TaxInclusive               bool           `xml:"TaxInclusive,attr,omitempty"`
	Percent                    float64        `xml:"Percent,attr,omitempty"`
	DiscountCode               string         `xml:"DiscountCode,attr,omitempty"`
	DiscountReason             *ParagraphType `xml:"DiscountReason,omitempty"`
}

// AdditionalGuestAmountType ...
type AdditionalGuestAmountType struct {
	Amount                  *TotalType       `xml:"Amount,omitempty"`
	AddlGuestAmtDescription []*ParagraphType `xml:"AddlGuestAmtDescription,omitempty"`
	MaxAdditionalGuests     int32            `xml:"MaxAdditionalGuests,attr,omitempty"`
	AgeQualifyingCode       string           `xml:"AgeQualifyingCode,attr,omitempty"`
	MinAge                  int32            `xml:"MinAge,attr,omitempty"`
	MaxAge                  int32            `xml:"MaxAge,attr,omitempty"`
	AgeTimeUnit             *TimeUnitType    `xml:"AgeTimeUnit,attr,omitempty"`
	Type                    string           `xml:"Type,attr,omitempty"`
	Percent                 float64          `xml:"Percent,attr,omitempty"`
	RPH                     string           `xml:"RPH,attr,omitempty"`
}

// ArrayOfFeeType ...
type ArrayOfFeeType struct {
	Fee []*FeeType `xml:"Fee,omitempty"`
}

// FeeType ...
type FeeType struct {
	Taxes                     *TaxesType               `xml:"Taxes,omitempty"`
	Description               []*ParagraphType         `xml:"Description,omitempty"`
	CurrencyCode              string                   `xml:"CurrencyCode,attr,omitempty"`
	Amount                    float64                  `xml:"Amount,attr,omitempty"`
	TaxInclusive              bool                     `xml:"TaxInclusive,attr,omitempty"`
	Type                      *AmountDeterminationType `xml:"Type,attr,omitempty"`
	Code                      string                   `xml:"Code,attr,omitempty"`
	Percent                   float64                  `xml:"Percent,attr,omitempty"`
	EffectiveDate             time.Time                `xml:"EffectiveDate,attr,omitempty"`
	ExpireDate                time.Time                `xml:"ExpireDate,attr,omitempty"`
	MandatoryIndicator        bool                     `xml:"MandatoryIndicator,attr,omitempty"`
	RPH                       string                   `xml:"RPH,attr,omitempty"`
	ChargeUnit                string                   `xml:"ChargeUnit,attr,omitempty"`
	ChargeFrequency           string                   `xml:"ChargeFrequency,attr,omitempty"`
	ChargeUnitExempt          int                      `xml:"ChargeUnitExempt,attr,omitempty"`
	ChargeFrequencyExempt     int                      `xml:"ChargeFrequencyExempt,attr,omitempty"`
	MaxChargeUnitApplies      int                      `xml:"MaxChargeUnitApplies,attr,omitempty"`
	MaxChargeFrequencyApplies int                      `xml:"MaxChargeFrequencyApplies,attr,omitempty"`
	TaxableIndicator          bool                     `xml:"TaxableIndicator,attr,omitempty"`
}

// ArrayOfRequiredPaymentsTypeGuaranteePayment ...
type ArrayOfRequiredPaymentsTypeGuaranteePayment struct {
	GuaranteePayment []struct {
		AcceptedPayments *ArrayOfPaymentFormType `xml:"AcceptedPayments,omitempty"`
		AmountPercent    *AmountPercentType      `xml:"AmountPercent,omitempty"`
		Deadline         []struct {
			AbsoluteDeadline     string        `xml:"AbsoluteDeadline,attr,omitempty"`
			OffsetTimeUnit       *TimeUnitType `xml:"OffsetTimeUnit,attr,omitempty"`
			OffsetUnitMultiplier int32         `xml:"OffsetUnitMultiplier,attr,omitempty"`
			OffsetDropTime       string        `xml:"OffsetDropTime,attr,omitempty"`
		} `xml:"Deadline,omitempty"`
		Description []*ParagraphType `xml:"Description,omitempty"`
		Address     []struct {
			*AddressInfoType
			AddresseeName string `xml:"AddresseeName,attr,omitempty"`
		} `xml:"Address,omitempty"`
		TPAExtensions          *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
		RetributionType        string             `xml:"RetributionType,attr,omitempty"`
		PaymentCode            string             `xml:"PaymentCode,attr,omitempty"`
		Type                   string             `xml:"Type,attr,omitempty"`
		GuaranteeCode          string             `xml:"GuaranteeCode,attr,omitempty"`
		GuaranteeType          string             `xml:"GuaranteeType,attr,omitempty"`
		HoldTime               time.Time          `xml:"HoldTime,attr,omitempty"`
		Mon                    bool               `xml:"Mon,attr,omitempty"`
		Tue                    bool               `xml:"Tue,attr,omitempty"`
		Weds                   bool               `xml:"Weds,attr,omitempty"`
		Thur                   bool               `xml:"Thur,attr,omitempty"`
		Fri                    bool               `xml:"Fri,attr,omitempty"`
		Sat                    bool               `xml:"Sat,attr,omitempty"`
		Sun                    bool               `xml:"Sun,attr,omitempty"`
		Start                  string             `xml:"Start,attr,omitempty"`
		Duration               string             `xml:"Duration,attr,omitempty"`
		End                    string             `xml:"End,attr,omitempty"`
		NoCardHolderInfoReqInd bool               `xml:"NoCardHolderInfoReqInd,attr,omitempty"`
		NameInd                bool               `xml:"NameInd,attr,omitempty"`
		AddressInd             bool               `xml:"AddressInd,attr,omitempty"`
		PhoneInd               bool               `xml:"PhoneInd,attr,omitempty"`
		InterbankNbrInd        bool               `xml:"InterbankNbrInd,attr,omitempty"`
		RoomTypeCode           string             `xml:"RoomTypeCode,attr,omitempty"`
		InfoSource             string             `xml:"InfoSource,attr,omitempty"`
		NonRefundableIndicator bool               `xml:"NonRefundableIndicator,attr,omitempty"`
		PolicyCode             string             `xml:"PolicyCode,attr,omitempty"`
	} `xml:"GuaranteePayment,omitempty"`
}

// ArrayOfPaymentFormType ...
type ArrayOfPaymentFormType struct {
	AcceptedPayment []*PaymentFormType `xml:"AcceptedPayment,omitempty"`
}

// ArrayOfRoomRateTypeFeature ...
type ArrayOfRoomRateTypeFeature struct {
	Feature []struct {
		Description  []*ParagraphType `xml:"Description,omitempty"`
		RoomAmenity  string           `xml:"RoomAmenity,attr,omitempty"`
		Quantity     int32            `xml:"Quantity,attr,omitempty"`
		RoomViewCode string           `xml:"RoomViewCode,attr,omitempty"`
	} `xml:"Feature,omitempty"`
}

// ArrayOfRoomStayTypeRoomRatesRoomRateRestriction ...
type ArrayOfRoomStayTypeRoomRatesRoomRateRestriction struct {
	Restriction []struct {
		DOW_Restrictions *DOWRestrictionsType `xml:"DOW_Restrictions,omitempty"`
		EffectiveDate    time.Time            `xml:"EffectiveDate,attr,omitempty"`
		ExpireDate       time.Time            `xml:"ExpireDate,attr,omitempty"`
	} `xml:"Restriction,omitempty"`
}

// DOWRestrictionsType ...
type DOWRestrictionsType struct {
	AvailableDaysOfWeek struct {
		Mon  bool `xml:"Mon,attr,omitempty"`
		Tue  bool `xml:"Tue,attr,omitempty"`
		Weds bool `xml:"Weds,attr,omitempty"`
		Thur bool `xml:"Thur,attr,omitempty"`
		Fri  bool `xml:"Fri,attr,omitempty"`
		Sat  bool `xml:"Sat,attr,omitempty"`
		Sun  bool `xml:"Sun,attr,omitempty"`
	} `xml:"AvailableDaysOfWeek,omitempty"`
	ArrivalDaysOfWeek struct {
		Mon  bool `xml:"Mon,attr,omitempty"`
		Tue  bool `xml:"Tue,attr,omitempty"`
		Weds bool `xml:"Weds,attr,omitempty"`
		Thur bool `xml:"Thur,attr,omitempty"`
		Fri  bool `xml:"Fri,attr,omitempty"`
		Sat  bool `xml:"Sat,attr,omitempty"`
		Sun  bool `xml:"Sun,attr,omitempty"`
	} `xml:"ArrivalDaysOfWeek,omitempty"`
	DepartureDaysOfWeek struct {
		Mon  bool `xml:"Mon,attr,omitempty"`
		Tue  bool `xml:"Tue,attr,omitempty"`
		Weds bool `xml:"Weds,attr,omitempty"`
		Thur bool `xml:"Thur,attr,omitempty"`
		Fri  bool `xml:"Fri,attr,omitempty"`
		Sat  bool `xml:"Sat,attr,omitempty"`
		Sun  bool `xml:"Sun,attr,omitempty"`
	} `xml:"DepartureDaysOfWeek,omitempty"`
	RequiredDaysOfWeek struct {
		Mon  bool `xml:"Mon,attr,omitempty"`
		Tue  bool `xml:"Tue,attr,omitempty"`
		Weds bool `xml:"Weds,attr,omitempty"`
		Thur bool `xml:"Thur,attr,omitempty"`
		Fri  bool `xml:"Fri,attr,omitempty"`
		Sat  bool `xml:"Sat,attr,omitempty"`
		Sun  bool `xml:"Sun,attr,omitempty"`
	} `xml:"RequiredDaysOfWeek,omitempty"`
}

// GuestCountType ...
type GuestCountType struct {
	GuestCount []struct {
		AgeQualifyingCode string `xml:"AgeQualifyingCode,attr,omitempty"`
		Age               int32  `xml:"Age,attr,omitempty"`
		Count             int32  `xml:"Count,attr,omitempty"`
	} `xml:"GuestCount,omitempty"`
	IsPerRoom bool `xml:"IsPerRoom,attr,omitempty"`
}

// DateTimeSpanType ...
type DateTimeSpanType struct {
	DateWindowRange *TimeInstantType `xml:"DateWindowRange,omitempty"`
	EndDateWindow   struct {
		EarliestDate string         `xml:"EarliestDate,attr,omitempty"`
		LatestDate   string         `xml:"LatestDate,attr,omitempty"`
		DOW          *DayOfWeekType `xml:"DOW,attr,omitempty"`
	} `xml:"EndDateWindow,omitempty"`
	StartDateWindow struct {
		EarliestDate string         `xml:"EarliestDate,attr,omitempty"`
		LatestDate   string         `xml:"LatestDate,attr,omitempty"`
		DOW          *DayOfWeekType `xml:"DOW,attr,omitempty"`
	} `xml:"StartDateWindow,omitempty"`
	Start     string    `xml:"Start,attr,omitempty"`
	StartDate time.Time `xml:"StartDate,attr,omitempty"`
	Duration  string    `xml:"Duration,attr,omitempty"`
	End       string    `xml:"End,attr,omitempty"`
	EndDate   time.Time `xml:"EndDate,attr,omitempty"`
}

// TimeInstantType ...
type TimeInstantType struct {
	Value                     string
	WindowBefore              time.Duration `xml:"WindowBefore,attr,omitempty"`
	WindowAfter               time.Duration `xml:"WindowAfter,attr,omitempty"`
	CrossDateAllowedIndicator bool          `xml:"CrossDateAllowedIndicator,attr,omitempty"`
}

// BasicPropertyInfoType ...
type BasicPropertyInfoType struct {
	Guarentee              string         `xml:"Guarentee,omitempty"`
	Stay                   string         `xml:"Stay,omitempty"`
	Deposit                string         `xml:"Deposit,omitempty"`
	Safety                 string         `xml:"Safety,omitempty"`
	FireSafety             string         `xml:"FireSafety,omitempty"`
	Transportation         string         `xml:"Transportation,omitempty"`
	Tax                    string         `xml:"Tax,omitempty"`
	Meals                  string         `xml:"Meals,omitempty"`
	Images                 *ArrayOfString `xml:"Images,omitempty"`
	ExtraCharges           string         `xml:"ExtraCharges,omitempty"`
	BestTransportationType string         `xml:"BestTransportationType,omitempty"`
	AddRoomOccupant        string         `xml:"AddRoomOccupant,omitempty"`
	ThumbnailImage         string         `xml:"ThumbnailImage,omitempty"`
	Position               struct {
		Latitude                  string `xml:"Latitude,attr,omitempty"`
		Longitude                 string `xml:"Longitude,attr,omitempty"`
		Altitude                  string `xml:"Altitude,attr,omitempty"`
		AltitudeUnitOfMeasureCode string `xml:"AltitudeUnitOfMeasureCode,attr,omitempty"`
	} `xml:"Position,omitempty"`
	Address        *AddressInfoType                           `xml:"Address,omitempty"`
	ContactNumbers *ArrayOfBasicPropertyInfoTypeContactNumber `xml:"ContactNumbers,omitempty"`
	Award          []struct {
		Provider string `xml:"Provider,attr,omitempty"`
		Rating   string `xml:"Rating,attr,omitempty"`
	} `xml:"Award,omitempty"`
	Introduction     string                `xml:"Introduction,omitempty"`
	RelativePosition *RelativePositionType `xml:"RelativePosition,omitempty"`
	HotelAmenity     []struct {
		Code string `xml:"Code,attr,omitempty"`
	} `xml:"HotelAmenity,omitempty"`
	Recreation []struct {
		Code string `xml:"Code,attr,omitempty"`
	} `xml:"Recreation,omitempty"`
	Service []struct {
		BusinessServiceCode string `xml:"BusinessServiceCode,attr,omitempty"`
	} `xml:"Service,omitempty"`
	Policy                   string `xml:"Policy,omitempty"`
	HotelID                  int32  `xml:"HotelID,attr,omitempty"`
	ChainCode                string `xml:"ChainCode,attr,omitempty"`
	BrandCode                string `xml:"BrandCode,attr,omitempty"`
	HotelCode                string `xml:"HotelCode,attr,omitempty"`
	HotelCityCode            string `xml:"HotelCityCode,attr,omitempty"`
	HotelName                string `xml:"HotelName,attr,omitempty"`
	HotelCodeContext         string `xml:"HotelCodeContext,attr,omitempty"`
	ChainName                string `xml:"ChainName,attr,omitempty"`
	BrandName                string `xml:"BrandName,attr,omitempty"`
	HotelSegmentCategoryCode string `xml:"HotelSegmentCategoryCode,attr,omitempty"`
	SupplierIntegrationLevel int    `xml:"SupplierIntegrationLevel,attr,omitempty"`
}

// ArrayOfString ...
type ArrayOfString struct {
	Image []string `xml:"Image,omitempty"`
}

// ArrayOfBasicPropertyInfoTypeContactNumber ...
type ArrayOfBasicPropertyInfoTypeContactNumber struct {
	ContactNumber []struct {
		RPH         string `xml:"RPH,attr,omitempty"`
		PhoneNumber string `xml:"PhoneNumber,attr,omitempty"`
	} `xml:"ContactNumber,omitempty"`
}

// RelativePositionType ...
type RelativePositionType struct {
	*TransportationsType
	Direction              string                `xml:"Direction,attr,omitempty"`
	Distance               string                `xml:"Distance,attr,omitempty"`
	DistanceUnitName       *DistanceUnitNameType `xml:"DistanceUnitName,attr,omitempty"`
	UnitOfMeasureCode      string                `xml:"UnitOfMeasureCode,attr,omitempty"`
	Nearest                bool                  `xml:"Nearest,attr,omitempty"`
	IndexPointCode         string                `xml:"IndexPointCode,attr,omitempty"`
	Name                   string                `xml:"Name,attr,omitempty"`
	PrimaryIndicator       bool                  `xml:"PrimaryIndicator,attr,omitempty"`
	ToFrom                 string                `xml:"ToFrom,attr,omitempty"`
	ApproximateDistanceInd bool                  `xml:"ApproximateDistanceInd,attr,omitempty"`
}

// TransportationsType ...
type TransportationsType struct {
	Transportations *ArrayOfTransportationTypeTransportation `xml:"Transportations,omitempty"`
}

// ArrayOfTransportationTypeTransportation ...
type ArrayOfTransportationTypeTransportation struct {
	Transportation []struct {
		MultimediaDescriptions *MultimediaDescriptionsType `xml:"MultimediaDescriptions,omitempty"`
		OperationSchedules     *OperationSchedulesType     `xml:"OperationSchedules,omitempty"`
		DescriptiveText        string                      `xml:"DescriptiveText,omitempty"`
		NotificationRequired   string                      `xml:"NotificationRequired,attr,omitempty"`
		TransportationCode     string                      `xml:"TransportationCode,attr,omitempty"`
		ChargeUnit             string                      `xml:"ChargeUnit,attr,omitempty"`
		Included               bool                        `xml:"Included,attr,omitempty"`
		CodeDetail             string                      `xml:"CodeDetail,attr,omitempty"`
		Description            string                      `xml:"Description,attr,omitempty"`
		TypicalTravelTime      string                      `xml:"TypicalTravelTime,attr,omitempty"`
		Amount                 float64                     `xml:"Amount,attr,omitempty"`
		ExistsCode             string                      `xml:"ExistsCode,attr,omitempty"`
		ID                     string                      `xml:"ID,attr,omitempty"`
	} `xml:"Transportation,omitempty"`
}

// MultimediaDescriptionsType ...
type MultimediaDescriptionsType struct {
	MultimediaDescription []*MultimediaDescriptionType `xml:"MultimediaDescription,omitempty"`
	LastUpdated           time.Time                    `xml:"LastUpdated,attr,omitempty"`
}

// MultimediaDescriptionType ...
type MultimediaDescriptionType struct {
	VideoItems           *ArrayOfVideoItemsTypeVideoItem `xml:"VideoItems,omitempty"`
	ImageItems           *ArrayOfImageItemsTypeImageItem `xml:"ImageItems,omitempty"`
	TextItems            *ArrayOfTextItemsTypeTextItem   `xml:"TextItems,omitempty"`
	InfoCode             string                          `xml:"InfoCode,attr,omitempty"`
	AdditionalDetailCode string                          `xml:"AdditionalDetailCode,attr,omitempty"`
	ID                   string                          `xml:"ID,attr,omitempty"`
}

// ArrayOfVideoItemsTypeVideoItem ...
type ArrayOfVideoItemsTypeVideoItem struct {
	VideoItemsTypeVideoItem []struct {
		*VideoDescriptionType
		Language           *Language `xml:"Language,attr,omitempty"`
		Caption            string    `xml:"Caption,attr,omitempty"`
		Removal            bool      `xml:"Removal,attr,omitempty"`
		Version            string    `xml:"Version,attr,omitempty"`
		CreateDateTime     time.Time `xml:"CreateDateTime,attr,omitempty"`
		CreatorID          string    `xml:"CreatorID,attr,omitempty"`
		LastModifyDateTime time.Time `xml:"LastModifyDateTime,attr,omitempty"`
		LastModifierID     string    `xml:"LastModifierID,attr,omitempty"`
		PurgeDate          time.Time `xml:"PurgeDate,attr,omitempty"`
	} `xml:"VideoItemsTypeVideoItem,omitempty"`
}

// VideoDescriptionType ...
type VideoDescriptionType struct {
	VideoFormat []struct {
		*VideoItemType
		ContentID       string `xml:"ContentID,attr,omitempty"`
		Title           string `xml:"Title,attr,omitempty"`
		Author          string `xml:"Author,attr,omitempty"`
		CopyrightNotice string `xml:"CopyrightNotice,attr,omitempty"`
		CopyrightOwner  string `xml:"CopyrightOwner,attr,omitempty"`
		CopyrightStart  string `xml:"CopyrightStart,attr,omitempty"`
		CopyrightEnd    string `xml:"CopyrightEnd,attr,omitempty"`
		EffectiveStart  string `xml:"EffectiveStart,attr,omitempty"`
		EffectiveEnd    string `xml:"EffectiveEnd,attr,omitempty"`
		ApplicableStart string `xml:"ApplicableStart,attr,omitempty"`
		ApplicableEnd   string `xml:"ApplicableEnd,attr,omitempty"`
		RecordID        string `xml:"RecordID,attr,omitempty"`
		SourceID        string `xml:"SourceID,attr,omitempty"`
		ID              string `xml:"ID,attr,omitempty"`
	} `xml:"VideoFormat,omitempty"`
	Category string `xml:"Category,attr,omitempty"`
}

// VideoItemType ...
type VideoItemType struct {
	URL               string    `xml:"URL,omitempty"`
	UnitOfMeasureCode string    `xml:"UnitOfMeasureCode,attr,omitempty"`
	Width             int       `xml:"Width,attr,omitempty"`
	Height            int       `xml:"Height,attr,omitempty"`
	BitRate           int       `xml:"BitRate,attr,omitempty"`
	Length            int       `xml:"Length,attr,omitempty"`
	Language          *Language `xml:"Language,attr,omitempty"`
	Format            string    `xml:"Format,attr,omitempty"`
	FileSize          int       `xml:"FileSize,attr,omitempty"`
	FileName          string    `xml:"FileName,attr,omitempty"`
}

// ArrayOfImageItemsTypeImageItem ...
type ArrayOfImageItemsTypeImageItem struct {
	ImageItemsTypeImageItem []struct {
		*ImageDescriptionType
		CreateDateTime     time.Time `xml:"CreateDateTime,attr,omitempty"`
		CreatorID          string    `xml:"CreatorID,attr,omitempty"`
		LastModifyDateTime time.Time `xml:"LastModifyDateTime,attr,omitempty"`
		LastModifierID     string    `xml:"LastModifierID,attr,omitempty"`
		PurgeDate          time.Time `xml:"PurgeDate,attr,omitempty"`
		Removal            bool      `xml:"Removal,attr,omitempty"`
		Version            string    `xml:"Version,attr,omitempty"`
		ID                 string    `xml:"ID,attr,omitempty"`
	} `xml:"ImageItemsTypeImageItem,omitempty"`
}

// ImageDescriptionType ...
type ImageDescriptionType struct {
	ImageFormat []struct {
		*ImageItemType
		ContentID           string    `xml:"ContentID,attr,omitempty"`
		Title               string    `xml:"Title,attr,omitempty"`
		Author              string    `xml:"Author,attr,omitempty"`
		CopyrightNotice     string    `xml:"CopyrightNotice,attr,omitempty"`
		CopyrightOwner      string    `xml:"CopyrightOwner,attr,omitempty"`
		CopyrightStart      string    `xml:"CopyrightStart,attr,omitempty"`
		CopyrightEnd        string    `xml:"CopyrightEnd,attr,omitempty"`
		EffectiveStart      string    `xml:"EffectiveStart,attr,omitempty"`
		EffectiveEnd        string    `xml:"EffectiveEnd,attr,omitempty"`
		ApplicableStart     string    `xml:"ApplicableStart,attr,omitempty"`
		ApplicableEnd       string    `xml:"ApplicableEnd,attr,omitempty"`
		RecordID            string    `xml:"RecordID,attr,omitempty"`
		SourceID            string    `xml:"SourceID,attr,omitempty"`
		Language            *Language `xml:"Language,attr,omitempty"`
		Format              string    `xml:"Format,attr,omitempty"`
		FileName            string    `xml:"FileName,attr,omitempty"`
		FileSize            int       `xml:"FileSize,attr,omitempty"`
		DimensionCategory   string    `xml:"DimensionCategory,attr,omitempty"`
		IsOriginalIndicator bool      `xml:"IsOriginalIndicator,attr,omitempty"`
	} `xml:"ImageFormat,omitempty"`
	Description []struct {
		*FormattedTextTextType
		Caption string `xml:"Caption,attr,omitempty"`
	} `xml:"Description,omitempty"`
	Category string `xml:"Category,attr,omitempty"`
}

// ImageItemType ...
type ImageItemType struct {
	URL               string `xml:"URL,omitempty"`
	UnitOfMeasureCode string `xml:"UnitOfMeasureCode,attr,omitempty"`
	Width             int    `xml:"Width,attr,omitempty"`
	Height            int    `xml:"Height,attr,omitempty"`
}

// ArrayOfTextItemsTypeTextItem ...
type ArrayOfTextItemsTypeTextItem struct {
	TextItemsTypeTextItem []struct {
		*TextDescriptionType
		CreateDateTime     time.Time `xml:"CreateDateTime,attr,omitempty"`
		CreatorID          string    `xml:"CreatorID,attr,omitempty"`
		LastModifyDateTime time.Time `xml:"LastModifyDateTime,attr,omitempty"`
		LastModifierID     string    `xml:"LastModifierID,attr,omitempty"`
		PurgeDate          time.Time `xml:"PurgeDate,attr,omitempty"`
		Removal            bool      `xml:"Removal,attr,omitempty"`
		Version            string    `xml:"Version,attr,omitempty"`
	} `xml:"TextItemsTypeTextItem,omitempty"`
}

// TextDescriptionType ...
type TextDescriptionType struct {
	Description struct {
		*FormattedTextTextType
		ListItem int32 `xml:"ListItem,attr,omitempty"`
	} `xml:"Description,omitempty"`
	URL             string    `xml:"URL,omitempty"`
	Category        string    `xml:"Category,attr,omitempty"`
	ContentID       string    `xml:"ContentID,attr,omitempty"`
	Title           string    `xml:"Title,attr,omitempty"`
	Author          string    `xml:"Author,attr,omitempty"`
	CopyrightNotice string    `xml:"CopyrightNotice,attr,omitempty"`
	CopyrightOwner  string    `xml:"CopyrightOwner,attr,omitempty"`
	CopyrightStart  string    `xml:"CopyrightStart,attr,omitempty"`
	CopyrightEnd    string    `xml:"CopyrightEnd,attr,omitempty"`
	EffectiveStart  string    `xml:"EffectiveStart,attr,omitempty"`
	EffectiveEnd    string    `xml:"EffectiveEnd,attr,omitempty"`
	ApplicableStart string    `xml:"ApplicableStart,attr,omitempty"`
	ApplicableEnd   string    `xml:"ApplicableEnd,attr,omitempty"`
	RecordID        string    `xml:"RecordID,attr,omitempty"`
	SourceID        string    `xml:"SourceID,attr,omitempty"`
	Language        *Language `xml:"Language,attr,omitempty"`
}

// OperationSchedulesType ...
type OperationSchedulesType struct {
	OperationSchedule []*OperationScheduleType `xml:"OperationSchedule,omitempty"`
	Start             string                   `xml:"Start,attr,omitempty"`
	Duration          string                   `xml:"Duration,attr,omitempty"`
	End               string                   `xml:"End,attr,omitempty"`
}

// OperationScheduleType ...
type OperationScheduleType struct {
	OperationTimes *ArrayOfOperationScheduleTypeOperationTime `xml:"OperationTimes,omitempty"`
	Start          string                                     `xml:"Start,attr,omitempty"`
	Duration       string                                     `xml:"Duration,attr,omitempty"`
	End            string                                     `xml:"End,attr,omitempty"`
}

type ArrayOfOperationScheduleTypeOperationTime struct {
	OperationTime []struct {
		Mon  bool `xml:"Mon,attr,omitempty"`
		Tue  bool `xml:"Tue,attr,omitempty"`
		Weds bool `xml:"Weds,attr,omitempty"`
		
		Thur bool `xml:"Thur,attr,omitempty"`
		
		Fri bool `xml:"Fri,attr,omitempty"`
		
		Sat bool `xml:"Sat,attr,omitempty"`
		
		Sun bool `xml:"Sun,attr,omitempty"`
		
		Start string `xml:"Start,attr,omitempty"`
		
		Duration string `xml:"Duration,attr,omitempty"`
		
		End string `xml:"End,attr,omitempty"`
		
		AdditionalOperationInfoCode string `xml:"AdditionalOperationInfoCode,attr,omitempty"`
		
		Frequency string `xml:"Frequency,attr,omitempty"`
		
		Text string `xml:"Text,attr,omitempty"`
	} `xml:"OperationTime,omitempty"`
}

type PropertyValueMatchType struct {
	*BasicPropertyInfoType
	
	SearchValueMatch []struct {
		Value string
		
		Match bool `xml:"Match,attr,omitempty"`
		
		Relevance float64 `xml:"Relevance,attr,omitempty"`
	} `xml:"SearchValueMatch,omitempty"`
	
	Amenities *ArrayOfPropertyValueMatchTypeAmenity `xml:"Amenities,omitempty"`
	
	RateRange struct {
		MinRate float64 `xml:"MinRate,attr,omitempty"`
		
		MaxRate float64 `xml:"MaxRate,attr,omitempty"`
		
		FixedRate float64 `xml:"FixedRate,attr,omitempty"`
		
		RateTimeUnit *TimeUnitType `xml:"RateTimeUnit,attr,omitempty"`
		
		InfoSource string `xml:"InfoSource,attr,omitempty"`
		
		TaxRate float64 `xml:"TaxRate,attr,omitempty"`
		
		RateInfoNotAvailableInd bool `xml:"RateInfoNotAvailableInd,attr,omitempty"`
	} `xml:"RateRange,omitempty"`
	
	MoreDataEchoToken string `xml:"MoreDataEchoToken,attr,omitempty"`
	
	SameCountryInd bool `xml:"SameCountryInd,attr,omitempty"`
	
	AvailabilityStatus *RateIndicatorType `xml:"AvailabilityStatus,attr,omitempty"`
	
	HotelCurrency string `xml:"HotelCurrency,attr,omitempty"`
}

type ArrayOfPropertyValueMatchTypeAmenity struct {
	Amenity []struct {
		PropertyAmenityType string `xml:"PropertyAmenityType,attr,omitempty"`
	} `xml:"Amenity,omitempty"`
}

type ServiceFeesType struct {
	ServiceFee []*ServiceFeeType `xml:"ServiceFee,omitempty"`
	
	BookingFee *BookingFeeType `xml:"BookingFee,omitempty"`
	
	DeliveryFee *DeliveryFeeType `xml:"DeliveryFee,omitempty"`
	
	PaymentFee *PaymentFeeType `xml:"PaymentFee,omitempty"`
	
	AirlineOBFee *AirlineOBFeeType `xml:"AirlineOBFee,omitempty"`
	
	CCOBFee *CCOBFeeType `xml:"CCOBFee,omitempty"`
	
	OtherExtraCost []*OtherExtraCostType `xml:"OtherExtraCost,omitempty"`
}

type ServiceFeeType struct {
	PassengerTypeQuantity *PassengerTypeQuantityType `xml:"PassengerTypeQuantity,omitempty"`
	
	Amount float64 `xml:"Amount,attr,omitempty"`
	
	SurChargeAmount float64 `xml:"SurChargeAmount,attr,omitempty"`
	
	ExternalServiceFee float64 `xml:"ExternalServiceFee,attr,omitempty"`
	
	ItineraryIndex int32 `xml:"ItineraryIndex,attr,omitempty"`
	
	TotalForAllPax bool `xml:"TotalForAllPax,attr,omitempty"`
	
	MarkupFeeAmount float64 `xml:"MarkupFeeAmount,attr,omitempty"`
	
	ExtraSeatAmount float64 `xml:"ExtraSeatAmount,attr,omitempty"`
	
	CabinBaggageAmount float64 `xml:"CabinBaggageAmount,attr,omitempty"`
}

type BookingFeeType struct {
	*ServiceFeeType
	
	Discount float64 `xml:"Discount,attr,omitempty"`
}

type OtherExtraCostType struct {
	*BookingFeeType
	
	CostType string `xml:"CostType,attr,omitempty"`
}

type CCOBFeeType struct {
	*BookingFeeType
	
	MinOBFee float64 `xml:"MinOBFee,attr,omitempty"`
	
	MaxOBFee float64 `xml:"MaxOBFee,attr,omitempty"`
}

type AirlineOBFeeType struct {
	*BookingFeeType
}

type PaymentFeeType struct {
	*BookingFeeType
}

type DeliveryFeeType struct {
	*BookingFeeType
}

type ArrayOfServiceRPHsTypeServiceRPH struct {
	ServiceRPH []struct {
		RPH string `xml:"RPH,attr,omitempty"`
		
		IsPerRoom bool `xml:"IsPerRoom,attr,omitempty"`
	} `xml:"ServiceRPH,omitempty"`
}

type ArrayOfServicesTypeService struct {
	Service []struct {
		Price []*AmountType `xml:"Price,omitempty"`
		
		ServiceDetails *ResCommonDetailType `xml:"ServiceDetails,omitempty"`
		
		TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
		
		ServicePricingType *PricingType `xml:"ServicePricingType,attr,omitempty"`
		
		ReservationStatusType *PMS_ResStatusType `xml:"ReservationStatusType,attr,omitempty"`
		
		ServiceRPH string `xml:"ServiceRPH,attr,omitempty"`
		
		ServiceInventoryCode string `xml:"ServiceInventoryCode,attr,omitempty"`
		
		RatePlanCode string `xml:"RatePlanCode,attr,omitempty"`
		
		InventoryBlockCode string `xml:"InventoryBlockCode,attr,omitempty"`
		
		PriceGuaranteed bool `xml:"PriceGuaranteed,attr,omitempty"`
		
		Inclusive bool `xml:"Inclusive,attr,omitempty"`
		
		Quantity int32 `xml:"Quantity,attr,omitempty"`
		
		RequestedIndicator bool `xml:"RequestedIndicator,attr,omitempty"`
	} `xml:"Service,omitempty"`
}

type ResCommonDetailType struct {
	GuestCounts *GuestCountType `xml:"GuestCounts,omitempty"`
	
	TimeSpan *DateTimeSpanType `xml:"TimeSpan,omitempty"`
	
	ResGuestRPHs *ArrayOfResGuestRPHsTypeResGuestRPH `xml:"ResGuestRPHs,omitempty"`
	
	Memberships *ArrayOfMembershipTypeMembership `xml:"Memberships,omitempty"`
	
	Comments *ArrayOfCommentTypeComment `xml:"Comments,omitempty"`
	
	SpecialRequests *ArrayOfSpecialRequestTypeSpecialRequest `xml:"SpecialRequests,omitempty"`
	
	Guarantee *GuaranteeType `xml:"Guarantee,omitempty"`
	
	DepositPayments *ArrayOfRequiredPaymentsTypeGuaranteePayment `xml:"DepositPayments,omitempty"`
	
	CancelPenalties *CancelPenaltiesType `xml:"CancelPenalties,omitempty"`
	
	Fees *ArrayOfFeeType `xml:"Fees,omitempty"`
	
	Total *TotalType `xml:"Total,omitempty"`
}

type ArrayOfResGuestRPHsTypeResGuestRPH struct {
	ResGuestRPH []struct {
		RPH string `xml:"RPH,attr,omitempty"`
	} `xml:"ResGuestRPH,omitempty"`
}

type ArrayOfMembershipTypeMembership struct {
	Membership []struct {
		ProgramCode string `xml:"ProgramCode,attr,omitempty"`
		
		BonusCode string `xml:"BonusCode,attr,omitempty"`
		
		AccountID string `xml:"AccountID,attr,omitempty"`
		
		PointsEarned int32 `xml:"PointsEarned,attr,omitempty"`
		
		TravelSector string `xml:"TravelSector,attr,omitempty"`
	} `xml:"Membership,omitempty"`
}

type ArrayOfSpecialRequestTypeSpecialRequest struct {
	SpecialRequest []struct {
		*ParagraphType
		
		RequestCode string `xml:"RequestCode,attr,omitempty"`
		
		CodeContext string `xml:"CodeContext,attr,omitempty"`
		
		NumberOfUnits int32 `xml:"NumberOfUnits,attr,omitempty"`
	} `xml:"SpecialRequest,omitempty"`
}

type ResGlobalInfoType struct {
	*ResCommonDetailType
	
	HotelReservationIDs *ArrayOfHotelReservationIDsTypeHotelReservationID `xml:"HotelReservationIDs,omitempty"`
	
	RoutingHops *ArrayOfRoutingHopTypeRoutingHop `xml:"RoutingHops,omitempty"`
	
	Profiles *ArrayOfProfilesTypeProfileInfo `xml:"Profiles,omitempty"`
	
	BookingRules *ArrayOfBookingRulesTypeBookingRule `xml:"BookingRules,omitempty"`
}

type ArrayOfHotelReservationIDsTypeHotelReservationID struct {
	HotelReservationID []struct {
		ResID_Type string `xml:"ResID_Type,attr,omitempty"`
		
		ResID_Value string `xml:"ResID_Value,attr,omitempty"`
		
		ResID_Source string `xml:"ResID_Source,attr,omitempty"`
		
		ResID_SourceContext string `xml:"ResID_SourceContext,attr,omitempty"`
		
		ResID_Date time.Time `xml:"ResID_Date,attr,omitempty"`
		
		ForGuest bool `xml:"ForGuest,attr,omitempty"`
		
		ResGuestRPH string `xml:"ResGuestRPH,attr,omitempty"`
		
		CancelOriginatorCode string `xml:"CancelOriginatorCode,attr,omitempty"`
		
		CancellationDate time.Time `xml:"CancellationDate,attr,omitempty"`
	} `xml:"HotelReservationID,omitempty"`
}

type ArrayOfRoutingHopTypeRoutingHop struct {
	RoutingHop []struct {
		SystemCode string `xml:"SystemCode,attr,omitempty"`
		
		LocalRefID string `xml:"LocalRefID,attr,omitempty"`
		
		TimeStamp string `xml:"TimeStamp,attr,omitempty"`
		
		Comment string `xml:"Comment,attr,omitempty"`
		
		SequenceNmbr int32 `xml:"SequenceNmbr,attr,omitempty"`
		
		Data string `xml:"Data,attr,omitempty"`
	} `xml:"RoutingHop,omitempty"`
}

type ArrayOfProfilesTypeProfileInfo struct {
	ProfileInfo []struct {
		UniqueID *UniqueIDType `xml:"UniqueID,omitempty"`
		
		Profile *ProfileType `xml:"Profile,omitempty"`
	} `xml:"ProfileInfo,omitempty"`
}

type ProfileType struct {
	Accesses           *AccessesType              `xml:"Accesses,omitempty"`
	Customer           *CustomerType              `xml:"Customer,omitempty"`
	PrefCollections    *PreferencesType           `xml:"PrefCollections,omitempty"`
	CompanyInfo        *CompanyInfoType           `xml:"CompanyInfo,omitempty"`
	Agreements         *AgreementsType            `xml:"Agreements,omitempty"`
	Comments           *ArrayOfCommentTypeComment `xml:"Comments,omitempty"`
	TPA_Extensions     *TPAExtensionsType         `xml:"TPA_Extensions,omitempty"`
	ShareAllSynchInd   *YesNoType                 `xml:"ShareAllSynchInd,attr,omitempty"`
	ShareAllMarketInd  *YesNoType                 `xml:"ShareAllMarketInd,attr,omitempty"`
	ProfileType        string                     `xml:"ProfileType,attr,omitempty"`
	CreateDateTime     time.Time                  `xml:"CreateDateTime,attr,omitempty"`
	CreatorID          string                     `xml:"CreatorID,attr,omitempty"`
	LastModifyDateTime time.Time                  `xml:"LastModifyDateTime,attr,omitempty"`
	LastModifierID     string                     `xml:"LastModifierID,attr,omitempty"`
	PurgeDate          time.Time                  `xml:"PurgeDate,attr,omitempty"`
	RPH                string                     `xml:"RPH,attr,omitempty"`
}

type AccessesType struct {
	Access []struct {
		AccessPerson *PersonNameType `xml:"AccessPerson,omitempty"`
		
		AccessComment *FreeTextType `xml:"AccessComment,omitempty"`
		
		ActionType string `xml:"ActionType,attr,omitempty"`
		
		ActionDateTime time.Time `xml:"ActionDateTime,attr,omitempty"`
		
		ID string `xml:"ID,attr,omitempty"`
	} `xml:"Access,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	CreateDateTime time.Time `xml:"CreateDateTime,attr,omitempty"`
}

type CustomerType struct {
	PersonName *PersonNameType `xml:"PersonName,omitempty"`
	
	Telephone []struct {
		EffectiveDate time.Time `xml:"EffectiveDate,attr,omitempty"`
		
		ExpireDate time.Time `xml:"ExpireDate,attr,omitempty"`
		
		RPH string `xml:"RPH,attr,omitempty"`
		
		PhoneNumber string `xml:"PhoneNumber,attr,omitempty"`
		
		AreaCityCode string `xml:"AreaCityCode,attr,omitempty"`
		
		CountryCode string `xml:"CountryCode,attr,omitempty"`
	} `xml:"Telephone,omitempty"`
	
	Email []*EmailType `xml:"Email,omitempty"`
	
	Address []struct {
		*AddressInfoType
		
		CompanyName *CompanyNameType `xml:"CompanyName,omitempty"`
		
		EffectiveDate time.Time `xml:"EffectiveDate,attr,omitempty"`
		
		ExpireDate time.Time `xml:"ExpireDate,attr,omitempty"`
	} `xml:"Address,omitempty"`
	
	URL []*URL_Type `xml:"URL,omitempty"`
	
	CitizenCountryName []struct {
		Code string `xml:"Code,attr,omitempty"`
	} `xml:"CitizenCountryName,omitempty"`
	
	PhysChallName []string `xml:"PhysChallName,omitempty"`
	
	PetInfo []string `xml:"PetInfo,omitempty"`
	
	PaymentForm []*PaymentFormType `xml:"PaymentForm,omitempty"`
	
	RelatedTraveler []*RelatedTravelerType `xml:"RelatedTraveler,omitempty"`
	
	ContactPerson []*ContactPersonType `xml:"ContactPerson,omitempty"`
	
	Document []*DocumentType `xml:"Document,omitempty"`
	
	CustLoyalty []struct {
		ProgramID string `xml:"ProgramID,attr,omitempty"`
		
		MembershipID string `xml:"MembershipID,attr,omitempty"`
		
		TravelSector string `xml:"TravelSector,attr,omitempty"`
		
		RPH string `xml:"RPH,attr,omitempty"`
		
		VendorCode []string `xml:"VendorCode,attr,omitempty"`
	} `xml:"CustLoyalty,omitempty"`
	
	EmployeeInfo []*EmployeeInfoType `xml:"EmployeeInfo,omitempty"`
	
	EmployerInfo *CompanyNameType `xml:"EmployerInfo,omitempty"`
	
	TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
	
	Gender string `xml:"Gender,attr,omitempty"`
	
	Deceased bool `xml:"Deceased,attr,omitempty"`
	
	LockoutType string `xml:"LockoutType,attr,omitempty"`
	
	BirthDate time.Time `xml:"BirthDate,attr,omitempty"`
	
	CurrencyCode string `xml:"CurrencyCode,attr,omitempty"`
	
	DecimalPlaces int `xml:"DecimalPlaces,attr,omitempty"`
	
	VIP_Indicator bool `xml:"VIP_Indicator,attr,omitempty"`
	
	Text string `xml:"Text,attr,omitempty"`
	
	Language *Language `xml:"Language,attr,omitempty"`
}

type URL_Type struct {
	Value string
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	Type string `xml:"Type,attr,omitempty"`
	
	DefaultInd bool `xml:"DefaultInd,attr,omitempty"`
}

type RelatedTravelerType struct {
	UniqueID *UniqueIDType `xml:"UniqueID,omitempty"`
	
	PersonName *PersonNameType `xml:"PersonName,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	Relation string `xml:"Relation,attr,omitempty"`
}

type ContactPersonType struct {
	PersonName *PersonNameType `xml:"PersonName,omitempty"`
	
	Telephone []struct {
		RPH string `xml:"RPH,attr,omitempty"`
	} `xml:"Telephone,omitempty"`
	
	Address []*AddressInfoType `xml:"Address,omitempty"`
	
	Email []*EmailType `xml:"Email,omitempty"`
	
	URL []*URL_Type `xml:"URL,omitempty"`
	
	CompanyName []*CompanyNameType `xml:"CompanyName,omitempty"`
	
	EmployeeInfo []*EmployeeInfoType `xml:"EmployeeInfo,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	DefaultInd bool `xml:"DefaultInd,attr,omitempty"`
	
	ContactType string `xml:"ContactType,attr,omitempty"`
	
	Relation string `xml:"Relation,attr,omitempty"`
	
	EmergencyFlag bool `xml:"EmergencyFlag,attr,omitempty"`
	
	RPH string `xml:"RPH,attr,omitempty"`
	
	CommunicationMethodCode string `xml:"CommunicationMethodCode,attr,omitempty"`
	
	DocumentDistribMethodCode string `xml:"DocumentDistribMethodCode,attr,omitempty"`
}

type EmployeeInfoType struct {
	Value string
	
	EmployeeId string `xml:"EmployeeId,attr,omitempty"`
	
	EmployeeLevel string `xml:"EmployeeLevel,attr,omitempty"`
	
	EmployeeTitle string `xml:"EmployeeTitle,attr,omitempty"`
	
	EmployeeStatus string `xml:"EmployeeStatus,attr,omitempty"`
}

type DocumentType struct {
	DocIssueLocation  string `xml:"DocIssueLocation,attr,omitempty"`
	DocID             string `xml:"DocID,attr,omitempty"`
	DocType           string `xml:"DocType,attr,omitempty"`
	InnerDocType      string `xml:"InnerDocType,attr,omitempty"`
	Gender            string `xml:"Gender,attr,omitempty"`
	BirthDate         string `xml:"BirthDate,attr,omitempty"`
	ExpireDate        string `xml:"ExpireDate,attr,omitempty"`
	DocIssueStateProv string `xml:"DocIssueStateProv,attr,omitempty"`
	DocIssueCountry   string `xml:"DocIssueCountry,attr,omitempty"`
	CityCode          string `xml:"CityCode,attr,omitempty"`
	State             string `xml:"State,attr,omitempty"`
	ZipPostalCode     string `xml:"ZipPostalCode,attr,omitempty"`
	BirthCountry      string `xml:"BirthCountry,attr,omitempty"`
}

type ArrayOfString1 struct {
	AdditionalPersonName []string `xml:"AdditionalPersonName,omitempty"`
}

type CustomerPrimaryAdditionalTypeAdditional struct {
	*CustomerType
	
	Start string `xml:"Start,attr,omitempty"`
	
	Duration string `xml:"Duration,attr,omitempty"`
	
	End string `xml:"End,attr,omitempty"`
	
	CorpDiscountName string `xml:"CorpDiscountName,attr,omitempty"`
	
	CorpDiscountNmbr string `xml:"CorpDiscountNmbr,attr,omitempty"`
	
	QualificationMethod *CustomerPrimaryAdditionalTypeAdditionalQualificationMethod `xml:"QualificationMethod,attr,omitempty"`
	
	Age int32 `xml:"Age,attr,omitempty"`
	
	Code string `xml:"Code,attr,omitempty"`
	
	CodeContext string `xml:"CodeContext,attr,omitempty"`
	
	URI string `xml:"URI,attr,omitempty"`
}

type CustomerPrimaryAdditionalTypePrimary struct {
	*CustomerType
	
	CustomerID *UniqueIDType `xml:"CustomerID,omitempty"`
}

type PreferencesType struct {
	PrefCollection []struct {
		CommonPref []*CommonPrefType `xml:"CommonPref,omitempty"`
		
		// VehicleRentalPref []*VehicleProfileRentalPrefType `xml:"VehicleRentalPref,omitempty"`
		
		AirlinePref []*AirlinePrefType `xml:"AirlinePref,omitempty"`
		
		// HotelPref []*HotelPrefType `xml:"HotelPref,omitempty"`
		
		// OtherSrvcPref []*OtherSrvcPrefType `xml:"OtherSrvcPref,omitempty"`
		
		TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
		
		ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
		
		ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
		
		TravelPurpose string `xml:"TravelPurpose,attr,omitempty"`
	} `xml:"PrefCollection,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
}

type CommonPrefType struct {
	NamePref []*NamePrefType `xml:"NamePref,omitempty"`
	
	PhonePref []*PhonePrefType `xml:"PhonePref,omitempty"`
	
	AddressPref []*AddressPrefType `xml:"AddressPref,omitempty"`
	
	PaymentFormPref []*PaymentFormPrefType `xml:"PaymentFormPref,omitempty"`
	
	InterestPref []*InterestPrefType `xml:"InterestPref,omitempty"`
	
	InsurancePref []*InsurancePrefType `xml:"InsurancePref,omitempty"`
	
	SeatingPref []*SeatingPrefType `xml:"SeatingPref,omitempty"`
	
	TicketDistribPref []*TicketDistribPrefType `xml:"TicketDistribPref,omitempty"`
	
	MediaEntertainPref []*MediaEntertainPrefType `xml:"MediaEntertainPref,omitempty"`
	
	PetInfoPref []*PetInfoPrefType `xml:"PetInfoPref,omitempty"`
	
	MealPref []*MealPrefType `xml:"MealPref,omitempty"`
	
	LoyaltyPref []*LoyaltyPrefType `xml:"LoyaltyPref,omitempty"`
	
	SpecRequestPref []*SpecRequestPrefType `xml:"SpecRequestPref,omitempty"`
	
	RelatedTravelerPref []*RelatedTravelerPrefType `xml:"RelatedTravelerPref,omitempty"`
	
	TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	SmokingAllowed bool `xml:"SmokingAllowed,attr,omitempty"`
	
	PrimaryLangID *Language `xml:"PrimaryLangID,attr,omitempty"`
	
	AltLangID *Language `xml:"AltLangID,attr,omitempty"`
}

type NamePrefType struct {
	UniqueID *UniqueIDType `xml:"UniqueID,omitempty"`
	
	PersonName *PersonNameType `xml:"PersonName,omitempty"`
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

type PhonePrefType struct {
	Telephone struct {
		RPH string `xml:"RPH,attr,omitempty"`
	} `xml:"Telephone,omitempty"`
}

type AddressPrefType struct {
	Address *AddressInfoType `xml:"Address,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
}

type PaymentFormPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	RPH string `xml:"RPH,attr,omitempty"`
}

type InterestPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

type InsurancePrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	RPH string `xml:"RPH,attr,omitempty"`
}

type SeatingPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	SeatDirection string `xml:"SeatDirection,attr,omitempty"`
	
	SeatLocation string `xml:"SeatLocation,attr,omitempty"`
	
	SeatPosition string `xml:"SeatPosition,attr,omitempty"`
	
	SeatRow string `xml:"SeatRow,attr,omitempty"`
}

type TicketDistribPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	DistribType string `xml:"DistribType,attr,omitempty"`
	
	TicketTime time.Duration `xml:"TicketTime,attr,omitempty"`
}

type MediaEntertainPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

type PetInfoPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

type MealPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	MealType *MealType `xml:"MealType,attr,omitempty"`
	
	FavoriteFood string `xml:"FavoriteFood,attr,omitempty"`
	
	Beverage string `xml:"Beverage,attr,omitempty"`
}

type LoyaltyPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	RPH string `xml:"RPH,attr,omitempty"`
}

type SpecRequestPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

type RelatedTravelerPrefType struct {
	UniqueID *UniqueIDType `xml:"UniqueID,omitempty"`
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

type AirlinePrefType struct {
	LoyaltyPref []*LoyaltyPrefType `xml:"LoyaltyPref,omitempty"`
	
	VendorPref []*CompanyNamePrefType `xml:"VendorPref,omitempty"`
	
	PaymentFormPref []*PaymentFormPrefType `xml:"PaymentFormPref,omitempty"`
	
	AirportOriginPref []*AirportPrefType `xml:"AirportOriginPref,omitempty"`
	
	AirportRoutePref []*AirportPrefType `xml:"AirportRoutePref,omitempty"`
	
	FareRestrictPref []struct {
		FareRestriction string `xml:"FareRestriction,attr,omitempty"`
		
		Date string `xml:"Date,attr,omitempty"`
	} `xml:"FareRestrictPref,omitempty"`
	
	FlightTypePref []struct {
		FlightType *FlightTypeType `xml:"FlightType,attr,omitempty"`
		
		MaxConnections int `xml:"MaxConnections,attr,omitempty"`
		
		NonScheduledFltInfo string `xml:"NonScheduledFltInfo,attr,omitempty"`
		
		BackhaulIndicator bool `xml:"BackhaulIndicator,attr,omitempty"`
		
		GroundTransportIndicator bool `xml:"GroundTransportIndicator,attr,omitempty"`
		
		DirectAndNonStopOnlyInd bool `xml:"DirectAndNonStopOnlyInd,attr,omitempty"`
		
		NonStopsOnlyInd bool `xml:"NonStopsOnlyInd,attr,omitempty"`
		
		OnlineConnectionsOnlyInd bool `xml:"OnlineConnectionsOnlyInd,attr,omitempty"`
		
		RoutingType string `xml:"RoutingType,attr,omitempty"`
		
		ExcludeTrainInd bool `xml:"ExcludeTrainInd,attr,omitempty"`
	} `xml:"FlightTypePref,omitempty"`
	
	EquipPref []*EquipmentTypePref `xml:"EquipPref,omitempty"`
	
	CabinPref []struct {
		Cabin *CabinType `xml:"Cabin,attr,omitempty"`
	} `xml:"CabinPref,omitempty"`
	
	SeatPref []struct {
	} `xml:"SeatPref,omitempty"`
	
	TicketDistribPref []*TicketDistribPrefType `xml:"TicketDistribPref,omitempty"`
	
	MealPref []*MealPrefType `xml:"MealPref,omitempty"`
	
	SpecRequestPref []*SpecRequestPrefType `xml:"SpecRequestPref,omitempty"`
	
	SSR_Pref []struct {
		SSR_Code string `xml:"SSR_Code,attr,omitempty"`
	} `xml:"SSR_Pref,omitempty"`
	
	TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
	
	MediaEntertainPref []*MediaEntertainPrefType `xml:"MediaEntertainPref,omitempty"`
	
	PetInfoPref []*PetInfoPrefType `xml:"PetInfoPref,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	SmokingAllowed bool `xml:"SmokingAllowed,attr,omitempty"`
	
	PassengerTypeCode string `xml:"PassengerTypeCode,attr,omitempty"`
	
	AirTicketType *TicketType `xml:"AirTicketType,attr,omitempty"`
}

type EquipmentTypePref struct {
	*EquipmentType
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	WideBody bool `xml:"WideBody,attr,omitempty"`
}

type EquipmentType struct {
	Value string
	
	AirEquipType string `xml:"AirEquipType,attr,omitempty"`
	
	ChangeofGauge bool `xml:"ChangeofGauge,attr,omitempty"`
}

type HotelPrefType struct {
	LoyaltyPref []*LoyaltyPrefType `xml:"LoyaltyPref,omitempty"`
	
	PaymentFormPref []*PaymentFormPrefType `xml:"PaymentFormPref,omitempty"`
	
	HotelChainPref []*CompanyNamePrefType `xml:"HotelChainPref,omitempty"`
	
	PropertyNamePref []*PropertyNamePrefType `xml:"PropertyNamePref,omitempty"`
	
	PropertyLocationPref []*PropertyLocationPrefType `xml:"PropertyLocationPref,omitempty"`
	
	PropertyTypePref []*PropertyTypePrefType `xml:"PropertyTypePref,omitempty"`
	
	PropertyClassPref []*PropertyClassPrefType `xml:"PropertyClassPref,omitempty"`
	
	PropertyAmenityPref []*PropertyAmenityPrefType `xml:"PropertyAmenityPref,omitempty"`
	
	RoomLocationPref []*RoomLocationPrefType `xml:"RoomLocationPref,omitempty"`
	
	BedTypePref []*BedTypePrefType `xml:"BedTypePref,omitempty"`
	
	FoodSrvcPref []*FoodSrvcPrefType `xml:"FoodSrvcPref,omitempty"`
	
	MediaEntertainPref []*MediaEntertainPrefType `xml:"MediaEntertainPref,omitempty"`
	
	PetInfoPref []*PetInfoPrefType `xml:"PetInfoPref,omitempty"`
	
	MealPref []*MealPrefType `xml:"MealPref,omitempty"`
	
	RecreationSrvcPref []*RecreationSrvcPrefType `xml:"RecreationSrvcPref,omitempty"`
	
	BusinessSrvcPref []*BusinessSrvcPrefType `xml:"BusinessSrvcPref,omitempty"`
	
	PersonalSrvcPref []*PersonalSrvcPrefType `xml:"PersonalSrvcPref,omitempty"`
	
	SecurityFeaturePref []*SecurityFeaturePrefType `xml:"SecurityFeaturePref,omitempty"`
	
	PhysChallFeaturePref []*PhysChallFeaturePrefType `xml:"PhysChallFeaturePref,omitempty"`
	
	SpecRequestPref []*SpecRequestPrefType `xml:"SpecRequestPref,omitempty"`
	
	TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	SmokingAllowed bool `xml:"SmokingAllowed,attr,omitempty"`
	
	RatePlanCode string `xml:"RatePlanCode,attr,omitempty"`
	
	HotelGuestType string `xml:"HotelGuestType,attr,omitempty"`
}

type PropertyNamePrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	HotelCode string `xml:"HotelCode,attr,omitempty"`
}

type PropertyLocationPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	PropertyLocationType string `xml:"PropertyLocationType,attr,omitempty"`
}

type PropertyTypePrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	PropertyType string `xml:"PropertyType,attr,omitempty"`
}

type PropertyClassPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	PropertyClassType string `xml:"PropertyClassType,attr,omitempty"`
}

type PropertyAmenityPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	PropertyAmenityType string `xml:"PropertyAmenityType,attr,omitempty"`
}

type RoomLocationPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	RoomLocationType string `xml:"RoomLocationType,attr,omitempty"`
}

type BedTypePrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	BedType string `xml:"BedType,attr,omitempty"`
}

type FoodSrvcPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	FoodSrvcType string `xml:"FoodSrvcType,attr,omitempty"`
}

type RecreationSrvcPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	RecreationSrvcType string `xml:"RecreationSrvcType,attr,omitempty"`
}

type BusinessSrvcPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	BusinessSrvcType string `xml:"BusinessSrvcType,attr,omitempty"`
}

type PersonalSrvcPrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

type SecurityFeaturePrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
}

type PhysChallFeaturePrefType struct {
	Value string
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	PhysChallFeatureType string `xml:"PhysChallFeatureType,attr,omitempty"`
}

type OtherSrvcPrefType struct {
	OtherSrvcName string `xml:"OtherSrvcName,omitempty"`
	
	VendorPref []*CompanyNamePrefType `xml:"VendorPref,omitempty"`
	
	LoyaltyPref []*LoyaltyPrefType `xml:"LoyaltyPref,omitempty"`
	
	PaymentFormPref []*PaymentFormPrefType `xml:"PaymentFormPref,omitempty"`
	
	SpecRequestPref []*SpecRequestPrefType `xml:"SpecRequestPref,omitempty"`
	
	TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
	
	PreferLevel *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	TravelPurpose string `xml:"TravelPurpose,attr,omitempty"`
}

type CompanyInfoType struct {
	CompanyName []*CompanyNameType `xml:"CompanyName,omitempty"`
	
	AddressInfo []*AddressInfoType `xml:"AddressInfo,omitempty"`
	
	TelephoneInfo []struct {
		RPH string `xml:"RPH,attr,omitempty"`
	} `xml:"TelephoneInfo,omitempty"`
	
	Email []*EmailType `xml:"Email,omitempty"`
	
	URL []*URL_Type `xml:"URL,omitempty"`
	
	BusinessLocale []*AddressType `xml:"BusinessLocale,omitempty"`
	
	PaymentForm []*PaymentFormType `xml:"PaymentForm,omitempty"`
	
	ContactPerson []*ContactPersonType `xml:"ContactPerson,omitempty"`
	
	TravelArranger []*TravelArrangerType `xml:"TravelArranger,omitempty"`
	
	LoyaltyProgram []*LoyaltyProgramType `xml:"LoyaltyProgram,omitempty"`
}

type LoyaltyProgramType struct {
	Value string
	
	ProgramCode string `xml:"ProgramCode,attr,omitempty"`
	
	SingleVendorInd string `xml:"SingleVendorInd,attr,omitempty"`
	
	LoyaltyLevel string `xml:"LoyaltyLevel,attr,omitempty"`
	
	RPH string `xml:"RPH,attr,omitempty"`
}

type OrganizationType struct {
	OrgMemberName struct {
		*PersonNameType
		
		Level string `xml:"Level,attr,omitempty"`
		
		Title string `xml:"Title,attr,omitempty"`
	} `xml:"OrgMemberName,omitempty"`
	
	OrgName *CompanyNameType `xml:"OrgName,omitempty"`
	
	RelatedOrgName []*CompanyNameType `xml:"RelatedOrgName,omitempty"`
	
	TravelArranger []*TravelArrangerType `xml:"TravelArranger,omitempty"`
	
	DefaultInd bool `xml:"DefaultInd,attr,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	EffectiveDate time.Time `xml:"EffectiveDate,attr,omitempty"`
	
	ExpireDate time.Time `xml:"ExpireDate,attr,omitempty"`
	
	OfficeType *OfficeLocationType `xml:"OfficeType,attr,omitempty"`
}

type EmployerType struct {
	CompanyName *CompanyNameType `xml:"CompanyName,omitempty"`
	
	RelatedEmployer []*CompanyNameType `xml:"RelatedEmployer,omitempty"`
	
	EmployeeInfo []*EmployeeInfoType `xml:"EmployeeInfo,omitempty"`
	
	InternalRefNmbr []*FreeTextType `xml:"InternalRefNmbr,omitempty"`
	
	TravelArranger []*TravelArrangerType `xml:"TravelArranger,omitempty"`
	
	LoyaltyProgram []*LoyaltyProgramType `xml:"LoyaltyProgram,omitempty"`
	
	DefaultInd bool `xml:"DefaultInd,attr,omitempty"`
	
	OfficeType *OfficeLocationType `xml:"OfficeType,attr,omitempty"`
	
	EffectiveDate time.Time `xml:"EffectiveDate,attr,omitempty"`
	
	ExpireDate time.Time `xml:"ExpireDate,attr,omitempty"`
}

type TravelClubType struct {
	TravelClubName *CompanyNameType `xml:"TravelClubName,omitempty"`
	
	ClubMemberName struct {
		*PersonNameType
		
		ID string `xml:"ID,attr,omitempty"`
	} `xml:"ClubMemberName,omitempty"`
	
	ShareSynchInd string `xml:"ShareSynchInd,attr,omitempty"`
	
	ShareMarketInd string `xml:"ShareMarketInd,attr,omitempty"`
	
	EffectiveDate time.Time `xml:"EffectiveDate,attr,omitempty"`
	
	ExpireDate time.Time `xml:"ExpireDate,attr,omitempty"`
}

// AgreementsType ..
type AgreementsType struct {
	ShareSynchInd      string                    `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd     string                    `xml:"ShareMarketInd,attr,omitempty"`
	TPAExtensions      *TPAExtensionsType        `xml:"TPA_Extensions,omitempty"`
	Certification      []*CertificationType      `xml:"Certification,omitempty"`
	CommissionInfo     []*CommissionInfoType     `xml:"CommissionInfo,omitempty"`
	AllianceConsortium []*AllianceConsortiumType `xml:"AllianceConsortium,omitempty"`
}

// CertificationType ...
type CertificationType struct {
	Value           *FreeTextType
	ID              string    `xml:"ID,attr,omitempty"`
	SingleVendorInd string    `xml:"SingleVendorInd,attr,omitempty"`
	EffectiveDate   time.Time `xml:"EffectiveDate,attr,omitempty"`
	ExpireDate      time.Time `xml:"ExpireDate,attr,omitempty"`
}

// AllianceConsortiumType ...
type AllianceConsortiumType struct {
	ID             string            `xml:"ID,attr,omitempty"`
	EffectiveDate  time.Time         `xml:"EffectiveDate,attr,omitempty"`
	ExpireDate     time.Time         `xml:"ExpireDate,attr,omitempty"`
	AllianceMember []*AllianceMember `xml:"AllianceMember,omitempty"`
}

// AllianceMember ...
type AllianceMember struct {
	Value      *CompanyNameType
	MemberCode string `xml:"MemberCode,attr,omitempty"`
}

// CommissionInfoType ...
type CommissionInfoType struct {
	Value              *FreeTextType
	ShareSynchInd      string  `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd     string  `xml:"ShareMarketInd,attr,omitempty"`
	CommissionPlanCode string  `xml:"CommissionPlanCode,attr,omitempty"`
	Amount             float64 `xml:"Amount,attr,omitempty"`
}

// ArrayOfBookingRulesTypeBookingRule ...
type ArrayOfBookingRulesTypeBookingRule struct {
	BookingRule []*BookingRule `xml:"BookingRule,omitempty"`
}

// BookingRule ...
type BookingRule struct {
	MinTotalOccupancy        uint                                                   `xml:"MinTotalOccupancy,attr,omitempty"`
	MaxTotalOccupancy        uint                                                   `xml:"MaxTotalOccupancy,attr,omitempty"`
	MaxContiguousBookings    uint                                                   `xml:"MaxContiguousBookings,attr,omitempty"`
	GenerallyBookable        bool                                                   `xml:"GenerallyBookable,attr,omitempty"`
	PriceViewable            bool                                                   `xml:"PriceViewable,attr,omitempty"`
	QualifiedRateYN          bool                                                   `xml:"QualifiedRateYN,attr,omitempty"`
	AddressRequired          bool                                                   `xml:"AddressRequired,attr,omitempty"`
	AbsoluteDropTime         string                                                 `xml:"AbsoluteDropTime,attr,omitempty"`
	AbsoluteCutoff           string                                                 `xml:"AbsoluteCutoff,attr,omitempty"`
	URI                      string                                                 `xml:"URI,attr,omitempty"`
	Start                    string                                                 `xml:"Start,attr,omitempty"`
	Duration                 string                                                 `xml:"Duration,attr,omitempty"`
	End                      string                                                 `xml:"End,attr,omitempty"`
	OffsetCalculationMode    string                                                 `xml:"OffsetCalculationMode,attr,omitempty"`
	MaxAdvancedBookingOffset time.Duration                                          `xml:"MaxAdvancedBookingOffset,attr,omitempty"`
	MinAdvancedBookingOffset time.Duration                                          `xml:"MinAdvancedBookingOffset,attr,omitempty"`
	DepositWaiverOffset      time.Duration                                          `xml:"DepositWaiverOffset,attr,omitempty"`
	ForceGuaranteeOffset     time.Duration                                          `xml:"ForceGuaranteeOffset,attr,omitempty"`
	OffsetDuration           time.Duration                                          `xml:"OffsetDuration,attr,omitempty"`
	UniqueID                 *UniqueIDType                                          `xml:"UniqueID,omitempty"`
	Description              []*ParagraphType                                       `xml:"Description,omitempty"`
	CheckoutCharge           []*CheckoutCharge                                      `xml:"CheckoutCharge,omitempty"`
	LengthsOfStay            *LengthsOfStayType                                     `xml:"LengthsOfStay,omitempty"`
	RestrictionStatus        *RestrictionStatus                                     `xml:"RestrictionStatus,omitempty"`
	CancelPenalties          *CancelPenaltiesType                                   `xml:"CancelPenalties,omitempty"`
	DOWRestrictions          *DOWRestrictionsType                                   `xml:"DOW_Restrictions,omitempty"`
	ViewerShips              *ArrayOfViewerShipsTypeViewership                      `xml:"Viewerships,omitempty"`
	RequiredPaymts           *ArrayOfRequiredPaymentsTypeGuaranteePayment           `xml:"RequiredPaymts,omitempty"`
	AdditionalRules          *ArrayOfBookingRulesTypeBookingRuleAdditionalRule      `xml:"AddtionalRules,omitempty"`
	AcceptableGuarantees     *ArrayOfBookingRulesTypeBookingRuleAcceptableGuarantee `xml:"AcceptableGuarantees,omitempty"`
}

// CheckoutCharge ...
type CheckoutCharge struct {
	Amount           float64 `xml:"Amount,attr,omitempty"`
	Percent          float64 `xml:"Percent,attr,omitempty"`
	Type             string  `xml:"Type,attr,omitempty"`
	CodeDetail       string  `xml:"CodeDetail,attr,omitempty"`
	NmbrOfNights     int     `xml:"NmbrOfNights,attr,omitempty"`
	ExistsCode       string  `xml:"ExistsCode,attr,omitempty"`
	BalanceOfStayInd bool    `xml:"BalanceOfStayInd,attr,omitempty"`
}

// RestrictionStatus ...
type RestrictionStatus struct {
	Restriction              string                  `xml:"Restriction,attr,omitempty"`
	SellThroughOpenIndicator bool                    `xml:"SellThroughOpenIndicator,attr,omitempty"`
	Status                   *AvailabilityStatusType `xml:"Status,attr,omitempty"`
}

// ArrayOfBookingRulesTypeBookingRuleAcceptableGuarantee ...
type ArrayOfBookingRulesTypeBookingRuleAcceptableGuarantee struct {
	AcceptableGuarantee []*AcceptableGuarantee `xml:"AcceptableGuarantee,omitempty"`
}

// AcceptableGuarantee ...
type AcceptableGuarantee struct {
	*GuaranteeType
	DecimalPlaces           uint   `xml:"DecimalPlaces,attr,omitempty"`
	GuaranteePolicyType     string `xml:"GuaranteePolicyType,attr,omitempty"`
	PaymentType             string `xml:"PaymentType,attr,omitempty"`
	CurrencyCode            string `xml:"CurrencyCode,attr,omitempty"`
	UnacceptablePaymentType string `xml:"UnacceptablePaymentType,attr,omitempty"`
}

// LengthsOfStayType ...
type LengthsOfStayType struct {
	LengthOfStay       []*LengthOfStay `xml:"LengthOfStay,omitempty"`
	ArrivalDateBased   bool            `xml:"ArrivalDateBased,attr,omitempty"`
	FixedPatternLength int32           `xml:"FixedPatternLength,attr,omitempty"`
}

// LengthOfStay ...
type LengthOfStay struct {
	OpenStatusIndicator bool          `xml:"OpenStatusIndicator,attr,omitempty"`
	Time                int32         `xml:"Time,attr,omitempty"`
	MinMaxMessageType   string        `xml:"MinMaxMessageType,attr,omitempty"`
	LOSPattern          *LOSPattern   `xml:"LOS_Pattern,omitempty"`
	TimeUnit            *TimeUnitType `xml:"TimeUnit,attr,omitempty"`
}

// LOSPattern ...
type LOSPattern struct {
	FullPatternLOS string `xml:"FullPatternLOS,attr,omitempty"`
}

// ArrayOfViewerShipsTypeViewership ...
type ArrayOfViewerShipsTypeViewership struct {
	ViewerShip []*ViewerShip `xml:"Viewership,omitempty"`
}

// ViewerShip ...
type ViewerShip struct {
	ViewershipRPH       int8                                             `xml:"ViewershipRPH,attr,omitempty"`
	ViewOnly            bool                                             `xml:"ViewOnly,attr,omitempty"`
	SystemCodes         *SystemCodes                                     `xml:"SystemCodes,omitempty"`
	LocationCodes       *LocationCodes                                   `xml:"LocationCodes,omitempty"`
	BookingChannelCodes *LocationCodes                                   `xml:"BookingChannelCodes,omitempty"`
	ViewershipCodes     *ViewerShipCodes                                 `xml:"ViewershipCodes,omitempty"`
	Profiles            *ArrayOfProfileType                              `xml:"Profiles,omitempty"`
	ProfileRefs         *ArrayOfUniqueIDType                             `xml:"ProfileRefs,omitempty"`
	ProfileTypes        *ArrayOfViewerShipsTypeViewershipProfileType     `xml:"ProfileTypes,omitempty"`
	DistributorTypes    *ArrayOfViewershipsTypeViewershipDistributorType `xml:"DistributorTypes,omitempty"`
}

// BookingChannelCodes ...
type BookingChannelCodes struct {
	ChannelCodesInclusive bool         `xml:"ChannelCodesInclusive,attr,omitempty"`
	BookingChannelCode    []SystemCode `xml:"BookingChannelCode,omitempty"`
}

// BookingChannelCode ...
type BookingChannelCode struct {
	Value                      string
	Sort                       uint `xml:"Sort,attr,omitempty"`
	RestrictedDisplayIndicator bool `xml:"RestrictedDisplayIndicator,attr,omitempty"`
}

// LocationCodes ...
type LocationCodes struct {
	LocationCode           []*LocationCode `xml:"LocationCode,omitempty"`
	LocationCodesInclusive bool            `xml:"LocationCodesInclusive,attr,omitempty"`
}

// LocationCode ...
type LocationCode struct {
	CityCode          string `xml:"CityCode,attr,omitempty"`
	StateProvinceCode string `xml:"StateProvinceCode,attr,omitempty"`
	CountryCode       string `xml:"CountryCode,attr,omitempty"`
}

// SystemCodes ...
type SystemCodes struct {
	SystemCodesInclusive bool          `xml:"SystemCodesInclusive,attr,omitempty"`
	SystemCode           []*SystemCode `xml:"SystemCode,omitempty"`
}

// SystemCode ...
type SystemCode struct {
	Sort                       uint `xml:"Sort,attr,omitempty"`
	RestrictedDisplayIndicator bool `xml:"RestrictedDisplayIndicator,attr,omitempty"`
	Value                      string
}

// ViewerShipCodes ...
type ViewerShipCodes struct {
	ViewerShipCode string `xml:"ViewershipCode,omitempty"`
}

// ArrayOfViewerShipsTypeViewershipProfileType ...
type ArrayOfViewerShipsTypeViewershipProfileType struct {
	ProfileType []*ProfileType `xml:"ProfileType,omitempty"`
}

// ViewerProfileType ...
type ViewerProfileType struct {
	ProfileType string `xml:"ProfileType,attr,omitempty"`
}

// ArrayOfUniqueIDType ...
type ArrayOfUniqueIDType struct {
	ProfileRef []*UniqueIDType `xml:"ProfileRef,omitempty"`
}

// ArrayOfProfileType ...
type ArrayOfProfileType struct {
	Profile []*ProfileType `xml:"Profile,omitempty"`
}

// ArrayOfViewershipsTypeViewershipDistributorType ...
type ArrayOfViewershipsTypeViewershipDistributorType struct {
	DistributorType []*DistributorType `xml:"DistributorType,omitempty"`
}

// DistributorType ...
type DistributorType struct {
	Value               string
	DistributorCode     string `xml:"DistributorCode,attr,omitempty"`
	DistributorTypeCode string `xml:"DistributorTypeCode,attr,omitempty"`
}

// ArrayOfBookingRulesTypeBookingRuleAdditionalRule ...
type ArrayOfBookingRulesTypeBookingRuleAdditionalRule struct {
	AdditionalRule []*AdditionalRule `xml:"AdditionalRule,omitempty"`
}

// AdditionalRule ...
type AdditionalRule struct {
	AdditionalRule string `xml:"AdditionalRule,attr,omitempty"`
}

// ArrayOfResGuestsTypeResGuest ...
type ArrayOfResGuestsTypeResGuest struct {
	ResGuest []*ResGuest `xml:"ResGuest,omitempty"`
}

// ResGuest ...
type ResGuest struct {
	ResGuestRPH        string                                   `xml:"ResGuestRPH,attr,omitempty"`
	AgeQualifyingCode  string                                   `xml:"AgeQualifyingCode,attr,omitempty"`
	ArrivalTime        time.Time                                `xml:"ArrivalTime,attr,omitempty"`
	DepartureTime      time.Time                                `xml:"DepartureTime,attr,omitempty"`
	GroupEventCode     string                                   `xml:"GroupEventCode,attr,omitempty"`
	VIP                bool                                     `xml:"VIP,attr,omitempty"`
	PrimaryIndicator   bool                                     `xml:"PrimaryIndicator,attr,omitempty"`
	GuestCounts        *GuestCountType                          `xml:"GuestCounts,omitempty"`
	InHouseTimeSpan    *DateTimeSpanType                        `xml:"InHouseTimeSpan,omitempty"`
	ArrivalTransport   *TransportInfoType                       `xml:"ArrivalTransport,omitempty"`
	DepartureTransport *TransportInfoType                       `xml:"DepartureTransport,omitempty"`
	TPAExtensions      *TPAExtensionsType                       `xml:"TPA_Extensions,omitempty"`
	Comments           *ArrayOfCommentTypeComment               `xml:"Comments,omitempty"`
	Profiles           *ArrayOfProfilesTypeProfileInfo          `xml:"Profiles,omitempty"`
	ServiceRPHs        *ArrayOfServiceRPHsTypeServiceRPH        `xml:"ServiceRPHs,omitempty"`
	ProfileRPHs        *ArrayOfResGuestsTypeResGuestProfileRPH  `xml:"ProfileRPHs,omitempty"`
	SpecialRequests    *ArrayOfSpecialRequestTypeSpecialRequest `xml:"SpecialRequests,omitempty"`
}

// ArrayOfResGuestsTypeResGuestProfileRPH ...
type ArrayOfResGuestsTypeResGuestProfileRPH struct {
	ProfileRPH []*ProfileRPH `xml:"ProfileRPH,omitempty"`
}

// ProfileRPH ...
type ProfileRPH []struct {
	RPH string `xml:"RPH,attr,omitempty"`
}

// TransportInfoType ...
type TransportInfoType struct {
	TransportInfo *TransportInfo `xml:"TransportInfo,omitempty"`
}

// TransportInfo ...
type TransportInfo struct {
	Type         string    `xml:"Type,attr,omitempty"`
	ID           string    `xml:"ID,attr,omitempty"`
	LocationCode string    `xml:"LocationCode,attr,omitempty"`
	Time         time.Time `xml:"Time,attr,omitempty"`
}

// WrittenConfInstType ...
type WrittenConfInstType struct {
	ConfirmInd       bool           `xml:"ConfirmInd,attr,omitempty"`
	LanguageID       string         `xml:"LanguageID,attr,omitempty"`
	AddresseeName    string         `xml:"AddresseeName,attr,omitempty"`
	Address          string         `xml:"Address,attr,omitempty"`
	Telephone        string         `xml:"Telephone,attr,omitempty"`
	Email            *EmailType     `xml:"Email,omitempty"`
	SupplementalData *ParagraphType `xml:"SupplementalData,omitempty"`
}

// AirItineraryPricingInfoType ...
type AirItineraryPricingInfoType struct {
	ITIndex           int32                                       `xml:"ITIndex,attr,omitempty"`
	MinCCOBFee        float64                                     `xml:"MinCCOBFee,omitempty"`
	MaxCCOBFee        float64                                     `xml:"MaxCCOBFee,omitempty"`
	CCOBFee           float64                                     `xml:"CCOBFee,omitempty"`
	ItinTotalFare     *FareType                                   `xml:"ItinTotalFare,omitempty"`
	ServiceFees       *ServiceFeesType                            `xml:"ServiceFees,omitempty"`
	PricingSource     *PricingSourceType                          `xml:"PricingSource,attr,omitempty"`
	PTCFareBreakdowns *ArrayOfPTCFareBreakdownType                `xml:"PTC_FareBreakdowns,omitempty"`
	FareInfos         *ArrayOfAirItineraryPricingInfoTypeFareInfo `xml:"FareInfos,omitempty"`
}

// FareType ...
type FareType struct {
	NegotiatedFare                 bool                  `xml:"NegotiatedFare,attr,omitempty"`
	NegotiatedFareCode             string                `xml:"NegotiatedFareCode,attr,omitempty"`
	TicketDesignatorCode           string                `xml:"TicketDesignatorCode,attr,omitempty"`
	TotalNbrTrips                  int32                 `xml:"TotalNbrTrips,attr,omitempty"`
	TotalNbrPTC                    int32                 `xml:"TotalNbrPTC,attr,omitempty"`
	BaseFare                       *BaseFare             `xml:"BaseFare,omitempty"`
	MarkupFare                     *BaseFare             `xml:"MarkupFare,omitempty"`
	EquivFare                      *BaseFare             `xml:"EquivFare,omitempty"`
	TotalFare                      *BaseFare             `xml:"TotalFare,omitempty"`
	TotalAmountInTicketingCurrency *BaseFare             `xml:"TotalAmountInTicketingCurrency,omitempty"`
	FareConstruction               *FareConstruction     `xml:"FareConstruction,omitempty"`
	Taxes                          *ArrayOfAirTaxType    `xml:"Taxes,omitempty"`
	Fees                           *ArrayOfAirFeeType    `xml:"Fees,omitempty"`
	TPAExtensions                  *TPAExtensionsType    `xml:"TPA_Extensions,omitempty"`
	UnstructuredFareCalc           *UnstructuredFareCalc `xml:"UnstructuredFareCalc,omitempty"`
}

// UnstructuredFareCalc ...
type UnstructuredFareCalc struct {
	Value        string
	FareCalcMode string `xml:"FareCalcMode,attr,omitempty"`
}

// FareConstruction ...
type FareConstruction struct {
	Language               *Language `xml:"Language,attr,omitempty"`
	OriginCityCode         string    `xml:"OriginCityCode,attr,omitempty"`
	OriginCodeContext      string    `xml:"OriginCodeContext,attr,omitempty"`
	FormattedIndicator     bool      `xml:"FormattedIndicator,attr,omitempty"`
	DestinationCityCode    string    `xml:"DestinationCityCode,attr,omitempty"`
	DestinationCodeContext string    `xml:"DestinationCodeContext,attr,omitempty"`
}

// BaseFare ...
type BaseFare struct {
	Currency string    `xml:"Currency,attr,omitempty"`
	Amount   float64   `xml:"Amount,attr,omitempty"`
	Rate     float64   `xml:"Rate,attr,omitempty"`
	Date     time.Time `xml:"Date,attr,omitempty"`
}

// ArrayOfAirTaxType ...
type ArrayOfAirTaxType struct {
	Tax []*AirTaxType `xml:"Tax,omitempty"`
}

// AirTaxType ...
type AirTaxType struct {
	Value      string
	TaxCode    string  `xml:"TaxCode,attr,omitempty"`
	Amount     float64 `xml:"Amount,attr,omitempty"`
	TaxCountry string  `xml:"TaxCountry,attr,omitempty"`
	TaxName    string  `xml:"TaxName,attr,omitempty"`
}

// ArrayOfAirFeeType ...
type ArrayOfAirFeeType struct {
	Fee []*AirFeeType `xml:"Fee,omitempty"`
}

// AirFeeType ...
type AirFeeType struct {
	Value   string
	FeeCode string  `xml:"FeeCode,attr,omitempty"`
	Amount  float64 `xml:"Amount,attr,omitempty"`
}

// ArrayOfPTCFareBreakdownType ...
type ArrayOfPTCFareBreakdownType struct {
	PTCFareBreakdown []*PTCFareBreakdownType `xml:"PTC_FareBreakdown,omitempty"`
}

// PTCFareBreakdownType ...
type PTCFareBreakdownType struct {
	FareBasisCodes        []string                                     `xml:"FareBasisCodes,omitempty"`
	PassengerFare         *FareType                                    `xml:"PassengerFare,omitempty"`
	Endorsements          *Endorsements                                `xml:"Endorsements,omitempty"`
	PricingUnit           []*PricingUnit                               `xml:"PricingUnit,omitempty"`
	PricingSource         *PricingSourceType                           `xml:"PricingSource,attr,omitempty"`
	FareInfo              []PassengerFareInfo                          `xml:"FareInfo,omitempty"`
	TravelerRefNumber     []*TravelerRefNumber                         `xml:"TravelerRefNumber,omitempty"`
	PassengerTypeQuantity *PassengerTypeQuantityType                   `xml:"PassengerTypeQuantity,omitempty"`
	TicketDesignators     *ArrayOfPTCFareBreakdownTypeTicketDesignator `xml:"TicketDesignators,omitempty"`
}

// Endorsements ....
type Endorsements struct {
	Endorsement            []*Endrosment `xml:"Endorsement,omitempty"`
	NonRefundableIndicator bool          `xml:"NonRefundableIndicator,attr,omitempty"`
	NonEndrosableIndicator bool          `xml:"NonEndorsableIndicator,attr,omitempty"`
}

// Endrosment ...
type Endrosment struct {
	*FreeTextType
}

// PassengerFareInfo ...
type PassengerFareInfo []struct {
	*FareInfoType
	PassengerFare *FareType `xml:"PassengerFare,omitempty"`
}

// PricingUnit ...
type PricingUnit struct {
	FareComponent []*FareComponent `xml:"FareComponent,omitempty"`
	UnitNumber    int32            `xml:"UnitNumber,attr,omitempty"`
}

// FareComponent ...
type FareComponent struct {
	FlightLeg []*FlightLeg `xml:"FlightLeg,omitempty"`
	
	Number int32 `xml:"Number,attr,omitempty"`
	
	Amount float64 `xml:"Amount,attr,omitempty"`
}

// FlightLeg ...
type FlightLeg struct {
	*BookFlightSegmentType
	SurchargeInd          bool    `xml:"SurchargeInd,attr,omitempty"`
	FareBasisCode         string  `xml:"FareBasisCode,attr,omitempty"`
	UnitOfMeasureQuantity float64 `xml:"UnitOfMeasureQuantity,attr,omitempty"`
	UnitOfMeasure         string  `xml:"UnitOfMeasure,attr,omitempty"`
	UnitOfMeasureCode     string  `xml:"UnitOfMeasureCode,attr,omitempty"`
}

// ArrayOfPTCFareBreakdownTypeTicketDesignator ...
type ArrayOfPTCFareBreakdownTypeTicketDesignator struct {
	TicketDesignator []*TicketDesignator `xml:"TicketDesignator,omitempty"`
}

// TicketDesignator ...
type TicketDesignator struct {
	FlightRefRPH              string `xml:"FlightRefRPH,attr,omitempty"`
	TicketDesignatorCode      string `xml:"TicketDesignatorCode,attr,omitempty"`
	TicketDesignatorExtension string `xml:"TicketDesignatorExtension,attr,omitempty"`
}

// FareInfoType ...
type FareInfoType struct {
	NegotiatedFare     bool                `xml:"NegotiatedFare,attr,omitempty"`
	NegotiatedFareCode string              `xml:"NegotiatedFareCode,attr,omitempty"`
	CurrencyCode       string              `xml:"CurrencyCode,attr,omitempty"`
	TariffNumber       string              `xml:"TariffNumber,attr,omitempty"`
	RuleNumber         string              `xml:"RuleNumber,attr,omitempty"`
	RoutingNumber      int32               `xml:"RoutingNumber,attr,omitempty"`
	Date               []*Date             `xml:"Date,omitempty"`
	DepartureDate      time.Time           `xml:"DepartureDate,omitempty"`
	RuleInfo           *RuleInfo           `xml:"RuleInfo,omitempty"`
	City               []*Airport          `xml:"City,omitempty"`
	Airport            []*Airport          `xml:"Airport,omitempty"`
	DepartureAirport   *LocationType       `xml:"DepartureAirport,omitempty"`
	ArrivalAirport     *LocationType       `xml:"ArrivalAirport,omitempty"`
	FilingAirline      *CompanyNameType    `xml:"FilingAirline,omitempty"`
	FareReference      []*FareReference    `xml:"FareReference,omitempty"`
	DiscountPricing    *DiscountPricing    `xml:"DiscountPricing,omitempty"`
	MarketingAirline   []*CompanyNameType  `xml:"MarketingAirline,omitempty"`
	FareInfo           []*FareInfoTypeFare `xml:"FareInfo,omitempty"`
}

// Date ...
type Date struct {
	Date string `xml:"Date,attr,omitempty"`
	Type string `xml:"Type,attr,omitempty"`
}

// FareInfoTypeFare ...
type FareInfoTypeFare struct {
	Fare                    *Fare                `xml:"Fare,omitempty"`
	MaximumPermittedMileage int32                `xml:"MaximumPermittedMileage,attr,omitempty"`
	PTC                     []*PTC               `xml:"PTC,omitempty"`
	FareBasisCode           string               `xml:"FareBasisCode,attr,omitempty"`
	FareType                string               `xml:"FareType,attr,omitempty"`
	Date                    []*Date              `xml:"Date,omitempty"`
	TripType                string               `xml:"TripType,attr,omitempty"`
	FareStatus              *FareStatusType      `xml:"FareStatus,attr,omitempty"`
	GlobalIndicatorCode     *GlobalIndicatorType `xml:"GlobalIndicatorCode,attr,omitempty"`
}

// PTC ...
type PTC struct {
	PassengerTypeCode string `xml:"PassengerTypeCode,attr,omitempty"`
}

// Fare ...
type Fare struct {
	FareDescription string  `xml:"FareDescription,attr,omitempty"`
	BaseAmount      float64 `xml:"BaseAmount,attr,omitempty"`
	BaseNUCAmount   float64 `xml:"BaseNUC_Amount,attr,omitempty"`
	TaxAmount       float64 `xml:"TaxAmount,attr,omitempty"`
	TotalFare       float64 `xml:"TotalFare,attr,omitempty"`
}

// DiscountPricing ...
type DiscountPricing struct {
	Purpose              string `xml:"Purpose,attr,omitempty"`
	TicketDesignatorCode string `xml:"TicketDesignatorCode,attr,omitempty"`
	Discount             string `xml:"Discount,attr,omitempty"`
	Type                 string `xml:"Type,attr,omitempty"`
	Usage                string `xml:"Usage,attr,omitempty"`
	Text                 string `xml:"Text,attr,omitempty"`
}

// RuleInfo ...
type RuleInfo struct {
	*RuleInfoType
	MoneySaverInd bool   `xml:"MoneySaverInd,attr,omitempty"`
	TripType      string `xml:"TripType,attr,omitempty"`
}

// FareReference ...
type FareReference struct {
	Value                string
	ResBookDesigCode     string `xml:"ResBookDesigCode,attr,omitempty"`
	TicketDesignatorCode string `xml:"TicketDesignatorCode,attr,omitempty"`
	AccountCode          string `xml:"AccountCode,attr,omitempty"`
}

// RuleInfoType ...
type RuleInfoType struct {
	ChargesRules      *ChargesRules         `xml:"ChargesRules,omitempty"`
	ResTicketingRules *ResTicketingRules    `xml:"ResTicketingRules,omitempty"`
	LengthOfStayRules *StayRestrictionsType `xml:"LengthOfStayRules,omitempty"`
}

// ChargesRules ...
type ChargesRules struct {
	VoluntaryChanges *VoluntaryChangesType `xml:"VoluntaryChanges,omitempty"`
	VoluntaryRefunds *VoluntaryChangesType `xml:"VoluntaryRefunds,omitempty"`
}

// ResTicketingRules ...
type ResTicketingRules struct {
	AdvResTicketing *AdvResTicketing `xml:"AdvResTicketing,omitempty"`
}

// AdvResTicketing ...
type AdvResTicketing struct {
	*AdvResTicketingType
	FirstTicketDate string `xml:"FirstTicketDate,attr,omitempty"`
	LastTicketDate  string `xml:"LastTicketDate,attr,omitempty"`
}

// AdvResTicketingType ...
type AdvResTicketingType struct {
	AdvResInd              bool            `xml:"AdvResInd,attr,omitempty"`
	AdvTicketingInd        bool            `xml:"AdvTicketingInd,attr,omitempty"`
	RequestedTicketingDate string          `xml:"RequestedTicketingDate,attr,omitempty"`
	AdvReservation         *AdvReservation `xml:"AdvReservation,omitempty"`
	AdvTicketing           *AdvTicketing   `xml:"AdvTicketing,omitempty"`
}

// AdvTicketing ...
type AdvTicketing struct {
	FromResTimeOfDay    string        `xml:"FromResTimeOfDay,attr,omitempty"`
	FromResPeriod       string        `xml:"FromResPeriod,attr,omitempty"`
	FromDepartTimeOfDay string        `xml:"FromDepartTimeOfDay,attr,omitempty"`
	FromDepartPeriod    string        `xml:"FromDepartPeriod,attr,omitempty"`
	FromResUnit         *StayUnitType `xml:"FromResUnit,attr,omitempty"`
	FromDepartUnit      *StayUnitType `xml:"FromDepartUnit,attr,omitempty"`
}

// AdvReservation ...
type AdvReservation struct {
	LatestTimeOfDay string        `xml:"LatestTimeOfDay,attr,omitempty"`
	LatestPeriod    string        `xml:"LatestPeriod,attr,omitempty"`
	LatestUnit      *StayUnitType `xml:"LatestUnit,attr,omitempty"`
}

// StayRestrictionsType ...
type StayRestrictionsType struct {
	StayRestrictionsInd bool      `xml:"StayRestrictionsInd,attr,omitempty"`
	MinimumStay         *StayType `xml:"MinimumStay,omitempty"`
	MaximumStay         *StayType `xml:"MaximumStay,omitempty"`
}

// StayType ...
type StayType struct {
	ComplicatedRulesInd bool          `xml:"ComplicatedRulesInd,attr,omitempty"`
	MinStay             int32         `xml:"MinStay,attr,omitempty"`
	ReturnTimeOfDay     string        `xml:"ReturnTimeOfDay,attr,omitempty"`
	MinStayDate         string        `xml:"MinStayDate,attr,omitempty"`
	StayUnit            *StayUnitType `xml:"StayUnit,attr,omitempty"`
}

// VoluntaryChangesType ...
type VoluntaryChangesType struct {
	Penalty struct {
		PenaltyType string `xml:"PenaltyType,attr,omitempty"`
		
		DepartureStatus string `xml:"DepartureStatus,attr,omitempty"`
		
		Amount float64 `xml:"Amount,attr,omitempty"`
		
		Percent float64 `xml:"Percent,attr,omitempty"`
	} `xml:"Penalty,omitempty"`
	VolChangeInd bool `xml:"VolChangeInd,attr,omitempty"`
}

// Penalty ...
type Penalty struct {
	PenaltyType     string  `xml:"PenaltyType,attr,omitempty"`
	DepartureStatus string  `xml:"DepartureStatus,attr,omitempty"`
	Amount          float64 `xml:"Amount,attr,omitempty"`
	Percent         float64 `xml:"Percent,attr,omitempty"`
}

// FareRuleResponseInfoTypeFareRuleInfo ...
type FareRuleResponseInfoTypeFareRuleInfo struct {
	// *FareInfoType
	NegotiatedFare     bool                `xml:"NegotiatedFare,attr,omitempty"`
	NegotiatedFareCode string              `xml:"NegotiatedFareCode,attr,omitempty"`
	CurrencyCode       string              `xml:"CurrencyCode,attr,omitempty"`
	TariffNumber       string              `xml:"TariffNumber,attr,omitempty"`
	RuleNumber         string              `xml:"RuleNumber,attr,omitempty"`
	RoutingNumber      int32               `xml:"RoutingNumber,attr,omitempty"`
	Date               []*Date             `xml:"Date,omitempty"`
	DepartureDate      time.Time           `xml:"DepartureDate,omitempty"`
	RuleInfo           *RuleInfo           `xml:"RuleInfo,omitempty"`
	City               []*Airport          `xml:"City,omitempty"`
	Airport            []*Airport          `xml:"Airport,omitempty"`
	DepartureAirport   *LocationType       `xml:"DepartureAirport,omitempty"`
	ArrivalAirport     *LocationType       `xml:"ArrivalAirport,omitempty"`
	FilingAirline      *CompanyNameType    `xml:"FilingAirline,omitempty"`
	FareReference      []*FareReference    `xml:"FareReference,omitempty"`
	DiscountPricing    *DiscountPricing    `xml:"DiscountPricing,omitempty"`
	MarketingAirline   []*CompanyNameType  `xml:"MarketingAirline,omitempty"`
	FareInfo           []*FareInfoTypeFare `xml:"FareInfo,omitempty"`
	LanguageRequested  string              `xml:"LanguageRequested,attr,omitempty"`
	LanguageReturned   string              `xml:"LanguageReturned,attr,omitempty"`
	FareRules          *FormattedTextType  `xml:"FareRules,omitempty"`
}

// FormattedTextType ...
type FormattedTextType struct {
	Title      string                         `xml:"Title,attr,omitempty"`
	Language   *Language                      `xml:"Language,attr,omitempty"`
	SubSection []*FormattedTextSubSectionType `xml:"SubSection,omitempty"`
}

// FormattedTextSubSectionType ...
type FormattedTextSubSectionType struct {
	Paragraph        *ParagraphType `xml:"Paragraph,omitempty"`
	SubTitle         string         `xml:"SubTitle,attr,omitempty"`
	SubCode          string         `xml:"SubCode,attr,omitempty"`
	SubSectionNumber int32          `xml:"SubSectionNumber,attr,omitempty"`
}

// VehicleLocationAdditionalDetailsTypeShuttleShuttleInfo ...
type VehicleLocationAdditionalDetailsTypeShuttleShuttleInfo struct {
	*FormattedTextType
	Type *LocationDetailShuttleInfoType `xml:"Type,attr,omitempty"`
}

// VehicleLocationInformationType ...
type VehicleLocationInformationType struct {
	*FormattedTextType
	Type string `xml:"Type,attr,omitempty"`
}

// BookFlightSegmentType ...
type BookFlightSegmentType struct {
	*FlightSegmentType
	NumberInParty          uint                                           `xml:"NumberInParty,attr,omitempty"`
	Distance               uint                                           `xml:"Distance,attr,omitempty"`
	StopoverInd            bool                                           `xml:"StopoverInd,attr,omitempty"`
	ValidConnectionInd     bool                                           `xml:"ValidConnectionInd,attr,omitempty"`
	LineNumber             int32                                          `xml:"LineNumber,attr,omitempty"`
	MarriageGrp            string                                         `xml:"MarriageGrp,omitempty"`
	ResBookDesigCode       string                                         `xml:"ResBookDesigCode,attr,omitempty"`
	Status                 string                                         `xml:"Status,attr,omitempty"`
	ETicketEligibility     string                                         `xml:"E_TicketEligibility,attr,omitempty"`
	MealCode               string                                         `xml:"MealCode,attr,omitempty"`
	ConnectionType         string                                         `xml:"ConnectionType,attr,omitempty"`
	ParticipationLevelCode string                                         `xml:"ParticipationLevelCode,attr,omitempty"`
	DateChangeNbr          string                                         `xml:"DateChangeNbr,attr,omitempty"`
	DepartureDay           *DayOfWeekType                                 `xml:"DepartureDay,attr,omitempty"`
	Comment                []*FreeTextType                                `xml:"Comment,omitempty"`
	StopLocation           []*StopLocation                                `xml:"StopLocation,omitempty"`
	BookingClassAvails     *ArrayOfBookFlightSegmentTypeBookingClassAvail `xml:"BookingClassAvails,omitempty"`
}

// StopLocation ...
type StopLocation struct {
	LocationCode      string `xml:"LocationCode,attr,omitempty"`
	CodeContext       string `xml:"CodeContext,attr,omitempty"`
	DepartureDateTime string `xml:"DepartureDateTime,attr,omitempty"`
	ArrivalDateTime   string `xml:"ArrivalDateTime,attr,omitempty"`
}

// FlightSegmentType ...
type FlightSegmentType struct {
	*FlightSegmentBaseType
	FlightNumber         string                         `xml:"FlightNumber,attr,omitempty"`
	TourOperatorFlightID string                         `xml:"TourOperatorFlightID,attr,omitempty"`
	MarketingAirline     *MarketingAirline              `xml:"MarketingAirline,omitempty"`
	Baggages             *ArrayOfFreeBaggageSegmentItem `xml:"Baggages,omitempty"`
}

// MarketingAirline ...
type MarketingAirline struct {
	*CompanyNameType
	SingleVendorInd string `xml:"SingleVendorInd,attr,omitempty"`
}

// FlightSegmentBaseType ...
type FlightSegmentBaseType struct {
	StopQuantity      uint                  `xml:"StopQuantity,attr,omitempty"`
	RPH               string                `xml:"RPH,attr,omitempty"`
	InfoSource        string                `xml:"InfoSource,attr,omitempty"`
	FlightDuration    time.Time             `xml:"FlightDuration,omitempty"`
	DepartureAirport  *Airport              `xml:"DepartureAirport,omitempty"`
	ArrivalAirport    *Airport              `xml:"ArrivalAirport,omitempty"`
	DepartureDateTime string                `xml:"DepartureDateTime,attr,omitempty"`
	ArrivalDateTime   string                `xml:"ArrivalDateTime,attr,omitempty"`
	Equipment         []*EquipmentType      `xml:"Equipment,omitempty"`
	OperatingAirline  *OperatingAirlineType `xml:"OperatingAirline,omitempty"`
}

// Airport ...
type Airport struct {
	LocationCode string `xml:"LocationCode,attr,omitempty"`
	CodeContext  string `xml:"CodeContext,attr,omitempty"`
	Terminal     string `xml:"Terminal,attr,omitempty"`
	Gate         string `xml:"Gate,attr,omitempty"`
}

// ArrayOfFreeBaggageSegmentItem ...
type ArrayOfFreeBaggageSegmentItem struct {
	Baggage []*Baggage `xml:"Baggage,omitempty"`
}

// ArrayOfBookFlightSegmentTypeBookingClassAvail ...
type ArrayOfBookFlightSegmentTypeBookingClassAvail struct {
	BookingClassAvail []*BookingClassAvail `xml:"BookingClassAvail,omitempty"`
}

// BookingClassAvail ...
type BookingClassAvail struct {
	ResBookDesigQuantity   uint   `xml:"ResBookDesigQuantity,attr,omitempty"`
	ResBookDesigCode       string `xml:"ResBookDesigCode,attr,omitempty"`
	ResBookDesigStatusCode string `xml:"ResBookDesigStatusCode,attr,omitempty"`
	RPH                    string `xml:"RPH,attr,omitempty"`
	AvailablePTC           string `xml:"AvailablePTC,attr,omitempty"`
	ResBookDesigCabinCode  string `xml:"ResBookDesigCabinCode,attr,omitempty"`
	FareBasis              string `xml:"FareBasis,attr,omitempty"`
}

// ArrayOfAirItineraryPricingInfoTypeFareInfo ...
type ArrayOfAirItineraryPricingInfoTypeFareInfo struct {
	FareInfo []*TFareInfo `xml:"FareInfo,omitempty"`
}

// TFareInfo ...
type TFareInfo struct {
	*FareInfoType
	TPAExtensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
}

// BookingPriceInfoType ...
type BookingPriceInfoType struct {
	*AirItineraryPricingInfoType
	RepriceRequired         bool                         `xml:"RepriceRequired,attr,omitempty"`
	PriceRequestInformation *PriceRequestInformationType `xml:"PriceRequestInformation,omitempty"`
}

// PriceRequestInformationType ...
type PriceRequestInformationType struct {
	NegotiatedFaresOnly bool                  `xml:"NegotiatedFaresOnly,attr,omitempty"`
	Reprice             bool                  `xml:"Reprice,attr,omitempty"`
	FareQualifier       string                `xml:"FareQualifier,attr,omitempty"`
	TicketingCountry    string                `xml:"TicketingCountry,attr,omitempty"`
	CurrencyCode        string                `xml:"CurrencyCode,attr,omitempty"`
	CabinType           *CabinType            `xml:"CabinType,attr,omitempty"`
	RebookOption        []*RebookOption       `xml:"RebookOption,omitempty"`
	PricingSource       *PricingSourceType    `xml:"PricingSource,attr,omitempty"`
	NegotiatedFareCode  []*NegotiatedFareCode `xml:"NegotiatedFareCode,omitempty"`
}

// RebookOption ...
type RebookOption struct {
	FlightSegmentRPH string `xml:"FlightSegmentRPH,attr,omitempty"`
	ResBookDesigCode string `xml:"ResBookDesigCode,attr,omitempty"`
}

// NegotiatedFareCode ...
type NegotiatedFareCode struct {
	SecondaryCode string `xml:"SecondaryCode,attr,omitempty"`
	
	SupplierCode string `xml:"SupplierCode,attr,omitempty"`
	
	Value string `xml:"Value,attr,omitempty"`
}

// AirReservationType ...
type AirReservationType struct {
	LastModified        string                          `xml:"LastModified,attr,omitempty"`
	Fulfillment         *Fulfillment                    `xml:"Fulfillment,omitempty"`
	BookingReferenceID  []*UniqueIDType                 `xml:"BookingReferenceID,omitempty"`
	PricingOverview     *PricingOverview                `xml:"PricingOverview,omitempty"`
	TravelerInfo        *TravelerInfoType               `xml:"TravelerInfo,omitempty"`
	AirItinerary        *AirItineraryType               `xml:"AirItinerary,omitempty"`
	OSIElements         *ArrayOfOSIelement              `xml:"OSIElements,omitempty"`
	Ticketing           []*TicketingInfoType            `xml:"Ticketing,omitempty"`
	PriceInfo           *BookingPriceInfoType           `xml:"PriceInfo,omitempty"`
	PNRRemarks          *ArrayOfPNRRemarkInfo           `xml:"PNRRemarks,omitempty"`
	PriceMessageInfo    *PriceMessageInfoType           `xml:"PriceMessageInfo,omitempty"`
	Comment             *FormattedTextTextType          `xml:"Comment,omitempty"`
	FlightMiniRules     *MiniRuleResponseInfoType       `xml:"FlightMiniRules,omitempty"`
	FlightRulePenalties []*FareRuleResponseInfoType     `xml:"FlightRulePenalties,omitempty"`
	Queues              *ArrayOfAirReservationTypeQueue `xml:"Queues,omitempty"`
}

// PricingOverview ...
type PricingOverview struct {
	PricingIndicator []struct {
		Type string `xml:"Type,attr,omitempty"`
		
		ExcludeInd bool `xml:"ExcludeInd,attr,omitempty"`
	} `xml:"PricingIndicator,omitempty"`
	
	Account []struct {
		Code string `xml:"Code,attr,omitempty"`
	} `xml:"Account,omitempty"`
	
	Comment []*FreeTextType `xml:"Comment,omitempty"`
	
	StatisticalCode string `xml:"StatisticalCode,attr,omitempty"`
	
	ValidatingAirlineCode string `xml:"ValidatingAirlineCode,attr,omitempty"`
	
	DepartureDate string `xml:"DepartureDate,attr,omitempty"`
	
	PriceType string `xml:"PriceType,attr,omitempty"`
	
	NUC_Rate float64 `xml:"NUC_Rate,attr,omitempty"`
	
	ExchangeRate float64 `xml:"ExchangeRate,attr,omitempty"`
}

// Fulfillment ...
type Fulfillment struct {
	PaymentDetails *ArrayOfPaymentDetailType `xml:"PaymentDetails,omitempty"`
	
	DeliveryAddress *AddressType `xml:"DeliveryAddress,omitempty"`
}

// ArrayOfOSIelement ...
type ArrayOfOSIelement struct {
	OSIelement []*OSIelement `xml:"OSIelement,omitempty"`
}

// OSIelement ...
type OSIelement struct {
	AirlineID string `xml:"AirlineID,omitempty"`
	
	Text string `xml:"Text,omitempty"`
	
	SubjectQualifier string `xml:"SubjectQualifier,omitempty"`
	
	Type string `xml:"Type,omitempty"`
	
	NumberOfAssociatedPassenger int32 `xml:"NumberOfAssociatedPassenger,omitempty"`
	
	NumberOfAssociatedSegments string `xml:"NumberOfAssociatedSegments,omitempty"`
	
	LineNumber int32 `xml:"LineNumber,omitempty"`
}

// ArrayOfAirItineraryTypeOriginDestinationOption ...
type ArrayOfAirItineraryTypeOriginDestinationOption struct {
	OriginDestinationOption []*OriginDestinationOption `xml:"OriginDestinationOption,omitempty"`
}

// OriginDestinationOption ...
type OriginDestinationOption struct {
	*OriginDestinationOptionType
	SequenceNumber    string                       `xml:"SequenceNumber,omitempty"`
	ItinBaseFare      *FareType                    `xml:"ItinBaseFare,omitempty"`
	RefNumber         int32                        `xml:"RefNumber,attr,omitempty"`
	DirectionId       int32                        `xml:"DirectionId,attr,omitempty"`
	ElapsedTime       string                       `xml:"ElapsedTime,attr,omitempty"`
	ProviderType      string                       `xml:"ProviderType,attr,omitempty"`
	DirectionInd      string                       `xml:"DirectionInd,attr,omitempty"`
	OptionPricingInfo *AirItineraryPricingInfoType `xml:"OptionPricingInfo,omitempty"`
}

// OriginDestinationOptionType ...
type OriginDestinationOptionType struct {
	FlightSegment []*FlightSegment `xml:"FlightSegment,omitempty"`
}

// FlightSegment ...
type FlightSegment struct {
	*BookFlightSegmentType
	TPA_Extensions *TPAExtensionsType `xml:"TPA_Extensions,omitempty"`
}

// ArrayOfOriginDestinationCombinationType ...
type ArrayOfOriginDestinationCombinationType struct {
	OriginDestinationCombination []*OriginDestinationCombinationType `xml:"OriginDestinationCombination,omitempty"`
}

// OriginDestinationCombinationType ...
type OriginDestinationCombinationType struct {
	ValidatingAirlineCode string  `xml:"ValidatingAirlineCode,attr,omitempty"`
	ForceETicket          bool    `xml:"ForceETicket,attr,omitempty"`
	OriginIndex           int32   `xml:"OriginIndex,attr,omitempty"`
	IndexList             string  `xml:"IndexList,attr,omitempty"`
	DestinationIndex      int32   `xml:"DestinationIndex,attr,omitempty"`
	CombinationID         int32   `xml:"CombinationID,attr,omitempty"`
	RecommendationID      int32   `xml:"RecommendationID,attr,omitempty"`
	ETicketEligibility    string  `xml:"E_TicketEligibility,attr,omitempty"`
	ServiceFeeAmount      float64 `xml:"ServiceFeeAmount,attr,omitempty"`
}

type TravelerInfoType struct {
	AirTraveler       []*AirTraveler           `xml:"AirTraveler,omitempty"`
	SpecialReqDetails []*SpecialReqDetailsType `xml:"SpecialReqDetails,omitempty"`
}

// AirTraveler ...
type AirTraveler struct {
	*AirTravelerType
	ETicketNumber        string                      `xml:"eTicketNumber,attr,omitempty"`
	OSIMessage           []string                    `xml:"OSIMessage,omitempty"`
	Comment              []*Comment                  `xml:"Comment,omitempty"`
	ETicketDocument      *ETicketInfo                `xml:"eTicketDocument,omitempty"`
	ETicketURLs          []*ETicketURLs              `xml:"ETicketURLs,omitempty"`
	ETicketDocuments     *ArrayOfETicketInfo         `xml:"eTicketDocuments,omitempty"`
	AncillaryProducts    *AncillaryProductsType      `xml:"AncillaryProducts,omitempty"`
	FreeBaggageAllowance *FreeBaggageAllowancesTypes `xml:"FreeBaggageAllowance,omitempty"`
}

// Phone ...
type Phone struct {
	AreaCityCode string         `xml:"AreaCityCode,attr,omitempty"`
	CountryCode  string         `xml:"CountryCode,attr,omitempty"`
	PhoneNumber  string         `xml:"PhoneNumber,attr,omitempty"`
	PhoneType    *PhoneTypeEnum `xml:"PhoneType,attr,omitempty"`
}

// Email ...
type Email struct {
	*EmailType
	Operation *ActionType `xml:"Operation,attr,omitempty"`
}

// Address ...
type Address struct {
	*AddressType
	Operation *ActionType `xml:"Operation,attr,omitempty"`
}

// Document ...
type Document struct {
	*DocumentType
	Operation *ActionType `xml:"Operation,attr,omitempty"`
}

// TravelerRefNumber ...
type TravelerRefNumber struct {
	RPH              string `xml:"RPH,attr,omitempty"`
	SurnameRefNumber string `xml:"SurnameRefNumber,attr,omitempty"`
}
type AirTravelerType struct {
	NumberOfBaggages      int32                      `xml:"NumberOfBaggages,omitempty"`
	PersonName            *PersonNameType            `xml:"PersonName,omitempty"`
	Telephone             []*Phone                   `xml:"Telephone,omitempty"`
	Email                 []*Email                   `xml:"Email,omitempty"`
	Address               []*Address                 `xml:"Address,omitempty"`
	Document              []*Document                `xml:"Document,omitempty"`
	PassengerTypeQuantity *PassengerTypeQuantityType `xml:"PassengerTypeQuantity,omitempty"`
	TravelerRefNumber     *TravelerRefNumber         `xml:"TravelerRefNumber,omitempty"`
	FlightSegmentRPHs     *FlightSegmentRPHs         `xml:"FlightSegmentRPHs,omitempty"`
	BirthDate             string                     `xml:"BirthDate,omitempty"`
	Gender                string                     `xml:"Gender,attr,omitempty"`
	ShareSynchInd         string                     `xml:"ShareSynchInd,attr,omitempty"`
	ShareMarketInd        string                     `xml:"ShareMarketInd,attr,omitempty"`
	CurrencyCode          string                     `xml:"CurrencyCode,attr,omitempty"`
	PassengerTypeCode     string                     `xml:"PassengerTypeCode,attr,omitempty"`
	AccompaniedByInfant   bool                       `xml:"AccompaniedByInfant,attr,omitempty"`
}

// Comment ...
type Comment struct {
	*FormattedTextTextType
	Name string `xml:"Name,attr,omitempty"`
}

// FlightSegmentRPHs ...
type FlightSegmentRPHs struct {
	FlightSegmentRPH []string `xml:"FlightSegmentRPH,omitempty"`
}

// ETicketInfo ...
type ETicketInfo struct {
	AgencyTitle            string                       `xml:"AgencyTitle,omitempty"`
	AgencyTelNo            string                       `xml:"AgencyTelNo,omitempty"`
	AgentID                string                       `xml:"AgentID,omitempty"`
	IATAID                 string                       `xml:"IATAID,omitempty"`
	PassengerNameSurname   string                       `xml:"PassengerName_Surname,omitempty"`
	FrequentFlyer          string                       `xml:"FrequentFlyer,omitempty"`
	IssuingAirline         string                       `xml:"IssuingAirline,omitempty"`
	BookingRef             string                       `xml:"BookingRef,omitempty"`
	FareCalculation        string                       `xml:"FareCalculation,omitempty"`
	Endorsements           string                       `xml:"Endorsements,omitempty"`
	ExchangeRate           float64                      `xml:"ExchangeRate,omitempty"`
	PaymentType            string                       `xml:"PaymentType,omitempty"`
	AirFare                float64                      `xml:"AirFare,omitempty"`
	AirFareCurrency        string                       `xml:"AirFareCurrency,omitempty"`
	EquivalentFare         float64                      `xml:"EquivalentFare,omitempty"`
	ServiceFee             float64                      `xml:"ServiceFee,omitempty"`
	EquivalentFareCurrency string                       `xml:"EquivalentFareCurrency,omitempty"`
	TotalFare              float64                      `xml:"TotalFare,omitempty"`
	TotalFareCurrency      string                       `xml:"TotalFareCurrency,omitempty"`
	TicketCurrency         string                       `xml:"TicketCurrency,omitempty"`
	ControlNumbers         string                       `xml:"ControlNumbers,omitempty"`
	TicketNumber           string                       `xml:"TicketNumber,attr,omitempty"`
	TicketingDate          time.Time                    `xml:"TicketingDate,omitempty"`
	AgencyAddress          *AddressType                 `xml:"AgencyAddress,omitempty"`
	Taxes                  *ArrayOfTaxInfo              `xml:"Taxes,omitempty"`
	Itineraries            *ArrayOfETicketItineraryInfo `xml:"Itineraries,omitempty"`
}

// ETicketURLs ...
type ETicketURLs struct {
	ETicketURL   string `xml:"ETicketURL,attr,omitempty"`
	ProviderType string `xml:"ProviderType,attr,omitempty"`
}

// ArrayOfTaxInfo ...
type ArrayOfTaxInfo struct {
	TaxInfo []*TaxInfo `xml:"TaxInfo,omitempty"`
}

// TaxInfo ...
type TaxInfo struct {
	TaxType     string  `xml:"TaxType,omitempty"`
	Currency    string  `xml:"Currency,omitempty"`
	TaxNature   string  `xml:"TaxNature,omitempty"`
	CountryCode string  `xml:"CountryCode,omitempty"`
	Amount      float64 `xml:"Amount,omitempty"`
}

// ArrayOfETicketItineraryInfo ...
type ArrayOfETicketItineraryInfo struct {
	ETicketItineraryInfo []*ETicketItineraryInfo `xml:"ETicketItineraryInfo,omitempty"`
}

// ETicketItineraryInfo ...
type ETicketItineraryInfo struct {
	From                     string    `xml:"From,omitempty"`
	FromTerminal             string    `xml:"FromTerminal,omitempty"`
	To                       string    `xml:"To,omitempty"`
	ToTerminal               string    `xml:"ToTerminal,omitempty"`
	Carrier                  string    `xml:"Carrier,omitempty"`
	FlightNo                 int32     `xml:"FlightNo,omitempty"`
	OperatingAirlineCode     string    `xml:"OperatingAirlineCode,omitempty"`
	MarketingAirlineCode     string    `xml:"MarketingAirlineCode,omitempty"`
	Class                    string    `xml:"Class,omitempty"`
	DepartureDate            time.Time `xml:"DepartureDate,omitempty"`
	ArrivalDate              time.Time `xml:"ArrivalDate,omitempty"`
	FareBasis                string    `xml:"FareBasis,omitempty"`
	NotValidBefore           time.Time `xml:"NotValidBefore,omitempty"`
	NotValidAfter            time.Time `xml:"NotValidAfter,omitempty"`
	BaggageWeight            float64   `xml:"BaggageWeight,omitempty"`
	BaggageWeightMeasureUnit string    `xml:"BaggageWeightMeasureUnit,omitempty"`
	BaggageQuantity          float64   `xml:"BaggageQuantity,omitempty"`
	Status                   string    `xml:"Status,omitempty"`
}

// ArrayOfETicketInfo ...
type ArrayOfETicketInfo struct {
	ETicketInfo []*ETicketInfo `xml:"ETicketInfo,omitempty"`
}

type AncillaryProductsType struct {
	AncillaryProduct []*AncillaryProductType `xml:"AncillaryProduct,omitempty"`
}

type AncillaryProductType struct {
	SSRCode string `xml:"SSRCode,omitempty"`
	
	AirlineCode string `xml:"AirlineCode,omitempty"`
	
	Amount float64 `xml:"Amount,omitempty"`
	
	PricedAmount float64 `xml:"PricedAmount,omitempty"`
	
	PricedAmountInTicketCurrency float64 `xml:"PricedAmountInTicketCurrency,omitempty"`
	
	Tax float64 `xml:"Tax,omitempty"`
	
	Currency string `xml:"Currency,omitempty"`
	
	SegmentReference int32 `xml:"SegmentReference,omitempty"`
	
	Description string `xml:"Description,omitempty"`
	
	Status string `xml:"Status,omitempty"`
	
	UniqueReference string `xml:"UniqueReference,omitempty"`
	
	ConfirmationNumber string `xml:"ConfirmationNumber,omitempty"`
	
	EticketInfo *AncillaryEticketInfoType `xml:"EticketInfo,omitempty"`
	
	PNRPassengerOrderID int32 `xml:"PNRPassengerOrderID,omitempty"`
	
	PNRSegmentOrderID int32 `xml:"PNRSegmentOrderID,omitempty"`
	
	AncillaryProductLineNumber int32 `xml:"AncillaryProductLineNumber,omitempty"`
}

type AncillaryEticketInfoType struct {
	DocumentNumber string `xml:"DocumentNumber,omitempty"`
	
	IssuanceTypeCode string `xml:"IssuanceTypeCode,omitempty"`
	
	IssuanceType string `xml:"IssuanceType,omitempty"`
	
	DateOfIssuance time.Time `xml:"DateOfIssuance,omitempty"`
	
	PlaceOfIssue string `xml:"PlaceOfIssue,omitempty"`
	
	InternationalIndicatorCode string `xml:"InternationalIndicatorCode,omitempty"`
	
	InternationalIndicator string `xml:"InternationalIndicator,omitempty"`
	
	IATANumber string `xml:"IATANumber,omitempty"`
	
	EndorsableCode string `xml:"EndorsableCode,omitempty"`
	
	Endorsable string `xml:"Endorsable,omitempty"`
	
	AncillaryTicketRemark []*AncillaryTicketRemarkType `xml:"AncillaryTicketRemark,omitempty"`
	
	AncillaryTicketCoupon []*AncillaryTicketCouponType `xml:"AncillaryTicketCoupon,omitempty"`
	
	AncillaryTicketFareInfo []*AncillaryTicketFareInfoType `xml:"AncillaryTicketFareInfo,omitempty"`
	
	AncillaryTicketFOP *AncillaryTicketFOPType `xml:"AncillaryTicketFOP,omitempty"`
}

type AncillaryTicketRemarkType struct {
	RemarkTypeCode string `xml:"RemarkTypeCode,omitempty"`
	
	RemarkType string `xml:"RemarkType,omitempty"`
	
	Remark string `xml:"Remark,omitempty"`
}

type AncillaryTicketCouponType struct {
	CpnNumber string `xml:"CpnNumber,omitempty"`
	
	IssuanceSubCode string `xml:"IssuanceSubCode,omitempty"`
	
	OperatingAirline string `xml:"OperatingAirline,omitempty"`
	
	OriginDestination string `xml:"OriginDestination,omitempty"`
	
	StatusCode string `xml:"StatusCode,omitempty"`
	
	Status string `xml:"Status,omitempty"`
	
	ConnectionWithEticketNumber string `xml:"ConnectionWithEticketNumber,omitempty"`
	
	ConnectionWithEticketType string `xml:"ConnectionWithEticketType,omitempty"`
	
	ConnectionWithEticket string `xml:"ConnectionWithEticket,omitempty"`
	
	CouponAmount string `xml:"CouponAmount,omitempty"`
	
	CouponCurrency string `xml:"CouponCurrency,omitempty"`
	
	CouponValue string `xml:"CouponValue,omitempty"`
	
	CouponDetail *AncillaryTicketCouponDetailType `xml:"CouponDetail,omitempty"`
}

type AncillaryTicketCouponDetailType struct {
	RefundableCode string `xml:"RefundableCode,omitempty"`
	
	Refundable string `xml:"Refundable,omitempty"`
	
	ExchangeableCode string `xml:"ExchangeableCode,omitempty"`
	
	Exchangeable string `xml:"Exchangeable,omitempty"`
	
	InterlinableCode string `xml:"InterlinableCode,omitempty"`
	
	Interlinable string `xml:"Interlinable,omitempty"`
	
	ConsummedIssuanceCode string `xml:"ConsummedIssuanceCode,omitempty"`
	
	ConsummedIssuance string `xml:"ConsummedIssuance,omitempty"`
	
	QuantityTypeCode string `xml:"QuantityTypeCode,omitempty"`
	
	QuantityType string `xml:"QuantityType,omitempty"`
	
	UnitQualifierCode string `xml:"UnitQualifierCode,omitempty"`
	
	UnitQualifier string `xml:"UnitQualifier,omitempty"`
	
	FreeAllowance string `xml:"FreeAllowance,omitempty"`
	
	ExcessBaggage string `xml:"ExcessBaggage,omitempty"`
	
	RatePerUnitAmount string `xml:"RatePerUnitAmount,omitempty"`
	
	RatePerUnitCurrency string `xml:"RatePerUnitCurrency,omitempty"`
	
	RatePerUnit string `xml:"RatePerUnit,omitempty"`
}

type AncillaryTicketFareInfoType struct {
	FareTypeCode string `xml:"FareTypeCode,omitempty"`
	
	FareType string `xml:"FareType,omitempty"`
	
	Amount string `xml:"Amount,omitempty"`
	
	Currency string `xml:"Currency,omitempty"`
}

type AncillaryTicketFOPType struct {
	PaymentTypeCode string `xml:"PaymentTypeCode,omitempty"`
	
	PaymentType string `xml:"PaymentType,omitempty"`
	
	VendorCode string `xml:"VendorCode,omitempty"`
	
	CCNumber string `xml:"CCNumber,omitempty"`
	
	PaymentDetail string `xml:"PaymentDetail,omitempty"`
	
	PaymentAmount string `xml:"PaymentAmount,omitempty"`
	
	PaymentCurrency string `xml:"PaymentCurrency,omitempty"`
	
	PaymentAmountDetail string `xml:"PaymentAmountDetail,omitempty"`
}

type FreeBaggageAllowancesTypes struct {
	FreeBaggageAllowance *ArrayOfFreeBaggageAllowanceTypes `xml:"FreeBaggageAllowance,omitempty"`
	
	Errors *ErrorsType `xml:"Errors,omitempty"`
}

type ArrayOfFreeBaggageAllowanceTypes struct {
	FreeBaggageAllowanceTypes []*FreeBaggageAllowanceTypes `xml:"FreeBaggageAllowanceTypes,omitempty"`
}

type FreeBaggageAllowanceTypes struct {
	ItineraryReference int32 `xml:"ItineraryReference,attr,omitempty"`
	
	SegmentReference int32 `xml:"SegmentReference,attr,omitempty"`
	
	PassengerType string `xml:"PassengerType,attr,omitempty"`
	
	Unit string `xml:"Unit,attr,omitempty"`
	
	Quantity int32 `xml:"Quantity,attr,omitempty"`
}

type SpecialReqDetailsType struct {
	SeatRequests *ArrayOfSpecialReqDetailsTypeSeatRequest `xml:"SeatRequests,omitempty"`
	
	SpecialServiceRequests *ArrayOfSpecialReqDetailsTypeSpecialServiceRequest `xml:"SpecialServiceRequests,omitempty"`
	
	OtherServiceInformations *ArrayOfSpecialReqDetailsTypeOtherServiceInformation `xml:"OtherServiceInformations,omitempty"`
	
	Remarks *ArrayOfSpecialReqDetailsTypeRemark `xml:"Remarks,omitempty"`
	
	SpecialRemarks *ArrayOfSpecialReqDetailsTypeSpecialRemark `xml:"SpecialRemarks,omitempty"`
}

type ArrayOfSpecialReqDetailsTypeSeatRequest struct {
	SeatRequest []struct {
		*SeatRequestType
		
		TravelerRefNumberRPHList []string `xml:"TravelerRefNumberRPHList,attr,omitempty"`
		
		FlightRefNumberRPHList []string `xml:"FlightRefNumberRPHList,attr,omitempty"`
		
		PartialSeatingInd bool `xml:"PartialSeatingInd,attr,omitempty"`
	} `xml:"SeatRequest,omitempty"`
}

type SeatRequestType struct {
	DepartureAirport *LocationType `xml:"DepartureAirport,omitempty"`
	
	ArrivalAirport *LocationType `xml:"ArrivalAirport,omitempty"`
	
	SeatNumber string `xml:"SeatNumber,attr,omitempty"`
	
	SeatPrice string `xml:"SeatPrice,attr,omitempty"`
	
	SeatPriceCurrency string `xml:"SeatPriceCurrency,attr,omitempty"`
	
	SeatPreference string `xml:"SeatPreference,attr,omitempty"`
	
	DepartureDate string `xml:"DepartureDate,attr,omitempty"`
	
	FlightNumber string `xml:"FlightNumber,attr,omitempty"`
	
	Status string `xml:"Status,attr,omitempty"`
	
	ItineraryID string `xml:"ItineraryID,attr,omitempty"`
	
	SegmentID string `xml:"SegmentID,attr,omitempty"`
}

type ArrayOfSpecialReqDetailsTypeSpecialServiceRequest struct {
	SpecialServiceRequest []struct {
		*SpecialServiceRequestType
		
		AnimalBoxType string `xml:"AnimalBoxType,omitempty"`
		
		FlightLeg *FlightLegType `xml:"FlightLeg,omitempty"`
		
		Height int32 `xml:"Height,attr,omitempty"`
		
		Width int32 `xml:"Width,attr,omitempty"`
		
		Length int32 `xml:"Length,attr,omitempty"`
		
		Weight int32 `xml:"Weight,attr,omitempty"`
		
		OtherEquipment string `xml:"OtherEquipment,attr,omitempty"`
		
		BaggageName string `xml:"BaggageName,attr,omitempty"`
		
		DisabledPassenger string `xml:"DisabledPassenger,attr,omitempty"`
		
		OtherAnimalType string `xml:"OtherAnimalType,attr,omitempty"`
		
		SSRType string `xml:"SSRType,attr,omitempty"`
		
		TravelerRefNumberRPHList []string `xml:"TravelerRefNumberRPHList,attr,omitempty"`
		
		FlightRefNumberRPHList []string `xml:"FlightRefNumberRPHList,attr,omitempty"`
		
		Notes string `xml:"Notes,attr,omitempty"`
		
		WithDog bool `xml:"WithDog,attr,omitempty"`
		
		WithDogSpecified bool `xml:"WithDogSpecified,attr,omitempty"`
		
		AirlineCode string `xml:"AirlineCode,attr,omitempty"`
	} `xml:"SpecialServiceRequest,omitempty"`
}

type SpecialServiceRequestType struct {
	Airline *CompanyNameType `xml:"Airline,omitempty"`
	
	Text string `xml:"Text,omitempty"`
	
	SSRCode string `xml:"SSRCode,attr,omitempty"`
	
	ServiceQuantity int32 `xml:"ServiceQuantity,attr,omitempty"`
	
	Status string `xml:"Status,attr,omitempty"`
}

type FlightLegType struct {
	DepartureAirport struct {
		LocationCode string `xml:"LocationCode,attr,omitempty"`
		
		CodeContext string `xml:"CodeContext,attr,omitempty"`
	} `xml:"DepartureAirport,omitempty"`
	
	ArrivalAirport struct {
		LocationCode string `xml:"LocationCode,attr,omitempty"`
		
		CodeContext string `xml:"CodeContext,attr,omitempty"`
	} `xml:"ArrivalAirport,omitempty"`
	
	FlightNumber string `xml:"FlightNumber,attr,omitempty"`
	
	ResBookDesigCode string `xml:"ResBookDesigCode,attr,omitempty"`
	
	Date time.Time `xml:"Date,attr,omitempty"`
}

type ArrayOfSpecialReqDetailsTypeOtherServiceInformation struct {
	OtherServiceInformation []struct {
		*OtherServiceInfoType
		
		RPH string `xml:"RPH,attr,omitempty"`
		
		Operation *ActionType `xml:"Operation,attr,omitempty"`
	} `xml:"OtherServiceInformation,omitempty"`
}

type OtherServiceInfoType struct {
	TravelerRefNumber []struct {
		RPH string `xml:"RPH,attr,omitempty"`
		
		SurnameRefNumber string `xml:"SurnameRefNumber,attr,omitempty"`
	} `xml:"TravelerRefNumber,omitempty"`
	
	Airline *CompanyNameType `xml:"Airline,omitempty"`
	
	Text string `xml:"Text,omitempty"`
}

type ArrayOfSpecialReqDetailsTypeRemark struct {
	Remark []struct {
		Value string
		
		RPH string `xml:"RPH,attr,omitempty"`
		
		Operation *ActionType `xml:"Operation,attr,omitempty"`
	} `xml:"Remark,omitempty"`
}

type ArrayOfSpecialReqDetailsTypeSpecialRemark struct {
	SpecialRemark []struct {
		*SpecialRemarkType
		
		FlightLeg *FlightLegType `xml:"FlightLeg,omitempty"`
		
		Operation *ActionType `xml:"Operation,attr,omitempty"`
		
		RPH string `xml:"RPH,attr,omitempty"`
	} `xml:"SpecialRemark,omitempty"`
}

type SpecialRemarkType struct {
	TravelerRefNumber []struct {
		RPH string `xml:"RPH,attr,omitempty"`
		
		SurnameRefNumber string `xml:"SurnameRefNumber,attr,omitempty"`
	} `xml:"TravelerRefNumber,omitempty"`
	
	FlightRefNumber []struct {
		RPH string `xml:"RPH,attr,omitempty"`
	} `xml:"FlightRefNumber,omitempty"`
	
	Text string `xml:"Text,omitempty"`
	
	Airline []*CompanyNameType `xml:"Airline,omitempty"`
	
	AuthorizedViewers *ArrayOfSpecialRemarkTypeAuthorizedViewer `xml:"AuthorizedViewers,omitempty"`
	
	RemarkType string `xml:"RemarkType,attr,omitempty"`
}

// ArrayOfSpecialRemarkTypeAuthorizedViewer ...
type ArrayOfSpecialRemarkTypeAuthorizedViewer struct {
	AuthorizedViewer []struct {
		ViewerCode string `xml:"ViewerCode,attr,omitempty"`
		
		ViewerCarrierCode string `xml:"ViewerCarrierCode,attr,omitempty"`
	} `xml:"AuthorizedViewer,omitempty"`
}

// ArrayOfPaymentDetailType ...
type ArrayOfPaymentDetailType struct {
	PaymentDetail []*PaymentDetailType `xml:"PaymentDetail,omitempty"`
}

// TicketingInfoType ...
type TicketingInfoType struct {
	TicketAdvisory         []*FreeTextType `xml:"TicketAdvisory,omitempty"`
	TicketTimeLimit        string          `xml:"TicketTimeLimit,attr,omitempty"`
	OldTicketTimeLimit     string          `xml:"OldTicketTimeLimit,attr,omitempty"`
	TicketType             TicketType      `xml:"TicketType,attr,omitempty"`
	TicketingStatus        string          `xml:"TicketingStatus,attr,omitempty"`
	FlightSegmentRefNumber []string        `xml:"FlightSegmentRefNumber,attr,omitempty"`
	TravelerRefNumber      []string        `xml:"TravelerRefNumber,attr,omitempty"`
	ReverseTktgSegmentsInd bool            `xml:"ReverseTktgSegmentsInd,attr,omitempty"`
	PseudoCityCode         string          `xml:"PseudoCityCode,attr,omitempty"`
	RequestedTicketingDate string          `xml:"RequestedTicketingDate,attr,omitempty"`
	TimeLimitMinutes       int32           `xml:"TimeLimitMinutes,attr,omitempty"`
}

type TicketingInfoRSType struct {
	*TicketingInfoType
	
	ETicketNumber string `xml:"eTicketNumber,attr,omitempty"`
}

type ArrayOfAirReservationTypeQueue struct {
	Queue []struct {
		PseudoCityCode string `xml:"PseudoCityCode,attr,omitempty"`
		
		QueueNumber string `xml:"QueueNumber,attr,omitempty"`
		
		QueueCategory string `xml:"QueueCategory,attr,omitempty"`
		
		SystemCode string `xml:"SystemCode,attr,omitempty"`
		
		QueueID string `xml:"QueueID,attr,omitempty"`
		
		DateTime string `xml:"DateTime,attr,omitempty"`
		
		Text string `xml:"Text,attr,omitempty"`
		
		CarrierCode string `xml:"CarrierCode,attr,omitempty"`
	} `xml:"Queue,omitempty"`
}

type FareRuleResponseInfoType struct {
	FareRuleInfo   []*FareRuleResponseInfoTypeFareRuleInfo `xml:"FareRuleInfo,omitempty"`
	Routing        []*FareRuleResponseInfoTypeRouting      `xml:"Routing,omitempty"`
	AdvisoryInfo   []*FareRuleResponseInfoTypeAdvisoryInfo `xml:"AdvisoryInfo,omitempty"`
	ItineraryIndex int32                                   `xml:"ItineraryIndex,attr,omitempty"`
}

type FareRuleResponseInfoTypeRouting struct {
	Info []*FareRuleResponseInfoTypeRoutingInfo `xml:"Info,omitempty"`
	
	Number int32 `xml:"Number,attr,omitempty"`
	
	RoutingConstructedInd bool `xml:"RoutingConstructedInd,attr,omitempty"`
	
	MaximumPermittedMileage int32 `xml:"MaximumPermittedMileage,attr,omitempty"`
	
	RoutingRestriction string `xml:"RoutingRestriction,attr,omitempty"`
}

type FareRuleResponseInfoTypeRoutingInfo struct {
	Direction string `xml:"Direction,attr,omitempty"`
	
	Text string `xml:"Text,attr,omitempty"`
}

type FareRuleResponseInfoTypeAdvisoryInfo struct {
	*FreeTextType
	
	AdvisoryCode string `xml:"AdvisoryCode,attr,omitempty"`
}

type MiniRuleResponseInfoType struct {
	MiniRuleInfo []*MiniRules `xml:"MiniRuleInfo,omitempty"`
}

type MiniRules struct {
	TermsAndCondition *ArrayOfFlightTermsAndCondition `xml:"TermsAndCondition,omitempty"`
}

type ArrayOfFlightTermsAndCondition struct {
	FlightTermsAndCondition []*FlightTermsAndCondition `xml:"FlightTermsAndCondition,omitempty"`
}

type FlightTermsAndCondition struct {
	RuleInfo *ArrayOfRuleInfos `xml:"RuleInfo,omitempty"`
	
	Destination string `xml:"Destination,attr,omitempty"`
	
	Origin string `xml:"Origin,attr,omitempty"`
	
	AirlineCode string `xml:"AirlineCode,attr,omitempty"`
	
	AirlineName string `xml:"AirlineName,attr,omitempty"`
	
	FareBasisCode string `xml:"FareBasisCode,attr,omitempty"`
}

type ArrayOfRuleInfos struct {
	RuleInfos []*RuleInfos `xml:"RuleInfos,omitempty"`
}

type RuleInfos struct {
	CategoryType string `xml:"CategoryType,omitempty"`
	
	RuleText []*RuleInfosDetails `xml:"RuleText,omitempty"`
}

type RuleInfosDetails struct {
	RuleTitle string `xml:"RuleTitle,attr,omitempty"`
	
	RuleText string `xml:"RuleText,attr,omitempty"`
	
	RuleDataList string `xml:"RuleDataList,attr,omitempty"`
}

type PriceMessageInfoType struct {
	PriceMessageInfo *ArrayOfMiniRulesPriceMessages `xml:"PriceMessageInfo,omitempty"`
}

type ArrayOfMiniRulesPriceMessages struct {
	MiniRulesPriceMessages []*MiniRulesPriceMessages `xml:"MiniRulesPriceMessages,omitempty"`
}

type MiniRulesPriceMessages struct {
	Text *ArrayOfMiniRulesPriceText `xml:"Text,omitempty"`
}

type ArrayOfMiniRulesPriceText struct {
	MiniRulesPriceText []*MiniRulesPriceText `xml:"MiniRulesPriceText,omitempty"`
}

type MiniRulesPriceText struct {
	PriceMessageValue string `xml:"PriceMessageValue,attr,omitempty"`
	
	PriceDataList string `xml:"PriceDataList,attr,omitempty"`
}

type ArrayOfPNRRemarkInfo struct {
	PNRRemark []*PNRRemarkInfo `xml:"PNRRemark,omitempty"`
}

type PNRRemarkInfo struct {
	RemarkType string `xml:"RemarkType,attr,omitempty"`
	
	RemarkCategory string `xml:"RemarkCategory,attr,omitempty"`
	
	Note string `xml:"Note,attr,omitempty"`
	
	PassengerOrderId int32 `xml:"PassengerOrderId,attr,omitempty"`
	
	SegmentsOrderId string `xml:"SegmentsOrderId,attr,omitempty"`
}

type BookingToBeCancelled struct {
	PNRNo string `xml:"PNRNo,attr,omitempty"`
	
	ProviderType string `xml:"ProviderType,attr,omitempty"`
}

type CustomerPrimaryAdditionalType struct {
	Primary *CustomerPrimaryAdditionalTypePrimary `xml:"Primary,omitempty"`
	
	Additional []*CustomerPrimaryAdditionalTypeAdditional `xml:"Additional,omitempty"`
}

type ArrayOfSpecialRemarkType struct {
	Remark []*SpecialRemarkType `xml:"Remark,omitempty"`
}

type RateQualifierType struct {
	PromoDesc string `xml:"PromoDesc,omitempty"`
	
	RateComments *ArrayOfRateQualifierTypeRateComment `xml:"RateComments,omitempty"`
	
	TravelPurpose string `xml:"TravelPurpose,attr,omitempty"`
	
	RateCategory string `xml:"RateCategory,attr,omitempty"`
	
	CorpDiscountNmbr string `xml:"CorpDiscountNmbr,attr,omitempty"`
	
	RateQualifier string `xml:"RateQualifier,attr,omitempty"`
	
	RatePeriod *RateQualifierTypeRatePeriod `xml:"RatePeriod,attr,omitempty"`
	
	GuaranteedInd bool `xml:"GuaranteedInd,attr,omitempty"`
	
	ArriveByFlight bool `xml:"ArriveByFlight,attr,omitempty"`
	
	RateAuthorizationCode string `xml:"RateAuthorizationCode,attr,omitempty"`
	
	VendorRateID string `xml:"VendorRateID,attr,omitempty"`
}

type ArrayOfRateQualifierTypeRateComment struct {
	RateComment []*RateQualifierTypeRateComment `xml:"RateComment,omitempty"`
}

type ArrayOfMonetaryRuleType struct {
	PaymentRule []*MonetaryRuleType `xml:"PaymentRule,omitempty"`
}

type MonetaryRuleType struct {
	Value string
	
	Amount float64 `xml:"Amount,attr,omitempty"`
	
	RuleType string `xml:"RuleType,attr,omitempty"`
	
	Percent float64 `xml:"Percent,attr,omitempty"`
	
	DateTime time.Time `xml:"DateTime,attr,omitempty"`
	
	PaymentType string `xml:"PaymentType,attr,omitempty"`
}

type OffLocationServiceType struct {
	*OffLocationServiceCoreType
	
	PersonName *PersonNameType `xml:"PersonName,omitempty"`
	
	Telephone *OffLocationServiceTypeTelephone `xml:"Telephone,omitempty"`
	
	SpecInstructions string `xml:"SpecInstructions,attr,omitempty"`
}

type OffLocationServiceCoreType struct {
	Address *OffLocationServiceCoreTypeAddress `xml:"Address,omitempty"`
	
	Type *OffLocationServiceID_Type `xml:"Type,attr,omitempty"`
}

type OffLocationServiceTypeTelephone struct {
}

type ArrayOfFormattedTextType struct {
	VendorMessage []*FormattedTextType `xml:"VendorMessage,omitempty"`
}

type QueueRQ struct {
	QueueNumber string `xml:"QueueNumber,omitempty"`
	
	CategoryNumber string `xml:"CategoryNumber,omitempty"`
}

type QueueRS struct {
	PNRList *PNRListType `xml:"PNRList,omitempty"`
	
	Warnings *WarningsType `xml:"Warnings,omitempty"`
	
	Errors *ErrorsType `xml:"Errors,omitempty"`
	
	Success *SuccessType `xml:"Success,omitempty"`
}

type PNRListType struct {
	PNRListItem []*PNRListItemType `xml:"PNRListItem,omitempty"`
}

type PNRListItemType struct {
	PNRNo string `xml:"PNRNo,omitempty"`
	
	SurName string `xml:"SurName,omitempty"`
	
	FlightNumber string `xml:"FlightNumber,omitempty"`
	
	Carrier string `xml:"Carrier,omitempty"`
	
	From string `xml:"From,omitempty"`
	
	To string `xml:"To,omitempty"`
}

// Location ....
type Location struct {
	LocationCode        string `xml:"LocationCode,attr,omitempty"`
	MultiAirportCityInd bool   `xml:"MultiAirportCityInd,attr,omitempty"`
	// AlternateLocationInd bool `xml:"AlternateLocationInd,attr,omitempty"`
	Value *LocationType
}

// OriginDestinationInformationType ...
type OriginDestinationInformationType struct {
	*TravelDateTimeType
	OriginLocation          *Location                                                   `xml:"OriginLocation,omitempty"`
	DestinationLocation     *Location                                                   `xml:"DestinationLocation,omitempty"`
	RadiusInformation       *RadiusItem                                                 `xml:"RadiusInformation,omitempty"`
	ConnectionLocations     *ArrayOfConnectionTypeConnectionLocation                    `xml:"ConnectionLocations,omitempty"`
	OriginLocationList      *ArrayOfOriginDestinationInformationTypeOriginLocation      `xml:"OriginLocationList,omitempty"`
	DestinationLocationList *ArrayOfOriginDestinationInformationTypeDestinationLocation `xml:"DestinationLocationList,omitempty"`
}

// TravelDateTimeType ...
type TravelDateTimeType struct {
	// ArrivalDateTime *TimeInstantType `xml:"ArrivalDateTime,omitempty"`
	DepartureDateTime string `xml:"DepartureDateTime,omitempty"`
}

// ArrayOfOriginDestinationInformationTypeOriginLocation ...
type ArrayOfOriginDestinationInformationTypeOriginLocation struct {
	OriginLocation []struct {
		*LocationType
		MultiAirportCityInd  bool `xml:"MultiAirportCityInd,attr,omitempty"`
		AlternateLocationInd bool `xml:"AlternateLocationInd,attr,omitempty"`
	} `xml:"OriginLocation,omitempty"`
}

// ArrayOfOriginDestinationInformationTypeDestinationLocation ...
type ArrayOfOriginDestinationInformationTypeDestinationLocation struct {
	DestinationLocation []struct {
		*LocationType
		MultiAirportCityInd  bool `xml:"MultiAirportCityInd,attr,omitempty"`
		AlternateLocationInd bool `xml:"AlternateLocationInd,attr,omitempty"`
	} `xml:"DestinationLocation,omitempty"`
}

// ArrayOfConnectionTypeConnectionLocation ...
type ArrayOfConnectionTypeConnectionLocation struct {
	ConnectionLocation []struct {
		Value               *LocationType
		Inclusive           bool             `xml:"Inclusive,attr,omitempty"`
		PreferLevel         *PreferLevelType `xml:"PreferLevel,attr,omitempty"`
		MinChangeTime       uint             `xml:"MinChangeTime,attr,omitempty"`
		ConnectionInfo      string           `xml:"ConnectionInfo,attr,omitempty"`
		MultiAirportCityInd bool             `xml:"MultiAirportCityInd,attr,omitempty"`
	} `xml:"ConnectionLocation,omitempty"`
}

// RadiusItem ...
type RadiusItem struct {
	Index int32 `xml:"Index,omitempty"`
	
	FromValue int32 `xml:"FromValue,attr,omitempty"`
	
	ToValue int32 `xml:"ToValue,attr,omitempty"`
	
	MeasureUnit string `xml:"MeasureUnit,attr,omitempty"`
}

// SpecificFlightInfoType ...
type SpecificFlightInfoType struct {
	FlightNumber string `xml:"FlightNumber,omitempty"`
	
	Airline []*CompanyNameType `xml:"Airline,omitempty"`
	
	ExcludedAirline []*CompanyNameType `xml:"ExcludedAirline,omitempty"`
	
	Alliance []*AllianceCode `xml:"Alliance,omitempty"`
	
	ExcludedAlliance []*AllianceCode `xml:"ExcludedAlliance,omitempty"`
	
	BookingClassPref []struct {
		ResBookDesigCode string `xml:"ResBookDesigCode,attr,omitempty"`
		
		ResBookDesigCodeType string `xml:"ResBookDesigCodeType,attr,omitempty"`
	} `xml:"BookingClassPref,omitempty"`
}

// AirSearchPrefsType ...
type AirSearchPrefsType struct {
	VendorPref []struct {
		*CompanyNamePrefType
		
		AllianceAllowedInd bool `xml:"AllianceAllowedInd,attr,omitempty"`
		
		LoyaltyAllowedInd bool `xml:"LoyaltyAllowedInd,attr,omitempty"`
	} `xml:"VendorPref,omitempty"`
	
	FlightTypePref []struct {
		FlightType *FlightTypeType `xml:"FlightType,attr,omitempty"`
		
		MaxConnections int `xml:"MaxConnections,attr,omitempty"`
		
		NonScheduledFltInfo string `xml:"NonScheduledFltInfo,attr,omitempty"`
		
		BackhaulIndicator bool `xml:"BackhaulIndicator,attr,omitempty"`
		
		GroundTransportIndicator bool `xml:"GroundTransportIndicator,attr,omitempty"`
		
		DirectAndNonStopOnlyInd bool `xml:"DirectAndNonStopOnlyInd,attr,omitempty"`
		
		NonStopsOnlyInd bool `xml:"NonStopsOnlyInd,attr,omitempty"`
		
		OnlineConnectionsOnlyInd bool `xml:"OnlineConnectionsOnlyInd,attr,omitempty"`
		
		RoutingType string `xml:"RoutingType,attr,omitempty"`
		
		ExcludeTrainInd bool `xml:"ExcludeTrainInd,attr,omitempty"`
	} `xml:"FlightTypePref,omitempty"`
	
	FareRestrictPref []struct {
		AdvResTicketing *AdvResTicketingType `xml:"AdvResTicketing,omitempty"`
		
		StayRestrictions *StayRestrictionsType `xml:"StayRestrictions,omitempty"`
		
		VoluntaryChanges *VoluntaryChangesType `xml:"VoluntaryChanges,omitempty"`
		
		FareRestriction string `xml:"FareRestriction,attr,omitempty"`
		
		Date string `xml:"Date,attr,omitempty"`
		
		FareDisplayCurrency string `xml:"FareDisplayCurrency,attr,omitempty"`
		
		CurrencyOverride string `xml:"CurrencyOverride,attr,omitempty"`
	} `xml:"FareRestrictPref,omitempty"`
	
	EquipPref []*EquipmentTypePref `xml:"EquipPref,omitempty"`
	
	CabinPref []struct {
		Cabin *CabinType `xml:"Cabin,attr,omitempty"`
		
		MajorityCabin bool `xml:"MajorityCabin,attr,omitempty"`
		
		MandatoryCabin bool `xml:"MandatoryCabin,attr,omitempty"`
		
		CabinSubtype string `xml:"CabinSubtype,attr,omitempty"`
	} `xml:"CabinPref,omitempty"`
	
	TicketDistribPref []struct {
		*TicketDistribPrefType
		
		LastTicketDate time.Time `xml:"LastTicketDate,attr,omitempty"`
		
		FirstTicketDate time.Time `xml:"FirstTicketDate,attr,omitempty"`
	} `xml:"TicketDistribPref,omitempty"`
	
	SmokingAllowed bool `xml:"SmokingAllowed,attr,omitempty"`
	
	OnTimeRate float64 `xml:"OnTimeRate,attr,omitempty"`
	
	ETicketDesired bool `xml:"ETicketDesired,attr,omitempty"`
	
	MaxStopsQuantity int32 `xml:"MaxStopsQuantity,attr,omitempty"`
	
	Start string `xml:"Start,attr,omitempty"`
	
	Duration string `xml:"Duration,attr,omitempty"`
	
	End string `xml:"End,attr,omitempty"`
}

// AirlineDiversityType ...
type AirlineDiversityType struct {
	DiversityType []*AirlineDiversityEnum `xml:"DiversityType,omitempty"`
}

// ExpandedParametersType ...
type ExpandedParametersType struct {
	Param []*ExpandedParameterEnum `xml:"Param,omitempty"`
}

// PricedItinerariesType ...
type PricedItinerariesType struct {
	PricedItinerary []*PricedItinerary             `xml:"PricedItinerary,omitempty"`
	FreeBaggages    *ArrayOfFreeBaggageSegmentItem `xml:"FreeBaggages"`
}

// PricedItinerary ...
type PricedItinerary struct {
	*PricedItineraryType
}

// PricedItineraryType ...
type PricedItineraryType struct {
	SequenceNumber          string                       `xml:"SequenceNumber,attr,omitempty"`
	IsOneWayCombinable      bool                         `xml:"IsOneWayCombinable,attr,omitempty"`
	Currency                string                       `xml:"Currency,attr,omitempty"`
	ProviderType            string                       `xml:"ProviderType,attr,omitempty"`
	TicketingInfo           *TicketingInfo               `xml:"TicketingInfo,omitempty"`
	Notes                   []*FreeTextType              `xml:"Notes,omitempty"`
	AirItinerary            *AirItineraryType            `xml:"AirItinerary,omitempty"`
	AirItineraryPricingInfo *AirItineraryPricingInfoType `xml:"AirItineraryPricingInfo,omitempty"`
}

// AirItineraryType ...
type AirItineraryType struct {
	CalculateSF                   bool                                            `xml:"CalculateSF,attr,omitempty"`
	ValidatingAirlineCode         string                                          `xml:"ValidatingAirlineCode,attr,omitempty"`
	DirectionInd                  string                                          `xml:"DirectionInd,attr,omitempty"`
	OriginDestinationCombinations *ArrayOfOriginDestinationCombinationType        `xml:"OriginDestinationCombinations,omitempty"`
	OriginDestinationOptions      *ArrayOfAirItineraryTypeOriginDestinationOption `xml:"OriginDestinationOptions,omitempty"`
}

// DeliveryInfo ...
type DeliveryInfo struct {
	DistributionType string  `xml:"DistribType,attr,omitempty"`
	Amount           float64 `xml:"Amount,attr,omitempty"`
}

// TicketingInfo ...
type TicketingInfo struct {
	*TicketingInfoRSType
	DeliveryInfo []*DeliveryInfo `xml:"DeliveryInfo,omitempty"`
	PaymentType  []*PaymentTypes `xml:"PaymentType,attr,omitempty"`
}

// Baggage ...
type Baggage struct {
	// XMLName xml.Name `xml:"Baggage"`
	Index    int32  `xml:"Index,attr,omitempty"`
	Quantity string `xml:"Quantity,attr,omitempty"`
	Unit     string `xml:"Unit,attr,omitempty"`
}

// ConnectedOperatingAirline ...
type ConnectedOperatingAirline struct {
	Code string `xml:"Code,attr,omitempty"`
	Text string `xml:"Text,attr,omitempty"`
}

// ServiceFeeInfoRQ ...
type ServiceFeeInfoRQ struct {
	*GenericFlightRQ
	Type       string                      `xml:"Type,attr,omitempty"`
	FlightInfo *ServiceFeeInfoRQFlightInfo `xml:"FlightInfo,omitempty"`
}

// GenericFlightRQ ...
type GenericFlightRQ struct {
	IsOneWayCombinable  bool   `xml:"IsOneWayCombinable,attr,omitempty"`
	RecommendationID    string `xml:"RecommendationID,attr,omitempty"`
	CombinationID       string `xml:"CombinationID,attr,omitempty"`
	FareFamilyPackageID string `xml:"FareFamilyPackageID,attr,omitempty"`
}

// ServiceFeeInfoRQFlightInfo ...
type ServiceFeeInfoRQFlightInfo struct {
	RecommendationID   int32 `xml:"RecommendationID,attr,omitempty"`
	CombinationID      int32 `xml:"CombinationID,attr,omitempty"`
	IsOneWayCombinable bool  `xml:"IsOneWayCombinable,attr,omitempty"`
}

// ServiceFeeInfoRS ...
type ServiceFeeInfoRS struct {
	Warnings    *WarningsType    `xml:"Warnings,omitempty"`
	Errors      *ErrorsType      `xml:"Errors,omitempty"`
	ServiceFees *ServiceFeesType `xml:"ServiceFees,omitempty"`
	Success     *SuccessType     `xml:"Success,omitempty"`
}

// BaggageInfoRQ ...
type BaggageInfoRQ struct {
	*ServiceFeeInfoRQ
	Passenger []*PassengerInfo `xml:"Passenger,omitempty"`
}

// PassengerInfo ...
type PassengerInfo struct {
	FrequentFlyerCard []*FrequentFlyerCardType `xml:"FrequentFlyerCard,omitempty"`
	PassengerType     string                   `xml:"PassengerType,attr,omitempty"`
	PassengerIndex    int32                    `xml:"PassengerIndex,attr,omitempty"`
}

// FrequentFlyerCardType ...
type FrequentFlyerCardType struct {
	AirlineCode string `xml:"AirlineCode,attr,omitempty"`
	CardNumber  string `xml:"CardNumber,attr,omitempty"`
}

// BaggageInfoRS ...
type BaggageInfoRS struct {
	AncillaryCatalogueItems *AncillaryCatalogueItemsType `xml:"AncillaryCatalogueItems,omitempty"`
	Errors                  *ErrorsType                  `xml:"Errors,omitempty"`
	FreeBaggageAllowances   *FreeBaggageAllowancesTypes  `xml:"FreeBaggageAllowances,omitempty"`
	Warnings                *WarningsType                `xml:"Warnings,omitempty"`
	Success                 *SuccessType                 `xml:"Success,omitempty"`
}

// AncillaryCatalogueItemsType ...
type AncillaryCatalogueItemsType struct {
	Errors                 *ErrorsType                   `xml:"Errors,omitempty"`
	AncillaryCatalogueItem []*AncillaryCatalogueItemType `xml:"AncillaryCatalogueItem,omitempty"`
}

// ArrayOfAncillaryCatalogueItemType ...
type ArrayOfAncillaryCatalogueItemType struct {
	AncillaryCatalogueItemType []*AncillaryCatalogueItemType `xml:"AncillaryCatalogueItemType,omitempty"`
}

// AncillaryCatalogueItemType ...
type AncillaryCatalogueItemType struct {
	SSRCode        string              `xml:"SSRCode,omitempty"`
	AirlineCode    string              `xml:"AirlineCode,omitempty"`
	AncillaryFares *AncillaryFaresType `xml:"AncillaryFares,omitempty"`
}

// AncillaryFaresType ...
type AncillaryFaresType struct {
	AncillaryFare []*AncillaryFareType `xml:"AncillaryFare,omitempty"`
}

// AncillaryFareType ...
type AncillaryFareType struct {
	Amount           float64 `xml:"Amount,omitempty"`
	Tax              float64 `xml:"Tax,omitempty"`
	Currency         string  `xml:"Currency,omitempty"`
	SegmentReference string  `xml:"SegmentReference,omitempty"`
	SegmentInfo      string  `xml:"SegmentInfo,omitempty"`
	PassengerType    string  `xml:"PassengerType,omitempty"`
	XMLPassengerID   int32   `xml:"XMLPassengerID,omitempty"`
	Description      string  `xml:"Description,omitempty"`
}

// AncillaryInfoRQ ...
type AncillaryInfoRQ struct {
	*BaggageInfoRQ
}

// AncillaryInfoRS ...
type AncillaryInfoRS struct {
	Success                 *SuccessType                 `xml:"Success,omitempty"`
	Warnings                *WarningsType                `xml:"Warnings,omitempty"`
	Errors                  *ErrorsType                  `xml:"Errors,omitempty"`
	AncillaryCatalogueItems *AncillaryCatalogueItemsType `xml:"AncillaryCatalogueItems,omitempty"`
}

// PriceInfoRQ ...
type PriceInfoRQ struct {
	*ServiceFeeInfoRQ
}

// RequiredParametersRQ ...
type RequiredParametersRQ struct {
	RecommendationID   int32  `xml:"RecommendationID,attr,omitempty"`
	CombinationID      int32  `xml:"CombinationID,attr,omitempty"`
	IsOneWayCombinable bool   `xml:"IsOneWayCombinable,attr,omitempty"`
	CCType             string `xml:"CCType,attr,omitempty"`
}

// ArrayOfLCCRequiredParameter ...
type ArrayOfLCCRequiredParameter struct {
	LCCRequiredParameter []*LCCRequiredParameter `xml:"LCCRequiredParameter,omitempty"`
}

// LCCRequiredParameter ...
type LCCRequiredParameter struct {
	Enabled      bool   `xml:"Enabled,omitempty"`
	Mandatory    bool   `xml:"Mandatory,omitempty"`
	Info         string `xml:"Info,omitempty"`
	PerPassenger bool   `xml:"PerPassenger,omitempty"`
}

// ArrayOfLCCSupportedCard ...
type ArrayOfLCCSupportedCard struct {
	LCCSupportedCard []*LCCSupportedCard `xml:"LCCSupportedCard,omitempty"`
}

// LCCSupportedCard ...
type LCCSupportedCard struct {
	Code     string  `xml:"Code,attr,omitempty"`
	Amount   float64 `xml:"Amount,attr,omitempty"`
	Currency string  `xml:"Currency,attr,omitempty"`
}

// ArrayOfLCCBaggageFee ...
type ArrayOfLCCBaggageFee struct {
	LCCBaggageFee []*LCCBaggageFee `xml:"LCCBaggageFee,omitempty"`
}

// LCCBaggageFee ...
type LCCBaggageFee struct {
	Amount    float64 `xml:"Amount,attr,omitempty"`
	Currency  string  `xml:"Currency,attr,omitempty"`
	Weight    string  `xml:"Weight,attr,omitempty"`
	MaxWeight string  `xml:"MaxWeight,attr,omitempty"`
	Quantity  string  `xml:"Quantity,attr,omitempty"`
}

// OTAAirCheckETicketRQ ...
type OTAAirCheckETicketRQ struct {
	*GenericFlightRQ
}

// OTAAirCheckETicketRS ...
type OTAAirCheckETicketRS struct {
	Errors  *ErrorsType  `xml:"Errors,omitempty"`
	Success *SuccessType `xml:"Success,omitempty"`
}

// OTAAirRulesRQ ...
type OTAAirRulesRQ struct {
	*GenericFlightRQ
	SequenceNmbr            uint                                 `xml:"SequenceNmbr,attr,omitempty"`
	RetransmissionIndicator bool                                 `xml:"RetransmissionIndicator,attr,omitempty"`
	MiniRuleEnabled         int32                                `xml:"MiniRuleEnabled,attr,omitempty"`
	PriceMessageEnabled     int32                                `xml:"PriceMessageEnabled,attr,omitempty"`
	FlightRuleEnabled       int32                                `xml:"FlightRuleEnabled,attr,omitempty"`
	PassengerType           string                               `xml:"PassengerType,attr,omitempty"`
	EchoToken               string                               `xml:"EchoToken,attr,omitempty"`
	TimeStamp               string                               `xml:"TimeStamp,attr,omitempty"`
	Version                 float64                              `xml:"Version,attr,omitempty"`
	TransactionIdentifier   string                               `xml:"TransactionIdentifier,attr,omitempty"`
	POS                     *ArrayOfSourceType                   `xml:"POS,omitempty"`
	Target                  *OTA_AirRulesRQTarget                `xml:"Target,attr,omitempty"`
	RuleReqInfo             *OTAAirRulesRQRuleReqInfo            `xml:"RuleReqInfo,omitempty"`
	TransactionStatusCode   *OTA_AirRulesRQTransactionStatusCode `xml:"TransactionStatusCode,attr,omitempty"`
}

// OTAAirRulesRQRuleReqInfo ...
type OTAAirRulesRQRuleReqInfo struct {
	*FareInfoType
	SubSection        []*OTAAirRulesRQRuleReqInfoSubSection `xml:"SubSection,omitempty"`
	LanguageRequested string                                `xml:"LanguageRequested,attr,omitempty"`
}

// OTAAirRulesRQRuleReqInfoSubSection ...
type OTAAirRulesRQRuleReqInfoSubSection struct {
	SubTitle string `xml:"SubTitle,attr,omitempty"`
	
	SubCode string `xml:"SubCode,attr,omitempty"`
	
	SubSectionNumber string `xml:"SubSectionNumber,attr,omitempty"`
}

// OTAAirRulesRS ...
type OTAAirRulesRS struct {
	FareRuleResponseInfo *FareRuleResponseInfoType `xml:"FareRuleResponseInfo,omitempty"`
	Success              *SuccessType              `xml:"Success,omitempty"`
	Errors               *ErrorsType               `xml:"Errors,omitempty"`
	Warnings             *WarningsType             `xml:"Warnings,omitempty"`
	RuleLink             *RuleLink                 `xml:"RuleLink,omitempty"`
	MiniRuleResponseInfo *MiniRuleResponseInfoType `xml:"MiniRuleResponseInfo,omitempty"`
	PriceMessageInfoType *PriceMessageInfoType     `xml:"PriceMessageInfoType,omitempty"`
	// EchoToken               string                               `xml:"EchoToken,attr,omitempty"`
	TimeStamp string `xml:"TimeStamp,attr,omitempty"`
	// Target                  *OTA_AirRulesRSTarget                `xml:"Target,attr,omitempty"`
	Version float64 `xml:"Version,attr,omitempty"`
	// TransactionIdentifier   string                               `xml:"TransactionIdentifier,attr,omitempty"`
	// SequenceNmbr            int                                  `xml:"SequenceNmbr,attr,omitempty"`
	// TransactionStatusCode   *OTA_AirRulesRSTransactionStatusCode `xml:"TransactionStatusCode,attr,omitempty"`
	// RetransmissionIndicator bool                                 `xml:"RetransmissionIndicator,attr,omitempty"`
}

// RuleLink ...
type RuleLink struct {
	Link string `xml:"Link,omitempty"`
	
	ItineraryIndex int32 `xml:"ItineraryIndex,attr,omitempty"`
}

// InsurancePlan ...
type InsurancePlan struct {
	PlanID int32 `xml:"PlanID,attr,omitempty"`
}

// RelatedMember ...
type RelatedMember struct {
	MemberID string `xml:"MemberID,attr,omitempty"`
	
	GuestMemberID string `xml:"GuestMemberID,attr,omitempty"`
}

// PassengerDetailChanges ...
type PassengerDetailChanges struct {
	ChangeDetails *ArrayOfChangeDetail `xml:"ChangeDetails,omitempty"`
}

// ArrayOfChangeDetail ...
type ArrayOfChangeDetail struct {
	ChangeDetail []*ChangeDetail `xml:"ChangeDetail,omitempty"`
}

// ChangeDetail ...
type ChangeDetail struct {
	SpecialRequests *ArrayOfSpecialRequestChange `xml:"SpecialRequests,omitempty"`
	
	PersonName *PersonNameType `xml:"PersonName,omitempty"`
	
	Document struct {
		*DocumentType
		
		Operation *ActionType `xml:"Operation,attr,omitempty"`
	} `xml:"Document,omitempty"`
	
	ChangeDetailAction *ChangeDetailActionTypes `xml:"ChangeDetailAction,attr,omitempty"`
	
	SegmentIndex int32 `xml:"SegmentIndex,attr,omitempty"`
}

// ArrayOfSpecialRequestChange ...
type ArrayOfSpecialRequestChange struct {
	SpecialRequestChange []*SpecialRequestChange `xml:"SpecialRequestChange,omitempty"`
}

// SpecialRequestChange ...
type SpecialRequestChange struct {
	SpecialReqDetails []*SpecialReqDetailsType `xml:"SpecialReqDetails,omitempty"`
}

// ArrayOfPNRRemarkChange ...
type ArrayOfPNRRemarkChange struct {
	PNRRemarkChange []*PNRRemarkChange `xml:"PNRRemarkChange,omitempty"`
}

// PNRRemarkChange ...
type PNRRemarkChange struct {
	PNRRemark *PNRRemarkInfo `xml:"PNRRemark,omitempty"`
	
	PersonName *PersonNameType `xml:"PersonName,omitempty"`
	
	ChangeDetailAction *ChangeDetailActionTypes `xml:"ChangeDetailAction,attr,omitempty"`
	
	SegmentIndex int32 `xml:"SegmentIndex,attr,omitempty"`
}

// ArrayOfOSIElementChange ...
type ArrayOfOSIElementChange struct {
	OSIElementChange []*OSIElementChange `xml:"OSIElementChange,omitempty"`
}

// OSIElementChange ...
type OSIElementChange struct {
	PersonName *PersonNameType `xml:"PersonName,omitempty"`
	
	OSIMessage string `xml:"OSIMessage,omitempty"`
	
	ChangeDetailAction *ChangeDetailActionTypes `xml:"ChangeDetailAction,attr,omitempty"`
}

// ArrayOfLastTicketingDateData ...
type ArrayOfLastTicketingDateData struct {
	LastTicketingDateData []struct {
		PnrNo string `xml:"PnrNo,attr,omitempty"`
		
		OldLastTicketingDate time.Time `xml:"OldLastTicketingDate,attr,omitempty"`
		
		LastTicketingDate time.Time `xml:"LastTicketingDate,attr,omitempty"`
	} `xml:"LastTicketingDateData,omitempty"`
}

// ArrayOfValidationError ...
type ArrayOfValidationError struct {
	ValidationError []*ValidationError `xml:"ValidationError,omitempty"`
}

// ValidationError ...
type ValidationError struct {
	ErrorMessage              string                     `xml:"ErrorMessage,omitempty"`
	MessageValue              string                     `xml:"MessageValue,omitempty"`
	MainMessage               string                     `xml:"MainMessage,omitempty"`
	ErrorField                string                     `xml:"ErrorField,omitempty"`
	GotoPage                  string                     `xml:"GotoPage,omitempty"`
	RelatedEntityPropertyName string                     `xml:"RelatedEntityPropertyName,omitempty"`
	MessageValueList          []string                   `xml:"MessageValueList,omitempty"`
	ErrorCode                 *ValidationErrorCodes      `xml:"ErrorCode,omitempty"`
	ErrorType                 *ValidationErrorTypes      `xml:"ErrorType,omitempty"`
	ErrorCategory             *ValidationErrorCategories `xml:"ErrorCategory,omitempty"`
}

// ArrayOfCorporateSetting ...
type ArrayOfCorporateSetting struct {
	CorporateSetting []*CorporateSetting `xml:"CorporateSetting,omitempty"`
}

// CorporateSetting ...
type CorporateSetting struct {
	FlightProviderType                        string                `xml:"FlightProviderType,omitempty"`
	RailProviderType                          *RailProviderTypeEnum `xml:"RailProviderType,omitempty"`
	DisplayRule                               *DisplayRuleEnum      `xml:"DisplayRule,omitempty"`
	TripApprovalRule                          *TripApprovalRuleEnum `xml:"TripApprovalRule,omitempty"`
	RequireJustificationRule                  bool                  `xml:"RequireJustificationRule,omitempty"`
	RequireJustificationOnlyFlightRule        bool                  `xml:"RequireJustificationOnlyFlightRule,omitempty"`
	RequireJusificationForDeclineProcess      bool                  `xml:"RequireJusificationForDeclineProcess,omitempty"`
	ApplyApprovalRules                        string                `xml:"ApplyApprovalRules,omitempty"`
	AllowBookingOnlyArrangerOrSelectedMembers bool                  `xml:"AllowBookingOnlyArrangerOrSelectedMembers,omitempty"`
	AllowETicketForNext                       int32                 `xml:"AllowETicketForNext,omitempty"`
}

// ArrayOfQueueSetting ...
type ArrayOfQueueSetting struct {
	QueueSetting []*QueueSetting `xml:"QueueSetting,omitempty"`
}

// QueueSetting ...
type QueueSetting struct {
	FlightProviderType string                `xml:"FlightProviderType,omitempty"`
	RailProviderType   *RailProviderTypeEnum `xml:"RailProviderType,omitempty"`
	DefaultQue         *QueSetting           `xml:"DefaultQue,omitempty"`
	CancellationQue    *QueSetting           `xml:"CancellationQue,omitempty"`
	PendingApproval    *QueSetting           `xml:"PendingApproval,omitempty"`
	ApprovedApproval   *QueSetting           `xml:"ApprovedApproval,omitempty"`
	DeclinedApproval   *QueSetting           `xml:"DeclinedApproval,omitempty"`
	TicketedQue        *QueSetting           `xml:"TicketedQue,omitempty"`
}

// QueSetting ...
type QueSetting struct {
	QueNo         int32  `xml:"QueNo,omitempty"`
	QueCategoryNo int32  `xml:"QueCategoryNo,omitempty"`
	OfficeID      string `xml:"OfficeID,omitempty"`
}

// APISRulesRQ ...
type APISRulesRQ struct {
	RecommendationID int32 `xml:"RecommendationID,attr,omitempty"`
	CombinationID    int32 `xml:"CombinationID,attr,omitempty"`
}

// APISRulesRS ...
type APISRulesRS struct {
	Errors    *ErrorsType   `xml:"Errors,omitempty"`
	APISRules *APISRules    `xml:"APISRules,omitempty"`
	Success   *SuccessType  `xml:"Success,omitempty"`
	Warnings  *WarningsType `xml:"Warnings,omitempty"`
}

// APISRules ...
type APISRules struct {
	APISRuleDocs *APISRuleDocs `xml:"APISRuleDocs,omitempty"`
	APISRuleDoco *APISRuleDoco `xml:"APISRuleDoco,omitempty"`
	APISRuleDoca *APISRuleDoca `xml:"APISRuleDoca,omitempty"`
}

// APISRuleDocs ...
type APISRuleDocs struct {
	MiddleName string `xml:"MiddleName,omitempty"`
	
	DocumentType string `xml:"DocumentType,omitempty"`
	
	DocumentNumber string `xml:"DocumentNumber,omitempty"`
	
	CountryOfIssuance string `xml:"CountryOfIssuance,omitempty"`
	
	Nationality string `xml:"Nationality,omitempty"`
	
	ExpiryDate string `xml:"ExpiryDate,omitempty"`
	
	DocumentTypeMandatory string `xml:"DocumentTypeMandatory,omitempty"`
}

// APISRuleDoco ...
type APISRuleDoco struct {
	VisaNumber string `xml:"VisaNumber,omitempty"`
	
	PlaceOfBirth string `xml:"PlaceOfBirth,omitempty"`
	
	DateOfIssuance string `xml:"DateOfIssuance,omitempty"`
	
	PlaceOfIssuance string `xml:"PlaceOfIssuance,omitempty"`
	
	CountryRequiringVisa string `xml:"CountryRequiringVisa,omitempty"`
}

// APISRuleDoca ...
type APISRuleDoca struct {
	AddressType string `xml:"AddressType,omitempty"`
	
	CountryCode string `xml:"CountryCode,omitempty"`
	
	AddressDetails string `xml:"AddressDetails,omitempty"`
	
	CityCode string `xml:"CityCode,omitempty"`
	
	State string `xml:"State,omitempty"`
	
	ZipPostalCode string `xml:"ZipPostalCode,omitempty"`
}

// FareFamilyInfoRQ ...
type FareFamilyInfoRQ struct {
	*GenericFlightRQ
	
	TravelerInfo *TravelerInfoType `xml:"TravelerInfo,omitempty"`
}

// FareFamilyInfoRS ...
type FareFamilyInfoRS struct {
	Warnings *WarningsType `xml:"Warnings,omitempty"`
	
	Errors *ErrorsType `xml:"Errors,omitempty"`
	
	FareFamilies *ArrayOfFareFamily `xml:"FareFamilies,omitempty"`
	
	Success *SuccessType `xml:"Success,omitempty"`
}

// ArrayOfFareFamily ...
type ArrayOfFareFamily struct {
	FareFamily []*FareFamily `xml:"FareFamily,omitempty"`
}

// FareFamily ...
type FareFamily struct {
	Errors *ErrorsType `xml:"Errors,omitempty"`
	
	FareFamilyDescriptions *ArrayOfFareFamilyDescription `xml:"FareFamilyDescriptions,omitempty"`
	
	FlightInfo *FlightInfo `xml:"FlightInfo,omitempty"`
	
	FareInfo *FareInfo `xml:"FareInfo,omitempty"`
	
	PackageId string `xml:"PackageId,attr,omitempty"`
	
	FareFamilyName string `xml:"FareFamilyName,attr,omitempty"`
	
	FareBasisCode string `xml:"FareBasisCode,attr,omitempty"`
	
	Class string `xml:"Class,attr,omitempty"`
}

// ArrayOfFareFamilyDescription ...
type ArrayOfFareFamilyDescription struct {
	FareFamilyDescription []*FareFamilyDescription `xml:"FareFamilyDescription,omitempty"`
}

// FareFamilyDescription ...
type FareFamilyDescription struct {
	FamilyDescName string `xml:"FamilyDescName,attr,omitempty"`
	
	FamilyTypeText string `xml:"FamilyTypeText,attr,omitempty"`
	
	PaymentType string `xml:"PaymentType,attr,omitempty"`
}

// FlightInfo ...
type FlightInfo struct {
	Airline string `xml:"Airline,attr,omitempty"`
	
	Routes string `xml:"Routes,attr,omitempty"`
}

// FareInfo ...
type FareInfo struct {
	ExtraFare  string `xml:"ExtraFare,attr,omitempty"`
	ServiceFee string `xml:"ServiceFee,attr,omitempty"`
	Currency   string `xml:"Currency,attr,omitempty"`
}

// Guid ...
type Guid string
