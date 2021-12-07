package models

// AirRulesRequest ...
type AirRulesRequest struct {
	ServiceId      int32  `validate:"required" json:"service_id"`
	SessionId      string `validate:"required" json:"session_id"`
	FareSourceCode string `validate:"required" json:"fare_source_code"`
	UniqueId       string `validate:"required" json:"unique_id"`
}

// AirRulesResponse ...
type AirRulesResponse struct {
	StdError []*StdError   `json:"errors"`
	Data     *AirRulesData `json:"data"`
	Meta     *Meta         `json:"meta"`
	Links    *Link         `json:"links"`
}

// AirRulesData ...
type AirRulesData struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    *AirRulesAttributes    `json:"attributes"`
	Relationships *AirRulesRelationships `json:"relationships"`
}

// AirRulesAttributes ...
type AirRulesAttributes struct {
	Success  bool  `json:"success"`
	FareType int32 `json:"fare_type"`
}

// AirRulesRelationships ...
type AirRulesRelationships struct {
	FareRules []*FareRules
}

// FareRules ...
type FareRules struct {
	ID            string                  `json:"id"`
	Type          string                  `json:"type"`
	Attributes    *FareRulesAttributes    `json:"attributes"`
	Relationships *FareRulesRelationships `json:"relationships"`
}

// FareRulesAttributes ...
type FareRulesAttributes struct {
	Airline  string `json:"airline"`
	CityPair string `json:"city_pair"`
}

// FareRulesRelationships ...
type FareRulesRelationships struct {
	RuleDetails []*RuleDetails `json:"rule_details"`
}

// RuleDetails ...
type RuleDetails struct {
	Category string `json:"category"`
	Rules    string `json:"rules"`
}
