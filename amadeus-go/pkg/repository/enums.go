package repository

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// Language ...
type Language string

// PreferLevelType ...
type PreferLevelType string

const (
	// PreferLevelTypeOnly ..
	PreferLevelTypeOnly PreferLevelType = "Only"
	
	// PreferLevelTypeUnacceptable ...
	PreferLevelTypeUnacceptable PreferLevelType = "Unacceptable"
	
	// PreferLevelTypePreferred ...
	PreferLevelTypePreferred PreferLevelType = "Preferred"
	
	// PreferLevelTypeRequired ...
	PreferLevelTypeRequired PreferLevelType = "Required"
	
	// PreferLevelTypeNoPreference ...
	PreferLevelTypeNoPreference PreferLevelType = "NoPreference"
)

// TransactionActionType ...
type TransactionActionType string

const (
	// TransactionActionTypeBook ...
	TransactionActionTypeBook TransactionActionType = "Book"
	
	// TransactionActionTypeQuote ...
	TransactionActionTypeQuote TransactionActionType = "Quote"
	
	// TransactionActionTypeHold ...
	TransactionActionTypeHold TransactionActionType = "Hold"
	
	// TransactionActionTypeInitiate ..
	TransactionActionTypeInitiate TransactionActionType = "Initiate"
	
	// TransactionActionTypeIgnore ...
	TransactionActionTypeIgnore TransactionActionType = "Ignore"
	
	// TransactionActionTypeModify ...
	TransactionActionTypeModify TransactionActionType = "Modify"
	
	// TransactionActionTypeCommit ...
	TransactionActionTypeCommit TransactionActionType = "Commit"
	
	// TransactionActionTypeCancel ...
	TransactionActionTypeCancel TransactionActionType = "Cancel"
	
	// TransactionActionTypeCommitOverrideEdits ...
	TransactionActionTypeCommitOverrideEdits TransactionActionType = "CommitOverrideEdits"
)

type CoverageTextType string

const (
	CoverageTextTypeSupplement CoverageTextType = "Supplement"
	
	CoverageTextTypeDescription CoverageTextType = "Description"
	
	CoverageTextTypeLimits CoverageTextType = "Limits"
)

// PaymentTypes ...
type PaymentTypes string

const (
	// PaymentTypesNone ...
	PaymentTypesNone PaymentTypes = "None"
	
	// PaymentTypesCallMe ...
	PaymentTypesCallMe PaymentTypes = "CallMe"
	
	// PaymentTypesCreditCard ...
	PaymentTypesCreditCard PaymentTypes = "CreditCard"
	
	// PaymentTypesInvoice ...
	PaymentTypesInvoice PaymentTypes = "Invoice"
	
	// PaymentTypesPayAtAgency ...
	PaymentTypesPayAtAgency PaymentTypes = "PayAtAgency"
	
	// PaymentTypesC ...
	PaymentTypesC PaymentTypes = "C"
	
	// PaymentTypesCC ...
	PaymentTypesCC PaymentTypes = "CC"
)

// PaymentFPTypes ...
type PaymentFPTypes string

const (
	// PaymentFPTypesFPCA ...
	PaymentFPTypesFPCA PaymentFPTypes = "FPCA"
	
	// PaymentFPTypesFPCC ...
	PaymentFPTypesFPCC PaymentFPTypes = "FPCC"
)

// TimeUnitType ...
type TimeUnitType string

const (
	// TimeUnitTypeYear ...
	TimeUnitTypeYear TimeUnitType = "Year"
	
	// TimeUnitTypeMonth ...
	TimeUnitTypeMonth TimeUnitType = "Month"
	
	// TimeUnitTypeWeek ...
	TimeUnitTypeWeek TimeUnitType = "Week"
	
	// TimeUnitTypeDay ...
	TimeUnitTypeDay TimeUnitType = "Day"
	
	// TimeUnitTypeHour ...
	TimeUnitTypeHour TimeUnitType = "Hour"
	
	// TimeUnitTypeSecond ...
	TimeUnitTypeSecond TimeUnitType = "Second"
	
	// TimeUnitTypeFullDuration ...
	TimeUnitTypeFullDuration TimeUnitType = "FullDuration"
)

// AmountDeterminationType ...
type AmountDeterminationType string

const (
	// AmountDeterminationTypeInclusive ...
	AmountDeterminationTypeInclusive AmountDeterminationType = "Inclusive"
	
	// AmountDeterminationTypeExclusive ...
	AmountDeterminationTypeExclusive AmountDeterminationType = "Exclusive"
	
	// AmountDeterminationTypeCumulative ...
	AmountDeterminationTypeCumulative AmountDeterminationType = "Cumulative"
)

// AvailabilityStatusType ...
type AvailabilityStatusType string

const (
	// AvailabilityStatusTypeOpen ...
	AvailabilityStatusTypeOpen AvailabilityStatusType = "Open"
	
	// AvailabilityStatusTypeClose ...
	AvailabilityStatusTypeClose AvailabilityStatusType = "Close"
	
	// AvailabilityStatusTypeClosedOnArrival ...
	AvailabilityStatusTypeClosedOnArrival AvailabilityStatusType = "ClosedOnArrival"
	
	// AvailabilityStatusTypeClosedOnArrivalOnRequest ...
	AvailabilityStatusTypeClosedOnArrivalOnRequest AvailabilityStatusType = "ClosedOnArrivalOnRequest"
	
	// AvailabilityStatusTypeOnRequest ...
	AvailabilityStatusTypeOnRequest AvailabilityStatusType = "OnRequest"
)

// RateIndicatorType ...
type RateIndicatorType string

const (
	// RateIndicatorTypeChangeDuringStay ...
	RateIndicatorTypeChangeDuringStay RateIndicatorType = "ChangeDuringStay"
	
	// RateIndicatorTypeMultipleNights ...
	RateIndicatorTypeMultipleNights RateIndicatorType = "MultipleNights"
	
	// RateIndicatorTypeExclusive ...
	RateIndicatorTypeExclusive RateIndicatorType = "Exclusive"
	
	// RateIndicatorTypeOnRequest ...
	RateIndicatorTypeOnRequest RateIndicatorType = "OnRequest"
	
	// RateIndicatorTypeLimitedAvailability ...
	RateIndicatorTypeLimitedAvailability RateIndicatorType = "LimitedAvailability"
	
	// RateIndicatorTypeAvailableForSale ...
	RateIndicatorTypeAvailableForSale RateIndicatorType = "AvailableForSale"
	
	// RateIndicatorTypeClosedOut ...
	RateIndicatorTypeClosedOut RateIndicatorType = "ClosedOut"
	
	// RateIndicatorTypeOtherAvailable ...
	RateIndicatorTypeOtherAvailable RateIndicatorType = "OtherAvailable"
	
	// RateIndicatorTypeUnableToProcess ...
	RateIndicatorTypeUnableToProcess RateIndicatorType = "UnableToProcess"
	
	// RateIndicatorTypeNoAvailability ...
	RateIndicatorTypeNoAvailability RateIndicatorType = "NoAvailability"
	
	// RateIndicatorTypeRoomTypeClosed ...
	RateIndicatorTypeRoomTypeClosed RateIndicatorType = "RoomTypeClosed"
	
	// RateIndicatorTypeRatePlanClosed ...
	RateIndicatorTypeRatePlanClosed RateIndicatorType = "RatePlanClosed"
	
	// RateIndicatorTypeLOS_Restricted ...
	RateIndicatorTypeLOS_Restricted RateIndicatorType = "LOS_Restricted"
	
	// RateIndicatorTypeRestricted ...
	RateIndicatorTypeRestricted RateIndicatorType = "Restricted"
	
	// RateIndicatorTypeDoesNotExist ...
	RateIndicatorTypeDoesNotExist RateIndicatorType = "DoesNotExist"
)

// DayOfWeekType ...
type DayOfWeekType string

const (
	// DayOfWeekTypeMon ...
	DayOfWeekTypeMon DayOfWeekType = "Mon"
	
	// DayOfWeekTypeTue ...
	DayOfWeekTypeTue DayOfWeekType = "Tue"
	
	// DayOfWeekTypeWed ...
	DayOfWeekTypeWed DayOfWeekType = "Wed"
	
	// DayOfWeekTypeThu ...
	DayOfWeekTypeThu DayOfWeekType = "Thu"
	
	// DayOfWeekTypeFri ...
	DayOfWeekTypeFri DayOfWeekType = "Fri"
	
	// DayOfWeekTypeSat ...
	DayOfWeekTypeSat DayOfWeekType = "Sat"
	
	// DayOfWeekTypeSun ...
	DayOfWeekTypeSun DayOfWeekType = "Sun"
)

