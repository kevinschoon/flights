package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const d2r = (math.Pi / 180.0)

// Haversine calculates linear distance between two pairs
func Haversine(start, dest [2]float64) float64 {
	dlong := (dest[1] - start[1]) * d2r
	dlat := (dest[0] - start[0]) * d2r
	a := math.Pow(math.Sin(dlat/2.0), 2) + math.Cos(start[0]*d2r)*math.Cos(dest[0]*d2r)*math.Pow(math.Sin(dlong/2.0), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := 6367 * c
	return distance
}

func GetDistance(start, dest Airport) float64 {
	return Haversine([2]float64{start.Latitude, start.Longitude}, [2]float64{dest.Latitude, dest.Longitude})
}

func NearBy(threshold float64, start Airport, airports AirportMap) []Airport {
	nearby := []Airport{}
	for _, dest := range airports {
		if GetDistance(start, dest) <= threshold {
			if dest.Code != start.Code {
				nearby = append(nearby, dest)
			}
		}
	}
	return nearby
}

func FindRoutes(threshold float64, airports map[string]Airport, departure, arrival Airport) ([]Airport, []Airport) {
	return NearBy(threshold, departure, airports), NearBy(threshold, arrival, airports)
}

// icao_code,iata_code,name,city,country,lat_deg,lat_min,lat_sec,lat_dir,lon_deg,lon_min,lon_sec,lon_dir,altitude,lat_decimal,lon_decimal
// AYGA:GKA:GOROKA:GOROKA:PAPUA NEW GUINEA:006:004:054:S:145:023:030:E:01610:-6.082:145.392
func parseAirport(row []string) *Airport {
	if len(row) != 16 {
		return nil
	}
	if row[1] == "N/A" {
		return nil
	}
	latitude, _ := strconv.ParseFloat(row[14], 64)
	longitude, _ := strconv.ParseFloat(row[15], 64)
	if latitude == 0 || longitude == 0 {
		return nil
	}
	airport := &Airport{
		Code:      row[1],
		Name:      row[2],
		City:      row[3],
		Country:   row[4],
		Latitude:  latitude,
		Longitude: longitude,
	}
	return airport
}

//2B,410,AER,2965,KZN,2990,,0,CR2
func parseRoute(airports AirportMap, row []string) *Route {
	if len(row) != 9 {
		return nil
	}
	start, end := row[2], row[4]
	if airports.HasCode(start) && airports.HasCode(end) {
		if start != end {
			return &Route{
				to:   airports.Airport(end),
				from: airports.Airport(start),
			}
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func Airports() AirportMap {
	airports := map[string]Airport{}
	reader := csv.NewReader(strings.NewReader(airportData))
	reader.Comma = ':'
	for {
		row, err := reader.Read()
		if err == io.EOF {
			return airports
		}
		airport := parseAirport(row)
		if airport != nil {
			airports[airport.Code] = *airport
		}
	}
}

func Routes(airports AirportMap) []Route {
	var routes []Route
	reader := csv.NewReader(strings.NewReader(routeData))
	reader.Comma = ','
	for {
		row, err := reader.Read()
		if err == io.EOF {
			return routes
		}
		route := parseRoute(airports, row)
		if route != nil {
			routes = append(routes, *route)
		}
	}
}

func maybe(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
