package service

import (
	"amadeus-go/config/env"
	"amadeus-go/models"
	"amadeus-go/pkg/repository"
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

// constructing an array type of travelers for performing search request
func travelerInfoConstructor(req *models.SearchRequest) []*repository.TravelerInformationType {
	travelersQuantity := make([]*repository.PassengerTypeQuantityType, 0)
	
	// adding adults
	adults := &repository.PassengerTypeQuantityType{
		Code:     "ADT",
		Quantity: req.AdultCount,
	}
	travelersQuantity = append(travelersQuantity, adults)
	
	// adding children
	if ch := req.ChildCount; ch > 0 {
		children := &repository.PassengerTypeQuantityType{
			Code:     "CHD",
			Quantity: ch,
		}
		travelersQuantity = append(travelersQuantity, children)
	}
	
	// adding infants
	if inf := req.InfantCount; inf > 0 {
		infants := &repository.PassengerTypeQuantityType{
			Code:     "INF",
			Quantity: inf,
		}
		travelersQuantity = append(travelersQuantity, infants)
	}
	
	travelers := []*repository.TravelerInformationType{
		{travelersQuantity},
	}
	
	return travelers
}

// returns the Refund type according to the project configuration
func getRefundableType(refundType string) repository.RefundableTypesEnum {
	switch refundType {
	case "refundable":
		return repository.RefundableTypesEnumOnlyRefundable
	case "nonrefundable":
		return repository.RefundableTypesEnumOnlyNonRefundable
	default:
		return repository.RefundableTypesEnumAllFlights
		
	}
}

// returns the type of search provider according to the project configuration
func getProviderType(provider string) repository.SearchTypeByProviderEnum {
	switch provider {
	case "external":
		return repository.SearchTypeByProviderEnumOnlyExtProviders
	case "amadeus_and_external":
		return repository.SearchTypeByProviderEnumAmadeusAndExtProviders
	default:
		return repository.SearchTypeByProviderEnumOnlyAmadeus
	}
}

// returns the type of booking according to the project configuration
func getTicketType(typ string) repository.TicketType {
	switch typ {
	case "eticket_request":
		return repository.TicketTypeETicketRequest
	case "eticket_instant":
		return repository.TicketTypeETicketInstant
	default:
		return repository.TicketTypeBookingOnly
		
	}
}

// Key , generates a unique key for the requested search, using request params
func key(r *models.SearchRequest) string {
	return fmt.Sprintf("%s%s%s%s%s%d%d%d",
		env.GetStringSlice("service.name"),
		r.Origin, r.Destination,
		r.DepartureTime, r.ReturnTime,
		r.AdultCount, r.ChildCount, r.InfantCount,
	)
}

// cacheTheSearchResult ...
func cacheTheSearchResult(req *models.SearchRequest, data []*models.Data) error {
	d := make([]redis.Z, 0)
	if len(data) == 0 {
		z := redis.Z{
			Score: 0,
			Member: &models.Data{
				ID:            "",
				Type:          env.GetString("service.name"),
				Attributes:    nil,
				Relationships: nil,
			},
		}
		d = append(d, z)
	} else {
		for _, v := range data {
			z := redis.Z{
				Score:  0,
				Member: v,
			}
			d = append(d, z)
		}
	}
	
	if redisErr := repository.Cache().ZAdd(key(req), d...).Err(); redisErr != nil {
		return redisErr
	}
	seconds := env.GetInt("service.cache.expire_search_key")
	var dur = int64(seconds)
	repository.Cache().Expire(key(req), time.Second*time.Duration(dur))
	return nil
}

// cacheTheReserveRequest ...
func cacheTheReserveRequest(key string, value string) error {
	// Create the lead passenger's surname with PNR key for getTicket (Book) method usage
	return repository.Cache().Set(key, value, time.Minute*10).Err()
}

// get title of passenger according to the his/her gender
func getTitle(gender int32) string {
	if gender == 0 {
		return "MR"
	}
	return "MS"
}

// converts the integer convenient gender into the sting one
func getGender(iGender int32) string {
	if iGender == 0 {
		return "Male"
	}
	return "Female"
}

// passengerConstructor ...
func passengerConstructor(req *models.ReserveRequest) []*repository.AirTraveler {
	var passengers = make([]*repository.AirTraveler, 0)
	for _, p := range req.TravelerInfo.AirTravelers {
		passenger := &repository.AirTraveler{
			
			AirTravelerType: &repository.AirTravelerType{
				PassengerTypeCode:getPassengerType(p.PassengerType),
				PersonName: &repository.PersonNameType{
					GivenName: p.PassengerName.PassengerFirstName,
					Surname:   p.PassengerName.PassengerLastName,
					NamePrefix: getTitle(p.Gender),
				},
				Document: []*repository.Document{
					{
						DocumentType: &repository.DocumentType{
							DocID:            getDocID(p),
							DocType:          getDocType(p),
							ExpireDate:       getDocExpireDate(p),
							DocIssueCountry:  getDocIssueCountry(p),
						},
					},
				},
				PassengerTypeQuantity: &repository.PassengerTypeQuantityType{
					Code:     getPassengerType(p.PassengerType),
					Quantity: getQuantity(req, p.PassengerType),
				},
				// Gender:    getGender(p.Gender),
			},
		}
		b,_:=time.Parse("2006-01-02T15:04:05",p.DateOfBirth)
		passenger.BirthDate = b.Format("2006-01-02")
		/*// Set email(s) if provided
		if email := req.TravelerInfo.Email; email != "" {
			passenger.Email = []*repository.Email{
				{
					EmailType: &repository.EmailType{
						EmailType: "1",
						Value:     email,
					},
				},
			}
		}
		
		// Set Phone(s) if provided
		if phone := req.TravelerInfo.PhoneNumber; phone != "" {
			passenger.Telephone = []*repository.Phone{
				{
					PhoneNumber: phone,
				},
			}
		}*/
		passengers = append(passengers, passenger)
	}
	return passengers
}

// get passengers quantity per each PTC
func getQuantity(request *models.ReserveRequest, typ int32) int32 {
	var count int32 = 0
	for _, v := range request.TravelerInfo.AirTravelers {
		if v.PassengerType == typ {
			count++
		}
	}
	
	return count
}

// returns expiration date of document
func getDocExpireDate(traveler *models.AirTraveler) string {
	if traveler.Passport != nil {
		return traveler.Passport.ExpiryDate
	}
	return ""
}

// returns issue country of document
func getDocIssueCountry(traveler *models.AirTraveler) string {
	if traveler.Passport != nil {
		return traveler.Passport.Country
	}
	
	return "Iran"
}

// returns the document type according the passenger info
func getDocType(traveler *models.AirTraveler) string {
	if traveler.NationalId != "" {
		return "NationalID"
	}
	return "Passport"
}

// returns the document id according the passenger info
func getDocID(traveler *models.AirTraveler) string {
	if id := traveler.NationalId; id != "" {
		return id
	}
	return traveler.Passport.PassportNumber
}

// converts integer type of passenger into the string for using in amadeus apis
func getPassengerType(typ int32) string {
	var strType string
	switch typ {
	case 1:
		strType = "ADT"
	case 2:
		strType = "CHD"
	case 3:
		strType = "INF"
	}
	return strType
	// 1 => Adult , 2 => Child  3 => Infant
}

// ID of each item of search result is in colon separated combination of
// RecommendationID and CombinationID
func getBookingIDs(id string) (string, string) {
	ids := strings.Split(id, ":")
	return ids[0], ids[1]
}

func getNationalID(documents []*repository.Document) string {
	for _, d := range documents {
		if strings.Contains(strings.ToLower(d.DocType), "national") {
			return d.DocID
		}
	}
	return ""
}

func getPassportExpireDate(documents []*repository.Document) string {
	for _, d := range documents {
		if strings.Contains(strings.ToLower(d.DocType), "passport") {
			return d.ExpireDate
		}
	}
	return ""
}

func getPassportNumber(documents []*repository.Document) string {
	for _, d := range documents {
		if strings.Contains(strings.ToLower(d.DocType), "passport") {
			return d.DocID
		}
	}
	return ""
}