// PricingType ...
type PricingType string

const (
	// PricingTypePerstay ...
	PricingTypePerstay PricingType = "Per stay"
	
	// PricingTypePerperson ...
	PricingTypePerperson PricingType = "Per person"
	
	// PricingTypePernight ...
	PricingTypePernight PricingType = "Per night"
	
	// PricingTypePerpersonpernight ...
	PricingTypePerpersonpernight PricingType = "Per person per night"
	
	// PricingTypePeruse ...
	PricingTypePeruse PricingType = "Per use"
)

// DistanceUnitNameType ...
type DistanceUnitNameType string

const (
	// DistanceUnitNameTypeMile ...
	DistanceUnitNameTypeMile DistanceUnitNameType = "Mile"
	
	// DistanceUnitNameTypeKm ...
	DistanceUnitNameTypeKm DistanceUnitNameType = "Km"
	
	// DistanceUnitNameTypeBlock ....
	DistanceUnitNameTypeBlock DistanceUnitNameType = "Block"
)

// CustomerPrimaryAdditionalTypeAdditionalQualificationMethod ...
type CustomerPrimaryAdditionalTypeAdditionalQualificationMethod string

const (
	// CustomerPrimaryAdditionalTypeAdditionalQualificationMethodRT_AirlineTicket ...
	CustomerPrimaryAdditionalTypeAdditionalQualificationMethodRT_AirlineTicket CustomerPrimaryAdditionalTypeAdditionalQualificationMethod = "RT_AirlineTicket"
	
	// CustomerPrimaryAdditionalTypeAdditionalQualificationMethodCreditCard ...
	CustomerPrimaryAdditionalTypeAdditionalQualificationMethodCreditCard CustomerPrimaryAdditionalTypeAdditionalQualificationMethod = "CreditCard"
	
	// CustomerPrimaryAdditionalTypeAdditionalQualificationMethodPassportAndReturnTkt ...
	CustomerPrimaryAdditionalTypeAdditionalQualificationMethodPassportAndReturnTkt CustomerPrimaryAdditionalTypeAdditionalQualificationMethod = "PassportAndReturnTkt"
)

// MealType ...
type MealType string

const (
	// MealTypeAVML ...
	MealTypeAVML MealType = "AVML"
	
	// MealTypeBBML ....
	MealTypeBBML MealType = "BBML"
	
	// MealTypeBLML ...
	MealTypeBLML MealType = "BLML"
	
	// MealTypeCHML ...
	MealTypeCHML MealType = "CHML"
	
	// MealTypeDBML ...
	MealTypeDBML MealType = "DBML"
	
	// MealTypeFPML ...
	MealTypeFPML MealType = "FPML"
	
	// MealTypeGFML ...
	MealTypeGFML MealType = "GFML"
	
	// MealTypeHFML ...
	MealTypeHFML MealType = "HFML"
	
	// MealTypeHNML ...
	MealTypeHNML MealType = "HNML"
	
	// MealTypeKSML ...
	MealTypeKSML MealType = "KSML"
	
	// MealTypeLCML ...
	MealTypeLCML MealType = "LCML"
	
	// MealTypeLFML ...
	MealTypeLFML MealType = "LFML"
	
	// MealTypeLPML ...
	MealTypeLPML MealType = "LPML"
	
	// MealTypeLSML ...
	MealTypeLSML MealType = "LSML"
	
	// MealTypeMOML ...
	MealTypeMOML MealType = "MOML"
	
	// MealTypeNLML ...
	MealTypeNLML MealType = "NLML"
	
	// MealTypeORML ...
	MealTypeORML MealType = "ORML"
	
	// MealTypePRML ...
	MealTypePRML MealType = "PRML"
	
	// MealTypeRVML ...
	MealTypeRVML MealType = "RVML"
	
	// MealTypeSFML ...
	MealTypeSFML MealType = "SFML"
	
	// MealTypeSPML ...
	MealTypeSPML MealType = "SPML"
	
	// MealTypeVGML ...
	MealTypeVGML MealType = "VGML"
	
	// MealTypeVLML ...
	MealTypeVLML MealType = "VLML"
	
	// MealTypeRGML ...
	MealTypeRGML MealType = "RGML"
)

// VehicleTransmissionType ...
type VehicleTransmissionType string

const (
	// VehicleTransmissionTypeA ...
	VehicleTransmissionTypeA VehicleTransmissionType = "A"
	
	// VehicleTransmissionTypeM ...
	VehicleTransmissionTypeM VehicleTransmissionType = "M"
)

// ActionType ...
type ActionType string

const (
	// ActionTypeAddUpdate ...
	ActionTypeAddUpdate ActionType = "Add-Update"
	
	// ActionTypeCancel ...
	ActionTypeCancel ActionType = "Cancel"
	
	// ActionTypeDelete ...
	ActionTypeDelete ActionType = "Delete"
	
	// ActionTypeAdd ...
	ActionTypeAdd ActionType = "Add"
	
	// ActionTypeReplace ...
	ActionTypeReplace ActionType = "Replace"
)

// FlightTypeType ...
type FlightTypeType string

const (
	// FlightTypeTypeNonstop ...
	FlightTypeTypeNonstop FlightTypeType = "Nonstop"
	
	// FlightTypeTypeDirect ...
	FlightTypeTypeDirect FlightTypeType = "Direct"
	
	// FlightTypeTypeConnection ...
	FlightTypeTypeConnection FlightTypeType = "Connection"
	
	// FlightTypeTypeSingleConnection ...
	FlightTypeTypeSingleConnection FlightTypeType = "SingleConnection"
	
	// FlightTypeTypeDoubleConnection ...
	FlightTypeTypeDoubleConnection FlightTypeType = "DoubleConnection"
	
	// FlightTypeTypeOneStopOnly ...
	FlightTypeTypeOneStopOnly FlightTypeType = "OneStopOnly"
)

// CabinType ...
type CabinType string

const (
	// CabinTypeFirst ...
	CabinTypeFirst CabinType = "First"
	
	// CabinTypeBusiness ...
	CabinTypeBusiness CabinType = "Business"
	
	// CabinTypeEconomy ...
	CabinTypeEconomy CabinType = "Economy"
	
	// CabinTypePremium ...
	CabinTypePremium CabinType = "Premium"
)

// TicketType ...
type TicketType string

const (
	// TicketTypeBookingOnly ...
	TicketTypeBookingOnly TicketType = "BookingOnly"
	
	// TicketTypeETicketInstant ...
	TicketTypeETicketInstant TicketType = "ETicketInstant"
	
	// TicketTypeETicketRequest ...
	TicketTypeETicketRequest TicketType = "ETicketRequest"
	
	// TicketTypeMyAddress ...
	TicketTypeMyAddress TicketType = "MyAddress"
	
	// TicketTypeTicketless ...
	TicketTypeTicketless TicketType = "Ticketless"
)

// OfficeLocationType ...
type OfficeLocationType string

const (
	// OfficeLocationTypeMain ...
	OfficeLocationTypeMain OfficeLocationType = "Main"
	
	// OfficeLocationTypeField ...
	OfficeLocationTypeField OfficeLocationType = "Field"
	
	// OfficeLocationTypeDivision ...
	OfficeLocationTypeDivision OfficeLocationType = "Division"
	
	// OfficeLocationTypeRegional ...
	OfficeLocationTypeRegional OfficeLocationType = "Regional"
	
	OfficeLocationTypeRemote OfficeLocationType = "Remote"
)

// YesNoType ...
type YesNoType string

const (
	// YesNoTypeYes ...
	YesNoTypeYes YesNoType = "Yes"
	
	// YesNoTypeNo ...
	YesNoTypeNo YesNoType = "No"
)

// PMS_ResStatusType ...
type PMS_ResStatusType string

const (
	// PMS_ResStatusTypeReserved ...
	PMS_ResStatusTypeReserved PMS_ResStatusType = "Reserved"
	
	// PMS_ResStatusTypeRequested ...
	PMS_ResStatusTypeRequested PMS_ResStatusType = "Requested"
	
	// PMS_ResStatusTypeRequestdenied ...
	PMS_ResStatusTypeRequestdenied PMS_ResStatusType = "Request denied"
	
	// PMS_ResStatusTypeNoshow ...
	PMS_ResStatusTypeNoshow PMS_ResStatusType = "No-show"
	
	// PMS_ResStatusTypeCancelled ...
	PMS_ResStatusTypeCancelled PMS_ResStatusType = "Cancelled"
	
	// PMS_ResStatusTypeInhouse ...
	PMS_ResStatusTypeInhouse PMS_ResStatusType = "In-house"
	
	// PMS_ResStatusTypeCheckedout ...
	PMS_ResStatusTypeCheckedout PMS_ResStatusType = "Checked out"
	
	// PMS_ResStatusTypeWaitlisted ...
	PMS_ResStatusTypeWaitlisted PMS_ResStatusType = "Waitlisted"
)

