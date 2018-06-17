package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("flights", " ✈️  Find The Best Possible Flight ✈️  ")

	app.Command("airports", "list all possible airports", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			airports := Airports()
			for _, airport := range airports.Ordered() {
				fmt.Println(airport)
			}
		}
	})

	app.Command("routes", "list all possible routes", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			for _, route := range Routes(Airports()) {
				fmt.Println(route)
			}
		}
	})

	app.Command("furthest", "show the furthest distance you can travel", func(cmd *cli.Cmd) {
		var (
			departureCode = cmd.StringArg("DEPARTURE", "", "starting airport")
		)
		cmd.Action = func() {
			airports := Airports()
			start := airports.Airport(*departureCode)
			end := start
			for _, other := range airports {
				if GetDistance(start, other) > GetDistance(start, end) {
					end = other
				}
			}
			fmt.Println(Itinerary{weight: GetDistance(start, end), stops: []Airport{start, end}})
		}
	})

	app.Command("nearby", "show airports nearby", func(cmd *cli.Cmd) {
		var (
			threshold     = cmd.IntOpt("t threshold", 500, "distance from starting airport")
			departureCode = cmd.StringArg("DEPARTURE", "", "starting airport")
		)
		cmd.Action = func() {
			airports := Airports()
			start := airports.Airport(*departureCode)
			for _, end := range NearBy(float64(*threshold), start, airports) {
				fmt.Println(Itinerary{weight: GetDistance(start, end), stops: []Airport{start, end}})
			}
		}
	})

	app.Command("route", "find the best possible routes", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS] DEPARTURE ARRIVAL"
		var (
			//threshold     = cmd.IntOpt("threshold", 100, "airport distance threshold")
			departureCode = cmd.StringArg("DEPARTURE", "", "starting airport")
			arrivalCode   = cmd.StringArg("ARRIVAL", "", "ending airport")
		)
		cmd.Action = func() {
			airports := Airports()
			//shortest, weight := Load(airports.Airport(*departureCode), airports.Airport(*arrivalCode), opts)
			fmt.Println(Find(
				airports.Airport(*departureCode),
				airports.Airport(*arrivalCode),
				Load(airports, Routes(airports), ByDistance()),
			))
			//fmt.Println(shortest, weight)
			//maybe(NewEncoder(Json, Load(airports, Routes(airports))).Encode(os.Stdout))
			/*
				departures, arrivals := FindRoutes(float64(*threshold), airports, airports[*arrivalCode], airports[*departureCode])
				if *asJson {
					data := [][]Airport{departures, arrivals}
					maybe(json.NewEncoder(os.Stdout).Encode(data))
				} else {
					fmt.Println("Departures: ")
					for _, departure := range departures {
						fmt.Println(departure)
					}
					fmt.Println("Arrivals: ")
					for _, arrival := range arrivals {
						fmt.Println(arrival)
					}
				}
			*/
		}
	})
	maybe(app.Run(os.Args))
}
