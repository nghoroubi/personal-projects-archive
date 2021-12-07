package helpers

import (
	"amadeus-go/config/env"
	"amadeus-go/models"
	"amadeus-go/pkg/repository"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// TIME , format of timestamp for parsing time strings
const TIME = "2006-01-02T15:04:05"

// ParseSearchResponse ...
func ParseSearchResponse(pricedItinerary *repository.PricedItinerary, items []*repository.OriginDestinationOption, pricedItin *repository.PricedItinerariesType, combinationId int32) *models.Data {
	var (
		services        = make([]*models.Service, 0)
		offerItems      = make([]*models.OfferItems, 0)
		priceInfoSource *repository.AirItineraryPricingInfoType
	)
	
	priceInfoSource = pricedItinerary.AirItineraryPricingInfo
	
	// Calculate and assign fares
	totalPrice := strconv.FormatFloat(priceInfoSource.ItinTotalFare.TotalFare.Amount, 'f', -1, 64)
	totalTax := priceInfoSource.ItinTotalFare.TotalFare.Amount - priceInfoSource.ItinTotalFare.BaseFare.Amount
	totalTaxStr := strconv.FormatFloat(totalTax, 'f', -1, 64)
	adultTotalFare := getTotalFare(priceInfoSource.PTCFareBreakdowns.PTCFareBreakdown, "ADT")
	adultTotalTax := getTotalTax(priceInfoSource.PTCFareBreakdowns.PTCFareBreakdown, "ADT")
	childTotalFare := getTotalFare(priceInfoSource.PTCFareBreakdowns.PTCFareBreakdown, "CHD")
	childTotalTax := getTotalTax(priceInfoSource.PTCFareBreakdowns.PTCFareBreakdown, "CHD")
	infantTotalFare := getTotalFare(priceInfoSource.PTCFareBreakdowns.PTCFareBreakdown, "INF")
	infantTotalTax := getTotalTax(priceInfoSource.PTCFareBreakdowns.PTCFareBreakdown, "INF")
	
	for _, item := range items {
		service := &models.Service{
			Attributes:           &models.ServiceAttributes{},
			ServiceRelationships: &models.ServiceRelationships{},
		}
		segments := make([]*models.Segment, 0)
		// parsing response to native response type
		for _, v := range item.FlightSegment {
			technicalStops := make([]*models.TechnicalStops, 0)
			// Calc technical stops
			for _, stop := range v.StopLocation {
				s := &models.TechnicalStops{Data: &models.TechnicalStopsData{
					ArrivalAirport:    stop.LocationCode,
					ArrivalDateTime:   stop.ArrivalDateTime,
					DepartureDateTime: stop.DepartureDateTime,
				}}
				technicalStops = append(technicalStops, s)
			}
			
			flightSegment := &models.Segment{
				ID:   "",
				Type: env.GetString("service.name"),
				Attributes: &models.SegmentAtt{
					TravelClass:    v.BookingClassAvails.BookingClassAvail[0].ResBookDesigCabinCode,
					FareClass:      GetClassByLetter(v.BookingClassAvails.BookingClassAvail[0].ResBookDesigCode),
					Availability:   int32(v.BookingClassAvails.BookingClassAvail[0].ResBookDesigQuantity),
					FareBasis:      v.BookingClassAvails.BookingClassAvail[0].FareBasis,
					Duration:       Duration(v),
					DurationPerMin: DurationInMinutes(v),
					Arrival: &models.DateDetail{
						IataCode: v.ArrivalAirport.LocationCode,
						At:       v.ArrivalDateTime,
					},
					Departure: &models.DateDetail{
						IataCode: v.DepartureAirport.LocationCode,
						At:       v.DepartureDateTime,
					},
					CarrierCode:  v.MarketingAirline.Code,
					Number:       v.FlightNumber,
					AircraftCode: v.Equipment[0].AirEquipType,
					Baggage:      getBaggage(v.Baggages.Baggage[0].Index, pricedItin),
					IsCharter:    false,
				},
				Relationships: &models.SegmentRelationships{
					Operating: &models.Operating{
						CarrierCode: v.OperatingAirline.Code,
						Number:      v.OperatingAirline.FlightNumber,
					},
					TechnicalStops: technicalStops,
				},
			}
			segments = append(segments, flightSegment)
			service.Attributes.Duration, service.Attributes.DurationPerMin = CalcDuration(item.ElapsedTime)
			service.ServiceRelationships.Segment = segments
			
		}
		services = append(services, service)
	}
	offerItems = []*models.OfferItems{
		{
			Relationships: &models.OfferItemRelationship{
				Price: &models.Price{
					Total:    totalPrice,
					TotalTax: totalTaxStr,
				},
				PricePerAdult: &models.Price{
					Total:    adultTotalFare,
					TotalTax: adultTotalTax,
				},
				PricePerChild: &models.Price{
					Total:    childTotalFare,
					TotalTax: childTotalTax,
				},
				PricePerInfant: &models.Price{
					Total:    infantTotalFare,
					TotalTax: infantTotalTax,
				},
				Service: services,
			},
		},
	}
	
	d := &models.Data{
		ID:   fmt.Sprintf("%s:%d", pricedItinerary.SequenceNumber, combinationId),
		Type: env.GetString("service.name"),
		Attributes: &models.Attributes{
			Currency: env.GetString("api.currency"),
			Stop:     false,
			Date: &models.Date{
				DepartureDate: services[0].ServiceRelationships.Segment[0].Attributes.Departure.At,
				ArrivalDate:   services[0].ServiceRelationships.Segment[0].Attributes.Arrival.At,
			},
		},
		Relationships: &models.Relationships{
			OfferItems: offerItems,
			Links:      nil,
		},
	}
	return d
}

// extracts the free baggage info from search result's Baggage section with the baggage index of flight
func getBaggage(index int32, itineraries *repository.PricedItinerariesType) string {
	var b string
	for _, baggage := range itineraries.FreeBaggages.Baggage {
		if baggage.Index == index {
			b = fmt.Sprintf("%s%s", baggage.Quantity, baggage.Unit)
			break
		}
	}
	return b
}

// GetTotalTax , calculates the total amount of tax per passenger type
func getTotalTax(taxes []*repository.PTCFareBreakdownType, passengerType string) string {
	var totalTax float64 = 0
	
	if passengerType != "" {
		for _, tax := range taxes {
			if tax.PassengerTypeQuantity.Code == passengerType {
				for _, taxType := range tax.PassengerFare.Taxes.Tax {
					totalTax += taxType.Amount
				}
			}
		}
		
		return strconv.FormatFloat(totalTax, 'f', -1, 64)
	}
	
	for _, tax := range taxes {
		for _, taxType := range tax.PassengerFare.Taxes.Tax {
			totalTax += taxType.Amount
		}
	}
	
	return strconv.FormatFloat(totalTax, 'f', -1, 64)
}

// GetTotalFare ...
// Calculate total fare per passenger type and return it as string
// Get passenger type as argument (e.g "ADT")
func getTotalFare(passengers []*repository.PTCFareBreakdownType, passengerType string) string {
	var totalFare float64 = 0
	
	for _, passenger := range passengers {
		if passenger.PassengerTypeQuantity.Code == passengerType {
			totalFare = passenger.PassengerFare.TotalFare.Amount
		}
	}
	
	return strconv.FormatFloat(totalFare, 'f', -1, 64)
}

// TimeStandardize , formats the time  in "2006-01-02 15:04" format into "2006-01-02T15:04:05"
func TimeStandardize(t string) string {
	tm, _ := time.Parse("2006-01-02 15:04", t)
	return tm.Format(TIME)
}

// Duration , calculates the duration of a flight in HH:MM format
func Duration(flight *repository.FlightSegment) string {
	arrTime, err := time.Parse(TIME, flight.ArrivalDateTime)
	if err != nil {
		return ""
	}
	depTime, err := time.Parse(TIME, flight.DepartureDateTime)
	if err != nil {
		return ""
	}
	
	hourDelta := arrTime.Hour() - depTime.Hour()
	if hourDelta < 0 {
		hourDelta = 24 - int(math.Abs(float64(hourDelta)))
	}
	
	minDelta := arrTime.Minute() - depTime.Minute()
	if minDelta < 0 {
		minDelta = 60 - int(math.Abs(float64(minDelta)))
		hourDelta -= 1
	}
	return fmt.Sprintf("%d:%d", hourDelta, minDelta)
	
}

// DurationInMinutes , calculates the duration of a flight in minutes
func DurationInMinutes(flight *repository.FlightSegment) int32 {
	arrTime, err := time.Parse(TIME, flight.ArrivalDateTime)
	if err != nil {
		return 0
	}
	depTime, err := time.Parse(TIME, flight.DepartureDateTime)
	if err != nil {
		return 0
	}
	
	hourDelta := arrTime.Hour() - depTime.Hour()
	if hourDelta < 0 {
		hourDelta = 24 - int(math.Abs(float64(hourDelta)))
	}
	
	minDelta := arrTime.Minute() - depTime.Minute()
	if minDelta < 0 {
		minDelta = 60 - int(math.Abs(float64(minDelta)))
		hourDelta -= 1
	}
	return int32(hourDelta*60 + minDelta)
	
}

// GetPassengerType , gets the type of passenger in number mode.
// Adult -> 1 , Child -> 2 , Infant -> 3
func GetPassengerType(birthdate string) int32 {
	c := GetAgeCategory(birthdate)
	if c == "ADL" {
		return 01
	} else if c == "CHD" {
		return 2
	} else {
		return 3
	}
}

// GetAgeCategory ...
func GetAgeCategory(birthDate string) string {
	age := getAge(birthDate)
	var cat = "ADL"
	if age <= 2 {
		cat = "INF"
	}
	if age > 2 && age <= 12 {
		cat = "CHD"
	}
	
	return cat
}

// GetAge , calculates age using birthdate and current date
func getAge(birthDate string) int32 {
	var (
		birth     time.Time
		err       error
		birthYear int
	)
	thisYear := time.Now().Year()
	birth, err = time.Parse("2006-01-02T15:04:05", birthDate)
	if err != nil {
		return 0
	}
	birthYear = birth.Year()
	
	return int32(thisYear - birthYear)
}

// GetTitle ,
//  Adult - Male=>0 , -Adult - FeMale =>1 , Adule - Female=>2 ,Child,Infant - Female =>4 , Child,Infant - Male =>5
func GetTitle(ageCat string, gender string) int32 {
	if strings.ToLower(ageCat) == "adult" {
		if strings.ToLower(gender) == "male" {
			return 0
		} else {
			return 1
		}
	} else {
		if strings.ToLower(gender) == "female" {
			return 5
		} else {
			return 4
		}
	}
}

// SeparateODOByDirection , this method takes an array of OriginDestinationOptions
// and separates them into a map  of originDestinationOptions by them direction id
func SeparateODOByDirection(ods []*repository.OriginDestinationOption) map[int32][]*repository.OriginDestinationOption {
	originDestinationOptions := make(map[int32][]*repository.OriginDestinationOption)
	for _, od := range ods {
		if od.DirectionId == 0 {
			originDestinationOptions[0] = append(originDestinationOptions[0], od)
		} else if od.DirectionId == 1 {
			originDestinationOptions[1] = append(originDestinationOptions[1], od)
		}
		
	}
	return originDestinationOptions
}

// DetectCombinableOriginDestinations , this function
// ODO stands for  OriginDestinationOptions
// ODC stands for OriginDestinationCombination
func DetectCombinableOriginDestinations(ODOsMap map[int32][]*repository.OriginDestinationOption, ODCs []*repository.OriginDestinationCombinationType) []*models.CombinedOriginDestinations {
	var result = make([]*models.CombinedOriginDestinations, 0)
	for _, odc := range ODCs {
		combined := &models.CombinedOriginDestinations{}
		originDestinationOptions := make([]*repository.OriginDestinationOption, 0)
		// assign combinationId
		combined.CombinationId = odc.CombinationID
		
		indices, _ := SeparateStringToIntArray(odc.IndexList, ";")
		for directionId, refNumber := range indices {
			ODOs, _ := ODOsMap[int32(directionId)]                        // An array with direction i id
			ODO := getOriginDestinationOptionByRefNumber(ODOs, refNumber) // returns one ODO
			originDestinationOptions = append(originDestinationOptions, ODO)
		}
		
		combined.OriginDestinationOptions = originDestinationOptions
		result = append(result, combined)
	}
	return result
}

// returns an ODO from array, which refNum is equal to passed param
func getOriginDestinationOptionByRefNumber(originDestinationOptions []*repository.OriginDestinationOption, refNumber int32) *repository.OriginDestinationOption {
	for _, ODO := range originDestinationOptions {
		if ODO.RefNumber == refNumber {
			return ODO
		}
	}
	return nil
}

// SeparateStringToIntArray , takes a semiColon separated string an returns an array of integer indices
func SeparateStringToIntArray(main string, splitter string) ([]int32, error) {
	result := make([]int32, 0)
	parts := strings.Split(main, splitter)
	for _, v := range parts {
		i, err := strconv.Atoi(v)
		if err != nil {
			return result, err
		}
		
		result = append(result, int32(i))
	}
	return result, nil
}

// CalcDuration , returns the duration of trip elapsed time as HH:MM format
// and , in minutes as an integer respectively
func CalcDuration(elapsed string) (string, int32) {
	x := elapsed
	H := x[:2]
	h, _ := strconv.Atoi(H)
	
	M := x[2:]
	m, _ := strconv.Atoi(M)
	return fmt.Sprintf("%s:%s", H, M), int32(h*60 + m)
}