// TransactionStatusType ...
type TransactionStatusType string

const (
	// TransactionStatusTypePending ...
	TransactionStatusTypePending TransactionStatusType = "Pending"
	
	// TransactionStatusTypeCancelled ...
	TransactionStatusTypeCancelled TransactionStatusType = "Cancelled"
	
	// TransactionStatusTypeModified ...
	TransactionStatusTypeModified TransactionStatusType = "Modified"
	
	// TransactionStatusTypeCommitted ...
	TransactionStatusTypeCommitted TransactionStatusType = "Committed"
	
	// TransactionStatusTypeIgnored ...
	TransactionStatusTypeIgnored TransactionStatusType = "Ignored"
	
	// TransactionStatusTypeOnHold ...
	TransactionStatusTypeOnHold TransactionStatusType = "On Hold"
	
	// TransactionStatusTypeUnsuccessful ...
	TransactionStatusTypeUnsuccessful TransactionStatusType = "Unsuccessful"
)

// StayUnitType ...
type StayUnitType string

const (
	// StayUnitTypeMinutes ...
	StayUnitTypeMinutes StayUnitType = "Minutes"
	// StayUnitTypeHours ...
	StayUnitTypeHours StayUnitType = "Hours"
	// StayUnitTypeDays ...
	StayUnitTypeDays StayUnitType = "Days"
	// StayUnitTypeMonths ...
	StayUnitTypeMonths StayUnitType = "Months"
	// StayUnitTypeMON ...
	StayUnitTypeMON StayUnitType = "MON"
	// StayUnitTypeTUES ...
	StayUnitTypeTUES StayUnitType = "TUES"
	// StayUnitTypeWED ...
	StayUnitTypeWED StayUnitType = "WED"
	// StayUnitTypeTHU ...
	StayUnitTypeTHU StayUnitType = "THU"
	// StayUnitTypeFRI ...
	StayUnitTypeFRI StayUnitType = "FRI"
	// StayUnitTypeSAT ...
	StayUnitTypeSAT StayUnitType = "SAT"
	// StayUnitTypeSUN ...
	StayUnitTypeSUN StayUnitType = "SUN"
)

// AirTripType ...
type AirTripType string

const (
	// AirTripTypeOneWay ...
	AirTripTypeOneWay AirTripType = "OneWay"
	
	// AirTripTypeOneWayOnly ...
	AirTripTypeOneWayOnly AirTripType = "OneWayOnly"
	
	// AirTripTypeReturn ...
	AirTripTypeReturn AirTripType = "Return"
	
	// AirTripTypeCircle ...
	AirTripTypeCircle AirTripType = "Circle"
	
	// AirTripTypeOpenJaw ...
	AirTripTypeOpenJaw AirTripType = "OpenJaw"
	
	// AirTripTypeOther ...
	AirTripTypeOther AirTripType = "Other"
	
	// AirTripTypeOutbound ...
	AirTripTypeOutbound AirTripType = "Outbound"
	
	// AirTripTypeOutboundSeasonRoundtrip ...
	AirTripTypeOutboundSeasonRoundtrip AirTripType = "OutboundSeasonRoundtrip"
	
	// AirTripTypeNondirectional ...
	AirTripTypeNondirectional AirTripType = "Non-directional"
	
	// AirTripTypeInbound ...
	AirTripTypeInbound AirTripType = "Inbound"
	
	// AirTripTypeRoundtrip ...
	AirTripTypeRoundtrip AirTripType = "Roundtrip"
)

// GlobalIndicatorType ...
type GlobalIndicatorType string

const (
	// GlobalIndicatorTypeAP ...
	GlobalIndicatorTypeAP GlobalIndicatorType = "AP"
	
	// GlobalIndicatorTypeAT ...
	GlobalIndicatorTypeAT GlobalIndicatorType = "AT"
	
	// GlobalIndicatorTypeCT ...
	GlobalIndicatorTypeCT GlobalIndicatorType = "CT"
	
	// GlobalIndicatorTypeDO ...
	GlobalIndicatorTypeDO GlobalIndicatorType = "DO"
	
	// GlobalIndicatorTypeEH ...
	GlobalIndicatorTypeEH GlobalIndicatorType = "EH"
	
	// GlobalIndicatorTypeFE ...
	GlobalIndicatorTypeFE GlobalIndicatorType = "FE"
	
	// GlobalIndicatorTypePA ...
	GlobalIndicatorTypePA GlobalIndicatorType = "PA"
	
	// GlobalIndicatorTypePN ...
	GlobalIndicatorTypePN GlobalIndicatorType = "PN"
	
	// GlobalIndicatorTypePO ...
	GlobalIndicatorTypePO GlobalIndicatorType = "PO"
	
	// GlobalIndicatorTypeRU ....
	GlobalIndicatorTypeRU GlobalIndicatorType = "RU"
	
	// GlobalIndicatorTypeRW ...
	GlobalIndicatorTypeRW GlobalIndicatorType = "RW"
	
	// GlobalIndicatorTypeSA ...
	GlobalIndicatorTypeSA GlobalIndicatorType = "SA"
	
	// GlobalIndicatorTypeTS ...
	GlobalIndicatorTypeTS GlobalIndicatorType = "TS"
	
	// GlobalIndicatorTypeWH ...
	GlobalIndicatorTypeWH GlobalIndicatorType = "WH"
)

// FareStatusType ...
type FareStatusType string

const (
	// FareStatusTypeConstructed ...
	FareStatusTypeConstructed FareStatusType = "constructed"
	
	// FareStatusTypePublished ...
	FareStatusTypePublished FareStatusType = "published"
	
	// FareStatusTypeCreated ...
	FareStatusTypeCreated FareStatusType = "created"
	
	// FareStatusTypeFareByRule ...
	FareStatusTypeFareByRule FareStatusType = "fareByRule"
	
	// FareStatusTypeFareByRulePrivate ...
	FareStatusTypeFareByRulePrivate FareStatusType = "fareByRulePrivate"
)

// LocationDetailShuttleInfoType ...
type LocationDetailShuttleInfoType string

const (
	// LocationDetailShuttleInfoTypeTransportation ...
	LocationDetailShuttleInfoTypeTransportation LocationDetailShuttleInfoType = "Transportation"
	
	// LocationDetailShuttleInfoTypeFrequency ....
	LocationDetailShuttleInfoTypeFrequency LocationDetailShuttleInfoType = "Frequency"
	
	// LocationDetailShuttleInfoTypePickupInfo ....
	LocationDetailShuttleInfoTypePickupInfo LocationDetailShuttleInfoType = "PickupInfo"
	
	// LocationDetailShuttleInfoTypeDistance ...
	LocationDetailShuttleInfoTypeDistance LocationDetailShuttleInfoType = "Distance"
	
	// LocationDetailShuttleInfoTypeElapsedTime ...
	LocationDetailShuttleInfoTypeElapsedTime LocationDetailShuttleInfoType = "ElapsedTime"
	
	// LocationDetailShuttleInfoTypeFee ...
	LocationDetailShuttleInfoTypeFee LocationDetailShuttleInfoType = "Fee"
	
	// LocationDetailShuttleInfoTypeMiscellaneous ...
	LocationDetailShuttleInfoTypeMiscellaneous LocationDetailShuttleInfoType = "Miscellaneous"
	
	// LocationDetailShuttleInfoTypeHours ...
	LocationDetailShuttleInfoTypeHours LocationDetailShuttleInfoType = "Hours"
)

// PricingSourceType ...
type PricingSourceType string

const (
	// PricingSourceTypePublished ...
	PricingSourceTypePublished PricingSourceType = "Published"
	
	// PricingSourceTypePrivate ...
	PricingSourceTypePrivate PricingSourceType = "Private"
	
	// PricingSourceTypeBoth ...
	PricingSourceTypeBoth PricingSourceType = "Both"
)

// PhoneTypeEnum ...
type PhoneTypeEnum string

const (
	// PhoneTypeEnumHome ...
	PhoneTypeEnumHome PhoneTypeEnum = "Home"
	
	// PhoneTypeEnumOffice ...
	PhoneTypeEnumOffice PhoneTypeEnum = "Office"
	
	// PhoneTypeEnumGsm ...
	PhoneTypeEnumGsm PhoneTypeEnum = "Gsm"
)

// VehiclePeriodUnitNameType ...
type VehiclePeriodUnitNameType string

