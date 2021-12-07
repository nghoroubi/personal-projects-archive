package helpers

import (
	"amadeus-go/config/env"
	"strings"
)

// GetClassByLetter , returns the type of fare class by assigned letter
func GetClassByLetter(l string) string {
	// all classes map
	classes := make(map[string][]string)
	
	economy := env.GetStringSlice("flight_classes.economy")
	premium := env.GetStringSlice("flight_classes.premium")
	business := env.GetStringSlice("flight_classes.business")
	first := env.GetStringSlice("flight_classes.first")
	
	classes["economy"] = economy
	classes["premium"] = premium
	classes["business"] = business
	classes["first"] = first
	
	for name, class := range classes {
		for _, letter := range class {
			if strings.EqualFold(strings.ToLower(l), strings.ToLower(letter)) {
				return name
			}
		}
	}
	return ""
}
