package main

import (
	"fmt"
	"encoding/base64"
)

type location struct {
	Name string
}

type airport struct {
	Iata string
	Location location
	Name string
	Stop bool
}

type airports []airport

type price struct {
	Currency string
	MaxWithTax float64
	MaxWithoutTax float64
	MinWithTax float64
	MinWithoutTax float64
}

type airCompany iataValue

type row struct {
	BestPrice bool
	NumberOfStops int
	Price float64
	PriceWithoutTax float64
}

type column struct {
	AirCompanies []airCompany
	Rows []row
}

type priceMatrix struct {
	Columns []column
	Currency string
}

type maxMinInt struct {
	Maximum int32
	Minimum int32
}

type maxMinString struct {
	Maximum string
	Minimum string
}

type departureDates maxMinString

type route struct {
	AirCompanies []airCompany
	Airports []airport
	BaggageIncluded []bool
	DepartureDates departureDates
	FareType []string
	FlightDuration maxMinInt
	NumberOfStops []int
	rph int
}

type statusExecution struct {
	CompletedExecution bool
	Executed int
	Total int
}

type meta struct {
	CountFlights int
	CountPriceGroups int
	Price price
	Airports airports
	PriceMatrix priceMatrix
	Routes []route
	StatusExecution statusExecution
}

type validatingBy iataValue

type tax struct {
	Code string
	Description string
	Percent float64
	Values []value
}

type fare struct {
	PassengersCount int
	PassengersType string
	PriceWithTax float64
	PriceWithoutTax float64
	Taxes []tax
}

type fareGroup struct {
	Currency string
	priceWithTax float64
	priceWithoutTax float64
	reCharging bool
	Fares []fare
}

type baggage struct {
	IsIncluded bool
	Quantity float64
	Type string
	Uom string
	Weight float64
}

type service struct {
	Description string
	IsIncluded bool
	Type string
}

type fareProfile struct {
	Baggage baggage
	FareFamily string
	Services []service
}

type leg struct {
	AircraftCode string
	Arrival string
	ArrivalDate string
	Departure string
	DepartureDate string
	Duration int
	FareBasis string
	FareClass string
	FlightNumber int
	ManagedBy iataValue
	NumberOfStops int
	OperatedBy iataValue
	SeatClass descriptionValue
	Status string
//	"stops": []
}

type segment struct {
	Arrival string
	ArrivalDate string
	Departure string
	DepartureDate string
	Duration int32
	FareProfile fareProfile
	Legs []leg
	NumberOfStops int
	PackageGroup string
	RateToken string
	RouteRPH int
	Rph int
}

type flight struct {
	FareGroup fareGroup
	Segments []segment
	ValidatingBy validatingBy
}

type flights []flight

type flightsResponse struct {
	Flights flights
	Meta meta
}






func (f flightsResponse) printResume() {

	//	fmt.Println("Source: ", f.)
	}
	
	func (f flightsResponse) printRecommendations() {
		for _, r := range f.Flights {
			rateToken := r.Segments[0].RateToken
			decoded, err := base64.URLEncoding.DecodeString(rateToken)
			isValid(err)
			rateTokenDecoded := string(decoded)
			fmt.Println(rateTokenDecoded)
		}
	}
	
	func (f flightsResponse) printMeta() {
		fmt.Println("-------------------------------------")
	
		for _, column := range f.Meta.PriceMatrix.Columns {
			fmt.Print(column.AirCompanies[0].Iata, " | ")
			for _, row := range column.Rows {
				fmt.Print(" Stops(", row.NumberOfStops, ") ", row.PriceWithoutTax, " -> ", row.Price, " | ")
			}
			fmt.Println()
		}
		
		
	}