const (
	// VehiclePeriodUnitNameTypeRentalPeriod ...
	VehiclePeriodUnitNameTypeRentalPeriod VehiclePeriodUnitNameType = "RentalPeriod"
	
	// VehiclePeriodUnitNameTypeYear ...
	VehiclePeriodUnitNameTypeYear VehiclePeriodUnitNameType = "Year"
	
	// VehiclePeriodUnitNameTypeMonth ...
	VehiclePeriodUnitNameTypeMonth VehiclePeriodUnitNameType = "Month"
	
	// VehiclePeriodUnitNameTypeWeek ....
	VehiclePeriodUnitNameTypeWeek VehiclePeriodUnitNameType = "Week"
	
	// VehiclePeriodUnitNameTypeDay ...
	VehiclePeriodUnitNameTypeDay VehiclePeriodUnitNameType = "Day"
	
	// VehiclePeriodUnitNameTypeHour ...
	VehiclePeriodUnitNameTypeHour VehiclePeriodUnitNameType = "Hour"
	
	// VehiclePeriodUnitNameTypeWeekend ...
	VehiclePeriodUnitNameTypeWeekend VehiclePeriodUnitNameType = "Weekend"
	
	// VehiclePeriodUnitNameTypeBundle ...
	VehiclePeriodUnitNameTypeBundle VehiclePeriodUnitNameType = "Bundle"
	
	// VehiclePeriodUnitNameTypePackage ...
	VehiclePeriodUnitNameTypePackage VehiclePeriodUnitNameType = "Package"
)

// VehicleChargeTypeCalculationApplicability ...
type VehicleChargeTypeCalculationApplicability string

const (
	// VehicleChargeTypeCalculationApplicabilityFromPickupLocation ...
	VehicleChargeTypeCalculationApplicabilityFromPickupLocation VehicleChargeTypeCalculationApplicability = "FromPickupLocation"
	
	// VehicleChargeTypeCalculationApplicabilityFromDropoffLocation ...
	VehicleChargeTypeCalculationApplicabilityFromDropoffLocation VehicleChargeTypeCalculationApplicability = "FromDropoffLocation"
	
	// VehicleChargeTypeCalculationApplicabilityBeforePickup ...
	VehicleChargeTypeCalculationApplicabilityBeforePickup VehicleChargeTypeCalculationApplicability = "BeforePickup"
	
	// VehicleChargeTypeCalculationApplicabilityAfterDropoff ...
	VehicleChargeTypeCalculationApplicabilityAfterDropoff VehicleChargeTypeCalculationApplicability = "AfterDropoff"
)

// RateQualifierTypeRatePeriod ...
type RateQualifierTypeRatePeriod string

const (
	// RateQualifierTypeRatePeriodHourly ...
	RateQualifierTypeRatePeriodHourly RateQualifierTypeRatePeriod = "Hourly"
	
	// RateQualifierTypeRatePeriodDaily ...
	RateQualifierTypeRatePeriodDaily RateQualifierTypeRatePeriod = "Daily"
	
	// RateQualifierTypeRatePeriodWeekly ...
	RateQualifierTypeRatePeriodWeekly RateQualifierTypeRatePeriod = "Weekly"
	
	// RateQualifierTypeRatePeriodMonthly ....
	RateQualifierTypeRatePeriodMonthly RateQualifierTypeRatePeriod = "Monthly"
	
	// RateQualifierTypeRatePeriodWeekendDay ...
	RateQualifierTypeRatePeriodWeekendDay RateQualifierTypeRatePeriod = "WeekendDay"
	
	// RateQualifierTypeRatePeriodOther ...
	RateQualifierTypeRatePeriodOther RateQualifierTypeRatePeriod = "Other"
	
	// RateQualifierTypeRatePeriodPackage ...
	RateQualifierTypeRatePeriodPackage RateQualifierTypeRatePeriod = "Package"
	
	// RateQualifierTypeRatePeriodBundle ...
	RateQualifierTypeRatePeriodBundle RateQualifierTypeRatePeriod = "Bundle"
	
	// RateQualifierTypeRatePeriodTotal ...
	RateQualifierTypeRatePeriodTotal RateQualifierTypeRatePeriod = "Total"
)

// VehicleRentalRateTypeRateRestrictionsOneWayPolicy ...
type VehicleRentalRateTypeRateRestrictionsOneWayPolicy string

const (
	// VehicleRentalRateTypeRateRestrictionsOneWayPolicyOneWayAllowed ...
	VehicleRentalRateTypeRateRestrictionsOneWayPolicyOneWayAllowed VehicleRentalRateTypeRateRestrictionsOneWayPolicy = "OneWayAllowed"
	
	// VehicleRentalRateTypeRateRestrictionsOneWayPolicyOneWayNotAllowed ...
	VehicleRentalRateTypeRateRestrictionsOneWayPolicyOneWayNotAllowed VehicleRentalRateTypeRateRestrictionsOneWayPolicy = "OneWayNotAllowed"
	
	// VehicleRentalRateTypeRateRestrictionsOneWayPolicyRestrictedOneWay ...
	VehicleRentalRateTypeRateRestrictionsOneWayPolicyRestrictedOneWay VehicleRentalRateTypeRateRestrictionsOneWayPolicy = "RestrictedOneWay"
)

// VehicleRentalRateTypeRateGuaranteeOffsetDropTime ...
type VehicleRentalRateTypeRateGuaranteeOffsetDropTime string

const (
	// VehicleRentalRateTypeRateGuaranteeOffsetDropTimeBeforeArrival ...
	VehicleRentalRateTypeRateGuaranteeOffsetDropTimeBeforeArrival VehicleRentalRateTypeRateGuaranteeOffsetDropTime = "BeforeArrival"
	
	// VehicleRentalRateTypeRateGuaranteeOffsetDropTimeAfterBooking ...
	VehicleRentalRateTypeRateGuaranteeOffsetDropTimeAfterBooking VehicleRentalRateTypeRateGuaranteeOffsetDropTime = "AfterBooking"
	
	// VehicleRentalRateTypeRateGuaranteeOffsetDropTimeAfterConfirmation ...
	VehicleRentalRateTypeRateGuaranteeOffsetDropTimeAfterConfirmation VehicleRentalRateTypeRateGuaranteeOffsetDropTime = "AfterConfirmation"
)

// VehicleRentalRateTypePickupReturnRuleRuleType ...
type VehicleRentalRateTypePickupReturnRuleRuleType string

const (
	// VehicleRentalRateTypePickupReturnRuleRuleTypeEarliestPickup ...
	VehicleRentalRateTypePickupReturnRuleRuleTypeEarliestPickup VehicleRentalRateTypePickupReturnRuleRuleType = "EarliestPickup"
	
	// VehicleRentalRateTypePickupReturnRuleRuleTypeLatestPickup ...
	VehicleRentalRateTypePickupReturnRuleRuleTypeLatestPickup VehicleRentalRateTypePickupReturnRuleRuleType = "LatestPickup"
	
	// VehicleRentalRateTypePickupReturnRuleRuleTypeLatestReturn ...
	VehicleRentalRateTypePickupReturnRuleRuleTypeLatestReturn VehicleRentalRateTypePickupReturnRuleRuleType = "LatestReturn"
)

// EquipmentRestrictionType ...
type EquipmentRestrictionType string

const (
	// EquipmentRestrictionTypeOneWayOnly ...
	EquipmentRestrictionTypeOneWayOnly EquipmentRestrictionType = "OneWayOnly"
	
	// EquipmentRestrictionTypeRoundTripOnly ...
	EquipmentRestrictionTypeRoundTripOnly EquipmentRestrictionType = "RoundTripOnly"
	
	// EquipmentRestrictionTypeAnyReservation ...
	EquipmentRestrictionTypeAnyReservation EquipmentRestrictionType = "AnyReservation"
)

// OffLocationServiceID_Type ...
type OffLocationServiceID_Type string

const (
	// OffLocationServiceID_TypeCustPickUp ...
	OffLocationServiceID_TypeCustPickUp OffLocationServiceID_Type = "CustPickUp"
	
	// OffLocationServiceID_TypeVehDelivery ...
	OffLocationServiceID_TypeVehDelivery OffLocationServiceID_Type = "VehDelivery"
	
	// OffLocationServiceID_TypeCustDropOff ...
	OffLocationServiceID_TypeCustDropOff OffLocationServiceID_Type = "CustDropOff"
	
	// OffLocationServiceID_TypeVehCollection ...
	OffLocationServiceID_TypeVehCollection OffLocationServiceID_Type = "VehCollection"
	
	// OffLocationServiceID_TypeExchange ...
	OffLocationServiceID_TypeExchange OffLocationServiceID_Type = "Exchange"
)

