package main

import (
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

// Weighter evaluates the weight of a route
// based on some criteria
type Weighter func(Route) float64

// ByDistance sets the weight of a route based on the
// distance between two airports.
func ByDistance() Weighter {
	return func(r Route) float64 {
		return GetDistance(r.from, r.to)
	}
}

func Load(airports AirportMap, routes []Route, weighters ...Weighter) graph.Graph {
	g := simple.NewWeightedDirectedGraph(0.0, 0.0)
	for _, airport := range airports {
		g.AddNode(airport)
	}
	for _, route := range routes {
		var values []float64
		for _, weighter := range weighters {
			values = append(values, weighter(route))
		}
		g.SetWeightedEdge(SetWeight(route, floats.Sum(values)))
	}
	return g
}

func Find(start, end Airport, g graph.Graph) Itinerary {
	shortest := path.DijkstraFrom(start, g)
	path, weight := shortest.To(end)
	var airports []Airport
	for _, p := range path {
		airports = append(airports, p.(Airport))
	}
	return Itinerary{stops: airports, weight: weight}
}
