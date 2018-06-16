package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"sort"

	"gonum.org/v1/gonum/graph"
)

// Map of route keys e.g. ORDJFK and
// slice of weights to apply
type WeightMap map[string][]float64

type AirportMap map[string]Airport

func (a AirportMap) Airport(id string) Airport {
	return a[id]
}

func (a AirportMap) HasCode(code string) bool {
	_, ok := a[code]
	return ok
}

func (a AirportMap) Ordered() []Airport {
	var (
		sorted []Airport
		codes  []string
	)
	for key, _ := range a {
		codes = append(codes, key)
	}
	sort.Strings(codes)
	for _, code := range codes {
		sorted = append(sorted, a[code])
	}
	return sorted
}

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
	return fmt.Sprintf(
		"%s:%s:%s",
		a.Code, a.City, a.Country,
	)
}

func (a Airport) DOTID() string { return fmt.Sprintf("\"%s\"", a.Code) }

func (a Airport) ID() int64 {
	h := fnv.New64a()
	h.Write([]byte(a.String()))
	return int64(h.Sum64())
}

type Route struct {
	to     Airport
	from   Airport
	weight float64
}

func (r Route) String() string {
	return fmt.Sprintf("%s->%s", r.from, r.to)
}

func (r Route) To() graph.Node   { return r.to }
func (r Route) From() graph.Node { return r.from }
func (r Route) Weight() float64  { return r.weight }

func SetWeight(r Route, weight float64) Route {
	return Route{to: r.to, from: r.from, weight: weight}
}

type Itinerary struct {
	// Ordered list of airports
	stops  []Airport
	weight float64
}

func (i Itinerary) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("%f:", i.weight))
	for j, stop := range i.stops {
		buf.WriteString(stop.String())
		if j+1 != len(i.stops) {
			buf.WriteString("-->")
		}
	}
	return buf.String()
}