// OTA_VehResRSTarget ...
type OTA_VehResRSTarget string

const (
	// OTA_VehResRSTargetTest ...
	OTA_VehResRSTargetTest OTA_VehResRSTarget = "Test"
	
	// OTA_VehResRSTargetProduction ...
	OTA_VehResRSTargetProduction OTA_VehResRSTarget = "Production"
)

// OTA_VehResRSTransactionStatusCode ...
type OTA_VehResRSTransactionStatusCode string

const (
	// OTA_VehResRSTransactionStatusCodeStart ...
	OTA_VehResRSTransactionStatusCodeStart OTA_VehResRSTransactionStatusCode = "Start"
	
	// OTA_VehResRSTransactionStatusCodeEnd ...
	OTA_VehResRSTransactionStatusCodeEnd OTA_VehResRSTransactionStatusCode = "End"
	
	// OTA_VehResRSTransactionStatusCodeRollback ...
	OTA_VehResRSTransactionStatusCodeRollback OTA_VehResRSTransactionStatusCode = "Rollback"
	
	// OTA_VehResRSTransactionStatusCodeInSeries ...
	OTA_VehResRSTransactionStatusCodeInSeries OTA_VehResRSTransactionStatusCode = "InSeries"
	
	// OTA_VehResRSTransactionStatusCodeContinuation ...
	OTA_VehResRSTransactionStatusCodeContinuation OTA_VehResRSTransactionStatusCode = "Continuation"
	
	// OTA_VehResRSTransactionStatusCodeSubsequent ...
	OTA_VehResRSTransactionStatusCodeSubsequent OTA_VehResRSTransactionStatusCode = "Subsequent"
)

// DisplayOrderType ...
type DisplayOrderType string

const (
	// DisplayOrderTypeByDepartureTime ...
	DisplayOrderTypeByDepartureTime DisplayOrderType = "ByDepartureTime"
	
	// DisplayOrderTypeByArrivalTime ...
	DisplayOrderTypeByArrivalTime DisplayOrderType = "ByArrivalTime"
	
	// DisplayOrderTypeByJourneyTime ...
	DisplayOrderTypeByJourneyTime DisplayOrderType = "ByJourneyTime"
	
	// DisplayOrderTypeByPriceHighToLow ...
	DisplayOrderTypeByPriceHighToLow DisplayOrderType = "ByPriceHighToLow"
	
	// DisplayOrderTypeByPriceLowToHigh ...
	DisplayOrderTypeByPriceLowToHigh DisplayOrderType = "ByPriceLowToHigh"
)

// AllianceCode ...
type AllianceCode string

const (
	// AllianceCodeONEWORLD ...
	AllianceCodeONEWORLD AllianceCode = "ONEWORLD"
	
	// AllianceCodeSKYTEAM ...
	AllianceCodeSKYTEAM AllianceCode = "SKYTEAM"
	
	// AllianceCodeSTAR_ALLIANCE ...
	AllianceCodeSTAR_ALLIANCE AllianceCode = "STAR_ALLIANCE"
)

// AirlineDiversityEnum ...
type AirlineDiversityEnum string

const (
	// AirlineDiversityEnumADC ...
	AirlineDiversityEnumADC AirlineDiversityEnum = "ADC"
	
	// AirlineDiversityEnumNAD ...
	AirlineDiversityEnumNAD AirlineDiversityEnum = "NAD"
	
	// AirlineDiversityEnumNDD ...
	AirlineDiversityEnumNDD AirlineDiversityEnum = "NDD"
)

// ExpandedParameterEnum ...
type ExpandedParameterEnum string

const (
	// ExpandedParameterEnumNAP ...
	ExpandedParameterEnumNAP ExpandedParameterEnum = "NAP"
	
	// ExpandedParameterEnumNPE ...
	ExpandedParameterEnumNPE ExpandedParameterEnum = "NPE"
	
	// ExpandedParameterEnumNR ...
	ExpandedParameterEnumNR ExpandedParameterEnum = "NR"
)

// DuplicatedRecommendationsBehaviorEnum ...
type DuplicatedRecommendationsBehaviorEnum string

const (
	// DuplicatedRecommedationsBehaviorEnumRDG ...
	DuplicatedRecommedationsBehaviorEnumRDG DuplicatedRecommendationsBehaviorEnum = "RDG"
	
	// DuplicatedRecommedationsBehaviorEnumRDM ...
	DuplicatedRecommedationsBehaviorEnumRDM DuplicatedRecommendationsBehaviorEnum = "RDM"
	
	// DuplicatedRecommedationsBehaviorEnumNONE ...
	DuplicatedRecommedationsBehaviorEnumNONE DuplicatedRecommendationsBehaviorEnum = "NONE"
)

// RefundableTypesEnum ...
type RefundableTypesEnum string

const (
	// RefundableTypesEnumAllFlights ....
	RefundableTypesEnumAllFlights RefundableTypesEnum = "AllFlights"
	
	// RefundableTypesEnumOnlyRefundable ...
	RefundableTypesEnumOnlyRefundable RefundableTypesEnum = "OnlyRefundable"
	
	// RefundableTypesEnumOnlyNonRefundable ...
	RefundableTypesEnumOnlyNonRefundable RefundableTypesEnum = "OnlyNonRefundable"
)

// SearchTypeByProviderEnum ...
type SearchTypeByProviderEnum string

const (
	// SearchTypeByProviderEnumOnlyAmadeus ...
	SearchTypeByProviderEnumOnlyAmadeus SearchTypeByProviderEnum = "OnlyAmadeus"
	
	// SearchTypeByProviderEnumAmadeusAndExtProviders ...
	SearchTypeByProviderEnumAmadeusAndExtProviders SearchTypeByProviderEnum = "AmadeusAndExtProviders"
	
	// SearchTypeByProviderEnumOnlyExtProviders ...
	SearchTypeByProviderEnumOnlyExtProviders SearchTypeByProviderEnum = "OnlyExtProviders"
)

// OTA_AirRulesRQTarget ....
type OTA_AirRulesRQTarget string

const (
	// OTA_AirRulesRQTargetTest ...
	OTA_AirRulesRQTargetTest OTA_AirRulesRQTarget = "Test"
	
	// OTA_AirRulesRQTargetProduction ...
	OTA_AirRulesRQTargetProduction OTA_AirRulesRQTarget = "Production"
)

// OTA_AirRulesRQTransactionStatusCode ...
type OTA_AirRulesRQTransactionStatusCode string

const (
	// OTA_AirRulesRQTransactionStatusCodeStart ...
	OTA_AirRulesRQTransactionStatusCodeStart OTA_AirRulesRQTransactionStatusCode = "Start"
	
	// OTA_AirRulesRQTransactionStatusCodeEnd ...
	OTA_AirRulesRQTransactionStatusCodeEnd OTA_AirRulesRQTransactionStatusCode = "End"
	
	// OTA_AirRulesRQTransactionStatusCodeRollback ...
	OTA_AirRulesRQTransactionStatusCodeRollback OTA_AirRulesRQTransactionStatusCode = "Rollback"
	
	// OTA_AirRulesRQTransactionStatusCodeInSeries ...
	OTA_AirRulesRQTransactionStatusCodeInSeries OTA_AirRulesRQTransactionStatusCode = "InSeries"
	
	// OTA_AirRulesRQTransactionStatusCodeContinuation ...
	OTA_AirRulesRQTransactionStatusCodeContinuation OTA_AirRulesRQTransactionStatusCode = "Continuation"
	
	// OTA_AirRulesRQTransactionStatusCodeSubsequent ...
	OTA_AirRulesRQTransactionStatusCodeSubsequent OTA_AirRulesRQTransactionStatusCode = "Subsequent"
)

// OTA_AirRulesRSTarget ...
type OTA_AirRulesRSTarget string

const (
	// OTA_AirRulesRSTargetTest ...
	OTA_AirRulesRSTargetTest OTA_AirRulesRSTarget = "Test"
	
	// OTA_AirRulesRSTargetProduction ...
	OTA_AirRulesRSTargetProduction OTA_AirRulesRSTarget = "Production"
)

// OTA_AirRulesRSTransactionStatusCode ...
type OTA_AirRulesRSTransactionStatusCode string

