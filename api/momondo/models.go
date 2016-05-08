package momondo

import "time"

type FlightDate time.Time
type LegDate time.Time

type SearchRequest struct {
	AdultCount  int
	Application string
	ChildAges   []int
	Culture     string
	Mobile      bool
	TicketClass string
	Segments    []Direction
}

type Direction struct {
	Origin      string
	Destination string
	Departure   FlightDate
}

type SearchSessionInfo struct {
	AdultCount	int
	ChildAges	[]int
	ChildCount	int
	Culture		string
	EngineId	int
	InfantCount	int
	SearchId	string
}

type SearchResult struct {
	Airlines      *[]Airline
	Airports      *[]Airport
	Done          bool
	EngineId      int
	Error         bool
	ErrorMessage  string
	Fees          *[]Fee
	Flights       *[]Flight
	Legs          *[]Leg
	Offers        *[]Offer
	ResultNumber  int
	SearchId      string
	Segments      *[]Segment
	Suppliers     *[]Supplier
	TicketClasses *[]TicketClass
}

type Airline struct {
	IATACode string `json:"iata"`
	ICAOCode string `json:"icao"`
	Name     string `json:"name"`
}

type Airport struct {
	CountryCode  string
	CountryName  string
	IATACode     string `json:"iata"`
	ICAOCode     string `json:"icao"`
	MainCityCode string
	MainCityName string
	Name         string
	TimeZone     int
}

type Fee struct {
	AirlineIndex int
	Description  string
	MaxAmountEUR float32
	MinAmountEUR float32
	PaymentId    int
	Type         string
}

type Flight struct {
	DirectAirlineIndex int
	Key                string
	SegmentIndexes     []int
	TicketClassIndex   int
}

type Leg struct {
	AirlineIndex        int
	ArrivalDateRaw      LegDate `json:"Arrival"`
	DepartureDateRaw    LegDate `json:"Departure"`
	DestinationIndex    int
	OriginIndex         int
	Duration            int
	FlightNumber        int
	Key                 string
	StopOverCodeIndexes []int
	StopOvers           int
}

type Offer struct {
	AdultPrice        float64
	AdultPriceEUR     float64
	AdultPriceExclTax float64
	Currency          string
	DeepLink	  string
	FeeIndexes        []int
	FlightIndex       int
	TicketClassIndex  int
	TotalIsCalculated bool
	TotalPrice        float64
	TotalPriceEUR     float64
	TotalPriceExclTax float64
}

type Segment struct {
	Duration   int
	Key        string
	LegIndexes []int
}

type Supplier struct {
	DisplayName  string
	OfferIndexes []int
	SupplierId   string
}

type TicketClass struct {
	Code string
	Name string
}

func (d FlightDate) MarshalJSON() ([]byte, error) {
	date := time.Time(d).Format("\"2006-01-02\"")
	return []byte(date), nil
}

func (date *LegDate) UnmarshalJSON(b []byte) error {
	if (b[0] == '"' && b[len(b) - 1] == '"') {
		b = b[1 : len(b) - 1]
	}

	d, err := time.Parse("2006-01-02T15:04:05", string(b))
	if (err == nil) {
		*date = LegDate(d)
	}

	return err
}

func (leg Leg) DepartureDate() time.Time {
	return time.Time(leg.DepartureDateRaw)
}

func (leg Leg) ArrivalDate() time.Time {
	return time.Time(leg.ArrivalDateRaw)
}