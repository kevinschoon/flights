package main

import (
	"fmt"
)

type Airport struct {
	// ICAO code
	Code      string
	Name      string
	City      string
	Country   string
	Latitude  float64
	Longitude float64
}

func (a Airport) String() string {
	return fmt.Sprintf("%s:%s:%s:%s:%f:%f", a.Code, a.Name, a.City, a.Country, a.Latitude, a.Longitude)
}