const (
	// OTA_AirRulesRSTransactionStatusCodeStart ...
	OTA_AirRulesRSTransactionStatusCodeStart OTA_AirRulesRSTransactionStatusCode = "Start"
	
	// OTA_AirRulesRSTransactionStatusCodeEnd ...
	OTA_AirRulesRSTransactionStatusCodeEnd OTA_AirRulesRSTransactionStatusCode = "End"
	
	// OTA_AirRulesRSTransactionStatusCodeRollback ...
	OTA_AirRulesRSTransactionStatusCodeRollback OTA_AirRulesRSTransactionStatusCode = "Rollback"
	
	// OTA_AirRulesRSTransactionStatusCodeInSeries ...
	OTA_AirRulesRSTransactionStatusCodeInSeries OTA_AirRulesRSTransactionStatusCode = "InSeries"
	
	// OTA_AirRulesRSTransactionStatusCodeContinuation ...
	OTA_AirRulesRSTransactionStatusCodeContinuation OTA_AirRulesRSTransactionStatusCode = "Continuation"
	
	// OTA_AirRulesRSTransactionStatusCodeSubsequent ...
	OTA_AirRulesRSTransactionStatusCodeSubsequent OTA_AirRulesRSTransactionStatusCode = "Subsequent"
)

// ChangeDetailActionTypes ...
type ChangeDetailActionTypes string

const (
	// ChangeDetailActionTypesAdd ...
	ChangeDetailActionTypesAdd ChangeDetailActionTypes = "Add"
	
	// ChangeDetailActionTypesModify ...
	ChangeDetailActionTypesModify ChangeDetailActionTypes = "Modify"
	
	// ChangeDetailActionTypesDelete ...
	ChangeDetailActionTypesDelete ChangeDetailActionTypes = "Delete"
)

// OTA_AirSeatMapRQAirTravelerGender ...
type OTA_AirSeatMapRQAirTravelerGender string

const (
	OTA_AirSeatMapRQAirTravelerGenderMale OTA_AirSeatMapRQAirTravelerGender = "Male"
	
	OTA_AirSeatMapRQAirTravelerGenderFemale OTA_AirSeatMapRQAirTravelerGender = "Female"
	
	OTA_AirSeatMapRQAirTravelerGenderUnknown OTA_AirSeatMapRQAirTravelerGender = "Unknown"
)

type OTA_AirSeatMapRQTarget string

const (
	OTA_AirSeatMapRQTargetTest OTA_AirSeatMapRQTarget = "Test"
	
	OTA_AirSeatMapRQTargetProduction OTA_AirSeatMapRQTarget = "Production"
)

type OTA_AirSeatMapRQTransactionStatusCode string

const (
	OTA_AirSeatMapRQTransactionStatusCodeStart OTA_AirSeatMapRQTransactionStatusCode = "Start"
	
	OTA_AirSeatMapRQTransactionStatusCodeEnd OTA_AirSeatMapRQTransactionStatusCode = "End"
	
	OTA_AirSeatMapRQTransactionStatusCodeRollback OTA_AirSeatMapRQTransactionStatusCode = "Rollback"
	
	OTA_AirSeatMapRQTransactionStatusCodeInSeries OTA_AirSeatMapRQTransactionStatusCode = "InSeries"
	
	OTA_AirSeatMapRQTransactionStatusCodeContinuation OTA_AirSeatMapRQTransactionStatusCode = "Continuation"
	
	OTA_AirSeatMapRQTransactionStatusCodeSubsequent OTA_AirSeatMapRQTransactionStatusCode = "Subsequent"
)

type OTA_AirSeatMapRSSeatMapResponsesAirTravelerGender string

const (
	OTA_AirSeatMapRSSeatMapResponsesAirTravelerGenderMale OTA_AirSeatMapRSSeatMapResponsesAirTravelerGender = "Male"
	
	OTA_AirSeatMapRSSeatMapResponsesAirTravelerGenderFemale OTA_AirSeatMapRSSeatMapResponsesAirTravelerGender = "Female"
	
	OTA_AirSeatMapRSSeatMapResponsesAirTravelerGenderUnknown OTA_AirSeatMapRSSeatMapResponsesAirTravelerGender = "Unknown"
)

type OTA_AirSeatMapRSTarget string

const (
	OTA_AirSeatMapRSTargetTest OTA_AirSeatMapRSTarget = "Test"
	
	OTA_AirSeatMapRSTargetProduction OTA_AirSeatMapRSTarget = "Production"
)

type OTA_AirSeatMapRSTransactionStatusCode string

const (
	OTA_AirSeatMapRSTransactionStatusCodeStart OTA_AirSeatMapRSTransactionStatusCode = "Start"
	
	OTA_AirSeatMapRSTransactionStatusCodeEnd OTA_AirSeatMapRSTransactionStatusCode = "End"
	
	OTA_AirSeatMapRSTransactionStatusCodeRollback OTA_AirSeatMapRSTransactionStatusCode = "Rollback"
	
	OTA_AirSeatMapRSTransactionStatusCodeInSeries OTA_AirSeatMapRSTransactionStatusCode = "InSeries"
	
	OTA_AirSeatMapRSTransactionStatusCodeContinuation OTA_AirSeatMapRSTransactionStatusCode = "Continuation"
	
	OTA_AirSeatMapRSTransactionStatusCodeSubsequent OTA_AirSeatMapRSTransactionStatusCode = "Subsequent"
)

type OTA_AirScheduleRQFlightTypePrefNonScheduledFltInfo string

const (
	OTA_AirScheduleRQFlightTypePrefNonScheduledFltInfoChartersOnly OTA_AirScheduleRQFlightTypePrefNonScheduledFltInfo = "ChartersOnly"
	
	OTA_AirScheduleRQFlightTypePrefNonScheduledFltInfoExcludeCharters OTA_AirScheduleRQFlightTypePrefNonScheduledFltInfo = "ExcludeCharters"
	
	OTA_AirScheduleRQFlightTypePrefNonScheduledFltInfoAll OTA_AirScheduleRQFlightTypePrefNonScheduledFltInfo = "All"
)

type OTA_AirScheduleRQFlightTypePrefRoutingType string

const (
	OTA_AirScheduleRQFlightTypePrefRoutingTypeNormal OTA_AirScheduleRQFlightTypePrefRoutingType = "Normal"
	
	OTA_AirScheduleRQFlightTypePrefRoutingTypeMirror OTA_AirScheduleRQFlightTypePrefRoutingType = "Mirror"
)

type OTA_AirScheduleRQTarget string

const (
	OTA_AirScheduleRQTargetTest OTA_AirScheduleRQTarget = "Test"
	
	OTA_AirScheduleRQTargetProduction OTA_AirScheduleRQTarget = "Production"
)

type OTA_AirScheduleRQTransactionStatusCode string

const (
	OTA_AirScheduleRQTransactionStatusCodeStart OTA_AirScheduleRQTransactionStatusCode = "Start"
	
	OTA_AirScheduleRQTransactionStatusCodeEnd OTA_AirScheduleRQTransactionStatusCode = "End"
	
	OTA_AirScheduleRQTransactionStatusCodeRollback OTA_AirScheduleRQTransactionStatusCode = "Rollback"
	
	OTA_AirScheduleRQTransactionStatusCodeInSeries OTA_AirScheduleRQTransactionStatusCode = "InSeries"
	
	OTA_AirScheduleRQTransactionStatusCodeContinuation OTA_AirScheduleRQTransactionStatusCode = "Continuation"
	
	OTA_AirScheduleRQTransactionStatusCodeSubsequent OTA_AirScheduleRQTransactionStatusCode = "Subsequent"
)

type OTA_AirScheduleRSTarget string

const (
	OTA_AirScheduleRSTargetTest OTA_AirScheduleRSTarget = "Test"
	
	OTA_AirScheduleRSTargetProduction OTA_AirScheduleRSTarget = "Production"
)

type OTA_AirScheduleRSTransactionStatusCode string

const (
	OTA_AirScheduleRSTransactionStatusCodeStart OTA_AirScheduleRSTransactionStatusCode = "Start"
	
	OTA_AirScheduleRSTransactionStatusCodeEnd OTA_AirScheduleRSTransactionStatusCode = "End"
	
	OTA_AirScheduleRSTransactionStatusCodeRollback OTA_AirScheduleRSTransactionStatusCode = "Rollback"
	
	OTA_AirScheduleRSTransactionStatusCodeInSeries OTA_AirScheduleRSTransactionStatusCode = "InSeries"
	
	OTA_AirScheduleRSTransactionStatusCodeContinuation OTA_AirScheduleRSTransactionStatusCode = "Continuation"
	
	OTA_AirScheduleRSTransactionStatusCodeSubsequent OTA_AirScheduleRSTransactionStatusCode = "Subsequent"
)

type VehicleAvailRQCoreTypeRateQualifierRatePeriod string

