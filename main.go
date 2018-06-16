package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("flights", "✈️ ✈️ ✈️ ✈️ ✈️ ")

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

	app.Command("route", "find the best possible routes", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS] DEPARTURE ARRIVAL"
		var (
			//threshold     = cmd.IntOpt("threshold", 100, "airport distance threshold")
			departureCode = cmd.StringArg("DEPARTURE", "", "starting airport")
			arrivalCode   = cmd.StringArg("ARRIVAL", "", "ending airport")
		)
		cmd.Action = func() {
			airports := Airports()
			if !airports.HasCode(*arrivalCode) {
				maybe(fmt.Errorf("bad arrival code: %s", *arrivalCode))
			}
			if !airports.HasCode(*departureCode) {
				maybe(fmt.Errorf("bad departure code: %s", *departureCode))
			}
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
