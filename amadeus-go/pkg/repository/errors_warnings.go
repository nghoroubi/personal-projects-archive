package repository

// ErrorsType ...
type ErrorsType struct {
	Error      []*ErrorType `xml:"Error,omitempty"`
	MyProperty string       `xml:"MyProperty,omitempty"`
}

// ErrorType ...
type ErrorType struct {
	*FreeTextType
	Type      string `xml:"Type,attr,omitempty"`
	ErrorCode string `xml:"ErrorCode,attr,omitempty"`
	ShortText string `xml:"ShortText,attr,omitempty"`
	Code      string `xml:"Code,attr,omitempty"`
	DocURL    string `xml:"DocURL,attr,omitempty"`
	Status    string `xml:"Status,attr,omitempty"`
	Tag       string `xml:"Tag,attr,omitempty"`
	RecordID  string `xml:"RecordID,attr,omitempty"`
	NodeList  string `xml:"NodeList,attr,omitempty"`
	BreakFlow bool   `xml:"BreakFlow,attr,omitempty"`
}

// FreeTextType ...
type FreeTextType struct {
	Value    string
	Language *Language `xml:"Language,attr,omitempty"`
}

// WarningType ...
type WarningType struct {
	*FreeTextType
	Type      string `xml:"Type,attr,omitempty"`
	ShortText string `xml:"ShortText,attr,omitempty"`
	Code      string `xml:"Code,attr,omitempty"`
	DocURL    string `xml:"DocURL,attr,omitempty"`
	Status    string `xml:"Status,attr,omitempty"`
	Tag       string `xml:"Tag,attr,omitempty"`
	RecordID  string `xml:"RecordID,attr,omitempty"`
	RPH       string `xml:"RPH,attr,omitempty"`
}

// SuccessType ...
type SuccessType struct {
}

// WarningsType ...
type WarningsType struct {
	Warning []*WarningType `xml:"Warning,omitempty"`
}