const (
	VehicleAvailRQCoreTypeRateQualifierRatePeriodHourly VehicleAvailRQCoreTypeRateQualifierRatePeriod = "Hourly"
	
	VehicleAvailRQCoreTypeRateQualifierRatePeriodDaily VehicleAvailRQCoreTypeRateQualifierRatePeriod = "Daily"
	
	VehicleAvailRQCoreTypeRateQualifierRatePeriodWeekly VehicleAvailRQCoreTypeRateQualifierRatePeriod = "Weekly"
	
	VehicleAvailRQCoreTypeRateQualifierRatePeriodMonthly VehicleAvailRQCoreTypeRateQualifierRatePeriod = "Monthly"
	
	VehicleAvailRQCoreTypeRateQualifierRatePeriodWeekendDay VehicleAvailRQCoreTypeRateQualifierRatePeriod = "WeekendDay"
	
	VehicleAvailRQCoreTypeRateQualifierRatePeriodOther VehicleAvailRQCoreTypeRateQualifierRatePeriod = "Other"
	
	VehicleAvailRQCoreTypeRateQualifierRatePeriodPackage VehicleAvailRQCoreTypeRateQualifierRatePeriod = "Package"
	
	VehicleAvailRQCoreTypeRateQualifierRatePeriodBundle VehicleAvailRQCoreTypeRateQualifierRatePeriod = "Bundle"
	
	VehicleAvailRQCoreTypeRateQualifierRatePeriodTotal VehicleAvailRQCoreTypeRateQualifierRatePeriod = "Total"
)

type InventoryStatusType string

const (
	InventoryStatusTypeAvailable InventoryStatusType = "Available"
	
	InventoryStatusTypeUnavailable InventoryStatusType = "Unavailable"
	
	InventoryStatusTypeOnRequest InventoryStatusType = "OnRequest"
	
	InventoryStatusTypeConfirmed InventoryStatusType = "Confirmed"
	
	InventoryStatusTypeAll InventoryStatusType = "All"
	
	InventoryStatusTypeWaitlist InventoryStatusType = "Waitlist"
	
	InventoryStatusTypeSupplierBooked InventoryStatusType = "SupplierBooked"
)

type OTA_VehAvailRateRQTarget string

const (
	OTA_VehAvailRateRQTargetTest OTA_VehAvailRateRQTarget = "Test"
	
	OTA_VehAvailRateRQTargetProduction OTA_VehAvailRateRQTarget = "Production"
)

type OTA_VehAvailRateRQTransactionStatusCode string

const (
	OTA_VehAvailRateRQTransactionStatusCodeStart OTA_VehAvailRateRQTransactionStatusCode = "Start"
	
	OTA_VehAvailRateRQTransactionStatusCodeEnd OTA_VehAvailRateRQTransactionStatusCode = "End"
	
	OTA_VehAvailRateRQTransactionStatusCodeRollback OTA_VehAvailRateRQTransactionStatusCode = "Rollback"
	
	OTA_VehAvailRateRQTransactionStatusCodeInSeries OTA_VehAvailRateRQTransactionStatusCode = "InSeries"
	
	OTA_VehAvailRateRQTransactionStatusCodeContinuation OTA_VehAvailRateRQTransactionStatusCode = "Continuation"
	
	OTA_VehAvailRateRQTransactionStatusCodeSubsequent OTA_VehAvailRateRQTransactionStatusCode = "Subsequent"
)

type VehicleAvailCoreTypeVendorLocationCounterLocInfo string

const (
	VehicleAvailCoreTypeVendorLocationCounterLocInfoWalkToCar VehicleAvailCoreTypeVendorLocationCounterLocInfo = "WalkToCar"
	
	VehicleAvailCoreTypeVendorLocationCounterLocInfoShuttleToCar VehicleAvailCoreTypeVendorLocationCounterLocInfo = "ShuttleToCar"
)

type OTA_VehAvailRateRSTarget string

const (
	OTA_VehAvailRateRSTargetTest OTA_VehAvailRateRSTarget = "Test"
	
	OTA_VehAvailRateRSTargetProduction OTA_VehAvailRateRSTarget = "Production"
)

type OTA_VehAvailRateRSTransactionStatusCode string

const (
	OTA_VehAvailRateRSTransactionStatusCodeStart OTA_VehAvailRateRSTransactionStatusCode = "Start"
	
	OTA_VehAvailRateRSTransactionStatusCodeEnd OTA_VehAvailRateRSTransactionStatusCode = "End"
	
	OTA_VehAvailRateRSTransactionStatusCodeRollback OTA_VehAvailRateRSTransactionStatusCode = "Rollback"
	
	OTA_VehAvailRateRSTransactionStatusCodeInSeries OTA_VehAvailRateRSTransactionStatusCode = "InSeries"
	
	OTA_VehAvailRateRSTransactionStatusCodeContinuation OTA_VehAvailRateRSTransactionStatusCode = "Continuation"
	
	OTA_VehAvailRateRSTransactionStatusCodeSubsequent OTA_VehAvailRateRSTransactionStatusCode = "Subsequent"
)

type OTA_VehLocSearchRQTarget string

const (
	OTA_VehLocSearchRQTargetTest OTA_VehLocSearchRQTarget = "Test"
	
	OTA_VehLocSearchRQTargetProduction OTA_VehLocSearchRQTarget = "Production"
)

type OTA_VehLocSearchRQTransactionStatusCode string

const (
	OTA_VehLocSearchRQTransactionStatusCodeStart OTA_VehLocSearchRQTransactionStatusCode = "Start"
	
	OTA_VehLocSearchRQTransactionStatusCodeEnd OTA_VehLocSearchRQTransactionStatusCode = "End"
	
	OTA_VehLocSearchRQTransactionStatusCodeRollback OTA_VehLocSearchRQTransactionStatusCode = "Rollback"
	
	OTA_VehLocSearchRQTransactionStatusCodeInSeries OTA_VehLocSearchRQTransactionStatusCode = "InSeries"
	
	OTA_VehLocSearchRQTransactionStatusCodeContinuation OTA_VehLocSearchRQTransactionStatusCode = "Continuation"
	
	OTA_VehLocSearchRQTransactionStatusCodeSubsequent OTA_VehLocSearchRQTransactionStatusCode = "Subsequent"
)

type OTA_VehLocSearchRSTarget string

const (
	OTA_VehLocSearchRSTargetTest OTA_VehLocSearchRSTarget = "Test"
	
	OTA_VehLocSearchRSTargetProduction OTA_VehLocSearchRSTarget = "Production"
)

type OTA_VehLocSearchRSTransactionStatusCode string

const (
	OTA_VehLocSearchRSTransactionStatusCodeStart OTA_VehLocSearchRSTransactionStatusCode = "Start"
	
	OTA_VehLocSearchRSTransactionStatusCodeEnd OTA_VehLocSearchRSTransactionStatusCode = "End"
	
	OTA_VehLocSearchRSTransactionStatusCodeRollback OTA_VehLocSearchRSTransactionStatusCode = "Rollback"
	
	OTA_VehLocSearchRSTransactionStatusCodeInSeries OTA_VehLocSearchRSTransactionStatusCode = "InSeries"
	
	OTA_VehLocSearchRSTransactionStatusCodeContinuation OTA_VehLocSearchRSTransactionStatusCode = "Continuation"
	
	OTA_VehLocSearchRSTransactionStatusCodeSubsequent OTA_VehLocSearchRSTransactionStatusCode = "Subsequent"
)

type OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod string

const (
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodHourly OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "Hourly"
	
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodDaily OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "Daily"
	
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodWeekly OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "Weekly"
	
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodMonthly OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "Monthly"
	
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodWeekendDay OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "WeekendDay"
	
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodOther OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "Other"
	
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodPackage OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "Package"
	
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodBundle OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "Bundle"
	
	OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriodTotal OTA_VehRateRuleRQRentalInfoRateQualifierRatePeriod = "Total"
)

type OTA_VehRateRuleRQTarget string

const (
	OTA_VehRateRuleRQTargetTest OTA_VehRateRuleRQTarget = "Test"
	
	OTA_VehRateRuleRQTargetProduction OTA_VehRateRuleRQTarget = "Production"
)

type OTA_VehRateRuleRQTransactionStatusCode string

const (
	OTA_VehRateRuleRQTransactionStatusCodeStart OTA_VehRateRuleRQTransactionStatusCode = "Start"
	
	OTA_VehRateRuleRQTransactionStatusCodeEnd OTA_VehRateRuleRQTransactionStatusCode = "End"
	
	OTA_VehRateRuleRQTransactionStatusCodeRollback OTA_VehRateRuleRQTransactionStatusCode = "Rollback"
	
	OTA_VehRateRuleRQTransactionStatusCodeInSeries OTA_VehRateRuleRQTransactionStatusCode = "InSeries"
	
	OTA_VehRateRuleRQTransactionStatusCodeContinuation OTA_VehRateRuleRQTransactionStatusCode = "Continuation"
	
	OTA_VehRateRuleRQTransactionStatusCodeSubsequent OTA_VehRateRuleRQTransactionStatusCode = "Subsequent"
)

