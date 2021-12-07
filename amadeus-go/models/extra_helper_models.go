package models

import "amadeus-go/pkg/repository"

// CombinedOriginDestinations , a set of one or more originDestinationOption
// plus combinationId of them for book process usage
type CombinedOriginDestinations struct {
	OriginDestinationOptions []*repository.OriginDestinationOption
	CombinationId            int32
}