type OTA_VehRateRuleRSTarget string

const (
	OTA_VehRateRuleRSTargetTest OTA_VehRateRuleRSTarget = "Test"
	
	OTA_VehRateRuleRSTargetProduction OTA_VehRateRuleRSTarget = "Production"
)

type OTA_VehRateRuleRSTransactionStatusCode string

const (
	OTA_VehRateRuleRSTransactionStatusCodeStart OTA_VehRateRuleRSTransactionStatusCode = "Start"
	
	OTA_VehRateRuleRSTransactionStatusCodeEnd OTA_VehRateRuleRSTransactionStatusCode = "End"
	
	OTA_VehRateRuleRSTransactionStatusCodeRollback OTA_VehRateRuleRSTransactionStatusCode = "Rollback"
	
	OTA_VehRateRuleRSTransactionStatusCodeInSeries OTA_VehRateRuleRSTransactionStatusCode = "InSeries"
	
	OTA_VehRateRuleRSTransactionStatusCodeContinuation OTA_VehRateRuleRSTransactionStatusCode = "Continuation"
	
	OTA_VehRateRuleRSTransactionStatusCodeSubsequent OTA_VehRateRuleRSTransactionStatusCode = "Subsequent"
)

type VehicleReservationRQAdditionalInfoTypeRentalPaymentPrefType string

const (
	VehicleReservationRQAdditionalInfoTypeRentalPaymentPrefTypeGuarantee VehicleReservationRQAdditionalInfoTypeRentalPaymentPrefType = "guarantee"
	
	VehicleReservationRQAdditionalInfoTypeRentalPaymentPrefTypePayment VehicleReservationRQAdditionalInfoTypeRentalPaymentPrefType = "payment"
)

type InsCoverageDetailTypeType string

const (
	InsCoverageDetailTypeTypeSingleTrip InsCoverageDetailTypeType = "SingleTrip"
	
	InsCoverageDetailTypeTypeMultiTrip InsCoverageDetailTypeType = "MultiTrip"
	
	InsCoverageDetailTypeTypeOther InsCoverageDetailTypeType = "Other"
)

type ValidationErrorCodes string

const (
	ValidationErrorCodesA000 ValidationErrorCodes = "A000"
	
	ValidationErrorCodesA001 ValidationErrorCodes = "A001"
	
	ValidationErrorCodesA002 ValidationErrorCodes = "A002"
	
	ValidationErrorCodesA003 ValidationErrorCodes = "A003"
	
	ValidationErrorCodesA004 ValidationErrorCodes = "A004"
	
	ValidationErrorCodesA005 ValidationErrorCodes = "A005"
	
	ValidationErrorCodesA006 ValidationErrorCodes = "A006"
	
	ValidationErrorCodesA007 ValidationErrorCodes = "A007"
	
	ValidationErrorCodesA008 ValidationErrorCodes = "A008"
	
	ValidationErrorCodesA009 ValidationErrorCodes = "A009"
	
	ValidationErrorCodesA010 ValidationErrorCodes = "A010"
	
	ValidationErrorCodesA011 ValidationErrorCodes = "A011"
	
	ValidationErrorCodesA012 ValidationErrorCodes = "A012"
	
	ValidationErrorCodesA013 ValidationErrorCodes = "A013"
	
	ValidationErrorCodesA014 ValidationErrorCodes = "A014"
	
	ValidationErrorCodesA015 ValidationErrorCodes = "A015"
	
	ValidationErrorCodesA016 ValidationErrorCodes = "A016"
	
	ValidationErrorCodesA017 ValidationErrorCodes = "A017"
	
	ValidationErrorCodesA018 ValidationErrorCodes = "A018"
	
	ValidationErrorCodesA019 ValidationErrorCodes = "A019"
	
	ValidationErrorCodesA020 ValidationErrorCodes = "A020"
	
	ValidationErrorCodesA021 ValidationErrorCodes = "A021"
	
	ValidationErrorCodesA022 ValidationErrorCodes = "A022"
	
	ValidationErrorCodesA023 ValidationErrorCodes = "A023"
	
	ValidationErrorCodesA024 ValidationErrorCodes = "A024"
	
	ValidationErrorCodesA025 ValidationErrorCodes = "A025"
	
	ValidationErrorCodesA026 ValidationErrorCodes = "A026"
	
	ValidationErrorCodesA027 ValidationErrorCodes = "A027"
	
	ValidationErrorCodesUSR000 ValidationErrorCodes = "USR000"
	
	ValidationErrorCodesUSR011 ValidationErrorCodes = "USR011"
	
	ValidationErrorCodesUSR012 ValidationErrorCodes = "USR012"
	
	ValidationErrorCodesUSR013 ValidationErrorCodes = "USR013"
	
	ValidationErrorCodesUSR014 ValidationErrorCodes = "USR014"
	
	ValidationErrorCodesUSR015 ValidationErrorCodes = "USR015"
	
	ValidationErrorCodesUSR016 ValidationErrorCodes = "USR016"
	
	ValidationErrorCodesUSR017 ValidationErrorCodes = "USR017"
	
	ValidationErrorCodesUSR018 ValidationErrorCodes = "USR018"
	
	ValidationErrorCodesUSR019 ValidationErrorCodes = "USR019"
	
	ValidationErrorCodesUSR020 ValidationErrorCodes = "USR020"
	
	ValidationErrorCodesUSR021 ValidationErrorCodes = "USR021"
	
	ValidationErrorCodesUSR022 ValidationErrorCodes = "USR022"
	
	ValidationErrorCodesUSR023 ValidationErrorCodes = "USR023"
	
	ValidationErrorCodesUSR024 ValidationErrorCodes = "USR024"
	
	ValidationErrorCodesUSR025 ValidationErrorCodes = "USR025"
	
	ValidationErrorCodesUSR026 ValidationErrorCodes = "USR026"
	
	ValidationErrorCodesUSR027 ValidationErrorCodes = "USR027"
	
	ValidationErrorCodesUSR028 ValidationErrorCodes = "USR028"
)

type ValidationErrorTypes string

const (
	ValidationErrorTypesNone ValidationErrorTypes = "None"
	
	ValidationErrorTypesValidationError ValidationErrorTypes = "ValidationError"
	
	ValidationErrorTypesAmadeusAPIError ValidationErrorTypes = "AmadeusAPIError"
	
	ValidationErrorTypesEpowerInternalError ValidationErrorTypes = "EpowerInternalError"
	
	ValidationErrorTypesEpowerUnhandledError ValidationErrorTypes = "EpowerUnhandledError"
)

type ValidationErrorCategories string

const (
	ValidationErrorCategoriesSystem ValidationErrorCategories = "System"
	
	ValidationErrorCategoriesAmadeusAPI ValidationErrorCategories = "AmadeusAPI"
	
	ValidationErrorCategoriesFatalError ValidationErrorCategories = "FatalError"
)

type RailProviderTypeEnum string

const (
	RailProviderTypeEnumNone RailProviderTypeEnum = "None"
	
	RailProviderTypeEnumSNCFProvider RailProviderTypeEnum = "SNCFProvider"
)

type DisplayRuleEnum string

const (
	DisplayRuleEnumDISPLAYALLOPTIONS DisplayRuleEnum = "DISPLAYALLOPTIONS"
	
	DisplayRuleEnumDISPLAYONLYINPOLICYOPTIONS DisplayRuleEnum = "DISPLAYONLYINPOLICYOPTIONS"
	
	DisplayRuleEnumDISPLAYONLYOUTPOLICYOPTIONSIFNOINPOLICYOPTION DisplayRuleEnum = "DISPLAYONLYOUTPOLICYOPTIONSIFNOINPOLICYOPTION"
)

type TripApprovalRuleEnum string

const (
	TripApprovalRuleEnumNEVERNEEDSAPPROVAL TripApprovalRuleEnum = "NEVERNEEDSAPPROVAL"
	
	TripApprovalRuleEnumALWAYSNEEDSAPPROVAL TripApprovalRuleEnum = "ALWAYSNEEDSAPPROVAL"
	
	TripApprovalRuleEnumNEEDSAPPROVALWHENTRIPISOUTOFPOLICY TripApprovalRuleEnum = "NEEDSAPPROVALWHENTRIPISOUTOFPOLICY"
)

